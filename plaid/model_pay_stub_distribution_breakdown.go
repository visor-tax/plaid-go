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

// PayStubDistributionBreakdown Information about the accounts that the payment was distributed to.
type PayStubDistributionBreakdown struct {
	// Name of the account for the given distribution.
	AccountName NullableString `json:"account_name"`
	// The name of the bank that the payment is being deposited to.
	BankName NullableString `json:"bank_name"`
	// The amount distributed to this account.
	CurrentAmount NullableFloat64 `json:"current_amount"`
	// The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The last 2-4 alphanumeric characters of an account's official account number.
	Mask NullableString `json:"mask"`
	// Type of the account that the paystub was sent to (e.g. 'checking').
	Type NullableString `json:"type"`
	// The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _PayStubDistributionBreakdown PayStubDistributionBreakdown

// NewPayStubDistributionBreakdown instantiates a new PayStubDistributionBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStubDistributionBreakdown(accountName NullableString, bankName NullableString, currentAmount NullableFloat64, isoCurrencyCode NullableString, mask NullableString, type_ NullableString, unofficialCurrencyCode NullableString) *PayStubDistributionBreakdown {
	this := PayStubDistributionBreakdown{}
	this.AccountName = accountName
	this.BankName = bankName
	this.CurrentAmount = currentAmount
	this.IsoCurrencyCode = isoCurrencyCode
	this.Mask = mask
	this.Type = type_
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewPayStubDistributionBreakdownWithDefaults instantiates a new PayStubDistributionBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubDistributionBreakdownWithDefaults() *PayStubDistributionBreakdown {
	this := PayStubDistributionBreakdown{}
	return &this
}

// GetAccountName returns the AccountName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDistributionBreakdown) GetAccountName() string {
	if o == nil || o.AccountName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountName.Get()
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountName.Get(), o.AccountName.IsSet()
}

// SetAccountName sets field value
func (o *PayStubDistributionBreakdown) SetAccountName(v string) {
	o.AccountName.Set(&v)
}

// GetBankName returns the BankName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDistributionBreakdown) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// SetBankName sets field value
func (o *PayStubDistributionBreakdown) SetBankName(v string) {
	o.BankName.Set(&v)
}

// GetCurrentAmount returns the CurrentAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubDistributionBreakdown) GetCurrentAmount() float64 {
	if o == nil || o.CurrentAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.CurrentAmount.Get()
}

// GetCurrentAmountOk returns a tuple with the CurrentAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetCurrentAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentAmount.Get(), o.CurrentAmount.IsSet()
}

// SetCurrentAmount sets field value
func (o *PayStubDistributionBreakdown) SetCurrentAmount(v float64) {
	o.CurrentAmount.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDistributionBreakdown) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *PayStubDistributionBreakdown) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDistributionBreakdown) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *PayStubDistributionBreakdown) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDistributionBreakdown) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *PayStubDistributionBreakdown) SetType(v string) {
	o.Type.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubDistributionBreakdown) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubDistributionBreakdown) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *PayStubDistributionBreakdown) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o PayStubDistributionBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_name"] = o.AccountName.Get()
	}
	if true {
		toSerialize["bank_name"] = o.BankName.Get()
	}
	if true {
		toSerialize["current_amount"] = o.CurrentAmount.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStubDistributionBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	varPayStubDistributionBreakdown := _PayStubDistributionBreakdown{}

	if err = json.Unmarshal(bytes, &varPayStubDistributionBreakdown); err == nil {
		*o = PayStubDistributionBreakdown(varPayStubDistributionBreakdown)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_name")
		delete(additionalProperties, "bank_name")
		delete(additionalProperties, "current_amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "type")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStubDistributionBreakdown struct {
	value *PayStubDistributionBreakdown
	isSet bool
}

func (v NullablePayStubDistributionBreakdown) Get() *PayStubDistributionBreakdown {
	return v.value
}

func (v *NullablePayStubDistributionBreakdown) Set(val *PayStubDistributionBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStubDistributionBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStubDistributionBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStubDistributionBreakdown(val *PayStubDistributionBreakdown) *NullablePayStubDistributionBreakdown {
	return &NullablePayStubDistributionBreakdown{value: val, isSet: true}
}

func (v NullablePayStubDistributionBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStubDistributionBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


