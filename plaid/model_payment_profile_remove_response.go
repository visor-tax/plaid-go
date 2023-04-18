/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.345.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaymentProfileRemoveResponse PaymentProfileRemoveResponse defines the response schema for `/payment_profile/remove`
type PaymentProfileRemoveResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _PaymentProfileRemoveResponse PaymentProfileRemoveResponse

// NewPaymentProfileRemoveResponse instantiates a new PaymentProfileRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProfileRemoveResponse(requestId string) *PaymentProfileRemoveResponse {
	this := PaymentProfileRemoveResponse{}
	this.RequestId = requestId
	return &this
}

// NewPaymentProfileRemoveResponseWithDefaults instantiates a new PaymentProfileRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProfileRemoveResponseWithDefaults() *PaymentProfileRemoveResponse {
	this := PaymentProfileRemoveResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *PaymentProfileRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PaymentProfileRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PaymentProfileRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o PaymentProfileRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentProfileRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentProfileRemoveResponse := _PaymentProfileRemoveResponse{}

	if err = json.Unmarshal(bytes, &varPaymentProfileRemoveResponse); err == nil {
		*o = PaymentProfileRemoveResponse(varPaymentProfileRemoveResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentProfileRemoveResponse struct {
	value *PaymentProfileRemoveResponse
	isSet bool
}

func (v NullablePaymentProfileRemoveResponse) Get() *PaymentProfileRemoveResponse {
	return v.value
}

func (v *NullablePaymentProfileRemoveResponse) Set(val *PaymentProfileRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileRemoveResponse(val *PaymentProfileRemoveResponse) *NullablePaymentProfileRemoveResponse {
	return &NullablePaymentProfileRemoveResponse{value: val, isSet: true}
}

func (v NullablePaymentProfileRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


