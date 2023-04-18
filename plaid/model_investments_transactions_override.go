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

// InvestmentsTransactionsOverride Specify the list of investments transactions on the account.
type InvestmentsTransactionsOverride struct {
	// Posting date for the transaction. Must be formatted as an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date.
	Date string `json:"date"`
	// The institution's description of the transaction.
	Name string `json:"name"`
	// The number of units of the security involved in this transaction. Must be positive if the type is a buy and negative if the type is a sell.
	Quantity float64 `json:"quantity"`
	// The price of the security at which this transaction occurred.
	Price float64 `json:"price"`
	// The combined value of all fees applied to this transaction.
	Fees *float64 `json:"fees,omitempty"`
	// The type of the investment transaction. Possible values are: `buy`: Buying an investment `sell`: Selling an investment `cash`: Activity that modifies a cash position `fee`: A fee on the account `transfer`: Activity that modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer
	Type string `json:"type"`
	// Either a valid `iso_currency_code` or `unofficial_currency_code`
	Currency string `json:"currency"`
	Security *SecurityOverride `json:"security,omitempty"`
}

// NewInvestmentsTransactionsOverride instantiates a new InvestmentsTransactionsOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsTransactionsOverride(date string, name string, quantity float64, price float64, type_ string, currency string) *InvestmentsTransactionsOverride {
	this := InvestmentsTransactionsOverride{}
	this.Date = date
	this.Name = name
	this.Quantity = quantity
	this.Price = price
	this.Type = type_
	this.Currency = currency
	return &this
}

// NewInvestmentsTransactionsOverrideWithDefaults instantiates a new InvestmentsTransactionsOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsTransactionsOverrideWithDefaults() *InvestmentsTransactionsOverride {
	this := InvestmentsTransactionsOverride{}
	return &this
}

// GetDate returns the Date field value
func (o *InvestmentsTransactionsOverride) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *InvestmentsTransactionsOverride) SetDate(v string) {
	o.Date = v
}

// GetName returns the Name field value
func (o *InvestmentsTransactionsOverride) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvestmentsTransactionsOverride) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *InvestmentsTransactionsOverride) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetQuantityOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InvestmentsTransactionsOverride) SetQuantity(v float64) {
	o.Quantity = v
}

// GetPrice returns the Price field value
func (o *InvestmentsTransactionsOverride) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InvestmentsTransactionsOverride) SetPrice(v float64) {
	o.Price = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *InvestmentsTransactionsOverride) GetFees() float64 {
	if o == nil || o.Fees == nil {
		var ret float64
		return ret
	}
	return *o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetFeesOk() (*float64, bool) {
	if o == nil || o.Fees == nil {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *InvestmentsTransactionsOverride) HasFees() bool {
	if o != nil && o.Fees != nil {
		return true
	}

	return false
}

// SetFees gets a reference to the given float64 and assigns it to the Fees field.
func (o *InvestmentsTransactionsOverride) SetFees(v float64) {
	o.Fees = &v
}

// GetType returns the Type field value
func (o *InvestmentsTransactionsOverride) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InvestmentsTransactionsOverride) SetType(v string) {
	o.Type = v
}

// GetCurrency returns the Currency field value
func (o *InvestmentsTransactionsOverride) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InvestmentsTransactionsOverride) SetCurrency(v string) {
	o.Currency = v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InvestmentsTransactionsOverride) GetSecurity() SecurityOverride {
	if o == nil || o.Security == nil {
		var ret SecurityOverride
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsTransactionsOverride) GetSecurityOk() (*SecurityOverride, bool) {
	if o == nil || o.Security == nil {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InvestmentsTransactionsOverride) HasSecurity() bool {
	if o != nil && o.Security != nil {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given SecurityOverride and assigns it to the Security field.
func (o *InvestmentsTransactionsOverride) SetSecurity(v SecurityOverride) {
	o.Security = &v
}

func (o InvestmentsTransactionsOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["price"] = o.Price
	}
	if o.Fees != nil {
		toSerialize["fees"] = o.Fees
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if o.Security != nil {
		toSerialize["security"] = o.Security
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentsTransactionsOverride struct {
	value *InvestmentsTransactionsOverride
	isSet bool
}

func (v NullableInvestmentsTransactionsOverride) Get() *InvestmentsTransactionsOverride {
	return v.value
}

func (v *NullableInvestmentsTransactionsOverride) Set(val *InvestmentsTransactionsOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsTransactionsOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsTransactionsOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsTransactionsOverride(val *InvestmentsTransactionsOverride) *NullableInvestmentsTransactionsOverride {
	return &NullableInvestmentsTransactionsOverride{value: val, isSet: true}
}

func (v NullableInvestmentsTransactionsOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsTransactionsOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


