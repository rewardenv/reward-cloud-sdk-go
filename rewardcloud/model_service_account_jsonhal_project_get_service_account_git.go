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

// ServiceAccountJsonhalProjectGetServiceAccountGit struct for ServiceAccountJsonhalProjectGetServiceAccountGit
type ServiceAccountJsonhalProjectGetServiceAccountGit struct {
	ServiceAccountGitJsonhalProjectGet *ServiceAccountGitJsonhalProjectGet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceAccountJsonhalProjectGetServiceAccountGit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ServiceAccountGitJsonhalProjectGet
	err = json.Unmarshal(data, &dst.ServiceAccountGitJsonhalProjectGet);
	if err == nil {
		jsonServiceAccountGitJsonhalProjectGet, _ := json.Marshal(dst.ServiceAccountGitJsonhalProjectGet)
		if string(jsonServiceAccountGitJsonhalProjectGet) == "{}" { // empty struct
			dst.ServiceAccountGitJsonhalProjectGet = nil
		} else {
			return nil // data stored in dst.ServiceAccountGitJsonhalProjectGet, return on the first match
		}
	} else {
		dst.ServiceAccountGitJsonhalProjectGet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceAccountJsonhalProjectGetServiceAccountGit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceAccountJsonhalProjectGetServiceAccountGit) MarshalJSON() ([]byte, error) {
	if src.ServiceAccountGitJsonhalProjectGet != nil {
		return json.Marshal(&src.ServiceAccountGitJsonhalProjectGet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceAccountJsonhalProjectGetServiceAccountGit struct {
	value *ServiceAccountJsonhalProjectGetServiceAccountGit
	isSet bool
}

func (v NullableServiceAccountJsonhalProjectGetServiceAccountGit) Get() *ServiceAccountJsonhalProjectGetServiceAccountGit {
	return v.value
}

func (v *NullableServiceAccountJsonhalProjectGetServiceAccountGit) Set(val *ServiceAccountJsonhalProjectGetServiceAccountGit) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountJsonhalProjectGetServiceAccountGit) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountJsonhalProjectGetServiceAccountGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountJsonhalProjectGetServiceAccountGit(val *ServiceAccountJsonhalProjectGetServiceAccountGit) *NullableServiceAccountJsonhalProjectGetServiceAccountGit {
	return &NullableServiceAccountJsonhalProjectGetServiceAccountGit{value: val, isSet: true}
}

func (v NullableServiceAccountJsonhalProjectGetServiceAccountGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountJsonhalProjectGetServiceAccountGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


