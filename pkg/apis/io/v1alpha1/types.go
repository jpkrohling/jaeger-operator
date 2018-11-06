package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JaegerList is a list of Jaeger structs
type JaegerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Jaeger `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Jaeger defines the main structure for the custom-resource
type Jaeger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              JaegerSpec   `json:"spec"`
	Status            JaegerStatus `json:"status,omitempty"`
}

// JaegerSpec defines the structure of the Jaeger JSON object from the CR
type JaegerSpec struct {
	Strategy  string              `json:"strategy"`
	AllInOne  JaegerAllInOneSpec  `json:"allInOne"`
	Query     JaegerQuerySpec     `json:"query"`
	Collector JaegerCollectorSpec `json:"collector"`
	Agent     JaegerAgentSpec     `json:"agent"`
	Storage   JaegerStorageSpec   `json:"storage"`
	Ingress   JaegerIngressSpec   `json:"ingress"`
	Route     JaegerRouteSpec     `json:"route"`
}

// JaegerStatus defines what is to be returned from a status query
type JaegerStatus struct {
	// Fill me
}

// JaegerQuerySpec defines the options to be used when deploying the query
type JaegerQuerySpec struct {
	Size        int               `json:"size"`
	Image       string            `json:"image"`
	Options     Options           `json:"options"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// JaegerIngressSpec defines the options to be used when deploying the query ingress
type JaegerIngressSpec struct {
	Enabled *bool `json:"enabled"`
}

// JaegerRouteSpec defines the options to be used when deploying the query route (OpenShift-specific)
type JaegerRouteSpec struct {
	Enabled *bool `json:"enabled"`
}

// JaegerAllInOneSpec defines the options to be used when deploying the query
type JaegerAllInOneSpec struct {
	Image       string            `json:"image"`
	Options     Options           `json:"options"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// JaegerCollectorSpec defines the options to be used when deploying the collector
type JaegerCollectorSpec struct {
	Size        int               `json:"size"`
	Image       string            `json:"image"`
	Options     Options           `json:"options"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// JaegerAgentSpec defines the options to be used when deploying the agent
type JaegerAgentSpec struct {
	Strategy    string            `json:"strategy"` // can be either 'DaemonSet' or 'Sidecar' (default)
	Image       string            `json:"image"`
	Options     Options           `json:"options"`
	Annotations map[string]string `json:"annotations,omitempty"`
}

// JaegerStorageSpec defines the common storage options to be used for the query and collector
type JaegerStorageSpec struct {
	Type                  string                          `json:"type"` // can be `memory` (default), `cassandra`, `elasticsearch`, `kafka` or `managed`
	Options               Options                         `json:"options"`
	CassandraCreateSchema JaegerCassandraCreateSchemaSpec `json:"cassandraCreateSchema"`
}

// JaegerCassandraCreateSchemaSpec holds the options related to the create-schema batch job
type JaegerCassandraCreateSchemaSpec struct {
	Enabled    *bool  `json:"enabled"`
	Image      string `json:"image"`
	Datacenter string `json:"datacenter"`
	Mode       string `json:"mode"`
}
