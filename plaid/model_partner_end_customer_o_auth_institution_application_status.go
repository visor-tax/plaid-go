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

// PartnerEndCustomerOAuthInstitutionApplicationStatus The registration status for the end customer's application.
type PartnerEndCustomerOAuthInstitutionApplicationStatus string

var _ = fmt.Printf

// List of PartnerEndCustomerOAuthInstitutionApplicationStatus
const (
	PARTNERENDCUSTOMEROAUTHINSTITUTIONAPPLICATIONSTATUS_NOT_STARTED PartnerEndCustomerOAuthInstitutionApplicationStatus = "NOT_STARTED"
	PARTNERENDCUSTOMEROAUTHINSTITUTIONAPPLICATIONSTATUS_PROCESSING PartnerEndCustomerOAuthInstitutionApplicationStatus = "PROCESSING"
	PARTNERENDCUSTOMEROAUTHINSTITUTIONAPPLICATIONSTATUS_APPROVED PartnerEndCustomerOAuthInstitutionApplicationStatus = "APPROVED"
	PARTNERENDCUSTOMEROAUTHINSTITUTIONAPPLICATIONSTATUS_ENABLED PartnerEndCustomerOAuthInstitutionApplicationStatus = "ENABLED"
	PARTNERENDCUSTOMEROAUTHINSTITUTIONAPPLICATIONSTATUS_ATTENTION_REQUIRED PartnerEndCustomerOAuthInstitutionApplicationStatus = "ATTENTION_REQUIRED"
)

var allowedPartnerEndCustomerOAuthInstitutionApplicationStatusEnumValues = []PartnerEndCustomerOAuthInstitutionApplicationStatus{
	"NOT_STARTED",
	"PROCESSING",
	"APPROVED",
	"ENABLED",
	"ATTENTION_REQUIRED",
}

func (v *PartnerEndCustomerOAuthInstitutionApplicationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PartnerEndCustomerOAuthInstitutionApplicationStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPartnerEndCustomerOAuthInstitutionApplicationStatusFromValue returns a pointer to a valid PartnerEndCustomerOAuthInstitutionApplicationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartnerEndCustomerOAuthInstitutionApplicationStatusFromValue(v string) (*PartnerEndCustomerOAuthInstitutionApplicationStatus, error) {
	ev := PartnerEndCustomerOAuthInstitutionApplicationStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartnerEndCustomerOAuthInstitutionApplicationStatus) IsValid() bool {
	for _, existing := range allowedPartnerEndCustomerOAuthInstitutionApplicationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartnerEndCustomerOAuthInstitutionApplicationStatus value
func (v PartnerEndCustomerOAuthInstitutionApplicationStatus) Ptr() *PartnerEndCustomerOAuthInstitutionApplicationStatus {
	return &v
}

type NullablePartnerEndCustomerOAuthInstitutionApplicationStatus struct {
	value *PartnerEndCustomerOAuthInstitutionApplicationStatus
	isSet bool
}

func (v NullablePartnerEndCustomerOAuthInstitutionApplicationStatus) Get() *PartnerEndCustomerOAuthInstitutionApplicationStatus {
	return v.value
}

func (v *NullablePartnerEndCustomerOAuthInstitutionApplicationStatus) Set(val *PartnerEndCustomerOAuthInstitutionApplicationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerOAuthInstitutionApplicationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerOAuthInstitutionApplicationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerOAuthInstitutionApplicationStatus(val *PartnerEndCustomerOAuthInstitutionApplicationStatus) *NullablePartnerEndCustomerOAuthInstitutionApplicationStatus {
	return &NullablePartnerEndCustomerOAuthInstitutionApplicationStatus{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerOAuthInstitutionApplicationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerOAuthInstitutionApplicationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

