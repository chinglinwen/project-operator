package v1alpha1

import (
	"wen/project-operator/pkg/project"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProjectSpec defines the desired state of Project
// +k8s:openapi-gen=true
type ProjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// project.Project

	// Project string `yaml:"project,omitempty" json:"project,omitempty"` // event.Project.PathWithNamespace
	Branch string `yaml:"branch,omitempty" json:"branch,omitempty"` // parseBranch(event.Ref)
	// Env       string    `yaml:"env,omitempty"`                              // default detect from branch, can be overwrite here
	UserName       string `yaml:"userName,omitempty" json:"userName,omitempty"`
	UserEmail      string `yaml:"userEmail,omitempty" json:"userEmail,omitempty"`
	ReleaseMessage string `yaml:"releaseMessage,omitempty" json:"releaseMessage,omitempty"`
	ReleaseAt      string `yaml:"releaseAt,omitempty" json:"releaseAt,omitempty"`
	CommitId       string `yaml:"commitid,omitempty" json:"commitid,omitempty"`
}

// ProjectStatus defines the observed state of Project
// +k8s:openapi-gen=true
type ProjectStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	project.ProjectStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Project is the Schema for the projects API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec,omitempty"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProjectList contains a list of Project
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
