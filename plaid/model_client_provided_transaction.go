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

// ClientProvidedTransaction A client-provided transaction for Plaid to enrich.
type ClientProvidedTransaction struct {
	// A unique ID for the transaction used to help you tie data back to your systems.
	Id string `json:"id"`
	// The raw description of the transaction. If you have location data in available an unstructured format, it may be appended to the `description` field.
	Description string `json:"description"`
	// The absolute value of the transaction (>= 0). When testing Enrich, note that `amount` data should be realistic. Unrealistic or inaccurate `amount` data may result in reduced quality output.
	Amount float64 `json:"amount"`
	Direction EnrichTransactionDirection `json:"direction"`
	// The ISO-4217 currency code of the transaction e.g. USD.
	IsoCurrencyCode string `json:"iso_currency_code"`
	Location *ClientProvidedTransactionLocation `json:"location,omitempty"`
	// Merchant category codes (MCCs) are four-digit numbers that describe a merchant's primary business activities.
	Mcc *string `json:"mcc,omitempty"`
	// The date the transaction posted, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) (YYYY-MM-DD) format.
	DatePosted *string `json:"date_posted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientProvidedTransaction ClientProvidedTransaction

// NewClientProvidedTransaction instantiates a new ClientProvidedTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProvidedTransaction(id string, description string, amount float64, direction EnrichTransactionDirection, isoCurrencyCode string) *ClientProvidedTransaction {
	this := ClientProvidedTransaction{}
	this.Id = id
	this.Description = description
	this.Amount = amount
	this.Direction = direction
	this.IsoCurrencyCode = isoCurrencyCode
	return &this
}

// NewClientProvidedTransactionWithDefaults instantiates a new ClientProvidedTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProvidedTransactionWithDefaults() *ClientProvidedTransaction {
	this := ClientProvidedTransaction{}
	return &this
}

// GetId returns the Id field value
func (o *ClientProvidedTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientProvidedTransaction) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *ClientProvidedTransaction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ClientProvidedTransaction) SetDescription(v string) {
	o.Description = v
}

// GetAmount returns the Amount field value
func (o *ClientProvidedTransaction) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ClientProvidedTransaction) SetAmount(v float64) {
	o.Amount = v
}

// GetDirection returns the Direction field value
func (o *ClientProvidedTransaction) GetDirection() EnrichTransactionDirection {
	if o == nil {
		var ret EnrichTransactionDirection
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetDirectionOk() (*EnrichTransactionDirection, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *ClientProvidedTransaction) SetDirection(v EnrichTransactionDirection) {
	o.Direction = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
func (o *ClientProvidedTransaction) GetIsoCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCurrencyCode
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCurrencyCode, true
}

// SetIsoCurrencyCode sets field value
func (o *ClientProvidedTransaction) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ClientProvidedTransaction) GetLocation() ClientProvidedTransactionLocation {
	if o == nil || o.Location == nil {
		var ret ClientProvidedTransactionLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetLocationOk() (*ClientProvidedTransactionLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ClientProvidedTransaction) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ClientProvidedTransactionLocation and assigns it to the Location field.
func (o *ClientProvidedTransaction) SetLocation(v ClientProvidedTransactionLocation) {
	o.Location = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *ClientProvidedTransaction) GetMcc() string {
	if o == nil || o.Mcc == nil {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetMccOk() (*string, bool) {
	if o == nil || o.Mcc == nil {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *ClientProvidedTransaction) HasMcc() bool {
	if o != nil && o.Mcc != nil {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *ClientProvidedTransaction) SetMcc(v string) {
	o.Mcc = &v
}

// GetDatePosted returns the DatePosted field value if set, zero value otherwise.
func (o *ClientProvidedTransaction) GetDatePosted() string {
	if o == nil || o.DatePosted == nil {
		var ret string
		return ret
	}
	return *o.DatePosted
}

// GetDatePostedOk returns a tuple with the DatePosted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransaction) GetDatePostedOk() (*string, bool) {
	if o == nil || o.DatePosted == nil {
		return nil, false
	}
	return o.DatePosted, true
}

// HasDatePosted returns a boolean if a field has been set.
func (o *ClientProvidedTransaction) HasDatePosted() bool {
	if o != nil && o.DatePosted != nil {
		return true
	}

	return false
}

// SetDatePosted gets a reference to the given string and assigns it to the DatePosted field.
func (o *ClientProvidedTransaction) SetDatePosted(v string) {
	o.DatePosted = &v
}

func (o ClientProvidedTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["direction"] = o.Direction
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Mcc != nil {
		toSerialize["mcc"] = o.Mcc
	}
	if o.DatePosted != nil {
		toSerialize["date_posted"] = o.DatePosted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientProvidedTransaction) UnmarshalJSON(bytes []byte) (err error) {
	varClientProvidedTransaction := _ClientProvidedTransaction{}

	if err = json.Unmarshal(bytes, &varClientProvidedTransaction); err == nil {
		*o = ClientProvidedTransaction(varClientProvidedTransaction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "location")
		delete(additionalProperties, "mcc")
		delete(additionalProperties, "date_posted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientProvidedTransaction struct {
	value *ClientProvidedTransaction
	isSet bool
}

func (v NullableClientProvidedTransaction) Get() *ClientProvidedTransaction {
	return v.value
}

func (v *NullableClientProvidedTransaction) Set(val *ClientProvidedTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProvidedTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProvidedTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProvidedTransaction(val *ClientProvidedTransaction) *NullableClientProvidedTransaction {
	return &NullableClientProvidedTransaction{value: val, isSet: true}
}

func (v NullableClientProvidedTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProvidedTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


