/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ServiceAccountApplyConfiguration represents a declarative configuration of the ServiceAccount type for use
// with apply.
type ServiceAccountApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration    `json:",inline"`
	*metav1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Secrets                              []ObjectReferenceApplyConfiguration      `json:"secrets,omitempty"`
	ImagePullSecrets                     []LocalObjectReferenceApplyConfiguration `json:"imagePullSecrets,omitempty"`
	AutomountServiceAccountToken         *bool                                    `json:"automountServiceAccountToken,omitempty"`
}

// ServiceAccount constructs a declarative configuration of the ServiceAccount type for use with
// apply.
func ServiceAccount(name, namespace string) *ServiceAccountApplyConfiguration {
	b := &ServiceAccountApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("ServiceAccount")
	b.WithAPIVersion("v1")
	return b
}

// ExtractServiceAccount extracts the applied configuration owned by fieldManager from
// serviceAccount. If no managedFields are found in serviceAccount for fieldManager, a
// ServiceAccountApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// serviceAccount must be a unmodified ServiceAccount API object that was retrieved from the Kubernetes API.
// ExtractServiceAccount provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractServiceAccount(serviceAccount *corev1.ServiceAccount, fieldManager string) (*ServiceAccountApplyConfiguration, error) {
	return extractServiceAccount(serviceAccount, fieldManager, "")
}

// ExtractServiceAccountStatus is the same as ExtractServiceAccount except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractServiceAccountStatus(serviceAccount *corev1.ServiceAccount, fieldManager string) (*ServiceAccountApplyConfiguration, error) {
	return extractServiceAccount(serviceAccount, fieldManager, "status")
}

func extractServiceAccount(serviceAccount *corev1.ServiceAccount, fieldManager string, subresource string) (*ServiceAccountApplyConfiguration, error) {
	b := &ServiceAccountApplyConfiguration{}
	err := managedfields.ExtractInto(serviceAccount, internal.Parser().Type("io.k8s.api.core.v1.ServiceAccount"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(serviceAccount.Name)
	b.WithNamespace(serviceAccount.Namespace)

	b.WithKind("ServiceAccount")
	b.WithAPIVersion("v1")
	return b, nil
}
func (b ServiceAccountApplyConfiguration) IsApplyConfiguration() {}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithKind(value string) *ServiceAccountApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithAPIVersion(value string) *ServiceAccountApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithName(value string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithGenerateName(value string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithNamespace(value string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithUID(value types.UID) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithResourceVersion(value string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithGeneration(value int64) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithCreationTimestamp(value apismetav1.Time) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithDeletionTimestamp(value apismetav1.Time) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ServiceAccountApplyConfiguration) WithLabels(entries map[string]string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *ServiceAccountApplyConfiguration) WithAnnotations(entries map[string]string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *ServiceAccountApplyConfiguration) WithOwnerReferences(values ...*metav1.OwnerReferenceApplyConfiguration) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *ServiceAccountApplyConfiguration) WithFinalizers(values ...string) *ServiceAccountApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *ServiceAccountApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &metav1.ObjectMetaApplyConfiguration{}
	}
}

// WithSecrets adds the given value to the Secrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Secrets field.
func (b *ServiceAccountApplyConfiguration) WithSecrets(values ...*ObjectReferenceApplyConfiguration) *ServiceAccountApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSecrets")
		}
		b.Secrets = append(b.Secrets, *values[i])
	}
	return b
}

// WithImagePullSecrets adds the given value to the ImagePullSecrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImagePullSecrets field.
func (b *ServiceAccountApplyConfiguration) WithImagePullSecrets(values ...*LocalObjectReferenceApplyConfiguration) *ServiceAccountApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImagePullSecrets")
		}
		b.ImagePullSecrets = append(b.ImagePullSecrets, *values[i])
	}
	return b
}

// WithAutomountServiceAccountToken sets the AutomountServiceAccountToken field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutomountServiceAccountToken field is set to the value of the last call.
func (b *ServiceAccountApplyConfiguration) WithAutomountServiceAccountToken(value bool) *ServiceAccountApplyConfiguration {
	b.AutomountServiceAccountToken = &value
	return b
}

// GetKind retrieves the value of the Kind field in the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetKind() *string {
	return b.TypeMetaApplyConfiguration.Kind
}

// GetAPIVersion retrieves the value of the APIVersion field in the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetAPIVersion() *string {
	return b.TypeMetaApplyConfiguration.APIVersion
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}

// GetNamespace retrieves the value of the Namespace field in the declarative configuration.
func (b *ServiceAccountApplyConfiguration) GetNamespace() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Namespace
}
