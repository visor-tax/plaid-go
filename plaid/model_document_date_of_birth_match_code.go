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

// DocumentDateOfBirthMatchCode A match summary describing the cross comparison between the subject's date of birth, extracted from the document image, and the date of birth they separately provided to the identity verification attempt.
type DocumentDateOfBirthMatchCode string

var _ = fmt.Printf

// List of DocumentDateOfBirthMatchCode
const (
	DOCUMENTDATEOFBIRTHMATCHCODE_MATCH DocumentDateOfBirthMatchCode = "match"
	DOCUMENTDATEOFBIRTHMATCHCODE_PARTIAL_MATCH DocumentDateOfBirthMatchCode = "partial_match"
	DOCUMENTDATEOFBIRTHMATCHCODE_NO_MATCH DocumentDateOfBirthMatchCode = "no_match"
)

var allowedDocumentDateOfBirthMatchCodeEnumValues = []DocumentDateOfBirthMatchCode{
	"match",
	"partial_match",
	"no_match",
}

func (v *DocumentDateOfBirthMatchCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := DocumentDateOfBirthMatchCode(value)


	*v = enumTypeValue
	return nil
}

// NewDocumentDateOfBirthMatchCodeFromValue returns a pointer to a valid DocumentDateOfBirthMatchCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentDateOfBirthMatchCodeFromValue(v string) (*DocumentDateOfBirthMatchCode, error) {
	ev := DocumentDateOfBirthMatchCode(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentDateOfBirthMatchCode) IsValid() bool {
	for _, existing := range allowedDocumentDateOfBirthMatchCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DocumentDateOfBirthMatchCode value
func (v DocumentDateOfBirthMatchCode) Ptr() *DocumentDateOfBirthMatchCode {
	return &v
}

type NullableDocumentDateOfBirthMatchCode struct {
	value *DocumentDateOfBirthMatchCode
	isSet bool
}

func (v NullableDocumentDateOfBirthMatchCode) Get() *DocumentDateOfBirthMatchCode {
	return v.value
}

func (v *NullableDocumentDateOfBirthMatchCode) Set(val *DocumentDateOfBirthMatchCode) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentDateOfBirthMatchCode) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentDateOfBirthMatchCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentDateOfBirthMatchCode(val *DocumentDateOfBirthMatchCode) *NullableDocumentDateOfBirthMatchCode {
	return &NullableDocumentDateOfBirthMatchCode{value: val, isSet: true}
}

func (v NullableDocumentDateOfBirthMatchCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentDateOfBirthMatchCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

