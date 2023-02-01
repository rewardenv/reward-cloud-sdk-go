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

// ResourceTypeResourceTypeOutput Class ResourceType
type ResourceTypeResourceTypeOutput struct {
	Name NullableString `json:"name,omitempty"`
	ResourceTypeLimit NullableResourceTypeResourceTypeOutputResourceTypeLimit `json:"resourceTypeLimit,omitempty"`
}

// NewResourceTypeResourceTypeOutput instantiates a new ResourceTypeResourceTypeOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeResourceTypeOutput() *ResourceTypeResourceTypeOutput {
	this := ResourceTypeResourceTypeOutput{}
	return &this
}

// NewResourceTypeResourceTypeOutputWithDefaults instantiates a new ResourceTypeResourceTypeOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeResourceTypeOutputWithDefaults() *ResourceTypeResourceTypeOutput {
	this := ResourceTypeResourceTypeOutput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeResourceTypeOutput) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeResourceTypeOutput) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ResourceTypeResourceTypeOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ResourceTypeResourceTypeOutput) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ResourceTypeResourceTypeOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ResourceTypeResourceTypeOutput) UnsetName() {
	o.Name.Unset()
}

// GetResourceTypeLimit returns the ResourceTypeLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeResourceTypeOutput) GetResourceTypeLimit() ResourceTypeResourceTypeOutputResourceTypeLimit {
	if o == nil || isNil(o.ResourceTypeLimit.Get()) {
		var ret ResourceTypeResourceTypeOutputResourceTypeLimit
		return ret
	}
	return *o.ResourceTypeLimit.Get()
}

// GetResourceTypeLimitOk returns a tuple with the ResourceTypeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeResourceTypeOutput) GetResourceTypeLimitOk() (*ResourceTypeResourceTypeOutputResourceTypeLimit, bool) {
	if o == nil {
    return nil, false
	}
	return o.ResourceTypeLimit.Get(), o.ResourceTypeLimit.IsSet()
}

// HasResourceTypeLimit returns a boolean if a field has been set.
func (o *ResourceTypeResourceTypeOutput) HasResourceTypeLimit() bool {
	if o != nil && o.ResourceTypeLimit.IsSet() {
		return true
	}

	return false
}

// SetResourceTypeLimit gets a reference to the given NullableResourceTypeResourceTypeOutputResourceTypeLimit and assigns it to the ResourceTypeLimit field.
func (o *ResourceTypeResourceTypeOutput) SetResourceTypeLimit(v ResourceTypeResourceTypeOutputResourceTypeLimit) {
	o.ResourceTypeLimit.Set(&v)
}
// SetResourceTypeLimitNil sets the value for ResourceTypeLimit to be an explicit nil
func (o *ResourceTypeResourceTypeOutput) SetResourceTypeLimitNil() {
	o.ResourceTypeLimit.Set(nil)
}

// UnsetResourceTypeLimit ensures that no value is present for ResourceTypeLimit, not even an explicit nil
func (o *ResourceTypeResourceTypeOutput) UnsetResourceTypeLimit() {
	o.ResourceTypeLimit.Unset()
}

func (o ResourceTypeResourceTypeOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ResourceTypeLimit.IsSet() {
		toSerialize["resourceTypeLimit"] = o.ResourceTypeLimit.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceTypeResourceTypeOutput struct {
	value *ResourceTypeResourceTypeOutput
	isSet bool
}

func (v NullableResourceTypeResourceTypeOutput) Get() *ResourceTypeResourceTypeOutput {
	return v.value
}

func (v *NullableResourceTypeResourceTypeOutput) Set(val *ResourceTypeResourceTypeOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeResourceTypeOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeResourceTypeOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeResourceTypeOutput(val *ResourceTypeResourceTypeOutput) *NullableResourceTypeResourceTypeOutput {
	return &NullableResourceTypeResourceTypeOutput{value: val, isSet: true}
}

func (v NullableResourceTypeResourceTypeOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeResourceTypeOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

