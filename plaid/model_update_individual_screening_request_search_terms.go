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

// UpdateIndividualScreeningRequestSearchTerms Search terms for editing an individual watchlist screening
type UpdateIndividualScreeningRequestSearchTerms struct {
	// ID of the associated program.
	WatchlistProgramId *string `json:"watchlist_program_id,omitempty"`
	// The legal name of the individual being screened.
	LegalName *string `json:"legal_name,omitempty"`
	// A date in the format YYYY-MM-DD (RFC 3339 Section 5.6).
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	// The numeric or alphanumeric identifier associated with this document.
	DocumentNumber *string `json:"document_number,omitempty"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country *string `json:"country,omitempty"`
}

// NewUpdateIndividualScreeningRequestSearchTerms instantiates a new UpdateIndividualScreeningRequestSearchTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIndividualScreeningRequestSearchTerms() *UpdateIndividualScreeningRequestSearchTerms {
	this := UpdateIndividualScreeningRequestSearchTerms{}
	return &this
}

// NewUpdateIndividualScreeningRequestSearchTermsWithDefaults instantiates a new UpdateIndividualScreeningRequestSearchTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIndividualScreeningRequestSearchTermsWithDefaults() *UpdateIndividualScreeningRequestSearchTerms {
	this := UpdateIndividualScreeningRequestSearchTerms{}
	return &this
}

// GetWatchlistProgramId returns the WatchlistProgramId field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetWatchlistProgramId() string {
	if o == nil || o.WatchlistProgramId == nil {
		var ret string
		return ret
	}
	return *o.WatchlistProgramId
}

// GetWatchlistProgramIdOk returns a tuple with the WatchlistProgramId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetWatchlistProgramIdOk() (*string, bool) {
	if o == nil || o.WatchlistProgramId == nil {
		return nil, false
	}
	return o.WatchlistProgramId, true
}

// HasWatchlistProgramId returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) HasWatchlistProgramId() bool {
	if o != nil && o.WatchlistProgramId != nil {
		return true
	}

	return false
}

// SetWatchlistProgramId gets a reference to the given string and assigns it to the WatchlistProgramId field.
func (o *UpdateIndividualScreeningRequestSearchTerms) SetWatchlistProgramId(v string) {
	o.WatchlistProgramId = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetLegalName() string {
	if o == nil || o.LegalName == nil {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetLegalNameOk() (*string, bool) {
	if o == nil || o.LegalName == nil {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) HasLegalName() bool {
	if o != nil && o.LegalName != nil {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *UpdateIndividualScreeningRequestSearchTerms) SetLegalName(v string) {
	o.LegalName = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetDateOfBirthOk() (*string, bool) {
	if o == nil || o.DateOfBirth == nil {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth != nil {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *UpdateIndividualScreeningRequestSearchTerms) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber == nil {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetDocumentNumberOk() (*string, bool) {
	if o == nil || o.DocumentNumber == nil {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) HasDocumentNumber() bool {
	if o != nil && o.DocumentNumber != nil {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *UpdateIndividualScreeningRequestSearchTerms) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *UpdateIndividualScreeningRequestSearchTerms) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *UpdateIndividualScreeningRequestSearchTerms) SetCountry(v string) {
	o.Country = &v
}

func (o UpdateIndividualScreeningRequestSearchTerms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WatchlistProgramId != nil {
		toSerialize["watchlist_program_id"] = o.WatchlistProgramId
	}
	if o.LegalName != nil {
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

type NullableUpdateIndividualScreeningRequestSearchTerms struct {
	value *UpdateIndividualScreeningRequestSearchTerms
	isSet bool
}

func (v NullableUpdateIndividualScreeningRequestSearchTerms) Get() *UpdateIndividualScreeningRequestSearchTerms {
	return v.value
}

func (v *NullableUpdateIndividualScreeningRequestSearchTerms) Set(val *UpdateIndividualScreeningRequestSearchTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIndividualScreeningRequestSearchTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIndividualScreeningRequestSearchTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIndividualScreeningRequestSearchTerms(val *UpdateIndividualScreeningRequestSearchTerms) *NullableUpdateIndividualScreeningRequestSearchTerms {
	return &NullableUpdateIndividualScreeningRequestSearchTerms{value: val, isSet: true}
}

func (v NullableUpdateIndividualScreeningRequestSearchTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIndividualScreeningRequestSearchTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


