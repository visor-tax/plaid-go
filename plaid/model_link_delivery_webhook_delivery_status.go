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
	"fmt"
)

// LinkDeliveryWebhookDeliveryStatus The status of the delivery of the hosted link to the user
type LinkDeliveryWebhookDeliveryStatus string

var _ = fmt.Printf

// List of LinkDeliveryWebhookDeliveryStatus
const (
	LINKDELIVERYWEBHOOKDELIVERYSTATUS_SUCCEEDED LinkDeliveryWebhookDeliveryStatus = "succeeded"
	LINKDELIVERYWEBHOOKDELIVERYSTATUS_FAILED LinkDeliveryWebhookDeliveryStatus = "failed"
)

var allowedLinkDeliveryWebhookDeliveryStatusEnumValues = []LinkDeliveryWebhookDeliveryStatus{
	"succeeded",
	"failed",
}

func (v *LinkDeliveryWebhookDeliveryStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := LinkDeliveryWebhookDeliveryStatus(value)


	*v = enumTypeValue
	return nil
}

// NewLinkDeliveryWebhookDeliveryStatusFromValue returns a pointer to a valid LinkDeliveryWebhookDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLinkDeliveryWebhookDeliveryStatusFromValue(v string) (*LinkDeliveryWebhookDeliveryStatus, error) {
	ev := LinkDeliveryWebhookDeliveryStatus(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LinkDeliveryWebhookDeliveryStatus) IsValid() bool {
	for _, existing := range allowedLinkDeliveryWebhookDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LinkDeliveryWebhookDeliveryStatus value
func (v LinkDeliveryWebhookDeliveryStatus) Ptr() *LinkDeliveryWebhookDeliveryStatus {
	return &v
}

type NullableLinkDeliveryWebhookDeliveryStatus struct {
	value *LinkDeliveryWebhookDeliveryStatus
	isSet bool
}

func (v NullableLinkDeliveryWebhookDeliveryStatus) Get() *LinkDeliveryWebhookDeliveryStatus {
	return v.value
}

func (v *NullableLinkDeliveryWebhookDeliveryStatus) Set(val *LinkDeliveryWebhookDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkDeliveryWebhookDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkDeliveryWebhookDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkDeliveryWebhookDeliveryStatus(val *LinkDeliveryWebhookDeliveryStatus) *NullableLinkDeliveryWebhookDeliveryStatus {
	return &NullableLinkDeliveryWebhookDeliveryStatus{value: val, isSet: true}
}

func (v NullableLinkDeliveryWebhookDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkDeliveryWebhookDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

