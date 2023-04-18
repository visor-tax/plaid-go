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

// DetailedOriginator Originator and their status.
type DetailedOriginator struct {
	// Originator’s client ID.
	ClientId string `json:"client_id"`
	TransferDiligenceStatus TransferDiligenceStatus `json:"transfer_diligence_status"`
	CompanyName string `json:"company_name"`
	AdditionalProperties map[string]interface{}
}

type _DetailedOriginator DetailedOriginator

// NewDetailedOriginator instantiates a new DetailedOriginator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedOriginator(clientId string, transferDiligenceStatus TransferDiligenceStatus, companyName string) *DetailedOriginator {
	this := DetailedOriginator{}
	this.ClientId = clientId
	this.TransferDiligenceStatus = transferDiligenceStatus
	this.CompanyName = companyName
	return &this
}

// NewDetailedOriginatorWithDefaults instantiates a new DetailedOriginator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedOriginatorWithDefaults() *DetailedOriginator {
	this := DetailedOriginator{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *DetailedOriginator) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *DetailedOriginator) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *DetailedOriginator) SetClientId(v string) {
	o.ClientId = v
}

// GetTransferDiligenceStatus returns the TransferDiligenceStatus field value
func (o *DetailedOriginator) GetTransferDiligenceStatus() TransferDiligenceStatus {
	if o == nil {
		var ret TransferDiligenceStatus
		return ret
	}

	return o.TransferDiligenceStatus
}

// GetTransferDiligenceStatusOk returns a tuple with the TransferDiligenceStatus field value
// and a boolean to check if the value has been set.
func (o *DetailedOriginator) GetTransferDiligenceStatusOk() (*TransferDiligenceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferDiligenceStatus, true
}

// SetTransferDiligenceStatus sets field value
func (o *DetailedOriginator) SetTransferDiligenceStatus(v TransferDiligenceStatus) {
	o.TransferDiligenceStatus = v
}

// GetCompanyName returns the CompanyName field value
func (o *DetailedOriginator) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *DetailedOriginator) GetCompanyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *DetailedOriginator) SetCompanyName(v string) {
	o.CompanyName = v
}

func (o DetailedOriginator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["transfer_diligence_status"] = o.TransferDiligenceStatus
	}
	if true {
		toSerialize["company_name"] = o.CompanyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DetailedOriginator) UnmarshalJSON(bytes []byte) (err error) {
	varDetailedOriginator := _DetailedOriginator{}

	if err = json.Unmarshal(bytes, &varDetailedOriginator); err == nil {
		*o = DetailedOriginator(varDetailedOriginator)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "transfer_diligence_status")
		delete(additionalProperties, "company_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetailedOriginator struct {
	value *DetailedOriginator
	isSet bool
}

func (v NullableDetailedOriginator) Get() *DetailedOriginator {
	return v.value
}

func (v *NullableDetailedOriginator) Set(val *DetailedOriginator) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedOriginator) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedOriginator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedOriginator(val *DetailedOriginator) *NullableDetailedOriginator {
	return &NullableDetailedOriginator{value: val, isSet: true}
}

func (v NullableDetailedOriginator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedOriginator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


