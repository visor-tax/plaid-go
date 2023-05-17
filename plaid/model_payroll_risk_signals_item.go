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

// PayrollRiskSignalsItem Object containing fraud risk data pertaining to the Item linked as part of the verification.
type PayrollRiskSignalsItem struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	// Array of payroll income document authenticity data retrieved for each of the user's accounts.
	VerificationRiskSignals []DocumentRiskSignalsObject `json:"verification_risk_signals"`
	AdditionalProperties map[string]interface{}
}

type _PayrollRiskSignalsItem PayrollRiskSignalsItem

// NewPayrollRiskSignalsItem instantiates a new PayrollRiskSignalsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayrollRiskSignalsItem(itemId string, verificationRiskSignals []DocumentRiskSignalsObject) *PayrollRiskSignalsItem {
	this := PayrollRiskSignalsItem{}
	this.ItemId = itemId
	this.VerificationRiskSignals = verificationRiskSignals
	return &this
}

// NewPayrollRiskSignalsItemWithDefaults instantiates a new PayrollRiskSignalsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayrollRiskSignalsItemWithDefaults() *PayrollRiskSignalsItem {
	this := PayrollRiskSignalsItem{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *PayrollRiskSignalsItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *PayrollRiskSignalsItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *PayrollRiskSignalsItem) SetItemId(v string) {
	o.ItemId = v
}

// GetVerificationRiskSignals returns the VerificationRiskSignals field value
func (o *PayrollRiskSignalsItem) GetVerificationRiskSignals() []DocumentRiskSignalsObject {
	if o == nil {
		var ret []DocumentRiskSignalsObject
		return ret
	}

	return o.VerificationRiskSignals
}

// GetVerificationRiskSignalsOk returns a tuple with the VerificationRiskSignals field value
// and a boolean to check if the value has been set.
func (o *PayrollRiskSignalsItem) GetVerificationRiskSignalsOk() (*[]DocumentRiskSignalsObject, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VerificationRiskSignals, true
}

// SetVerificationRiskSignals sets field value
func (o *PayrollRiskSignalsItem) SetVerificationRiskSignals(v []DocumentRiskSignalsObject) {
	o.VerificationRiskSignals = v
}

func (o PayrollRiskSignalsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["verification_risk_signals"] = o.VerificationRiskSignals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayrollRiskSignalsItem) UnmarshalJSON(bytes []byte) (err error) {
	varPayrollRiskSignalsItem := _PayrollRiskSignalsItem{}

	if err = json.Unmarshal(bytes, &varPayrollRiskSignalsItem); err == nil {
		*o = PayrollRiskSignalsItem(varPayrollRiskSignalsItem)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "verification_risk_signals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayrollRiskSignalsItem struct {
	value *PayrollRiskSignalsItem
	isSet bool
}

func (v NullablePayrollRiskSignalsItem) Get() *PayrollRiskSignalsItem {
	return v.value
}

func (v *NullablePayrollRiskSignalsItem) Set(val *PayrollRiskSignalsItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePayrollRiskSignalsItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePayrollRiskSignalsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayrollRiskSignalsItem(val *PayrollRiskSignalsItem) *NullablePayrollRiskSignalsItem {
	return &NullablePayrollRiskSignalsItem{value: val, isSet: true}
}

func (v NullablePayrollRiskSignalsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayrollRiskSignalsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


