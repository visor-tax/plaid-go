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

// OwnershipType How an asset is owned.  `association`: Ownership by a corporation, partnership, or unincorporated association, including for-profit and not-for-profit organizations. `individual`: Ownership by an individual. `joint`: Joint ownership by multiple parties. `trust`: Ownership by a revocable or irrevocable trust.
type OwnershipType string

var _ = fmt.Printf

// List of OwnershipType
const (
	OWNERSHIPTYPE_NULL OwnershipType = "null"
	OWNERSHIPTYPE_INDIVIDUAL OwnershipType = "individual"
	OWNERSHIPTYPE_JOINT OwnershipType = "joint"
	OWNERSHIPTYPE_ASSOCIATION OwnershipType = "association"
	OWNERSHIPTYPE_TRUST OwnershipType = "trust"
)

var allowedOwnershipTypeEnumValues = []OwnershipType{
	"null",
	"individual",
	"joint",
	"association",
	"trust",
}

func (v *OwnershipType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := OwnershipType(value)


	*v = enumTypeValue
	return nil
}

// NewOwnershipTypeFromValue returns a pointer to a valid OwnershipType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOwnershipTypeFromValue(v string) (*OwnershipType, error) {
	ev := OwnershipType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OwnershipType) IsValid() bool {
	for _, existing := range allowedOwnershipTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OwnershipType value
func (v OwnershipType) Ptr() *OwnershipType {
	return &v
}

type NullableOwnershipType struct {
	value *OwnershipType
	isSet bool
}

func (v NullableOwnershipType) Get() *OwnershipType {
	return v.value
}

func (v *NullableOwnershipType) Set(val *OwnershipType) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnershipType) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnershipType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnershipType(val *OwnershipType) *NullableOwnershipType {
	return &NullableOwnershipType{value: val, isSet: true}
}

func (v NullableOwnershipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnershipType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

