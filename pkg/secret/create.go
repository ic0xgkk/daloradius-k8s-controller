package secret

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create() *corev1.Secret {
	immutable := true

	res := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "daloradius",
			Namespace: "default",
			Labels: map[string]string{
				"secret": "daloradius",
			},
		},
		Immutable: &immutable,
		Data: map[string][]byte{
			"username": []byte("root"),
			"password": []byte(""),
		},
		StringData: nil,
		Type:       "kubernetes.io/basic-auth",
	}

	return res
}
