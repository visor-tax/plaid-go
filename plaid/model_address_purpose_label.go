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

// AddressPurposeLabel Field describing whether the associated address is being used for commercial or residential purposes.  Note: This value will be `no_data` when Plaid does not have sufficient data to determine the address's use.
type AddressPurposeLabel string

var _ = fmt.Printf

// List of AddressPurposeLabel
const (
	ADDRESSPURPOSELABEL_RESIDENTIAL AddressPurposeLabel = "residential"
	ADDRESSPURPOSELABEL_COMMERCIAL AddressPurposeLabel = "commercial"
	ADDRESSPURPOSELABEL_NO_DATA AddressPurposeLabel = "no_data"
)

var allowedAddressPurposeLabelEnumValues = []AddressPurposeLabel{
	"residential",
	"commercial",
	"no_data",
}

func (v *AddressPurposeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := AddressPurposeLabel(value)


	*v = enumTypeValue
	return nil
}

// NewAddressPurposeLabelFromValue returns a pointer to a valid AddressPurposeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAddressPurposeLabelFromValue(v string) (*AddressPurposeLabel, error) {
	ev := AddressPurposeLabel(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AddressPurposeLabel) IsValid() bool {
	for _, existing := range allowedAddressPurposeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AddressPurposeLabel value
func (v AddressPurposeLabel) Ptr() *AddressPurposeLabel {
	return &v
}

type NullableAddressPurposeLabel struct {
	value *AddressPurposeLabel
	isSet bool
}

func (v NullableAddressPurposeLabel) Get() *AddressPurposeLabel {
	return v.value
}

func (v *NullableAddressPurposeLabel) Set(val *AddressPurposeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressPurposeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressPurposeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressPurposeLabel(val *AddressPurposeLabel) *NullableAddressPurposeLabel {
	return &NullableAddressPurposeLabel{value: val, isSet: true}
}

func (v NullableAddressPurposeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressPurposeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

