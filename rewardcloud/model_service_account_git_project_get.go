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

// checks if the ServiceAccountGitProjectGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountGitProjectGet{}

// ServiceAccountGitProjectGet 
type ServiceAccountGitProjectGet struct {
	IsExternal NullableBool `json:"isExternal,omitempty"`
	Url NullableString `json:"url,omitempty"`
	SshPrivateKey NullableString `json:"sshPrivateKey,omitempty"`
}

// NewServiceAccountGitProjectGet instantiates a new ServiceAccountGitProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountGitProjectGet() *ServiceAccountGitProjectGet {
	this := ServiceAccountGitProjectGet{}
	return &this
}

// NewServiceAccountGitProjectGetWithDefaults instantiates a new ServiceAccountGitProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountGitProjectGetWithDefaults() *ServiceAccountGitProjectGet {
	this := ServiceAccountGitProjectGet{}
	return &this
}

// GetIsExternal returns the IsExternal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountGitProjectGet) GetIsExternal() bool {
	if o == nil || IsNil(o.IsExternal.Get()) {
		var ret bool
		return ret
	}
	return *o.IsExternal.Get()
}

// GetIsExternalOk returns a tuple with the IsExternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountGitProjectGet) GetIsExternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsExternal.Get(), o.IsExternal.IsSet()
}

// HasIsExternal returns a boolean if a field has been set.
func (o *ServiceAccountGitProjectGet) HasIsExternal() bool {
	if o != nil && o.IsExternal.IsSet() {
		return true
	}

	return false
}

// SetIsExternal gets a reference to the given NullableBool and assigns it to the IsExternal field.
func (o *ServiceAccountGitProjectGet) SetIsExternal(v bool) {
	o.IsExternal.Set(&v)
}
// SetIsExternalNil sets the value for IsExternal to be an explicit nil
func (o *ServiceAccountGitProjectGet) SetIsExternalNil() {
	o.IsExternal.Set(nil)
}

// UnsetIsExternal ensures that no value is present for IsExternal, not even an explicit nil
func (o *ServiceAccountGitProjectGet) UnsetIsExternal() {
	o.IsExternal.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountGitProjectGet) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountGitProjectGet) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceAccountGitProjectGet) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ServiceAccountGitProjectGet) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ServiceAccountGitProjectGet) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ServiceAccountGitProjectGet) UnsetUrl() {
	o.Url.Unset()
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountGitProjectGet) GetSshPrivateKey() string {
	if o == nil || IsNil(o.SshPrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey.Get()
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountGitProjectGet) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPrivateKey.Get(), o.SshPrivateKey.IsSet()
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *ServiceAccountGitProjectGet) HasSshPrivateKey() bool {
	if o != nil && o.SshPrivateKey.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given NullableString and assigns it to the SshPrivateKey field.
func (o *ServiceAccountGitProjectGet) SetSshPrivateKey(v string) {
	o.SshPrivateKey.Set(&v)
}
// SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil
func (o *ServiceAccountGitProjectGet) SetSshPrivateKeyNil() {
	o.SshPrivateKey.Set(nil)
}

// UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
func (o *ServiceAccountGitProjectGet) UnsetSshPrivateKey() {
	o.SshPrivateKey.Unset()
}

func (o ServiceAccountGitProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountGitProjectGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsExternal.IsSet() {
		toSerialize["isExternal"] = o.IsExternal.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.SshPrivateKey.IsSet() {
		toSerialize["sshPrivateKey"] = o.SshPrivateKey.Get()
	}
	return toSerialize, nil
}

type NullableServiceAccountGitProjectGet struct {
	value *ServiceAccountGitProjectGet
	isSet bool
}

func (v NullableServiceAccountGitProjectGet) Get() *ServiceAccountGitProjectGet {
	return v.value
}

func (v *NullableServiceAccountGitProjectGet) Set(val *ServiceAccountGitProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountGitProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountGitProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountGitProjectGet(val *ServiceAccountGitProjectGet) *NullableServiceAccountGitProjectGet {
	return &NullableServiceAccountGitProjectGet{value: val, isSet: true}
}

func (v NullableServiceAccountGitProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountGitProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


