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

// TransactionsRulesCreateRequest TransactionsRulesCreateRequest defines the request schema for `beta/transactions/rules/v1/create`
type TransactionsRulesCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Personal finance detailed category.  See the [`taxonomy csv file`](https://plaid.com/documents/transactions-personal-finance-category-taxonomy.csv) for a full list of personal finance categories. 
	PersonalFinanceCategory string `json:"personal_finance_category"`
	RuleDetails TransactionsRuleDetails `json:"rule_details"`
}

// NewTransactionsRulesCreateRequest instantiates a new TransactionsRulesCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRulesCreateRequest(accessToken string, personalFinanceCategory string, ruleDetails TransactionsRuleDetails) *TransactionsRulesCreateRequest {
	this := TransactionsRulesCreateRequest{}
	this.AccessToken = accessToken
	this.PersonalFinanceCategory = personalFinanceCategory
	this.RuleDetails = ruleDetails
	return &this
}

// NewTransactionsRulesCreateRequestWithDefaults instantiates a new TransactionsRulesCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRulesCreateRequestWithDefaults() *TransactionsRulesCreateRequest {
	this := TransactionsRulesCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *TransactionsRulesCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *TransactionsRulesCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *TransactionsRulesCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetAccessToken returns the AccessToken field value
func (o *TransactionsRulesCreateRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TransactionsRulesCreateRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *TransactionsRulesCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *TransactionsRulesCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *TransactionsRulesCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value
func (o *TransactionsRulesCreateRequest) GetPersonalFinanceCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PersonalFinanceCategory
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateRequest) GetPersonalFinanceCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PersonalFinanceCategory, true
}

// SetPersonalFinanceCategory sets field value
func (o *TransactionsRulesCreateRequest) SetPersonalFinanceCategory(v string) {
	o.PersonalFinanceCategory = v
}

// GetRuleDetails returns the RuleDetails field value
func (o *TransactionsRulesCreateRequest) GetRuleDetails() TransactionsRuleDetails {
	if o == nil {
		var ret TransactionsRuleDetails
		return ret
	}

	return o.RuleDetails
}

// GetRuleDetailsOk returns a tuple with the RuleDetails field value
// and a boolean to check if the value has been set.
func (o *TransactionsRulesCreateRequest) GetRuleDetailsOk() (*TransactionsRuleDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RuleDetails, true
}

// SetRuleDetails sets field value
func (o *TransactionsRulesCreateRequest) SetRuleDetails(v TransactionsRuleDetails) {
	o.RuleDetails = v
}

func (o TransactionsRulesCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory
	}
	if true {
		toSerialize["rule_details"] = o.RuleDetails
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsRulesCreateRequest struct {
	value *TransactionsRulesCreateRequest
	isSet bool
}

func (v NullableTransactionsRulesCreateRequest) Get() *TransactionsRulesCreateRequest {
	return v.value
}

func (v *NullableTransactionsRulesCreateRequest) Set(val *TransactionsRulesCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRulesCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRulesCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRulesCreateRequest(val *TransactionsRulesCreateRequest) *NullableTransactionsRulesCreateRequest {
	return &NullableTransactionsRulesCreateRequest{value: val, isSet: true}
}

func (v NullableTransactionsRulesCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRulesCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


