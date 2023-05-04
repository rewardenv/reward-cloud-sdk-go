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

// ProjectEnvVarProject struct for ProjectEnvVarProject
type ProjectEnvVarProject struct {
	AbstractProject *AbstractProject
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProjectEnvVarProject) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into AbstractProject
	err = json.Unmarshal(data, &dst.AbstractProject);
	if err == nil {
		jsonAbstractProject, _ := json.Marshal(dst.AbstractProject)
		if string(jsonAbstractProject) == "{}" { // empty struct
			dst.AbstractProject = nil
		} else {
			return nil // data stored in dst.AbstractProject, return on the first match
		}
	} else {
		dst.AbstractProject = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProjectEnvVarProject)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProjectEnvVarProject) MarshalJSON() ([]byte, error) {
	if src.AbstractProject != nil {
		return json.Marshal(&src.AbstractProject)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProjectEnvVarProject struct {
	value *ProjectEnvVarProject
	isSet bool
}

func (v NullableProjectEnvVarProject) Get() *ProjectEnvVarProject {
	return v.value
}

func (v *NullableProjectEnvVarProject) Set(val *ProjectEnvVarProject) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectEnvVarProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectEnvVarProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectEnvVarProject(val *ProjectEnvVarProject) *NullableProjectEnvVarProject {
	return &NullableProjectEnvVarProject{value: val, isSet: true}
}

func (v NullableProjectEnvVarProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectEnvVarProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

