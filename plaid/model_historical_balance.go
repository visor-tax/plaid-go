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

// HistoricalBalance An object representing a balance held by an account in the past
type HistoricalBalance struct {
	// The date of the calculated historical balance, in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (YYYY-MM-DD)
	Date string `json:"date"`
	// The total amount of funds in the account, calculated from the `current` balance in the `balance` object by subtracting inflows and adding back outflows according to the posted date of each transaction.  If the account has any pending transactions, historical balance amounts on or after the date of the earliest pending transaction may differ if retrieved in subsequent Asset Reports as a result of those pending transactions posting.
	Current float64 `json:"current"`
	// The ISO-4217 currency code of the balance. Always `null` if `unofficial_currency_code` is non-`null`.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the balance. Always `null` if `iso_currency_code` is non-`null`.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _HistoricalBalance HistoricalBalance

// NewHistoricalBalance instantiates a new HistoricalBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalBalance(date string, current float64, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString) *HistoricalBalance {
	this := HistoricalBalance{}
	this.Date = date
	this.Current = current
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewHistoricalBalanceWithDefaults instantiates a new HistoricalBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalBalanceWithDefaults() *HistoricalBalance {
	this := HistoricalBalance{}
	return &this
}

// GetDate returns the Date field value
func (o *HistoricalBalance) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *HistoricalBalance) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *HistoricalBalance) SetDate(v string) {
	o.Date = v
}

// GetCurrent returns the Current field value
func (o *HistoricalBalance) GetCurrent() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *HistoricalBalance) GetCurrentOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *HistoricalBalance) SetCurrent(v float64) {
	o.Current = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HistoricalBalance) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricalBalance) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *HistoricalBalance) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HistoricalBalance) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricalBalance) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *HistoricalBalance) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o HistoricalBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["current"] = o.Current
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HistoricalBalance) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalBalance := _HistoricalBalance{}

	if err = json.Unmarshal(bytes, &varHistoricalBalance); err == nil {
		*o = HistoricalBalance(varHistoricalBalance)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "current")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistoricalBalance struct {
	value *HistoricalBalance
	isSet bool
}

func (v NullableHistoricalBalance) Get() *HistoricalBalance {
	return v.value
}

func (v *NullableHistoricalBalance) Set(val *HistoricalBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalBalance(val *HistoricalBalance) *NullableHistoricalBalance {
	return &NullableHistoricalBalance{value: val, isSet: true}
}

func (v NullableHistoricalBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


