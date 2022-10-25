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

// ReportType The report type. It can be `assets` or `income`.
type ReportType string

// List of ReportType
const (
	REPORTTYPE_ASSETS ReportType = "assets"
	REPORTTYPE_INCOME ReportType = "income"
)

var allowedReportTypeEnumValues = []ReportType{
	"assets",
	"income",
}

func (v *ReportType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportType(value)
	for _, existing := range allowedReportTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportType", value)
}

// NewReportTypeFromValue returns a pointer to a valid ReportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportTypeFromValue(v string) (*ReportType, error) {
	ev := ReportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportType: valid values are %v", v, allowedReportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportType) IsValid() bool {
	for _, existing := range allowedReportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportType value
func (v ReportType) Ptr() *ReportType {
	return &v
}

type NullableReportType struct {
	value *ReportType
	isSet bool
}

func (v NullableReportType) Get() *ReportType {
	return v.value
}

func (v *NullableReportType) Set(val *ReportType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportType(val *ReportType) *NullableReportType {
	return &NullableReportType{value: val, isSet: true}
}

func (v NullableReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

