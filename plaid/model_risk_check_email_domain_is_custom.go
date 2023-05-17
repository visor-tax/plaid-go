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

// RiskCheckEmailDomainIsCustom Indicates whether the email address domain is custom if known, i.e. a company domain and not free or disposable.
type RiskCheckEmailDomainIsCustom string

var _ = fmt.Printf

// List of RiskCheckEmailDomainIsCustom
const (
	RISKCHECKEMAILDOMAINISCUSTOM_YES RiskCheckEmailDomainIsCustom = "yes"
	RISKCHECKEMAILDOMAINISCUSTOM_NO RiskCheckEmailDomainIsCustom = "no"
	RISKCHECKEMAILDOMAINISCUSTOM_NO_DATA RiskCheckEmailDomainIsCustom = "no_data"
)

var allowedRiskCheckEmailDomainIsCustomEnumValues = []RiskCheckEmailDomainIsCustom{
	"yes",
	"no",
	"no_data",
}

func (v *RiskCheckEmailDomainIsCustom) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := RiskCheckEmailDomainIsCustom(value)


	*v = enumTypeValue
	return nil
}

// NewRiskCheckEmailDomainIsCustomFromValue returns a pointer to a valid RiskCheckEmailDomainIsCustom
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRiskCheckEmailDomainIsCustomFromValue(v string) (*RiskCheckEmailDomainIsCustom, error) {
	ev := RiskCheckEmailDomainIsCustom(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RiskCheckEmailDomainIsCustom) IsValid() bool {
	for _, existing := range allowedRiskCheckEmailDomainIsCustomEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RiskCheckEmailDomainIsCustom value
func (v RiskCheckEmailDomainIsCustom) Ptr() *RiskCheckEmailDomainIsCustom {
	return &v
}

type NullableRiskCheckEmailDomainIsCustom struct {
	value *RiskCheckEmailDomainIsCustom
	isSet bool
}

func (v NullableRiskCheckEmailDomainIsCustom) Get() *RiskCheckEmailDomainIsCustom {
	return v.value
}

func (v *NullableRiskCheckEmailDomainIsCustom) Set(val *RiskCheckEmailDomainIsCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskCheckEmailDomainIsCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskCheckEmailDomainIsCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskCheckEmailDomainIsCustom(val *RiskCheckEmailDomainIsCustom) *NullableRiskCheckEmailDomainIsCustom {
	return &NullableRiskCheckEmailDomainIsCustom{value: val, isSet: true}
}

func (v NullableRiskCheckEmailDomainIsCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskCheckEmailDomainIsCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

