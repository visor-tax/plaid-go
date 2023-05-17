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

// CreditFreddieMacAssetsVOE25 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacAssetsVOE25 struct {
	// Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
	ASSET []CreditFreddieMacAssetVOE25 `json:"ASSET"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacAssetsVOE25 CreditFreddieMacAssetsVOE25

// NewCreditFreddieMacAssetsVOE25 instantiates a new CreditFreddieMacAssetsVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacAssetsVOE25(aSSET []CreditFreddieMacAssetVOE25) *CreditFreddieMacAssetsVOE25 {
	this := CreditFreddieMacAssetsVOE25{}
	this.ASSET = aSSET
	return &this
}

// NewCreditFreddieMacAssetsVOE25WithDefaults instantiates a new CreditFreddieMacAssetsVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacAssetsVOE25WithDefaults() *CreditFreddieMacAssetsVOE25 {
	this := CreditFreddieMacAssetsVOE25{}
	return &this
}

// GetASSET returns the ASSET field value
func (o *CreditFreddieMacAssetsVOE25) GetASSET() []CreditFreddieMacAssetVOE25 {
	if o == nil {
		var ret []CreditFreddieMacAssetVOE25
		return ret
	}

	return o.ASSET
}

// GetASSETOk returns a tuple with the ASSET field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacAssetsVOE25) GetASSETOk() (*[]CreditFreddieMacAssetVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASSET, true
}

// SetASSET sets field value
func (o *CreditFreddieMacAssetsVOE25) SetASSET(v []CreditFreddieMacAssetVOE25) {
	o.ASSET = v
}

func (o CreditFreddieMacAssetsVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ASSET"] = o.ASSET
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacAssetsVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacAssetsVOE25 := _CreditFreddieMacAssetsVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacAssetsVOE25); err == nil {
		*o = CreditFreddieMacAssetsVOE25(varCreditFreddieMacAssetsVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ASSET")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacAssetsVOE25 struct {
	value *CreditFreddieMacAssetsVOE25
	isSet bool
}

func (v NullableCreditFreddieMacAssetsVOE25) Get() *CreditFreddieMacAssetsVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacAssetsVOE25) Set(val *CreditFreddieMacAssetsVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacAssetsVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacAssetsVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacAssetsVOE25(val *CreditFreddieMacAssetsVOE25) *NullableCreditFreddieMacAssetsVOE25 {
	return &NullableCreditFreddieMacAssetsVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacAssetsVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacAssetsVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


