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

// LinkTokenCreateRequestUserStatedIncomeSource Specifies user stated income sources for the Income product
type LinkTokenCreateRequestUserStatedIncomeSource struct {
	// The employer corresponding to an income source specified by the user
	Employer *string `json:"employer,omitempty"`
	Category *UserStatedIncomeSourceCategory `json:"category,omitempty"`
	// The income amount paid per cycle for a specified income source
	PayPerCycle *float64 `json:"pay_per_cycle,omitempty"`
	// The income amount paid annually for a specified income source
	PayAnnual *float64 `json:"pay_annual,omitempty"`
	PayType *UserStatedIncomeSourcePayType `json:"pay_type,omitempty"`
	PayFrequency *UserStatedIncomeSourceFrequency `json:"pay_frequency,omitempty"`
}

// NewLinkTokenCreateRequestUserStatedIncomeSource instantiates a new LinkTokenCreateRequestUserStatedIncomeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestUserStatedIncomeSource() *LinkTokenCreateRequestUserStatedIncomeSource {
	this := LinkTokenCreateRequestUserStatedIncomeSource{}
	return &this
}

// NewLinkTokenCreateRequestUserStatedIncomeSourceWithDefaults instantiates a new LinkTokenCreateRequestUserStatedIncomeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestUserStatedIncomeSourceWithDefaults() *LinkTokenCreateRequestUserStatedIncomeSource {
	this := LinkTokenCreateRequestUserStatedIncomeSource{}
	return &this
}

// GetEmployer returns the Employer field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetEmployer() string {
	if o == nil || o.Employer == nil {
		var ret string
		return ret
	}
	return *o.Employer
}

// GetEmployerOk returns a tuple with the Employer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetEmployerOk() (*string, bool) {
	if o == nil || o.Employer == nil {
		return nil, false
	}
	return o.Employer, true
}

// HasEmployer returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) HasEmployer() bool {
	if o != nil && o.Employer != nil {
		return true
	}

	return false
}

// SetEmployer gets a reference to the given string and assigns it to the Employer field.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) SetEmployer(v string) {
	o.Employer = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetCategory() UserStatedIncomeSourceCategory {
	if o == nil || o.Category == nil {
		var ret UserStatedIncomeSourceCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetCategoryOk() (*UserStatedIncomeSourceCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given UserStatedIncomeSourceCategory and assigns it to the Category field.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) SetCategory(v UserStatedIncomeSourceCategory) {
	o.Category = &v
}

// GetPayPerCycle returns the PayPerCycle field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayPerCycle() float64 {
	if o == nil || o.PayPerCycle == nil {
		var ret float64
		return ret
	}
	return *o.PayPerCycle
}

// GetPayPerCycleOk returns a tuple with the PayPerCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayPerCycleOk() (*float64, bool) {
	if o == nil || o.PayPerCycle == nil {
		return nil, false
	}
	return o.PayPerCycle, true
}

// HasPayPerCycle returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) HasPayPerCycle() bool {
	if o != nil && o.PayPerCycle != nil {
		return true
	}

	return false
}

// SetPayPerCycle gets a reference to the given float64 and assigns it to the PayPerCycle field.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) SetPayPerCycle(v float64) {
	o.PayPerCycle = &v
}

// GetPayAnnual returns the PayAnnual field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayAnnual() float64 {
	if o == nil || o.PayAnnual == nil {
		var ret float64
		return ret
	}
	return *o.PayAnnual
}

// GetPayAnnualOk returns a tuple with the PayAnnual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayAnnualOk() (*float64, bool) {
	if o == nil || o.PayAnnual == nil {
		return nil, false
	}
	return o.PayAnnual, true
}

// HasPayAnnual returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) HasPayAnnual() bool {
	if o != nil && o.PayAnnual != nil {
		return true
	}

	return false
}

// SetPayAnnual gets a reference to the given float64 and assigns it to the PayAnnual field.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) SetPayAnnual(v float64) {
	o.PayAnnual = &v
}

// GetPayType returns the PayType field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayType() UserStatedIncomeSourcePayType {
	if o == nil || o.PayType == nil {
		var ret UserStatedIncomeSourcePayType
		return ret
	}
	return *o.PayType
}

// GetPayTypeOk returns a tuple with the PayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayTypeOk() (*UserStatedIncomeSourcePayType, bool) {
	if o == nil || o.PayType == nil {
		return nil, false
	}
	return o.PayType, true
}

// HasPayType returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) HasPayType() bool {
	if o != nil && o.PayType != nil {
		return true
	}

	return false
}

// SetPayType gets a reference to the given UserStatedIncomeSourcePayType and assigns it to the PayType field.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) SetPayType(v UserStatedIncomeSourcePayType) {
	o.PayType = &v
}

// GetPayFrequency returns the PayFrequency field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayFrequency() UserStatedIncomeSourceFrequency {
	if o == nil || o.PayFrequency == nil {
		var ret UserStatedIncomeSourceFrequency
		return ret
	}
	return *o.PayFrequency
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) GetPayFrequencyOk() (*UserStatedIncomeSourceFrequency, bool) {
	if o == nil || o.PayFrequency == nil {
		return nil, false
	}
	return o.PayFrequency, true
}

// HasPayFrequency returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) HasPayFrequency() bool {
	if o != nil && o.PayFrequency != nil {
		return true
	}

	return false
}

// SetPayFrequency gets a reference to the given UserStatedIncomeSourceFrequency and assigns it to the PayFrequency field.
func (o *LinkTokenCreateRequestUserStatedIncomeSource) SetPayFrequency(v UserStatedIncomeSourceFrequency) {
	o.PayFrequency = &v
}

func (o LinkTokenCreateRequestUserStatedIncomeSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Employer != nil {
		toSerialize["employer"] = o.Employer
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.PayPerCycle != nil {
		toSerialize["pay_per_cycle"] = o.PayPerCycle
	}
	if o.PayAnnual != nil {
		toSerialize["pay_annual"] = o.PayAnnual
	}
	if o.PayType != nil {
		toSerialize["pay_type"] = o.PayType
	}
	if o.PayFrequency != nil {
		toSerialize["pay_frequency"] = o.PayFrequency
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestUserStatedIncomeSource struct {
	value *LinkTokenCreateRequestUserStatedIncomeSource
	isSet bool
}

func (v NullableLinkTokenCreateRequestUserStatedIncomeSource) Get() *LinkTokenCreateRequestUserStatedIncomeSource {
	return v.value
}

func (v *NullableLinkTokenCreateRequestUserStatedIncomeSource) Set(val *LinkTokenCreateRequestUserStatedIncomeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestUserStatedIncomeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestUserStatedIncomeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestUserStatedIncomeSource(val *LinkTokenCreateRequestUserStatedIncomeSource) *NullableLinkTokenCreateRequestUserStatedIncomeSource {
	return &NullableLinkTokenCreateRequestUserStatedIncomeSource{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestUserStatedIncomeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestUserStatedIncomeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


