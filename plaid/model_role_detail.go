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

// RoleDetail Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type RoleDetail struct {
	PartyRoleType PartyRoleType `json:"PartyRoleType"`
	AdditionalProperties map[string]interface{}
}

type _RoleDetail RoleDetail

// NewRoleDetail instantiates a new RoleDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleDetail(partyRoleType PartyRoleType) *RoleDetail {
	this := RoleDetail{}
	this.PartyRoleType = partyRoleType
	return &this
}

// NewRoleDetailWithDefaults instantiates a new RoleDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleDetailWithDefaults() *RoleDetail {
	this := RoleDetail{}
	return &this
}

// GetPartyRoleType returns the PartyRoleType field value
func (o *RoleDetail) GetPartyRoleType() PartyRoleType {
	if o == nil {
		var ret PartyRoleType
		return ret
	}

	return o.PartyRoleType
}

// GetPartyRoleTypeOk returns a tuple with the PartyRoleType field value
// and a boolean to check if the value has been set.
func (o *RoleDetail) GetPartyRoleTypeOk() (*PartyRoleType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartyRoleType, true
}

// SetPartyRoleType sets field value
func (o *RoleDetail) SetPartyRoleType(v PartyRoleType) {
	o.PartyRoleType = v
}

func (o RoleDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["PartyRoleType"] = o.PartyRoleType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleDetail) UnmarshalJSON(bytes []byte) (err error) {
	varRoleDetail := _RoleDetail{}

	if err = json.Unmarshal(bytes, &varRoleDetail); err == nil {
		*o = RoleDetail(varRoleDetail)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "PartyRoleType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoleDetail struct {
	value *RoleDetail
	isSet bool
}

func (v NullableRoleDetail) Get() *RoleDetail {
	return v.value
}

func (v *NullableRoleDetail) Set(val *RoleDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleDetail(val *RoleDetail) *NullableRoleDetail {
	return &NullableRoleDetail{value: val, isSet: true}
}

func (v NullableRoleDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


