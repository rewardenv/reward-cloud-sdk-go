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

// checks if the ResourceTypeJsonhalResourceTypeOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceTypeJsonhalResourceTypeOutput{}

// ResourceTypeJsonhalResourceTypeOutput Class ResourceType
type ResourceTypeJsonhalResourceTypeOutput struct {
	Links             *AbstractEnvironmentJsonhalLinks                               `json:"_links,omitempty"`
	Name              NullableString                                                 `json:"name,omitempty"`
	ResourceTypeLimit NullableResourceTypeJsonhalResourceTypeOutputResourceTypeLimit `json:"resourceTypeLimit,omitempty"`
}

// NewResourceTypeJsonhalResourceTypeOutput instantiates a new ResourceTypeJsonhalResourceTypeOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeJsonhalResourceTypeOutput() *ResourceTypeJsonhalResourceTypeOutput {
	this := ResourceTypeJsonhalResourceTypeOutput{}
	return &this
}

// NewResourceTypeJsonhalResourceTypeOutputWithDefaults instantiates a new ResourceTypeJsonhalResourceTypeOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeJsonhalResourceTypeOutputWithDefaults() *ResourceTypeJsonhalResourceTypeOutput {
	this := ResourceTypeJsonhalResourceTypeOutput{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceTypeJsonhalResourceTypeOutput) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeJsonhalResourceTypeOutput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceTypeJsonhalResourceTypeOutput) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *ResourceTypeJsonhalResourceTypeOutput) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeJsonhalResourceTypeOutput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeJsonhalResourceTypeOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ResourceTypeJsonhalResourceTypeOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ResourceTypeJsonhalResourceTypeOutput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ResourceTypeJsonhalResourceTypeOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ResourceTypeJsonhalResourceTypeOutput) UnsetName() {
	o.Name.Unset()
}

// GetResourceTypeLimit returns the ResourceTypeLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeJsonhalResourceTypeOutput) GetResourceTypeLimit() ResourceTypeJsonhalResourceTypeOutputResourceTypeLimit {
	if o == nil || IsNil(o.ResourceTypeLimit.Get()) {
		var ret ResourceTypeJsonhalResourceTypeOutputResourceTypeLimit
		return ret
	}
	return *o.ResourceTypeLimit.Get()
}

// GetResourceTypeLimitOk returns a tuple with the ResourceTypeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeJsonhalResourceTypeOutput) GetResourceTypeLimitOk() (*ResourceTypeJsonhalResourceTypeOutputResourceTypeLimit, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceTypeLimit.Get(), o.ResourceTypeLimit.IsSet()
}

// HasResourceTypeLimit returns a boolean if a field has been set.
func (o *ResourceTypeJsonhalResourceTypeOutput) HasResourceTypeLimit() bool {
	if o != nil && o.ResourceTypeLimit.IsSet() {
		return true
	}

	return false
}

// SetResourceTypeLimit gets a reference to the given NullableResourceTypeJsonhalResourceTypeOutputResourceTypeLimit and assigns it to the ResourceTypeLimit field.
func (o *ResourceTypeJsonhalResourceTypeOutput) SetResourceTypeLimit(v ResourceTypeJsonhalResourceTypeOutputResourceTypeLimit) {
	o.ResourceTypeLimit.Set(&v)
}

// SetResourceTypeLimitNil sets the value for ResourceTypeLimit to be an explicit nil
func (o *ResourceTypeJsonhalResourceTypeOutput) SetResourceTypeLimitNil() {
	o.ResourceTypeLimit.Set(nil)
}

// UnsetResourceTypeLimit ensures that no value is present for ResourceTypeLimit, not even an explicit nil
func (o *ResourceTypeJsonhalResourceTypeOutput) UnsetResourceTypeLimit() {
	o.ResourceTypeLimit.Unset()
}

func (o ResourceTypeJsonhalResourceTypeOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceTypeJsonhalResourceTypeOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ResourceTypeLimit.IsSet() {
		toSerialize["resourceTypeLimit"] = o.ResourceTypeLimit.Get()
	}
	return toSerialize, nil
}

type NullableResourceTypeJsonhalResourceTypeOutput struct {
	value *ResourceTypeJsonhalResourceTypeOutput
	isSet bool
}

func (v NullableResourceTypeJsonhalResourceTypeOutput) Get() *ResourceTypeJsonhalResourceTypeOutput {
	return v.value
}

func (v *NullableResourceTypeJsonhalResourceTypeOutput) Set(val *ResourceTypeJsonhalResourceTypeOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeJsonhalResourceTypeOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeJsonhalResourceTypeOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeJsonhalResourceTypeOutput(val *ResourceTypeJsonhalResourceTypeOutput) *NullableResourceTypeJsonhalResourceTypeOutput {
	return &NullableResourceTypeJsonhalResourceTypeOutput{value: val, isSet: true}
}

func (v NullableResourceTypeJsonhalResourceTypeOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeJsonhalResourceTypeOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
