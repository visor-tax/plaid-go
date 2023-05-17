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

// SelfieCheckSelfie Catpures and analysis from a user's selfie.
type SelfieCheckSelfie struct {
	Status SelfieStatus `json:"status"`
	// The `attempt` field begins with 1 and increments with each subsequent selfie upload.
	Attempt int32 `json:"attempt"`
	Capture SelfieCapture `json:"capture"`
	Analysis SelfieAnalysis `json:"analysis"`
	AdditionalProperties map[string]interface{}
}

type _SelfieCheckSelfie SelfieCheckSelfie

// NewSelfieCheckSelfie instantiates a new SelfieCheckSelfie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfieCheckSelfie(status SelfieStatus, attempt int32, capture SelfieCapture, analysis SelfieAnalysis) *SelfieCheckSelfie {
	this := SelfieCheckSelfie{}
	this.Status = status
	this.Attempt = attempt
	this.Capture = capture
	this.Analysis = analysis
	return &this
}

// NewSelfieCheckSelfieWithDefaults instantiates a new SelfieCheckSelfie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfieCheckSelfieWithDefaults() *SelfieCheckSelfie {
	this := SelfieCheckSelfie{}
	return &this
}

// GetStatus returns the Status field value
func (o *SelfieCheckSelfie) GetStatus() SelfieStatus {
	if o == nil {
		var ret SelfieStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SelfieCheckSelfie) GetStatusOk() (*SelfieStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SelfieCheckSelfie) SetStatus(v SelfieStatus) {
	o.Status = v
}

// GetAttempt returns the Attempt field value
func (o *SelfieCheckSelfie) GetAttempt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attempt
}

// GetAttemptOk returns a tuple with the Attempt field value
// and a boolean to check if the value has been set.
func (o *SelfieCheckSelfie) GetAttemptOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attempt, true
}

// SetAttempt sets field value
func (o *SelfieCheckSelfie) SetAttempt(v int32) {
	o.Attempt = v
}

// GetCapture returns the Capture field value
func (o *SelfieCheckSelfie) GetCapture() SelfieCapture {
	if o == nil {
		var ret SelfieCapture
		return ret
	}

	return o.Capture
}

// GetCaptureOk returns a tuple with the Capture field value
// and a boolean to check if the value has been set.
func (o *SelfieCheckSelfie) GetCaptureOk() (*SelfieCapture, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Capture, true
}

// SetCapture sets field value
func (o *SelfieCheckSelfie) SetCapture(v SelfieCapture) {
	o.Capture = v
}

// GetAnalysis returns the Analysis field value
func (o *SelfieCheckSelfie) GetAnalysis() SelfieAnalysis {
	if o == nil {
		var ret SelfieAnalysis
		return ret
	}

	return o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value
// and a boolean to check if the value has been set.
func (o *SelfieCheckSelfie) GetAnalysisOk() (*SelfieAnalysis, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Analysis, true
}

// SetAnalysis sets field value
func (o *SelfieCheckSelfie) SetAnalysis(v SelfieAnalysis) {
	o.Analysis = v
}

func (o SelfieCheckSelfie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["attempt"] = o.Attempt
	}
	if true {
		toSerialize["capture"] = o.Capture
	}
	if true {
		toSerialize["analysis"] = o.Analysis
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfieCheckSelfie) UnmarshalJSON(bytes []byte) (err error) {
	varSelfieCheckSelfie := _SelfieCheckSelfie{}

	if err = json.Unmarshal(bytes, &varSelfieCheckSelfie); err == nil {
		*o = SelfieCheckSelfie(varSelfieCheckSelfie)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "attempt")
		delete(additionalProperties, "capture")
		delete(additionalProperties, "analysis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfieCheckSelfie struct {
	value *SelfieCheckSelfie
	isSet bool
}

func (v NullableSelfieCheckSelfie) Get() *SelfieCheckSelfie {
	return v.value
}

func (v *NullableSelfieCheckSelfie) Set(val *SelfieCheckSelfie) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfieCheckSelfie) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfieCheckSelfie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfieCheckSelfie(val *SelfieCheckSelfie) *NullableSelfieCheckSelfie {
	return &NullableSelfieCheckSelfie{value: val, isSet: true}
}

func (v NullableSelfieCheckSelfie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfieCheckSelfie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


