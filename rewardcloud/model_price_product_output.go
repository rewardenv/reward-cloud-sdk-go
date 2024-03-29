/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// checks if the PriceProductOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceProductOutput{}

// PriceProductOutput Class Price
type PriceProductOutput struct {
	PaymentId NullableString  `json:"paymentId,omitempty"`
	Amount    NullableFloat32 `json:"amount,omitempty"`
	Currency  NullableString  `json:"currency,omitempty"`
	IsActive  NullableBool    `json:"isActive,omitempty"`
}

// NewPriceProductOutput instantiates a new PriceProductOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceProductOutput() *PriceProductOutput {
	this := PriceProductOutput{}
	return &this
}

// NewPriceProductOutputWithDefaults instantiates a new PriceProductOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceProductOutputWithDefaults() *PriceProductOutput {
	this := PriceProductOutput{}
	return &this
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceProductOutput) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceProductOutput) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *PriceProductOutput) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *PriceProductOutput) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}

// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *PriceProductOutput) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *PriceProductOutput) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceProductOutput) GetAmount() float32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceProductOutput) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *PriceProductOutput) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat32 and assigns it to the Amount field.
func (o *PriceProductOutput) SetAmount(v float32) {
	o.Amount.Set(&v)
}

// SetAmountNil sets the value for Amount to be an explicit nil
func (o *PriceProductOutput) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *PriceProductOutput) UnsetAmount() {
	o.Amount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceProductOutput) GetCurrency() string {
	if o == nil || IsNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceProductOutput) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PriceProductOutput) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *PriceProductOutput) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PriceProductOutput) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PriceProductOutput) UnsetCurrency() {
	o.Currency.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceProductOutput) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceProductOutput) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *PriceProductOutput) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *PriceProductOutput) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}

// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *PriceProductOutput) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *PriceProductOutput) UnsetIsActive() {
	o.IsActive.Unset()
}

func (o PriceProductOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceProductOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentId.IsSet() {
		toSerialize["paymentId"] = o.PaymentId.Get()
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	return toSerialize, nil
}

type NullablePriceProductOutput struct {
	value *PriceProductOutput
	isSet bool
}

func (v NullablePriceProductOutput) Get() *PriceProductOutput {
	return v.value
}

func (v *NullablePriceProductOutput) Set(val *PriceProductOutput) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceProductOutput) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceProductOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceProductOutput(val *PriceProductOutput) *NullablePriceProductOutput {
	return &NullablePriceProductOutput{value: val, isSet: true}
}

func (v NullablePriceProductOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceProductOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
