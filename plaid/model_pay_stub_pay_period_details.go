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

// PayStubPayPeriodDetails Details about the pay period.
type PayStubPayPeriodDetails struct {
	// The amount of the paycheck.
	PayAmount NullableFloat64 `json:"pay_amount"`
	DistributionBreakdown []PayStubDistributionBreakdown `json:"distribution_breakdown"`
	// The date on which the pay period ended, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (\"yyyy-mm-dd\").
	EndDate NullableString `json:"end_date"`
	// Total earnings before tax/deductions.
	GrossEarnings NullableFloat64 `json:"gross_earnings"`
	// The ISO-4217 currency code of the net pay. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The date on which the pay stub was issued, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (\"yyyy-mm-dd\").
	PayDate NullableString `json:"pay_date"`
	// The frequency at which an individual is paid.
	PayFrequency NullableString `json:"pay_frequency"`
	PayBasis *CreditPayStubPayBasisType `json:"pay_basis,omitempty"`
	// The date on which the pay period started, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (\"yyyy-mm-dd\").
	StartDate NullableString `json:"start_date"`
	// The unofficial currency code associated with the net pay. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	AdditionalProperties map[string]interface{}
}

type _PayStubPayPeriodDetails PayStubPayPeriodDetails

// NewPayStubPayPeriodDetails instantiates a new PayStubPayPeriodDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayStubPayPeriodDetails(payAmount NullableFloat64, distributionBreakdown []PayStubDistributionBreakdown, endDate NullableString, grossEarnings NullableFloat64, isoCurrencyCode NullableString, payDate NullableString, payFrequency NullableString, startDate NullableString, unofficialCurrencyCode NullableString) *PayStubPayPeriodDetails {
	this := PayStubPayPeriodDetails{}
	this.PayAmount = payAmount
	this.DistributionBreakdown = distributionBreakdown
	this.EndDate = endDate
	this.GrossEarnings = grossEarnings
	this.IsoCurrencyCode = isoCurrencyCode
	this.PayDate = payDate
	this.PayFrequency = payFrequency
	this.StartDate = startDate
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	return &this
}

// NewPayStubPayPeriodDetailsWithDefaults instantiates a new PayStubPayPeriodDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayStubPayPeriodDetailsWithDefaults() *PayStubPayPeriodDetails {
	this := PayStubPayPeriodDetails{}
	return &this
}

// GetPayAmount returns the PayAmount field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubPayPeriodDetails) GetPayAmount() float64 {
	if o == nil || o.PayAmount.Get() == nil {
		var ret float64
		return ret
	}

	return *o.PayAmount.Get()
}

// GetPayAmountOk returns a tuple with the PayAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetPayAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayAmount.Get(), o.PayAmount.IsSet()
}

// SetPayAmount sets field value
func (o *PayStubPayPeriodDetails) SetPayAmount(v float64) {
	o.PayAmount.Set(&v)
}

// GetDistributionBreakdown returns the DistributionBreakdown field value
func (o *PayStubPayPeriodDetails) GetDistributionBreakdown() []PayStubDistributionBreakdown {
	if o == nil {
		var ret []PayStubDistributionBreakdown
		return ret
	}

	return o.DistributionBreakdown
}

// GetDistributionBreakdownOk returns a tuple with the DistributionBreakdown field value
// and a boolean to check if the value has been set.
func (o *PayStubPayPeriodDetails) GetDistributionBreakdownOk() (*[]PayStubDistributionBreakdown, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DistributionBreakdown, true
}

// SetDistributionBreakdown sets field value
func (o *PayStubPayPeriodDetails) SetDistributionBreakdown(v []PayStubDistributionBreakdown) {
	o.DistributionBreakdown = v
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubPayPeriodDetails) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *PayStubPayPeriodDetails) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetGrossEarnings returns the GrossEarnings field value
// If the value is explicit nil, the zero value for float64 will be returned
func (o *PayStubPayPeriodDetails) GetGrossEarnings() float64 {
	if o == nil || o.GrossEarnings.Get() == nil {
		var ret float64
		return ret
	}

	return *o.GrossEarnings.Get()
}

// GetGrossEarningsOk returns a tuple with the GrossEarnings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetGrossEarningsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GrossEarnings.Get(), o.GrossEarnings.IsSet()
}

// SetGrossEarnings sets field value
func (o *PayStubPayPeriodDetails) SetGrossEarnings(v float64) {
	o.GrossEarnings.Set(&v)
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubPayPeriodDetails) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *PayStubPayPeriodDetails) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetPayDate returns the PayDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubPayPeriodDetails) GetPayDate() string {
	if o == nil || o.PayDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayDate.Get()
}

// GetPayDateOk returns a tuple with the PayDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetPayDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayDate.Get(), o.PayDate.IsSet()
}

// SetPayDate sets field value
func (o *PayStubPayPeriodDetails) SetPayDate(v string) {
	o.PayDate.Set(&v)
}

// GetPayFrequency returns the PayFrequency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubPayPeriodDetails) GetPayFrequency() string {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetPayFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// SetPayFrequency sets field value
func (o *PayStubPayPeriodDetails) SetPayFrequency(v string) {
	o.PayFrequency.Set(&v)
}

// GetPayBasis returns the PayBasis field value if set, zero value otherwise.
func (o *PayStubPayPeriodDetails) GetPayBasis() CreditPayStubPayBasisType {
	if o == nil || o.PayBasis == nil {
		var ret CreditPayStubPayBasisType
		return ret
	}
	return *o.PayBasis
}

// GetPayBasisOk returns a tuple with the PayBasis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayStubPayPeriodDetails) GetPayBasisOk() (*CreditPayStubPayBasisType, bool) {
	if o == nil || o.PayBasis == nil {
		return nil, false
	}
	return o.PayBasis, true
}

// HasPayBasis returns a boolean if a field has been set.
func (o *PayStubPayPeriodDetails) HasPayBasis() bool {
	if o != nil && o.PayBasis != nil {
		return true
	}

	return false
}

// SetPayBasis gets a reference to the given CreditPayStubPayBasisType and assigns it to the PayBasis field.
func (o *PayStubPayPeriodDetails) SetPayBasis(v CreditPayStubPayBasisType) {
	o.PayBasis = &v
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubPayPeriodDetails) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *PayStubPayPeriodDetails) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayStubPayPeriodDetails) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayStubPayPeriodDetails) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *PayStubPayPeriodDetails) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

func (o PayStubPayPeriodDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pay_amount"] = o.PayAmount.Get()
	}
	if true {
		toSerialize["distribution_breakdown"] = o.DistributionBreakdown
	}
	if true {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if true {
		toSerialize["gross_earnings"] = o.GrossEarnings.Get()
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["pay_date"] = o.PayDate.Get()
	}
	if true {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}
	if o.PayBasis != nil {
		toSerialize["pay_basis"] = o.PayBasis
	}
	if true {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PayStubPayPeriodDetails) UnmarshalJSON(bytes []byte) (err error) {
	varPayStubPayPeriodDetails := _PayStubPayPeriodDetails{}

	if err = json.Unmarshal(bytes, &varPayStubPayPeriodDetails); err == nil {
		*o = PayStubPayPeriodDetails(varPayStubPayPeriodDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pay_amount")
		delete(additionalProperties, "distribution_breakdown")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "gross_earnings")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "pay_date")
		delete(additionalProperties, "pay_frequency")
		delete(additionalProperties, "pay_basis")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "unofficial_currency_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayStubPayPeriodDetails struct {
	value *PayStubPayPeriodDetails
	isSet bool
}

func (v NullablePayStubPayPeriodDetails) Get() *PayStubPayPeriodDetails {
	return v.value
}

func (v *NullablePayStubPayPeriodDetails) Set(val *PayStubPayPeriodDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePayStubPayPeriodDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePayStubPayPeriodDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayStubPayPeriodDetails(val *PayStubPayPeriodDetails) *NullablePayStubPayPeriodDetails {
	return &NullablePayStubPayPeriodDetails{value: val, isSet: true}
}

func (v NullablePayStubPayPeriodDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayStubPayPeriodDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


