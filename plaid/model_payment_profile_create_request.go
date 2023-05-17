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

// PaymentProfileCreateRequest PaymentProfileCreateRequest defines the request schema for `/payment_profile/create`
type PaymentProfileCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewPaymentProfileCreateRequest instantiates a new PaymentProfileCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentProfileCreateRequest() *PaymentProfileCreateRequest {
	this := PaymentProfileCreateRequest{}
	return &this
}

// NewPaymentProfileCreateRequestWithDefaults instantiates a new PaymentProfileCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentProfileCreateRequestWithDefaults() *PaymentProfileCreateRequest {
	this := PaymentProfileCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *PaymentProfileCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentProfileCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *PaymentProfileCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *PaymentProfileCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PaymentProfileCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentProfileCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PaymentProfileCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *PaymentProfileCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o PaymentProfileCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentProfileCreateRequest struct {
	value *PaymentProfileCreateRequest
	isSet bool
}

func (v NullablePaymentProfileCreateRequest) Get() *PaymentProfileCreateRequest {
	return v.value
}

func (v *NullablePaymentProfileCreateRequest) Set(val *PaymentProfileCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentProfileCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentProfileCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentProfileCreateRequest(val *PaymentProfileCreateRequest) *NullablePaymentProfileCreateRequest {
	return &NullablePaymentProfileCreateRequest{value: val, isSet: true}
}

func (v NullablePaymentProfileCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentProfileCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


