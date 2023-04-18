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

// CreditPayrollIncomeRiskSignalsGetRequest CreditPayrollIncomeRiskSignalsGetRequest defines the request schema for `/beta/credit/payroll_income/risk_signals/get`
type CreditPayrollIncomeRiskSignalsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The user token associated with the User data is being requested for.
	UserToken *string `json:"user_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreditPayrollIncomeRiskSignalsGetRequest CreditPayrollIncomeRiskSignalsGetRequest

// NewCreditPayrollIncomeRiskSignalsGetRequest instantiates a new CreditPayrollIncomeRiskSignalsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPayrollIncomeRiskSignalsGetRequest() *CreditPayrollIncomeRiskSignalsGetRequest {
	this := CreditPayrollIncomeRiskSignalsGetRequest{}
	return &this
}

// NewCreditPayrollIncomeRiskSignalsGetRequestWithDefaults instantiates a new CreditPayrollIncomeRiskSignalsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPayrollIncomeRiskSignalsGetRequestWithDefaults() *CreditPayrollIncomeRiskSignalsGetRequest {
	this := CreditPayrollIncomeRiskSignalsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) GetUserToken() string {
	if o == nil || o.UserToken == nil {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) GetUserTokenOk() (*string, bool) {
	if o == nil || o.UserToken == nil {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) HasUserToken() bool {
	if o != nil && o.UserToken != nil {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *CreditPayrollIncomeRiskSignalsGetRequest) SetUserToken(v string) {
	o.UserToken = &v
}

func (o CreditPayrollIncomeRiskSignalsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.UserToken != nil {
		toSerialize["user_token"] = o.UserToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditPayrollIncomeRiskSignalsGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreditPayrollIncomeRiskSignalsGetRequest := _CreditPayrollIncomeRiskSignalsGetRequest{}

	if err = json.Unmarshal(bytes, &varCreditPayrollIncomeRiskSignalsGetRequest); err == nil {
		*o = CreditPayrollIncomeRiskSignalsGetRequest(varCreditPayrollIncomeRiskSignalsGetRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "user_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditPayrollIncomeRiskSignalsGetRequest struct {
	value *CreditPayrollIncomeRiskSignalsGetRequest
	isSet bool
}

func (v NullableCreditPayrollIncomeRiskSignalsGetRequest) Get() *CreditPayrollIncomeRiskSignalsGetRequest {
	return v.value
}

func (v *NullableCreditPayrollIncomeRiskSignalsGetRequest) Set(val *CreditPayrollIncomeRiskSignalsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPayrollIncomeRiskSignalsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPayrollIncomeRiskSignalsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPayrollIncomeRiskSignalsGetRequest(val *CreditPayrollIncomeRiskSignalsGetRequest) *NullableCreditPayrollIncomeRiskSignalsGetRequest {
	return &NullableCreditPayrollIncomeRiskSignalsGetRequest{value: val, isSet: true}
}

func (v NullableCreditPayrollIncomeRiskSignalsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPayrollIncomeRiskSignalsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


