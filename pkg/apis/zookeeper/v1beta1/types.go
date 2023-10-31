package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ZookeeperCluster is the Schema for the ZookeeperClusters API
type ZookeeperCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZookeeperClusterSpec   `json:"spec,omitempty"`
	Status ZookeeperClusterStatus `json:"status,omitempty"`
}

// ZookeeperClusterSpec defines the desired state of ZookeeperCluster
type ZookeeperClusterSpec struct {
	// Add custom validation using kubebuilder tags:
	// https://book.kubebuilder.io/beyond_basics/generating_crd.html
	TimeoutMins   string     `json:"timeout_mins,omitempty"`
	ClusterType   string     `json:"clusterType,omitempty"`
	ContainerCIDR string     `json:"containerCIDR,omitempty"`
	ServiceCIDR   string     `json:"serviceCIDR,omitempty"`
	MasterList    []Node     `json:"masterList" tag:"required"`
	MasterVIP     string     `json:"masterVIP,omitempty"`
	NodeList      []Node     `json:"nodeList" tag:"required"`
	EtcdList      []Node     `json:"etcdList,omitempty"`
	Region        string     `json:"region,omitempty"`
	AuthConfig    AuthConfig `json:"authConfig,omitempty"`
}

// AuthConfig defines the nodes peer authentication
type AuthConfig struct {
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	PrivateSSHKey string `json:"privateSSHKey,omitempty"`
}

// ZookeeperClusterStatus defines the observed state of ZookeeperCluster
type ZookeeperClusterStatus struct {
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Phase KubernetesOperatorPhase `json:"phase,omitempty"`

	// when job failed callback or job timeout used
	Reason string `json:"reason,omitempty"`

	// JobName is store each job name
	JobName string `json:"jobName,omitempty"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ZookeeperClusterList contains a list of ZookeeperCluster
type ZookeeperClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZookeeperCluster `json:"items"`
}

// users
// "None,Creating,Running,Failed,Scaling"
type KubernetesOperatorPhase string

type Node struct {
	IP string `json:"ip,omitempty"`
}
