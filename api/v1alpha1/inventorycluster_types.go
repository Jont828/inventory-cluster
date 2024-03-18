/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InventoryClusterSpec defines the desired state of InventoryCluster.
type InventoryClusterSpec struct {
	// DisplayName is the human-readable name of the cluster set by the consumer of the cluster.
	DisplayName string `json:"displayName"`

	// ClusterManager is an reference to the ClusterManager that manages the cluster.
	// Note: There must also be a `multicluster.x-k8s.io/cluster-manager` label on the InventoryCluster object matching the name of the ClusterManager.
	ClusterManager ClusterManager `json:"clusterManager"`
}

type ClusterManager struct {
	// Name is the name of the cluster manager.
	Name string `json:"name"`
}

// InventoryClusterStatus defines the observed state of InventoryCluster.
type InventoryClusterStatus struct {
	// Version is the Kubernetes version of the cluster.
	// +optional
	Version string `json:"version,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// InventoryCluster is the Schema for the inventoryclusters API
type InventoryCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InventoryClusterSpec   `json:"spec,omitempty"`
	Status InventoryClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InventoryClusterList contains a list of InventoryCluster
type InventoryClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InventoryCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InventoryCluster{}, &InventoryClusterList{})
}
