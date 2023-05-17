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

// FDXInitiatorFiAttribute Initiator Fi Attribute
type FDXInitiatorFiAttribute struct {
	Name *string `json:"name,omitempty"`
	Value *FDXPartyType `json:"value,omitempty"`
}

// NewFDXInitiatorFiAttribute instantiates a new FDXInitiatorFiAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFDXInitiatorFiAttribute() *FDXInitiatorFiAttribute {
	this := FDXInitiatorFiAttribute{}
	return &this
}

// NewFDXInitiatorFiAttributeWithDefaults instantiates a new FDXInitiatorFiAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFDXInitiatorFiAttributeWithDefaults() *FDXInitiatorFiAttribute {
	this := FDXInitiatorFiAttribute{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FDXInitiatorFiAttribute) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXInitiatorFiAttribute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FDXInitiatorFiAttribute) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FDXInitiatorFiAttribute) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FDXInitiatorFiAttribute) GetValue() FDXPartyType {
	if o == nil || o.Value == nil {
		var ret FDXPartyType
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FDXInitiatorFiAttribute) GetValueOk() (*FDXPartyType, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FDXInitiatorFiAttribute) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given FDXPartyType and assigns it to the Value field.
func (o *FDXInitiatorFiAttribute) SetValue(v FDXPartyType) {
	o.Value = &v
}

func (o FDXInitiatorFiAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFDXInitiatorFiAttribute struct {
	value *FDXInitiatorFiAttribute
	isSet bool
}

func (v NullableFDXInitiatorFiAttribute) Get() *FDXInitiatorFiAttribute {
	return v.value
}

func (v *NullableFDXInitiatorFiAttribute) Set(val *FDXInitiatorFiAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableFDXInitiatorFiAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableFDXInitiatorFiAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFDXInitiatorFiAttribute(val *FDXInitiatorFiAttribute) *NullableFDXInitiatorFiAttribute {
	return &NullableFDXInitiatorFiAttribute{value: val, isSet: true}
}

func (v NullableFDXInitiatorFiAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFDXInitiatorFiAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


