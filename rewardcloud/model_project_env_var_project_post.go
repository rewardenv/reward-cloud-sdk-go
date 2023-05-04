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

// checks if the ProjectEnvVarProjectPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEnvVarProjectPost{}

// ProjectEnvVarProjectPost 
type ProjectEnvVarProjectPost struct {
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
	IsEncrypted NullableBool `json:"isEncrypted,omitempty"`
	EnvVarType NullableString `json:"envVarType,omitempty"`
}

// NewProjectEnvVarProjectPost instantiates a new ProjectEnvVarProjectPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEnvVarProjectPost() *ProjectEnvVarProjectPost {
	this := ProjectEnvVarProjectPost{}
	return &this
}

// NewProjectEnvVarProjectPostWithDefaults instantiates a new ProjectEnvVarProjectPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEnvVarProjectPostWithDefaults() *ProjectEnvVarProjectPost {
	this := ProjectEnvVarProjectPost{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectPost) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectPost) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectPost) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ProjectEnvVarProjectPost) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProjectEnvVarProjectPost) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProjectEnvVarProjectPost) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectPost) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectPost) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectPost) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProjectEnvVarProjectPost) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ProjectEnvVarProjectPost) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProjectEnvVarProjectPost) UnsetValue() {
	o.Value.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectPost) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectPost) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectPost) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *ProjectEnvVarProjectPost) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}
// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *ProjectEnvVarProjectPost) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *ProjectEnvVarProjectPost) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarProjectPost) GetEnvVarType() string {
	if o == nil || IsNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarProjectPost) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *ProjectEnvVarProjectPost) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *ProjectEnvVarProjectPost) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}
// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *ProjectEnvVarProjectPost) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *ProjectEnvVarProjectPost) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

func (o ProjectEnvVarProjectPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEnvVarProjectPost) ToMap() (map[string]interface{}, error) {
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

type NullableProjectEnvVarProjectPost struct {
	value *ProjectEnvVarProjectPost
	isSet bool
}

func (v NullableProjectEnvVarProjectPost) Get() *ProjectEnvVarProjectPost {
	return v.value
}

func (v *NullableProjectEnvVarProjectPost) Set(val *ProjectEnvVarProjectPost) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEnvVarProjectPost) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEnvVarProjectPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEnvVarProjectPost(val *ProjectEnvVarProjectPost) *NullableProjectEnvVarProjectPost {
	return &NullableProjectEnvVarProjectPost{value: val, isSet: true}
}

func (v NullableProjectEnvVarProjectPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEnvVarProjectPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


