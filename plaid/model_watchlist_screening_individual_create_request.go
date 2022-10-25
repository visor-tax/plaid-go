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

// WatchlistScreeningIndividualCreateRequest Request input for creating an individual watchlist screening
type WatchlistScreeningIndividualCreateRequest struct {
	SearchTerms WatchlistScreeningRequestSearchTerms `json:"search_terms"`
	// An identifier to help you connect this object to your internal systems. For example, your database ID corresponding to this object.
	ClientUserId *string `json:"client_user_id,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
}

// NewWatchlistScreeningIndividualCreateRequest instantiates a new WatchlistScreeningIndividualCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualCreateRequest(searchTerms WatchlistScreeningRequestSearchTerms) *WatchlistScreeningIndividualCreateRequest {
	this := WatchlistScreeningIndividualCreateRequest{}
	this.SearchTerms = searchTerms
	return &this
}

// NewWatchlistScreeningIndividualCreateRequestWithDefaults instantiates a new WatchlistScreeningIndividualCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualCreateRequestWithDefaults() *WatchlistScreeningIndividualCreateRequest {
	this := WatchlistScreeningIndividualCreateRequest{}
	return &this
}

// GetSearchTerms returns the SearchTerms field value
func (o *WatchlistScreeningIndividualCreateRequest) GetSearchTerms() WatchlistScreeningRequestSearchTerms {
	if o == nil {
		var ret WatchlistScreeningRequestSearchTerms
		return ret
	}

	return o.SearchTerms
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualCreateRequest) GetSearchTermsOk() (*WatchlistScreeningRequestSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SearchTerms, true
}

// SetSearchTerms sets field value
func (o *WatchlistScreeningIndividualCreateRequest) SetSearchTerms(v WatchlistScreeningRequestSearchTerms) {
	o.SearchTerms = v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualCreateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualCreateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || o.ClientUserId == nil {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualCreateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId != nil {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *WatchlistScreeningIndividualCreateRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WatchlistScreeningIndividualCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WatchlistScreeningIndividualCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

func (o WatchlistScreeningIndividualCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["search_terms"] = o.SearchTerms
	}
	if o.ClientUserId != nil {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningIndividualCreateRequest struct {
	value *WatchlistScreeningIndividualCreateRequest
	isSet bool
}

func (v NullableWatchlistScreeningIndividualCreateRequest) Get() *WatchlistScreeningIndividualCreateRequest {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualCreateRequest) Set(val *WatchlistScreeningIndividualCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualCreateRequest(val *WatchlistScreeningIndividualCreateRequest) *NullableWatchlistScreeningIndividualCreateRequest {
	return &NullableWatchlistScreeningIndividualCreateRequest{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


