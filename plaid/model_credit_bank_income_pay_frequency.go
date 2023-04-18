/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.345.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// CreditBankIncomePayFrequency The income pay frequency.
type CreditBankIncomePayFrequency string

var _ = fmt.Printf

// List of CreditBankIncomePayFrequency
const (
	CREDITBANKINCOMEPAYFREQUENCY_WEEKLY CreditBankIncomePayFrequency = "WEEKLY"
	CREDITBANKINCOMEPAYFREQUENCY_BIWEEKLY CreditBankIncomePayFrequency = "BIWEEKLY"
	CREDITBANKINCOMEPAYFREQUENCY_SEMI_MONTHLY CreditBankIncomePayFrequency = "SEMI_MONTHLY"
	CREDITBANKINCOMEPAYFREQUENCY_MONTHLY CreditBankIncomePayFrequency = "MONTHLY"
	CREDITBANKINCOMEPAYFREQUENCY_DAILY CreditBankIncomePayFrequency = "DAILY"
	CREDITBANKINCOMEPAYFREQUENCY_UNKNOWN CreditBankIncomePayFrequency = "UNKNOWN"
)

var allowedCreditBankIncomePayFrequencyEnumValues = []CreditBankIncomePayFrequency{
	"WEEKLY",
	"BIWEEKLY",
	"SEMI_MONTHLY",
	"MONTHLY",
	"DAILY",
	"UNKNOWN",
}

func (v *CreditBankIncomePayFrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditBankIncomePayFrequency(value)


	*v = enumTypeValue
	return nil
}

// NewCreditBankIncomePayFrequencyFromValue returns a pointer to a valid CreditBankIncomePayFrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditBankIncomePayFrequencyFromValue(v string) (*CreditBankIncomePayFrequency, error) {
	ev := CreditBankIncomePayFrequency(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditBankIncomePayFrequency) IsValid() bool {
	for _, existing := range allowedCreditBankIncomePayFrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditBankIncomePayFrequency value
func (v CreditBankIncomePayFrequency) Ptr() *CreditBankIncomePayFrequency {
	return &v
}

type NullableCreditBankIncomePayFrequency struct {
	value *CreditBankIncomePayFrequency
	isSet bool
}

func (v NullableCreditBankIncomePayFrequency) Get() *CreditBankIncomePayFrequency {
	return v.value
}

func (v *NullableCreditBankIncomePayFrequency) Set(val *CreditBankIncomePayFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomePayFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomePayFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomePayFrequency(val *CreditBankIncomePayFrequency) *NullableCreditBankIncomePayFrequency {
	return &NullableCreditBankIncomePayFrequency{value: val, isSet: true}
}

func (v NullableCreditBankIncomePayFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomePayFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

