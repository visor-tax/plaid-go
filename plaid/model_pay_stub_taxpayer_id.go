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

// PayStubTaxpayerID Taxpayer ID of the individual receiving the paystub.
type PayStubTaxpayerID struct {
	// Type of ID, e.g. 'SSN'.
	IdType NullableString `json:"id_type"`
	// ID mask; i.e. last 4 digits of the taxpayer ID.
	IdMask NullableString `json:"id_mask"`
	AdditionalProperties map[string]interface{}
}

type _PayStubTaxpayerID PayStubTaxpayerID

// NewPayStubTaxpayerID instantiates a new PayStubTaxpayerID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStubTaxpayerID(idType NullableString, idMask NullableString) *PayStubTaxpayerID {
	this := PayStubTaxpayerID{}
	this.IdType = idType
	this.IdMask = idMask
	return &this
}

// NewPayStubTaxpayerIDWithDefaults instantiates a new PayStubTaxpayerID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubTaxpayerIDWithDefaults() *PayStubTaxpayerID {
	this := PayStubTaxpayerID{}
	return &this
}

// GetIdType returns the IdType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubTaxpayerID) GetIdType() string {
	if o == nil || o.IdType.Get() == nil {
		var ret string
		return ret
	}

	return *o.IdType.Get()
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubTaxpayerID) GetIdTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdType.Get(), o.IdType.IsSet()
}

// SetIdType sets field value
func (o *PayStubTaxpayerID) SetIdType(v string) {
	o.IdType.Set(&v)
}

// GetIdMask returns the IdMask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubTaxpayerID) GetIdMask() string {
	if o == nil || o.IdMask.Get() == nil {
		var ret string
		return ret
	}

	return *o.IdMask.Get()
}

// GetIdMaskOk returns a tuple with the IdMask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubTaxpayerID) GetIdMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdMask.Get(), o.IdMask.IsSet()
}

// SetIdMask sets field value
func (o *PayStubTaxpayerID) SetIdMask(v string) {
	o.IdMask.Set(&v)
}

func (o PayStubTaxpayerID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id_type"] = o.IdType.Get()
	}
	if true {
		toSerialize["id_mask"] = o.IdMask.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStubTaxpayerID) UnmarshalJSON(bytes []byte) (err error) {
	varPayStubTaxpayerID := _PayStubTaxpayerID{}

	if err = json.Unmarshal(bytes, &varPayStubTaxpayerID); err == nil {
		*o = PayStubTaxpayerID(varPayStubTaxpayerID)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id_type")
		delete(additionalProperties, "id_mask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStubTaxpayerID struct {
	value *PayStubTaxpayerID
	isSet bool
}

func (v NullablePayStubTaxpayerID) Get() *PayStubTaxpayerID {
	return v.value
}

func (v *NullablePayStubTaxpayerID) Set(val *PayStubTaxpayerID) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStubTaxpayerID) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStubTaxpayerID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStubTaxpayerID(val *PayStubTaxpayerID) *NullablePayStubTaxpayerID {
	return &NullablePayStubTaxpayerID{value: val, isSet: true}
}

func (v NullablePayStubTaxpayerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStubTaxpayerID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


