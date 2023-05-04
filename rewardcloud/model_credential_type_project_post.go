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

// CredentialTypeProjectPost
type CredentialTypeProjectPost struct {
	Uuid NullableString `json:"uuid,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewCredentialTypeProjectPost instantiates a new CredentialTypeProjectPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialTypeProjectPost() *CredentialTypeProjectPost {
	this := CredentialTypeProjectPost{}
	return &this
}

// NewCredentialTypeProjectPostWithDefaults instantiates a new CredentialTypeProjectPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialTypeProjectPostWithDefaults() *CredentialTypeProjectPost {
	this := CredentialTypeProjectPost{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialTypeProjectPost) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialTypeProjectPost) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *CredentialTypeProjectPost) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *CredentialTypeProjectPost) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *CredentialTypeProjectPost) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *CredentialTypeProjectPost) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialTypeProjectPost) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialTypeProjectPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CredentialTypeProjectPost) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CredentialTypeProjectPost) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CredentialTypeProjectPost) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CredentialTypeProjectPost) UnsetName() {
	o.Name.Unset()
}

func (o CredentialTypeProjectPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialTypeProjectPost struct {
	value *CredentialTypeProjectPost
	isSet bool
}

func (v NullableCredentialTypeProjectPost) Get() *CredentialTypeProjectPost {
	return v.value
}

func (v *NullableCredentialTypeProjectPost) Set(val *CredentialTypeProjectPost) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialTypeProjectPost) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialTypeProjectPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialTypeProjectPost(val *CredentialTypeProjectPost) *NullableCredentialTypeProjectPost {
	return &NullableCredentialTypeProjectPost{value: val, isSet: true}
}

func (v NullableCredentialTypeProjectPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialTypeProjectPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
