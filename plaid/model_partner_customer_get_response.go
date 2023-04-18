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

// PartnerCustomerGetResponse Response schema for `/partner/customer/get`.
type PartnerCustomerGetResponse struct {
	EndCustomer *PartnerEndCustomer `json:"end_customer,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerCustomerGetResponse PartnerCustomerGetResponse

// NewPartnerCustomerGetResponse instantiates a new PartnerCustomerGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCustomerGetResponse() *PartnerCustomerGetResponse {
	this := PartnerCustomerGetResponse{}
	return &this
}

// NewPartnerCustomerGetResponseWithDefaults instantiates a new PartnerCustomerGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCustomerGetResponseWithDefaults() *PartnerCustomerGetResponse {
	this := PartnerCustomerGetResponse{}
	return &this
}

// GetEndCustomer returns the EndCustomer field value if set, zero value otherwise.
func (o *PartnerCustomerGetResponse) GetEndCustomer() PartnerEndCustomer {
	if o == nil || o.EndCustomer == nil {
		var ret PartnerEndCustomer
		return ret
	}
	return *o.EndCustomer
}

// GetEndCustomerOk returns a tuple with the EndCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerGetResponse) GetEndCustomerOk() (*PartnerEndCustomer, bool) {
	if o == nil || o.EndCustomer == nil {
		return nil, false
	}
	return o.EndCustomer, true
}

// HasEndCustomer returns a boolean if a field has been set.
func (o *PartnerCustomerGetResponse) HasEndCustomer() bool {
	if o != nil && o.EndCustomer != nil {
		return true
	}

	return false
}

// SetEndCustomer gets a reference to the given PartnerEndCustomer and assigns it to the EndCustomer field.
func (o *PartnerCustomerGetResponse) SetEndCustomer(v PartnerEndCustomer) {
	o.EndCustomer = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PartnerCustomerGetResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PartnerCustomerGetResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PartnerCustomerGetResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PartnerCustomerGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndCustomer != nil {
		toSerialize["end_customer"] = o.EndCustomer
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerCustomerGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerCustomerGetResponse := _PartnerCustomerGetResponse{}

	if err = json.Unmarshal(bytes, &varPartnerCustomerGetResponse); err == nil {
		*o = PartnerCustomerGetResponse(varPartnerCustomerGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "end_customer")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerCustomerGetResponse struct {
	value *PartnerCustomerGetResponse
	isSet bool
}

func (v NullablePartnerCustomerGetResponse) Get() *PartnerCustomerGetResponse {
	return v.value
}

func (v *NullablePartnerCustomerGetResponse) Set(val *PartnerCustomerGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCustomerGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCustomerGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCustomerGetResponse(val *PartnerCustomerGetResponse) *NullablePartnerCustomerGetResponse {
	return &NullablePartnerCustomerGetResponse{value: val, isSet: true}
}

func (v NullablePartnerCustomerGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCustomerGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


