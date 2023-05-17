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

// PartnerEndCustomerFlowdownStatus The status of the addendum to the Plaid MSA (\"flowdown\") for the end customer.
type PartnerEndCustomerFlowdownStatus string

var _ = fmt.Printf

// List of PartnerEndCustomerFlowdownStatus
const (
	PARTNERENDCUSTOMERFLOWDOWNSTATUS_NOT_STARTED PartnerEndCustomerFlowdownStatus = "NOT_STARTED"
	PARTNERENDCUSTOMERFLOWDOWNSTATUS_IN_REVIEW PartnerEndCustomerFlowdownStatus = "IN_REVIEW"
	PARTNERENDCUSTOMERFLOWDOWNSTATUS_NEGOTIATION PartnerEndCustomerFlowdownStatus = "NEGOTIATION"
	PARTNERENDCUSTOMERFLOWDOWNSTATUS_COMPLETE PartnerEndCustomerFlowdownStatus = "COMPLETE"
)

var allowedPartnerEndCustomerFlowdownStatusEnumValues = []PartnerEndCustomerFlowdownStatus{
	"NOT_STARTED",
	"IN_REVIEW",
	"NEGOTIATION",
	"COMPLETE",
}

func (v *PartnerEndCustomerFlowdownStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := PartnerEndCustomerFlowdownStatus(value)


	*v = enumTypeValue
	return nil
}

// NewPartnerEndCustomerFlowdownStatusFromValue returns a pointer to a valid PartnerEndCustomerFlowdownStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartnerEndCustomerFlowdownStatusFromValue(v string) (*PartnerEndCustomerFlowdownStatus, error) {
	ev := PartnerEndCustomerFlowdownStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartnerEndCustomerFlowdownStatus) IsValid() bool {
	for _, existing := range allowedPartnerEndCustomerFlowdownStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartnerEndCustomerFlowdownStatus value
func (v PartnerEndCustomerFlowdownStatus) Ptr() *PartnerEndCustomerFlowdownStatus {
	return &v
}

type NullablePartnerEndCustomerFlowdownStatus struct {
	value *PartnerEndCustomerFlowdownStatus
	isSet bool
}

func (v NullablePartnerEndCustomerFlowdownStatus) Get() *PartnerEndCustomerFlowdownStatus {
	return v.value
}

func (v *NullablePartnerEndCustomerFlowdownStatus) Set(val *PartnerEndCustomerFlowdownStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerFlowdownStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerFlowdownStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerFlowdownStatus(val *PartnerEndCustomerFlowdownStatus) *NullablePartnerEndCustomerFlowdownStatus {
	return &NullablePartnerEndCustomerFlowdownStatus{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerFlowdownStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerFlowdownStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

