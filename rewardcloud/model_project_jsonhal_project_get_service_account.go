/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
	"fmt"
)

// ProjectJsonhalProjectGetServiceAccount struct for ProjectJsonhalProjectGetServiceAccount
type ProjectJsonhalProjectGetServiceAccount struct {
	ServiceAccountJsonhalProjectGet *ServiceAccountJsonhalProjectGet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProjectJsonhalProjectGetServiceAccount) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ServiceAccountJsonhalProjectGet
	err = json.Unmarshal(data, &dst.ServiceAccountJsonhalProjectGet);
	if err == nil {
		jsonServiceAccountJsonhalProjectGet, _ := json.Marshal(dst.ServiceAccountJsonhalProjectGet)
		if string(jsonServiceAccountJsonhalProjectGet) == "{}" { // empty struct
			dst.ServiceAccountJsonhalProjectGet = nil
		} else {
			return nil // data stored in dst.ServiceAccountJsonhalProjectGet, return on the first match
		}
	} else {
		dst.ServiceAccountJsonhalProjectGet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProjectJsonhalProjectGetServiceAccount)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProjectJsonhalProjectGetServiceAccount) MarshalJSON() ([]byte, error) {
	if src.ServiceAccountJsonhalProjectGet != nil {
		return json.Marshal(&src.ServiceAccountJsonhalProjectGet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProjectJsonhalProjectGetServiceAccount struct {
	value *ProjectJsonhalProjectGetServiceAccount
	isSet bool
}

func (v NullableProjectJsonhalProjectGetServiceAccount) Get() *ProjectJsonhalProjectGetServiceAccount {
	return v.value
}

func (v *NullableProjectJsonhalProjectGetServiceAccount) Set(val *ProjectJsonhalProjectGetServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectJsonhalProjectGetServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectJsonhalProjectGetServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectJsonhalProjectGetServiceAccount(val *ProjectJsonhalProjectGetServiceAccount) *NullableProjectJsonhalProjectGetServiceAccount {
	return &NullableProjectJsonhalProjectGetServiceAccount{value: val, isSet: true}
}

func (v NullableProjectJsonhalProjectGetServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectJsonhalProjectGetServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


