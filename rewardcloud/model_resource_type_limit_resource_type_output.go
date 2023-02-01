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

// ResourceTypeLimitResourceTypeOutput 
type ResourceTypeLimitResourceTypeOutput struct {
	ProjectMinValue NullableInt32 `json:"projectMinValue,omitempty"`
	ProjectMaxValue NullableInt32 `json:"projectMaxValue,omitempty"`
	EnvironmentMinValue NullableInt32 `json:"environmentMinValue,omitempty"`
	EnvironmentMaxValue NullableInt32 `json:"environmentMaxValue,omitempty"`
	ComponentMinValue NullableInt32 `json:"componentMinValue,omitempty"`
	ComponentMaxValue NullableInt32 `json:"componentMaxValue,omitempty"`
}

// NewResourceTypeLimitResourceTypeOutput instantiates a new ResourceTypeLimitResourceTypeOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeLimitResourceTypeOutput() *ResourceTypeLimitResourceTypeOutput {
	this := ResourceTypeLimitResourceTypeOutput{}
	return &this
}

// NewResourceTypeLimitResourceTypeOutputWithDefaults instantiates a new ResourceTypeLimitResourceTypeOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeLimitResourceTypeOutputWithDefaults() *ResourceTypeLimitResourceTypeOutput {
	this := ResourceTypeLimitResourceTypeOutput{}
	return &this
}

// GetProjectMinValue returns the ProjectMinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimitResourceTypeOutput) GetProjectMinValue() int32 {
	if o == nil || isNil(o.ProjectMinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectMinValue.Get()
}

// GetProjectMinValueOk returns a tuple with the ProjectMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimitResourceTypeOutput) GetProjectMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectMinValue.Get(), o.ProjectMinValue.IsSet()
}

// HasProjectMinValue returns a boolean if a field has been set.
func (o *ResourceTypeLimitResourceTypeOutput) HasProjectMinValue() bool {
	if o != nil && o.ProjectMinValue.IsSet() {
		return true
	}

	return false
}

// SetProjectMinValue gets a reference to the given NullableInt32 and assigns it to the ProjectMinValue field.
func (o *ResourceTypeLimitResourceTypeOutput) SetProjectMinValue(v int32) {
	o.ProjectMinValue.Set(&v)
}
// SetProjectMinValueNil sets the value for ProjectMinValue to be an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) SetProjectMinValueNil() {
	o.ProjectMinValue.Set(nil)
}

// UnsetProjectMinValue ensures that no value is present for ProjectMinValue, not even an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) UnsetProjectMinValue() {
	o.ProjectMinValue.Unset()
}

// GetProjectMaxValue returns the ProjectMaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimitResourceTypeOutput) GetProjectMaxValue() int32 {
	if o == nil || isNil(o.ProjectMaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectMaxValue.Get()
}

// GetProjectMaxValueOk returns a tuple with the ProjectMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimitResourceTypeOutput) GetProjectMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectMaxValue.Get(), o.ProjectMaxValue.IsSet()
}

// HasProjectMaxValue returns a boolean if a field has been set.
func (o *ResourceTypeLimitResourceTypeOutput) HasProjectMaxValue() bool {
	if o != nil && o.ProjectMaxValue.IsSet() {
		return true
	}

	return false
}

// SetProjectMaxValue gets a reference to the given NullableInt32 and assigns it to the ProjectMaxValue field.
func (o *ResourceTypeLimitResourceTypeOutput) SetProjectMaxValue(v int32) {
	o.ProjectMaxValue.Set(&v)
}
// SetProjectMaxValueNil sets the value for ProjectMaxValue to be an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) SetProjectMaxValueNil() {
	o.ProjectMaxValue.Set(nil)
}

// UnsetProjectMaxValue ensures that no value is present for ProjectMaxValue, not even an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) UnsetProjectMaxValue() {
	o.ProjectMaxValue.Unset()
}

// GetEnvironmentMinValue returns the EnvironmentMinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimitResourceTypeOutput) GetEnvironmentMinValue() int32 {
	if o == nil || isNil(o.EnvironmentMinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.EnvironmentMinValue.Get()
}

// GetEnvironmentMinValueOk returns a tuple with the EnvironmentMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimitResourceTypeOutput) GetEnvironmentMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvironmentMinValue.Get(), o.EnvironmentMinValue.IsSet()
}

// HasEnvironmentMinValue returns a boolean if a field has been set.
func (o *ResourceTypeLimitResourceTypeOutput) HasEnvironmentMinValue() bool {
	if o != nil && o.EnvironmentMinValue.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentMinValue gets a reference to the given NullableInt32 and assigns it to the EnvironmentMinValue field.
func (o *ResourceTypeLimitResourceTypeOutput) SetEnvironmentMinValue(v int32) {
	o.EnvironmentMinValue.Set(&v)
}
// SetEnvironmentMinValueNil sets the value for EnvironmentMinValue to be an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) SetEnvironmentMinValueNil() {
	o.EnvironmentMinValue.Set(nil)
}

// UnsetEnvironmentMinValue ensures that no value is present for EnvironmentMinValue, not even an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) UnsetEnvironmentMinValue() {
	o.EnvironmentMinValue.Unset()
}

// GetEnvironmentMaxValue returns the EnvironmentMaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimitResourceTypeOutput) GetEnvironmentMaxValue() int32 {
	if o == nil || isNil(o.EnvironmentMaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.EnvironmentMaxValue.Get()
}

// GetEnvironmentMaxValueOk returns a tuple with the EnvironmentMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimitResourceTypeOutput) GetEnvironmentMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvironmentMaxValue.Get(), o.EnvironmentMaxValue.IsSet()
}

// HasEnvironmentMaxValue returns a boolean if a field has been set.
func (o *ResourceTypeLimitResourceTypeOutput) HasEnvironmentMaxValue() bool {
	if o != nil && o.EnvironmentMaxValue.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentMaxValue gets a reference to the given NullableInt32 and assigns it to the EnvironmentMaxValue field.
func (o *ResourceTypeLimitResourceTypeOutput) SetEnvironmentMaxValue(v int32) {
	o.EnvironmentMaxValue.Set(&v)
}
// SetEnvironmentMaxValueNil sets the value for EnvironmentMaxValue to be an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) SetEnvironmentMaxValueNil() {
	o.EnvironmentMaxValue.Set(nil)
}

// UnsetEnvironmentMaxValue ensures that no value is present for EnvironmentMaxValue, not even an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) UnsetEnvironmentMaxValue() {
	o.EnvironmentMaxValue.Unset()
}

// GetComponentMinValue returns the ComponentMinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimitResourceTypeOutput) GetComponentMinValue() int32 {
	if o == nil || isNil(o.ComponentMinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ComponentMinValue.Get()
}

// GetComponentMinValueOk returns a tuple with the ComponentMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimitResourceTypeOutput) GetComponentMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentMinValue.Get(), o.ComponentMinValue.IsSet()
}

// HasComponentMinValue returns a boolean if a field has been set.
func (o *ResourceTypeLimitResourceTypeOutput) HasComponentMinValue() bool {
	if o != nil && o.ComponentMinValue.IsSet() {
		return true
	}

	return false
}

// SetComponentMinValue gets a reference to the given NullableInt32 and assigns it to the ComponentMinValue field.
func (o *ResourceTypeLimitResourceTypeOutput) SetComponentMinValue(v int32) {
	o.ComponentMinValue.Set(&v)
}
// SetComponentMinValueNil sets the value for ComponentMinValue to be an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) SetComponentMinValueNil() {
	o.ComponentMinValue.Set(nil)
}

// UnsetComponentMinValue ensures that no value is present for ComponentMinValue, not even an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) UnsetComponentMinValue() {
	o.ComponentMinValue.Unset()
}

// GetComponentMaxValue returns the ComponentMaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimitResourceTypeOutput) GetComponentMaxValue() int32 {
	if o == nil || isNil(o.ComponentMaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ComponentMaxValue.Get()
}

// GetComponentMaxValueOk returns a tuple with the ComponentMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimitResourceTypeOutput) GetComponentMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentMaxValue.Get(), o.ComponentMaxValue.IsSet()
}

// HasComponentMaxValue returns a boolean if a field has been set.
func (o *ResourceTypeLimitResourceTypeOutput) HasComponentMaxValue() bool {
	if o != nil && o.ComponentMaxValue.IsSet() {
		return true
	}

	return false
}

// SetComponentMaxValue gets a reference to the given NullableInt32 and assigns it to the ComponentMaxValue field.
func (o *ResourceTypeLimitResourceTypeOutput) SetComponentMaxValue(v int32) {
	o.ComponentMaxValue.Set(&v)
}
// SetComponentMaxValueNil sets the value for ComponentMaxValue to be an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) SetComponentMaxValueNil() {
	o.ComponentMaxValue.Set(nil)
}

// UnsetComponentMaxValue ensures that no value is present for ComponentMaxValue, not even an explicit nil
func (o *ResourceTypeLimitResourceTypeOutput) UnsetComponentMaxValue() {
	o.ComponentMaxValue.Unset()
}

func (o ResourceTypeLimitResourceTypeOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectMinValue.IsSet() {
		toSerialize["projectMinValue"] = o.ProjectMinValue.Get()
	}
	if o.ProjectMaxValue.IsSet() {
		toSerialize["projectMaxValue"] = o.ProjectMaxValue.Get()
	}
	if o.EnvironmentMinValue.IsSet() {
		toSerialize["environmentMinValue"] = o.EnvironmentMinValue.Get()
	}
	if o.EnvironmentMaxValue.IsSet() {
		toSerialize["environmentMaxValue"] = o.EnvironmentMaxValue.Get()
	}
	if o.ComponentMinValue.IsSet() {
		toSerialize["componentMinValue"] = o.ComponentMinValue.Get()
	}
	if o.ComponentMaxValue.IsSet() {
		toSerialize["componentMaxValue"] = o.ComponentMaxValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceTypeLimitResourceTypeOutput struct {
	value *ResourceTypeLimitResourceTypeOutput
	isSet bool
}

func (v NullableResourceTypeLimitResourceTypeOutput) Get() *ResourceTypeLimitResourceTypeOutput {
	return v.value
}

func (v *NullableResourceTypeLimitResourceTypeOutput) Set(val *ResourceTypeLimitResourceTypeOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeLimitResourceTypeOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeLimitResourceTypeOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeLimitResourceTypeOutput(val *ResourceTypeLimitResourceTypeOutput) *NullableResourceTypeLimitResourceTypeOutput {
	return &NullableResourceTypeLimitResourceTypeOutput{value: val, isSet: true}
}

func (v NullableResourceTypeLimitResourceTypeOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeLimitResourceTypeOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


