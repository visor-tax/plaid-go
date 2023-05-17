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

// CreditBankEmploymentGetResponse CreditBankEmploymentGetResponse defines the response schema for `/beta/credit/v1/bank_employment/get`.
type CreditBankEmploymentGetResponse struct {
	// Bank Employment data. Each entry in the array will be a distinct bank employment report.
	BankEmploymentReports []CreditBankEmploymentReport `json:"bank_employment_reports"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _CreditBankEmploymentGetResponse CreditBankEmploymentGetResponse

// NewCreditBankEmploymentGetResponse instantiates a new CreditBankEmploymentGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankEmploymentGetResponse(bankEmploymentReports []CreditBankEmploymentReport, requestId string) *CreditBankEmploymentGetResponse {
	this := CreditBankEmploymentGetResponse{}
	this.BankEmploymentReports = bankEmploymentReports
	this.RequestId = requestId
	return &this
}

// NewCreditBankEmploymentGetResponseWithDefaults instantiates a new CreditBankEmploymentGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankEmploymentGetResponseWithDefaults() *CreditBankEmploymentGetResponse {
	this := CreditBankEmploymentGetResponse{}
	return &this
}

// GetBankEmploymentReports returns the BankEmploymentReports field value
func (o *CreditBankEmploymentGetResponse) GetBankEmploymentReports() []CreditBankEmploymentReport {
	if o == nil {
		var ret []CreditBankEmploymentReport
		return ret
	}

	return o.BankEmploymentReports
}

// GetBankEmploymentReportsOk returns a tuple with the BankEmploymentReports field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentGetResponse) GetBankEmploymentReportsOk() (*[]CreditBankEmploymentReport, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankEmploymentReports, true
}

// SetBankEmploymentReports sets field value
func (o *CreditBankEmploymentGetResponse) SetBankEmploymentReports(v []CreditBankEmploymentReport) {
	o.BankEmploymentReports = v
}

// GetRequestId returns the RequestId field value
func (o *CreditBankEmploymentGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *CreditBankEmploymentGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreditBankEmploymentGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_employment_reports"] = o.BankEmploymentReports
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditBankEmploymentGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCreditBankEmploymentGetResponse := _CreditBankEmploymentGetResponse{}

	if err = json.Unmarshal(bytes, &varCreditBankEmploymentGetResponse); err == nil {
		*o = CreditBankEmploymentGetResponse(varCreditBankEmploymentGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_employment_reports")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditBankEmploymentGetResponse struct {
	value *CreditBankEmploymentGetResponse
	isSet bool
}

func (v NullableCreditBankEmploymentGetResponse) Get() *CreditBankEmploymentGetResponse {
	return v.value
}

func (v *NullableCreditBankEmploymentGetResponse) Set(val *CreditBankEmploymentGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmploymentGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmploymentGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmploymentGetResponse(val *CreditBankEmploymentGetResponse) *NullableCreditBankEmploymentGetResponse {
	return &NullableCreditBankEmploymentGetResponse{value: val, isSet: true}
}

func (v NullableCreditBankEmploymentGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmploymentGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


