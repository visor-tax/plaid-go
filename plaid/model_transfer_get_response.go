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

// TransferGetResponse Defines the response schema for `/transfer/get`
type TransferGetResponse struct {
	Transfer Transfer `json:"transfer"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferGetResponse TransferGetResponse

// NewTransferGetResponse instantiates a new TransferGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferGetResponse(transfer Transfer, requestId string) *TransferGetResponse {
	this := TransferGetResponse{}
	this.Transfer = transfer
	this.RequestId = requestId
	return &this
}

// NewTransferGetResponseWithDefaults instantiates a new TransferGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferGetResponseWithDefaults() *TransferGetResponse {
	this := TransferGetResponse{}
	return &this
}

// GetTransfer returns the Transfer field value
func (o *TransferGetResponse) GetTransfer() Transfer {
	if o == nil {
		var ret Transfer
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *TransferGetResponse) GetTransferOk() (*Transfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *TransferGetResponse) SetTransfer(v Transfer) {
	o.Transfer = v
}

// GetRequestId returns the RequestId field value
func (o *TransferGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer"] = o.Transfer
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferGetResponse := _TransferGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferGetResponse); err == nil {
		*o = TransferGetResponse(varTransferGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferGetResponse struct {
	value *TransferGetResponse
	isSet bool
}

func (v NullableTransferGetResponse) Get() *TransferGetResponse {
	return v.value
}

func (v *NullableTransferGetResponse) Set(val *TransferGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferGetResponse(val *TransferGetResponse) *NullableTransferGetResponse {
	return &NullableTransferGetResponse{value: val, isSet: true}
}

func (v NullableTransferGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


