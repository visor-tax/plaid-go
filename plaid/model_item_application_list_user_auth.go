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

// ItemApplicationListUserAuth User authentication parameters, for clients making a request without an `access_token`. This is only allowed for select clients and will not be supported in the future. Most clients should call /item/import to obtain an access token before making a request.
type ItemApplicationListUserAuth struct {
	// Account username.
	UserId NullableString `json:"user_id,omitempty"`
	// Account username hashed by FI.
	FiUsernameHash NullableString `json:"fi_username_hash,omitempty"`
}

// NewItemApplicationListUserAuth instantiates a new ItemApplicationListUserAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemApplicationListUserAuth() *ItemApplicationListUserAuth {
	this := ItemApplicationListUserAuth{}
	return &this
}

// NewItemApplicationListUserAuthWithDefaults instantiates a new ItemApplicationListUserAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemApplicationListUserAuthWithDefaults() *ItemApplicationListUserAuth {
	this := ItemApplicationListUserAuth{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemApplicationListUserAuth) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemApplicationListUserAuth) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ItemApplicationListUserAuth) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *ItemApplicationListUserAuth) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ItemApplicationListUserAuth) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ItemApplicationListUserAuth) UnsetUserId() {
	o.UserId.Unset()
}

// GetFiUsernameHash returns the FiUsernameHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemApplicationListUserAuth) GetFiUsernameHash() string {
	if o == nil || o.FiUsernameHash.Get() == nil {
		var ret string
		return ret
	}
	return *o.FiUsernameHash.Get()
}

// GetFiUsernameHashOk returns a tuple with the FiUsernameHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemApplicationListUserAuth) GetFiUsernameHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FiUsernameHash.Get(), o.FiUsernameHash.IsSet()
}

// HasFiUsernameHash returns a boolean if a field has been set.
func (o *ItemApplicationListUserAuth) HasFiUsernameHash() bool {
	if o != nil && o.FiUsernameHash.IsSet() {
		return true
	}

	return false
}

// SetFiUsernameHash gets a reference to the given NullableString and assigns it to the FiUsernameHash field.
func (o *ItemApplicationListUserAuth) SetFiUsernameHash(v string) {
	o.FiUsernameHash.Set(&v)
}
// SetFiUsernameHashNil sets the value for FiUsernameHash to be an explicit nil
func (o *ItemApplicationListUserAuth) SetFiUsernameHashNil() {
	o.FiUsernameHash.Set(nil)
}

// UnsetFiUsernameHash ensures that no value is present for FiUsernameHash, not even an explicit nil
func (o *ItemApplicationListUserAuth) UnsetFiUsernameHash() {
	o.FiUsernameHash.Unset()
}

func (o ItemApplicationListUserAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	if o.FiUsernameHash.IsSet() {
		toSerialize["fi_username_hash"] = o.FiUsernameHash.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableItemApplicationListUserAuth struct {
	value *ItemApplicationListUserAuth
	isSet bool
}

func (v NullableItemApplicationListUserAuth) Get() *ItemApplicationListUserAuth {
	return v.value
}

func (v *NullableItemApplicationListUserAuth) Set(val *ItemApplicationListUserAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableItemApplicationListUserAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableItemApplicationListUserAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemApplicationListUserAuth(val *ItemApplicationListUserAuth) *NullableItemApplicationListUserAuth {
	return &NullableItemApplicationListUserAuth{value: val, isSet: true}
}

func (v NullableItemApplicationListUserAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemApplicationListUserAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


