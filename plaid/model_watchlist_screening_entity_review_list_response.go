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

// WatchlistScreeningEntityReviewListResponse Paginated list of entity watchlist screening reviews
type WatchlistScreeningEntityReviewListResponse struct {
	// List of entity watchlist screening reviews
	EntityWatchlistScreeningReviews []EntityWatchlistScreeningReview `json:"entity_watchlist_screening_reviews"`
	// An identifier that determines which page of results you receive.
	NextCursor NullableString `json:"next_cursor"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningEntityReviewListResponse WatchlistScreeningEntityReviewListResponse

// NewWatchlistScreeningEntityReviewListResponse instantiates a new WatchlistScreeningEntityReviewListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningEntityReviewListResponse(entityWatchlistScreeningReviews []EntityWatchlistScreeningReview, nextCursor NullableString, requestId string) *WatchlistScreeningEntityReviewListResponse {
	this := WatchlistScreeningEntityReviewListResponse{}
	this.EntityWatchlistScreeningReviews = entityWatchlistScreeningReviews
	this.NextCursor = nextCursor
	this.RequestId = requestId
	return &this
}

// NewWatchlistScreeningEntityReviewListResponseWithDefaults instantiates a new WatchlistScreeningEntityReviewListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningEntityReviewListResponseWithDefaults() *WatchlistScreeningEntityReviewListResponse {
	this := WatchlistScreeningEntityReviewListResponse{}
	return &this
}

// GetEntityWatchlistScreeningReviews returns the EntityWatchlistScreeningReviews field value
func (o *WatchlistScreeningEntityReviewListResponse) GetEntityWatchlistScreeningReviews() []EntityWatchlistScreeningReview {
	if o == nil {
		var ret []EntityWatchlistScreeningReview
		return ret
	}

	return o.EntityWatchlistScreeningReviews
}

// GetEntityWatchlistScreeningReviewsOk returns a tuple with the EntityWatchlistScreeningReviews field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityReviewListResponse) GetEntityWatchlistScreeningReviewsOk() (*[]EntityWatchlistScreeningReview, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EntityWatchlistScreeningReviews, true
}

// SetEntityWatchlistScreeningReviews sets field value
func (o *WatchlistScreeningEntityReviewListResponse) SetEntityWatchlistScreeningReviews(v []EntityWatchlistScreeningReview) {
	o.EntityWatchlistScreeningReviews = v
}

// GetNextCursor returns the NextCursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningEntityReviewListResponse) GetNextCursor() string {
	if o == nil || o.NextCursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningEntityReviewListResponse) GetNextCursorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// SetNextCursor sets field value
func (o *WatchlistScreeningEntityReviewListResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *WatchlistScreeningEntityReviewListResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningEntityReviewListResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WatchlistScreeningEntityReviewListResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WatchlistScreeningEntityReviewListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entity_watchlist_screening_reviews"] = o.EntityWatchlistScreeningReviews
	}
	if true {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningEntityReviewListResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningEntityReviewListResponse := _WatchlistScreeningEntityReviewListResponse{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningEntityReviewListResponse); err == nil {
		*o = WatchlistScreeningEntityReviewListResponse(varWatchlistScreeningEntityReviewListResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "entity_watchlist_screening_reviews")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningEntityReviewListResponse struct {
	value *WatchlistScreeningEntityReviewListResponse
	isSet bool
}

func (v NullableWatchlistScreeningEntityReviewListResponse) Get() *WatchlistScreeningEntityReviewListResponse {
	return v.value
}

func (v *NullableWatchlistScreeningEntityReviewListResponse) Set(val *WatchlistScreeningEntityReviewListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningEntityReviewListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningEntityReviewListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningEntityReviewListResponse(val *WatchlistScreeningEntityReviewListResponse) *NullableWatchlistScreeningEntityReviewListResponse {
	return &NullableWatchlistScreeningEntityReviewListResponse{value: val, isSet: true}
}

func (v NullableWatchlistScreeningEntityReviewListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningEntityReviewListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


