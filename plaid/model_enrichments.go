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

// Enrichments A grouping of the Plaid produced transaction enrichment fields.
type Enrichments struct {
	// The check number of the transaction. This field is only populated for check transactions.
	CheckNumber NullableString `json:"check_number,omitempty"`
	// The counterparties present in the transaction. Counterparties, such as the merchant or the financial institution, are extracted by Plaid from the raw description.
	Counterparties []Counterparty `json:"counterparties"`
	// A unique, stable, Plaid-generated id that maps to the primary counterparty.
	EntityId NullableString `json:"entity_id,omitempty"`
	// The ID of the legacy category to which this transaction belongs. For a full list of legacy categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).  We recommend using the `personal_finance_category` for transaction categorization to obtain the best results.
	LegacyCategoryId NullableString `json:"legacy_category_id,omitempty"`
	// A hierarchical array of the legacy categories to which this transaction belongs. For a full list of legacy categories, see [`/categories/get`](https://plaid.com/docs/api/products/transactions/#categoriesget).  We recommend using the `personal_finance_category` for transaction categorization to obtain the best results.
	LegacyCategory *[]string `json:"legacy_category,omitempty"`
	Location Location `json:"location"`
	// The URL of a logo associated with this transaction, if available. The logo is formatted as a 100x100 pixel PNG file.
	LogoUrl NullableString `json:"logo_url"`
	// The name of the primary counterparty, such as the merchant or the financial institution, as extracted by Plaid from the raw description.
	MerchantName NullableString `json:"merchant_name"`
	PaymentChannel PaymentChannel `json:"payment_channel"`
	PersonalFinanceCategory NullablePersonalFinanceCategory `json:"personal_finance_category"`
	// A link to the icon associated with the primary personal finance category. The logo will always be 100x100 pixels.
	PersonalFinanceCategoryIconUrl string `json:"personal_finance_category_icon_url"`
	Recurrence NullableRecurrence `json:"recurrence,omitempty"`
	// The website associated with this transaction.
	Website NullableString `json:"website"`
	AdditionalProperties map[string]interface{}
}

type _Enrichments Enrichments

// NewEnrichments instantiates a new Enrichments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichments(counterparties []Counterparty, location Location, logoUrl NullableString, merchantName NullableString, paymentChannel PaymentChannel, personalFinanceCategory NullablePersonalFinanceCategory, personalFinanceCategoryIconUrl string, website NullableString) *Enrichments {
	this := Enrichments{}
	this.Counterparties = counterparties
	this.Location = location
	this.LogoUrl = logoUrl
	this.MerchantName = merchantName
	this.PaymentChannel = paymentChannel
	this.PersonalFinanceCategory = personalFinanceCategory
	this.PersonalFinanceCategoryIconUrl = personalFinanceCategoryIconUrl
	this.Website = website
	return &this
}

// NewEnrichmentsWithDefaults instantiates a new Enrichments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichmentsWithDefaults() *Enrichments {
	this := Enrichments{}
	return &this
}

// GetCheckNumber returns the CheckNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enrichments) GetCheckNumber() string {
	if o == nil || o.CheckNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.CheckNumber.Get()
}

// GetCheckNumberOk returns a tuple with the CheckNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetCheckNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CheckNumber.Get(), o.CheckNumber.IsSet()
}

// HasCheckNumber returns a boolean if a field has been set.
func (o *Enrichments) HasCheckNumber() bool {
	if o != nil && o.CheckNumber.IsSet() {
		return true
	}

	return false
}

// SetCheckNumber gets a reference to the given NullableString and assigns it to the CheckNumber field.
func (o *Enrichments) SetCheckNumber(v string) {
	o.CheckNumber.Set(&v)
}
// SetCheckNumberNil sets the value for CheckNumber to be an explicit nil
func (o *Enrichments) SetCheckNumberNil() {
	o.CheckNumber.Set(nil)
}

// UnsetCheckNumber ensures that no value is present for CheckNumber, not even an explicit nil
func (o *Enrichments) UnsetCheckNumber() {
	o.CheckNumber.Unset()
}

// GetCounterparties returns the Counterparties field value
func (o *Enrichments) GetCounterparties() []Counterparty {
	if o == nil {
		var ret []Counterparty
		return ret
	}

	return o.Counterparties
}

// GetCounterpartiesOk returns a tuple with the Counterparties field value
// and a boolean to check if the value has been set.
func (o *Enrichments) GetCounterpartiesOk() (*[]Counterparty, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Counterparties, true
}

// SetCounterparties sets field value
func (o *Enrichments) SetCounterparties(v []Counterparty) {
	o.Counterparties = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enrichments) GetEntityId() string {
	if o == nil || o.EntityId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EntityId.Get()
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetEntityIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EntityId.Get(), o.EntityId.IsSet()
}

// HasEntityId returns a boolean if a field has been set.
func (o *Enrichments) HasEntityId() bool {
	if o != nil && o.EntityId.IsSet() {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given NullableString and assigns it to the EntityId field.
func (o *Enrichments) SetEntityId(v string) {
	o.EntityId.Set(&v)
}
// SetEntityIdNil sets the value for EntityId to be an explicit nil
func (o *Enrichments) SetEntityIdNil() {
	o.EntityId.Set(nil)
}

// UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
func (o *Enrichments) UnsetEntityId() {
	o.EntityId.Unset()
}

// GetLegacyCategoryId returns the LegacyCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enrichments) GetLegacyCategoryId() string {
	if o == nil || o.LegacyCategoryId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LegacyCategoryId.Get()
}

// GetLegacyCategoryIdOk returns a tuple with the LegacyCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetLegacyCategoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LegacyCategoryId.Get(), o.LegacyCategoryId.IsSet()
}

// HasLegacyCategoryId returns a boolean if a field has been set.
func (o *Enrichments) HasLegacyCategoryId() bool {
	if o != nil && o.LegacyCategoryId.IsSet() {
		return true
	}

	return false
}

// SetLegacyCategoryId gets a reference to the given NullableString and assigns it to the LegacyCategoryId field.
func (o *Enrichments) SetLegacyCategoryId(v string) {
	o.LegacyCategoryId.Set(&v)
}
// SetLegacyCategoryIdNil sets the value for LegacyCategoryId to be an explicit nil
func (o *Enrichments) SetLegacyCategoryIdNil() {
	o.LegacyCategoryId.Set(nil)
}

// UnsetLegacyCategoryId ensures that no value is present for LegacyCategoryId, not even an explicit nil
func (o *Enrichments) UnsetLegacyCategoryId() {
	o.LegacyCategoryId.Unset()
}

// GetLegacyCategory returns the LegacyCategory field value if set, zero value otherwise.
func (o *Enrichments) GetLegacyCategory() []string {
	if o == nil || o.LegacyCategory == nil {
		var ret []string
		return ret
	}
	return *o.LegacyCategory
}

// GetLegacyCategoryOk returns a tuple with the LegacyCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Enrichments) GetLegacyCategoryOk() (*[]string, bool) {
	if o == nil || o.LegacyCategory == nil {
		return nil, false
	}
	return o.LegacyCategory, true
}

// HasLegacyCategory returns a boolean if a field has been set.
func (o *Enrichments) HasLegacyCategory() bool {
	if o != nil && o.LegacyCategory != nil {
		return true
	}

	return false
}

// SetLegacyCategory gets a reference to the given []string and assigns it to the LegacyCategory field.
func (o *Enrichments) SetLegacyCategory(v []string) {
	o.LegacyCategory = &v
}

// GetLocation returns the Location field value
func (o *Enrichments) GetLocation() Location {
	if o == nil {
		var ret Location
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Enrichments) GetLocationOk() (*Location, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Enrichments) SetLocation(v Location) {
	o.Location = v
}

// GetLogoUrl returns the LogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Enrichments) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// SetLogoUrl sets field value
func (o *Enrichments) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// GetMerchantName returns the MerchantName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Enrichments) GetMerchantName() string {
	if o == nil || o.MerchantName.Get() == nil {
		var ret string
		return ret
	}

	return *o.MerchantName.Get()
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetMerchantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantName.Get(), o.MerchantName.IsSet()
}

// SetMerchantName sets field value
func (o *Enrichments) SetMerchantName(v string) {
	o.MerchantName.Set(&v)
}

// GetPaymentChannel returns the PaymentChannel field value
func (o *Enrichments) GetPaymentChannel() PaymentChannel {
	if o == nil {
		var ret PaymentChannel
		return ret
	}

	return o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value
// and a boolean to check if the value has been set.
func (o *Enrichments) GetPaymentChannelOk() (*PaymentChannel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentChannel, true
}

// SetPaymentChannel sets field value
func (o *Enrichments) SetPaymentChannel(v PaymentChannel) {
	o.PaymentChannel = v
}

// GetPersonalFinanceCategory returns the PersonalFinanceCategory field value
// If the value is explicit nil, the zero value for PersonalFinanceCategory will be returned
func (o *Enrichments) GetPersonalFinanceCategory() PersonalFinanceCategory {
	if o == nil || o.PersonalFinanceCategory.Get() == nil {
		var ret PersonalFinanceCategory
		return ret
	}

	return *o.PersonalFinanceCategory.Get()
}

// GetPersonalFinanceCategoryOk returns a tuple with the PersonalFinanceCategory field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetPersonalFinanceCategoryOk() (*PersonalFinanceCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalFinanceCategory.Get(), o.PersonalFinanceCategory.IsSet()
}

// SetPersonalFinanceCategory sets field value
func (o *Enrichments) SetPersonalFinanceCategory(v PersonalFinanceCategory) {
	o.PersonalFinanceCategory.Set(&v)
}

// GetPersonalFinanceCategoryIconUrl returns the PersonalFinanceCategoryIconUrl field value
func (o *Enrichments) GetPersonalFinanceCategoryIconUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PersonalFinanceCategoryIconUrl
}

// GetPersonalFinanceCategoryIconUrlOk returns a tuple with the PersonalFinanceCategoryIconUrl field value
// and a boolean to check if the value has been set.
func (o *Enrichments) GetPersonalFinanceCategoryIconUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PersonalFinanceCategoryIconUrl, true
}

// SetPersonalFinanceCategoryIconUrl sets field value
func (o *Enrichments) SetPersonalFinanceCategoryIconUrl(v string) {
	o.PersonalFinanceCategoryIconUrl = v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Enrichments) GetRecurrence() Recurrence {
	if o == nil || o.Recurrence.Get() == nil {
		var ret Recurrence
		return ret
	}
	return *o.Recurrence.Get()
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetRecurrenceOk() (*Recurrence, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Recurrence.Get(), o.Recurrence.IsSet()
}

// HasRecurrence returns a boolean if a field has been set.
func (o *Enrichments) HasRecurrence() bool {
	if o != nil && o.Recurrence.IsSet() {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given NullableRecurrence and assigns it to the Recurrence field.
func (o *Enrichments) SetRecurrence(v Recurrence) {
	o.Recurrence.Set(&v)
}
// SetRecurrenceNil sets the value for Recurrence to be an explicit nil
func (o *Enrichments) SetRecurrenceNil() {
	o.Recurrence.Set(nil)
}

// UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
func (o *Enrichments) UnsetRecurrence() {
	o.Recurrence.Unset()
}

// GetWebsite returns the Website field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Enrichments) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}

	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Enrichments) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// SetWebsite sets field value
func (o *Enrichments) SetWebsite(v string) {
	o.Website.Set(&v)
}

func (o Enrichments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckNumber.IsSet() {
		toSerialize["check_number"] = o.CheckNumber.Get()
	}
	if true {
		toSerialize["counterparties"] = o.Counterparties
	}
	if o.EntityId.IsSet() {
		toSerialize["entity_id"] = o.EntityId.Get()
	}
	if o.LegacyCategoryId.IsSet() {
		toSerialize["legacy_category_id"] = o.LegacyCategoryId.Get()
	}
	if o.LegacyCategory != nil {
		toSerialize["legacy_category"] = o.LegacyCategory
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if true {
		toSerialize["merchant_name"] = o.MerchantName.Get()
	}
	if true {
		toSerialize["payment_channel"] = o.PaymentChannel
	}
	if true {
		toSerialize["personal_finance_category"] = o.PersonalFinanceCategory.Get()
	}
	if true {
		toSerialize["personal_finance_category_icon_url"] = o.PersonalFinanceCategoryIconUrl
	}
	if o.Recurrence.IsSet() {
		toSerialize["recurrence"] = o.Recurrence.Get()
	}
	if true {
		toSerialize["website"] = o.Website.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Enrichments) UnmarshalJSON(bytes []byte) (err error) {
	varEnrichments := _Enrichments{}

	if err = json.Unmarshal(bytes, &varEnrichments); err == nil {
		*o = Enrichments(varEnrichments)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "check_number")
		delete(additionalProperties, "counterparties")
		delete(additionalProperties, "entity_id")
		delete(additionalProperties, "legacy_category_id")
		delete(additionalProperties, "legacy_category")
		delete(additionalProperties, "location")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "merchant_name")
		delete(additionalProperties, "payment_channel")
		delete(additionalProperties, "personal_finance_category")
		delete(additionalProperties, "personal_finance_category_icon_url")
		delete(additionalProperties, "recurrence")
		delete(additionalProperties, "website")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnrichments struct {
	value *Enrichments
	isSet bool
}

func (v NullableEnrichments) Get() *Enrichments {
	return v.value
}

func (v *NullableEnrichments) Set(val *Enrichments) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichments) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichments(val *Enrichments) *NullableEnrichments {
	return &NullableEnrichments{value: val, isSet: true}
}

func (v NullableEnrichments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


