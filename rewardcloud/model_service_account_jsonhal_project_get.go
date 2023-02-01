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

// ServiceAccountJsonhalProjectGet 
type ServiceAccountJsonhalProjectGet struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	ServiceAccountGit NullableServiceAccountJsonhalProjectGetServiceAccountGit `json:"serviceAccountGit,omitempty"`
	ServiceAccountRegistry NullableServiceAccountJsonhalProjectGetServiceAccountRegistry `json:"serviceAccountRegistry,omitempty"`
}

// NewServiceAccountJsonhalProjectGet instantiates a new ServiceAccountJsonhalProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountJsonhalProjectGet() *ServiceAccountJsonhalProjectGet {
	this := ServiceAccountJsonhalProjectGet{}
	return &this
}

// NewServiceAccountJsonhalProjectGetWithDefaults instantiates a new ServiceAccountJsonhalProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountJsonhalProjectGetWithDefaults() *ServiceAccountJsonhalProjectGet {
	this := ServiceAccountJsonhalProjectGet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceAccountJsonhalProjectGet) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountJsonhalProjectGet) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceAccountJsonhalProjectGet) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *ServiceAccountJsonhalProjectGet) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetServiceAccountGit returns the ServiceAccountGit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountJsonhalProjectGet) GetServiceAccountGit() ServiceAccountJsonhalProjectGetServiceAccountGit {
	if o == nil || isNil(o.ServiceAccountGit.Get()) {
		var ret ServiceAccountJsonhalProjectGetServiceAccountGit
		return ret
	}
	return *o.ServiceAccountGit.Get()
}

// GetServiceAccountGitOk returns a tuple with the ServiceAccountGit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountJsonhalProjectGet) GetServiceAccountGitOk() (*ServiceAccountJsonhalProjectGetServiceAccountGit, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceAccountGit.Get(), o.ServiceAccountGit.IsSet()
}

// HasServiceAccountGit returns a boolean if a field has been set.
func (o *ServiceAccountJsonhalProjectGet) HasServiceAccountGit() bool {
	if o != nil && o.ServiceAccountGit.IsSet() {
		return true
	}

	return false
}

// SetServiceAccountGit gets a reference to the given NullableServiceAccountJsonhalProjectGetServiceAccountGit and assigns it to the ServiceAccountGit field.
func (o *ServiceAccountJsonhalProjectGet) SetServiceAccountGit(v ServiceAccountJsonhalProjectGetServiceAccountGit) {
	o.ServiceAccountGit.Set(&v)
}
// SetServiceAccountGitNil sets the value for ServiceAccountGit to be an explicit nil
func (o *ServiceAccountJsonhalProjectGet) SetServiceAccountGitNil() {
	o.ServiceAccountGit.Set(nil)
}

// UnsetServiceAccountGit ensures that no value is present for ServiceAccountGit, not even an explicit nil
func (o *ServiceAccountJsonhalProjectGet) UnsetServiceAccountGit() {
	o.ServiceAccountGit.Unset()
}

// GetServiceAccountRegistry returns the ServiceAccountRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountJsonhalProjectGet) GetServiceAccountRegistry() ServiceAccountJsonhalProjectGetServiceAccountRegistry {
	if o == nil || isNil(o.ServiceAccountRegistry.Get()) {
		var ret ServiceAccountJsonhalProjectGetServiceAccountRegistry
		return ret
	}
	return *o.ServiceAccountRegistry.Get()
}

// GetServiceAccountRegistryOk returns a tuple with the ServiceAccountRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountJsonhalProjectGet) GetServiceAccountRegistryOk() (*ServiceAccountJsonhalProjectGetServiceAccountRegistry, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceAccountRegistry.Get(), o.ServiceAccountRegistry.IsSet()
}

// HasServiceAccountRegistry returns a boolean if a field has been set.
func (o *ServiceAccountJsonhalProjectGet) HasServiceAccountRegistry() bool {
	if o != nil && o.ServiceAccountRegistry.IsSet() {
		return true
	}

	return false
}

// SetServiceAccountRegistry gets a reference to the given NullableServiceAccountJsonhalProjectGetServiceAccountRegistry and assigns it to the ServiceAccountRegistry field.
func (o *ServiceAccountJsonhalProjectGet) SetServiceAccountRegistry(v ServiceAccountJsonhalProjectGetServiceAccountRegistry) {
	o.ServiceAccountRegistry.Set(&v)
}
// SetServiceAccountRegistryNil sets the value for ServiceAccountRegistry to be an explicit nil
func (o *ServiceAccountJsonhalProjectGet) SetServiceAccountRegistryNil() {
	o.ServiceAccountRegistry.Set(nil)
}

// UnsetServiceAccountRegistry ensures that no value is present for ServiceAccountRegistry, not even an explicit nil
func (o *ServiceAccountJsonhalProjectGet) UnsetServiceAccountRegistry() {
	o.ServiceAccountRegistry.Unset()
}

func (o ServiceAccountJsonhalProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if o.ServiceAccountGit.IsSet() {
		toSerialize["serviceAccountGit"] = o.ServiceAccountGit.Get()
	}
	if o.ServiceAccountRegistry.IsSet() {
		toSerialize["serviceAccountRegistry"] = o.ServiceAccountRegistry.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountJsonhalProjectGet struct {
	value *ServiceAccountJsonhalProjectGet
	isSet bool
}

func (v NullableServiceAccountJsonhalProjectGet) Get() *ServiceAccountJsonhalProjectGet {
	return v.value
}

func (v *NullableServiceAccountJsonhalProjectGet) Set(val *ServiceAccountJsonhalProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountJsonhalProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountJsonhalProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountJsonhalProjectGet(val *ServiceAccountJsonhalProjectGet) *NullableServiceAccountJsonhalProjectGet {
	return &NullableServiceAccountJsonhalProjectGet{value: val, isSet: true}
}

func (v NullableServiceAccountJsonhalProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountJsonhalProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


