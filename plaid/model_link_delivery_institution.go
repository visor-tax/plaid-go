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

// LinkDeliveryInstitution Information related to the financial institution.
type LinkDeliveryInstitution struct {
	// The full institution name, such as 'Wells Fargo'
	Name *string `json:"name,omitempty"`
	// The Plaid institution identifier
	InstitutionId *string `json:"institution_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkDeliveryInstitution LinkDeliveryInstitution

// NewLinkDeliveryInstitution instantiates a new LinkDeliveryInstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryInstitution() *LinkDeliveryInstitution {
	this := LinkDeliveryInstitution{}
	return &this
}

// NewLinkDeliveryInstitutionWithDefaults instantiates a new LinkDeliveryInstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryInstitutionWithDefaults() *LinkDeliveryInstitution {
	this := LinkDeliveryInstitution{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinkDeliveryInstitution) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryInstitution) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinkDeliveryInstitution) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinkDeliveryInstitution) SetName(v string) {
	o.Name = &v
}

// GetInstitutionId returns the InstitutionId field value if set, zero value otherwise.
func (o *LinkDeliveryInstitution) GetInstitutionId() string {
	if o == nil || o.InstitutionId == nil {
		var ret string
		return ret
	}
	return *o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryInstitution) GetInstitutionIdOk() (*string, bool) {
	if o == nil || o.InstitutionId == nil {
		return nil, false
	}
	return o.InstitutionId, true
}

// HasInstitutionId returns a boolean if a field has been set.
func (o *LinkDeliveryInstitution) HasInstitutionId() bool {
	if o != nil && o.InstitutionId != nil {
		return true
	}

	return false
}

// SetInstitutionId gets a reference to the given string and assigns it to the InstitutionId field.
func (o *LinkDeliveryInstitution) SetInstitutionId(v string) {
	o.InstitutionId = &v
}

func (o LinkDeliveryInstitution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.InstitutionId != nil {
		toSerialize["institution_id"] = o.InstitutionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkDeliveryInstitution) UnmarshalJSON(bytes []byte) (err error) {
	varLinkDeliveryInstitution := _LinkDeliveryInstitution{}

	if err = json.Unmarshal(bytes, &varLinkDeliveryInstitution); err == nil {
		*o = LinkDeliveryInstitution(varLinkDeliveryInstitution)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "institution_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkDeliveryInstitution struct {
	value *LinkDeliveryInstitution
	isSet bool
}

func (v NullableLinkDeliveryInstitution) Get() *LinkDeliveryInstitution {
	return v.value
}

func (v *NullableLinkDeliveryInstitution) Set(val *LinkDeliveryInstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryInstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryInstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryInstitution(val *LinkDeliveryInstitution) *NullableLinkDeliveryInstitution {
	return &NullableLinkDeliveryInstitution{value: val, isSet: true}
}

func (v NullableLinkDeliveryInstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryInstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


