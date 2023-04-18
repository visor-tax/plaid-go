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

// PartnerCustomerEnableResponse Response schema for `/partner/customer/enable`.
type PartnerCustomerEnableResponse struct {
	// The end customer's secret key for the Production environment.
	ProductionSecret *string `json:"production_secret,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId *string `json:"request_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerCustomerEnableResponse PartnerCustomerEnableResponse

// NewPartnerCustomerEnableResponse instantiates a new PartnerCustomerEnableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerCustomerEnableResponse() *PartnerCustomerEnableResponse {
	this := PartnerCustomerEnableResponse{}
	return &this
}

// NewPartnerCustomerEnableResponseWithDefaults instantiates a new PartnerCustomerEnableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerCustomerEnableResponseWithDefaults() *PartnerCustomerEnableResponse {
	this := PartnerCustomerEnableResponse{}
	return &this
}

// GetProductionSecret returns the ProductionSecret field value if set, zero value otherwise.
func (o *PartnerCustomerEnableResponse) GetProductionSecret() string {
	if o == nil || o.ProductionSecret == nil {
		var ret string
		return ret
	}
	return *o.ProductionSecret
}

// GetProductionSecretOk returns a tuple with the ProductionSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerEnableResponse) GetProductionSecretOk() (*string, bool) {
	if o == nil || o.ProductionSecret == nil {
		return nil, false
	}
	return o.ProductionSecret, true
}

// HasProductionSecret returns a boolean if a field has been set.
func (o *PartnerCustomerEnableResponse) HasProductionSecret() bool {
	if o != nil && o.ProductionSecret != nil {
		return true
	}

	return false
}

// SetProductionSecret gets a reference to the given string and assigns it to the ProductionSecret field.
func (o *PartnerCustomerEnableResponse) SetProductionSecret(v string) {
	o.ProductionSecret = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *PartnerCustomerEnableResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerCustomerEnableResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *PartnerCustomerEnableResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *PartnerCustomerEnableResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o PartnerCustomerEnableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProductionSecret != nil {
		toSerialize["production_secret"] = o.ProductionSecret
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerCustomerEnableResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerCustomerEnableResponse := _PartnerCustomerEnableResponse{}

	if err = json.Unmarshal(bytes, &varPartnerCustomerEnableResponse); err == nil {
		*o = PartnerCustomerEnableResponse(varPartnerCustomerEnableResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "production_secret")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerCustomerEnableResponse struct {
	value *PartnerCustomerEnableResponse
	isSet bool
}

func (v NullablePartnerCustomerEnableResponse) Get() *PartnerCustomerEnableResponse {
	return v.value
}

func (v *NullablePartnerCustomerEnableResponse) Set(val *PartnerCustomerEnableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerCustomerEnableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerCustomerEnableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerCustomerEnableResponse(val *PartnerCustomerEnableResponse) *NullablePartnerCustomerEnableResponse {
	return &NullablePartnerCustomerEnableResponse{value: val, isSet: true}
}

func (v NullablePartnerCustomerEnableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerCustomerEnableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


