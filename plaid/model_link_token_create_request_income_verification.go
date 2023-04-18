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

// LinkTokenCreateRequestIncomeVerification Specifies options for initializing Link for use with the Income product. This field is required if `income_verification` is included in the `products` array.
type LinkTokenCreateRequestIncomeVerification struct {
	// The `income_verification_id` of the verification instance, as provided by `/income/verification/create`.
	IncomeVerificationId *string `json:"income_verification_id,omitempty"`
	// The `asset_report_id` of an asset report associated with the user, as provided by `/asset_report/create`. Providing an `asset_report_id` is optional and can be used to verify the user through a streamlined flow. If provided, the bank linking flow will be skipped.
	AssetReportId *string `json:"asset_report_id,omitempty"`
	// The ID of a precheck created with `/income/verification/precheck`. Will be used to improve conversion of the income verification flow by streamlining the Link interface presented to the end user.
	PrecheckId *string `json:"precheck_id,omitempty"`
	// An array of access tokens corresponding to Items that a user has previously connected with. Data from these institutions will be cross-referenced with document data received during the Document Income flow to help verify that the uploaded documents are accurate. If the `transactions` product was not initialized for these Items during link, it will be initialized after this Link session.  This field should only be used with the `payroll` income source type.
	AccessTokens *[]string `json:"access_tokens,omitempty"`
	// The types of source income data that users will be permitted to share. Options include `bank` and `payroll`. Currently you can only specify one of these options.
	IncomeSourceTypes *[]IncomeVerificationSourceType `json:"income_source_types,omitempty"`
	BankIncome *LinkTokenCreateRequestIncomeVerificationBankIncome `json:"bank_income,omitempty"`
	PayrollIncome *LinkTokenCreateRequestIncomeVerificationPayrollIncome `json:"payroll_income,omitempty"`
	// A list of user stated income sources
	StatedIncomeSources *[]LinkTokenCreateRequestUserStatedIncomeSource `json:"stated_income_sources,omitempty"`
}

// NewLinkTokenCreateRequestIncomeVerification instantiates a new LinkTokenCreateRequestIncomeVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestIncomeVerification() *LinkTokenCreateRequestIncomeVerification {
	this := LinkTokenCreateRequestIncomeVerification{}
	return &this
}

// NewLinkTokenCreateRequestIncomeVerificationWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestIncomeVerificationWithDefaults() *LinkTokenCreateRequestIncomeVerification {
	this := LinkTokenCreateRequestIncomeVerification{}
	return &this
}

// GetIncomeVerificationId returns the IncomeVerificationId field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationId() string {
	if o == nil || o.IncomeVerificationId == nil {
		var ret string
		return ret
	}
	return *o.IncomeVerificationId
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil || o.IncomeVerificationId == nil {
		return nil, false
	}
	return o.IncomeVerificationId, true
}

// HasIncomeVerificationId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasIncomeVerificationId() bool {
	if o != nil && o.IncomeVerificationId != nil {
		return true
	}

	return false
}

// SetIncomeVerificationId gets a reference to the given string and assigns it to the IncomeVerificationId field.
func (o *LinkTokenCreateRequestIncomeVerification) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId = &v
}

// GetAssetReportId returns the AssetReportId field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportId() string {
	if o == nil || o.AssetReportId == nil {
		var ret string
		return ret
	}
	return *o.AssetReportId
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportIdOk() (*string, bool) {
	if o == nil || o.AssetReportId == nil {
		return nil, false
	}
	return o.AssetReportId, true
}

// HasAssetReportId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasAssetReportId() bool {
	if o != nil && o.AssetReportId != nil {
		return true
	}

	return false
}

// SetAssetReportId gets a reference to the given string and assigns it to the AssetReportId field.
func (o *LinkTokenCreateRequestIncomeVerification) SetAssetReportId(v string) {
	o.AssetReportId = &v
}

// GetPrecheckId returns the PrecheckId field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetPrecheckId() string {
	if o == nil || o.PrecheckId == nil {
		var ret string
		return ret
	}
	return *o.PrecheckId
}

// GetPrecheckIdOk returns a tuple with the PrecheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetPrecheckIdOk() (*string, bool) {
	if o == nil || o.PrecheckId == nil {
		return nil, false
	}
	return o.PrecheckId, true
}

// HasPrecheckId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasPrecheckId() bool {
	if o != nil && o.PrecheckId != nil {
		return true
	}

	return false
}

// SetPrecheckId gets a reference to the given string and assigns it to the PrecheckId field.
func (o *LinkTokenCreateRequestIncomeVerification) SetPrecheckId(v string) {
	o.PrecheckId = &v
}

// GetAccessTokens returns the AccessTokens field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetAccessTokens() []string {
	if o == nil || o.AccessTokens == nil {
		var ret []string
		return ret
	}
	return *o.AccessTokens
}

// GetAccessTokensOk returns a tuple with the AccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetAccessTokensOk() (*[]string, bool) {
	if o == nil || o.AccessTokens == nil {
		return nil, false
	}
	return o.AccessTokens, true
}

// HasAccessTokens returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasAccessTokens() bool {
	if o != nil && o.AccessTokens != nil {
		return true
	}

	return false
}

// SetAccessTokens gets a reference to the given []string and assigns it to the AccessTokens field.
func (o *LinkTokenCreateRequestIncomeVerification) SetAccessTokens(v []string) {
	o.AccessTokens = &v
}

// GetIncomeSourceTypes returns the IncomeSourceTypes field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeSourceTypes() []IncomeVerificationSourceType {
	if o == nil || o.IncomeSourceTypes == nil {
		var ret []IncomeVerificationSourceType
		return ret
	}
	return *o.IncomeSourceTypes
}

// GetIncomeSourceTypesOk returns a tuple with the IncomeSourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeSourceTypesOk() (*[]IncomeVerificationSourceType, bool) {
	if o == nil || o.IncomeSourceTypes == nil {
		return nil, false
	}
	return o.IncomeSourceTypes, true
}

// HasIncomeSourceTypes returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasIncomeSourceTypes() bool {
	if o != nil && o.IncomeSourceTypes != nil {
		return true
	}

	return false
}

// SetIncomeSourceTypes gets a reference to the given []IncomeVerificationSourceType and assigns it to the IncomeSourceTypes field.
func (o *LinkTokenCreateRequestIncomeVerification) SetIncomeSourceTypes(v []IncomeVerificationSourceType) {
	o.IncomeSourceTypes = &v
}

// GetBankIncome returns the BankIncome field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetBankIncome() LinkTokenCreateRequestIncomeVerificationBankIncome {
	if o == nil || o.BankIncome == nil {
		var ret LinkTokenCreateRequestIncomeVerificationBankIncome
		return ret
	}
	return *o.BankIncome
}

// GetBankIncomeOk returns a tuple with the BankIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetBankIncomeOk() (*LinkTokenCreateRequestIncomeVerificationBankIncome, bool) {
	if o == nil || o.BankIncome == nil {
		return nil, false
	}
	return o.BankIncome, true
}

// HasBankIncome returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasBankIncome() bool {
	if o != nil && o.BankIncome != nil {
		return true
	}

	return false
}

// SetBankIncome gets a reference to the given LinkTokenCreateRequestIncomeVerificationBankIncome and assigns it to the BankIncome field.
func (o *LinkTokenCreateRequestIncomeVerification) SetBankIncome(v LinkTokenCreateRequestIncomeVerificationBankIncome) {
	o.BankIncome = &v
}

// GetPayrollIncome returns the PayrollIncome field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetPayrollIncome() LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	if o == nil || o.PayrollIncome == nil {
		var ret LinkTokenCreateRequestIncomeVerificationPayrollIncome
		return ret
	}
	return *o.PayrollIncome
}

// GetPayrollIncomeOk returns a tuple with the PayrollIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetPayrollIncomeOk() (*LinkTokenCreateRequestIncomeVerificationPayrollIncome, bool) {
	if o == nil || o.PayrollIncome == nil {
		return nil, false
	}
	return o.PayrollIncome, true
}

// HasPayrollIncome returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasPayrollIncome() bool {
	if o != nil && o.PayrollIncome != nil {
		return true
	}

	return false
}

// SetPayrollIncome gets a reference to the given LinkTokenCreateRequestIncomeVerificationPayrollIncome and assigns it to the PayrollIncome field.
func (o *LinkTokenCreateRequestIncomeVerification) SetPayrollIncome(v LinkTokenCreateRequestIncomeVerificationPayrollIncome) {
	o.PayrollIncome = &v
}

// GetStatedIncomeSources returns the StatedIncomeSources field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerification) GetStatedIncomeSources() []LinkTokenCreateRequestUserStatedIncomeSource {
	if o == nil || o.StatedIncomeSources == nil {
		var ret []LinkTokenCreateRequestUserStatedIncomeSource
		return ret
	}
	return *o.StatedIncomeSources
}

// GetStatedIncomeSourcesOk returns a tuple with the StatedIncomeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetStatedIncomeSourcesOk() (*[]LinkTokenCreateRequestUserStatedIncomeSource, bool) {
	if o == nil || o.StatedIncomeSources == nil {
		return nil, false
	}
	return o.StatedIncomeSources, true
}

// HasStatedIncomeSources returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasStatedIncomeSources() bool {
	if o != nil && o.StatedIncomeSources != nil {
		return true
	}

	return false
}

// SetStatedIncomeSources gets a reference to the given []LinkTokenCreateRequestUserStatedIncomeSource and assigns it to the StatedIncomeSources field.
func (o *LinkTokenCreateRequestIncomeVerification) SetStatedIncomeSources(v []LinkTokenCreateRequestUserStatedIncomeSource) {
	o.StatedIncomeSources = &v
}

func (o LinkTokenCreateRequestIncomeVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncomeVerificationId != nil {
		toSerialize["income_verification_id"] = o.IncomeVerificationId
	}
	if o.AssetReportId != nil {
		toSerialize["asset_report_id"] = o.AssetReportId
	}
	if o.PrecheckId != nil {
		toSerialize["precheck_id"] = o.PrecheckId
	}
	if o.AccessTokens != nil {
		toSerialize["access_tokens"] = o.AccessTokens
	}
	if o.IncomeSourceTypes != nil {
		toSerialize["income_source_types"] = o.IncomeSourceTypes
	}
	if o.BankIncome != nil {
		toSerialize["bank_income"] = o.BankIncome
	}
	if o.PayrollIncome != nil {
		toSerialize["payroll_income"] = o.PayrollIncome
	}
	if o.StatedIncomeSources != nil {
		toSerialize["stated_income_sources"] = o.StatedIncomeSources
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestIncomeVerification struct {
	value *LinkTokenCreateRequestIncomeVerification
	isSet bool
}

func (v NullableLinkTokenCreateRequestIncomeVerification) Get() *LinkTokenCreateRequestIncomeVerification {
	return v.value
}

func (v *NullableLinkTokenCreateRequestIncomeVerification) Set(val *LinkTokenCreateRequestIncomeVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestIncomeVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestIncomeVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestIncomeVerification(val *LinkTokenCreateRequestIncomeVerification) *NullableLinkTokenCreateRequestIncomeVerification {
	return &NullableLinkTokenCreateRequestIncomeVerification{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestIncomeVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestIncomeVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


