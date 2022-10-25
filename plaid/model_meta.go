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

// Meta Allows specifying the metadata of the test account
type Meta struct {
	// The account's name
	Name string `json:"name"`
	// The account's official name
	OfficialName string `json:"official_name"`
	// The account's limit
	Limit float64 `json:"limit"`
	// The account's mask. Should be a string of 2-4 alphanumeric characters. This allows you to model a mask which does not match the account number (such as with a virtual account number).
	Mask string `json:"mask"`
	AdditionalProperties map[string]interface{}
}

type _Meta Meta

// NewMeta instantiates a new Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeta(name string, officialName string, limit float64, mask string) *Meta {
	this := Meta{}
	this.Name = name
	this.OfficialName = officialName
	this.Limit = limit
	this.Mask = mask
	return &this
}

// NewMetaWithDefaults instantiates a new Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaWithDefaults() *Meta {
	this := Meta{}
	return &this
}

// GetName returns the Name field value
func (o *Meta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Meta) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Meta) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
func (o *Meta) GetOfficialName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfficialName
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
func (o *Meta) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OfficialName, true
}

// SetOfficialName sets field value
func (o *Meta) SetOfficialName(v string) {
	o.OfficialName = v
}

// GetLimit returns the Limit field value
func (o *Meta) GetLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *Meta) GetLimitOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *Meta) SetLimit(v float64) {
	o.Limit = v
}

// GetMask returns the Mask field value
func (o *Meta) GetMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
func (o *Meta) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mask, true
}

// SetMask sets field value
func (o *Meta) SetMask(v string) {
	o.Mask = v
}

func (o Meta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["official_name"] = o.OfficialName
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["mask"] = o.Mask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Meta) UnmarshalJSON(bytes []byte) (err error) {
	varMeta := _Meta{}

	if err = json.Unmarshal(bytes, &varMeta); err == nil {
		*o = Meta(varMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "official_name")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "mask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeta struct {
	value *Meta
	isSet bool
}

func (v NullableMeta) Get() *Meta {
	return v.value
}

func (v *NullableMeta) Set(val *Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeta(val *Meta) *NullableMeta {
	return &NullableMeta{value: val, isSet: true}
}

func (v NullableMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


