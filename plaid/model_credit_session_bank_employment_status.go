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

// CreditSessionBankEmploymentStatus The terminal status of the bank employment verification.  `APPROVED`: User has approved and verified their employment.  `NO_EMPLOYMENTS_FOUND`: We attempted, but were unable to find any employment in the connected account.  `EMPLOYER_NOT_LISTED`: The user explicitly indicated that they did not see their current or previous employer in the list of employer names found.
type CreditSessionBankEmploymentStatus string

var _ = fmt.Printf

// List of CreditSessionBankEmploymentStatus
const (
	CREDITSESSIONBANKEMPLOYMENTSTATUS_APPROVED CreditSessionBankEmploymentStatus = "APPROVED"
	CREDITSESSIONBANKEMPLOYMENTSTATUS_NO_EMPLOYERS_FOUND CreditSessionBankEmploymentStatus = "NO_EMPLOYERS_FOUND"
	CREDITSESSIONBANKEMPLOYMENTSTATUS_EMPLOYER_NOT_LISTED CreditSessionBankEmploymentStatus = "EMPLOYER_NOT_LISTED"
)

var allowedCreditSessionBankEmploymentStatusEnumValues = []CreditSessionBankEmploymentStatus{
	"APPROVED",
	"NO_EMPLOYERS_FOUND",
	"EMPLOYER_NOT_LISTED",
}

func (v *CreditSessionBankEmploymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := CreditSessionBankEmploymentStatus(value)


	*v = enumTypeValue
	return nil
}

// NewCreditSessionBankEmploymentStatusFromValue returns a pointer to a valid CreditSessionBankEmploymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditSessionBankEmploymentStatusFromValue(v string) (*CreditSessionBankEmploymentStatus, error) {
	ev := CreditSessionBankEmploymentStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditSessionBankEmploymentStatus) IsValid() bool {
	for _, existing := range allowedCreditSessionBankEmploymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditSessionBankEmploymentStatus value
func (v CreditSessionBankEmploymentStatus) Ptr() *CreditSessionBankEmploymentStatus {
	return &v
}

type NullableCreditSessionBankEmploymentStatus struct {
	value *CreditSessionBankEmploymentStatus
	isSet bool
}

func (v NullableCreditSessionBankEmploymentStatus) Get() *CreditSessionBankEmploymentStatus {
	return v.value
}

func (v *NullableCreditSessionBankEmploymentStatus) Set(val *CreditSessionBankEmploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionBankEmploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionBankEmploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionBankEmploymentStatus(val *CreditSessionBankEmploymentStatus) *NullableCreditSessionBankEmploymentStatus {
	return &NullableCreditSessionBankEmploymentStatus{value: val, isSet: true}
}

func (v NullableCreditSessionBankEmploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionBankEmploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

