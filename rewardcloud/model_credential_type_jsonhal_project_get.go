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

// CredentialTypeJsonhalProjectGet
type CredentialTypeJsonhalProjectGet struct {
	Links *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Uuid  NullableString                   `json:"uuid,omitempty"`
	Name  NullableString                   `json:"name,omitempty"`
}

// NewCredentialTypeJsonhalProjectGet instantiates a new CredentialTypeJsonhalProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialTypeJsonhalProjectGet() *CredentialTypeJsonhalProjectGet {
	this := CredentialTypeJsonhalProjectGet{}
	return &this
}

// NewCredentialTypeJsonhalProjectGetWithDefaults instantiates a new CredentialTypeJsonhalProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialTypeJsonhalProjectGetWithDefaults() *CredentialTypeJsonhalProjectGet {
	this := CredentialTypeJsonhalProjectGet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CredentialTypeJsonhalProjectGet) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeJsonhalProjectGet) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CredentialTypeJsonhalProjectGet) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *CredentialTypeJsonhalProjectGet) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialTypeJsonhalProjectGet) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialTypeJsonhalProjectGet) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *CredentialTypeJsonhalProjectGet) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *CredentialTypeJsonhalProjectGet) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *CredentialTypeJsonhalProjectGet) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *CredentialTypeJsonhalProjectGet) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialTypeJsonhalProjectGet) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialTypeJsonhalProjectGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CredentialTypeJsonhalProjectGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CredentialTypeJsonhalProjectGet) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CredentialTypeJsonhalProjectGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CredentialTypeJsonhalProjectGet) UnsetName() {
	o.Name.Unset()
}

func (o CredentialTypeJsonhalProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCredentialTypeJsonhalProjectGet struct {
	value *CredentialTypeJsonhalProjectGet
	isSet bool
}

func (v NullableCredentialTypeJsonhalProjectGet) Get() *CredentialTypeJsonhalProjectGet {
	return v.value
}

func (v *NullableCredentialTypeJsonhalProjectGet) Set(val *CredentialTypeJsonhalProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialTypeJsonhalProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialTypeJsonhalProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialTypeJsonhalProjectGet(val *CredentialTypeJsonhalProjectGet) *NullableCredentialTypeJsonhalProjectGet {
	return &NullableCredentialTypeJsonhalProjectGet{value: val, isSet: true}
}

func (v NullableCredentialTypeJsonhalProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialTypeJsonhalProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
