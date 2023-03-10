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

// ServiceAccountProjectGetServiceAccountRegistry struct for ServiceAccountProjectGetServiceAccountRegistry
type ServiceAccountProjectGetServiceAccountRegistry struct {
	ServiceAccountRegistryProjectGet *ServiceAccountRegistryProjectGet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceAccountProjectGetServiceAccountRegistry) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ServiceAccountRegistryProjectGet
	err = json.Unmarshal(data, &dst.ServiceAccountRegistryProjectGet);
	if err == nil {
		jsonServiceAccountRegistryProjectGet, _ := json.Marshal(dst.ServiceAccountRegistryProjectGet)
		if string(jsonServiceAccountRegistryProjectGet) == "{}" { // empty struct
			dst.ServiceAccountRegistryProjectGet = nil
		} else {
			return nil // data stored in dst.ServiceAccountRegistryProjectGet, return on the first match
		}
	} else {
		dst.ServiceAccountRegistryProjectGet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceAccountProjectGetServiceAccountRegistry)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceAccountProjectGetServiceAccountRegistry) MarshalJSON() ([]byte, error) {
	if src.ServiceAccountRegistryProjectGet != nil {
		return json.Marshal(&src.ServiceAccountRegistryProjectGet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceAccountProjectGetServiceAccountRegistry struct {
	value *ServiceAccountProjectGetServiceAccountRegistry
	isSet bool
}

func (v NullableServiceAccountProjectGetServiceAccountRegistry) Get() *ServiceAccountProjectGetServiceAccountRegistry {
	return v.value
}

func (v *NullableServiceAccountProjectGetServiceAccountRegistry) Set(val *ServiceAccountProjectGetServiceAccountRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountProjectGetServiceAccountRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountProjectGetServiceAccountRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountProjectGetServiceAccountRegistry(val *ServiceAccountProjectGetServiceAccountRegistry) *NullableServiceAccountProjectGetServiceAccountRegistry {
	return &NullableServiceAccountProjectGetServiceAccountRegistry{value: val, isSet: true}
}

func (v NullableServiceAccountProjectGetServiceAccountRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountProjectGetServiceAccountRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


