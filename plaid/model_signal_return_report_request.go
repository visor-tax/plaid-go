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
	"time"
)

// SignalReturnReportRequest SignalReturnReportRequest defines the request schema for `/signal/return/report`
type SignalReturnReportRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Must be the same as the `client_transaction_id` supplied when calling `/signal/evaluate`
	ClientTransactionId string `json:"client_transaction_id"`
	// Must be a valid ACH return code (e.g. \"R01\")  If formatted incorrectly, this will result in an [`INVALID_FIELD`](/docs/errors/invalid-request/#invalid_field) error.
	ReturnCode string `json:"return_code"`
	// Date and time when you receive the returns from your payment processors, in ISO 8601 format (`YYYY-MM-DDTHH:mm:ssZ`).
	ReturnedAt NullableTime `json:"returned_at,omitempty"`
}

// NewSignalReturnReportRequest instantiates a new SignalReturnReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignalReturnReportRequest(clientTransactionId string, returnCode string) *SignalReturnReportRequest {
	this := SignalReturnReportRequest{}
	this.ClientTransactionId = clientTransactionId
	this.ReturnCode = returnCode
	return &this
}

// NewSignalReturnReportRequestWithDefaults instantiates a new SignalReturnReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignalReturnReportRequestWithDefaults() *SignalReturnReportRequest {
	this := SignalReturnReportRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SignalReturnReportRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalReturnReportRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SignalReturnReportRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SignalReturnReportRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SignalReturnReportRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignalReturnReportRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SignalReturnReportRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SignalReturnReportRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientTransactionId returns the ClientTransactionId field value
func (o *SignalReturnReportRequest) GetClientTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientTransactionId
}

// GetClientTransactionIdOk returns a tuple with the ClientTransactionId field value
// and a boolean to check if the value has been set.
func (o *SignalReturnReportRequest) GetClientTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientTransactionId, true
}

// SetClientTransactionId sets field value
func (o *SignalReturnReportRequest) SetClientTransactionId(v string) {
	o.ClientTransactionId = v
}

// GetReturnCode returns the ReturnCode field value
func (o *SignalReturnReportRequest) GetReturnCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value
// and a boolean to check if the value has been set.
func (o *SignalReturnReportRequest) GetReturnCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReturnCode, true
}

// SetReturnCode sets field value
func (o *SignalReturnReportRequest) SetReturnCode(v string) {
	o.ReturnCode = v
}

// GetReturnedAt returns the ReturnedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignalReturnReportRequest) GetReturnedAt() time.Time {
	if o == nil || o.ReturnedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ReturnedAt.Get()
}

// GetReturnedAtOk returns a tuple with the ReturnedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignalReturnReportRequest) GetReturnedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReturnedAt.Get(), o.ReturnedAt.IsSet()
}

// HasReturnedAt returns a boolean if a field has been set.
func (o *SignalReturnReportRequest) HasReturnedAt() bool {
	if o != nil && o.ReturnedAt.IsSet() {
		return true
	}

	return false
}

// SetReturnedAt gets a reference to the given NullableTime and assigns it to the ReturnedAt field.
func (o *SignalReturnReportRequest) SetReturnedAt(v time.Time) {
	o.ReturnedAt.Set(&v)
}
// SetReturnedAtNil sets the value for ReturnedAt to be an explicit nil
func (o *SignalReturnReportRequest) SetReturnedAtNil() {
	o.ReturnedAt.Set(nil)
}

// UnsetReturnedAt ensures that no value is present for ReturnedAt, not even an explicit nil
func (o *SignalReturnReportRequest) UnsetReturnedAt() {
	o.ReturnedAt.Unset()
}

func (o SignalReturnReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["client_transaction_id"] = o.ClientTransactionId
	}
	if true {
		toSerialize["return_code"] = o.ReturnCode
	}
	if o.ReturnedAt.IsSet() {
		toSerialize["returned_at"] = o.ReturnedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSignalReturnReportRequest struct {
	value *SignalReturnReportRequest
	isSet bool
}

func (v NullableSignalReturnReportRequest) Get() *SignalReturnReportRequest {
	return v.value
}

func (v *NullableSignalReturnReportRequest) Set(val *SignalReturnReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSignalReturnReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSignalReturnReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignalReturnReportRequest(val *SignalReturnReportRequest) *NullableSignalReturnReportRequest {
	return &NullableSignalReturnReportRequest{value: val, isSet: true}
}

func (v NullableSignalReturnReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignalReturnReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


