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

// DashboardUserGetRequest Request input for fetching a dashboard user
type DashboardUserGetRequest struct {
	// ID of the associated user.
	DashboardUserId string `json:"dashboard_user_id"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
}

// NewDashboardUserGetRequest instantiates a new DashboardUserGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardUserGetRequest(dashboardUserId string) *DashboardUserGetRequest {
	this := DashboardUserGetRequest{}
	this.DashboardUserId = dashboardUserId
	return &this
}

// NewDashboardUserGetRequestWithDefaults instantiates a new DashboardUserGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardUserGetRequestWithDefaults() *DashboardUserGetRequest {
	this := DashboardUserGetRequest{}
	return &this
}

// GetDashboardUserId returns the DashboardUserId field value
func (o *DashboardUserGetRequest) GetDashboardUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardUserId
}

// GetDashboardUserIdOk returns a tuple with the DashboardUserId field value
// and a boolean to check if the value has been set.
func (o *DashboardUserGetRequest) GetDashboardUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DashboardUserId, true
}

// SetDashboardUserId sets field value
func (o *DashboardUserGetRequest) SetDashboardUserId(v string) {
	o.DashboardUserId = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *DashboardUserGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *DashboardUserGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *DashboardUserGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DashboardUserGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardUserGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DashboardUserGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DashboardUserGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

func (o DashboardUserGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dashboard_user_id"] = o.DashboardUserId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardUserGetRequest struct {
	value *DashboardUserGetRequest
	isSet bool
}

func (v NullableDashboardUserGetRequest) Get() *DashboardUserGetRequest {
	return v.value
}

func (v *NullableDashboardUserGetRequest) Set(val *DashboardUserGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardUserGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardUserGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardUserGetRequest(val *DashboardUserGetRequest) *NullableDashboardUserGetRequest {
	return &NullableDashboardUserGetRequest{value: val, isSet: true}
}

func (v NullableDashboardUserGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardUserGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


