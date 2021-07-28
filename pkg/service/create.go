package service

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Create() *corev1.Service {
	res := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "daloradius",
			Namespace: "default",
			Labels: map[string]string{
				"service": "daloradius",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:        "http",
					Protocol:    "TCP",
					AppProtocol: nil,
					Port:        80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
						StrVal: "",
					},
					NodePort: int32(80),
				},
				{
					Name:        "auth",
					Protocol:    "UDP",
					AppProtocol: nil,
					Port:        1812,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 1812,
						StrVal: "",
					},
					NodePort: int32(1812),
				},
				{
					Name:        "acct",
					Protocol:    "UDP",
					AppProtocol: nil,
					Port:        1813,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 1813,
						StrVal: "",
					},
					NodePort: int32(1813),
				},
			},
			Selector: map[string]string{
				"deployment": "daloradius",
			},
			ClusterIP:                     "",
			ClusterIPs:                    nil,
			Type:                          "NodePort",
			ExternalIPs:                   nil,
			SessionAffinity:               "ClientIP",
			LoadBalancerIP:                "",
			LoadBalancerSourceRanges:      nil,
			ExternalName:                  "",
			ExternalTrafficPolicy:         "",
			HealthCheckNodePort:           0,
			PublishNotReadyAddresses:      false,
			SessionAffinityConfig:         nil,
			TopologyKeys:                  nil,
			IPFamilies:                    nil,
			IPFamilyPolicy:                nil,
			AllocateLoadBalancerNodePorts: nil,
		},
		Status: corev1.ServiceStatus{},
	}

	return res
}
