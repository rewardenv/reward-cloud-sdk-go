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

// ApiComponentResourceLimitsGetCollection200ResponseLinks struct for ApiComponentResourceLimitsGetCollection200ResponseLinks
type ApiComponentResourceLimitsGetCollection200ResponseLinks struct {
	Self     *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf `json:"self,omitempty"`
	First    *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf `json:"first,omitempty"`
	Last     *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf `json:"last,omitempty"`
	Next     *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf `json:"next,omitempty"`
	Previous *ApiComponentResourceLimitsGetCollection200ResponseLinksSelf `json:"previous,omitempty"`
}

// NewApiComponentResourceLimitsGetCollection200ResponseLinks instantiates a new ApiComponentResourceLimitsGetCollection200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentResourceLimitsGetCollection200ResponseLinks() *ApiComponentResourceLimitsGetCollection200ResponseLinks {
	this := ApiComponentResourceLimitsGetCollection200ResponseLinks{}
	return &this
}

// NewApiComponentResourceLimitsGetCollection200ResponseLinksWithDefaults instantiates a new ApiComponentResourceLimitsGetCollection200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentResourceLimitsGetCollection200ResponseLinksWithDefaults() *ApiComponentResourceLimitsGetCollection200ResponseLinks {
	this := ApiComponentResourceLimitsGetCollection200ResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetSelf() ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	if o == nil || isNil(o.Self) {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetSelfOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinksSelf, bool) {
	if o == nil || isNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given ApiComponentResourceLimitsGetCollection200ResponseLinksSelf and assigns it to the Self field.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) SetSelf(v ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) {
	o.Self = &v
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetFirst() ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	if o == nil || isNil(o.First) {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetFirstOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinksSelf, bool) {
	if o == nil || isNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) HasFirst() bool {
	if o != nil && !isNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given ApiComponentResourceLimitsGetCollection200ResponseLinksSelf and assigns it to the First field.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) SetFirst(v ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetLast() ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	if o == nil || isNil(o.Last) {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetLastOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinksSelf, bool) {
	if o == nil || isNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) HasLast() bool {
	if o != nil && !isNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given ApiComponentResourceLimitsGetCollection200ResponseLinksSelf and assigns it to the Last field.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) SetLast(v ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetNext() ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	if o == nil || isNil(o.Next) {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetNextOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinksSelf, bool) {
	if o == nil || isNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) HasNext() bool {
	if o != nil && !isNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given ApiComponentResourceLimitsGetCollection200ResponseLinksSelf and assigns it to the Next field.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) SetNext(v ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetPrevious() ApiComponentResourceLimitsGetCollection200ResponseLinksSelf {
	if o == nil || isNil(o.Previous) {
		var ret ApiComponentResourceLimitsGetCollection200ResponseLinksSelf
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) GetPreviousOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinksSelf, bool) {
	if o == nil || isNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) HasPrevious() bool {
	if o != nil && !isNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given ApiComponentResourceLimitsGetCollection200ResponseLinksSelf and assigns it to the Previous field.
func (o *ApiComponentResourceLimitsGetCollection200ResponseLinks) SetPrevious(v ApiComponentResourceLimitsGetCollection200ResponseLinksSelf) {
	o.Previous = &v
}

func (o ApiComponentResourceLimitsGetCollection200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !isNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !isNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !isNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !isNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	return json.Marshal(toSerialize)
}

type NullableApiComponentResourceLimitsGetCollection200ResponseLinks struct {
	value *ApiComponentResourceLimitsGetCollection200ResponseLinks
	isSet bool
}

func (v NullableApiComponentResourceLimitsGetCollection200ResponseLinks) Get() *ApiComponentResourceLimitsGetCollection200ResponseLinks {
	return v.value
}

func (v *NullableApiComponentResourceLimitsGetCollection200ResponseLinks) Set(val *ApiComponentResourceLimitsGetCollection200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentResourceLimitsGetCollection200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentResourceLimitsGetCollection200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentResourceLimitsGetCollection200ResponseLinks(val *ApiComponentResourceLimitsGetCollection200ResponseLinks) *NullableApiComponentResourceLimitsGetCollection200ResponseLinks {
	return &NullableApiComponentResourceLimitsGetCollection200ResponseLinks{value: val, isSet: true}
}

func (v NullableApiComponentResourceLimitsGetCollection200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentResourceLimitsGetCollection200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
