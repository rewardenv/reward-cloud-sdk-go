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

// ServiceAccountProjectGetServiceAccountGit struct for ServiceAccountProjectGetServiceAccountGit
type ServiceAccountProjectGetServiceAccountGit struct {
	ServiceAccountGitProjectGet *ServiceAccountGitProjectGet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceAccountProjectGetServiceAccountGit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ServiceAccountGitProjectGet
	err = json.Unmarshal(data, &dst.ServiceAccountGitProjectGet);
	if err == nil {
		jsonServiceAccountGitProjectGet, _ := json.Marshal(dst.ServiceAccountGitProjectGet)
		if string(jsonServiceAccountGitProjectGet) == "{}" { // empty struct
			dst.ServiceAccountGitProjectGet = nil
		} else {
			return nil // data stored in dst.ServiceAccountGitProjectGet, return on the first match
		}
	} else {
		dst.ServiceAccountGitProjectGet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceAccountProjectGetServiceAccountGit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceAccountProjectGetServiceAccountGit) MarshalJSON() ([]byte, error) {
	if src.ServiceAccountGitProjectGet != nil {
		return json.Marshal(&src.ServiceAccountGitProjectGet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceAccountProjectGetServiceAccountGit struct {
	value *ServiceAccountProjectGetServiceAccountGit
	isSet bool
}

func (v NullableServiceAccountProjectGetServiceAccountGit) Get() *ServiceAccountProjectGetServiceAccountGit {
	return v.value
}

func (v *NullableServiceAccountProjectGetServiceAccountGit) Set(val *ServiceAccountProjectGetServiceAccountGit) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountProjectGetServiceAccountGit) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountProjectGetServiceAccountGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountProjectGetServiceAccountGit(val *ServiceAccountProjectGetServiceAccountGit) *NullableServiceAccountProjectGetServiceAccountGit {
	return &NullableServiceAccountProjectGetServiceAccountGit{value: val, isSet: true}
}

func (v NullableServiceAccountProjectGetServiceAccountGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountProjectGetServiceAccountGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

