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

// RemovedTransaction A representation of a removed transaction
type RemovedTransaction struct {
	// The ID of the removed transaction.
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewRemovedTransaction instantiates a new RemovedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemovedTransaction() *RemovedTransaction {
	this := RemovedTransaction{}
	return &this
}

// NewRemovedTransactionWithDefaults instantiates a new RemovedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemovedTransactionWithDefaults() *RemovedTransaction {
	this := RemovedTransaction{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *RemovedTransaction) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemovedTransaction) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *RemovedTransaction) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *RemovedTransaction) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o RemovedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableRemovedTransaction struct {
	value *RemovedTransaction
	isSet bool
}

func (v NullableRemovedTransaction) Get() *RemovedTransaction {
	return v.value
}

func (v *NullableRemovedTransaction) Set(val *RemovedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableRemovedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableRemovedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemovedTransaction(val *RemovedTransaction) *NullableRemovedTransaction {
	return &NullableRemovedTransaction{value: val, isSet: true}
}

func (v NullableRemovedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemovedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


