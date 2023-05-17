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

// AccountBase A single account at a financial institution.
type AccountBase struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  If an account with a specific `account_id` disappears instead of changing, the account is likely closed. Closed accounts are not returned by the Plaid API.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances AccountBalance `json:"balances"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask"`
	// The name of the account, either assigned by the user or by the financial institution itself
	Name string `json:"name"`
	// The official name of the account as given by the financial institution
	OfficialName NullableString `json:"official_name"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The current verification status of an Auth Item initiated through Automated or Manual micro-deposits.  Returned for Auth Items only.  `pending_automatic_verification`: The Item is pending automatic verification  `pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the two amounts.  `automatically_verified`: The Item has successfully been automatically verified   `manually_verified`: The Item has successfully been manually verified  `verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  `verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.   
	VerificationStatus *string `json:"verification_status,omitempty"`
	// A unique and persistent identifier for accounts that can be used to trace multiple instances of the same account across different Items for depository accounts. This is currently an opt-in field and only supported for Chase Items.
	PersistentAccountId *string `json:"persistent_account_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountBase AccountBase

// NewAccountBase instantiates a new AccountBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBase(accountId string, balances AccountBalance, mask NullableString, name string, officialName NullableString, type_ AccountType, subtype NullableAccountSubtype) *AccountBase {
	this := AccountBase{}
	this.AccountId = accountId
	this.Balances = balances
	this.Mask = mask
	this.Name = name
	this.OfficialName = officialName
	this.Type = type_
	this.Subtype = subtype
	return &this
}

// NewAccountBaseWithDefaults instantiates a new AccountBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBaseWithDefaults() *AccountBase {
	this := AccountBase{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountBase) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountBase) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountBase) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *AccountBase) GetBalances() AccountBalance {
	if o == nil {
		var ret AccountBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *AccountBase) GetBalancesOk() (*AccountBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *AccountBase) SetBalances(v AccountBalance) {
	o.Balances = v
}

// GetMask returns the Mask field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountBase) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBase) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// SetMask sets field value
func (o *AccountBase) SetMask(v string) {
	o.Mask.Set(&v)
}

// GetName returns the Name field value
func (o *AccountBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountBase) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountBase) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountBase) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBase) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// SetOfficialName sets field value
func (o *AccountBase) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}

// GetType returns the Type field value
func (o *AccountBase) GetType() AccountType {
	if o == nil {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountBase) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountBase) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *AccountBase) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountBase) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *AccountBase) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountBase) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountBase) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountBase) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetPersistentAccountId returns the PersistentAccountId field value if set, zero value otherwise.
func (o *AccountBase) GetPersistentAccountId() string {
	if o == nil || o.PersistentAccountId == nil {
		var ret string
		return ret
	}
	return *o.PersistentAccountId
}

// GetPersistentAccountIdOk returns a tuple with the PersistentAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBase) GetPersistentAccountIdOk() (*string, bool) {
	if o == nil || o.PersistentAccountId == nil {
		return nil, false
	}
	return o.PersistentAccountId, true
}

// HasPersistentAccountId returns a boolean if a field has been set.
func (o *AccountBase) HasPersistentAccountId() bool {
	if o != nil && o.PersistentAccountId != nil {
		return true
	}

	return false
}

// SetPersistentAccountId gets a reference to the given string and assigns it to the PersistentAccountId field.
func (o *AccountBase) SetPersistentAccountId(v string) {
	o.PersistentAccountId = &v
}

func (o AccountBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if true {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["official_name"] = o.OfficialName.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.VerificationStatus != nil {
		toSerialize["verification_status"] = o.VerificationStatus
	}
	if o.PersistentAccountId != nil {
		toSerialize["persistent_account_id"] = o.PersistentAccountId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountBase) UnmarshalJSON(bytes []byte) (err error) {
	varAccountBase := _AccountBase{}

	if err = json.Unmarshal(bytes, &varAccountBase); err == nil {
		*o = AccountBase(varAccountBase)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "balances")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "name")
		delete(additionalProperties, "official_name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "verification_status")
		delete(additionalProperties, "persistent_account_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountBase struct {
	value *AccountBase
	isSet bool
}

func (v NullableAccountBase) Get() *AccountBase {
	return v.value
}

func (v *NullableAccountBase) Set(val *AccountBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBase(val *AccountBase) *NullableAccountBase {
	return &NullableAccountBase{value: val, isSet: true}
}

func (v NullableAccountBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


