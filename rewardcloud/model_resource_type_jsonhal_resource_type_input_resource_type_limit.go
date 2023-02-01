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

// ResourceTypeJsonhalResourceTypeInputResourceTypeLimit struct for ResourceTypeJsonhalResourceTypeInputResourceTypeLimit
type ResourceTypeJsonhalResourceTypeInputResourceTypeLimit struct {
	ResourceTypeLimitJsonhalResourceTypeInput *ResourceTypeLimitJsonhalResourceTypeInput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResourceTypeJsonhalResourceTypeInputResourceTypeLimit) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ResourceTypeLimitJsonhalResourceTypeInput
	err = json.Unmarshal(data, &dst.ResourceTypeLimitJsonhalResourceTypeInput);
	if err == nil {
		jsonResourceTypeLimitJsonhalResourceTypeInput, _ := json.Marshal(dst.ResourceTypeLimitJsonhalResourceTypeInput)
		if string(jsonResourceTypeLimitJsonhalResourceTypeInput) == "{}" { // empty struct
			dst.ResourceTypeLimitJsonhalResourceTypeInput = nil
		} else {
			return nil // data stored in dst.ResourceTypeLimitJsonhalResourceTypeInput, return on the first match
		}
	} else {
		dst.ResourceTypeLimitJsonhalResourceTypeInput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ResourceTypeJsonhalResourceTypeInputResourceTypeLimit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResourceTypeJsonhalResourceTypeInputResourceTypeLimit) MarshalJSON() ([]byte, error) {
	if src.ResourceTypeLimitJsonhalResourceTypeInput != nil {
		return json.Marshal(&src.ResourceTypeLimitJsonhalResourceTypeInput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit struct {
	value *ResourceTypeJsonhalResourceTypeInputResourceTypeLimit
	isSet bool
}

func (v NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit) Get() *ResourceTypeJsonhalResourceTypeInputResourceTypeLimit {
	return v.value
}

func (v *NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit) Set(val *ResourceTypeJsonhalResourceTypeInputResourceTypeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit(val *ResourceTypeJsonhalResourceTypeInputResourceTypeLimit) *NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit {
	return &NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit{value: val, isSet: true}
}

func (v NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeJsonhalResourceTypeInputResourceTypeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

