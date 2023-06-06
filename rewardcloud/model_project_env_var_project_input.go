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

// checks if the ProjectEnvVarProjectInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEnvVarProjectInput{}

// ProjectEnvVarProjectInput Class ProjectEnvVar
type ProjectEnvVarProjectInput struct {
	Key         NullableString `json:"key,omitempty"`
	Value       NullableString `json:"value,omitempty"`
	IsEncrypted NullableBool   `json:"isEncrypted,omitempty"`
	EnvVarType  NullableString `json:"envVarType,omitempty"`
}

// NewProjectEnvVarProjectInput instantiates a new ProjectEnvVarProjectInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEnvVarProjectInput() *ProjectEnvVarProjectInput {
	this := ProjectEnvVarProjectInput{}
	return &this
}

// NewProjectEnvVarProjectInputWithDefaults instantiates a new ProjectEnvVarProjectInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEnvVarProjectInputWithDefaults() *ProjectEnvVarProjectInput {
	this := ProjectEnvVarProjectInput{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectInput) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectInput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectInput) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ProjectEnvVarProjectInput) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProjectEnvVarProjectInput) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProjectEnvVarProjectInput) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectInput) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectInput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectInput) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProjectEnvVarProjectInput) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ProjectEnvVarProjectInput) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProjectEnvVarProjectInput) UnsetValue() {
	o.Value.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectInput) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectInput) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectInput) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *ProjectEnvVarProjectInput) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}

// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *ProjectEnvVarProjectInput) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *ProjectEnvVarProjectInput) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectInput) GetEnvVarType() string {
	if o == nil || IsNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectInput) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectInput) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *ProjectEnvVarProjectInput) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}

// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *ProjectEnvVarProjectInput) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *ProjectEnvVarProjectInput) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

func (o ProjectEnvVarProjectInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEnvVarProjectInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsEncrypted.IsSet() {
		toSerialize["isEncrypted"] = o.IsEncrypted.Get()
	}
	if o.EnvVarType.IsSet() {
		toSerialize["envVarType"] = o.EnvVarType.Get()
	}
	return toSerialize, nil
}

type NullableProjectEnvVarProjectInput struct {
	value *ProjectEnvVarProjectInput
	isSet bool
}

func (v NullableProjectEnvVarProjectInput) Get() *ProjectEnvVarProjectInput {
	return v.value
}

func (v *NullableProjectEnvVarProjectInput) Set(val *ProjectEnvVarProjectInput) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEnvVarProjectInput) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEnvVarProjectInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEnvVarProjectInput(val *ProjectEnvVarProjectInput) *NullableProjectEnvVarProjectInput {
	return &NullableProjectEnvVarProjectInput{value: val, isSet: true}
}

func (v NullableProjectEnvVarProjectInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEnvVarProjectInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}