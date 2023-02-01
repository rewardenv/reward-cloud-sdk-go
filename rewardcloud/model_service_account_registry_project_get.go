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

// ServiceAccountRegistryProjectGet 
type ServiceAccountRegistryProjectGet struct {
	Url NullableString `json:"url,omitempty"`
}

// NewServiceAccountRegistryProjectGet instantiates a new ServiceAccountRegistryProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountRegistryProjectGet() *ServiceAccountRegistryProjectGet {
	this := ServiceAccountRegistryProjectGet{}
	return &this
}

// NewServiceAccountRegistryProjectGetWithDefaults instantiates a new ServiceAccountRegistryProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountRegistryProjectGetWithDefaults() *ServiceAccountRegistryProjectGet {
	this := ServiceAccountRegistryProjectGet{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceAccountRegistryProjectGet) GetUrl() string {
	if o == nil || isNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceAccountRegistryProjectGet) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceAccountRegistryProjectGet) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ServiceAccountRegistryProjectGet) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ServiceAccountRegistryProjectGet) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ServiceAccountRegistryProjectGet) UnsetUrl() {
	o.Url.Unset()
}

func (o ServiceAccountRegistryProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAccountRegistryProjectGet struct {
	value *ServiceAccountRegistryProjectGet
	isSet bool
}

func (v NullableServiceAccountRegistryProjectGet) Get() *ServiceAccountRegistryProjectGet {
	return v.value
}

func (v *NullableServiceAccountRegistryProjectGet) Set(val *ServiceAccountRegistryProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountRegistryProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountRegistryProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountRegistryProjectGet(val *ServiceAccountRegistryProjectGet) *NullableServiceAccountRegistryProjectGet {
	return &NullableServiceAccountRegistryProjectGet{value: val, isSet: true}
}

func (v NullableServiceAccountRegistryProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountRegistryProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


