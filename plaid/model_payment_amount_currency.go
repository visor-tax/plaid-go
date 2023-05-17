/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.370.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// PaymentAmountCurrency The ISO-4217 currency code of the payment. For standing orders and payment consents, `\"GBP\"` must be used.
type PaymentAmountCurrency string

var _ = fmt.Printf

// List of PaymentAmountCurrency
const (
	PAYMENTAMOUNTCURRENCY_GBP PaymentAmountCurrency = "GBP"
	PAYMENTAMOUNTCURRENCY_EUR PaymentAmountCurrency = "EUR"
	PAYMENTAMOUNTCURRENCY_PLN PaymentAmountCurrency = "PLN"
	PAYMENTAMOUNTCURRENCY_SEK PaymentAmountCurrency = "SEK"
	PAYMENTAMOUNTCURRENCY_DKK PaymentAmountCurrency = "DKK"
	PAYMENTAMOUNTCURRENCY_NOK PaymentAmountCurrency = "NOK"
)

var allowedPaymentAmountCurrencyEnumValues = []PaymentAmountCurrency{
	"GBP",
	"EUR",
	"PLN",
	"SEK",
	"DKK",
	"NOK",
}

func (v *PaymentAmountCurrency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PaymentAmountCurrency(value)


	*v = enumTypeValue
	return nil
}

// NewPaymentAmountCurrencyFromValue returns a pointer to a valid PaymentAmountCurrency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentAmountCurrencyFromValue(v string) (*PaymentAmountCurrency, error) {
	ev := PaymentAmountCurrency(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentAmountCurrency) IsValid() bool {
	for _, existing := range allowedPaymentAmountCurrencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentAmountCurrency value
func (v PaymentAmountCurrency) Ptr() *PaymentAmountCurrency {
	return &v
}

type NullablePaymentAmountCurrency struct {
	value *PaymentAmountCurrency
	isSet bool
}

func (v NullablePaymentAmountCurrency) Get() *PaymentAmountCurrency {
	return v.value
}

func (v *NullablePaymentAmountCurrency) Set(val *PaymentAmountCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAmountCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAmountCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAmountCurrency(val *PaymentAmountCurrency) *NullablePaymentAmountCurrency {
	return &NullablePaymentAmountCurrency{value: val, isSet: true}
}

func (v NullablePaymentAmountCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAmountCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

