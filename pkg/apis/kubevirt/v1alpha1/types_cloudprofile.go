// Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudProfileConfig contains provider-specific configuration that is embedded into Gardener's `CloudProfile`
// resource.
type CloudProfileConfig struct {
	metav1.TypeMeta `json:",inline"`
	// MachineImages is the list of machine images that are understood by the controller. It maps
	// logical names and versions to provider-specific identifiers.
	MachineImages []MachineImages `json:"machineImages"`
	// MachineTypes extend machineType object to KubeVirt provider specific config
	// +optional
	MachineTypes []MachineType `json:"machineTypes,omitempty"`
}

// MachineImages is a mapping from logical names and versions to provider-specific identifiers.
type MachineImages struct {
	// Name is the logical name of the machine image.
	Name string `json:"name"`
	// Versions contains versions and a provider-specific identifier.
	Versions []MachineImageVersion `json:"versions"`
}

// MachineImageVersion contains a version and a provider-specific identifier.
type MachineImageVersion struct {
	// Version is the version of the image.
	Version string `json:"version"`
	// SourceURL is the url of the image
	SourceURL string `json:"sourceURL"`
}

// MachineType extend machineType object to KubeVirt provider specific config
type MachineType struct {
	// Name is a reference for a machine type object
	Name string `json:"name"`
	// Limits defines resources limits of KubeVirt VMs
	// +optional
	Limits *ResourcesLimits `json:"limits,omitempty"`
}

// ResourceLimits define resources limits of KubeVirt VMs
type ResourcesLimits struct {
	// CPU limits
	// +optional
	CPU resource.Quantity `json:"cpu,omitempty"`
	// Memory limits
	// +optional
	Memory resource.Quantity `json:"memory,omitempty"`
}
