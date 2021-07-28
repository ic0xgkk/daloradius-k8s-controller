package service

import corev1 "k8s.io/api/core/v1"

func PatchHttpPort(svc *corev1.Service, port uint16) {
	svc.Spec.Ports[0].NodePort = int32(port)

	if svc.Spec.Ports[0].Name != "http" {
		panic("http port order wrong")
	}
}

func PatchAuthPort(svc *corev1.Service, port uint16) {
	svc.Spec.Ports[1].NodePort = int32(port)

	if svc.Spec.Ports[1].Name != "auth" {
		panic("auth port order wrong")
	}
}

func PatchAcctPort(svc *corev1.Service, port uint16) {
	svc.Spec.Ports[2].NodePort = int32(port)

	if svc.Spec.Ports[2].Name != "acct" {
		panic("acct port order wrong")
	}
}
