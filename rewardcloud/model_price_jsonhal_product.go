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

// PriceJsonhalProduct struct for PriceJsonhalProduct
type PriceJsonhalProduct struct {
	ProductJsonhal *ProductJsonhal
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PriceJsonhalProduct) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into ProductJsonhal
	err = json.Unmarshal(data, &dst.ProductJsonhal);
	if err == nil {
		jsonProductJsonhal, _ := json.Marshal(dst.ProductJsonhal)
		if string(jsonProductJsonhal) == "{}" { // empty struct
			dst.ProductJsonhal = nil
		} else {
			return nil // data stored in dst.ProductJsonhal, return on the first match
		}
	} else {
		dst.ProductJsonhal = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PriceJsonhalProduct)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PriceJsonhalProduct) MarshalJSON() ([]byte, error) {
	if src.ProductJsonhal != nil {
		return json.Marshal(&src.ProductJsonhal)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePriceJsonhalProduct struct {
	value *PriceJsonhalProduct
	isSet bool
}

func (v NullablePriceJsonhalProduct) Get() *PriceJsonhalProduct {
	return v.value
}

func (v *NullablePriceJsonhalProduct) Set(val *PriceJsonhalProduct) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceJsonhalProduct) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceJsonhalProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceJsonhalProduct(val *PriceJsonhalProduct) *NullablePriceJsonhalProduct {
	return &NullablePriceJsonhalProduct{value: val, isSet: true}
}

func (v NullablePriceJsonhalProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceJsonhalProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

