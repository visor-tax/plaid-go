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
	"fmt"
)

// TransactionsRuleField Transaction field for which the rule is defined.
type TransactionsRuleField string

// List of TransactionsRuleField
const (
	TRANSACTIONSRULEFIELD_TRANSACTION_ID TransactionsRuleField = "TRANSACTION_ID"
	TRANSACTIONSRULEFIELD_NAME TransactionsRuleField = "NAME"
)

var allowedTransactionsRuleFieldEnumValues = []TransactionsRuleField{
	"TRANSACTION_ID",
	"NAME",
}

func (v *TransactionsRuleField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionsRuleField(value)
	for _, existing := range allowedTransactionsRuleFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionsRuleField", value)
}

// NewTransactionsRuleFieldFromValue returns a pointer to a valid TransactionsRuleField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransactionsRuleFieldFromValue(v string) (*TransactionsRuleField, error) {
	ev := TransactionsRuleField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransactionsRuleField: valid values are %v", v, allowedTransactionsRuleFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransactionsRuleField) IsValid() bool {
	for _, existing := range allowedTransactionsRuleFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransactionsRuleField value
func (v TransactionsRuleField) Ptr() *TransactionsRuleField {
	return &v
}

type NullableTransactionsRuleField struct {
	value *TransactionsRuleField
	isSet bool
}

func (v NullableTransactionsRuleField) Get() *TransactionsRuleField {
	return v.value
}

func (v *NullableTransactionsRuleField) Set(val *TransactionsRuleField) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRuleField) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRuleField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRuleField(val *TransactionsRuleField) *NullableTransactionsRuleField {
	return &NullableTransactionsRuleField{value: val, isSet: true}
}

func (v NullableTransactionsRuleField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRuleField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

