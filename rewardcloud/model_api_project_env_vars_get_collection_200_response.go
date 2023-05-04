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

// checks if the ApiProjectEnvVarsGetCollection200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiProjectEnvVarsGetCollection200Response{}

// ApiProjectEnvVarsGetCollection200Response struct for ApiProjectEnvVarsGetCollection200Response
type ApiProjectEnvVarsGetCollection200Response struct {
	Embedded []ProjectEnvVarJsonhal `json:"_embedded"`
	TotalItems *int32 `json:"totalItems,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	Links ApiComponentResourceLimitsGetCollection200ResponseLinks `json:"_links"`
}

// NewApiProjectEnvVarsGetCollection200Response instantiates a new ApiProjectEnvVarsGetCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiProjectEnvVarsGetCollection200Response(embedded []ProjectEnvVarJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks) *ApiProjectEnvVarsGetCollection200Response {
	this := ApiProjectEnvVarsGetCollection200Response{}
	this.Embedded = embedded
	this.Links = links
	return &this
}

// NewApiProjectEnvVarsGetCollection200ResponseWithDefaults instantiates a new ApiProjectEnvVarsGetCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiProjectEnvVarsGetCollection200ResponseWithDefaults() *ApiProjectEnvVarsGetCollection200Response {
	this := ApiProjectEnvVarsGetCollection200Response{}
	return &this
}

// GetEmbedded returns the Embedded field value
func (o *ApiProjectEnvVarsGetCollection200Response) GetEmbedded() []ProjectEnvVarJsonhal {
	if o == nil {
		var ret []ProjectEnvVarJsonhal
		return ret
	}

	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value
// and a boolean to check if the value has been set.
func (o *ApiProjectEnvVarsGetCollection200Response) GetEmbeddedOk() ([]ProjectEnvVarJsonhal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embedded, true
}

// SetEmbedded sets field value
func (o *ApiProjectEnvVarsGetCollection200Response) SetEmbedded(v []ProjectEnvVarJsonhal) {
	o.Embedded = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *ApiProjectEnvVarsGetCollection200Response) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProjectEnvVarsGetCollection200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *ApiProjectEnvVarsGetCollection200Response) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *ApiProjectEnvVarsGetCollection200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *ApiProjectEnvVarsGetCollection200Response) GetItemsPerPage() int32 {
	if o == nil || IsNil(o.ItemsPerPage) {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiProjectEnvVarsGetCollection200Response) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemsPerPage) {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *ApiProjectEnvVarsGetCollection200Response) HasItemsPerPage() bool {
	if o != nil && !IsNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *ApiProjectEnvVarsGetCollection200Response) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetLinks returns the Links field value
func (o *ApiProjectEnvVarsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks {
	if o == nil {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApiProjectEnvVarsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ApiProjectEnvVarsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks) {
	o.Links = v
}

func (o ApiProjectEnvVarsGetCollection200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiProjectEnvVarsGetCollection200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_embedded"] = o.Embedded
	if !IsNil(o.TotalItems) {
		toSerialize["totalItems"] = o.TotalItems
	}
	if !IsNil(o.ItemsPerPage) {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	toSerialize["_links"] = o.Links
	return toSerialize, nil
}

type NullableApiProjectEnvVarsGetCollection200Response struct {
	value *ApiProjectEnvVarsGetCollection200Response
	isSet bool
}

func (v NullableApiProjectEnvVarsGetCollection200Response) Get() *ApiProjectEnvVarsGetCollection200Response {
	return v.value
}

func (v *NullableApiProjectEnvVarsGetCollection200Response) Set(val *ApiProjectEnvVarsGetCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiProjectEnvVarsGetCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiProjectEnvVarsGetCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiProjectEnvVarsGetCollection200Response(val *ApiProjectEnvVarsGetCollection200Response) *NullableApiProjectEnvVarsGetCollection200Response {
	return &NullableApiProjectEnvVarsGetCollection200Response{value: val, isSet: true}
}

func (v NullableApiProjectEnvVarsGetCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiProjectEnvVarsGetCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


