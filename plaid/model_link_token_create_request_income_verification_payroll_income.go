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

// LinkTokenCreateRequestIncomeVerificationPayrollIncome Specifies options for initializing Link for use with Payroll Income.
type LinkTokenCreateRequestIncomeVerificationPayrollIncome struct {
	// The types of payroll income verification to enable for the Link session. If none are specified, then users will see both document and digital payroll income.
	FlowTypes []IncomeVerificationPayrollFlowType `json:"flow_types,omitempty"`
	// An identifier to indicate whether the income verification Link token will be used for an update or not
	IsUpdateMode *bool `json:"is_update_mode,omitempty"`
	// Uniquely identify a payroll income item to update with. It should only be used for update mode.
	ItemIdToUpdate NullableString `json:"item_id_to_update,omitempty"`
}

// NewLinkTokenCreateRequestIncomeVerificationPayrollIncome instantiates a new LinkTokenCreateRequestIncomeVerificationPayrollIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestIncomeVerificationPayrollIncome() *LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	this := LinkTokenCreateRequestIncomeVerificationPayrollIncome{}
	var isUpdateMode bool = false
	this.IsUpdateMode = &isUpdateMode
	return &this
}

// NewLinkTokenCreateRequestIncomeVerificationPayrollIncomeWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerificationPayrollIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestIncomeVerificationPayrollIncomeWithDefaults() *LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	this := LinkTokenCreateRequestIncomeVerificationPayrollIncome{}
	var isUpdateMode bool = false
	this.IsUpdateMode = &isUpdateMode
	return &this
}

// GetFlowTypes returns the FlowTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetFlowTypes() []IncomeVerificationPayrollFlowType {
	if o == nil  {
		var ret []IncomeVerificationPayrollFlowType
		return ret
	}
	return o.FlowTypes
}

// GetFlowTypesOk returns a tuple with the FlowTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetFlowTypesOk() (*[]IncomeVerificationPayrollFlowType, bool) {
	if o == nil || o.FlowTypes == nil {
		return nil, false
	}
	return &o.FlowTypes, true
}

// HasFlowTypes returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) HasFlowTypes() bool {
	if o != nil && o.FlowTypes != nil {
		return true
	}

	return false
}

// SetFlowTypes gets a reference to the given []IncomeVerificationPayrollFlowType and assigns it to the FlowTypes field.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) SetFlowTypes(v []IncomeVerificationPayrollFlowType) {
	o.FlowTypes = v
}

// GetIsUpdateMode returns the IsUpdateMode field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetIsUpdateMode() bool {
	if o == nil || o.IsUpdateMode == nil {
		var ret bool
		return ret
	}
	return *o.IsUpdateMode
}

// GetIsUpdateModeOk returns a tuple with the IsUpdateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetIsUpdateModeOk() (*bool, bool) {
	if o == nil || o.IsUpdateMode == nil {
		return nil, false
	}
	return o.IsUpdateMode, true
}

// HasIsUpdateMode returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) HasIsUpdateMode() bool {
	if o != nil && o.IsUpdateMode != nil {
		return true
	}

	return false
}

// SetIsUpdateMode gets a reference to the given bool and assigns it to the IsUpdateMode field.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) SetIsUpdateMode(v bool) {
	o.IsUpdateMode = &v
}

// GetItemIdToUpdate returns the ItemIdToUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetItemIdToUpdate() string {
	if o == nil || o.ItemIdToUpdate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ItemIdToUpdate.Get()
}

// GetItemIdToUpdateOk returns a tuple with the ItemIdToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) GetItemIdToUpdateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ItemIdToUpdate.Get(), o.ItemIdToUpdate.IsSet()
}

// HasItemIdToUpdate returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) HasItemIdToUpdate() bool {
	if o != nil && o.ItemIdToUpdate.IsSet() {
		return true
	}

	return false
}

// SetItemIdToUpdate gets a reference to the given NullableString and assigns it to the ItemIdToUpdate field.
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) SetItemIdToUpdate(v string) {
	o.ItemIdToUpdate.Set(&v)
}
// SetItemIdToUpdateNil sets the value for ItemIdToUpdate to be an explicit nil
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) SetItemIdToUpdateNil() {
	o.ItemIdToUpdate.Set(nil)
}

// UnsetItemIdToUpdate ensures that no value is present for ItemIdToUpdate, not even an explicit nil
func (o *LinkTokenCreateRequestIncomeVerificationPayrollIncome) UnsetItemIdToUpdate() {
	o.ItemIdToUpdate.Unset()
}

func (o LinkTokenCreateRequestIncomeVerificationPayrollIncome) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FlowTypes != nil {
		toSerialize["flow_types"] = o.FlowTypes
	}
	if o.IsUpdateMode != nil {
		toSerialize["is_update_mode"] = o.IsUpdateMode
	}
	if o.ItemIdToUpdate.IsSet() {
		toSerialize["item_id_to_update"] = o.ItemIdToUpdate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome struct {
	value *LinkTokenCreateRequestIncomeVerificationPayrollIncome
	isSet bool
}

func (v NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) Get() *LinkTokenCreateRequestIncomeVerificationPayrollIncome {
	return v.value
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) Set(val *LinkTokenCreateRequestIncomeVerificationPayrollIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestIncomeVerificationPayrollIncome(val *LinkTokenCreateRequestIncomeVerificationPayrollIncome) *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome {
	return &NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestIncomeVerificationPayrollIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


