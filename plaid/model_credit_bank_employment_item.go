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
	"time"
)

// CreditBankEmploymentItem The details and metadata for an end user's Item.
type CreditBankEmploymentItem struct {
	// The unique identifier for the Item.
	ItemId string `json:"item_id"`
	// The time when this Item's data was last retrieved from the financial institution, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (e.g. \"2018-04-12T03:32:11Z\").
	LastUpdatedTime time.Time `json:"last_updated_time"`
	// The unique identifier of the institution associated with the Item.
	InstitutionId string `json:"institution_id"`
	// The name of the institution associated with the Item.
	InstitutionName string `json:"institution_name"`
	// The bank employment information for this Item. Each entry in the array is a different employer found.
	BankEmployments []CreditBankEmployment `json:"bank_employments"`
	// The Item's accounts that have Bank Employment data.
	BankEmploymentAccounts []CreditBankIncomeAccount `json:"bank_employment_accounts"`
}

// NewCreditBankEmploymentItem instantiates a new CreditBankEmploymentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditBankEmploymentItem(itemId string, lastUpdatedTime time.Time, institutionId string, institutionName string, bankEmployments []CreditBankEmployment, bankEmploymentAccounts []CreditBankIncomeAccount) *CreditBankEmploymentItem {
	this := CreditBankEmploymentItem{}
	this.ItemId = itemId
	this.LastUpdatedTime = lastUpdatedTime
	this.InstitutionId = institutionId
	this.InstitutionName = institutionName
	this.BankEmployments = bankEmployments
	this.BankEmploymentAccounts = bankEmploymentAccounts
	return &this
}

// NewCreditBankEmploymentItemWithDefaults instantiates a new CreditBankEmploymentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditBankEmploymentItemWithDefaults() *CreditBankEmploymentItem {
	this := CreditBankEmploymentItem{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *CreditBankEmploymentItem) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentItem) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *CreditBankEmploymentItem) SetItemId(v string) {
	o.ItemId = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value
func (o *CreditBankEmploymentItem) GetLastUpdatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentItem) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastUpdatedTime, true
}

// SetLastUpdatedTime sets field value
func (o *CreditBankEmploymentItem) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = v
}

// GetInstitutionId returns the InstitutionId field value
func (o *CreditBankEmploymentItem) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentItem) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *CreditBankEmploymentItem) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetInstitutionName returns the InstitutionName field value
func (o *CreditBankEmploymentItem) GetInstitutionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionName
}

// GetInstitutionNameOk returns a tuple with the InstitutionName field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentItem) GetInstitutionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionName, true
}

// SetInstitutionName sets field value
func (o *CreditBankEmploymentItem) SetInstitutionName(v string) {
	o.InstitutionName = v
}

// GetBankEmployments returns the BankEmployments field value
func (o *CreditBankEmploymentItem) GetBankEmployments() []CreditBankEmployment {
	if o == nil {
		var ret []CreditBankEmployment
		return ret
	}

	return o.BankEmployments
}

// GetBankEmploymentsOk returns a tuple with the BankEmployments field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentItem) GetBankEmploymentsOk() (*[]CreditBankEmployment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankEmployments, true
}

// SetBankEmployments sets field value
func (o *CreditBankEmploymentItem) SetBankEmployments(v []CreditBankEmployment) {
	o.BankEmployments = v
}

// GetBankEmploymentAccounts returns the BankEmploymentAccounts field value
func (o *CreditBankEmploymentItem) GetBankEmploymentAccounts() []CreditBankIncomeAccount {
	if o == nil {
		var ret []CreditBankIncomeAccount
		return ret
	}

	return o.BankEmploymentAccounts
}

// GetBankEmploymentAccountsOk returns a tuple with the BankEmploymentAccounts field value
// and a boolean to check if the value has been set.
func (o *CreditBankEmploymentItem) GetBankEmploymentAccountsOk() (*[]CreditBankIncomeAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankEmploymentAccounts, true
}

// SetBankEmploymentAccounts sets field value
func (o *CreditBankEmploymentItem) SetBankEmploymentAccounts(v []CreditBankIncomeAccount) {
	o.BankEmploymentAccounts = v
}

func (o CreditBankEmploymentItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["last_updated_time"] = o.LastUpdatedTime
	}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["institution_name"] = o.InstitutionName
	}
	if true {
		toSerialize["bank_employments"] = o.BankEmployments
	}
	if true {
		toSerialize["bank_employment_accounts"] = o.BankEmploymentAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableCreditBankEmploymentItem struct {
	value *CreditBankEmploymentItem
	isSet bool
}

func (v NullableCreditBankEmploymentItem) Get() *CreditBankEmploymentItem {
	return v.value
}

func (v *NullableCreditBankEmploymentItem) Set(val *CreditBankEmploymentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditBankEmploymentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditBankEmploymentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditBankEmploymentItem(val *CreditBankEmploymentItem) *NullableCreditBankEmploymentItem {
	return &NullableCreditBankEmploymentItem{value: val, isSet: true}
}

func (v NullableCreditBankEmploymentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditBankEmploymentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


