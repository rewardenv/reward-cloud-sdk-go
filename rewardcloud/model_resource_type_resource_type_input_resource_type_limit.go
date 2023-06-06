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

// ResourceTypeResourceTypeInputResourceTypeLimit struct for ResourceTypeResourceTypeInputResourceTypeLimit
type ResourceTypeResourceTypeInputResourceTypeLimit struct {
	ResourceTypeLimitResourceTypeInput *ResourceTypeLimitResourceTypeInput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResourceTypeResourceTypeInputResourceTypeLimit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ResourceTypeLimitResourceTypeInput
	err = json.Unmarshal(data, &dst.ResourceTypeLimitResourceTypeInput)
	if err == nil {
		jsonResourceTypeLimitResourceTypeInput, _ := json.Marshal(dst.ResourceTypeLimitResourceTypeInput)
		if string(jsonResourceTypeLimitResourceTypeInput) == "{}" { // empty struct
			dst.ResourceTypeLimitResourceTypeInput = nil
		} else {
			return nil // data stored in dst.ResourceTypeLimitResourceTypeInput, return on the first match
		}
	} else {
		dst.ResourceTypeLimitResourceTypeInput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResourceTypeResourceTypeInputResourceTypeLimit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResourceTypeResourceTypeInputResourceTypeLimit) MarshalJSON() ([]byte, error) {
	if src.ResourceTypeLimitResourceTypeInput != nil {
		return json.Marshal(&src.ResourceTypeLimitResourceTypeInput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResourceTypeResourceTypeInputResourceTypeLimit struct {
	value *ResourceTypeResourceTypeInputResourceTypeLimit
	isSet bool
}

func (v NullableResourceTypeResourceTypeInputResourceTypeLimit) Get() *ResourceTypeResourceTypeInputResourceTypeLimit {
	return v.value
}

func (v *NullableResourceTypeResourceTypeInputResourceTypeLimit) Set(val *ResourceTypeResourceTypeInputResourceTypeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeResourceTypeInputResourceTypeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeResourceTypeInputResourceTypeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeResourceTypeInputResourceTypeLimit(val *ResourceTypeResourceTypeInputResourceTypeLimit) *NullableResourceTypeResourceTypeInputResourceTypeLimit {
	return &NullableResourceTypeResourceTypeInputResourceTypeLimit{value: val, isSet: true}
}

func (v NullableResourceTypeResourceTypeInputResourceTypeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeResourceTypeInputResourceTypeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
