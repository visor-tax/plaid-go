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

// AssetTransactionType Asset Transaction Type.
type AssetTransactionType string

var _ = fmt.Printf

// List of AssetTransactionType
const (
	ASSETTRANSACTIONTYPE_CREDIT AssetTransactionType = "Credit"
	ASSETTRANSACTIONTYPE_DEBIT AssetTransactionType = "Debit"
)

var allowedAssetTransactionTypeEnumValues = []AssetTransactionType{
	"Credit",
	"Debit",
}

func (v *AssetTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := AssetTransactionType(value)


	*v = enumTypeValue
	return nil
}

// NewAssetTransactionTypeFromValue returns a pointer to a valid AssetTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetTransactionTypeFromValue(v string) (*AssetTransactionType, error) {
	ev := AssetTransactionType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetTransactionType) IsValid() bool {
	for _, existing := range allowedAssetTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetTransactionType value
func (v AssetTransactionType) Ptr() *AssetTransactionType {
	return &v
}

type NullableAssetTransactionType struct {
	value *AssetTransactionType
	isSet bool
}

func (v NullableAssetTransactionType) Get() *AssetTransactionType {
	return v.value
}

func (v *NullableAssetTransactionType) Set(val *AssetTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetTransactionType(val *AssetTransactionType) *NullableAssetTransactionType {
	return &NullableAssetTransactionType{value: val, isSet: true}
}

func (v NullableAssetTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

