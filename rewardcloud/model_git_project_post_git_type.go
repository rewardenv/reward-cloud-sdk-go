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

// GitProjectPostGitType struct for GitProjectPostGitType
type GitProjectPostGitType struct {
	GitTypeProjectPost *GitTypeProjectPost
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GitProjectPostGitType) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into GitTypeProjectPost
	err = json.Unmarshal(data, &dst.GitTypeProjectPost);
	if err == nil {
		jsonGitTypeProjectPost, _ := json.Marshal(dst.GitTypeProjectPost)
		if string(jsonGitTypeProjectPost) == "{}" { // empty struct
			dst.GitTypeProjectPost = nil
		} else {
			return nil // data stored in dst.GitTypeProjectPost, return on the first match
		}
	} else {
		dst.GitTypeProjectPost = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GitProjectPostGitType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GitProjectPostGitType) MarshalJSON() ([]byte, error) {
	if src.GitTypeProjectPost != nil {
		return json.Marshal(&src.GitTypeProjectPost)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGitProjectPostGitType struct {
	value *GitProjectPostGitType
	isSet bool
}

func (v NullableGitProjectPostGitType) Get() *GitProjectPostGitType {
	return v.value
}

func (v *NullableGitProjectPostGitType) Set(val *GitProjectPostGitType) {
	v.value = val
	v.isSet = true
}

func (v NullableGitProjectPostGitType) IsSet() bool {
	return v.isSet
}

func (v *NullableGitProjectPostGitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitProjectPostGitType(val *GitProjectPostGitType) *NullableGitProjectPostGitType {
	return &NullableGitProjectPostGitType{value: val, isSet: true}
}

func (v NullableGitProjectPostGitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitProjectPostGitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


