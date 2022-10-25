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

// IdentityVerificationRequestUser User information collected outside of Link, most likely via your own onboarding process.  Each of the following identity fields are optional:  `email_address`  `phone_number`  `date_of_birth`  `name`  `address`  `id_number`  Specifically, these fields are optional in that they can either be fully provided (satisfying every required field in their subschema) or omitted from the request entirely by not providing the key or value. Providing these fields via the API will result in Link skipping the data collection process for the associated user. All verification steps enabled in the associated Identity Verification Template will still be run. Verification steps will either be run immediately, or once the user completes the `accept_tos` step, depending on the value provided to the `gave_consent` field.
type IdentityVerificationRequestUser struct {
	// An identifier to help you connect this object to your internal systems. For example, your database ID corresponding to this object.
	ClientUserId string `json:"client_user_id"`
	// A valid email address.
	EmailAddress *string `json:"email_address,omitempty"`
	// A phone number in E.164 format.
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	Name NullableUserName `json:"name,omitempty"`
	Address NullableUserAddress `json:"address,omitempty"`
	IdNumber NullableUserIDNumber `json:"id_number,omitempty"`
}

// NewIdentityVerificationRequestUser instantiates a new IdentityVerificationRequestUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationRequestUser(clientUserId string) *IdentityVerificationRequestUser {
	this := IdentityVerificationRequestUser{}
	this.ClientUserId = clientUserId
	return &this
}

// NewIdentityVerificationRequestUserWithDefaults instantiates a new IdentityVerificationRequestUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationRequestUserWithDefaults() *IdentityVerificationRequestUser {
	this := IdentityVerificationRequestUser{}
	return &this
}

// GetClientUserId returns the ClientUserId field value
func (o *IdentityVerificationRequestUser) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRequestUser) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *IdentityVerificationRequestUser) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *IdentityVerificationRequestUser) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRequestUser) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *IdentityVerificationRequestUser) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *IdentityVerificationRequestUser) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationRequestUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationRequestUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *IdentityVerificationRequestUser) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *IdentityVerificationRequestUser) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *IdentityVerificationRequestUser) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *IdentityVerificationRequestUser) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *IdentityVerificationRequestUser) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityVerificationRequestUser) GetDateOfBirthOk() (*string, bool) {
	if o == nil || o.DateOfBirth == nil {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *IdentityVerificationRequestUser) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth != nil {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *IdentityVerificationRequestUser) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationRequestUser) GetName() UserName {
	if o == nil || o.Name.Get() == nil {
		var ret UserName
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationRequestUser) GetNameOk() (*UserName, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdentityVerificationRequestUser) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableUserName and assigns it to the Name field.
func (o *IdentityVerificationRequestUser) SetName(v UserName) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IdentityVerificationRequestUser) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdentityVerificationRequestUser) UnsetName() {
	o.Name.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationRequestUser) GetAddress() UserAddress {
	if o == nil || o.Address.Get() == nil {
		var ret UserAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationRequestUser) GetAddressOk() (*UserAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *IdentityVerificationRequestUser) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableUserAddress and assigns it to the Address field.
func (o *IdentityVerificationRequestUser) SetAddress(v UserAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *IdentityVerificationRequestUser) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *IdentityVerificationRequestUser) UnsetAddress() {
	o.Address.Unset()
}

// GetIdNumber returns the IdNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityVerificationRequestUser) GetIdNumber() UserIDNumber {
	if o == nil || o.IdNumber.Get() == nil {
		var ret UserIDNumber
		return ret
	}
	return *o.IdNumber.Get()
}

// GetIdNumberOk returns a tuple with the IdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationRequestUser) GetIdNumberOk() (*UserIDNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IdNumber.Get(), o.IdNumber.IsSet()
}

// HasIdNumber returns a boolean if a field has been set.
func (o *IdentityVerificationRequestUser) HasIdNumber() bool {
	if o != nil && o.IdNumber.IsSet() {
		return true
	}

	return false
}

// SetIdNumber gets a reference to the given NullableUserIDNumber and assigns it to the IdNumber field.
func (o *IdentityVerificationRequestUser) SetIdNumber(v UserIDNumber) {
	o.IdNumber.Set(&v)
}
// SetIdNumberNil sets the value for IdNumber to be an explicit nil
func (o *IdentityVerificationRequestUser) SetIdNumberNil() {
	o.IdNumber.Set(nil)
}

// UnsetIdNumber ensures that no value is present for IdNumber, not even an explicit nil
func (o *IdentityVerificationRequestUser) UnsetIdNumber() {
	o.IdNumber.Unset()
}

func (o IdentityVerificationRequestUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.DateOfBirth != nil {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.IdNumber.IsSet() {
		toSerialize["id_number"] = o.IdNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityVerificationRequestUser struct {
	value *IdentityVerificationRequestUser
	isSet bool
}

func (v NullableIdentityVerificationRequestUser) Get() *IdentityVerificationRequestUser {
	return v.value
}

func (v *NullableIdentityVerificationRequestUser) Set(val *IdentityVerificationRequestUser) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationRequestUser) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationRequestUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationRequestUser(val *IdentityVerificationRequestUser) *NullableIdentityVerificationRequestUser {
	return &NullableIdentityVerificationRequestUser{value: val, isSet: true}
}

func (v NullableIdentityVerificationRequestUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationRequestUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


