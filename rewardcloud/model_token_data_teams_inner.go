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

// checks if the TokenDataTeamsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenDataTeamsInner{}

// TokenDataTeamsInner struct for TokenDataTeamsInner
type TokenDataTeamsInner struct {
	Uuid *string `json:"uuid,omitempty"`
}

// NewTokenDataTeamsInner instantiates a new TokenDataTeamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenDataTeamsInner() *TokenDataTeamsInner {
	this := TokenDataTeamsInner{}
	return &this
}

// NewTokenDataTeamsInnerWithDefaults instantiates a new TokenDataTeamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDataTeamsInnerWithDefaults() *TokenDataTeamsInner {
	this := TokenDataTeamsInner{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TokenDataTeamsInner) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenDataTeamsInner) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TokenDataTeamsInner) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TokenDataTeamsInner) SetUuid(v string) {
	o.Uuid = &v
}

func (o TokenDataTeamsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenDataTeamsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableTokenDataTeamsInner struct {
	value *TokenDataTeamsInner
	isSet bool
}

func (v NullableTokenDataTeamsInner) Get() *TokenDataTeamsInner {
	return v.value
}

func (v *NullableTokenDataTeamsInner) Set(val *TokenDataTeamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenDataTeamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenDataTeamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenDataTeamsInner(val *TokenDataTeamsInner) *NullableTokenDataTeamsInner {
	return &NullableTokenDataTeamsInner{value: val, isSet: true}
}

func (v NullableTokenDataTeamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenDataTeamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


