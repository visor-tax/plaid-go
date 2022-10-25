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
)

// CreditSessionDocumentIncomeResult The details of a document income verification in Link
type CreditSessionDocumentIncomeResult struct {
	// The number of paystubs uploaded by the user.
	NumPaystubsUploaded int32 `json:"num_paystubs_uploaded"`
	// The number of w2s uploaded by the user.
	NumW2sUploaded int32 `json:"num_w2s_uploaded"`
}

// NewCreditSessionDocumentIncomeResult instantiates a new CreditSessionDocumentIncomeResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditSessionDocumentIncomeResult(numPaystubsUploaded int32, numW2sUploaded int32) *CreditSessionDocumentIncomeResult {
	this := CreditSessionDocumentIncomeResult{}
	this.NumPaystubsUploaded = numPaystubsUploaded
	this.NumW2sUploaded = numW2sUploaded
	return &this
}

// NewCreditSessionDocumentIncomeResultWithDefaults instantiates a new CreditSessionDocumentIncomeResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditSessionDocumentIncomeResultWithDefaults() *CreditSessionDocumentIncomeResult {
	this := CreditSessionDocumentIncomeResult{}
	return &this
}

// GetNumPaystubsUploaded returns the NumPaystubsUploaded field value
func (o *CreditSessionDocumentIncomeResult) GetNumPaystubsUploaded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPaystubsUploaded
}

// GetNumPaystubsUploadedOk returns a tuple with the NumPaystubsUploaded field value
// and a boolean to check if the value has been set.
func (o *CreditSessionDocumentIncomeResult) GetNumPaystubsUploadedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumPaystubsUploaded, true
}

// SetNumPaystubsUploaded sets field value
func (o *CreditSessionDocumentIncomeResult) SetNumPaystubsUploaded(v int32) {
	o.NumPaystubsUploaded = v
}

// GetNumW2sUploaded returns the NumW2sUploaded field value
func (o *CreditSessionDocumentIncomeResult) GetNumW2sUploaded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumW2sUploaded
}

// GetNumW2sUploadedOk returns a tuple with the NumW2sUploaded field value
// and a boolean to check if the value has been set.
func (o *CreditSessionDocumentIncomeResult) GetNumW2sUploadedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NumW2sUploaded, true
}

// SetNumW2sUploaded sets field value
func (o *CreditSessionDocumentIncomeResult) SetNumW2sUploaded(v int32) {
	o.NumW2sUploaded = v
}

func (o CreditSessionDocumentIncomeResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["num_paystubs_uploaded"] = o.NumPaystubsUploaded
	}
	if true {
		toSerialize["num_w2s_uploaded"] = o.NumW2sUploaded
	}
	return json.Marshal(toSerialize)
}

type NullableCreditSessionDocumentIncomeResult struct {
	value *CreditSessionDocumentIncomeResult
	isSet bool
}

func (v NullableCreditSessionDocumentIncomeResult) Get() *CreditSessionDocumentIncomeResult {
	return v.value
}

func (v *NullableCreditSessionDocumentIncomeResult) Set(val *CreditSessionDocumentIncomeResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditSessionDocumentIncomeResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditSessionDocumentIncomeResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditSessionDocumentIncomeResult(val *CreditSessionDocumentIncomeResult) *NullableCreditSessionDocumentIncomeResult {
	return &NullableCreditSessionDocumentIncomeResult{value: val, isSet: true}
}

func (v NullableCreditSessionDocumentIncomeResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditSessionDocumentIncomeResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


