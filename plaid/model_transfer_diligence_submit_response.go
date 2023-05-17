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

// TransferDiligenceSubmitResponse Defines the response schema for `/transfer/diligence/submit`
type TransferDiligenceSubmitResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferDiligenceSubmitResponse TransferDiligenceSubmitResponse

// NewTransferDiligenceSubmitResponse instantiates a new TransferDiligenceSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferDiligenceSubmitResponse(requestId string) *TransferDiligenceSubmitResponse {
	this := TransferDiligenceSubmitResponse{}
	this.RequestId = requestId
	return &this
}

// NewTransferDiligenceSubmitResponseWithDefaults instantiates a new TransferDiligenceSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDiligenceSubmitResponseWithDefaults() *TransferDiligenceSubmitResponse {
	this := TransferDiligenceSubmitResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *TransferDiligenceSubmitResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferDiligenceSubmitResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferDiligenceSubmitResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferDiligenceSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferDiligenceSubmitResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferDiligenceSubmitResponse := _TransferDiligenceSubmitResponse{}

	if err = json.Unmarshal(bytes, &varTransferDiligenceSubmitResponse); err == nil {
		*o = TransferDiligenceSubmitResponse(varTransferDiligenceSubmitResponse)
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

type NullableTransferDiligenceSubmitResponse struct {
	value *TransferDiligenceSubmitResponse
	isSet bool
}

func (v NullableTransferDiligenceSubmitResponse) Get() *TransferDiligenceSubmitResponse {
	return v.value
}

func (v *NullableTransferDiligenceSubmitResponse) Set(val *TransferDiligenceSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDiligenceSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDiligenceSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDiligenceSubmitResponse(val *TransferDiligenceSubmitResponse) *NullableTransferDiligenceSubmitResponse {
	return &NullableTransferDiligenceSubmitResponse{value: val, isSet: true}
}

func (v NullableTransferDiligenceSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDiligenceSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


