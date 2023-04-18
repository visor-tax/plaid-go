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

// Taxform Data about an official document used to report the user's income to the IRS.
type Taxform struct {
	// An identifier of the document referenced by the document metadata.
	DocId *string `json:"doc_id,omitempty"`
	// The type of tax document. Currently, the only supported value is `w2`.
	DocumentType string `json:"document_type"`
	W2 *W2 `json:"w2,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Taxform Taxform

// NewTaxform instantiates a new Taxform object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxform(documentType string) *Taxform {
	this := Taxform{}
	this.DocumentType = documentType
	return &this
}

// NewTaxformWithDefaults instantiates a new Taxform object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxformWithDefaults() *Taxform {
	this := Taxform{}
	return &this
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *Taxform) GetDocId() string {
	if o == nil || o.DocId == nil {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxform) GetDocIdOk() (*string, bool) {
	if o == nil || o.DocId == nil {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *Taxform) HasDocId() bool {
	if o != nil && o.DocId != nil {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *Taxform) SetDocId(v string) {
	o.DocId = &v
}

// GetDocumentType returns the DocumentType field value
func (o *Taxform) GetDocumentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
func (o *Taxform) GetDocumentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentType, true
}

// SetDocumentType sets field value
func (o *Taxform) SetDocumentType(v string) {
	o.DocumentType = v
}

// GetW2 returns the W2 field value if set, zero value otherwise.
func (o *Taxform) GetW2() W2 {
	if o == nil || o.W2 == nil {
		var ret W2
		return ret
	}
	return *o.W2
}

// GetW2Ok returns a tuple with the W2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taxform) GetW2Ok() (*W2, bool) {
	if o == nil || o.W2 == nil {
		return nil, false
	}
	return o.W2, true
}

// HasW2 returns a boolean if a field has been set.
func (o *Taxform) HasW2() bool {
	if o != nil && o.W2 != nil {
		return true
	}

	return false
}

// SetW2 gets a reference to the given W2 and assigns it to the W2 field.
func (o *Taxform) SetW2(v W2) {
	o.W2 = &v
}

func (o Taxform) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocId != nil {
		toSerialize["doc_id"] = o.DocId
	}
	if true {
		toSerialize["document_type"] = o.DocumentType
	}
	if o.W2 != nil {
		toSerialize["w2"] = o.W2
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Taxform) UnmarshalJSON(bytes []byte) (err error) {
	varTaxform := _Taxform{}

	if err = json.Unmarshal(bytes, &varTaxform); err == nil {
		*o = Taxform(varTaxform)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "doc_id")
		delete(additionalProperties, "document_type")
		delete(additionalProperties, "w2")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaxform struct {
	value *Taxform
	isSet bool
}

func (v NullableTaxform) Get() *Taxform {
	return v.value
}

func (v *NullableTaxform) Set(val *Taxform) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxform) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxform) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxform(val *Taxform) *NullableTaxform {
	return &NullableTaxform{value: val, isSet: true}
}

func (v NullableTaxform) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxform) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


