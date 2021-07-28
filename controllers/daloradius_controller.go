/*
Copyright 2021 Harris<i@xuegaogg.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	appv1 "github.com/XUEGAONET/daloradius-k8s-controller/api/v1"
	"github.com/XUEGAONET/daloradius-k8s-controller/pkg/deployment"
	"github.com/XUEGAONET/daloradius-k8s-controller/pkg/metadata"
	"github.com/XUEGAONET/daloradius-k8s-controller/pkg/secret"
	"github.com/XUEGAONET/daloradius-k8s-controller/pkg/service"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// DaloRadiusReconciler reconciles a DaloRadius object
type DaloRadiusReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.k8s-cluster.net.xuegaogg.com,resources=daloradius,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.k8s-cluster.net.xuegaogg.com,resources=daloradius/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.k8s-cluster.net.xuegaogg.com,resources=daloradius/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DaloRadius object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DaloRadiusReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	dr := &appv1.DaloRadius{}
	err := r.Get(ctx, req.NamespacedName, dr)
	if err != nil {
		if k8serr.IsNotFound(err) {
			return ctrl.Result{}, nil
		} else {
			return ctrl.Result{}, err
		}
	}

	logger.Info(fmt.Sprintf("DaloRadius name: %s, namespace: %s", dr.Name, dr.Namespace))

	var (
		sr  *corev1.Secret     = &corev1.Secret{}
		dep *appsv1.Deployment = &appsv1.Deployment{}
		svc *corev1.Service    = &corev1.Service{}
	)

	nsn := types.NamespacedName{
		Namespace: req.Namespace,
		Name:      "daloradius",
	}

	var needCreate bool = false
	if dr.Status.Status != "OK" {
		sr = secret.Create()
		dep = deployment.Create()
		svc = service.Create()
		needCreate = true
	} else {
		err = r.Get(ctx, nsn, sr)
		if err != nil {
			return ctrl.Result{}, err
		}

		err = r.Get(ctx, nsn, dep)
		if err != nil {
			return ctrl.Result{}, err
		}

		err = r.Get(ctx, nsn, svc)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	secret.PatchUsername(sr, dr.Spec.MysqlUsername)
	secret.PatchPassword(sr, dr.Spec.MysqlPassword)
	metadata.PatchNamespace(sr, req.Namespace)
	err = r.CreateOrUpdate(ctx, needCreate, sr)
	if err != nil {
		logger.Error(err, "create/update secret failed")
		return ctrl.Result{}, err
	}

	deployment.PatchTimezone(dep, dr.Spec.Timezone)
	deployment.PatchDatabaseHost(dep, dr.Spec.MysqlHost)
	deployment.PatchDatabasePort(dep, dr.Spec.MysqlPort)
	deployment.PatchDatabaseName(dep, dr.Spec.MysqlDatabase)
	deployment.PatchReplicas(dep, dr.Spec.Replicas)
	deployment.PatchImageTag(dep, dr.Spec.ImageTag)
	metadata.PatchNamespace(dep, req.Namespace)
	err = r.CreateOrUpdate(ctx, needCreate, dep)
	if err != nil {
		logger.Error(err, "create/update deployment failed")
		return ctrl.Result{}, err
	}

	service.PatchHttpPort(svc, dr.Spec.HttpPort)
	service.PatchAuthPort(svc, dr.Spec.AuthPort)
	service.PatchAcctPort(svc, dr.Spec.AcctPort)
	metadata.PatchNamespace(svc, req.Namespace)
	err = r.CreateOrUpdate(ctx, needCreate, svc)
	if err != nil {
		logger.Error(err, "create/update service failed")
		return ctrl.Result{}, err
	}

	dr.Status.Status = "OK"
	err = r.Status().Update(ctx, dr)
	if err != nil {
		logger.Error(err, "update status failed")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *DaloRadiusReconciler) CreateOrUpdate(ctx context.Context, needCreate bool, obj client.Object) error {
	logger := log.FromContext(ctx)

	var err error = nil
	if needCreate {
		err = r.Create(ctx, obj)
		logger.Info(fmt.Sprintf("created"))
	} else {
		err = r.Update(ctx, obj)
		logger.Info(fmt.Sprintf("updated"))
	}
	return err
}

// SetupWithManager sets up the controller with the Manager.
func (r *DaloRadiusReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.DaloRadius{}).
		Complete(r)
}
