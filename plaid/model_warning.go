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

// Warning It is possible for an Asset Report to be returned with missing account owner information. In such cases, the Asset Report will contain warning data in the response, indicating why obtaining the owner information failed.
type Warning struct {
	// The warning type, which will always be `ASSET_REPORT_WARNING`
	WarningType string `json:"warning_type"`
	// The warning code identifies a specific kind of warning. Currently, the only possible warning code is `OWNERS_UNAVAILABLE`, which indicates that account-owner information is not available.
	WarningCode string `json:"warning_code"`
	Cause Cause `json:"cause"`
	AdditionalProperties map[string]interface{}
}

type _Warning Warning

// NewWarning instantiates a new Warning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarning(warningType string, warningCode string, cause Cause) *Warning {
	this := Warning{}
	this.WarningType = warningType
	this.WarningCode = warningCode
	this.Cause = cause
	return &this
}

// NewWarningWithDefaults instantiates a new Warning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningWithDefaults() *Warning {
	this := Warning{}
	return &this
}

// GetWarningType returns the WarningType field value
func (o *Warning) GetWarningType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarningType
}

// GetWarningTypeOk returns a tuple with the WarningType field value
// and a boolean to check if the value has been set.
func (o *Warning) GetWarningTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WarningType, true
}

// SetWarningType sets field value
func (o *Warning) SetWarningType(v string) {
	o.WarningType = v
}

// GetWarningCode returns the WarningCode field value
func (o *Warning) GetWarningCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WarningCode
}

// GetWarningCodeOk returns a tuple with the WarningCode field value
// and a boolean to check if the value has been set.
func (o *Warning) GetWarningCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WarningCode, true
}

// SetWarningCode sets field value
func (o *Warning) SetWarningCode(v string) {
	o.WarningCode = v
}

// GetCause returns the Cause field value
func (o *Warning) GetCause() Cause {
	if o == nil {
		var ret Cause
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *Warning) GetCauseOk() (*Cause, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *Warning) SetCause(v Cause) {
	o.Cause = v
}

func (o Warning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warning_type"] = o.WarningType
	}
	if true {
		toSerialize["warning_code"] = o.WarningCode
	}
	if true {
		toSerialize["cause"] = o.Cause
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Warning) UnmarshalJSON(bytes []byte) (err error) {
	varWarning := _Warning{}

	if err = json.Unmarshal(bytes, &varWarning); err == nil {
		*o = Warning(varWarning)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "warning_type")
		delete(additionalProperties, "warning_code")
		delete(additionalProperties, "cause")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWarning struct {
	value *Warning
	isSet bool
}

func (v NullableWarning) Get() *Warning {
	return v.value
}

func (v *NullableWarning) Set(val *Warning) {
	v.value = val
	v.isSet = true
}

func (v NullableWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarning(val *Warning) *NullableWarning {
	return &NullableWarning{value: val, isSet: true}
}

func (v NullableWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


