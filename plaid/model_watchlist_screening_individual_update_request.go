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

// WatchlistScreeningIndividualUpdateRequest Request input for editing an individual watchlist screening
type WatchlistScreeningIndividualUpdateRequest struct {
	// ID of the associated screening.
	WatchlistScreeningId string `json:"watchlist_screening_id"`
	SearchTerms NullableUpdateIndividualScreeningRequestSearchTerms `json:"search_terms,omitempty"`
	// ID of the associated user.
	Assignee *string `json:"assignee,omitempty"`
	Status *WatchlistScreeningStatus `json:"status,omitempty"`
	// An identifier to help you connect this object to your internal systems. For example, your database ID corresponding to this object.
	ClientUserId *string `json:"client_user_id,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// A list of fields to reset back to null
	ResetFields []WatchlistScreeningIndividualUpdateRequestResettableField `json:"reset_fields,omitempty"`
}

// NewWatchlistScreeningIndividualUpdateRequest instantiates a new WatchlistScreeningIndividualUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualUpdateRequest(watchlistScreeningId string) *WatchlistScreeningIndividualUpdateRequest {
	this := WatchlistScreeningIndividualUpdateRequest{}
	this.WatchlistScreeningId = watchlistScreeningId
	return &this
}

// NewWatchlistScreeningIndividualUpdateRequestWithDefaults instantiates a new WatchlistScreeningIndividualUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualUpdateRequestWithDefaults() *WatchlistScreeningIndividualUpdateRequest {
	this := WatchlistScreeningIndividualUpdateRequest{}
	return &this
}

// GetWatchlistScreeningId returns the WatchlistScreeningId field value
func (o *WatchlistScreeningIndividualUpdateRequest) GetWatchlistScreeningId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistScreeningId
}

// GetWatchlistScreeningIdOk returns a tuple with the WatchlistScreeningId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) GetWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistScreeningId, true
}

// SetWatchlistScreeningId sets field value
func (o *WatchlistScreeningIndividualUpdateRequest) SetWatchlistScreeningId(v string) {
	o.WatchlistScreeningId = v
}

// GetSearchTerms returns the SearchTerms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatchlistScreeningIndividualUpdateRequest) GetSearchTerms() UpdateIndividualScreeningRequestSearchTerms {
	if o == nil || o.SearchTerms.Get() == nil {
		var ret UpdateIndividualScreeningRequestSearchTerms
		return ret
	}
	return *o.SearchTerms.Get()
}

// GetSearchTermsOk returns a tuple with the SearchTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualUpdateRequest) GetSearchTermsOk() (*UpdateIndividualScreeningRequestSearchTerms, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SearchTerms.Get(), o.SearchTerms.IsSet()
}

// HasSearchTerms returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasSearchTerms() bool {
	if o != nil && o.SearchTerms.IsSet() {
		return true
	}

	return false
}

// SetSearchTerms gets a reference to the given NullableUpdateIndividualScreeningRequestSearchTerms and assigns it to the SearchTerms field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetSearchTerms(v UpdateIndividualScreeningRequestSearchTerms) {
	o.SearchTerms.Set(&v)
}
// SetSearchTermsNil sets the value for SearchTerms to be an explicit nil
func (o *WatchlistScreeningIndividualUpdateRequest) SetSearchTermsNil() {
	o.SearchTerms.Set(nil)
}

// UnsetSearchTerms ensures that no value is present for SearchTerms, not even an explicit nil
func (o *WatchlistScreeningIndividualUpdateRequest) UnsetSearchTerms() {
	o.SearchTerms.Unset()
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualUpdateRequest) GetAssignee() string {
	if o == nil || o.Assignee == nil {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) GetAssigneeOk() (*string, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasAssignee() bool {
	if o != nil && o.Assignee != nil {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetAssignee(v string) {
	o.Assignee = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualUpdateRequest) GetStatus() WatchlistScreeningStatus {
	if o == nil || o.Status == nil {
		var ret WatchlistScreeningStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) GetStatusOk() (*WatchlistScreeningStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WatchlistScreeningStatus and assigns it to the Status field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetStatus(v WatchlistScreeningStatus) {
	o.Status = &v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualUpdateRequest) GetClientUserId() string {
	if o == nil || o.ClientUserId == nil {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || o.ClientUserId == nil {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasClientUserId() bool {
	if o != nil && o.ClientUserId != nil {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualUpdateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WatchlistScreeningIndividualUpdateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetResetFields returns the ResetFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WatchlistScreeningIndividualUpdateRequest) GetResetFields() []WatchlistScreeningIndividualUpdateRequestResettableField {
	if o == nil  {
		var ret []WatchlistScreeningIndividualUpdateRequestResettableField
		return ret
	}
	return o.ResetFields
}

// GetResetFieldsOk returns a tuple with the ResetFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualUpdateRequest) GetResetFieldsOk() (*[]WatchlistScreeningIndividualUpdateRequestResettableField, bool) {
	if o == nil || o.ResetFields == nil {
		return nil, false
	}
	return &o.ResetFields, true
}

// HasResetFields returns a boolean if a field has been set.
func (o *WatchlistScreeningIndividualUpdateRequest) HasResetFields() bool {
	if o != nil && o.ResetFields != nil {
		return true
	}

	return false
}

// SetResetFields gets a reference to the given []WatchlistScreeningIndividualUpdateRequestResettableField and assigns it to the ResetFields field.
func (o *WatchlistScreeningIndividualUpdateRequest) SetResetFields(v []WatchlistScreeningIndividualUpdateRequestResettableField) {
	o.ResetFields = v
}

func (o WatchlistScreeningIndividualUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_screening_id"] = o.WatchlistScreeningId
	}
	if o.SearchTerms.IsSet() {
		toSerialize["search_terms"] = o.SearchTerms.Get()
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
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
	if o.ResetFields != nil {
		toSerialize["reset_fields"] = o.ResetFields
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningIndividualUpdateRequest struct {
	value *WatchlistScreeningIndividualUpdateRequest
	isSet bool
}

func (v NullableWatchlistScreeningIndividualUpdateRequest) Get() *WatchlistScreeningIndividualUpdateRequest {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualUpdateRequest) Set(val *WatchlistScreeningIndividualUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualUpdateRequest(val *WatchlistScreeningIndividualUpdateRequest) *NullableWatchlistScreeningIndividualUpdateRequest {
	return &NullableWatchlistScreeningIndividualUpdateRequest{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


