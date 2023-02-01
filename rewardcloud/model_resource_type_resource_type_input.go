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

// ResourceTypeResourceTypeInput Class ResourceType
type ResourceTypeResourceTypeInput struct {
	Name NullableString `json:"name,omitempty"`
	ResourceTypeLimit NullableResourceTypeResourceTypeInputResourceTypeLimit `json:"resourceTypeLimit,omitempty"`
}

// NewResourceTypeResourceTypeInput instantiates a new ResourceTypeResourceTypeInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeResourceTypeInput() *ResourceTypeResourceTypeInput {
	this := ResourceTypeResourceTypeInput{}
	return &this
}

// NewResourceTypeResourceTypeInputWithDefaults instantiates a new ResourceTypeResourceTypeInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeResourceTypeInputWithDefaults() *ResourceTypeResourceTypeInput {
	this := ResourceTypeResourceTypeInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeResourceTypeInput) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeResourceTypeInput) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ResourceTypeResourceTypeInput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ResourceTypeResourceTypeInput) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ResourceTypeResourceTypeInput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ResourceTypeResourceTypeInput) UnsetName() {
	o.Name.Unset()
}

// GetResourceTypeLimit returns the ResourceTypeLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeResourceTypeInput) GetResourceTypeLimit() ResourceTypeResourceTypeInputResourceTypeLimit {
	if o == nil || isNil(o.ResourceTypeLimit.Get()) {
		var ret ResourceTypeResourceTypeInputResourceTypeLimit
		return ret
	}
	return *o.ResourceTypeLimit.Get()
}

// GetResourceTypeLimitOk returns a tuple with the ResourceTypeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeResourceTypeInput) GetResourceTypeLimitOk() (*ResourceTypeResourceTypeInputResourceTypeLimit, bool) {
	if o == nil {
    return nil, false
	}
	return o.ResourceTypeLimit.Get(), o.ResourceTypeLimit.IsSet()
}

// HasResourceTypeLimit returns a boolean if a field has been set.
func (o *ResourceTypeResourceTypeInput) HasResourceTypeLimit() bool {
	if o != nil && o.ResourceTypeLimit.IsSet() {
		return true
	}

	return false
}

// SetResourceTypeLimit gets a reference to the given NullableResourceTypeResourceTypeInputResourceTypeLimit and assigns it to the ResourceTypeLimit field.
func (o *ResourceTypeResourceTypeInput) SetResourceTypeLimit(v ResourceTypeResourceTypeInputResourceTypeLimit) {
	o.ResourceTypeLimit.Set(&v)
}
// SetResourceTypeLimitNil sets the value for ResourceTypeLimit to be an explicit nil
func (o *ResourceTypeResourceTypeInput) SetResourceTypeLimitNil() {
	o.ResourceTypeLimit.Set(nil)
}

// UnsetResourceTypeLimit ensures that no value is present for ResourceTypeLimit, not even an explicit nil
func (o *ResourceTypeResourceTypeInput) UnsetResourceTypeLimit() {
	o.ResourceTypeLimit.Unset()
}

func (o ResourceTypeResourceTypeInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ResourceTypeLimit.IsSet() {
		toSerialize["resourceTypeLimit"] = o.ResourceTypeLimit.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceTypeResourceTypeInput struct {
	value *ResourceTypeResourceTypeInput
	isSet bool
}

func (v NullableResourceTypeResourceTypeInput) Get() *ResourceTypeResourceTypeInput {
	return v.value
}

func (v *NullableResourceTypeResourceTypeInput) Set(val *ResourceTypeResourceTypeInput) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeResourceTypeInput) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeResourceTypeInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeResourceTypeInput(val *ResourceTypeResourceTypeInput) *NullableResourceTypeResourceTypeInput {
	return &NullableResourceTypeResourceTypeInput{value: val, isSet: true}
}

func (v NullableResourceTypeResourceTypeInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeResourceTypeInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

