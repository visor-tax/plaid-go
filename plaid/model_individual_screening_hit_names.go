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

// IndividualScreeningHitNames Name information for the associated individual watchlist hit
type IndividualScreeningHitNames struct {
	// The full name of the individual, including all parts.
	Full string `json:"full"`
	// Primary names are those most commonly used to refer to this person. Only one name will ever be marked as primary.
	IsPrimary bool `json:"is_primary"`
	WeakAliasDetermination WeakAliasDetermination `json:"weak_alias_determination"`
}

// NewIndividualScreeningHitNames instantiates a new IndividualScreeningHitNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualScreeningHitNames(full string, isPrimary bool, weakAliasDetermination WeakAliasDetermination) *IndividualScreeningHitNames {
	this := IndividualScreeningHitNames{}
	this.Full = full
	this.IsPrimary = isPrimary
	this.WeakAliasDetermination = weakAliasDetermination
	return &this
}

// NewIndividualScreeningHitNamesWithDefaults instantiates a new IndividualScreeningHitNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualScreeningHitNamesWithDefaults() *IndividualScreeningHitNames {
	this := IndividualScreeningHitNames{}
	return &this
}

// GetFull returns the Full field value
func (o *IndividualScreeningHitNames) GetFull() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Full
}

// GetFullOk returns a tuple with the Full field value
// and a boolean to check if the value has been set.
func (o *IndividualScreeningHitNames) GetFullOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Full, true
}

// SetFull sets field value
func (o *IndividualScreeningHitNames) SetFull(v string) {
	o.Full = v
}

// GetIsPrimary returns the IsPrimary field value
func (o *IndividualScreeningHitNames) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *IndividualScreeningHitNames) GetIsPrimaryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *IndividualScreeningHitNames) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetWeakAliasDetermination returns the WeakAliasDetermination field value
func (o *IndividualScreeningHitNames) GetWeakAliasDetermination() WeakAliasDetermination {
	if o == nil {
		var ret WeakAliasDetermination
		return ret
	}

	return o.WeakAliasDetermination
}

// GetWeakAliasDeterminationOk returns a tuple with the WeakAliasDetermination field value
// and a boolean to check if the value has been set.
func (o *IndividualScreeningHitNames) GetWeakAliasDeterminationOk() (*WeakAliasDetermination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WeakAliasDetermination, true
}

// SetWeakAliasDetermination sets field value
func (o *IndividualScreeningHitNames) SetWeakAliasDetermination(v WeakAliasDetermination) {
	o.WeakAliasDetermination = v
}

func (o IndividualScreeningHitNames) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["full"] = o.Full
	}
	if true {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if true {
		toSerialize["weak_alias_determination"] = o.WeakAliasDetermination
	}
	return json.Marshal(toSerialize)
}

type NullableIndividualScreeningHitNames struct {
	value *IndividualScreeningHitNames
	isSet bool
}

func (v NullableIndividualScreeningHitNames) Get() *IndividualScreeningHitNames {
	return v.value
}

func (v *NullableIndividualScreeningHitNames) Set(val *IndividualScreeningHitNames) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualScreeningHitNames) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualScreeningHitNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualScreeningHitNames(val *IndividualScreeningHitNames) *NullableIndividualScreeningHitNames {
	return &NullableIndividualScreeningHitNames{value: val, isSet: true}
}

func (v NullableIndividualScreeningHitNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualScreeningHitNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


