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

// PartnerEndCustomerStatus The status of the given end customer.  `UNDER REVIEW`: The end customer has been created and enabled in the non-Production environments. The end customer must be manually reviewed by the Plaid team before it can be enabled in production, at which point its status will automatically transition to `READY FOR ENABLEMENT` or `DENIED`.  `READY FOR ENABLEMENT`: The end customer is ready to be enabled in the Production environment. Call the `/partner/customer/enable` endpoint to enable the end customer in Production.  `ACTIVE`: The end customer has been enabled in all environments.   `DENIED`: The end customer has been created and enabled in the non-Production environments, but it did not pass review by the Plaid team and therefore cannot be enabled in the Production environment. Talk to your Account Manager for more information.
type PartnerEndCustomerStatus string

// List of PartnerEndCustomerStatus
const (
	PARTNERENDCUSTOMERSTATUS_UNDER_REVIEW PartnerEndCustomerStatus = "UNDER REVIEW"
	PARTNERENDCUSTOMERSTATUS_READY_FOR_ENABLEMENT PartnerEndCustomerStatus = "READY FOR ENABLEMENT"
	PARTNERENDCUSTOMERSTATUS_ACTIVE PartnerEndCustomerStatus = "ACTIVE"
	PARTNERENDCUSTOMERSTATUS_DENIED PartnerEndCustomerStatus = "DENIED"
)

var allowedPartnerEndCustomerStatusEnumValues = []PartnerEndCustomerStatus{
	"UNDER REVIEW",
	"READY FOR ENABLEMENT",
	"ACTIVE",
	"DENIED",
}

func (v *PartnerEndCustomerStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartnerEndCustomerStatus(value)
	for _, existing := range allowedPartnerEndCustomerStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartnerEndCustomerStatus", value)
}

// NewPartnerEndCustomerStatusFromValue returns a pointer to a valid PartnerEndCustomerStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartnerEndCustomerStatusFromValue(v string) (*PartnerEndCustomerStatus, error) {
	ev := PartnerEndCustomerStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartnerEndCustomerStatus: valid values are %v", v, allowedPartnerEndCustomerStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartnerEndCustomerStatus) IsValid() bool {
	for _, existing := range allowedPartnerEndCustomerStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartnerEndCustomerStatus value
func (v PartnerEndCustomerStatus) Ptr() *PartnerEndCustomerStatus {
	return &v
}

type NullablePartnerEndCustomerStatus struct {
	value *PartnerEndCustomerStatus
	isSet bool
}

func (v NullablePartnerEndCustomerStatus) Get() *PartnerEndCustomerStatus {
	return v.value
}

func (v *NullablePartnerEndCustomerStatus) Set(val *PartnerEndCustomerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerStatus(val *PartnerEndCustomerStatus) *NullablePartnerEndCustomerStatus {
	return &NullablePartnerEndCustomerStatus{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

