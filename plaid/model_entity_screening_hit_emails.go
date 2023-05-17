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

// EntityScreeningHitEmails Email address information for the associated entity watchlist hit
type EntityScreeningHitEmails struct {
	// A valid email address.
	EmailAddress string `json:"email_address"`
	AdditionalProperties map[string]interface{}
}

type _EntityScreeningHitEmails EntityScreeningHitEmails

// NewEntityScreeningHitEmails instantiates a new EntityScreeningHitEmails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityScreeningHitEmails(emailAddress string) *EntityScreeningHitEmails {
	this := EntityScreeningHitEmails{}
	this.EmailAddress = emailAddress
	return &this
}

// NewEntityScreeningHitEmailsWithDefaults instantiates a new EntityScreeningHitEmails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityScreeningHitEmailsWithDefaults() *EntityScreeningHitEmails {
	this := EntityScreeningHitEmails{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *EntityScreeningHitEmails) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *EntityScreeningHitEmails) GetEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *EntityScreeningHitEmails) SetEmailAddress(v string) {
	o.EmailAddress = v
}

func (o EntityScreeningHitEmails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email_address"] = o.EmailAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntityScreeningHitEmails) UnmarshalJSON(bytes []byte) (err error) {
	varEntityScreeningHitEmails := _EntityScreeningHitEmails{}

	if err = json.Unmarshal(bytes, &varEntityScreeningHitEmails); err == nil {
		*o = EntityScreeningHitEmails(varEntityScreeningHitEmails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "email_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntityScreeningHitEmails struct {
	value *EntityScreeningHitEmails
	isSet bool
}

func (v NullableEntityScreeningHitEmails) Get() *EntityScreeningHitEmails {
	return v.value
}

func (v *NullableEntityScreeningHitEmails) Set(val *EntityScreeningHitEmails) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityScreeningHitEmails) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityScreeningHitEmails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityScreeningHitEmails(val *EntityScreeningHitEmails) *NullableEntityScreeningHitEmails {
	return &NullableEntityScreeningHitEmails{value: val, isSet: true}
}

func (v NullableEntityScreeningHitEmails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityScreeningHitEmails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


