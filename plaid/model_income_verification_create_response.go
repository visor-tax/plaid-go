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

// IncomeVerificationCreateResponse IncomeVerificationCreateResponse defines the response schema for `/income/verification/create`.
type IncomeVerificationCreateResponse struct {
	// ID of the verification. This ID is persisted throughout the lifetime of the verification.
	IncomeVerificationId string `json:"income_verification_id"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _IncomeVerificationCreateResponse IncomeVerificationCreateResponse

// NewIncomeVerificationCreateResponse instantiates a new IncomeVerificationCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationCreateResponse(incomeVerificationId string, requestId string) *IncomeVerificationCreateResponse {
	this := IncomeVerificationCreateResponse{}
	this.IncomeVerificationId = incomeVerificationId
	this.RequestId = requestId
	return &this
}

// NewIncomeVerificationCreateResponseWithDefaults instantiates a new IncomeVerificationCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationCreateResponseWithDefaults() *IncomeVerificationCreateResponse {
	this := IncomeVerificationCreateResponse{}
	return &this
}

// GetIncomeVerificationId returns the IncomeVerificationId field value
func (o *IncomeVerificationCreateResponse) GetIncomeVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IncomeVerificationId
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateResponse) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomeVerificationId, true
}

// SetIncomeVerificationId sets field value
func (o *IncomeVerificationCreateResponse) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId = v
}

// GetRequestId returns the RequestId field value
func (o *IncomeVerificationCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IncomeVerificationCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o IncomeVerificationCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["income_verification_id"] = o.IncomeVerificationId
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeVerificationCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeVerificationCreateResponse := _IncomeVerificationCreateResponse{}

	if err = json.Unmarshal(bytes, &varIncomeVerificationCreateResponse); err == nil {
		*o = IncomeVerificationCreateResponse(varIncomeVerificationCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "income_verification_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeVerificationCreateResponse struct {
	value *IncomeVerificationCreateResponse
	isSet bool
}

func (v NullableIncomeVerificationCreateResponse) Get() *IncomeVerificationCreateResponse {
	return v.value
}

func (v *NullableIncomeVerificationCreateResponse) Set(val *IncomeVerificationCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationCreateResponse(val *IncomeVerificationCreateResponse) *NullableIncomeVerificationCreateResponse {
	return &NullableIncomeVerificationCreateResponse{value: val, isSet: true}
}

func (v NullableIncomeVerificationCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


