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

// ProductProductOutputDefaultPrice struct for ProductProductOutputDefaultPrice
type ProductProductOutputDefaultPrice struct {
	PriceProductOutput *PriceProductOutput
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProductProductOutputDefaultPrice) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PriceProductOutput
	err = json.Unmarshal(data, &dst.PriceProductOutput)
	if err == nil {
		jsonPriceProductOutput, _ := json.Marshal(dst.PriceProductOutput)
		if string(jsonPriceProductOutput) == "{}" { // empty struct
			dst.PriceProductOutput = nil
		} else {
			return nil // data stored in dst.PriceProductOutput, return on the first match
		}
	} else {
		dst.PriceProductOutput = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ProductProductOutputDefaultPrice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProductProductOutputDefaultPrice) MarshalJSON() ([]byte, error) {
	if src.PriceProductOutput != nil {
		return json.Marshal(&src.PriceProductOutput)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProductProductOutputDefaultPrice struct {
	value *ProductProductOutputDefaultPrice
	isSet bool
}

func (v NullableProductProductOutputDefaultPrice) Get() *ProductProductOutputDefaultPrice {
	return v.value
}

func (v *NullableProductProductOutputDefaultPrice) Set(val *ProductProductOutputDefaultPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProductProductOutputDefaultPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProductProductOutputDefaultPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductProductOutputDefaultPrice(val *ProductProductOutputDefaultPrice) *NullableProductProductOutputDefaultPrice {
	return &NullableProductProductOutputDefaultPrice{value: val, isSet: true}
}

func (v NullableProductProductOutputDefaultPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductProductOutputDefaultPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
