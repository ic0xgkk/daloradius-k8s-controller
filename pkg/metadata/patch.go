package metadata

import "sigs.k8s.io/controller-runtime/pkg/client"

func PatchNamespace(obj client.Object, ns string) {
	obj.SetNamespace(ns)
}
