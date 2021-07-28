package deployment

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

func Create() *appsv1.Deployment {
	replicas := int32(3)

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "daloradius",
			Namespace: "default",
			Labels: map[string]string{
				"deployment": "daloradius",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"deployment": "daloradius",
				},
				MatchExpressions: nil,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"deployment": "daloradius",
					},
				},
				Spec: corev1.PodSpec{
					Volumes:        nil,
					InitContainers: nil,
					Containers: []corev1.Container{
						{
							Name:       "daloradius",
							Image:      fmt.Sprintf("%s:%s", "a980883231/daloradius-docker", "latest"),
							Command:    nil,
							Args:       nil,
							WorkingDir: "",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									HostPort:      0,
									ContainerPort: 80,
									Protocol:      "TCP",
									HostIP:        "",
								},
								{
									Name:          "auth",
									HostPort:      0,
									ContainerPort: 1812,
									Protocol:      "UDP",
									HostIP:        "",
								},
								{
									Name:          "acct",
									HostPort:      0,
									ContainerPort: 1813,
									Protocol:      "UDP",
									HostIP:        "",
								},
							},
							EnvFrom: nil,
							Env: []corev1.EnvVar{
								{
									Name:      "TZ",
									Value:     "Asia/Shanghai",
									ValueFrom: nil,
								},
								{
									Name:      "MYSQL_DATABASE",
									Value:     "daloradius",
									ValueFrom: nil,
								},
								{
									Name:      "MYSQL_PORT",
									Value:     strconv.Itoa(3306),
									ValueFrom: nil,
								},
								{
									Name:      "MYSQL_HOST",
									Value:     "localhost",
									ValueFrom: nil,
								},
								{
									Name:  "MYSQL_USER",
									Value: "",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef:         nil,
										ResourceFieldRef: nil,
										ConfigMapKeyRef:  nil,
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "daloradius",
											},
											Key:      "username",
											Optional: nil,
										},
									},
								},
								{
									Name:  "MYSQL_PASSWORD",
									Value: "",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef:         nil,
										ResourceFieldRef: nil,
										ConfigMapKeyRef:  nil,
										SecretKeyRef: &corev1.SecretKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "daloradius",
											},
											Key:      "password",
											Optional: nil,
										},
									},
								},
							},
							Resources:     corev1.ResourceRequirements{},
							VolumeMounts:  nil,
							VolumeDevices: nil,
							LivenessProbe: &corev1.Probe{
								Handler: corev1.Handler{
									Exec:    nil,
									HTTPGet: nil,
									TCPSocket: &corev1.TCPSocketAction{
										Port: intstr.IntOrString{
											Type:   intstr.Int,
											IntVal: 80,
											StrVal: "",
										},
										Host: "",
									},
								},
								InitialDelaySeconds: 10,
								TimeoutSeconds:      5,
								PeriodSeconds:       3,
								SuccessThreshold:    1,
								FailureThreshold:    2,
							},
							ReadinessProbe:           nil,
							StartupProbe:             nil,
							Lifecycle:                nil,
							TerminationMessagePath:   "",
							TerminationMessagePolicy: "",
							ImagePullPolicy:          "",
							SecurityContext:          nil,
							Stdin:                    false,
							StdinOnce:                false,
							TTY:                      false,
						},
					},
					EphemeralContainers:           nil,
					RestartPolicy:                 "",
					TerminationGracePeriodSeconds: nil,
					ActiveDeadlineSeconds:         nil,
					DNSPolicy:                     "",
					NodeSelector:                  nil,
					ServiceAccountName:            "",
					AutomountServiceAccountToken:  nil,
					ShareProcessNamespace:         nil,
					SecurityContext:               nil,
					ImagePullSecrets:              nil,
					Hostname:                      "",
					Subdomain:                     "",
					Affinity:                      nil,
					SchedulerName:                 "",
					Tolerations:                   nil,
					HostAliases:                   nil,
					PriorityClassName:             "",
					Priority:                      nil,
					DNSConfig:                     nil,
					ReadinessGates:                nil,
					RuntimeClassName:              nil,
					EnableServiceLinks:            nil,
					PreemptionPolicy:              nil,
					Overhead:                      nil,
					TopologySpreadConstraints:     nil,
					SetHostnameAsFQDN:             nil,
				},
			},
			Strategy:                appsv1.DeploymentStrategy{},
			MinReadySeconds:         0,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
		Status: appsv1.DeploymentStatus{},
	}
}
