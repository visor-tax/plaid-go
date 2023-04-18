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
	"time"
)

// BankTransferSweep BankTransferSweep describes a sweep transfer.
type BankTransferSweep struct {
	// Identifier of the sweep.
	Id string `json:"id"`
	// The datetime when the sweep occurred, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The amount of the sweep.
	Amount string `json:"amount"`
	// The currency of the sweep, e.g. \"USD\".
	IsoCurrencyCode string `json:"iso_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferSweep BankTransferSweep

// NewBankTransferSweep instantiates a new BankTransferSweep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferSweep(id string, createdAt time.Time, amount string, isoCurrencyCode string) *BankTransferSweep {
	this := BankTransferSweep{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewBankTransferSweepWithDefaults instantiates a new BankTransferSweep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferSweepWithDefaults() *BankTransferSweep {
	this := BankTransferSweep{}
	return &this
}

// GetId returns the Id field value
func (o *BankTransferSweep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweep) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BankTransferSweep) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BankTransferSweep) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweep) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BankTransferSweep) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetAmount returns the Amount field value
func (o *BankTransferSweep) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweep) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BankTransferSweep) SetAmount(v string) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *BankTransferSweep) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *BankTransferSweep) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *BankTransferSweep) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

func (o BankTransferSweep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferSweep) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferSweep := _BankTransferSweep{}

	if err = json.Unmarshal(bytes, &varBankTransferSweep); err == nil {
		*o = BankTransferSweep(varBankTransferSweep)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferSweep struct {
	value *BankTransferSweep
	isSet bool
}

func (v NullableBankTransferSweep) Get() *BankTransferSweep {
	return v.value
}

func (v *NullableBankTransferSweep) Set(val *BankTransferSweep) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferSweep) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferSweep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferSweep(val *BankTransferSweep) *NullableBankTransferSweep {
	return &NullableBankTransferSweep{value: val, isSet: true}
}

func (v NullableBankTransferSweep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferSweep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


