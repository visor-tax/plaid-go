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

// IndividualName Parent container for name that allows for choice group between parsed and unparsed containers.Parent container for name that allows for choice group between parsed and unparsed containers.
type IndividualName struct {
	// The first name of the individual represented by the parent object.
	FirstName string `json:"FirstName"`
	// The last name of the individual represented by the parent object.
	LastName string `json:"LastName"`
	AdditionalProperties map[string]interface{}
}

type _IndividualName IndividualName

// NewIndividualName instantiates a new IndividualName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualName(firstName string, lastName string) *IndividualName {
	this := IndividualName{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewIndividualNameWithDefaults instantiates a new IndividualName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualNameWithDefaults() *IndividualName {
	this := IndividualName{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *IndividualName) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *IndividualName) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *IndividualName) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *IndividualName) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *IndividualName) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *IndividualName) SetLastName(v string) {
	o.LastName = v
}

func (o IndividualName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["FirstName"] = o.FirstName
	}
	if true {
		toSerialize["LastName"] = o.LastName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IndividualName) UnmarshalJSON(bytes []byte) (err error) {
	varIndividualName := _IndividualName{}

	if err = json.Unmarshal(bytes, &varIndividualName); err == nil {
		*o = IndividualName(varIndividualName)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "FirstName")
		delete(additionalProperties, "LastName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndividualName struct {
	value *IndividualName
	isSet bool
}

func (v NullableIndividualName) Get() *IndividualName {
	return v.value
}

func (v *NullableIndividualName) Set(val *IndividualName) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualName) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualName(val *IndividualName) *NullableIndividualName {
	return &NullableIndividualName{value: val, isSet: true}
}

func (v NullableIndividualName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


