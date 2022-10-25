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

// TransactionsRuleDetails A representation of transactions rule details.
type TransactionsRuleDetails struct {
	Field TransactionsRuleField `json:"field"`
	Type TransactionsRuleType `json:"type"`
	// For TRANSACTION_ID field, provide transaction_id. For NAME field, provide a string pattern. 
	Query string `json:"query"`
}

// NewTransactionsRuleDetails instantiates a new TransactionsRuleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRuleDetails(field TransactionsRuleField, type_ TransactionsRuleType, query string) *TransactionsRuleDetails {
	this := TransactionsRuleDetails{}
	this.Field = field
	this.Type = type_
	this.Query = query
	return &this
}

// NewTransactionsRuleDetailsWithDefaults instantiates a new TransactionsRuleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRuleDetailsWithDefaults() *TransactionsRuleDetails {
	this := TransactionsRuleDetails{}
	return &this
}

// GetField returns the Field field value
func (o *TransactionsRuleDetails) GetField() TransactionsRuleField {
	if o == nil {
		var ret TransactionsRuleField
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *TransactionsRuleDetails) GetFieldOk() (*TransactionsRuleField, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *TransactionsRuleDetails) SetField(v TransactionsRuleField) {
	o.Field = v
}

// GetType returns the Type field value
func (o *TransactionsRuleDetails) GetType() TransactionsRuleType {
	if o == nil {
		var ret TransactionsRuleType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransactionsRuleDetails) GetTypeOk() (*TransactionsRuleType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransactionsRuleDetails) SetType(v TransactionsRuleType) {
	o.Type = v
}

// GetQuery returns the Query field value
func (o *TransactionsRuleDetails) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *TransactionsRuleDetails) GetQueryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *TransactionsRuleDetails) SetQuery(v string) {
	o.Query = v
}

func (o TransactionsRuleDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRuleDetails struct {
	value *TransactionsRuleDetails
	isSet bool
}

func (v NullableTransactionsRuleDetails) Get() *TransactionsRuleDetails {
	return v.value
}

func (v *NullableTransactionsRuleDetails) Set(val *TransactionsRuleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRuleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRuleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRuleDetails(val *TransactionsRuleDetails) *NullableTransactionsRuleDetails {
	return &NullableTransactionsRuleDetails{value: val, isSet: true}
}

func (v NullableTransactionsRuleDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRuleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


