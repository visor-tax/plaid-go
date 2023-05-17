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

// Owner Data returned from the financial institution about the owner or owners of an account. Only the `names` array must be non-empty.
type Owner struct {
	// A list of names associated with the account by the financial institution. In the case of a joint account, Plaid will make a best effort to report the names of all account holders.  If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account's `names` array.
	Names []string `json:"names"`
	// A list of phone numbers associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	// A list of email addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.
	Emails []Email `json:"emails"`
	// Data about the various addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution.
	Addresses []Address `json:"addresses"`
	AdditionalProperties map[string]interface{}
}

type _Owner Owner

// NewOwner instantiates a new Owner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwner(names []string, phoneNumbers []PhoneNumber, emails []Email, addresses []Address) *Owner {
	this := Owner{}
	this.Names = names
	this.PhoneNumbers = phoneNumbers
	this.Emails = emails
	this.Addresses = addresses
	return &this
}

// NewOwnerWithDefaults instantiates a new Owner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerWithDefaults() *Owner {
	this := Owner{}
	return &this
}

// GetNames returns the Names field value
func (o *Owner) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *Owner) GetNamesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value
func (o *Owner) SetNames(v []string) {
	o.Names = v
}

// GetPhoneNumbers returns the PhoneNumbers field value
func (o *Owner) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		var ret []PhoneNumber
		return ret
	}

	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value
// and a boolean to check if the value has been set.
func (o *Owner) GetPhoneNumbersOk() (*[]PhoneNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PhoneNumbers, true
}

// SetPhoneNumbers sets field value
func (o *Owner) SetPhoneNumbers(v []PhoneNumber) {
	o.PhoneNumbers = v
}

// GetEmails returns the Emails field value
func (o *Owner) GetEmails() []Email {
	if o == nil {
		var ret []Email
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *Owner) GetEmailsOk() (*[]Email, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Emails, true
}

// SetEmails sets field value
func (o *Owner) SetEmails(v []Email) {
	o.Emails = v
}

// GetAddresses returns the Addresses field value
func (o *Owner) GetAddresses() []Address {
	if o == nil {
		var ret []Address
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *Owner) GetAddressesOk() (*[]Address, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *Owner) SetAddresses(v []Address) {
	o.Addresses = v
}

func (o Owner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["names"] = o.Names
	}
	if true {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Owner) UnmarshalJSON(bytes []byte) (err error) {
	varOwner := _Owner{}

	if err = json.Unmarshal(bytes, &varOwner); err == nil {
		*o = Owner(varOwner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "names")
		delete(additionalProperties, "phone_numbers")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOwner struct {
	value *Owner
	isSet bool
}

func (v NullableOwner) Get() *Owner {
	return v.value
}

func (v *NullableOwner) Set(val *Owner) {
	v.value = val
	v.isSet = true
}

func (v NullableOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwner(val *Owner) *NullableOwner {
	return &NullableOwner{value: val, isSet: true}
}

func (v NullableOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


