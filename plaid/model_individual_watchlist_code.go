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

// IndividualWatchlistCode Shorthand identifier for a specific screening list for individuals.
type IndividualWatchlistCode string

// List of IndividualWatchlistCode
const (
	INDIVIDUALWATCHLISTCODE_AU_CON IndividualWatchlistCode = "AU_CON"
	INDIVIDUALWATCHLISTCODE_CA_CON IndividualWatchlistCode = "CA_CON"
	INDIVIDUALWATCHLISTCODE_EU_CON IndividualWatchlistCode = "EU_CON"
	INDIVIDUALWATCHLISTCODE_IZ_CIA IndividualWatchlistCode = "IZ_CIA"
	INDIVIDUALWATCHLISTCODE_IZ_IPL IndividualWatchlistCode = "IZ_IPL"
	INDIVIDUALWATCHLISTCODE_IZ_PEP IndividualWatchlistCode = "IZ_PEP"
	INDIVIDUALWATCHLISTCODE_IZ_UNC IndividualWatchlistCode = "IZ_UNC"
	INDIVIDUALWATCHLISTCODE_UK_HMC IndividualWatchlistCode = "UK_HMC"
	INDIVIDUALWATCHLISTCODE_US_DPL IndividualWatchlistCode = "US_DPL"
	INDIVIDUALWATCHLISTCODE_US_DTC IndividualWatchlistCode = "US_DTC"
	INDIVIDUALWATCHLISTCODE_US_FBI IndividualWatchlistCode = "US_FBI"
	INDIVIDUALWATCHLISTCODE_US_FSE IndividualWatchlistCode = "US_FSE"
	INDIVIDUALWATCHLISTCODE_US_ISN IndividualWatchlistCode = "US_ISN"
	INDIVIDUALWATCHLISTCODE_US_MBC IndividualWatchlistCode = "US_MBC"
	INDIVIDUALWATCHLISTCODE_US_PLC IndividualWatchlistCode = "US_PLC"
	INDIVIDUALWATCHLISTCODE_US_SDN IndividualWatchlistCode = "US_SDN"
	INDIVIDUALWATCHLISTCODE_US_SSI IndividualWatchlistCode = "US_SSI"
	INDIVIDUALWATCHLISTCODE_SG_SOF IndividualWatchlistCode = "SG_SOF"
	INDIVIDUALWATCHLISTCODE_TR_TWL IndividualWatchlistCode = "TR_TWL"
	INDIVIDUALWATCHLISTCODE_TR_DFD IndividualWatchlistCode = "TR_DFD"
	INDIVIDUALWATCHLISTCODE_TR_FOR IndividualWatchlistCode = "TR_FOR"
	INDIVIDUALWATCHLISTCODE_TR_WMD IndividualWatchlistCode = "TR_WMD"
)

var allowedIndividualWatchlistCodeEnumValues = []IndividualWatchlistCode{
	"AU_CON",
	"CA_CON",
	"EU_CON",
	"IZ_CIA",
	"IZ_IPL",
	"IZ_PEP",
	"IZ_UNC",
	"UK_HMC",
	"US_DPL",
	"US_DTC",
	"US_FBI",
	"US_FSE",
	"US_ISN",
	"US_MBC",
	"US_PLC",
	"US_SDN",
	"US_SSI",
	"SG_SOF",
	"TR_TWL",
	"TR_DFD",
	"TR_FOR",
	"TR_WMD",
}

func (v *IndividualWatchlistCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IndividualWatchlistCode(value)
	for _, existing := range allowedIndividualWatchlistCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IndividualWatchlistCode", value)
}

// NewIndividualWatchlistCodeFromValue returns a pointer to a valid IndividualWatchlistCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndividualWatchlistCodeFromValue(v string) (*IndividualWatchlistCode, error) {
	ev := IndividualWatchlistCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IndividualWatchlistCode: valid values are %v", v, allowedIndividualWatchlistCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IndividualWatchlistCode) IsValid() bool {
	for _, existing := range allowedIndividualWatchlistCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IndividualWatchlistCode value
func (v IndividualWatchlistCode) Ptr() *IndividualWatchlistCode {
	return &v
}

type NullableIndividualWatchlistCode struct {
	value *IndividualWatchlistCode
	isSet bool
}

func (v NullableIndividualWatchlistCode) Get() *IndividualWatchlistCode {
	return v.value
}

func (v *NullableIndividualWatchlistCode) Set(val *IndividualWatchlistCode) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualWatchlistCode) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualWatchlistCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualWatchlistCode(val *IndividualWatchlistCode) *NullableIndividualWatchlistCode {
	return &NullableIndividualWatchlistCode{value: val, isSet: true}
}

func (v NullableIndividualWatchlistCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualWatchlistCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

