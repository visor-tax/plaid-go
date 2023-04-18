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

// CreditFreddieMacVerificationOfAssetVOE25 Documentation not found in the MISMO model viewer and not provided by Freddie Mac.
type CreditFreddieMacVerificationOfAssetVOE25 struct {
	REPORTING_INFORMATION CreditFreddieMacReportingInformationVOA24 `json:"REPORTING_INFORMATION"`
	SERVICE_PRODUCT_FULFILLMENT ServiceProductFulfillment `json:"SERVICE_PRODUCT_FULFILLMENT"`
	VERIFICATION_OF_ASSET_RESPONSE CreditFreddieMacVerificationOfAssetResponseVOE25 `json:"VERIFICATION_OF_ASSET_RESPONSE"`
	AdditionalProperties map[string]interface{}
}

type _CreditFreddieMacVerificationOfAssetVOE25 CreditFreddieMacVerificationOfAssetVOE25

// NewCreditFreddieMacVerificationOfAssetVOE25 instantiates a new CreditFreddieMacVerificationOfAssetVOE25 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFreddieMacVerificationOfAssetVOE25(rEPORTINGINFORMATION CreditFreddieMacReportingInformationVOA24, sERVICEPRODUCTFULFILLMENT ServiceProductFulfillment, vERIFICATIONOFASSETRESPONSE CreditFreddieMacVerificationOfAssetResponseVOE25) *CreditFreddieMacVerificationOfAssetVOE25 {
	this := CreditFreddieMacVerificationOfAssetVOE25{}
	this.REPORTING_INFORMATION = rEPORTINGINFORMATION
	this.SERVICE_PRODUCT_FULFILLMENT = sERVICEPRODUCTFULFILLMENT
	this.VERIFICATION_OF_ASSET_RESPONSE = vERIFICATIONOFASSETRESPONSE
	return &this
}

// NewCreditFreddieMacVerificationOfAssetVOE25WithDefaults instantiates a new CreditFreddieMacVerificationOfAssetVOE25 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFreddieMacVerificationOfAssetVOE25WithDefaults() *CreditFreddieMacVerificationOfAssetVOE25 {
	this := CreditFreddieMacVerificationOfAssetVOE25{}
	return &this
}

// GetREPORTING_INFORMATION returns the REPORTING_INFORMATION field value
func (o *CreditFreddieMacVerificationOfAssetVOE25) GetREPORTING_INFORMATION() CreditFreddieMacReportingInformationVOA24 {
	if o == nil {
		var ret CreditFreddieMacReportingInformationVOA24
		return ret
	}

	return o.REPORTING_INFORMATION
}

// GetREPORTING_INFORMATIONOk returns a tuple with the REPORTING_INFORMATION field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetVOE25) GetREPORTING_INFORMATIONOk() (*CreditFreddieMacReportingInformationVOA24, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.REPORTING_INFORMATION, true
}

// SetREPORTING_INFORMATION sets field value
func (o *CreditFreddieMacVerificationOfAssetVOE25) SetREPORTING_INFORMATION(v CreditFreddieMacReportingInformationVOA24) {
	o.REPORTING_INFORMATION = v
}

// GetSERVICE_PRODUCT_FULFILLMENT returns the SERVICE_PRODUCT_FULFILLMENT field value
func (o *CreditFreddieMacVerificationOfAssetVOE25) GetSERVICE_PRODUCT_FULFILLMENT() ServiceProductFulfillment {
	if o == nil {
		var ret ServiceProductFulfillment
		return ret
	}

	return o.SERVICE_PRODUCT_FULFILLMENT
}

// GetSERVICE_PRODUCT_FULFILLMENTOk returns a tuple with the SERVICE_PRODUCT_FULFILLMENT field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetVOE25) GetSERVICE_PRODUCT_FULFILLMENTOk() (*ServiceProductFulfillment, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SERVICE_PRODUCT_FULFILLMENT, true
}

// SetSERVICE_PRODUCT_FULFILLMENT sets field value
func (o *CreditFreddieMacVerificationOfAssetVOE25) SetSERVICE_PRODUCT_FULFILLMENT(v ServiceProductFulfillment) {
	o.SERVICE_PRODUCT_FULFILLMENT = v
}

// GetVERIFICATION_OF_ASSET_RESPONSE returns the VERIFICATION_OF_ASSET_RESPONSE field value
func (o *CreditFreddieMacVerificationOfAssetVOE25) GetVERIFICATION_OF_ASSET_RESPONSE() CreditFreddieMacVerificationOfAssetResponseVOE25 {
	if o == nil {
		var ret CreditFreddieMacVerificationOfAssetResponseVOE25
		return ret
	}

	return o.VERIFICATION_OF_ASSET_RESPONSE
}

// GetVERIFICATION_OF_ASSET_RESPONSEOk returns a tuple with the VERIFICATION_OF_ASSET_RESPONSE field value
// and a boolean to check if the value has been set.
func (o *CreditFreddieMacVerificationOfAssetVOE25) GetVERIFICATION_OF_ASSET_RESPONSEOk() (*CreditFreddieMacVerificationOfAssetResponseVOE25, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VERIFICATION_OF_ASSET_RESPONSE, true
}

// SetVERIFICATION_OF_ASSET_RESPONSE sets field value
func (o *CreditFreddieMacVerificationOfAssetVOE25) SetVERIFICATION_OF_ASSET_RESPONSE(v CreditFreddieMacVerificationOfAssetResponseVOE25) {
	o.VERIFICATION_OF_ASSET_RESPONSE = v
}

func (o CreditFreddieMacVerificationOfAssetVOE25) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["REPORTING_INFORMATION"] = o.REPORTING_INFORMATION
	}
	if true {
		toSerialize["SERVICE_PRODUCT_FULFILLMENT"] = o.SERVICE_PRODUCT_FULFILLMENT
	}
	if true {
		toSerialize["VERIFICATION_OF_ASSET_RESPONSE"] = o.VERIFICATION_OF_ASSET_RESPONSE
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFreddieMacVerificationOfAssetVOE25) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFreddieMacVerificationOfAssetVOE25 := _CreditFreddieMacVerificationOfAssetVOE25{}

	if err = json.Unmarshal(bytes, &varCreditFreddieMacVerificationOfAssetVOE25); err == nil {
		*o = CreditFreddieMacVerificationOfAssetVOE25(varCreditFreddieMacVerificationOfAssetVOE25)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "REPORTING_INFORMATION")
		delete(additionalProperties, "SERVICE_PRODUCT_FULFILLMENT")
		delete(additionalProperties, "VERIFICATION_OF_ASSET_RESPONSE")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFreddieMacVerificationOfAssetVOE25 struct {
	value *CreditFreddieMacVerificationOfAssetVOE25
	isSet bool
}

func (v NullableCreditFreddieMacVerificationOfAssetVOE25) Get() *CreditFreddieMacVerificationOfAssetVOE25 {
	return v.value
}

func (v *NullableCreditFreddieMacVerificationOfAssetVOE25) Set(val *CreditFreddieMacVerificationOfAssetVOE25) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFreddieMacVerificationOfAssetVOE25) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFreddieMacVerificationOfAssetVOE25) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFreddieMacVerificationOfAssetVOE25(val *CreditFreddieMacVerificationOfAssetVOE25) *NullableCreditFreddieMacVerificationOfAssetVOE25 {
	return &NullableCreditFreddieMacVerificationOfAssetVOE25{value: val, isSet: true}
}

func (v NullableCreditFreddieMacVerificationOfAssetVOE25) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFreddieMacVerificationOfAssetVOE25) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


