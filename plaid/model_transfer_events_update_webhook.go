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

// TransferEventsUpdateWebhook Fired when new transfer events are available. Receiving this webhook indicates you should fetch the new events from `/transfer/event/sync`.
type TransferEventsUpdateWebhook struct {
	// `TRANSFER`
	WebhookType string `json:"webhook_type"`
	// `TRANSFER_EVENTS_UPDATE`
	WebhookCode string `json:"webhook_code"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _TransferEventsUpdateWebhook TransferEventsUpdateWebhook

// NewTransferEventsUpdateWebhook instantiates a new TransferEventsUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEventsUpdateWebhook(webhookType string, webhookCode string, environment WebhookEnvironmentValues) *TransferEventsUpdateWebhook {
	this := TransferEventsUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.Environment = environment
	return &this
}

// NewTransferEventsUpdateWebhookWithDefaults instantiates a new TransferEventsUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventsUpdateWebhookWithDefaults() *TransferEventsUpdateWebhook {
	this := TransferEventsUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *TransferEventsUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *TransferEventsUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *TransferEventsUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *TransferEventsUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *TransferEventsUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *TransferEventsUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetEnvironment returns the Environment field value
func (o *TransferEventsUpdateWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *TransferEventsUpdateWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *TransferEventsUpdateWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o TransferEventsUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferEventsUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varTransferEventsUpdateWebhook := _TransferEventsUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varTransferEventsUpdateWebhook); err == nil {
		*o = TransferEventsUpdateWebhook(varTransferEventsUpdateWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferEventsUpdateWebhook struct {
	value *TransferEventsUpdateWebhook
	isSet bool
}

func (v NullableTransferEventsUpdateWebhook) Get() *TransferEventsUpdateWebhook {
	return v.value
}

func (v *NullableTransferEventsUpdateWebhook) Set(val *TransferEventsUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventsUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventsUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventsUpdateWebhook(val *TransferEventsUpdateWebhook) *NullableTransferEventsUpdateWebhook {
	return &NullableTransferEventsUpdateWebhook{value: val, isSet: true}
}

func (v NullableTransferEventsUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventsUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


