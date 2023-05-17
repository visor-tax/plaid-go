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

// PaymentInitiationRecipientListRequest PaymentInitiationRecipientListRequest defines the request schema for `/payment_initiation/recipient/list`
type PaymentInitiationRecipientListRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewPaymentInitiationRecipientListRequest instantiates a new PaymentInitiationRecipientListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientListRequest() *PaymentInitiationRecipientListRequest {
	this := PaymentInitiationRecipientListRequest{}
	return &this
}

// NewPaymentInitiationRecipientListRequestWithDefaults instantiates a new PaymentInitiationRecipientListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientListRequestWithDefaults() *PaymentInitiationRecipientListRequest {
	this := PaymentInitiationRecipientListRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientListRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientListRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientListRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationRecipientListRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientListRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientListRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientListRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationRecipientListRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o PaymentInitiationRecipientListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationRecipientListRequest struct {
	value *PaymentInitiationRecipientListRequest
	isSet bool
}

func (v NullablePaymentInitiationRecipientListRequest) Get() *PaymentInitiationRecipientListRequest {
	return v.value
}

func (v *NullablePaymentInitiationRecipientListRequest) Set(val *PaymentInitiationRecipientListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientListRequest(val *PaymentInitiationRecipientListRequest) *NullablePaymentInitiationRecipientListRequest {
	return &NullablePaymentInitiationRecipientListRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


