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

// InvestmentTransactionType Value is one of the following: `buy`: Buying an investment `sell`: Selling an investment `cancel`: A cancellation of a pending transaction `cash`: Activity that modifies a cash position `fee`: A fee on the account `transfer`: Activity which modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer  For descriptions of possible transaction types and subtypes, see the [Investment transaction types schema](https://plaid.com/docs/api/accounts/#investment-transaction-types-schema).
type InvestmentTransactionType string

var _ = fmt.Printf

// List of InvestmentTransactionType
const (
	INVESTMENTTRANSACTIONTYPE_BUY InvestmentTransactionType = "buy"
	INVESTMENTTRANSACTIONTYPE_SELL InvestmentTransactionType = "sell"
	INVESTMENTTRANSACTIONTYPE_CANCEL InvestmentTransactionType = "cancel"
	INVESTMENTTRANSACTIONTYPE_CASH InvestmentTransactionType = "cash"
	INVESTMENTTRANSACTIONTYPE_FEE InvestmentTransactionType = "fee"
	INVESTMENTTRANSACTIONTYPE_TRANSFER InvestmentTransactionType = "transfer"
)

var allowedInvestmentTransactionTypeEnumValues = []InvestmentTransactionType{
	"buy",
	"sell",
	"cancel",
	"cash",
	"fee",
	"transfer",
}

func (v *InvestmentTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := InvestmentTransactionType(value)


	*v = enumTypeValue
	return nil
}

// NewInvestmentTransactionTypeFromValue returns a pointer to a valid InvestmentTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInvestmentTransactionTypeFromValue(v string) (*InvestmentTransactionType, error) {
	ev := InvestmentTransactionType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InvestmentTransactionType) IsValid() bool {
	for _, existing := range allowedInvestmentTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InvestmentTransactionType value
func (v InvestmentTransactionType) Ptr() *InvestmentTransactionType {
	return &v
}

type NullableInvestmentTransactionType struct {
	value *InvestmentTransactionType
	isSet bool
}

func (v NullableInvestmentTransactionType) Get() *InvestmentTransactionType {
	return v.value
}

func (v *NullableInvestmentTransactionType) Set(val *InvestmentTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentTransactionType(val *InvestmentTransactionType) *NullableInvestmentTransactionType {
	return &NullableInvestmentTransactionType{value: val, isSet: true}
}

func (v NullableInvestmentTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

