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
)

// CreditPayStubEarnings An object representing both a breakdown of earnings on a pay stub and the total earnings.
type CreditPayStubEarnings struct {
	Breakdown []PayStubEarningsBreakdown `json:"breakdown"`
	Total PayStubEarningsTotal `json:"total"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayStubEarnings CreditPayStubEarnings

// NewCreditPayStubEarnings instantiates a new CreditPayStubEarnings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayStubEarnings(breakdown []PayStubEarningsBreakdown, total PayStubEarningsTotal) *CreditPayStubEarnings {
	this := CreditPayStubEarnings{}
	this.Breakdown = breakdown
	this.Total = total
	return &this
}

// NewCreditPayStubEarningsWithDefaults instantiates a new CreditPayStubEarnings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayStubEarningsWithDefaults() *CreditPayStubEarnings {
	this := CreditPayStubEarnings{}
	return &this
}

// GetBreakdown returns the Breakdown field value
func (o *CreditPayStubEarnings) GetBreakdown() []PayStubEarningsBreakdown {
	if o == nil {
		var ret []PayStubEarningsBreakdown
		return ret
	}

	return o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value
// and a boolean to check if the value has been set.
func (o *CreditPayStubEarnings) GetBreakdownOk() (*[]PayStubEarningsBreakdown, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Breakdown, true
}

// SetBreakdown sets field value
func (o *CreditPayStubEarnings) SetBreakdown(v []PayStubEarningsBreakdown) {
	o.Breakdown = v
}

// GetTotal returns the Total field value
func (o *CreditPayStubEarnings) GetTotal() PayStubEarningsTotal {
	if o == nil {
		var ret PayStubEarningsTotal
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CreditPayStubEarnings) GetTotalOk() (*PayStubEarningsTotal, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CreditPayStubEarnings) SetTotal(v PayStubEarningsTotal) {
	o.Total = v
}

func (o CreditPayStubEarnings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["breakdown"] = o.Breakdown
	}
	if true {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayStubEarnings) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayStubEarnings := _CreditPayStubEarnings{}

	if err = json.Unmarshal(bytes, &varCreditPayStubEarnings); err == nil {
		*o = CreditPayStubEarnings(varCreditPayStubEarnings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "breakdown")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayStubEarnings struct {
	value *CreditPayStubEarnings
	isSet bool
}

func (v NullableCreditPayStubEarnings) Get() *CreditPayStubEarnings {
	return v.value
}

func (v *NullableCreditPayStubEarnings) Set(val *CreditPayStubEarnings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayStubEarnings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayStubEarnings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayStubEarnings(val *CreditPayStubEarnings) *NullableCreditPayStubEarnings {
	return &NullableCreditPayStubEarnings{value: val, isSet: true}
}

func (v NullableCreditPayStubEarnings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayStubEarnings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


