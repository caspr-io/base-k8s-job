package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func NewIngressRoute(name string, labels map[string]string) *IngressRoute {
	return &IngressRoute{
		TypeMeta: metav1.TypeMeta{
			Kind:       "IngressRoute",
			APIVersion: "traefik.containo.us/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Spec: IngressRouteSpec{},
	}
}
