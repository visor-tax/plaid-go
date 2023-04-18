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

// LinkTokenGetResponse LinkTokenGetResponse defines the response schema for `/link/token/get`
type LinkTokenGetResponse struct {
	// A `link_token`, which can be supplied to Link in order to initialize it and receive a `public_token`, which can be exchanged for an `access_token`.
	LinkToken string `json:"link_token"`
	// The creation timestamp for the `link_token`, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	CreatedAt NullableTime `json:"created_at"`
	// The expiration timestamp for the `link_token`, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	Expiration NullableTime `json:"expiration"`
	Metadata LinkTokenGetMetadataResponse `json:"metadata"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenGetResponse LinkTokenGetResponse

// NewLinkTokenGetResponse instantiates a new LinkTokenGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenGetResponse(linkToken string, createdAt NullableTime, expiration NullableTime, metadata LinkTokenGetMetadataResponse, requestId string) *LinkTokenGetResponse {
	this := LinkTokenGetResponse{}
	this.LinkToken = linkToken
	this.CreatedAt = createdAt
	this.Expiration = expiration
	this.Metadata = metadata
	this.RequestId = requestId
	return &this
}

// NewLinkTokenGetResponseWithDefaults instantiates a new LinkTokenGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenGetResponseWithDefaults() *LinkTokenGetResponse {
	this := LinkTokenGetResponse{}
	return &this
}

// GetLinkToken returns the LinkToken field value
func (o *LinkTokenGetResponse) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *LinkTokenGetResponse) GetLinkTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *LinkTokenGetResponse) SetLinkToken(v string) {
	o.LinkToken = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LinkTokenGetResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *LinkTokenGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetExpiration returns the Expiration field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LinkTokenGetResponse) GetExpiration() time.Time {
	if o == nil || o.Expiration.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenGetResponse) GetExpirationOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// SetExpiration sets field value
func (o *LinkTokenGetResponse) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// GetMetadata returns the Metadata field value
func (o *LinkTokenGetResponse) GetMetadata() LinkTokenGetMetadataResponse {
	if o == nil {
		var ret LinkTokenGetMetadataResponse
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *LinkTokenGetResponse) GetMetadataOk() (*LinkTokenGetMetadataResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *LinkTokenGetResponse) SetMetadata(v LinkTokenGetMetadataResponse) {
	o.Metadata = v
}

// GetRequestId returns the RequestId field value
func (o *LinkTokenGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *LinkTokenGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o LinkTokenGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["link_token"] = o.LinkToken
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if true {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenGetResponse := _LinkTokenGetResponse{}

	if err = json.Unmarshal(bytes, &varLinkTokenGetResponse); err == nil {
		*o = LinkTokenGetResponse(varLinkTokenGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "link_token")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenGetResponse struct {
	value *LinkTokenGetResponse
	isSet bool
}

func (v NullableLinkTokenGetResponse) Get() *LinkTokenGetResponse {
	return v.value
}

func (v *NullableLinkTokenGetResponse) Set(val *LinkTokenGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenGetResponse(val *LinkTokenGetResponse) *NullableLinkTokenGetResponse {
	return &NullableLinkTokenGetResponse{value: val, isSet: true}
}

func (v NullableLinkTokenGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


