/*
Copyright 2021 Harris<i@xuegaogg.com>.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DaloRadiusSpec defines the desired state of DaloRadius
type DaloRadiusSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// TZ database name in https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	// Default to Asia/Shanghai
	Timezone string `json:"timezone,omitempty"`
	// Default to 3
	Replicas uint32 `json:"replicas,omitempty"`
	// NodePort for DaloRadius port
	HttpPort uint16 `json:"http_port,omitempty"`
	// NodePort for FreeRadius auth port
	AuthPort uint16 `json:"auth_port,omitempty"`
	// NodePort for FreeRadius acct port
	AcctPort uint16 `json:"acct_port,omitempty"`
	// Must set mysql host
	MysqlHost string `json:"mysql_host,omitempty"`
	// Default to 3306
	MysqlPort uint16 `json:"mysql_port,omitempty"`
	// Must set mysql username
	MysqlUsername string `json:"mysql_username,omitempty"`
	// Must set mysql password
	MysqlPassword string `json:"mysql_password,omitempty"`
	// Must set mysql database name
	MysqlDatabase string `json:"mysql_database,omitempty"`
	// Default to latest
	ImageTag string `json:"image_tag,omitempty"`
}

// DaloRadiusStatus defines the observed state of DaloRadius
type DaloRadiusStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DaloRadius is the Schema for the daloradius API
type DaloRadius struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DaloRadiusSpec   `json:"spec,omitempty"`
	Status DaloRadiusStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DaloRadiusList contains a list of DaloRadius
type DaloRadiusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaloRadius `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DaloRadius{}, &DaloRadiusList{})
}
