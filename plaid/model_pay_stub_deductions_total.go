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

// PayStubDeductionsTotal An object representing the total deductions for the pay period
type PayStubDeductionsTotal struct {
	// Raw amount of the deduction
	CurrentAmount NullableFloat64 `json:"current_amount"`
	// The ISO-4217 currency code of the line item. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the line item. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// The year-to-date total amount of the deductions
	YtdAmount NullableFloat64 `json:"ytd_amount"`
	AdditionalProperties map[string]interface{}
}

type _PayStubDeductionsTotal PayStubDeductionsTotal

// NewPayStubDeductionsTotal instantiates a new PayStubDeductionsTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStubDeductionsTotal(currentAmount NullableFloat64, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, ytdAmount NullableFloat64) *PayStubDeductionsTotal {
	this := PayStubDeductionsTotal{}
	this.CurrentAmount = currentAmount
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	this.YtdAmount = ytdAmount
	return &this
}

// NewPayStubDeductionsTotalWithDefaults instantiates a new PayStubDeductionsTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubDeductionsTotalWithDefaults() *PayStubDeductionsTotal {
	this := PayStubDeductionsTotal{}
	return &this
}

// GetCurrentAmount returns the CurrentAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubDeductionsTotal) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDeductionsTotal) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// SetCurrentAmount sets field value
func (o *PayStubDeductionsTotal) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDeductionsTotal) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDeductionsTotal) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *PayStubDeductionsTotal) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDeductionsTotal) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDeductionsTotal) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *PayStubDeductionsTotal) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetYtdAmount returns the YtdAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubDeductionsTotal) GetYtdAmount() float64 {
	if o == nil || o.YtdAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.YtdAmount.Get()
}

// GetYtdAmountOk returns a tuple with the YtdAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDeductionsTotal) GetYtdAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YtdAmount.Get(), o.YtdAmount.IsSet()
}

// SetYtdAmount sets field value
func (o *PayStubDeductionsTotal) SetYtdAmount(v float64) {
	o.YtdAmount.Set(&v)
}

func (o PayStubDeductionsTotal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if true {
		toSerialize["ytd_amount"] = o.YtdAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStubDeductionsTotal) UnmarshalJSON(bytes []byte) (err error) {
	varPayStubDeductionsTotal := _PayStubDeductionsTotal{}

	if err = json.Unmarshal(bytes, &varPayStubDeductionsTotal); err == nil {
		*o = PayStubDeductionsTotal(varPayStubDeductionsTotal)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "ytd_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStubDeductionsTotal struct {
	value *PayStubDeductionsTotal
	isSet bool
}

func (v NullablePayStubDeductionsTotal) Get() *PayStubDeductionsTotal {
	return v.value
}

func (v *NullablePayStubDeductionsTotal) Set(val *PayStubDeductionsTotal) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStubDeductionsTotal) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStubDeductionsTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStubDeductionsTotal(val *PayStubDeductionsTotal) *NullablePayStubDeductionsTotal {
	return &NullablePayStubDeductionsTotal{value: val, isSet: true}
}

func (v NullablePayStubDeductionsTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStubDeductionsTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


