/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.197.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WalletGetRequest WalletGetRequest defines the request schema for `/wallet/get`
type WalletGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the e-wallet
	WalletId string `json:"wallet_id"`
}

// NewWalletGetRequest instantiates a new WalletGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletGetRequest(walletId string) *WalletGetRequest {
	this := WalletGetRequest{}
	this.WalletId = walletId
	return &this
}

// NewWalletGetRequestWithDefaults instantiates a new WalletGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletGetRequestWithDefaults() *WalletGetRequest {
	this := WalletGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WalletGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WalletGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WalletGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WalletGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WalletGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WalletGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetWalletId returns the WalletId field value
func (o *WalletGetRequest) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletGetRequest) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletGetRequest) SetWalletId(v string) {
	o.WalletId = v
}

func (o WalletGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	return json.Marshal(toSerialize)
}

type NullableWalletGetRequest struct {
	value *WalletGetRequest
	isSet bool
}

func (v NullableWalletGetRequest) Get() *WalletGetRequest {
	return v.value
}

func (v *NullableWalletGetRequest) Set(val *WalletGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletGetRequest(val *WalletGetRequest) *NullableWalletGetRequest {
	return &NullableWalletGetRequest{value: val, isSet: true}
}

func (v NullableWalletGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


