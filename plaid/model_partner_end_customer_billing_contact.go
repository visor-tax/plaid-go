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

// PartnerEndCustomerBillingContact The billing contact for the end customer. Defaults to partner's technical contact if omitted.
type PartnerEndCustomerBillingContact struct {
	GivenName *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	Email *string `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartnerEndCustomerBillingContact PartnerEndCustomerBillingContact

// NewPartnerEndCustomerBillingContact instantiates a new PartnerEndCustomerBillingContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerEndCustomerBillingContact() *PartnerEndCustomerBillingContact {
	this := PartnerEndCustomerBillingContact{}
	return &this
}

// NewPartnerEndCustomerBillingContactWithDefaults instantiates a new PartnerEndCustomerBillingContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerEndCustomerBillingContactWithDefaults() *PartnerEndCustomerBillingContact {
	this := PartnerEndCustomerBillingContact{}
	return &this
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *PartnerEndCustomerBillingContact) GetGivenName() string {
	if o == nil || o.GivenName == nil {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerBillingContact) GetGivenNameOk() (*string, bool) {
	if o == nil || o.GivenName == nil {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *PartnerEndCustomerBillingContact) HasGivenName() bool {
	if o != nil && o.GivenName != nil {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *PartnerEndCustomerBillingContact) SetGivenName(v string) {
	o.GivenName = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *PartnerEndCustomerBillingContact) GetFamilyName() string {
	if o == nil || o.FamilyName == nil {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerBillingContact) GetFamilyNameOk() (*string, bool) {
	if o == nil || o.FamilyName == nil {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *PartnerEndCustomerBillingContact) HasFamilyName() bool {
	if o != nil && o.FamilyName != nil {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *PartnerEndCustomerBillingContact) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PartnerEndCustomerBillingContact) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerEndCustomerBillingContact) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PartnerEndCustomerBillingContact) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PartnerEndCustomerBillingContact) SetEmail(v string) {
	o.Email = &v
}

func (o PartnerEndCustomerBillingContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GivenName != nil {
		toSerialize["given_name"] = o.GivenName
	}
	if o.FamilyName != nil {
		toSerialize["family_name"] = o.FamilyName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PartnerEndCustomerBillingContact) UnmarshalJSON(bytes []byte) (err error) {
	varPartnerEndCustomerBillingContact := _PartnerEndCustomerBillingContact{}

	if err = json.Unmarshal(bytes, &varPartnerEndCustomerBillingContact); err == nil {
		*o = PartnerEndCustomerBillingContact(varPartnerEndCustomerBillingContact)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "given_name")
		delete(additionalProperties, "family_name")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartnerEndCustomerBillingContact struct {
	value *PartnerEndCustomerBillingContact
	isSet bool
}

func (v NullablePartnerEndCustomerBillingContact) Get() *PartnerEndCustomerBillingContact {
	return v.value
}

func (v *NullablePartnerEndCustomerBillingContact) Set(val *PartnerEndCustomerBillingContact) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerEndCustomerBillingContact) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerEndCustomerBillingContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerEndCustomerBillingContact(val *PartnerEndCustomerBillingContact) *NullablePartnerEndCustomerBillingContact {
	return &NullablePartnerEndCustomerBillingContact{value: val, isSet: true}
}

func (v NullablePartnerEndCustomerBillingContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerEndCustomerBillingContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


