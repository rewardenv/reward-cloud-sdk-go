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

// ApiComponentResourceLimitsGetCollection200ResponseLinksSelf struct for ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
type ApiComponentResourceLimitsGetCollection200ResponseLinksSelf struct {
	Href *string `json:"href,omitempty"`
}

// NewApiComponentResourceLimitsGetCollection200ResponseLinksSelf instantiates a new ApiComponentResourceLimitsGetCollection200ResponseLinksSelf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentResourceLimitsGetCollection200ResponseLinksSelf() *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	this := ApiComponentResourceLimitsGetCollection200ResponseLinksSelf{}
	return &this
}

// NewApiComponentResourceLimitsGetCollection200ResponseLinksSelfWithDefaults instantiates a new ApiComponentResourceLimitsGetCollection200ResponseLinksSelf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentResourceLimitsGetCollection200ResponseLinksSelfWithDefaults() *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	this := ApiComponentResourceLimitsGetCollection200ResponseLinksSelf{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
    return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) SetHref(v string) {
	o.Href = &v
}

func (o ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf struct {
	value *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
	isSet bool
}

func (v NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf) Get() *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	return v.value
}

func (v *NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf) Set(val *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf(val *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) *NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	return &NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf{value: val, isSet: true}
}

func (v NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentResourceLimitsGetCollection200ResponseLinksSelf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


