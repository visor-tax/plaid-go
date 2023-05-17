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

// EntityScreeningHitNames Name information for the associated entity watchlist hit
type EntityScreeningHitNames struct {
	// The full name of the entity.
	Full string `json:"full"`
	// Primary names are those most commonly used to refer to this entity. Only one name will ever be marked as primary.
	IsPrimary bool `json:"is_primary"`
	WeakAliasDetermination WeakAliasDetermination `json:"weak_alias_determination"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningHitNames EntityScreeningHitNames

// NewEntityScreeningHitNames instantiates a new EntityScreeningHitNames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitNames(full string, isPrimary bool, weakAliasDetermination WeakAliasDetermination) *EntityScreeningHitNames {
	this := EntityScreeningHitNames{}
	this.Full = full
	this.IsPrimary = isPrimary
	this.WeakAliasDetermination = weakAliasDetermination
	return &this
}

// NewEntityScreeningHitNamesWithDefaults instantiates a new EntityScreeningHitNames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitNamesWithDefaults() *EntityScreeningHitNames {
	this := EntityScreeningHitNames{}
	return &this
}

// GetFull returns the Full field value
func (o *EntityScreeningHitNames) GetFull() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Full
}

// GetFullOk returns a tuple with the Full field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitNames) GetFullOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Full, true
}

// SetFull sets field value
func (o *EntityScreeningHitNames) SetFull(v string) {
	o.Full = v
}

// GetIsPrimary returns the IsPrimary field value
func (o *EntityScreeningHitNames) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitNames) GetIsPrimaryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value
func (o *EntityScreeningHitNames) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetWeakAliasDetermination returns the WeakAliasDetermination field value
func (o *EntityScreeningHitNames) GetWeakAliasDetermination() WeakAliasDetermination {
	if o == nil {
		var ret WeakAliasDetermination
		return ret
	}

	return o.WeakAliasDetermination
}

// GetWeakAliasDeterminationOk returns a tuple with the WeakAliasDetermination field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitNames) GetWeakAliasDeterminationOk() (*WeakAliasDetermination, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WeakAliasDetermination, true
}

// SetWeakAliasDetermination sets field value
func (o *EntityScreeningHitNames) SetWeakAliasDetermination(v WeakAliasDetermination) {
	o.WeakAliasDetermination = v
}

func (o EntityScreeningHitNames) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningHitNames) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningHitNames := _EntityScreeningHitNames{}

	if err = json.Unmarshal(bytes, &varEntityScreeningHitNames); err == nil {
		*o = EntityScreeningHitNames(varEntityScreeningHitNames)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "full")
		delete(additionalProperties, "is_primary")
		delete(additionalProperties, "weak_alias_determination")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningHitNames struct {
	value *EntityScreeningHitNames
	isSet bool
}

func (v NullableEntityScreeningHitNames) Get() *EntityScreeningHitNames {
	return v.value
}

func (v *NullableEntityScreeningHitNames) Set(val *EntityScreeningHitNames) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitNames) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitNames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitNames(val *EntityScreeningHitNames) *NullableEntityScreeningHitNames {
	return &NullableEntityScreeningHitNames{value: val, isSet: true}
}

func (v NullableEntityScreeningHitNames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitNames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


