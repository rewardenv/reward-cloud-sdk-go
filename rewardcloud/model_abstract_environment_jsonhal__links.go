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

// checks if the AbstractEnvironmentJsonhalLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbstractEnvironmentJsonhalLinks{}

// AbstractEnvironmentJsonhalLinks struct for AbstractEnvironmentJsonhalLinks
type AbstractEnvironmentJsonhalLinks struct {
	Self *AbstractEnvironmentJsonhalLinksSelf `json:"self,omitempty"`
}

// NewAbstractEnvironmentJsonhalLinks instantiates a new AbstractEnvironmentJsonhalLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractEnvironmentJsonhalLinks() *AbstractEnvironmentJsonhalLinks {
	this := AbstractEnvironmentJsonhalLinks{}
	return &this
}

// NewAbstractEnvironmentJsonhalLinksWithDefaults instantiates a new AbstractEnvironmentJsonhalLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractEnvironmentJsonhalLinksWithDefaults() *AbstractEnvironmentJsonhalLinks {
	this := AbstractEnvironmentJsonhalLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AbstractEnvironmentJsonhalLinks) GetSelf() AbstractEnvironmentJsonhalLinksSelf {
	if o == nil || IsNil(o.Self) {
		var ret AbstractEnvironmentJsonhalLinksSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractEnvironmentJsonhalLinks) GetSelfOk() (*AbstractEnvironmentJsonhalLinksSelf, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AbstractEnvironmentJsonhalLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given AbstractEnvironmentJsonhalLinksSelf and assigns it to the Self field.
func (o *AbstractEnvironmentJsonhalLinks) SetSelf(v AbstractEnvironmentJsonhalLinksSelf) {
	o.Self = &v
}

func (o AbstractEnvironmentJsonhalLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractEnvironmentJsonhalLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableAbstractEnvironmentJsonhalLinks struct {
	value *AbstractEnvironmentJsonhalLinks
	isSet bool
}

func (v NullableAbstractEnvironmentJsonhalLinks) Get() *AbstractEnvironmentJsonhalLinks {
	return v.value
}

func (v *NullableAbstractEnvironmentJsonhalLinks) Set(val *AbstractEnvironmentJsonhalLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractEnvironmentJsonhalLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractEnvironmentJsonhalLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractEnvironmentJsonhalLinks(val *AbstractEnvironmentJsonhalLinks) *NullableAbstractEnvironmentJsonhalLinks {
	return &NullableAbstractEnvironmentJsonhalLinks{value: val, isSet: true}
}

func (v NullableAbstractEnvironmentJsonhalLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractEnvironmentJsonhalLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
