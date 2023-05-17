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

// RecurringCancelledWebhook Fired when a recurring transfer is cancelled by Plaid.
type RecurringCancelledWebhook struct {
	// `TRANSFER`
	WebhookType string `json:"webhook_type"`
	// `RECURRING_CANCELLED`
	WebhookCode string `json:"webhook_code"`
	// Plaid’s unique identifier for a recurring transfer.
	RecurringTransferId string `json:"recurring_transfer_id"`
	Environment WebhookEnvironmentValues `json:"environment"`
	AdditionalProperties map[string]interface{}
}

type _RecurringCancelledWebhook RecurringCancelledWebhook

// NewRecurringCancelledWebhook instantiates a new RecurringCancelledWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringCancelledWebhook(webhookType string, webhookCode string, recurringTransferId string, environment WebhookEnvironmentValues) *RecurringCancelledWebhook {
	this := RecurringCancelledWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.RecurringTransferId = recurringTransferId
	this.Environment = environment
	return &this
}

// NewRecurringCancelledWebhookWithDefaults instantiates a new RecurringCancelledWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringCancelledWebhookWithDefaults() *RecurringCancelledWebhook {
	this := RecurringCancelledWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *RecurringCancelledWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *RecurringCancelledWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *RecurringCancelledWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *RecurringCancelledWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *RecurringCancelledWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *RecurringCancelledWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetRecurringTransferId returns the RecurringTransferId field value
func (o *RecurringCancelledWebhook) GetRecurringTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurringTransferId
}

// GetRecurringTransferIdOk returns a tuple with the RecurringTransferId field value
// and a boolean to check if the value has been set.
func (o *RecurringCancelledWebhook) GetRecurringTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecurringTransferId, true
}

// SetRecurringTransferId sets field value
func (o *RecurringCancelledWebhook) SetRecurringTransferId(v string) {
	o.RecurringTransferId = v
}

// GetEnvironment returns the Environment field value
func (o *RecurringCancelledWebhook) GetEnvironment() WebhookEnvironmentValues {
	if o == nil {
		var ret WebhookEnvironmentValues
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *RecurringCancelledWebhook) GetEnvironmentOk() (*WebhookEnvironmentValues, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *RecurringCancelledWebhook) SetEnvironment(v WebhookEnvironmentValues) {
	o.Environment = v
}

func (o RecurringCancelledWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["recurring_transfer_id"] = o.RecurringTransferId
	}
	if true {
		toSerialize["environment"] = o.Environment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecurringCancelledWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varRecurringCancelledWebhook := _RecurringCancelledWebhook{}

	if err = json.Unmarshal(bytes, &varRecurringCancelledWebhook); err == nil {
		*o = RecurringCancelledWebhook(varRecurringCancelledWebhook)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "recurring_transfer_id")
		delete(additionalProperties, "environment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecurringCancelledWebhook struct {
	value *RecurringCancelledWebhook
	isSet bool
}

func (v NullableRecurringCancelledWebhook) Get() *RecurringCancelledWebhook {
	return v.value
}

func (v *NullableRecurringCancelledWebhook) Set(val *RecurringCancelledWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringCancelledWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringCancelledWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringCancelledWebhook(val *RecurringCancelledWebhook) *NullableRecurringCancelledWebhook {
	return &NullableRecurringCancelledWebhook{value: val, isSet: true}
}

func (v NullableRecurringCancelledWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringCancelledWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


