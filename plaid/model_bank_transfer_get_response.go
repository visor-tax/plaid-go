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

// BankTransferGetResponse Defines the response schema for `/bank_transfer/get`
type BankTransferGetResponse struct {
	BankTransfer BankTransfer `json:"bank_transfer"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferGetResponse BankTransferGetResponse

// NewBankTransferGetResponse instantiates a new BankTransferGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferGetResponse(bankTransfer BankTransfer, requestId string) *BankTransferGetResponse {
	this := BankTransferGetResponse{}
	this.BankTransfer = bankTransfer
	this.RequestId = requestId
	return &this
}

// NewBankTransferGetResponseWithDefaults instantiates a new BankTransferGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferGetResponseWithDefaults() *BankTransferGetResponse {
	this := BankTransferGetResponse{}
	return &this
}

// GetBankTransfer returns the BankTransfer field value
func (o *BankTransferGetResponse) GetBankTransfer() BankTransfer {
	if o == nil {
		var ret BankTransfer
		return ret
	}

	return o.BankTransfer
}

// GetBankTransferOk returns a tuple with the BankTransfer field value
// and a boolean to check if the value has been set.
func (o *BankTransferGetResponse) GetBankTransferOk() (*BankTransfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransfer, true
}

// SetBankTransfer sets field value
func (o *BankTransferGetResponse) SetBankTransfer(v BankTransfer) {
	o.BankTransfer = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_transfer"] = o.BankTransfer
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferGetResponse := _BankTransferGetResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferGetResponse); err == nil {
		*o = BankTransferGetResponse(varBankTransferGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_transfer")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferGetResponse struct {
	value *BankTransferGetResponse
	isSet bool
}

func (v NullableBankTransferGetResponse) Get() *BankTransferGetResponse {
	return v.value
}

func (v *NullableBankTransferGetResponse) Set(val *BankTransferGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferGetResponse(val *BankTransferGetResponse) *NullableBankTransferGetResponse {
	return &NullableBankTransferGetResponse{value: val, isSet: true}
}

func (v NullableBankTransferGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


