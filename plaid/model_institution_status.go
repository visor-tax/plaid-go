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

// InstitutionStatus The status of an institution is determined by the health of its Item logins, Transactions updates, Investments updates, Liabilities updates, Auth requests, Balance requests, Identity requests, Investments requests, and Liabilities requests. A login attempt is conducted during the initial Item add in Link. If there is not enough traffic to accurately calculate an institution's status, Plaid will return null rather than potentially inaccurate data.  Institution status is accessible in the Dashboard and via the API using the `/institutions/get_by_id` endpoint with the `include_status` option set to true. Note that institution status is not available in the Sandbox environment. 
type InstitutionStatus struct {
	ItemLogins NullableProductStatus `json:"item_logins,omitempty"`
	TransactionsUpdates NullableProductStatus `json:"transactions_updates,omitempty"`
	Auth NullableProductStatus `json:"auth,omitempty"`
	Identity NullableProductStatus `json:"identity,omitempty"`
	InvestmentsUpdates NullableProductStatus `json:"investments_updates,omitempty"`
	LiabilitiesUpdates NullableProductStatus `json:"liabilities_updates,omitempty"`
	Liabilities NullableProductStatus `json:"liabilities,omitempty"`
	Investments NullableProductStatus `json:"investments,omitempty"`
	// Details of recent health incidents associated with the institution.
	HealthIncidents []HealthIncident `json:"health_incidents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InstitutionStatus InstitutionStatus

// NewInstitutionStatus instantiates a new InstitutionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionStatus() *InstitutionStatus {
	this := InstitutionStatus{}
	return &this
}

// NewInstitutionStatusWithDefaults instantiates a new InstitutionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionStatusWithDefaults() *InstitutionStatus {
	this := InstitutionStatus{}
	return &this
}

// GetItemLogins returns the ItemLogins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetItemLogins() ProductStatus {
	if o == nil || o.ItemLogins.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.ItemLogins.Get()
}

// GetItemLoginsOk returns a tuple with the ItemLogins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetItemLoginsOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ItemLogins.Get(), o.ItemLogins.IsSet()
}

// HasItemLogins returns a boolean if a field has been set.
func (o *InstitutionStatus) HasItemLogins() bool {
	if o != nil && o.ItemLogins.IsSet() {
		return true
	}

	return false
}

// SetItemLogins gets a reference to the given NullableProductStatus and assigns it to the ItemLogins field.
func (o *InstitutionStatus) SetItemLogins(v ProductStatus) {
	o.ItemLogins.Set(&v)
}
// SetItemLoginsNil sets the value for ItemLogins to be an explicit nil
func (o *InstitutionStatus) SetItemLoginsNil() {
	o.ItemLogins.Set(nil)
}

// UnsetItemLogins ensures that no value is present for ItemLogins, not even an explicit nil
func (o *InstitutionStatus) UnsetItemLogins() {
	o.ItemLogins.Unset()
}

// GetTransactionsUpdates returns the TransactionsUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetTransactionsUpdates() ProductStatus {
	if o == nil || o.TransactionsUpdates.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.TransactionsUpdates.Get()
}

// GetTransactionsUpdatesOk returns a tuple with the TransactionsUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetTransactionsUpdatesOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionsUpdates.Get(), o.TransactionsUpdates.IsSet()
}

// HasTransactionsUpdates returns a boolean if a field has been set.
func (o *InstitutionStatus) HasTransactionsUpdates() bool {
	if o != nil && o.TransactionsUpdates.IsSet() {
		return true
	}

	return false
}

// SetTransactionsUpdates gets a reference to the given NullableProductStatus and assigns it to the TransactionsUpdates field.
func (o *InstitutionStatus) SetTransactionsUpdates(v ProductStatus) {
	o.TransactionsUpdates.Set(&v)
}
// SetTransactionsUpdatesNil sets the value for TransactionsUpdates to be an explicit nil
func (o *InstitutionStatus) SetTransactionsUpdatesNil() {
	o.TransactionsUpdates.Set(nil)
}

// UnsetTransactionsUpdates ensures that no value is present for TransactionsUpdates, not even an explicit nil
func (o *InstitutionStatus) UnsetTransactionsUpdates() {
	o.TransactionsUpdates.Unset()
}

// GetAuth returns the Auth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetAuth() ProductStatus {
	if o == nil || o.Auth.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.Auth.Get()
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetAuthOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Auth.Get(), o.Auth.IsSet()
}

// HasAuth returns a boolean if a field has been set.
func (o *InstitutionStatus) HasAuth() bool {
	if o != nil && o.Auth.IsSet() {
		return true
	}

	return false
}

// SetAuth gets a reference to the given NullableProductStatus and assigns it to the Auth field.
func (o *InstitutionStatus) SetAuth(v ProductStatus) {
	o.Auth.Set(&v)
}
// SetAuthNil sets the value for Auth to be an explicit nil
func (o *InstitutionStatus) SetAuthNil() {
	o.Auth.Set(nil)
}

// UnsetAuth ensures that no value is present for Auth, not even an explicit nil
func (o *InstitutionStatus) UnsetAuth() {
	o.Auth.Unset()
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetIdentity() ProductStatus {
	if o == nil || o.Identity.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetIdentityOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *InstitutionStatus) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableProductStatus and assigns it to the Identity field.
func (o *InstitutionStatus) SetIdentity(v ProductStatus) {
	o.Identity.Set(&v)
}
// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *InstitutionStatus) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *InstitutionStatus) UnsetIdentity() {
	o.Identity.Unset()
}

// GetInvestmentsUpdates returns the InvestmentsUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetInvestmentsUpdates() ProductStatus {
	if o == nil || o.InvestmentsUpdates.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.InvestmentsUpdates.Get()
}

// GetInvestmentsUpdatesOk returns a tuple with the InvestmentsUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetInvestmentsUpdatesOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InvestmentsUpdates.Get(), o.InvestmentsUpdates.IsSet()
}

// HasInvestmentsUpdates returns a boolean if a field has been set.
func (o *InstitutionStatus) HasInvestmentsUpdates() bool {
	if o != nil && o.InvestmentsUpdates.IsSet() {
		return true
	}

	return false
}

// SetInvestmentsUpdates gets a reference to the given NullableProductStatus and assigns it to the InvestmentsUpdates field.
func (o *InstitutionStatus) SetInvestmentsUpdates(v ProductStatus) {
	o.InvestmentsUpdates.Set(&v)
}
// SetInvestmentsUpdatesNil sets the value for InvestmentsUpdates to be an explicit nil
func (o *InstitutionStatus) SetInvestmentsUpdatesNil() {
	o.InvestmentsUpdates.Set(nil)
}

// UnsetInvestmentsUpdates ensures that no value is present for InvestmentsUpdates, not even an explicit nil
func (o *InstitutionStatus) UnsetInvestmentsUpdates() {
	o.InvestmentsUpdates.Unset()
}

// GetLiabilitiesUpdates returns the LiabilitiesUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetLiabilitiesUpdates() ProductStatus {
	if o == nil || o.LiabilitiesUpdates.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.LiabilitiesUpdates.Get()
}

// GetLiabilitiesUpdatesOk returns a tuple with the LiabilitiesUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetLiabilitiesUpdatesOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LiabilitiesUpdates.Get(), o.LiabilitiesUpdates.IsSet()
}

// HasLiabilitiesUpdates returns a boolean if a field has been set.
func (o *InstitutionStatus) HasLiabilitiesUpdates() bool {
	if o != nil && o.LiabilitiesUpdates.IsSet() {
		return true
	}

	return false
}

// SetLiabilitiesUpdates gets a reference to the given NullableProductStatus and assigns it to the LiabilitiesUpdates field.
func (o *InstitutionStatus) SetLiabilitiesUpdates(v ProductStatus) {
	o.LiabilitiesUpdates.Set(&v)
}
// SetLiabilitiesUpdatesNil sets the value for LiabilitiesUpdates to be an explicit nil
func (o *InstitutionStatus) SetLiabilitiesUpdatesNil() {
	o.LiabilitiesUpdates.Set(nil)
}

// UnsetLiabilitiesUpdates ensures that no value is present for LiabilitiesUpdates, not even an explicit nil
func (o *InstitutionStatus) UnsetLiabilitiesUpdates() {
	o.LiabilitiesUpdates.Unset()
}

// GetLiabilities returns the Liabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetLiabilities() ProductStatus {
	if o == nil || o.Liabilities.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.Liabilities.Get()
}

// GetLiabilitiesOk returns a tuple with the Liabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetLiabilitiesOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Liabilities.Get(), o.Liabilities.IsSet()
}

// HasLiabilities returns a boolean if a field has been set.
func (o *InstitutionStatus) HasLiabilities() bool {
	if o != nil && o.Liabilities.IsSet() {
		return true
	}

	return false
}

// SetLiabilities gets a reference to the given NullableProductStatus and assigns it to the Liabilities field.
func (o *InstitutionStatus) SetLiabilities(v ProductStatus) {
	o.Liabilities.Set(&v)
}
// SetLiabilitiesNil sets the value for Liabilities to be an explicit nil
func (o *InstitutionStatus) SetLiabilitiesNil() {
	o.Liabilities.Set(nil)
}

// UnsetLiabilities ensures that no value is present for Liabilities, not even an explicit nil
func (o *InstitutionStatus) UnsetLiabilities() {
	o.Liabilities.Unset()
}

// GetInvestments returns the Investments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetInvestments() ProductStatus {
	if o == nil || o.Investments.Get() == nil {
		var ret ProductStatus
		return ret
	}
	return *o.Investments.Get()
}

// GetInvestmentsOk returns a tuple with the Investments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetInvestmentsOk() (*ProductStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Investments.Get(), o.Investments.IsSet()
}

// HasInvestments returns a boolean if a field has been set.
func (o *InstitutionStatus) HasInvestments() bool {
	if o != nil && o.Investments.IsSet() {
		return true
	}

	return false
}

// SetInvestments gets a reference to the given NullableProductStatus and assigns it to the Investments field.
func (o *InstitutionStatus) SetInvestments(v ProductStatus) {
	o.Investments.Set(&v)
}
// SetInvestmentsNil sets the value for Investments to be an explicit nil
func (o *InstitutionStatus) SetInvestmentsNil() {
	o.Investments.Set(nil)
}

// UnsetInvestments ensures that no value is present for Investments, not even an explicit nil
func (o *InstitutionStatus) UnsetInvestments() {
	o.Investments.Unset()
}

// GetHealthIncidents returns the HealthIncidents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionStatus) GetHealthIncidents() []HealthIncident {
	if o == nil  {
		var ret []HealthIncident
		return ret
	}
	return o.HealthIncidents
}

// GetHealthIncidentsOk returns a tuple with the HealthIncidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionStatus) GetHealthIncidentsOk() (*[]HealthIncident, bool) {
	if o == nil || o.HealthIncidents == nil {
		return nil, false
	}
	return &o.HealthIncidents, true
}

// HasHealthIncidents returns a boolean if a field has been set.
func (o *InstitutionStatus) HasHealthIncidents() bool {
	if o != nil && o.HealthIncidents != nil {
		return true
	}

	return false
}

// SetHealthIncidents gets a reference to the given []HealthIncident and assigns it to the HealthIncidents field.
func (o *InstitutionStatus) SetHealthIncidents(v []HealthIncident) {
	o.HealthIncidents = v
}

func (o InstitutionStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemLogins.IsSet() {
		toSerialize["item_logins"] = o.ItemLogins.Get()
	}
	if o.TransactionsUpdates.IsSet() {
		toSerialize["transactions_updates"] = o.TransactionsUpdates.Get()
	}
	if o.Auth.IsSet() {
		toSerialize["auth"] = o.Auth.Get()
	}
	if o.Identity.IsSet() {
		toSerialize["identity"] = o.Identity.Get()
	}
	if o.InvestmentsUpdates.IsSet() {
		toSerialize["investments_updates"] = o.InvestmentsUpdates.Get()
	}
	if o.LiabilitiesUpdates.IsSet() {
		toSerialize["liabilities_updates"] = o.LiabilitiesUpdates.Get()
	}
	if o.Liabilities.IsSet() {
		toSerialize["liabilities"] = o.Liabilities.Get()
	}
	if o.Investments.IsSet() {
		toSerialize["investments"] = o.Investments.Get()
	}
	if o.HealthIncidents != nil {
		toSerialize["health_incidents"] = o.HealthIncidents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionStatus) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionStatus := _InstitutionStatus{}

	if err = json.Unmarshal(bytes, &varInstitutionStatus); err == nil {
		*o = InstitutionStatus(varInstitutionStatus)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_logins")
		delete(additionalProperties, "transactions_updates")
		delete(additionalProperties, "auth")
		delete(additionalProperties, "identity")
		delete(additionalProperties, "investments_updates")
		delete(additionalProperties, "liabilities_updates")
		delete(additionalProperties, "liabilities")
		delete(additionalProperties, "investments")
		delete(additionalProperties, "health_incidents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionStatus struct {
	value *InstitutionStatus
	isSet bool
}

func (v NullableInstitutionStatus) Get() *InstitutionStatus {
	return v.value
}

func (v *NullableInstitutionStatus) Set(val *InstitutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionStatus(val *InstitutionStatus) *NullableInstitutionStatus {
	return &NullableInstitutionStatus{value: val, isSet: true}
}

func (v NullableInstitutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


