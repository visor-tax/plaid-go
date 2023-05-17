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

// VerificationOfAssetResponse Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type VerificationOfAssetResponse struct {
	ASSETS Assets `json:"ASSETS"`
	AdditionalProperties map[string]interface{}
}

type _VerificationOfAssetResponse VerificationOfAssetResponse

// NewVerificationOfAssetResponse instantiates a new VerificationOfAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationOfAssetResponse(aSSETS Assets) *VerificationOfAssetResponse {
	this := VerificationOfAssetResponse{}
	this.ASSETS = aSSETS
	return &this
}

// NewVerificationOfAssetResponseWithDefaults instantiates a new VerificationOfAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationOfAssetResponseWithDefaults() *VerificationOfAssetResponse {
	this := VerificationOfAssetResponse{}
	return &this
}

// GetASSETS returns the ASSETS field value
func (o *VerificationOfAssetResponse) GetASSETS() Assets {
	if o == nil {
		var ret Assets
		return ret
	}

	return o.ASSETS
}

// GetASSETSOk returns a tuple with the ASSETS field value
// and a boolean to check if the value has been set.
func (o *VerificationOfAssetResponse) GetASSETSOk() (*Assets, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSETS, true
}

// SetASSETS sets field value
func (o *VerificationOfAssetResponse) SetASSETS(v Assets) {
	o.ASSETS = v
}

func (o VerificationOfAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSETS"] = o.ASSETS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerificationOfAssetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varVerificationOfAssetResponse := _VerificationOfAssetResponse{}

	if err = json.Unmarshal(bytes, &varVerificationOfAssetResponse); err == nil {
		*o = VerificationOfAssetResponse(varVerificationOfAssetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSETS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerificationOfAssetResponse struct {
	value *VerificationOfAssetResponse
	isSet bool
}

func (v NullableVerificationOfAssetResponse) Get() *VerificationOfAssetResponse {
	return v.value
}

func (v *NullableVerificationOfAssetResponse) Set(val *VerificationOfAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationOfAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationOfAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationOfAssetResponse(val *VerificationOfAssetResponse) *NullableVerificationOfAssetResponse {
	return &NullableVerificationOfAssetResponse{value: val, isSet: true}
}

func (v NullableVerificationOfAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationOfAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


