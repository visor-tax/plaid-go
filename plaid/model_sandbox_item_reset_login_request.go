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

// SandboxItemResetLoginRequest SandboxItemResetLoginRequest defines the request schema for `/sandbox/item/reset_login`
type SandboxItemResetLoginRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
}

// NewSandboxItemResetLoginRequest instantiates a new SandboxItemResetLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemResetLoginRequest(accessToken string) *SandboxItemResetLoginRequest {
	this := SandboxItemResetLoginRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewSandboxItemResetLoginRequestWithDefaults instantiates a new SandboxItemResetLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemResetLoginRequestWithDefaults() *SandboxItemResetLoginRequest {
	this := SandboxItemResetLoginRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxItemResetLoginRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemResetLoginRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxItemResetLoginRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxItemResetLoginRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxItemResetLoginRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemResetLoginRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxItemResetLoginRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxItemResetLoginRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *SandboxItemResetLoginRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SandboxItemResetLoginRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SandboxItemResetLoginRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o SandboxItemResetLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxItemResetLoginRequest struct {
	value *SandboxItemResetLoginRequest
	isSet bool
}

func (v NullableSandboxItemResetLoginRequest) Get() *SandboxItemResetLoginRequest {
	return v.value
}

func (v *NullableSandboxItemResetLoginRequest) Set(val *SandboxItemResetLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemResetLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemResetLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemResetLoginRequest(val *SandboxItemResetLoginRequest) *NullableSandboxItemResetLoginRequest {
	return &NullableSandboxItemResetLoginRequest{value: val, isSet: true}
}

func (v NullableSandboxItemResetLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemResetLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


