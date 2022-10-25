/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.197.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// UserIDNumber ID number submitted by the user, currently used only for the Identity Verification product. If the user has not submitted this data yet, this field will be `null`. Otherwise, both fields are guaranteed to be filled.
type UserIDNumber struct {
	// Value of identity document value typed in by user. Alpha-numeric, with all formatting characters stripped.
	Value string `json:"value"`
	Type IDNumberType `json:"type"`
}

// NewUserIDNumber instantiates a new UserIDNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIDNumber(value string, type_ IDNumberType) *UserIDNumber {
	this := UserIDNumber{}
	this.Value = value
	this.Type = type_
	return &this
}

// NewUserIDNumberWithDefaults instantiates a new UserIDNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIDNumberWithDefaults() *UserIDNumber {
	this := UserIDNumber{}
	return &this
}

// GetValue returns the Value field value
func (o *UserIDNumber) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UserIDNumber) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *UserIDNumber) SetValue(v string) {
	o.Value = v
}

// GetType returns the Type field value
func (o *UserIDNumber) GetType() IDNumberType {
	if o == nil {
		var ret IDNumberType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserIDNumber) GetTypeOk() (*IDNumberType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserIDNumber) SetType(v IDNumberType) {
	o.Type = v
}

func (o UserIDNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUserIDNumber struct {
	value *UserIDNumber
	isSet bool
}

func (v NullableUserIDNumber) Get() *UserIDNumber {
	return v.value
}

func (v *NullableUserIDNumber) Set(val *UserIDNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIDNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIDNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIDNumber(val *UserIDNumber) *NullableUserIDNumber {
	return &NullableUserIDNumber{value: val, isSet: true}
}

func (v NullableUserIDNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIDNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


