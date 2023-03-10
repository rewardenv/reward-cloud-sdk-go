/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// ServiceAccountProjectGet 
type ServiceAccountProjectGet struct {
	ServiceAccountGit NullableServiceAccountProjectGetServiceAccountGit `json:"serviceAccountGit,omitempty"`
	ServiceAccountRegistry NullableServiceAccountProjectGetServiceAccountRegistry `json:"serviceAccountRegistry,omitempty"`
}

// NewServiceAccountProjectGet instantiates a new ServiceAccountProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountProjectGet() *ServiceAccountProjectGet {
	this := ServiceAccountProjectGet{}
	return &this
}

// NewServiceAccountProjectGetWithDefaults instantiates a new ServiceAccountProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountProjectGetWithDefaults() *ServiceAccountProjectGet {
	this := ServiceAccountProjectGet{}
	return &this
}

// GetServiceAccountGit returns the ServiceAccountGit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountProjectGet) GetServiceAccountGit() ServiceAccountProjectGetServiceAccountGit {
	if o == nil || isNil(o.ServiceAccountGit.Get()) {
		var ret ServiceAccountProjectGetServiceAccountGit
		return ret
	}
	return *o.ServiceAccountGit.Get()
}

// GetServiceAccountGitOk returns a tuple with the ServiceAccountGit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountProjectGet) GetServiceAccountGitOk() (*ServiceAccountProjectGetServiceAccountGit, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceAccountGit.Get(), o.ServiceAccountGit.IsSet()
}

// HasServiceAccountGit returns a boolean if a field has been set.
func (o *ServiceAccountProjectGet) HasServiceAccountGit() bool {
	if o != nil && o.ServiceAccountGit.IsSet() {
		return true
	}

	return false
}

// SetServiceAccountGit gets a reference to the given NullableServiceAccountProjectGetServiceAccountGit and assigns it to the ServiceAccountGit field.
func (o *ServiceAccountProjectGet) SetServiceAccountGit(v ServiceAccountProjectGetServiceAccountGit) {
	o.ServiceAccountGit.Set(&v)
}
// SetServiceAccountGitNil sets the value for ServiceAccountGit to be an explicit nil
func (o *ServiceAccountProjectGet) SetServiceAccountGitNil() {
	o.ServiceAccountGit.Set(nil)
}

// UnsetServiceAccountGit ensures that no value is present for ServiceAccountGit, not even an explicit nil
func (o *ServiceAccountProjectGet) UnsetServiceAccountGit() {
	o.ServiceAccountGit.Unset()
}

// GetServiceAccountRegistry returns the ServiceAccountRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountProjectGet) GetServiceAccountRegistry() ServiceAccountProjectGetServiceAccountRegistry {
	if o == nil || isNil(o.ServiceAccountRegistry.Get()) {
		var ret ServiceAccountProjectGetServiceAccountRegistry
		return ret
	}
	return *o.ServiceAccountRegistry.Get()
}

// GetServiceAccountRegistryOk returns a tuple with the ServiceAccountRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountProjectGet) GetServiceAccountRegistryOk() (*ServiceAccountProjectGetServiceAccountRegistry, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceAccountRegistry.Get(), o.ServiceAccountRegistry.IsSet()
}

// HasServiceAccountRegistry returns a boolean if a field has been set.
func (o *ServiceAccountProjectGet) HasServiceAccountRegistry() bool {
	if o != nil && o.ServiceAccountRegistry.IsSet() {
		return true
	}

	return false
}

// SetServiceAccountRegistry gets a reference to the given NullableServiceAccountProjectGetServiceAccountRegistry and assigns it to the ServiceAccountRegistry field.
func (o *ServiceAccountProjectGet) SetServiceAccountRegistry(v ServiceAccountProjectGetServiceAccountRegistry) {
	o.ServiceAccountRegistry.Set(&v)
}
// SetServiceAccountRegistryNil sets the value for ServiceAccountRegistry to be an explicit nil
func (o *ServiceAccountProjectGet) SetServiceAccountRegistryNil() {
	o.ServiceAccountRegistry.Set(nil)
}

// UnsetServiceAccountRegistry ensures that no value is present for ServiceAccountRegistry, not even an explicit nil
func (o *ServiceAccountProjectGet) UnsetServiceAccountRegistry() {
	o.ServiceAccountRegistry.Unset()
}

func (o ServiceAccountProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceAccountGit.IsSet() {
		toSerialize["serviceAccountGit"] = o.ServiceAccountGit.Get()
	}
	if o.ServiceAccountRegistry.IsSet() {
		toSerialize["serviceAccountRegistry"] = o.ServiceAccountRegistry.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountProjectGet struct {
	value *ServiceAccountProjectGet
	isSet bool
}

func (v NullableServiceAccountProjectGet) Get() *ServiceAccountProjectGet {
	return v.value
}

func (v *NullableServiceAccountProjectGet) Set(val *ServiceAccountProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountProjectGet(val *ServiceAccountProjectGet) *NullableServiceAccountProjectGet {
	return &NullableServiceAccountProjectGet{value: val, isSet: true}
}

func (v NullableServiceAccountProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


