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

// WatchlistScreeningIndividualReviewCreateResponse A review submitted by a team member for an individual watchlist screening. A review can be either a comment on the current screening state, actions taken against hits attached to the watchlist screening, or both.
type WatchlistScreeningIndividualReviewCreateResponse struct {
	// ID of the associated review.
	Id string `json:"id"`
	// Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.
	ConfirmedHits []string `json:"confirmed_hits"`
	// Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.
	DismissedHits []string `json:"dismissed_hits"`
	// A comment submitted by a team member as part of reviewing a watchlist screening.
	Comment NullableString `json:"comment"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningIndividualReviewCreateResponse WatchlistScreeningIndividualReviewCreateResponse

// NewWatchlistScreeningIndividualReviewCreateResponse instantiates a new WatchlistScreeningIndividualReviewCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningIndividualReviewCreateResponse(id string, confirmedHits []string, dismissedHits []string, comment NullableString, auditTrail WatchlistScreeningAuditTrail, requestId string) *WatchlistScreeningIndividualReviewCreateResponse {
	this := WatchlistScreeningIndividualReviewCreateResponse{}
	this.Id = id
	this.ConfirmedHits = confirmedHits
	this.DismissedHits = dismissedHits
	this.Comment = comment
	this.AuditTrail = auditTrail
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningIndividualReviewCreateResponseWithDefaults instantiates a new WatchlistScreeningIndividualReviewCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningIndividualReviewCreateResponseWithDefaults() *WatchlistScreeningIndividualReviewCreateResponse {
	this := WatchlistScreeningIndividualReviewCreateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) SetId(v string) {
	o.Id = v
}

// GetConfirmedHits returns the ConfirmedHits field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetConfirmedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfirmedHits
}

// GetConfirmedHitsOk returns a tuple with the ConfirmedHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetConfirmedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedHits, true
}

// SetConfirmedHits sets field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) SetConfirmedHits(v []string) {
	o.ConfirmedHits = v
}

// GetDismissedHits returns the DismissedHits field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetDismissedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DismissedHits
}

// GetDismissedHitsOk returns a tuple with the DismissedHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetDismissedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DismissedHits, true
}

// SetDismissedHits sets field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) SetDismissedHits(v []string) {
	o.DismissedHits = v
}

// GetComment returns the Comment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// SetComment sets field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) SetComment(v string) {
	o.Comment.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningIndividualReviewCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningIndividualReviewCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningIndividualReviewCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["confirmed_hits"] = o.ConfirmedHits
	}
	if true {
		toSerialize["dismissed_hits"] = o.DismissedHits
	}
	if true {
		toSerialize["comment"] = o.Comment.Get()
	}
	if true {
		toSerialize["audit_trail"] = o.AuditTrail
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningIndividualReviewCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningIndividualReviewCreateResponse := _WatchlistScreeningIndividualReviewCreateResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningIndividualReviewCreateResponse); err == nil {
		*o = WatchlistScreeningIndividualReviewCreateResponse(varWatchlistScreeningIndividualReviewCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "confirmed_hits")
		delete(additionalProperties, "dismissed_hits")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "audit_trail")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningIndividualReviewCreateResponse struct {
	value *WatchlistScreeningIndividualReviewCreateResponse
	isSet bool
}

func (v NullableWatchlistScreeningIndividualReviewCreateResponse) Get() *WatchlistScreeningIndividualReviewCreateResponse {
	return v.value
}

func (v *NullableWatchlistScreeningIndividualReviewCreateResponse) Set(val *WatchlistScreeningIndividualReviewCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningIndividualReviewCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningIndividualReviewCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningIndividualReviewCreateResponse(val *WatchlistScreeningIndividualReviewCreateResponse) *NullableWatchlistScreeningIndividualReviewCreateResponse {
	return &NullableWatchlistScreeningIndividualReviewCreateResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningIndividualReviewCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningIndividualReviewCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


