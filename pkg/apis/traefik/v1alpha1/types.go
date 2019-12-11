package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IngressRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec IngressRouteSpec `json:"spec"`
}

type IngressRouteSpec struct {
	Entrypoints []string               `json:"entryPoints"`
	Routes      []IngressRouteRuleSpec `json:"routes"`
	// +optional
	TLS *IngressRouteTLSSpec `json:"tls,omitempty"`
}

type IngressRouteRuleSpec struct {
	Kind     string                        `json:"kind"`
	Match    string                        `json:"match"`
	Services []IngressRouteRuleServiceSpec `json:"services"`
}

type IngressRouteRuleServiceSpec struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

type IngressRouteTLSSpec struct {
	CertResolver string `json:"certResolver,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// no client needed for list as it's been created in above
type IngressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IngressRoute `json:"items"`
}
