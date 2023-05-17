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

// PaymentInitiationRecipientGetRequest PaymentInitiationRecipientGetRequest defines the request schema for `/payment_initiation/recipient/get`
type PaymentInitiationRecipientGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the recipient
	RecipientId string `json:"recipient_id"`
}

// NewPaymentInitiationRecipientGetRequest instantiates a new PaymentInitiationRecipientGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientGetRequest(recipientId string) *PaymentInitiationRecipientGetRequest {
	this := PaymentInitiationRecipientGetRequest{}
	this.RecipientId = recipientId
	return &this
}

// NewPaymentInitiationRecipientGetRequestWithDefaults instantiates a new PaymentInitiationRecipientGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientGetRequestWithDefaults() *PaymentInitiationRecipientGetRequest {
	this := PaymentInitiationRecipientGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentInitiationRecipientGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentInitiationRecipientGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentInitiationRecipientGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentInitiationRecipientGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationRecipientGetRequest) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientGetRequest) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationRecipientGetRequest) SetRecipientId(v string) {
	o.RecipientId = v
}

func (o PaymentInitiationRecipientGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentInitiationRecipientGetRequest struct {
	value *PaymentInitiationRecipientGetRequest
	isSet bool
}

func (v NullablePaymentInitiationRecipientGetRequest) Get() *PaymentInitiationRecipientGetRequest {
	return v.value
}

func (v *NullablePaymentInitiationRecipientGetRequest) Set(val *PaymentInitiationRecipientGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientGetRequest(val *PaymentInitiationRecipientGetRequest) *NullablePaymentInitiationRecipientGetRequest {
	return &NullablePaymentInitiationRecipientGetRequest{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


