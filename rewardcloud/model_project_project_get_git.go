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

// ProjectProjectGetGit struct for ProjectProjectGetGit
type ProjectProjectGetGit struct {
	GitProjectGet *GitProjectGet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProjectProjectGetGit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into GitProjectGet
	err = json.Unmarshal(data, &dst.GitProjectGet);
	if err == nil {
		jsonGitProjectGet, _ := json.Marshal(dst.GitProjectGet)
		if string(jsonGitProjectGet) == "{}" { // empty struct
			dst.GitProjectGet = nil
		} else {
			return nil // data stored in dst.GitProjectGet, return on the first match
		}
	} else {
		dst.GitProjectGet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProjectProjectGetGit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProjectProjectGetGit) MarshalJSON() ([]byte, error) {
	if src.GitProjectGet != nil {
		return json.Marshal(&src.GitProjectGet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProjectProjectGetGit struct {
	value *ProjectProjectGetGit
	isSet bool
}

func (v NullableProjectProjectGetGit) Get() *ProjectProjectGetGit {
	return v.value
}

func (v *NullableProjectProjectGetGit) Set(val *ProjectProjectGetGit) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectGetGit) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectGetGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectGetGit(val *ProjectProjectGetGit) *NullableProjectProjectGetGit {
	return &NullableProjectProjectGetGit{value: val, isSet: true}
}

func (v NullableProjectProjectGetGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectGetGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


