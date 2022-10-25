/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.197.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// Institution Details relating to a specific financial institution
type Institution struct {
	// Unique identifier for the institution
	InstitutionId string `json:"institution_id"`
	// The official name of the institution
	Name string `json:"name"`
	// A list of the Plaid products supported by the institution. Note that only institutions that support Instant Auth will return `auth` in the product array; institutions that do not list `auth` may still support other Auth methods such as Instant Match or Automated Micro-deposit Verification. To identify institutions that support those methods, use the `auth_metadata` object. For more details, see [Full Auth coverage](https://plaid.com/docs/auth/coverage/).
	Products []Products `json:"products"`
	// A list of the country codes supported by the institution.
	CountryCodes []CountryCode `json:"country_codes"`
	// The URL for the institution's website
	Url NullableString `json:"url,omitempty"`
	// Hexadecimal representation of the primary color used by the institution
	PrimaryColor NullableString `json:"primary_color,omitempty"`
	// Base64 encoded representation of the institution's logo
	Logo NullableString `json:"logo,omitempty"`
	// A partial list of routing numbers associated with the institution. This list is provided for the purpose of looking up institutions by routing number. It is not comprehensive and should never be used as a complete list of routing numbers for an institution.
	RoutingNumbers []string `json:"routing_numbers"`
	// Indicates that the institution has a mandatory OAuth login flow. Note that `oauth` may be `false` even for institutions that support OAuth, if the institution is in the process of migrating to OAuth and some active Items still exist that do not use OAuth.
	Oauth bool `json:"oauth"`
	Status NullableInstitutionStatus `json:"status,omitempty"`
	PaymentInitiationMetadata NullablePaymentInitiationMetadata `json:"payment_initiation_metadata,omitempty"`
	AuthMetadata NullableAuthMetadata `json:"auth_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Institution Institution

// NewInstitution instantiates a new Institution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitution(institutionId string, name string, products []Products, countryCodes []CountryCode, routingNumbers []string, oauth bool) *Institution {
	this := Institution{}
	this.InstitutionId = institutionId
	this.Name = name
	this.Products = products
	this.CountryCodes = countryCodes
	this.RoutingNumbers = routingNumbers
	this.Oauth = oauth
	return &this
}

// NewInstitutionWithDefaults instantiates a new Institution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionWithDefaults() *Institution {
	this := Institution{}
	return &this
}

// GetInstitutionId returns the InstitutionId field value
func (o *Institution) GetInstitutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstitutionId
}

// GetInstitutionIdOk returns a tuple with the InstitutionId field value
// and a boolean to check if the value has been set.
func (o *Institution) GetInstitutionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstitutionId, true
}

// SetInstitutionId sets field value
func (o *Institution) SetInstitutionId(v string) {
	o.InstitutionId = v
}

// GetName returns the Name field value
func (o *Institution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Institution) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Institution) SetName(v string) {
	o.Name = v
}

// GetProducts returns the Products field value
func (o *Institution) GetProducts() []Products {
	if o == nil {
		var ret []Products
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *Institution) GetProductsOk() (*[]Products, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Products, true
}

// SetProducts sets field value
func (o *Institution) SetProducts(v []Products) {
	o.Products = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *Institution) GetCountryCodes() []CountryCode {
	if o == nil {
		var ret []CountryCode
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *Institution) GetCountryCodesOk() (*[]CountryCode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *Institution) SetCountryCodes(v []CountryCode) {
	o.CountryCodes = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Institution) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Institution) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Institution) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Institution) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Institution) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Institution) UnsetUrl() {
	o.Url.Unset()
}

// GetPrimaryColor returns the PrimaryColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Institution) GetPrimaryColor() string {
	if o == nil || o.PrimaryColor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryColor.Get()
}

// GetPrimaryColorOk returns a tuple with the PrimaryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Institution) GetPrimaryColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrimaryColor.Get(), o.PrimaryColor.IsSet()
}

// HasPrimaryColor returns a boolean if a field has been set.
func (o *Institution) HasPrimaryColor() bool {
	if o != nil && o.PrimaryColor.IsSet() {
		return true
	}

	return false
}

// SetPrimaryColor gets a reference to the given NullableString and assigns it to the PrimaryColor field.
func (o *Institution) SetPrimaryColor(v string) {
	o.PrimaryColor.Set(&v)
}
// SetPrimaryColorNil sets the value for PrimaryColor to be an explicit nil
func (o *Institution) SetPrimaryColorNil() {
	o.PrimaryColor.Set(nil)
}

// UnsetPrimaryColor ensures that no value is present for PrimaryColor, not even an explicit nil
func (o *Institution) UnsetPrimaryColor() {
	o.PrimaryColor.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Institution) GetLogo() string {
	if o == nil || o.Logo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Institution) GetLogoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *Institution) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *Institution) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *Institution) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *Institution) UnsetLogo() {
	o.Logo.Unset()
}

// GetRoutingNumbers returns the RoutingNumbers field value
func (o *Institution) GetRoutingNumbers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RoutingNumbers
}

// GetRoutingNumbersOk returns a tuple with the RoutingNumbers field value
// and a boolean to check if the value has been set.
func (o *Institution) GetRoutingNumbersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingNumbers, true
}

// SetRoutingNumbers sets field value
func (o *Institution) SetRoutingNumbers(v []string) {
	o.RoutingNumbers = v
}

// GetOauth returns the Oauth field value
func (o *Institution) GetOauth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value
// and a boolean to check if the value has been set.
func (o *Institution) GetOauthOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Oauth, true
}

// SetOauth sets field value
func (o *Institution) SetOauth(v bool) {
	o.Oauth = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Institution) GetStatus() InstitutionStatus {
	if o == nil || o.Status.Get() == nil {
		var ret InstitutionStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Institution) GetStatusOk() (*InstitutionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Institution) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInstitutionStatus and assigns it to the Status field.
func (o *Institution) SetStatus(v InstitutionStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Institution) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Institution) UnsetStatus() {
	o.Status.Unset()
}

// GetPaymentInitiationMetadata returns the PaymentInitiationMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Institution) GetPaymentInitiationMetadata() PaymentInitiationMetadata {
	if o == nil || o.PaymentInitiationMetadata.Get() == nil {
		var ret PaymentInitiationMetadata
		return ret
	}
	return *o.PaymentInitiationMetadata.Get()
}

// GetPaymentInitiationMetadataOk returns a tuple with the PaymentInitiationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Institution) GetPaymentInitiationMetadataOk() (*PaymentInitiationMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentInitiationMetadata.Get(), o.PaymentInitiationMetadata.IsSet()
}

// HasPaymentInitiationMetadata returns a boolean if a field has been set.
func (o *Institution) HasPaymentInitiationMetadata() bool {
	if o != nil && o.PaymentInitiationMetadata.IsSet() {
		return true
	}

	return false
}

// SetPaymentInitiationMetadata gets a reference to the given NullablePaymentInitiationMetadata and assigns it to the PaymentInitiationMetadata field.
func (o *Institution) SetPaymentInitiationMetadata(v PaymentInitiationMetadata) {
	o.PaymentInitiationMetadata.Set(&v)
}
// SetPaymentInitiationMetadataNil sets the value for PaymentInitiationMetadata to be an explicit nil
func (o *Institution) SetPaymentInitiationMetadataNil() {
	o.PaymentInitiationMetadata.Set(nil)
}

// UnsetPaymentInitiationMetadata ensures that no value is present for PaymentInitiationMetadata, not even an explicit nil
func (o *Institution) UnsetPaymentInitiationMetadata() {
	o.PaymentInitiationMetadata.Unset()
}

// GetAuthMetadata returns the AuthMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Institution) GetAuthMetadata() AuthMetadata {
	if o == nil || o.AuthMetadata.Get() == nil {
		var ret AuthMetadata
		return ret
	}
	return *o.AuthMetadata.Get()
}

// GetAuthMetadataOk returns a tuple with the AuthMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Institution) GetAuthMetadataOk() (*AuthMetadata, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AuthMetadata.Get(), o.AuthMetadata.IsSet()
}

// HasAuthMetadata returns a boolean if a field has been set.
func (o *Institution) HasAuthMetadata() bool {
	if o != nil && o.AuthMetadata.IsSet() {
		return true
	}

	return false
}

// SetAuthMetadata gets a reference to the given NullableAuthMetadata and assigns it to the AuthMetadata field.
func (o *Institution) SetAuthMetadata(v AuthMetadata) {
	o.AuthMetadata.Set(&v)
}
// SetAuthMetadataNil sets the value for AuthMetadata to be an explicit nil
func (o *Institution) SetAuthMetadataNil() {
	o.AuthMetadata.Set(nil)
}

// UnsetAuthMetadata ensures that no value is present for AuthMetadata, not even an explicit nil
func (o *Institution) UnsetAuthMetadata() {
	o.AuthMetadata.Unset()
}

func (o Institution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["institution_id"] = o.InstitutionId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["country_codes"] = o.CountryCodes
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.PrimaryColor.IsSet() {
		toSerialize["primary_color"] = o.PrimaryColor.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if true {
		toSerialize["routing_numbers"] = o.RoutingNumbers
	}
	if true {
		toSerialize["oauth"] = o.Oauth
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.PaymentInitiationMetadata.IsSet() {
		toSerialize["payment_initiation_metadata"] = o.PaymentInitiationMetadata.Get()
	}
	if o.AuthMetadata.IsSet() {
		toSerialize["auth_metadata"] = o.AuthMetadata.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Institution) UnmarshalJSON(bytes []byte) (err error) {
	varInstitution := _Institution{}

	if err = json.Unmarshal(bytes, &varInstitution); err == nil {
		*o = Institution(varInstitution)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "institution_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "products")
		delete(additionalProperties, "country_codes")
		delete(additionalProperties, "url")
		delete(additionalProperties, "primary_color")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "routing_numbers")
		delete(additionalProperties, "oauth")
		delete(additionalProperties, "status")
		delete(additionalProperties, "payment_initiation_metadata")
		delete(additionalProperties, "auth_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitution struct {
	value *Institution
	isSet bool
}

func (v NullableInstitution) Get() *Institution {
	return v.value
}

func (v *NullableInstitution) Set(val *Institution) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitution(val *Institution) *NullableInstitution {
	return &NullableInstitution{value: val, isSet: true}
}

func (v NullableInstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


