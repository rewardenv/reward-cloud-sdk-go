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

// ApiProjectTypeVersionEnvVarExamplesGetCollection200Response struct for ApiProjectTypeVersionEnvVarExamplesGetCollection200Response
type ApiProjectTypeVersionEnvVarExamplesGetCollection200Response struct {
	Embedded []ProjectTypeVersionEnvVarExampleJsonhal `json:"_embedded"`
	TotalItems *int32 `json:"totalItems,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	Links ApiComponentResourceLimitsGetCollection200ResponseLinks `json:"_links"`
}

// NewApiProjectTypeVersionEnvVarExamplesGetCollection200Response instantiates a new ApiProjectTypeVersionEnvVarExamplesGetCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiProjectTypeVersionEnvVarExamplesGetCollection200Response(embedded []ProjectTypeVersionEnvVarExampleJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks) *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response {
	this := ApiProjectTypeVersionEnvVarExamplesGetCollection200Response{}
	this.Embedded = embedded
	this.Links = links
	return &this
}

// NewApiProjectTypeVersionEnvVarExamplesGetCollection200ResponseWithDefaults instantiates a new ApiProjectTypeVersionEnvVarExamplesGetCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiProjectTypeVersionEnvVarExamplesGetCollection200ResponseWithDefaults() *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response {
	this := ApiProjectTypeVersionEnvVarExamplesGetCollection200Response{}
	return &this
}

// GetEmbedded returns the Embedded field value
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetEmbedded() []ProjectTypeVersionEnvVarExampleJsonhal {
	if o == nil {
		var ret []ProjectTypeVersionEnvVarExampleJsonhal
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetEmbeddedOk() ([]ProjectTypeVersionEnvVarExampleJsonhal, bool) {
	if o == nil {
    return nil, false
	}
	return o.Embedded, true
}

// SetEmbedded sets field value
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) SetEmbedded(v []ProjectTypeVersionEnvVarExampleJsonhal) {
	o.Embedded = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetTotalItems() int32 {
	if o == nil || isNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalItems) {
    return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) HasTotalItems() bool {
	if o != nil && !isNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetItemsPerPage() int32 {
	if o == nil || isNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || isNil(o.ItemsPerPage) {
    return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) HasItemsPerPage() bool {
	if o != nil && !isNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetLinks returns the Links field value
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks {
	if o == nil {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks) {
	o.Links = v
}

func (o ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["_embedded"] = o.Embedded
	}
	if !isNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !isNil(o.ItemsPerPage) {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response struct {
	value *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response
	isSet bool
}

func (v NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response) Get() *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response {
	return v.value
}

func (v *NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response) Set(val *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response(val *ApiProjectTypeVersionEnvVarExamplesGetCollection200Response) *NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response {
	return &NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response{value: val, isSet: true}
}

func (v NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiProjectTypeVersionEnvVarExamplesGetCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


