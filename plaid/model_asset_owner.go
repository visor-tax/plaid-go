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

// AssetOwner No documentation available
type AssetOwner struct {
	// Account Owner Full Name.
	AssetOwnerText NullableString `json:"AssetOwnerText"`
	AdditionalProperties map[string]interface{}
}

type _AssetOwner AssetOwner

// NewAssetOwner instantiates a new AssetOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOwner(assetOwnerText NullableString) *AssetOwner {
	this := AssetOwner{}
	this.AssetOwnerText = assetOwnerText
	return &this
}

// NewAssetOwnerWithDefaults instantiates a new AssetOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOwnerWithDefaults() *AssetOwner {
	this := AssetOwner{}
	return &this
}

// GetAssetOwnerText returns the AssetOwnerText field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetOwner) GetAssetOwnerText() string {
	if o == nil || o.AssetOwnerText.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetOwnerText.Get()
}

// GetAssetOwnerTextOk returns a tuple with the AssetOwnerText field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetOwner) GetAssetOwnerTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AssetOwnerText.Get(), o.AssetOwnerText.IsSet()
}

// SetAssetOwnerText sets field value
func (o *AssetOwner) SetAssetOwnerText(v string) {
	o.AssetOwnerText.Set(&v)
}

func (o AssetOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["AssetOwnerText"] = o.AssetOwnerText.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOwner) UnmarshalJSON(bytes []byte) (err error) {
	varAssetOwner := _AssetOwner{}

	if err = json.Unmarshal(bytes, &varAssetOwner); err == nil {
		*o = AssetOwner(varAssetOwner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AssetOwnerText")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetOwner struct {
	value *AssetOwner
	isSet bool
}

func (v NullableAssetOwner) Get() *AssetOwner {
	return v.value
}

func (v *NullableAssetOwner) Set(val *AssetOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOwner(val *AssetOwner) *NullableAssetOwner {
	return &NullableAssetOwner{value: val, isSet: true}
}

func (v NullableAssetOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


