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

// PayrollIncomeRateOfPay An object representing the rate at which an individual is paid.
type PayrollIncomeRateOfPay struct {
	// The rate at which an employee is paid.
	PayRate NullableString `json:"pay_rate,omitempty"`
	// The amount at which an employee is paid.
	PayAmount NullableFloat64 `json:"pay_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PayrollIncomeRateOfPay PayrollIncomeRateOfPay

// NewPayrollIncomeRateOfPay instantiates a new PayrollIncomeRateOfPay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayrollIncomeRateOfPay() *PayrollIncomeRateOfPay {
	this := PayrollIncomeRateOfPay{}
	return &this
}

// NewPayrollIncomeRateOfPayWithDefaults instantiates a new PayrollIncomeRateOfPay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayrollIncomeRateOfPayWithDefaults() *PayrollIncomeRateOfPay {
	this := PayrollIncomeRateOfPay{}
	return &this
}

// GetPayRate returns the PayRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayrollIncomeRateOfPay) GetPayRate() string {
	if o == nil || o.PayRate.Get() == nil {
		var ret string
		return ret
	}
	return *o.PayRate.Get()
}

// GetPayRateOk returns a tuple with the PayRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayrollIncomeRateOfPay) GetPayRateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayRate.Get(), o.PayRate.IsSet()
}

// HasPayRate returns a boolean if a field has been set.
func (o *PayrollIncomeRateOfPay) HasPayRate() bool {
	if o != nil && o.PayRate.IsSet() {
		return true
	}

	return false
}

// SetPayRate gets a reference to the given NullableString and assigns it to the PayRate field.
func (o *PayrollIncomeRateOfPay) SetPayRate(v string) {
	o.PayRate.Set(&v)
}
// SetPayRateNil sets the value for PayRate to be an explicit nil
func (o *PayrollIncomeRateOfPay) SetPayRateNil() {
	o.PayRate.Set(nil)
}

// UnsetPayRate ensures that no value is present for PayRate, not even an explicit nil
func (o *PayrollIncomeRateOfPay) UnsetPayRate() {
	o.PayRate.Unset()
}

// GetPayAmount returns the PayAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PayrollIncomeRateOfPay) GetPayAmount() float64 {
	if o == nil || o.PayAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PayAmount.Get()
}

// GetPayAmountOk returns a tuple with the PayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayrollIncomeRateOfPay) GetPayAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayAmount.Get(), o.PayAmount.IsSet()
}

// HasPayAmount returns a boolean if a field has been set.
func (o *PayrollIncomeRateOfPay) HasPayAmount() bool {
	if o != nil && o.PayAmount.IsSet() {
		return true
	}

	return false
}

// SetPayAmount gets a reference to the given NullableFloat64 and assigns it to the PayAmount field.
func (o *PayrollIncomeRateOfPay) SetPayAmount(v float64) {
	o.PayAmount.Set(&v)
}
// SetPayAmountNil sets the value for PayAmount to be an explicit nil
func (o *PayrollIncomeRateOfPay) SetPayAmountNil() {
	o.PayAmount.Set(nil)
}

// UnsetPayAmount ensures that no value is present for PayAmount, not even an explicit nil
func (o *PayrollIncomeRateOfPay) UnsetPayAmount() {
	o.PayAmount.Unset()
}

func (o PayrollIncomeRateOfPay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayRate.IsSet() {
		toSerialize["pay_rate"] = o.PayRate.Get()
	}
	if o.PayAmount.IsSet() {
		toSerialize["pay_amount"] = o.PayAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayrollIncomeRateOfPay) UnmarshalJSON(bytes []byte) (err error) {
	varPayrollIncomeRateOfPay := _PayrollIncomeRateOfPay{}

	if err = json.Unmarshal(bytes, &varPayrollIncomeRateOfPay); err == nil {
		*o = PayrollIncomeRateOfPay(varPayrollIncomeRateOfPay)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pay_rate")
		delete(additionalProperties, "pay_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayrollIncomeRateOfPay struct {
	value *PayrollIncomeRateOfPay
	isSet bool
}

func (v NullablePayrollIncomeRateOfPay) Get() *PayrollIncomeRateOfPay {
	return v.value
}

func (v *NullablePayrollIncomeRateOfPay) Set(val *PayrollIncomeRateOfPay) {
	v.value = val
	v.isSet = true
}

func (v NullablePayrollIncomeRateOfPay) IsSet() bool {
	return v.isSet
}

func (v *NullablePayrollIncomeRateOfPay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayrollIncomeRateOfPay(val *PayrollIncomeRateOfPay) *NullablePayrollIncomeRateOfPay {
	return &NullablePayrollIncomeRateOfPay{value: val, isSet: true}
}

func (v NullablePayrollIncomeRateOfPay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayrollIncomeRateOfPay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


