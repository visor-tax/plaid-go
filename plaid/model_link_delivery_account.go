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

// LinkDeliveryAccount Information related to account attached to the connected Item
type LinkDeliveryAccount struct {
	// The Plaid `account_id`
	Id *string `json:"id,omitempty"`
	// The official account name
	Name *string `json:"name,omitempty"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts. It may also not match the mask that the bank displays to the user.
	Mask *string `json:"mask,omitempty"`
	// The account type. See the [Account schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full list of possible values
	Type *string `json:"type,omitempty"`
	// The account subtype. See the [Account schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full list of possible values
	Subtype *string `json:"subtype,omitempty"`
	VerificationStatus *LinkDeliveryVerificationStatus `json:"verification_status,omitempty"`
	// If micro-deposit verification is being used, indicates whether the account being verified is a `business` or `personal` account.
	ClassType *string `json:"class_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkDeliveryAccount LinkDeliveryAccount

// NewLinkDeliveryAccount instantiates a new LinkDeliveryAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkDeliveryAccount() *LinkDeliveryAccount {
	this := LinkDeliveryAccount{}
	return &this
}

// NewLinkDeliveryAccountWithDefaults instantiates a new LinkDeliveryAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkDeliveryAccountWithDefaults() *LinkDeliveryAccount {
	this := LinkDeliveryAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LinkDeliveryAccount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinkDeliveryAccount) SetName(v string) {
	o.Name = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetMask() string {
	if o == nil || o.Mask == nil {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetMaskOk() (*string, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *LinkDeliveryAccount) SetMask(v string) {
	o.Mask = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LinkDeliveryAccount) SetType(v string) {
	o.Type = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *LinkDeliveryAccount) SetSubtype(v string) {
	o.Subtype = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetVerificationStatus() LinkDeliveryVerificationStatus {
	if o == nil || o.VerificationStatus == nil {
		var ret LinkDeliveryVerificationStatus
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetVerificationStatusOk() (*LinkDeliveryVerificationStatus, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given LinkDeliveryVerificationStatus and assigns it to the VerificationStatus field.
func (o *LinkDeliveryAccount) SetVerificationStatus(v LinkDeliveryVerificationStatus) {
	o.VerificationStatus = &v
}

// GetClassType returns the ClassType field value if set, zero value otherwise.
func (o *LinkDeliveryAccount) GetClassType() string {
	if o == nil || o.ClassType == nil {
		var ret string
		return ret
	}
	return *o.ClassType
}

// GetClassTypeOk returns a tuple with the ClassType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkDeliveryAccount) GetClassTypeOk() (*string, bool) {
	if o == nil || o.ClassType == nil {
		return nil, false
	}
	return o.ClassType, true
}

// HasClassType returns a boolean if a field has been set.
func (o *LinkDeliveryAccount) HasClassType() bool {
	if o != nil && o.ClassType != nil {
		return true
	}

	return false
}

// SetClassType gets a reference to the given string and assigns it to the ClassType field.
func (o *LinkDeliveryAccount) SetClassType(v string) {
	o.ClassType = &v
}

func (o LinkDeliveryAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Subtype != nil {
		toSerialize["subtype"] = o.Subtype
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if o.ClassType != nil {
		toSerialize["class_type"] = o.ClassType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkDeliveryAccount) UnmarshalJSON(bytes []byte) (err error) {
	varLinkDeliveryAccount := _LinkDeliveryAccount{}

	if err = json.Unmarshal(bytes, &varLinkDeliveryAccount); err == nil {
		*o = LinkDeliveryAccount(varLinkDeliveryAccount)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "verification_status")
		delete(additionalProperties, "class_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkDeliveryAccount struct {
	value *LinkDeliveryAccount
	isSet bool
}

func (v NullableLinkDeliveryAccount) Get() *LinkDeliveryAccount {
	return v.value
}

func (v *NullableLinkDeliveryAccount) Set(val *LinkDeliveryAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryAccount(val *LinkDeliveryAccount) *NullableLinkDeliveryAccount {
	return &NullableLinkDeliveryAccount{value: val, isSet: true}
}

func (v NullableLinkDeliveryAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


