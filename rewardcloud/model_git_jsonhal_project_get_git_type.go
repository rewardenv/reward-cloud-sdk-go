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

// GitJsonhalProjectGetGitType struct for GitJsonhalProjectGetGitType
type GitJsonhalProjectGetGitType struct {
	GitTypeJsonhalProjectGet *GitTypeJsonhalProjectGet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GitJsonhalProjectGetGitType) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into GitTypeJsonhalProjectGet
	err = json.Unmarshal(data, &dst.GitTypeJsonhalProjectGet);
	if err == nil {
		jsonGitTypeJsonhalProjectGet, _ := json.Marshal(dst.GitTypeJsonhalProjectGet)
		if string(jsonGitTypeJsonhalProjectGet) == "{}" { // empty struct
			dst.GitTypeJsonhalProjectGet = nil
		} else {
			return nil // data stored in dst.GitTypeJsonhalProjectGet, return on the first match
		}
	} else {
		dst.GitTypeJsonhalProjectGet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GitJsonhalProjectGetGitType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GitJsonhalProjectGetGitType) MarshalJSON() ([]byte, error) {
	if src.GitTypeJsonhalProjectGet != nil {
		return json.Marshal(&src.GitTypeJsonhalProjectGet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGitJsonhalProjectGetGitType struct {
	value *GitJsonhalProjectGetGitType
	isSet bool
}

func (v NullableGitJsonhalProjectGetGitType) Get() *GitJsonhalProjectGetGitType {
	return v.value
}

func (v *NullableGitJsonhalProjectGetGitType) Set(val *GitJsonhalProjectGetGitType) {
	v.value = val
	v.isSet = true
}

func (v NullableGitJsonhalProjectGetGitType) IsSet() bool {
	return v.isSet
}

func (v *NullableGitJsonhalProjectGetGitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitJsonhalProjectGetGitType(val *GitJsonhalProjectGetGitType) *NullableGitJsonhalProjectGetGitType {
	return &NullableGitJsonhalProjectGetGitType{value: val, isSet: true}
}

func (v NullableGitJsonhalProjectGetGitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitJsonhalProjectGetGitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

