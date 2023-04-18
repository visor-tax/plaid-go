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

// PaymentInitiationConsentRevokeResponse PaymentInitiationConsentRevokeResponse defines the response schema for `/payment_initation/consent/revoke`
type PaymentInitiationConsentRevokeResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationConsentRevokeResponse PaymentInitiationConsentRevokeResponse

// NewPaymentInitiationConsentRevokeResponse instantiates a new PaymentInitiationConsentRevokeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationConsentRevokeResponse() *PaymentInitiationConsentRevokeResponse {
	this := PaymentInitiationConsentRevokeResponse{}
	return &this
}

// NewPaymentInitiationConsentRevokeResponseWithDefaults instantiates a new PaymentInitiationConsentRevokeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationConsentRevokeResponseWithDefaults() *PaymentInitiationConsentRevokeResponse {
	this := PaymentInitiationConsentRevokeResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PaymentInitiationConsentRevokeResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationConsentRevokeResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PaymentInitiationConsentRevokeResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PaymentInitiationConsentRevokeResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PaymentInitiationConsentRevokeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationConsentRevokeResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationConsentRevokeResponse := _PaymentInitiationConsentRevokeResponse{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationConsentRevokeResponse); err == nil {
		*o = PaymentInitiationConsentRevokeResponse(varPaymentInitiationConsentRevokeResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationConsentRevokeResponse struct {
	value *PaymentInitiationConsentRevokeResponse
	isSet bool
}

func (v NullablePaymentInitiationConsentRevokeResponse) Get() *PaymentInitiationConsentRevokeResponse {
	return v.value
}

func (v *NullablePaymentInitiationConsentRevokeResponse) Set(val *PaymentInitiationConsentRevokeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationConsentRevokeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationConsentRevokeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationConsentRevokeResponse(val *PaymentInitiationConsentRevokeResponse) *NullablePaymentInitiationConsentRevokeResponse {
	return &NullablePaymentInitiationConsentRevokeResponse{value: val, isSet: true}
}

func (v NullablePaymentInitiationConsentRevokeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationConsentRevokeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


