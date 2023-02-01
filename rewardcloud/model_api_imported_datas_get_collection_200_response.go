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

// ApiImportedDatasGetCollection200Response struct for ApiImportedDatasGetCollection200Response
type ApiImportedDatasGetCollection200Response struct {
	Embedded []ImportedDataJsonhal `json:"_embedded"`
	TotalItems *int32 `json:"totalItems,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	Links ApiComponentResourceLimitsGetCollection200ResponseLinks `json:"_links"`
}

// NewApiImportedDatasGetCollection200Response instantiates a new ApiImportedDatasGetCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiImportedDatasGetCollection200Response(embedded []ImportedDataJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks) *ApiImportedDatasGetCollection200Response {
	this := ApiImportedDatasGetCollection200Response{}
	this.Embedded = embedded
	this.Links = links
	return &this
}

// NewApiImportedDatasGetCollection200ResponseWithDefaults instantiates a new ApiImportedDatasGetCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiImportedDatasGetCollection200ResponseWithDefaults() *ApiImportedDatasGetCollection200Response {
	this := ApiImportedDatasGetCollection200Response{}
	return &this
}

// GetEmbedded returns the Embedded field value
func (o *ApiImportedDatasGetCollection200Response) GetEmbedded() []ImportedDataJsonhal {
	if o == nil {
		var ret []ImportedDataJsonhal
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *ApiImportedDatasGetCollection200Response) GetEmbeddedOk() ([]ImportedDataJsonhal, bool) {
	if o == nil {
    return nil, false
	}
	return o.Embedded, true
}

// SetEmbedded sets field value
func (o *ApiImportedDatasGetCollection200Response) SetEmbedded(v []ImportedDataJsonhal) {
	o.Embedded = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ApiImportedDatasGetCollection200Response) GetTotalItems() int32 {
	if o == nil || isNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiImportedDatasGetCollection200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalItems) {
    return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ApiImportedDatasGetCollection200Response) HasTotalItems() bool {
	if o != nil && !isNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ApiImportedDatasGetCollection200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *ApiImportedDatasGetCollection200Response) GetItemsPerPage() int32 {
	if o == nil || isNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiImportedDatasGetCollection200Response) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || isNil(o.ItemsPerPage) {
    return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *ApiImportedDatasGetCollection200Response) HasItemsPerPage() bool {
	if o != nil && !isNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *ApiImportedDatasGetCollection200Response) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetLinks returns the Links field value
func (o *ApiImportedDatasGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks {
	if o == nil {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiImportedDatasGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ApiImportedDatasGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks) {
	o.Links = v
}

func (o ApiImportedDatasGetCollection200Response) MarshalJSON() ([]byte, error) {
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

type NullableApiImportedDatasGetCollection200Response struct {
	value *ApiImportedDatasGetCollection200Response
	isSet bool
}

func (v NullableApiImportedDatasGetCollection200Response) Get() *ApiImportedDatasGetCollection200Response {
	return v.value
}

func (v *NullableApiImportedDatasGetCollection200Response) Set(val *ApiImportedDatasGetCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiImportedDatasGetCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiImportedDatasGetCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiImportedDatasGetCollection200Response(val *ApiImportedDatasGetCollection200Response) *NullableApiImportedDatasGetCollection200Response {
	return &NullableApiImportedDatasGetCollection200Response{value: val, isSet: true}
}

func (v NullableApiImportedDatasGetCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiImportedDatasGetCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


