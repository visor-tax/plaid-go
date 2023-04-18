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

// WatchlistScreeningRequestSearchTerms Search inputs for creating a watchlist screening
type WatchlistScreeningRequestSearchTerms struct {
	// ID of the associated program.
	WatchlistProgramId string `json:"watchlist_program_id"`
	// The legal name of the individual being screened.
	LegalName string `json:"legal_name"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	// The numeric or alphanumeric identifier associated with this document.
	DocumentNumber *string `json:"document_number,omitempty"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country *string `json:"country,omitempty"`
}

// NewWatchlistScreeningRequestSearchTerms instantiates a new WatchlistScreeningRequestSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistScreeningRequestSearchTerms(watchlistProgramId string, legalName string) *WatchlistScreeningRequestSearchTerms {
	this := WatchlistScreeningRequestSearchTerms{}
	this.WatchlistProgramId = watchlistProgramId
	this.LegalName = legalName
	return &this
}

// NewWatchlistScreeningRequestSearchTermsWithDefaults instantiates a new WatchlistScreeningRequestSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistScreeningRequestSearchTermsWithDefaults() *WatchlistScreeningRequestSearchTerms {
	this := WatchlistScreeningRequestSearchTerms{}
	return &this
}

// GetWatchlistProgramId returns the WatchlistProgramId field value
func (o *WatchlistScreeningRequestSearchTerms) GetWatchlistProgramId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WatchlistProgramId
}

// GetWatchlistProgramIdOk returns a tuple with the WatchlistProgramId field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningRequestSearchTerms) GetWatchlistProgramIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WatchlistProgramId, true
}

// SetWatchlistProgramId sets field value
func (o *WatchlistScreeningRequestSearchTerms) SetWatchlistProgramId(v string) {
	o.WatchlistProgramId = v
}

// GetLegalName returns the LegalName field value
func (o *WatchlistScreeningRequestSearchTerms) GetLegalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningRequestSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LegalName, true
}

// SetLegalName sets field value
func (o *WatchlistScreeningRequestSearchTerms) SetLegalName(v string) {
	o.LegalName = v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *WatchlistScreeningRequestSearchTerms) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningRequestSearchTerms) GetDateOfBirthOk() (*string, bool) {
	if o == nil || o.DateOfBirth == nil {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *WatchlistScreeningRequestSearchTerms) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth != nil {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *WatchlistScreeningRequestSearchTerms) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *WatchlistScreeningRequestSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber == nil {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningRequestSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil || o.DocumentNumber == nil {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *WatchlistScreeningRequestSearchTerms) HasDocumentNumber() bool {
	if o != nil && o.DocumentNumber != nil {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *WatchlistScreeningRequestSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *WatchlistScreeningRequestSearchTerms) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistScreeningRequestSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *WatchlistScreeningRequestSearchTerms) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *WatchlistScreeningRequestSearchTerms) SetCountry(v string) {
	o.Country = &v
}

func (o WatchlistScreeningRequestSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["watchlist_program_id"] = o.WatchlistProgramId
	}
	if true {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.DateOfBirth != nil {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if o.DocumentNumber != nil {
		toSerialize["document_number"] = o.DocumentNumber
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableWatchlistScreeningRequestSearchTerms struct {
	value *WatchlistScreeningRequestSearchTerms
	isSet bool
}

func (v NullableWatchlistScreeningRequestSearchTerms) Get() *WatchlistScreeningRequestSearchTerms {
	return v.value
}

func (v *NullableWatchlistScreeningRequestSearchTerms) Set(val *WatchlistScreeningRequestSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistScreeningRequestSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistScreeningRequestSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistScreeningRequestSearchTerms(val *WatchlistScreeningRequestSearchTerms) *NullableWatchlistScreeningRequestSearchTerms {
	return &NullableWatchlistScreeningRequestSearchTerms{value: val, isSet: true}
}

func (v NullableWatchlistScreeningRequestSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistScreeningRequestSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


