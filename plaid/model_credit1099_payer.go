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

// Credit1099Payer An object representing a payer used by 1099-MISC tax documents.
type Credit1099Payer struct {
	Address *CreditPayStubAddress `json:"address,omitempty"`
	// Name of payer.
	Name NullableString `json:"name,omitempty"`
	// Tax identification number of payer.
	Tin NullableString `json:"tin,omitempty"`
	// Telephone number of payer.
	TelephoneNumber NullableString `json:"telephone_number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Credit1099Payer Credit1099Payer

// NewCredit1099Payer instantiates a new Credit1099Payer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredit1099Payer() *Credit1099Payer {
	this := Credit1099Payer{}
	return &this
}

// NewCredit1099PayerWithDefaults instantiates a new Credit1099Payer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredit1099PayerWithDefaults() *Credit1099Payer {
	this := Credit1099Payer{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Credit1099Payer) GetAddress() CreditPayStubAddress {
	if o == nil || o.Address == nil {
		var ret CreditPayStubAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit1099Payer) GetAddressOk() (*CreditPayStubAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Credit1099Payer) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given CreditPayStubAddress and assigns it to the Address field.
func (o *Credit1099Payer) SetAddress(v CreditPayStubAddress) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Payer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Payer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Credit1099Payer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Credit1099Payer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Credit1099Payer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Credit1099Payer) UnsetName() {
	o.Name.Unset()
}

// GetTin returns the Tin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Payer) GetTin() string {
	if o == nil || o.Tin.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tin.Get()
}

// GetTinOk returns a tuple with the Tin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Payer) GetTinOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tin.Get(), o.Tin.IsSet()
}

// HasTin returns a boolean if a field has been set.
func (o *Credit1099Payer) HasTin() bool {
	if o != nil && o.Tin.IsSet() {
		return true
	}

	return false
}

// SetTin gets a reference to the given NullableString and assigns it to the Tin field.
func (o *Credit1099Payer) SetTin(v string) {
	o.Tin.Set(&v)
}
// SetTinNil sets the value for Tin to be an explicit nil
func (o *Credit1099Payer) SetTinNil() {
	o.Tin.Set(nil)
}

// UnsetTin ensures that no value is present for Tin, not even an explicit nil
func (o *Credit1099Payer) UnsetTin() {
	o.Tin.Unset()
}

// GetTelephoneNumber returns the TelephoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Credit1099Payer) GetTelephoneNumber() string {
	if o == nil || o.TelephoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.TelephoneNumber.Get()
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Credit1099Payer) GetTelephoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TelephoneNumber.Get(), o.TelephoneNumber.IsSet()
}

// HasTelephoneNumber returns a boolean if a field has been set.
func (o *Credit1099Payer) HasTelephoneNumber() bool {
	if o != nil && o.TelephoneNumber.IsSet() {
		return true
	}

	return false
}

// SetTelephoneNumber gets a reference to the given NullableString and assigns it to the TelephoneNumber field.
func (o *Credit1099Payer) SetTelephoneNumber(v string) {
	o.TelephoneNumber.Set(&v)
}
// SetTelephoneNumberNil sets the value for TelephoneNumber to be an explicit nil
func (o *Credit1099Payer) SetTelephoneNumberNil() {
	o.TelephoneNumber.Set(nil)
}

// UnsetTelephoneNumber ensures that no value is present for TelephoneNumber, not even an explicit nil
func (o *Credit1099Payer) UnsetTelephoneNumber() {
	o.TelephoneNumber.Unset()
}

func (o Credit1099Payer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Tin.IsSet() {
		toSerialize["tin"] = o.Tin.Get()
	}
	if o.TelephoneNumber.IsSet() {
		toSerialize["telephone_number"] = o.TelephoneNumber.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Credit1099Payer) UnmarshalJSON(bytes []byte) (err error) {
	varCredit1099Payer := _Credit1099Payer{}

	if err = json.Unmarshal(bytes, &varCredit1099Payer); err == nil {
		*o = Credit1099Payer(varCredit1099Payer)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tin")
		delete(additionalProperties, "telephone_number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCredit1099Payer struct {
	value *Credit1099Payer
	isSet bool
}

func (v NullableCredit1099Payer) Get() *Credit1099Payer {
	return v.value
}

func (v *NullableCredit1099Payer) Set(val *Credit1099Payer) {
	v.value = val
	v.isSet = true
}

func (v NullableCredit1099Payer) IsSet() bool {
	return v.isSet
}

func (v *NullableCredit1099Payer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredit1099Payer(val *Credit1099Payer) *NullableCredit1099Payer {
	return &NullableCredit1099Payer{value: val, isSet: true}
}

func (v NullableCredit1099Payer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredit1099Payer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


