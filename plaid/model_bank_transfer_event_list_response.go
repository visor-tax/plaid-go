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

// BankTransferEventListResponse Defines the response schema for `/bank_transfer/event/list`
type BankTransferEventListResponse struct {
	BankTransferEvents []BankTransferEvent `json:"bank_transfer_events"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferEventListResponse BankTransferEventListResponse

// NewBankTransferEventListResponse instantiates a new BankTransferEventListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferEventListResponse(bankTransferEvents []BankTransferEvent, requestId string) *BankTransferEventListResponse {
	this := BankTransferEventListResponse{}
	this.BankTransferEvents = bankTransferEvents
	this.RequestId = requestId
	return &this
}

// NewBankTransferEventListResponseWithDefaults instantiates a new BankTransferEventListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferEventListResponseWithDefaults() *BankTransferEventListResponse {
	this := BankTransferEventListResponse{}
	return &this
}

// GetBankTransferEvents returns the BankTransferEvents field value
func (o *BankTransferEventListResponse) GetBankTransferEvents() []BankTransferEvent {
	if o == nil {
		var ret []BankTransferEvent
		return ret
	}

	return o.BankTransferEvents
}

// GetBankTransferEventsOk returns a tuple with the BankTransferEvents field value
// and a boolean to check if the value has been set.
func (o *BankTransferEventListResponse) GetBankTransferEventsOk() (*[]BankTransferEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferEvents, true
}

// SetBankTransferEvents sets field value
func (o *BankTransferEventListResponse) SetBankTransferEvents(v []BankTransferEvent) {
	o.BankTransferEvents = v
}

// GetRequestId returns the RequestId field value
func (o *BankTransferEventListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferEventListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferEventListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferEventListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bank_transfer_events"] = o.BankTransferEvents
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferEventListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferEventListResponse := _BankTransferEventListResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferEventListResponse); err == nil {
		*o = BankTransferEventListResponse(varBankTransferEventListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bank_transfer_events")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferEventListResponse struct {
	value *BankTransferEventListResponse
	isSet bool
}

func (v NullableBankTransferEventListResponse) Get() *BankTransferEventListResponse {
	return v.value
}

func (v *NullableBankTransferEventListResponse) Set(val *BankTransferEventListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferEventListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferEventListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferEventListResponse(val *BankTransferEventListResponse) *NullableBankTransferEventListResponse {
	return &NullableBankTransferEventListResponse{value: val, isSet: true}
}

func (v NullableBankTransferEventListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferEventListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


