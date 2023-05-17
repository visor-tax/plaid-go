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

// IdentityMatchRequestOptions An optional object to filter /identity/match results
type IdentityMatchRequestOptions struct {
	// An array of `account_ids` to perform fuzzy match
	AccountIds *[]string `json:"account_ids,omitempty"`
}

// NewIdentityMatchRequestOptions instantiates a new IdentityMatchRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityMatchRequestOptions() *IdentityMatchRequestOptions {
	this := IdentityMatchRequestOptions{}
	return &this
}

// NewIdentityMatchRequestOptionsWithDefaults instantiates a new IdentityMatchRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityMatchRequestOptionsWithDefaults() *IdentityMatchRequestOptions {
	this := IdentityMatchRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *IdentityMatchRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityMatchRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *IdentityMatchRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *IdentityMatchRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o IdentityMatchRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityMatchRequestOptions struct {
	value *IdentityMatchRequestOptions
	isSet bool
}

func (v NullableIdentityMatchRequestOptions) Get() *IdentityMatchRequestOptions {
	return v.value
}

func (v *NullableIdentityMatchRequestOptions) Set(val *IdentityMatchRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityMatchRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityMatchRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityMatchRequestOptions(val *IdentityMatchRequestOptions) *NullableIdentityMatchRequestOptions {
	return &NullableIdentityMatchRequestOptions{value: val, isSet: true}
}

func (v NullableIdentityMatchRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityMatchRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


