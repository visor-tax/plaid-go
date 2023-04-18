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

// TransactionsEnhanceGetRequest TransactionsEnhanceGetRequest defines the request schema for `/transactions/enhance`.
type TransactionsEnhanceGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The type of account for the requested transactions (`depository` or `credit`).
	AccountType string `json:"account_type"`
	// An array of raw transactions to be enhanced.
	Transactions []ClientProvidedRawTransaction `json:"transactions"`
}

// NewTransactionsEnhanceGetRequest instantiates a new TransactionsEnhanceGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsEnhanceGetRequest(accountType string, transactions []ClientProvidedRawTransaction) *TransactionsEnhanceGetRequest {
	this := TransactionsEnhanceGetRequest{}
	this.AccountType = accountType
	this.Transactions = transactions
	return &this
}

// NewTransactionsEnhanceGetRequestWithDefaults instantiates a new TransactionsEnhanceGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsEnhanceGetRequestWithDefaults() *TransactionsEnhanceGetRequest {
	this := TransactionsEnhanceGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsEnhanceGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnhanceGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsEnhanceGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsEnhanceGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsEnhanceGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsEnhanceGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsEnhanceGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsEnhanceGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccountType returns the AccountType field value
func (o *TransactionsEnhanceGetRequest) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *TransactionsEnhanceGetRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *TransactionsEnhanceGetRequest) SetAccountType(v string) {
	o.AccountType = v
}

// GetTransactions returns the Transactions field value
func (o *TransactionsEnhanceGetRequest) GetTransactions() []ClientProvidedRawTransaction {
	if o == nil {
		var ret []ClientProvidedRawTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsEnhanceGetRequest) GetTransactionsOk() (*[]ClientProvidedRawTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *TransactionsEnhanceGetRequest) SetTransactions(v []ClientProvidedRawTransaction) {
	o.Transactions = v
}

func (o TransactionsEnhanceGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsEnhanceGetRequest struct {
	value *TransactionsEnhanceGetRequest
	isSet bool
}

func (v NullableTransactionsEnhanceGetRequest) Get() *TransactionsEnhanceGetRequest {
	return v.value
}

func (v *NullableTransactionsEnhanceGetRequest) Set(val *TransactionsEnhanceGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsEnhanceGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsEnhanceGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsEnhanceGetRequest(val *TransactionsEnhanceGetRequest) *NullableTransactionsEnhanceGetRequest {
	return &NullableTransactionsEnhanceGetRequest{value: val, isSet: true}
}

func (v NullableTransactionsEnhanceGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsEnhanceGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


