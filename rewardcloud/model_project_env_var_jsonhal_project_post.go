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

// checks if the ProjectEnvVarJsonhalProjectPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectEnvVarJsonhalProjectPost{}

// ProjectEnvVarJsonhalProjectPost
type ProjectEnvVarJsonhalProjectPost struct {
	Links       *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Key         NullableString                   `json:"key,omitempty"`
	Value       NullableString                   `json:"value,omitempty"`
	IsEncrypted NullableBool                     `json:"isEncrypted,omitempty"`
	EnvVarType  NullableString                   `json:"envVarType,omitempty"`
}

// NewProjectEnvVarJsonhalProjectPost instantiates a new ProjectEnvVarJsonhalProjectPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectEnvVarJsonhalProjectPost() *ProjectEnvVarJsonhalProjectPost {
	this := ProjectEnvVarJsonhalProjectPost{}
	return &this
}

// NewProjectEnvVarJsonhalProjectPostWithDefaults instantiates a new ProjectEnvVarJsonhalProjectPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectEnvVarJsonhalProjectPostWithDefaults() *ProjectEnvVarJsonhalProjectPost {
	this := ProjectEnvVarJsonhalProjectPost{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProjectEnvVarJsonhalProjectPost) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectEnvVarJsonhalProjectPost) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProjectEnvVarJsonhalProjectPost) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *ProjectEnvVarJsonhalProjectPost) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarJsonhalProjectPost) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarJsonhalProjectPost) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectEnvVarJsonhalProjectPost) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ProjectEnvVarJsonhalProjectPost) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarJsonhalProjectPost) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarJsonhalProjectPost) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectEnvVarJsonhalProjectPost) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProjectEnvVarJsonhalProjectPost) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) UnsetValue() {
	o.Value.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarJsonhalProjectPost) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarJsonhalProjectPost) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *ProjectEnvVarJsonhalProjectPost) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *ProjectEnvVarJsonhalProjectPost) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}

// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectEnvVarJsonhalProjectPost) GetEnvVarType() string {
	if o == nil || IsNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectEnvVarJsonhalProjectPost) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *ProjectEnvVarJsonhalProjectPost) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *ProjectEnvVarJsonhalProjectPost) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}

// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *ProjectEnvVarJsonhalProjectPost) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

func (o ProjectEnvVarJsonhalProjectPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectEnvVarJsonhalProjectPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
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

type NullableProjectEnvVarJsonhalProjectPost struct {
	value *ProjectEnvVarJsonhalProjectPost
	isSet bool
}

func (v NullableProjectEnvVarJsonhalProjectPost) Get() *ProjectEnvVarJsonhalProjectPost {
	return v.value
}

func (v *NullableProjectEnvVarJsonhalProjectPost) Set(val *ProjectEnvVarJsonhalProjectPost) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEnvVarJsonhalProjectPost) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEnvVarJsonhalProjectPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEnvVarJsonhalProjectPost(val *ProjectEnvVarJsonhalProjectPost) *NullableProjectEnvVarJsonhalProjectPost {
	return &NullableProjectEnvVarJsonhalProjectPost{value: val, isSet: true}
}

func (v NullableProjectEnvVarJsonhalProjectPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEnvVarJsonhalProjectPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
