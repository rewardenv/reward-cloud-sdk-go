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

// GitJsonhalProjectPostGitType struct for GitJsonhalProjectPostGitType
type GitJsonhalProjectPostGitType struct {
	GitTypeJsonhalProjectPost *GitTypeJsonhalProjectPost
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GitJsonhalProjectPostGitType) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into GitTypeJsonhalProjectPost
	err = json.Unmarshal(data, &dst.GitTypeJsonhalProjectPost);
	if err == nil {
		jsonGitTypeJsonhalProjectPost, _ := json.Marshal(dst.GitTypeJsonhalProjectPost)
		if string(jsonGitTypeJsonhalProjectPost) == "{}" { // empty struct
			dst.GitTypeJsonhalProjectPost = nil
		} else {
			return nil // data stored in dst.GitTypeJsonhalProjectPost, return on the first match
		}
	} else {
		dst.GitTypeJsonhalProjectPost = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GitJsonhalProjectPostGitType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GitJsonhalProjectPostGitType) MarshalJSON() ([]byte, error) {
	if src.GitTypeJsonhalProjectPost != nil {
		return json.Marshal(&src.GitTypeJsonhalProjectPost)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGitJsonhalProjectPostGitType struct {
	value *GitJsonhalProjectPostGitType
	isSet bool
}

func (v NullableGitJsonhalProjectPostGitType) Get() *GitJsonhalProjectPostGitType {
	return v.value
}

func (v *NullableGitJsonhalProjectPostGitType) Set(val *GitJsonhalProjectPostGitType) {
	v.value = val
	v.isSet = true
}

func (v NullableGitJsonhalProjectPostGitType) IsSet() bool {
	return v.isSet
}

func (v *NullableGitJsonhalProjectPostGitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitJsonhalProjectPostGitType(val *GitJsonhalProjectPostGitType) *NullableGitJsonhalProjectPostGitType {
	return &NullableGitJsonhalProjectPostGitType{value: val, isSet: true}
}

func (v NullableGitJsonhalProjectPostGitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitJsonhalProjectPostGitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


