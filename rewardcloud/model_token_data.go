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

// TokenData struct for TokenData
type TokenData struct {
	Id *int32 `json:"id,omitempty"`
	Fullname *string `json:"fullname,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Teams []TokenDataTeamsInner `json:"teams,omitempty"`
}

// NewTokenData instantiates a new TokenData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenData() *TokenData {
	this := TokenData{}
	return &this
}

// NewTokenDataWithDefaults instantiates a new TokenData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDataWithDefaults() *TokenData {
	this := TokenData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenData) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenData) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TokenData) SetId(v int32) {
	o.Id = &v
}

// GetFullname returns the Fullname field value if set, zero value otherwise.
func (o *TokenData) GetFullname() string {
	if o == nil || isNil(o.Fullname) {
		var ret string
		return ret
	}
	return *o.Fullname
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetFullnameOk() (*string, bool) {
	if o == nil || isNil(o.Fullname) {
    return nil, false
	}
	return o.Fullname, true
}

// HasFullname returns a boolean if a field has been set.
func (o *TokenData) HasFullname() bool {
	if o != nil && !isNil(o.Fullname) {
		return true
	}

	return false
}

// SetFullname gets a reference to the given string and assigns it to the Fullname field.
func (o *TokenData) SetFullname(v string) {
	o.Fullname = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *TokenData) GetRoles() []string {
	if o == nil || isNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetRolesOk() ([]string, bool) {
	if o == nil || isNil(o.Roles) {
    return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *TokenData) HasRoles() bool {
	if o != nil && !isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *TokenData) SetRoles(v []string) {
	o.Roles = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *TokenData) GetTeams() []TokenDataTeamsInner {
	if o == nil || isNil(o.Teams) {
		var ret []TokenDataTeamsInner
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetTeamsOk() ([]TokenDataTeamsInner, bool) {
	if o == nil || isNil(o.Teams) {
    return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *TokenData) HasTeams() bool {
	if o != nil && !isNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []TokenDataTeamsInner and assigns it to the Teams field.
func (o *TokenData) SetTeams(v []TokenDataTeamsInner) {
	o.Teams = v
}

func (o TokenData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Fullname) {
		toSerialize["fullname"] = o.Fullname
	}
	if !isNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !isNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return json.Marshal(toSerialize)
}

type NullableTokenData struct {
	value *TokenData
	isSet bool
}

func (v NullableTokenData) Get() *TokenData {
	return v.value
}

func (v *NullableTokenData) Set(val *TokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenData(val *TokenData) *NullableTokenData {
	return &NullableTokenData{value: val, isSet: true}
}

func (v NullableTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


