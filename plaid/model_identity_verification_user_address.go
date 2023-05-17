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

// IdentityVerificationUserAddress Even if an address has been collected, some fields may be null depending on the region's addressing system. For example:  Addresses from the United Kingdom will not include a region  Addresses from Hong Kong will not include postal code
type IdentityVerificationUserAddress struct {
	// The primary street portion of an address. If the user has submitted their address, this field will always be filled.
	Street NullableString `json:"street"`
	// Extra street information, like an apartment or suite number.
	Street2 NullableString `json:"street2"`
	// City from the end user's address
	City NullableString `json:"city"`
	// An ISO 3166-2 subdivision code. Related terms would be \"state\", \"province\", \"prefecture\", \"zone\", \"subdivision\", etc.
	Region NullableString `json:"region"`
	// The postal code for the associated address. Between 2 and 10 alphanumeric characters. For US-based addresses this must be 5 numeric digits.
	PostalCode NullableString `json:"postal_code"`
	// Valid, capitalized, two-letter ISO code representing the country of this object. Must be in ISO 3166-1 alpha-2 form.
	Country string `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationUserAddress IdentityVerificationUserAddress

// NewIdentityVerificationUserAddress instantiates a new IdentityVerificationUserAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationUserAddress(street NullableString, street2 NullableString, city NullableString, region NullableString, postalCode NullableString, country string) *IdentityVerificationUserAddress {
	this := IdentityVerificationUserAddress{}
	this.Street = street
	this.Street2 = street2
	this.City = city
	this.Region = region
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewIdentityVerificationUserAddressWithDefaults instantiates a new IdentityVerificationUserAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationUserAddressWithDefaults() *IdentityVerificationUserAddress {
	this := IdentityVerificationUserAddress{}
	return &this
}

// GetStreet returns the Street field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserAddress) GetStreet() string {
	if o == nil || o.Street.Get() == nil {
		var ret string
		return ret
	}

	return *o.Street.Get()
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserAddress) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street.Get(), o.Street.IsSet()
}

// SetStreet sets field value
func (o *IdentityVerificationUserAddress) SetStreet(v string) {
	o.Street.Set(&v)
}

// GetStreet2 returns the Street2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserAddress) GetStreet2() string {
	if o == nil || o.Street2.Get() == nil {
		var ret string
		return ret
	}

	return *o.Street2.Get()
}

// GetStreet2Ok returns a tuple with the Street2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserAddress) GetStreet2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street2.Get(), o.Street2.IsSet()
}

// SetStreet2 sets field value
func (o *IdentityVerificationUserAddress) SetStreet2(v string) {
	o.Street2.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *IdentityVerificationUserAddress) SetCity(v string) {
	o.City.Set(&v)
}

// GetRegion returns the Region field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserAddress) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}

	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserAddress) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// SetRegion sets field value
func (o *IdentityVerificationUserAddress) SetRegion(v string) {
	o.Region.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationUserAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationUserAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *IdentityVerificationUserAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
func (o *IdentityVerificationUserAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationUserAddress) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *IdentityVerificationUserAddress) SetCountry(v string) {
	o.Country = v
}

func (o IdentityVerificationUserAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["street"] = o.Street.Get()
	}
	if true {
		toSerialize["street2"] = o.Street2.Get()
	}
	if true {
		toSerialize["city"] = o.City.Get()
	}
	if true {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationUserAddress) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationUserAddress := _IdentityVerificationUserAddress{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationUserAddress); err == nil {
		*o = IdentityVerificationUserAddress(varIdentityVerificationUserAddress)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "street")
		delete(additionalProperties, "street2")
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationUserAddress struct {
	value *IdentityVerificationUserAddress
	isSet bool
}

func (v NullableIdentityVerificationUserAddress) Get() *IdentityVerificationUserAddress {
	return v.value
}

func (v *NullableIdentityVerificationUserAddress) Set(val *IdentityVerificationUserAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationUserAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationUserAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationUserAddress(val *IdentityVerificationUserAddress) *NullableIdentityVerificationUserAddress {
	return &NullableIdentityVerificationUserAddress{value: val, isSet: true}
}

func (v NullableIdentityVerificationUserAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationUserAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


