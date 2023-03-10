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

// OrganisationEnvVarJsonhalOrganisationGet 
type OrganisationEnvVarJsonhalOrganisationGet struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
	IsEncrypted NullableBool `json:"isEncrypted,omitempty"`
	EnvVarType NullableString `json:"envVarType,omitempty"`
}

// NewOrganisationEnvVarJsonhalOrganisationGet instantiates a new OrganisationEnvVarJsonhalOrganisationGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationEnvVarJsonhalOrganisationGet() *OrganisationEnvVarJsonhalOrganisationGet {
	this := OrganisationEnvVarJsonhalOrganisationGet{}
	return &this
}

// NewOrganisationEnvVarJsonhalOrganisationGetWithDefaults instantiates a new OrganisationEnvVarJsonhalOrganisationGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationEnvVarJsonhalOrganisationGetWithDefaults() *OrganisationEnvVarJsonhalOrganisationGet {
	this := OrganisationEnvVarJsonhalOrganisationGet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhalOrganisationGet) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetKey() string {
	if o == nil || isNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhalOrganisationGet) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhalOrganisationGet) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) UnsetValue() {
	o.Value.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetIsEncrypted() bool {
	if o == nil || isNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhalOrganisationGet) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}
// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetEnvVarType() string {
	if o == nil || isNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhalOrganisationGet) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhalOrganisationGet) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}
// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *OrganisationEnvVarJsonhalOrganisationGet) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

func (o OrganisationEnvVarJsonhalOrganisationGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsEncrypted.IsSet() {
		toSerialize["isEncrypted"] = o.IsEncrypted.Get()
	}
	if o.EnvVarType.IsSet() {
		toSerialize["envVarType"] = o.EnvVarType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganisationEnvVarJsonhalOrganisationGet struct {
	value *OrganisationEnvVarJsonhalOrganisationGet
	isSet bool
}

func (v NullableOrganisationEnvVarJsonhalOrganisationGet) Get() *OrganisationEnvVarJsonhalOrganisationGet {
	return v.value
}

func (v *NullableOrganisationEnvVarJsonhalOrganisationGet) Set(val *OrganisationEnvVarJsonhalOrganisationGet) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationEnvVarJsonhalOrganisationGet) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationEnvVarJsonhalOrganisationGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationEnvVarJsonhalOrganisationGet(val *OrganisationEnvVarJsonhalOrganisationGet) *NullableOrganisationEnvVarJsonhalOrganisationGet {
	return &NullableOrganisationEnvVarJsonhalOrganisationGet{value: val, isSet: true}
}

func (v NullableOrganisationEnvVarJsonhalOrganisationGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationEnvVarJsonhalOrganisationGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


