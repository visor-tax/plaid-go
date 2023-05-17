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

// BankTransferCancelRequest Defines the request schema for `/bank_transfer/cancel`
type BankTransferCancelRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Plaid’s unique identifier for a bank transfer.
	BankTransferId string `json:"bank_transfer_id"`
}

// NewBankTransferCancelRequest instantiates a new BankTransferCancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferCancelRequest(bankTransferId string) *BankTransferCancelRequest {
	this := BankTransferCancelRequest{}
	this.BankTransferId = bankTransferId
	return &this
}

// NewBankTransferCancelRequestWithDefaults instantiates a new BankTransferCancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferCancelRequestWithDefaults() *BankTransferCancelRequest {
	this := BankTransferCancelRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransferCancelRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferCancelRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransferCancelRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransferCancelRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BankTransferCancelRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferCancelRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BankTransferCancelRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BankTransferCancelRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetBankTransferId returns the BankTransferId field value
func (o *BankTransferCancelRequest) GetBankTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankTransferId
}

// GetBankTransferIdOk returns a tuple with the BankTransferId field value
// and a boolean to check if the value has been set.
func (o *BankTransferCancelRequest) GetBankTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankTransferId, true
}

// SetBankTransferId sets field value
func (o *BankTransferCancelRequest) SetBankTransferId(v string) {
	o.BankTransferId = v
}

func (o BankTransferCancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["bank_transfer_id"] = o.BankTransferId
	}
	return json.Marshal(toSerialize)
}

type NullableBankTransferCancelRequest struct {
	value *BankTransferCancelRequest
	isSet bool
}

func (v NullableBankTransferCancelRequest) Get() *BankTransferCancelRequest {
	return v.value
}

func (v *NullableBankTransferCancelRequest) Set(val *BankTransferCancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferCancelRequest(val *BankTransferCancelRequest) *NullableBankTransferCancelRequest {
	return &NullableBankTransferCancelRequest{value: val, isSet: true}
}

func (v NullableBankTransferCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


