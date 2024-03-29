/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// checks if the ProjectEnvVarTemplateProjectInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEnvVarTemplateProjectInput{}

// ProjectEnvVarTemplateProjectInput Class ProjectEnvVar
type ProjectEnvVarTemplateProjectInput struct {
	Key        NullableString `json:"key,omitempty"`
	Value      NullableString `json:"value,omitempty"`
	EnvVarType NullableString `json:"envVarType,omitempty"`
}

// NewProjectEnvVarTemplateProjectInput instantiates a new ProjectEnvVarTemplateProjectInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEnvVarTemplateProjectInput() *ProjectEnvVarTemplateProjectInput {
	this := ProjectEnvVarTemplateProjectInput{}
	return &this
}

// NewProjectEnvVarTemplateProjectInputWithDefaults instantiates a new ProjectEnvVarTemplateProjectInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEnvVarTemplateProjectInputWithDefaults() *ProjectEnvVarTemplateProjectInput {
	this := ProjectEnvVarTemplateProjectInput{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarTemplateProjectInput) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarTemplateProjectInput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectEnvVarTemplateProjectInput) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ProjectEnvVarTemplateProjectInput) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProjectEnvVarTemplateProjectInput) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProjectEnvVarTemplateProjectInput) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarTemplateProjectInput) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarTemplateProjectInput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectEnvVarTemplateProjectInput) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProjectEnvVarTemplateProjectInput) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ProjectEnvVarTemplateProjectInput) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProjectEnvVarTemplateProjectInput) UnsetValue() {
	o.Value.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarTemplateProjectInput) GetEnvVarType() string {
	if o == nil || IsNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarTemplateProjectInput) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *ProjectEnvVarTemplateProjectInput) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *ProjectEnvVarTemplateProjectInput) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}

// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *ProjectEnvVarTemplateProjectInput) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *ProjectEnvVarTemplateProjectInput) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

func (o ProjectEnvVarTemplateProjectInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEnvVarTemplateProjectInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.EnvVarType.IsSet() {
		toSerialize["envVarType"] = o.EnvVarType.Get()
	}
	return toSerialize, nil
}

type NullableProjectEnvVarTemplateProjectInput struct {
	value *ProjectEnvVarTemplateProjectInput
	isSet bool
}

func (v NullableProjectEnvVarTemplateProjectInput) Get() *ProjectEnvVarTemplateProjectInput {
	return v.value
}

func (v *NullableProjectEnvVarTemplateProjectInput) Set(val *ProjectEnvVarTemplateProjectInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEnvVarTemplateProjectInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEnvVarTemplateProjectInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEnvVarTemplateProjectInput(val *ProjectEnvVarTemplateProjectInput) *NullableProjectEnvVarTemplateProjectInput {
	return &NullableProjectEnvVarTemplateProjectInput{value: val, isSet: true}
}

func (v NullableProjectEnvVarTemplateProjectInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEnvVarTemplateProjectInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
