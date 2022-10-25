/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.197.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentConsentMaxPaymentAmount Maximum amount of a single payment initiated using the payment consent.
type PaymentConsentMaxPaymentAmount struct {
	Currency PaymentAmountCurrency `json:"currency"`
	// The amount of the payment. Must contain at most two digits of precision e.g. `1.23`. Minimum accepted value is `1`.
	Value float64 `json:"value"`
}

// NewPaymentConsentMaxPaymentAmount instantiates a new PaymentConsentMaxPaymentAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentConsentMaxPaymentAmount(currency PaymentAmountCurrency, value float64) *PaymentConsentMaxPaymentAmount {
	this := PaymentConsentMaxPaymentAmount{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewPaymentConsentMaxPaymentAmountWithDefaults instantiates a new PaymentConsentMaxPaymentAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentConsentMaxPaymentAmountWithDefaults() *PaymentConsentMaxPaymentAmount {
	this := PaymentConsentMaxPaymentAmount{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *PaymentConsentMaxPaymentAmount) GetCurrency() PaymentAmountCurrency {
	if o == nil {
		var ret PaymentAmountCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentMaxPaymentAmount) GetCurrencyOk() (*PaymentAmountCurrency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentConsentMaxPaymentAmount) SetCurrency(v PaymentAmountCurrency) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *PaymentConsentMaxPaymentAmount) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PaymentConsentMaxPaymentAmount) GetValueOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PaymentConsentMaxPaymentAmount) SetValue(v float64) {
	o.Value = v
}

func (o PaymentConsentMaxPaymentAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentConsentMaxPaymentAmount struct {
	value *PaymentConsentMaxPaymentAmount
	isSet bool
}

func (v NullablePaymentConsentMaxPaymentAmount) Get() *PaymentConsentMaxPaymentAmount {
	return v.value
}

func (v *NullablePaymentConsentMaxPaymentAmount) Set(val *PaymentConsentMaxPaymentAmount) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentConsentMaxPaymentAmount) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentConsentMaxPaymentAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentConsentMaxPaymentAmount(val *PaymentConsentMaxPaymentAmount) *NullablePaymentConsentMaxPaymentAmount {
	return &NullablePaymentConsentMaxPaymentAmount{value: val, isSet: true}
}

func (v NullablePaymentConsentMaxPaymentAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentConsentMaxPaymentAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


