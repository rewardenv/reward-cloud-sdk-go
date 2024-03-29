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

// PriceProduct struct for PriceProduct
type PriceProduct struct {
	Product *Product
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PriceProduct) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Product
	err = json.Unmarshal(data, &dst.Product)
	if err == nil {
		jsonProduct, _ := json.Marshal(dst.Product)
		if string(jsonProduct) == "{}" { // empty struct
			dst.Product = nil
		} else {
			return nil // data stored in dst.Product, return on the first match
		}
	} else {
		dst.Product = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PriceProduct)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PriceProduct) MarshalJSON() ([]byte, error) {
	if src.Product != nil {
		return json.Marshal(&src.Product)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePriceProduct struct {
	value *PriceProduct
	isSet bool
}

func (v NullablePriceProduct) Get() *PriceProduct {
	return v.value
}

func (v *NullablePriceProduct) Set(val *PriceProduct) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceProduct) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceProduct(val *PriceProduct) *NullablePriceProduct {
	return &NullablePriceProduct{value: val, isSet: true}
}

func (v NullablePriceProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
