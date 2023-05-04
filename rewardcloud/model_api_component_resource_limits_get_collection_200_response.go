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

// ApiComponentResourceLimitsGetCollection200Response struct for ApiComponentResourceLimitsGetCollection200Response
type ApiComponentResourceLimitsGetCollection200Response struct {
	Embedded     []ComponentResourceLimitJsonhal                         `json:"_embedded"`
	TotalItems   *int32                                                  `json:"totalItems,omitempty"`
	ItemsPerPage *int32                                                  `json:"itemsPerPage,omitempty"`
	Links        ApiComponentResourceLimitsGetCollection200ResponseLinks `json:"_links"`
}

// NewApiComponentResourceLimitsGetCollection200Response instantiates a new ApiComponentResourceLimitsGetCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentResourceLimitsGetCollection200Response(embedded []ComponentResourceLimitJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks) *ApiComponentResourceLimitsGetCollection200Response {
	this := ApiComponentResourceLimitsGetCollection200Response{}
	this.Embedded = embedded
	this.Links = links
	return &this
}

// NewApiComponentResourceLimitsGetCollection200ResponseWithDefaults instantiates a new ApiComponentResourceLimitsGetCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentResourceLimitsGetCollection200ResponseWithDefaults() *ApiComponentResourceLimitsGetCollection200Response {
	this := ApiComponentResourceLimitsGetCollection200Response{}
	return &this
}

// GetEmbedded returns the Embedded field value
func (o *ApiComponentResourceLimitsGetCollection200Response) GetEmbedded() []ComponentResourceLimitJsonhal {
	if o == nil {
		var ret []ComponentResourceLimitJsonhal
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200Response) GetEmbeddedOk() ([]ComponentResourceLimitJsonhal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embedded, true
}

// SetEmbedded sets field value
func (o *ApiComponentResourceLimitsGetCollection200Response) SetEmbedded(v []ComponentResourceLimitJsonhal) {
	o.Embedded = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200Response) GetTotalItems() int32 {
	if o == nil || isNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200Response) HasTotalItems() bool {
	if o != nil && !isNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ApiComponentResourceLimitsGetCollection200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200Response) GetItemsPerPage() int32 {
	if o == nil || isNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200Response) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || isNil(o.ItemsPerPage) {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200Response) HasItemsPerPage() bool {
	if o != nil && !isNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *ApiComponentResourceLimitsGetCollection200Response) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetLinks returns the Links field value
func (o *ApiComponentResourceLimitsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks {
	if o == nil {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ApiComponentResourceLimitsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks) {
	o.Links = v
}

func (o ApiComponentResourceLimitsGetCollection200Response) MarshalJSON() ([]byte, error) {
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

type NullableApiComponentResourceLimitsGetCollection200Response struct {
	value *ApiComponentResourceLimitsGetCollection200Response
	isSet bool
}

func (v NullableApiComponentResourceLimitsGetCollection200Response) Get() *ApiComponentResourceLimitsGetCollection200Response {
	return v.value
}

func (v *NullableApiComponentResourceLimitsGetCollection200Response) Set(val *ApiComponentResourceLimitsGetCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentResourceLimitsGetCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentResourceLimitsGetCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentResourceLimitsGetCollection200Response(val *ApiComponentResourceLimitsGetCollection200Response) *NullableApiComponentResourceLimitsGetCollection200Response {
	return &NullableApiComponentResourceLimitsGetCollection200Response{value: val, isSet: true}
}

func (v NullableApiComponentResourceLimitsGetCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentResourceLimitsGetCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
