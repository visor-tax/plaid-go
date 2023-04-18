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
	"time"
)

// WatchlistScreeningAuditTrail Information about the last change made to the parent object specifying what caused the change as well as when it occurred.
type WatchlistScreeningAuditTrail struct {
	Source Source `json:"source"`
	// ID of the associated user.
	DashboardUserId NullableString `json:"dashboard_user_id"`
	// An ISO8601 formatted timestamp.
	Timestamp time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _WatchlistScreeningAuditTrail WatchlistScreeningAuditTrail

// NewWatchlistScreeningAuditTrail instantiates a new WatchlistScreeningAuditTrail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningAuditTrail(source Source, dashboardUserId NullableString, timestamp time.Time) *WatchlistScreeningAuditTrail {
	this := WatchlistScreeningAuditTrail{}
	this.Source = source
	this.DashboardUserId = dashboardUserId
	this.Timestamp = timestamp
	return &this
}

// NewWatchlistScreeningAuditTrailWithDefaults instantiates a new WatchlistScreeningAuditTrail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningAuditTrailWithDefaults() *WatchlistScreeningAuditTrail {
	this := WatchlistScreeningAuditTrail{}
	return &this
}

// GetSource returns the Source field value
func (o *WatchlistScreeningAuditTrail) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningAuditTrail) GetSourceOk() (*Source, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *WatchlistScreeningAuditTrail) SetSource(v Source) {
	o.Source = v
}

// GetDashboardUserId returns the DashboardUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WatchlistScreeningAuditTrail) GetDashboardUserId() string {
	if o == nil || o.DashboardUserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DashboardUserId.Get()
}

// GetDashboardUserIdOk returns a tuple with the DashboardUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchlistScreeningAuditTrail) GetDashboardUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DashboardUserId.Get(), o.DashboardUserId.IsSet()
}

// SetDashboardUserId sets field value
func (o *WatchlistScreeningAuditTrail) SetDashboardUserId(v string) {
	o.DashboardUserId.Set(&v)
}

// GetTimestamp returns the Timestamp field value
func (o *WatchlistScreeningAuditTrail) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningAuditTrail) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *WatchlistScreeningAuditTrail) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o WatchlistScreeningAuditTrail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["dashboard_user_id"] = o.DashboardUserId.Get()
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WatchlistScreeningAuditTrail) UnmarshalJSON(bytes []byte) (err error) {
	varWatchlistScreeningAuditTrail := _WatchlistScreeningAuditTrail{}

	if err = json.Unmarshal(bytes, &varWatchlistScreeningAuditTrail); err == nil {
		*o = WatchlistScreeningAuditTrail(varWatchlistScreeningAuditTrail)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source")
		delete(additionalProperties, "dashboard_user_id")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWatchlistScreeningAuditTrail struct {
	value *WatchlistScreeningAuditTrail
	isSet bool
}

func (v NullableWatchlistScreeningAuditTrail) Get() *WatchlistScreeningAuditTrail {
	return v.value
}

func (v *NullableWatchlistScreeningAuditTrail) Set(val *WatchlistScreeningAuditTrail) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningAuditTrail) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningAuditTrail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningAuditTrail(val *WatchlistScreeningAuditTrail) *NullableWatchlistScreeningAuditTrail {
	return &NullableWatchlistScreeningAuditTrail{value: val, isSet: true}
}

func (v NullableWatchlistScreeningAuditTrail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningAuditTrail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


