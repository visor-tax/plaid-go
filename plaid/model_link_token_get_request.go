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

// LinkTokenGetRequest LinkTokenGetRequest defines the request schema for `/link/token/get`
type LinkTokenGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A `link_token` from a previous invocation of `/link/token/create`
	LinkToken string `json:"link_token"`
}

// NewLinkTokenGetRequest instantiates a new LinkTokenGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenGetRequest(linkToken string) *LinkTokenGetRequest {
	this := LinkTokenGetRequest{}
	this.LinkToken = linkToken
	return &this
}

// NewLinkTokenGetRequestWithDefaults instantiates a new LinkTokenGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenGetRequestWithDefaults() *LinkTokenGetRequest {
	this := LinkTokenGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LinkTokenGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LinkTokenGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LinkTokenGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *LinkTokenGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *LinkTokenGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *LinkTokenGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetLinkToken returns the LinkToken field value
func (o *LinkTokenGetRequest) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *LinkTokenGetRequest) GetLinkTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *LinkTokenGetRequest) SetLinkToken(v string) {
	o.LinkToken = v
}

func (o LinkTokenGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["link_token"] = o.LinkToken
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenGetRequest struct {
	value *LinkTokenGetRequest
	isSet bool
}

func (v NullableLinkTokenGetRequest) Get() *LinkTokenGetRequest {
	return v.value
}

func (v *NullableLinkTokenGetRequest) Set(val *LinkTokenGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenGetRequest(val *LinkTokenGetRequest) *NullableLinkTokenGetRequest {
	return &NullableLinkTokenGetRequest{value: val, isSet: true}
}

func (v NullableLinkTokenGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


