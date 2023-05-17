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

// TransferScheduleIntervalUnit The unit of the recurring interval.
type TransferScheduleIntervalUnit string

var _ = fmt.Printf

// List of TransferScheduleIntervalUnit
const (
	TRANSFERSCHEDULEINTERVALUNIT_WEEK TransferScheduleIntervalUnit = "week"
	TRANSFERSCHEDULEINTERVALUNIT_MONTH TransferScheduleIntervalUnit = "month"
)

var allowedTransferScheduleIntervalUnitEnumValues = []TransferScheduleIntervalUnit{
	"week",
	"month",
}

func (v *TransferScheduleIntervalUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := TransferScheduleIntervalUnit(value)


	*v = enumTypeValue
	return nil
}

// NewTransferScheduleIntervalUnitFromValue returns a pointer to a valid TransferScheduleIntervalUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransferScheduleIntervalUnitFromValue(v string) (*TransferScheduleIntervalUnit, error) {
	ev := TransferScheduleIntervalUnit(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransferScheduleIntervalUnit) IsValid() bool {
	for _, existing := range allowedTransferScheduleIntervalUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransferScheduleIntervalUnit value
func (v TransferScheduleIntervalUnit) Ptr() *TransferScheduleIntervalUnit {
	return &v
}

type NullableTransferScheduleIntervalUnit struct {
	value *TransferScheduleIntervalUnit
	isSet bool
}

func (v NullableTransferScheduleIntervalUnit) Get() *TransferScheduleIntervalUnit {
	return v.value
}

func (v *NullableTransferScheduleIntervalUnit) Set(val *TransferScheduleIntervalUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferScheduleIntervalUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferScheduleIntervalUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferScheduleIntervalUnit(val *TransferScheduleIntervalUnit) *NullableTransferScheduleIntervalUnit {
	return &NullableTransferScheduleIntervalUnit{value: val, isSet: true}
}

func (v NullableTransferScheduleIntervalUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferScheduleIntervalUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

