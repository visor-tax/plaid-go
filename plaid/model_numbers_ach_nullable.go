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

// NumbersACHNullable Identifying information for transferring money to or from a US account via ACH or wire transfer.
type NumbersACHNullable struct {
	// The Plaid account ID associated with the account numbers
	AccountId string `json:"account_id"`
	// The ACH account number for the account.  Note that when using OAuth with Chase Bank (`ins_56`), Chase will issue \"tokenized\" routing and account numbers, which are not the user's actual account and routing numbers. These tokenized numbers should work identically to normal account and routing numbers. The digits returned in the `mask` field will continue to reflect the actual account number, rather than the tokenized account number; for this reason, when displaying account numbers to the user to help them identify their account in your UI, always use the `mask` rather than truncating the `account` number. If a user revokes their permissions to your app, the tokenized numbers will continue to work for ACH deposits, but not withdrawals.
	Account string `json:"account"`
	// The ACH routing number for the account. If the institution is `ins_56`, this may be a tokenized routing number. For more information, see the description of the `account` field.
	Routing string `json:"routing"`
	// The wire transfer routing number for the account, if available
	WireRouting NullableString `json:"wire_routing"`
	// Whether the account supports ACH transfers into the account
	CanTransferIn NullableBool `json:"can_transfer_in,omitempty"`
	// Whether the account supports ACH transfers out of the account
	CanTransferOut NullableBool `json:"can_transfer_out,omitempty"`
}

// NewNumbersACHNullable instantiates a new NumbersACHNullable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumbersACHNullable(accountId string, account string, routing string, wireRouting NullableString) *NumbersACHNullable {
	this := NumbersACHNullable{}
	this.AccountId = accountId
	this.Account = account
	this.Routing = routing
	this.WireRouting = wireRouting
	return &this
}

// NewNumbersACHNullableWithDefaults instantiates a new NumbersACHNullable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumbersACHNullableWithDefaults() *NumbersACHNullable {
	this := NumbersACHNullable{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *NumbersACHNullable) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *NumbersACHNullable) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *NumbersACHNullable) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccount returns the Account field value
func (o *NumbersACHNullable) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NumbersACHNullable) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NumbersACHNullable) SetAccount(v string) {
	o.Account = v
}

// GetRouting returns the Routing field value
func (o *NumbersACHNullable) GetRouting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Routing
}

// GetRoutingOk returns a tuple with the Routing field value
// and a boolean to check if the value has been set.
func (o *NumbersACHNullable) GetRoutingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Routing, true
}

// SetRouting sets field value
func (o *NumbersACHNullable) SetRouting(v string) {
	o.Routing = v
}

// GetWireRouting returns the WireRouting field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NumbersACHNullable) GetWireRouting() string {
	if o == nil || o.WireRouting.Get() == nil {
		var ret string
		return ret
	}

	return *o.WireRouting.Get()
}

// GetWireRoutingOk returns a tuple with the WireRouting field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NumbersACHNullable) GetWireRoutingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WireRouting.Get(), o.WireRouting.IsSet()
}

// SetWireRouting sets field value
func (o *NumbersACHNullable) SetWireRouting(v string) {
	o.WireRouting.Set(&v)
}

// GetCanTransferIn returns the CanTransferIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NumbersACHNullable) GetCanTransferIn() bool {
	if o == nil || o.CanTransferIn.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CanTransferIn.Get()
}

// GetCanTransferInOk returns a tuple with the CanTransferIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NumbersACHNullable) GetCanTransferInOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CanTransferIn.Get(), o.CanTransferIn.IsSet()
}

// HasCanTransferIn returns a boolean if a field has been set.
func (o *NumbersACHNullable) HasCanTransferIn() bool {
	if o != nil && o.CanTransferIn.IsSet() {
		return true
	}

	return false
}

// SetCanTransferIn gets a reference to the given NullableBool and assigns it to the CanTransferIn field.
func (o *NumbersACHNullable) SetCanTransferIn(v bool) {
	o.CanTransferIn.Set(&v)
}
// SetCanTransferInNil sets the value for CanTransferIn to be an explicit nil
func (o *NumbersACHNullable) SetCanTransferInNil() {
	o.CanTransferIn.Set(nil)
}

// UnsetCanTransferIn ensures that no value is present for CanTransferIn, not even an explicit nil
func (o *NumbersACHNullable) UnsetCanTransferIn() {
	o.CanTransferIn.Unset()
}

// GetCanTransferOut returns the CanTransferOut field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NumbersACHNullable) GetCanTransferOut() bool {
	if o == nil || o.CanTransferOut.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CanTransferOut.Get()
}

// GetCanTransferOutOk returns a tuple with the CanTransferOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NumbersACHNullable) GetCanTransferOutOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CanTransferOut.Get(), o.CanTransferOut.IsSet()
}

// HasCanTransferOut returns a boolean if a field has been set.
func (o *NumbersACHNullable) HasCanTransferOut() bool {
	if o != nil && o.CanTransferOut.IsSet() {
		return true
	}

	return false
}

// SetCanTransferOut gets a reference to the given NullableBool and assigns it to the CanTransferOut field.
func (o *NumbersACHNullable) SetCanTransferOut(v bool) {
	o.CanTransferOut.Set(&v)
}
// SetCanTransferOutNil sets the value for CanTransferOut to be an explicit nil
func (o *NumbersACHNullable) SetCanTransferOutNil() {
	o.CanTransferOut.Set(nil)
}

// UnsetCanTransferOut ensures that no value is present for CanTransferOut, not even an explicit nil
func (o *NumbersACHNullable) UnsetCanTransferOut() {
	o.CanTransferOut.Unset()
}

func (o NumbersACHNullable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["routing"] = o.Routing
	}
	if true {
		toSerialize["wire_routing"] = o.WireRouting.Get()
	}
	if o.CanTransferIn.IsSet() {
		toSerialize["can_transfer_in"] = o.CanTransferIn.Get()
	}
	if o.CanTransferOut.IsSet() {
		toSerialize["can_transfer_out"] = o.CanTransferOut.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNumbersACHNullable struct {
	value *NumbersACHNullable
	isSet bool
}

func (v NullableNumbersACHNullable) Get() *NumbersACHNullable {
	return v.value
}

func (v *NullableNumbersACHNullable) Set(val *NumbersACHNullable) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersACHNullable) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersACHNullable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersACHNullable(val *NumbersACHNullable) *NullableNumbersACHNullable {
	return &NullableNumbersACHNullable{value: val, isSet: true}
}

func (v NullableNumbersACHNullable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersACHNullable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


