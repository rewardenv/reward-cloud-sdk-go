/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// checks if the PostValidateEnvVarRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostValidateEnvVarRequest{}

// PostValidateEnvVarRequest struct for PostValidateEnvVarRequest
type PostValidateEnvVarRequest struct {
	Word *string `json:"word,omitempty"`
	Iri  *string `json:"iri,omitempty"`
}

// NewPostValidateEnvVarRequest instantiates a new PostValidateEnvVarRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostValidateEnvVarRequest() *PostValidateEnvVarRequest {
	this := PostValidateEnvVarRequest{}
	return &this
}

// NewPostValidateEnvVarRequestWithDefaults instantiates a new PostValidateEnvVarRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostValidateEnvVarRequestWithDefaults() *PostValidateEnvVarRequest {
	this := PostValidateEnvVarRequest{}
	return &this
}

// GetWord returns the Word field value if set, zero value otherwise.
func (o *PostValidateEnvVarRequest) GetWord() string {
	if o == nil || IsNil(o.Word) {
		var ret string
		return ret
	}
	return *o.Word
}

// GetWordOk returns a tuple with the Word field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostValidateEnvVarRequest) GetWordOk() (*string, bool) {
	if o == nil || IsNil(o.Word) {
		return nil, false
	}
	return o.Word, true
}

// HasWord returns a boolean if a field has been set.
func (o *PostValidateEnvVarRequest) HasWord() bool {
	if o != nil && !IsNil(o.Word) {
		return true
	}

	return false
}

// SetWord gets a reference to the given string and assigns it to the Word field.
func (o *PostValidateEnvVarRequest) SetWord(v string) {
	o.Word = &v
}

// GetIri returns the Iri field value if set, zero value otherwise.
func (o *PostValidateEnvVarRequest) GetIri() string {
	if o == nil || IsNil(o.Iri) {
		var ret string
		return ret
	}
	return *o.Iri
}

// GetIriOk returns a tuple with the Iri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostValidateEnvVarRequest) GetIriOk() (*string, bool) {
	if o == nil || IsNil(o.Iri) {
		return nil, false
	}
	return o.Iri, true
}

// HasIri returns a boolean if a field has been set.
func (o *PostValidateEnvVarRequest) HasIri() bool {
	if o != nil && !IsNil(o.Iri) {
		return true
	}

	return false
}

// SetIri gets a reference to the given string and assigns it to the Iri field.
func (o *PostValidateEnvVarRequest) SetIri(v string) {
	o.Iri = &v
}

func (o PostValidateEnvVarRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostValidateEnvVarRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Word) {
		toSerialize["word"] = o.Word
	}
	if !IsNil(o.Iri) {
		toSerialize["iri"] = o.Iri
	}
	return toSerialize, nil
}

type NullablePostValidateEnvVarRequest struct {
	value *PostValidateEnvVarRequest
	isSet bool
}

func (v NullablePostValidateEnvVarRequest) Get() *PostValidateEnvVarRequest {
	return v.value
}

func (v *NullablePostValidateEnvVarRequest) Set(val *PostValidateEnvVarRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostValidateEnvVarRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostValidateEnvVarRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostValidateEnvVarRequest(val *PostValidateEnvVarRequest) *NullablePostValidateEnvVarRequest {
	return &NullablePostValidateEnvVarRequest{value: val, isSet: true}
}

func (v NullablePostValidateEnvVarRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostValidateEnvVarRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
