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

// CreditFreddieMacServiceVOE25 A collection of details related to a fulfillment service or product in terms of request, process and result.
type CreditFreddieMacServiceVOE25 struct {
	VERIFICATION_OF_ASSET []CreditFreddieMacVerificationOfAssetVOE25 `json:"VERIFICATION_OF_ASSET"`
	STATUSES Statuses `json:"STATUSES"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacServiceVOE25 CreditFreddieMacServiceVOE25

// NewCreditFreddieMacServiceVOE25 instantiates a new CreditFreddieMacServiceVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacServiceVOE25(vERIFICATIONOFASSET []CreditFreddieMacVerificationOfAssetVOE25, sTATUSES Statuses) *CreditFreddieMacServiceVOE25 {
	this := CreditFreddieMacServiceVOE25{}
	this.VERIFICATION_OF_ASSET = vERIFICATIONOFASSET
	this.STATUSES = sTATUSES
	return &this
}

// NewCreditFreddieMacServiceVOE25WithDefaults instantiates a new CreditFreddieMacServiceVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacServiceVOE25WithDefaults() *CreditFreddieMacServiceVOE25 {
	this := CreditFreddieMacServiceVOE25{}
	return &this
}

// GetVERIFICATION_OF_ASSET returns the VERIFICATION_OF_ASSET field value
func (o *CreditFreddieMacServiceVOE25) GetVERIFICATION_OF_ASSET() []CreditFreddieMacVerificationOfAssetVOE25 {
	if o == nil {
		var ret []CreditFreddieMacVerificationOfAssetVOE25
		return ret
	}

	return o.VERIFICATION_OF_ASSET
}

// GetVERIFICATION_OF_ASSETOk returns a tuple with the VERIFICATION_OF_ASSET field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacServiceVOE25) GetVERIFICATION_OF_ASSETOk() (*[]CreditFreddieMacVerificationOfAssetVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VERIFICATION_OF_ASSET, true
}

// SetVERIFICATION_OF_ASSET sets field value
func (o *CreditFreddieMacServiceVOE25) SetVERIFICATION_OF_ASSET(v []CreditFreddieMacVerificationOfAssetVOE25) {
	o.VERIFICATION_OF_ASSET = v
}

// GetSTATUSES returns the STATUSES field value
func (o *CreditFreddieMacServiceVOE25) GetSTATUSES() Statuses {
	if o == nil {
		var ret Statuses
		return ret
	}

	return o.STATUSES
}

// GetSTATUSESOk returns a tuple with the STATUSES field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacServiceVOE25) GetSTATUSESOk() (*Statuses, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.STATUSES, true
}

// SetSTATUSES sets field value
func (o *CreditFreddieMacServiceVOE25) SetSTATUSES(v Statuses) {
	o.STATUSES = v
}

func (o CreditFreddieMacServiceVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["VERIFICATION_OF_ASSET"] = o.VERIFICATION_OF_ASSET
	}
	if true {
		toSerialize["STATUSES"] = o.STATUSES
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacServiceVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacServiceVOE25 := _CreditFreddieMacServiceVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacServiceVOE25); err == nil {
		*o = CreditFreddieMacServiceVOE25(varCreditFreddieMacServiceVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "VERIFICATION_OF_ASSET")
		delete(additionalProperties, "STATUSES")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacServiceVOE25 struct {
	value *CreditFreddieMacServiceVOE25
	isSet bool
}

func (v NullableCreditFreddieMacServiceVOE25) Get() *CreditFreddieMacServiceVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacServiceVOE25) Set(val *CreditFreddieMacServiceVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacServiceVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacServiceVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacServiceVOE25(val *CreditFreddieMacServiceVOE25) *NullableCreditFreddieMacServiceVOE25 {
	return &NullableCreditFreddieMacServiceVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacServiceVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacServiceVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


