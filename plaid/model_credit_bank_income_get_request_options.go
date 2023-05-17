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

// CreditBankIncomeGetRequestOptions An optional object for `/credit/bank_income/get` request options.
type CreditBankIncomeGetRequestOptions struct {
	// How many Bank Income Reports should be fetched. Multiple reports may be available if the report has been re-created or refreshed. If more than one report is available, the most recent reports will be returned first.
	Count *int32 `json:"count,omitempty"`
}

// NewCreditBankIncomeGetRequestOptions instantiates a new CreditBankIncomeGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankIncomeGetRequestOptions() *CreditBankIncomeGetRequestOptions {
	this := CreditBankIncomeGetRequestOptions{}
	var count int32 = 1
	this.Count = &count
	return &this
}

// NewCreditBankIncomeGetRequestOptionsWithDefaults instantiates a new CreditBankIncomeGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankIncomeGetRequestOptionsWithDefaults() *CreditBankIncomeGetRequestOptions {
	this := CreditBankIncomeGetRequestOptions{}
	var count int32 = 1
	this.Count = &count
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CreditBankIncomeGetRequestOptions) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditBankIncomeGetRequestOptions) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CreditBankIncomeGetRequestOptions) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CreditBankIncomeGetRequestOptions) SetCount(v int32) {
	o.Count = &v
}

func (o CreditBankIncomeGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankIncomeGetRequestOptions struct {
	value *CreditBankIncomeGetRequestOptions
	isSet bool
}

func (v NullableCreditBankIncomeGetRequestOptions) Get() *CreditBankIncomeGetRequestOptions {
	return v.value
}

func (v *NullableCreditBankIncomeGetRequestOptions) Set(val *CreditBankIncomeGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankIncomeGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankIncomeGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankIncomeGetRequestOptions(val *CreditBankIncomeGetRequestOptions) *NullableCreditBankIncomeGetRequestOptions {
	return &NullableCreditBankIncomeGetRequestOptions{value: val, isSet: true}
}

func (v NullableCreditBankIncomeGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankIncomeGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


