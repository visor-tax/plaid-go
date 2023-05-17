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

// ClientProvidedTransactionLocation A representation of where a transaction took place.  Use this field to pass in structured location information you may have about your transactions. Providing location data is optional but can increase result quality. If you have unstructured location information, it may be appended to the `description` field.
type ClientProvidedTransactionLocation struct {
	// The country where the transaction occurred.
	Country *string `json:"country,omitempty"`
	// The region or state where the transaction occurred.
	Region *string `json:"region,omitempty"`
	// The city where the transaction occurred.
	City *string `json:"city,omitempty"`
	// The street address where the transaction occurred.
	Address *string `json:"address,omitempty"`
	// The postal code where the transaction occurred.
	PostalCode *string `json:"postal_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClientProvidedTransactionLocation ClientProvidedTransactionLocation

// NewClientProvidedTransactionLocation instantiates a new ClientProvidedTransactionLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProvidedTransactionLocation() *ClientProvidedTransactionLocation {
	this := ClientProvidedTransactionLocation{}
	return &this
}

// NewClientProvidedTransactionLocationWithDefaults instantiates a new ClientProvidedTransactionLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProvidedTransactionLocationWithDefaults() *ClientProvidedTransactionLocation {
	this := ClientProvidedTransactionLocation{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ClientProvidedTransactionLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransactionLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ClientProvidedTransactionLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ClientProvidedTransactionLocation) SetCountry(v string) {
	o.Country = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClientProvidedTransactionLocation) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransactionLocation) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClientProvidedTransactionLocation) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClientProvidedTransactionLocation) SetRegion(v string) {
	o.Region = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *ClientProvidedTransactionLocation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransactionLocation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *ClientProvidedTransactionLocation) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *ClientProvidedTransactionLocation) SetCity(v string) {
	o.City = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ClientProvidedTransactionLocation) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransactionLocation) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ClientProvidedTransactionLocation) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ClientProvidedTransactionLocation) SetAddress(v string) {
	o.Address = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *ClientProvidedTransactionLocation) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientProvidedTransactionLocation) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ClientProvidedTransactionLocation) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *ClientProvidedTransactionLocation) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o ClientProvidedTransactionLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ClientProvidedTransactionLocation) UnmarshalJSON(bytes []byte) (err error) {
	varClientProvidedTransactionLocation := _ClientProvidedTransactionLocation{}

	if err = json.Unmarshal(bytes, &varClientProvidedTransactionLocation); err == nil {
		*o = ClientProvidedTransactionLocation(varClientProvidedTransactionLocation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "country")
		delete(additionalProperties, "region")
		delete(additionalProperties, "city")
		delete(additionalProperties, "address")
		delete(additionalProperties, "postal_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClientProvidedTransactionLocation struct {
	value *ClientProvidedTransactionLocation
	isSet bool
}

func (v NullableClientProvidedTransactionLocation) Get() *ClientProvidedTransactionLocation {
	return v.value
}

func (v *NullableClientProvidedTransactionLocation) Set(val *ClientProvidedTransactionLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProvidedTransactionLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProvidedTransactionLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProvidedTransactionLocation(val *ClientProvidedTransactionLocation) *NullableClientProvidedTransactionLocation {
	return &NullableClientProvidedTransactionLocation{value: val, isSet: true}
}

func (v NullableClientProvidedTransactionLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProvidedTransactionLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


