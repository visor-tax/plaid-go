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

// CreditSessionBankIncomeStatus Status of the Bank Income Link session.  `APPROVED`: User has approved and verified their income  `NO_DEPOSITS_FOUND`: We attempted, but were unable to find any income in the connected account.  `USER_REPORTED_NO_INCOME`: The user explicitly indicated that they don't receive income in the connected account.  `STARTED`: The user began the bank income portion of the link flow.  `INTERNAL_ERROR`: The user encountered an internal error.
type CreditSessionBankIncomeStatus string

var _ = fmt.Printf

// List of CreditSessionBankIncomeStatus
const (
	CREDITSESSIONBANKINCOMESTATUS_APPROVED CreditSessionBankIncomeStatus = "APPROVED"
	CREDITSESSIONBANKINCOMESTATUS_NO_DEPOSITS_FOUND CreditSessionBankIncomeStatus = "NO_DEPOSITS_FOUND"
	CREDITSESSIONBANKINCOMESTATUS_USER_REPORTED_NO_INCOME CreditSessionBankIncomeStatus = "USER_REPORTED_NO_INCOME"
)

var allowedCreditSessionBankIncomeStatusEnumValues = []CreditSessionBankIncomeStatus{
	"APPROVED",
	"NO_DEPOSITS_FOUND",
	"USER_REPORTED_NO_INCOME",
}

func (v *CreditSessionBankIncomeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditSessionBankIncomeStatus(value)


	*v = enumTypeValue
	return nil
}

// NewCreditSessionBankIncomeStatusFromValue returns a pointer to a valid CreditSessionBankIncomeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditSessionBankIncomeStatusFromValue(v string) (*CreditSessionBankIncomeStatus, error) {
	ev := CreditSessionBankIncomeStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditSessionBankIncomeStatus) IsValid() bool {
	for _, existing := range allowedCreditSessionBankIncomeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditSessionBankIncomeStatus value
func (v CreditSessionBankIncomeStatus) Ptr() *CreditSessionBankIncomeStatus {
	return &v
}

type NullableCreditSessionBankIncomeStatus struct {
	value *CreditSessionBankIncomeStatus
	isSet bool
}

func (v NullableCreditSessionBankIncomeStatus) Get() *CreditSessionBankIncomeStatus {
	return v.value
}

func (v *NullableCreditSessionBankIncomeStatus) Set(val *CreditSessionBankIncomeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionBankIncomeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionBankIncomeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionBankIncomeStatus(val *CreditSessionBankIncomeStatus) *NullableCreditSessionBankIncomeStatus {
	return &NullableCreditSessionBankIncomeStatus{value: val, isSet: true}
}

func (v NullableCreditSessionBankIncomeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionBankIncomeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

