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

// TransferCapabilitiesGetResponse Defines the response schema for `/transfer/capabilities/get`
type TransferCapabilitiesGetResponse struct {
	InstitutionSupportedNetworks InstitutionSupportedNetworks `json:"institution_supported_networks"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferCapabilitiesGetResponse TransferCapabilitiesGetResponse

// NewTransferCapabilitiesGetResponse instantiates a new TransferCapabilitiesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCapabilitiesGetResponse(institutionSupportedNetworks InstitutionSupportedNetworks, requestId string) *TransferCapabilitiesGetResponse {
	this := TransferCapabilitiesGetResponse{}
	this.InstitutionSupportedNetworks = institutionSupportedNetworks
	this.RequestId = requestId
	return &this
}

// NewTransferCapabilitiesGetResponseWithDefaults instantiates a new TransferCapabilitiesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCapabilitiesGetResponseWithDefaults() *TransferCapabilitiesGetResponse {
	this := TransferCapabilitiesGetResponse{}
	return &this
}

// GetInstitutionSupportedNetworks returns the InstitutionSupportedNetworks field value
func (o *TransferCapabilitiesGetResponse) GetInstitutionSupportedNetworks() InstitutionSupportedNetworks {
	if o == nil {
		var ret InstitutionSupportedNetworks
		return ret
	}

	return o.InstitutionSupportedNetworks
}

// GetInstitutionSupportedNetworksOk returns a tuple with the InstitutionSupportedNetworks field value
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetResponse) GetInstitutionSupportedNetworksOk() (*InstitutionSupportedNetworks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionSupportedNetworks, true
}

// SetInstitutionSupportedNetworks sets field value
func (o *TransferCapabilitiesGetResponse) SetInstitutionSupportedNetworks(v InstitutionSupportedNetworks) {
	o.InstitutionSupportedNetworks = v
}

// GetRequestId returns the RequestId field value
func (o *TransferCapabilitiesGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferCapabilitiesGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferCapabilitiesGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferCapabilitiesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institution_supported_networks"] = o.InstitutionSupportedNetworks
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferCapabilitiesGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferCapabilitiesGetResponse := _TransferCapabilitiesGetResponse{}

	if err = json.Unmarshal(bytes, &varTransferCapabilitiesGetResponse); err == nil {
		*o = TransferCapabilitiesGetResponse(varTransferCapabilitiesGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "institution_supported_networks")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferCapabilitiesGetResponse struct {
	value *TransferCapabilitiesGetResponse
	isSet bool
}

func (v NullableTransferCapabilitiesGetResponse) Get() *TransferCapabilitiesGetResponse {
	return v.value
}

func (v *NullableTransferCapabilitiesGetResponse) Set(val *TransferCapabilitiesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCapabilitiesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCapabilitiesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCapabilitiesGetResponse(val *TransferCapabilitiesGetResponse) *NullableTransferCapabilitiesGetResponse {
	return &NullableTransferCapabilitiesGetResponse{value: val, isSet: true}
}

func (v NullableTransferCapabilitiesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCapabilitiesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


