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

// InvestmentsTransactionsGetRequestOptions An optional object to filter `/investments/transactions/get` results. If provided, must be non-`null`.
type InvestmentsTransactionsGetRequestOptions struct {
	// An array of `account_ids` to retrieve for the Item.
	AccountIds *[]string `json:"account_ids,omitempty"`
	// The number of transactions to fetch. 
	Count *int32 `json:"count,omitempty"`
	// The number of transactions to skip when fetching transaction history
	Offset *int32 `json:"offset,omitempty"`
}

// NewInvestmentsTransactionsGetRequestOptions instantiates a new InvestmentsTransactionsGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsTransactionsGetRequestOptions() *InvestmentsTransactionsGetRequestOptions {
	this := InvestmentsTransactionsGetRequestOptions{}
	var count int32 = 100
	this.Count = &count
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewInvestmentsTransactionsGetRequestOptionsWithDefaults instantiates a new InvestmentsTransactionsGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsTransactionsGetRequestOptionsWithDefaults() *InvestmentsTransactionsGetRequestOptions {
	this := InvestmentsTransactionsGetRequestOptions{}
	var count int32 = 100
	this.Count = &count
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *InvestmentsTransactionsGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *InvestmentsTransactionsGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *InvestmentsTransactionsGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InvestmentsTransactionsGetRequestOptions) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequestOptions) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InvestmentsTransactionsGetRequestOptions) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InvestmentsTransactionsGetRequestOptions) SetCount(v int32) {
	o.Count = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *InvestmentsTransactionsGetRequestOptions) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsGetRequestOptions) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *InvestmentsTransactionsGetRequestOptions) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *InvestmentsTransactionsGetRequestOptions) SetOffset(v int32) {
	o.Offset = &v
}

func (o InvestmentsTransactionsGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentsTransactionsGetRequestOptions struct {
	value *InvestmentsTransactionsGetRequestOptions
	isSet bool
}

func (v NullableInvestmentsTransactionsGetRequestOptions) Get() *InvestmentsTransactionsGetRequestOptions {
	return v.value
}

func (v *NullableInvestmentsTransactionsGetRequestOptions) Set(val *InvestmentsTransactionsGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsTransactionsGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsTransactionsGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsTransactionsGetRequestOptions(val *InvestmentsTransactionsGetRequestOptions) *NullableInvestmentsTransactionsGetRequestOptions {
	return &NullableInvestmentsTransactionsGetRequestOptions{value: val, isSet: true}
}

func (v NullableInvestmentsTransactionsGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsTransactionsGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


