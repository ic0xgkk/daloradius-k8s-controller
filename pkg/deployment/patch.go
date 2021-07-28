package deployment

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"strconv"
)

func PatchTimezone(dep *appsv1.Deployment, tz string) {
	dep.Spec.Template.Spec.Containers[0].Env[0].Value = tz

	if dep.Spec.Template.Spec.Containers[0].Env[0].Name != "TZ" {
		panic("Timezone order wrong!")
	}
}

func PatchDatabaseName(dep *appsv1.Deployment, s string) {
	dep.Spec.Template.Spec.Containers[0].Env[1].Value = s

	if dep.Spec.Template.Spec.Containers[0].Env[1].Name != "MYSQL_DATABASE" {
		panic("MYSQL_DATABASE order wrong")
	}
}

func PatchDatabasePort(dep *appsv1.Deployment, port uint16) {
	dep.Spec.Template.Spec.Containers[0].Env[2].Value = strconv.Itoa(int(port))

	if dep.Spec.Template.Spec.Containers[0].Env[2].Name != "MYSQL_PORT" {
		panic("MYSQL_PORT order wrong")
	}
}

func PatchDatabaseHost(dep *appsv1.Deployment, host string) {
	dep.Spec.Template.Spec.Containers[0].Env[3].Value = host

	if dep.Spec.Template.Spec.Containers[0].Env[3].Name != "MYSQL_HOST" {
		panic("MYSQL_HOST order wrong")
	}
}

func PatchReplicas(dep *appsv1.Deployment, replicas uint32) {
	r := int32(replicas)
	dep.Spec.Replicas = &r
}

func PatchImageTag(dep *appsv1.Deployment, tag string) {
	dep.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", "a980883231/daloradius-docker", tag)
}
