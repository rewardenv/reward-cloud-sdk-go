/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
	"fmt"
)

// ResourceTypeResourceTypeOutputResourceTypeLimit struct for ResourceTypeResourceTypeOutputResourceTypeLimit
type ResourceTypeResourceTypeOutputResourceTypeLimit struct {
	ResourceTypeLimitResourceTypeOutput *ResourceTypeLimitResourceTypeOutput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResourceTypeResourceTypeOutputResourceTypeLimit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ResourceTypeLimitResourceTypeOutput
	err = json.Unmarshal(data, &dst.ResourceTypeLimitResourceTypeOutput)
	if err == nil {
		jsonResourceTypeLimitResourceTypeOutput, _ := json.Marshal(dst.ResourceTypeLimitResourceTypeOutput)
		if string(jsonResourceTypeLimitResourceTypeOutput) == "{}" { // empty struct
			dst.ResourceTypeLimitResourceTypeOutput = nil
		} else {
			return nil // data stored in dst.ResourceTypeLimitResourceTypeOutput, return on the first match
		}
	} else {
		dst.ResourceTypeLimitResourceTypeOutput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResourceTypeResourceTypeOutputResourceTypeLimit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResourceTypeResourceTypeOutputResourceTypeLimit) MarshalJSON() ([]byte, error) {
	if src.ResourceTypeLimitResourceTypeOutput != nil {
		return json.Marshal(&src.ResourceTypeLimitResourceTypeOutput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResourceTypeResourceTypeOutputResourceTypeLimit struct {
	value *ResourceTypeResourceTypeOutputResourceTypeLimit
	isSet bool
}

func (v NullableResourceTypeResourceTypeOutputResourceTypeLimit) Get() *ResourceTypeResourceTypeOutputResourceTypeLimit {
	return v.value
}

func (v *NullableResourceTypeResourceTypeOutputResourceTypeLimit) Set(val *ResourceTypeResourceTypeOutputResourceTypeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeResourceTypeOutputResourceTypeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeResourceTypeOutputResourceTypeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeResourceTypeOutputResourceTypeLimit(val *ResourceTypeResourceTypeOutputResourceTypeLimit) *NullableResourceTypeResourceTypeOutputResourceTypeLimit {
	return &NullableResourceTypeResourceTypeOutputResourceTypeLimit{value: val, isSet: true}
}

func (v NullableResourceTypeResourceTypeOutputResourceTypeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeResourceTypeOutputResourceTypeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
