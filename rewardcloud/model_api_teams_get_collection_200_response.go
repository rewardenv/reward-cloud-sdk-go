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

// ApiTeamsGetCollection200Response struct for ApiTeamsGetCollection200Response
type ApiTeamsGetCollection200Response struct {
	Embedded []TeamJsonhalTeamGet `json:"_embedded"`
	TotalItems *int32 `json:"totalItems,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	Links ApiComponentResourceLimitsGetCollection200ResponseLinks `json:"_links"`
}

// NewApiTeamsGetCollection200Response instantiates a new ApiTeamsGetCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTeamsGetCollection200Response(embedded []TeamJsonhalTeamGet, links ApiComponentResourceLimitsGetCollection200ResponseLinks) *ApiTeamsGetCollection200Response {
	this := ApiTeamsGetCollection200Response{}
	this.Embedded = embedded
	this.Links = links
	return &this
}

// NewApiTeamsGetCollection200ResponseWithDefaults instantiates a new ApiTeamsGetCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTeamsGetCollection200ResponseWithDefaults() *ApiTeamsGetCollection200Response {
	this := ApiTeamsGetCollection200Response{}
	return &this
}

// GetEmbedded returns the Embedded field value
func (o *ApiTeamsGetCollection200Response) GetEmbedded() []TeamJsonhalTeamGet {
	if o == nil {
		var ret []TeamJsonhalTeamGet
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *ApiTeamsGetCollection200Response) GetEmbeddedOk() ([]TeamJsonhalTeamGet, bool) {
	if o == nil {
    return nil, false
	}
	return o.Embedded, true
}

// SetEmbedded sets field value
func (o *ApiTeamsGetCollection200Response) SetEmbedded(v []TeamJsonhalTeamGet) {
	o.Embedded = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ApiTeamsGetCollection200Response) GetTotalItems() int32 {
	if o == nil || isNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTeamsGetCollection200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || isNil(o.TotalItems) {
    return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ApiTeamsGetCollection200Response) HasTotalItems() bool {
	if o != nil && !isNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ApiTeamsGetCollection200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *ApiTeamsGetCollection200Response) GetItemsPerPage() int32 {
	if o == nil || isNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTeamsGetCollection200Response) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || isNil(o.ItemsPerPage) {
    return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *ApiTeamsGetCollection200Response) HasItemsPerPage() bool {
	if o != nil && !isNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *ApiTeamsGetCollection200Response) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetLinks returns the Links field value
func (o *ApiTeamsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks {
	if o == nil {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiTeamsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ApiTeamsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks) {
	o.Links = v
}

func (o ApiTeamsGetCollection200Response) MarshalJSON() ([]byte, error) {
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

type NullableApiTeamsGetCollection200Response struct {
	value *ApiTeamsGetCollection200Response
	isSet bool
}

func (v NullableApiTeamsGetCollection200Response) Get() *ApiTeamsGetCollection200Response {
	return v.value
}

func (v *NullableApiTeamsGetCollection200Response) Set(val *ApiTeamsGetCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTeamsGetCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTeamsGetCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTeamsGetCollection200Response(val *ApiTeamsGetCollection200Response) *NullableApiTeamsGetCollection200Response {
	return &NullableApiTeamsGetCollection200Response{value: val, isSet: true}
}

func (v NullableApiTeamsGetCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTeamsGetCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


