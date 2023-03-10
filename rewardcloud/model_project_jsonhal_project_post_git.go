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

// ProjectJsonhalProjectPostGit struct for ProjectJsonhalProjectPostGit
type ProjectJsonhalProjectPostGit struct {
	GitJsonhalProjectPost *GitJsonhalProjectPost
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProjectJsonhalProjectPostGit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into GitJsonhalProjectPost
	err = json.Unmarshal(data, &dst.GitJsonhalProjectPost);
	if err == nil {
		jsonGitJsonhalProjectPost, _ := json.Marshal(dst.GitJsonhalProjectPost)
		if string(jsonGitJsonhalProjectPost) == "{}" { // empty struct
			dst.GitJsonhalProjectPost = nil
		} else {
			return nil // data stored in dst.GitJsonhalProjectPost, return on the first match
		}
	} else {
		dst.GitJsonhalProjectPost = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProjectJsonhalProjectPostGit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProjectJsonhalProjectPostGit) MarshalJSON() ([]byte, error) {
	if src.GitJsonhalProjectPost != nil {
		return json.Marshal(&src.GitJsonhalProjectPost)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProjectJsonhalProjectPostGit struct {
	value *ProjectJsonhalProjectPostGit
	isSet bool
}

func (v NullableProjectJsonhalProjectPostGit) Get() *ProjectJsonhalProjectPostGit {
	return v.value
}

func (v *NullableProjectJsonhalProjectPostGit) Set(val *ProjectJsonhalProjectPostGit) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectJsonhalProjectPostGit) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectJsonhalProjectPostGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectJsonhalProjectPostGit(val *ProjectJsonhalProjectPostGit) *NullableProjectJsonhalProjectPostGit {
	return &NullableProjectJsonhalProjectPostGit{value: val, isSet: true}
}

func (v NullableProjectJsonhalProjectPostGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectJsonhalProjectPostGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


