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

// WatchlistScreeningReview A review submitted by a team member for an individual watchlist screening. A review can be either a comment on the current screening state, actions taken against hits attached to the watchlist screening, or both.
type WatchlistScreeningReview struct {
	// ID of the associated review.
	Id string `json:"id"`
	// Hits marked as a true positive after thorough manual review. These hits will never recur or be updated once dismissed. In most cases, confirmed hits indicate that the customer should be rejected.
	ConfirmedHits []string `json:"confirmed_hits"`
	// Hits marked as a false positive after thorough manual review. These hits will never recur or be updated once dismissed.
	DismissedHits []string `json:"dismissed_hits"`
	// A comment submitted by a team member as part of reviewing a watchlist screening.
	Comment NullableString `json:"comment"`
	AuditTrail WatchlistScreeningAuditTrail `json:"audit_trail"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningReview WatchlistScreeningReview

// NewWatchlistScreeningReview instantiates a new WatchlistScreeningReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningReview(id string, confirmedHits []string, dismissedHits []string, comment NullableString, auditTrail WatchlistScreeningAuditTrail) *WatchlistScreeningReview {
	this := WatchlistScreeningReview{}
	this.Id = id
	this.ConfirmedHits = confirmedHits
	this.DismissedHits = dismissedHits
	this.Comment = comment
	this.AuditTrail = auditTrail
	return &this
}

// NewWatchlistScreeningReviewWithDefaults instantiates a new WatchlistScreeningReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningReviewWithDefaults() *WatchlistScreeningReview {
	this := WatchlistScreeningReview{}
	return &this
}

// GetId returns the Id field value
func (o *WatchlistScreeningReview) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningReview) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WatchlistScreeningReview) SetId(v string) {
	o.Id = v
}

// GetConfirmedHits returns the ConfirmedHits field value
func (o *WatchlistScreeningReview) GetConfirmedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ConfirmedHits
}

// GetConfirmedHitsOk returns a tuple with the ConfirmedHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningReview) GetConfirmedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfirmedHits, true
}

// SetConfirmedHits sets field value
func (o *WatchlistScreeningReview) SetConfirmedHits(v []string) {
	o.ConfirmedHits = v
}

// GetDismissedHits returns the DismissedHits field value
func (o *WatchlistScreeningReview) GetDismissedHits() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DismissedHits
}

// GetDismissedHitsOk returns a tuple with the DismissedHits field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningReview) GetDismissedHitsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DismissedHits, true
}

// SetDismissedHits sets field value
func (o *WatchlistScreeningReview) SetDismissedHits(v []string) {
	o.DismissedHits = v
}

// GetComment returns the Comment field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningReview) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}

	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningReview) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// SetComment sets field value
func (o *WatchlistScreeningReview) SetComment(v string) {
	o.Comment.Set(&v)
}

// GetAuditTrail returns the AuditTrail field value
func (o *WatchlistScreeningReview) GetAuditTrail() WatchlistScreeningAuditTrail {
	if o == nil {
		var ret WatchlistScreeningAuditTrail
		return ret
	}

	return o.AuditTrail
}

// GetAuditTrailOk returns a tuple with the AuditTrail field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningReview) GetAuditTrailOk() (*WatchlistScreeningAuditTrail, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuditTrail, true
}

// SetAuditTrail sets field value
func (o *WatchlistScreeningReview) SetAuditTrail(v WatchlistScreeningAuditTrail) {
	o.AuditTrail = v
}

func (o WatchlistScreeningReview) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningReview) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningReview := _WatchlistScreeningReview{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningReview); err == nil {
		*o = WatchlistScreeningReview(varWatchlistScreeningReview)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningReview struct {
	value *WatchlistScreeningReview
	isSet bool
}

func (v NullableWatchlistScreeningReview) Get() *WatchlistScreeningReview {
	return v.value
}

func (v *NullableWatchlistScreeningReview) Set(val *WatchlistScreeningReview) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningReview) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningReview(val *WatchlistScreeningReview) *NullableWatchlistScreeningReview {
	return &NullableWatchlistScreeningReview{value: val, isSet: true}
}

func (v NullableWatchlistScreeningReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


