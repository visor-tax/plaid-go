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

// FDXNotificationCategory Category of Notification
type FDXNotificationCategory string

var _ = fmt.Printf

// List of FDXNotificationCategory
const (
	FDXNOTIFICATIONCATEGORY_SECURITY FDXNotificationCategory = "SECURITY"
	FDXNOTIFICATIONCATEGORY_MAINTENANCE FDXNotificationCategory = "MAINTENANCE"
	FDXNOTIFICATIONCATEGORY_FRAUD FDXNotificationCategory = "FRAUD"
	FDXNOTIFICATIONCATEGORY_CONSENT FDXNotificationCategory = "CONSENT"
	FDXNOTIFICATIONCATEGORY_NEW_DATA FDXNotificationCategory = "NEW_DATA"
)

var allowedFDXNotificationCategoryEnumValues = []FDXNotificationCategory{
	"SECURITY",
	"MAINTENANCE",
	"FRAUD",
	"CONSENT",
	"NEW_DATA",
}

func (v *FDXNotificationCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := FDXNotificationCategory(value)


	*v = enumTypeValue
	return nil
}

// NewFDXNotificationCategoryFromValue returns a pointer to a valid FDXNotificationCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFDXNotificationCategoryFromValue(v string) (*FDXNotificationCategory, error) {
	ev := FDXNotificationCategory(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FDXNotificationCategory) IsValid() bool {
	for _, existing := range allowedFDXNotificationCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FDXNotificationCategory value
func (v FDXNotificationCategory) Ptr() *FDXNotificationCategory {
	return &v
}

type NullableFDXNotificationCategory struct {
	value *FDXNotificationCategory
	isSet bool
}

func (v NullableFDXNotificationCategory) Get() *FDXNotificationCategory {
	return v.value
}

func (v *NullableFDXNotificationCategory) Set(val *FDXNotificationCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXNotificationCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXNotificationCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXNotificationCategory(val *FDXNotificationCategory) *NullableFDXNotificationCategory {
	return &NullableFDXNotificationCategory{value: val, isSet: true}
}

func (v NullableFDXNotificationCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXNotificationCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

