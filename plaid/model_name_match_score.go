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

// NameMatchScore Score found by matching name provided by the API with the name on the account at the financial institution. If the account contains multiple owners, the maximum match score is filled.
type NameMatchScore struct {
	// Represents the match score for name. 100 is a perfect score, 85-99 means a strong match, 50-84 is a partial match, less than 50 is a weak match and 0 is a complete mismatch. If the name is missing from either the API or financial institution, this is empty.
	Score NullableInt32 `json:"score,omitempty"`
	// first or last name completely matched, likely a family member
	IsFirstNameOrLastNameMatch NullableBool `json:"is_first_name_or_last_name_match,omitempty"`
	// nickname matched, example Jennifer and Jenn.
	IsNicknameMatch NullableBool `json:"is_nickname_match,omitempty"`
	// Is `true` if the name on either of the names that was matched for the score contained strings indicative of a business name, such as \"CORP\", \"LLC\", \"INC\", or \"LTD\". A `true` result generally indicates the entity is a business. However, a `false` result does not mean the entity is not a business, as some businesses do not use these strings in the names used for their financial institution accounts.
	IsBusinessNameDetected NullableBool `json:"is_business_name_detected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NameMatchScore NameMatchScore

// NewNameMatchScore instantiates a new NameMatchScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameMatchScore() *NameMatchScore {
	this := NameMatchScore{}
	return &this
}

// NewNameMatchScoreWithDefaults instantiates a new NameMatchScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameMatchScoreWithDefaults() *NameMatchScore {
	this := NameMatchScore{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameMatchScore) GetScore() int32 {
	if o == nil || o.Score.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameMatchScore) GetScoreOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *NameMatchScore) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableInt32 and assigns it to the Score field.
func (o *NameMatchScore) SetScore(v int32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *NameMatchScore) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *NameMatchScore) UnsetScore() {
	o.Score.Unset()
}

// GetIsFirstNameOrLastNameMatch returns the IsFirstNameOrLastNameMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameMatchScore) GetIsFirstNameOrLastNameMatch() bool {
	if o == nil || o.IsFirstNameOrLastNameMatch.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsFirstNameOrLastNameMatch.Get()
}

// GetIsFirstNameOrLastNameMatchOk returns a tuple with the IsFirstNameOrLastNameMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameMatchScore) GetIsFirstNameOrLastNameMatchOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsFirstNameOrLastNameMatch.Get(), o.IsFirstNameOrLastNameMatch.IsSet()
}

// HasIsFirstNameOrLastNameMatch returns a boolean if a field has been set.
func (o *NameMatchScore) HasIsFirstNameOrLastNameMatch() bool {
	if o != nil && o.IsFirstNameOrLastNameMatch.IsSet() {
		return true
	}

	return false
}

// SetIsFirstNameOrLastNameMatch gets a reference to the given NullableBool and assigns it to the IsFirstNameOrLastNameMatch field.
func (o *NameMatchScore) SetIsFirstNameOrLastNameMatch(v bool) {
	o.IsFirstNameOrLastNameMatch.Set(&v)
}
// SetIsFirstNameOrLastNameMatchNil sets the value for IsFirstNameOrLastNameMatch to be an explicit nil
func (o *NameMatchScore) SetIsFirstNameOrLastNameMatchNil() {
	o.IsFirstNameOrLastNameMatch.Set(nil)
}

// UnsetIsFirstNameOrLastNameMatch ensures that no value is present for IsFirstNameOrLastNameMatch, not even an explicit nil
func (o *NameMatchScore) UnsetIsFirstNameOrLastNameMatch() {
	o.IsFirstNameOrLastNameMatch.Unset()
}

// GetIsNicknameMatch returns the IsNicknameMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameMatchScore) GetIsNicknameMatch() bool {
	if o == nil || o.IsNicknameMatch.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsNicknameMatch.Get()
}

// GetIsNicknameMatchOk returns a tuple with the IsNicknameMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameMatchScore) GetIsNicknameMatchOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsNicknameMatch.Get(), o.IsNicknameMatch.IsSet()
}

// HasIsNicknameMatch returns a boolean if a field has been set.
func (o *NameMatchScore) HasIsNicknameMatch() bool {
	if o != nil && o.IsNicknameMatch.IsSet() {
		return true
	}

	return false
}

// SetIsNicknameMatch gets a reference to the given NullableBool and assigns it to the IsNicknameMatch field.
func (o *NameMatchScore) SetIsNicknameMatch(v bool) {
	o.IsNicknameMatch.Set(&v)
}
// SetIsNicknameMatchNil sets the value for IsNicknameMatch to be an explicit nil
func (o *NameMatchScore) SetIsNicknameMatchNil() {
	o.IsNicknameMatch.Set(nil)
}

// UnsetIsNicknameMatch ensures that no value is present for IsNicknameMatch, not even an explicit nil
func (o *NameMatchScore) UnsetIsNicknameMatch() {
	o.IsNicknameMatch.Unset()
}

// GetIsBusinessNameDetected returns the IsBusinessNameDetected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameMatchScore) GetIsBusinessNameDetected() bool {
	if o == nil || o.IsBusinessNameDetected.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsBusinessNameDetected.Get()
}

// GetIsBusinessNameDetectedOk returns a tuple with the IsBusinessNameDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameMatchScore) GetIsBusinessNameDetectedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsBusinessNameDetected.Get(), o.IsBusinessNameDetected.IsSet()
}

// HasIsBusinessNameDetected returns a boolean if a field has been set.
func (o *NameMatchScore) HasIsBusinessNameDetected() bool {
	if o != nil && o.IsBusinessNameDetected.IsSet() {
		return true
	}

	return false
}

// SetIsBusinessNameDetected gets a reference to the given NullableBool and assigns it to the IsBusinessNameDetected field.
func (o *NameMatchScore) SetIsBusinessNameDetected(v bool) {
	o.IsBusinessNameDetected.Set(&v)
}
// SetIsBusinessNameDetectedNil sets the value for IsBusinessNameDetected to be an explicit nil
func (o *NameMatchScore) SetIsBusinessNameDetectedNil() {
	o.IsBusinessNameDetected.Set(nil)
}

// UnsetIsBusinessNameDetected ensures that no value is present for IsBusinessNameDetected, not even an explicit nil
func (o *NameMatchScore) UnsetIsBusinessNameDetected() {
	o.IsBusinessNameDetected.Unset()
}

func (o NameMatchScore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	if o.IsFirstNameOrLastNameMatch.IsSet() {
		toSerialize["is_first_name_or_last_name_match"] = o.IsFirstNameOrLastNameMatch.Get()
	}
	if o.IsNicknameMatch.IsSet() {
		toSerialize["is_nickname_match"] = o.IsNicknameMatch.Get()
	}
	if o.IsBusinessNameDetected.IsSet() {
		toSerialize["is_business_name_detected"] = o.IsBusinessNameDetected.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NameMatchScore) UnmarshalJSON(bytes []byte) (err error) {
	varNameMatchScore := _NameMatchScore{}

	if err = json.Unmarshal(bytes, &varNameMatchScore); err == nil {
		*o = NameMatchScore(varNameMatchScore)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "score")
		delete(additionalProperties, "is_first_name_or_last_name_match")
		delete(additionalProperties, "is_nickname_match")
		delete(additionalProperties, "is_business_name_detected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNameMatchScore struct {
	value *NameMatchScore
	isSet bool
}

func (v NullableNameMatchScore) Get() *NameMatchScore {
	return v.value
}

func (v *NullableNameMatchScore) Set(val *NameMatchScore) {
	v.value = val
	v.isSet = true
}

func (v NullableNameMatchScore) IsSet() bool {
	return v.isSet
}

func (v *NullableNameMatchScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameMatchScore(val *NameMatchScore) *NullableNameMatchScore {
	return &NullableNameMatchScore{value: val, isSet: true}
}

func (v NullableNameMatchScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameMatchScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


