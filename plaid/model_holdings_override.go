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

// HoldingsOverride Specify the holdings on the account.
type HoldingsOverride struct {
	// The last price given by the institution for this security
	InstitutionPrice float64 `json:"institution_price"`
	// The date at which `institution_price` was current. Must be formatted as an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) date.
	InstitutionPriceAsOf *string `json:"institution_price_as_of,omitempty"`
	// The average original value of the holding. Multiple cost basis values for the same security purchased at different prices are not supported.
	CostBasis *float64 `json:"cost_basis,omitempty"`
	// The total quantity of the asset held, as reported by the financial institution.
	Quantity float64 `json:"quantity"`
	// Either a valid `iso_currency_code` or `unofficial_currency_code`
	Currency string `json:"currency"`
	Security SecurityOverride `json:"security"`
}

// NewHoldingsOverride instantiates a new HoldingsOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHoldingsOverride(institutionPrice float64, quantity float64, currency string, security SecurityOverride) *HoldingsOverride {
	this := HoldingsOverride{}
	this.InstitutionPrice = institutionPrice
	this.Quantity = quantity
	this.Currency = currency
	this.Security = security
	return &this
}

// NewHoldingsOverrideWithDefaults instantiates a new HoldingsOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoldingsOverrideWithDefaults() *HoldingsOverride {
	this := HoldingsOverride{}
	return &this
}

// GetInstitutionPrice returns the InstitutionPrice field value
func (o *HoldingsOverride) GetInstitutionPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.InstitutionPrice
}

// GetInstitutionPriceOk returns a tuple with the InstitutionPrice field value
// and a boolean to check if the value has been set.
func (o *HoldingsOverride) GetInstitutionPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionPrice, true
}

// SetInstitutionPrice sets field value
func (o *HoldingsOverride) SetInstitutionPrice(v float64) {
	o.InstitutionPrice = v
}

// GetInstitutionPriceAsOf returns the InstitutionPriceAsOf field value if set, zero value otherwise.
func (o *HoldingsOverride) GetInstitutionPriceAsOf() string {
	if o == nil || o.InstitutionPriceAsOf == nil {
		var ret string
		return ret
	}
	return *o.InstitutionPriceAsOf
}

// GetInstitutionPriceAsOfOk returns a tuple with the InstitutionPriceAsOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HoldingsOverride) GetInstitutionPriceAsOfOk() (*string, bool) {
	if o == nil || o.InstitutionPriceAsOf == nil {
		return nil, false
	}
	return o.InstitutionPriceAsOf, true
}

// HasInstitutionPriceAsOf returns a boolean if a field has been set.
func (o *HoldingsOverride) HasInstitutionPriceAsOf() bool {
	if o != nil && o.InstitutionPriceAsOf != nil {
		return true
	}

	return false
}

// SetInstitutionPriceAsOf gets a reference to the given string and assigns it to the InstitutionPriceAsOf field.
func (o *HoldingsOverride) SetInstitutionPriceAsOf(v string) {
	o.InstitutionPriceAsOf = &v
}

// GetCostBasis returns the CostBasis field value if set, zero value otherwise.
func (o *HoldingsOverride) GetCostBasis() float64 {
	if o == nil || o.CostBasis == nil {
		var ret float64
		return ret
	}
	return *o.CostBasis
}

// GetCostBasisOk returns a tuple with the CostBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HoldingsOverride) GetCostBasisOk() (*float64, bool) {
	if o == nil || o.CostBasis == nil {
		return nil, false
	}
	return o.CostBasis, true
}

// HasCostBasis returns a boolean if a field has been set.
func (o *HoldingsOverride) HasCostBasis() bool {
	if o != nil && o.CostBasis != nil {
		return true
	}

	return false
}

// SetCostBasis gets a reference to the given float64 and assigns it to the CostBasis field.
func (o *HoldingsOverride) SetCostBasis(v float64) {
	o.CostBasis = &v
}

// GetQuantity returns the Quantity field value
func (o *HoldingsOverride) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *HoldingsOverride) GetQuantityOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *HoldingsOverride) SetQuantity(v float64) {
	o.Quantity = v
}

// GetCurrency returns the Currency field value
func (o *HoldingsOverride) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *HoldingsOverride) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *HoldingsOverride) SetCurrency(v string) {
	o.Currency = v
}

// GetSecurity returns the Security field value
func (o *HoldingsOverride) GetSecurity() SecurityOverride {
	if o == nil {
		var ret SecurityOverride
		return ret
	}

	return o.Security
}

// GetSecurityOk returns a tuple with the Security field value
// and a boolean to check if the value has been set.
func (o *HoldingsOverride) GetSecurityOk() (*SecurityOverride, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Security, true
}

// SetSecurity sets field value
func (o *HoldingsOverride) SetSecurity(v SecurityOverride) {
	o.Security = v
}

func (o HoldingsOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institution_price"] = o.InstitutionPrice
	}
	if o.InstitutionPriceAsOf != nil {
		toSerialize["institution_price_as_of"] = o.InstitutionPriceAsOf
	}
	if o.CostBasis != nil {
		toSerialize["cost_basis"] = o.CostBasis
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["security"] = o.Security
	}
	return json.Marshal(toSerialize)
}

type NullableHoldingsOverride struct {
	value *HoldingsOverride
	isSet bool
}

func (v NullableHoldingsOverride) Get() *HoldingsOverride {
	return v.value
}

func (v *NullableHoldingsOverride) Set(val *HoldingsOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableHoldingsOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableHoldingsOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoldingsOverride(val *HoldingsOverride) *NullableHoldingsOverride {
	return &NullableHoldingsOverride{value: val, isSet: true}
}

func (v NullableHoldingsOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoldingsOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


