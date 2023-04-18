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
	"time"
)

// WalletTransactionGetResponse WalletTransactionGetResponse defines the response schema for `/wallet/transaction/get`
type WalletTransactionGetResponse struct {
	// A unique ID identifying the transaction
	TransactionId string `json:"transaction_id"`
	// The EMI (E-Money Institution) wallet that this payment is associated with, if any. This wallet is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.
	WalletId string `json:"wallet_id"`
	// A reference for the transaction
	Reference string `json:"reference"`
	// The type of the transaction. The supported transaction types that are returned are: `BANK_TRANSFER:` a transaction which credits an e-wallet through an external bank transfer.  `PAYOUT:` a transaction which debits an e-wallet by disbursing funds to a counterparty.  `PIS_PAY_IN:` a payment which credits an e-wallet through Plaid's Payment Initiation Services (PIS) APIs. For more information see the [Payment Initiation endpoints](https://plaid.com/docs/api/products/payment-initiation/).  `REFUND:` a transaction which debits an e-wallet by refunding a previously initiated payment made through Plaid's [PIS APIs](https://plaid.com/docs/api/products/payment-initiation/).  `FUNDS_SWEEP`: an automated transaction which debits funds from an e-wallet to a designated client-owned account.
	Type string `json:"type"`
	Amount WalletTransactionAmount `json:"amount"`
	Counterparty WalletTransactionCounterparty `json:"counterparty"`
	Status WalletTransactionStatus `json:"status"`
	// Timestamp when the transaction was created, in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time of the last time the `status` was updated, in IS0 8601 format
	LastStatusUpdate time.Time `json:"last_status_update"`
	// The payment id that this transaction is associated with, if any. This is present only for transaction types `PIS_PAY_IN` and `REFUND`.
	PaymentId NullableString `json:"payment_id,omitempty"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _WalletTransactionGetResponse WalletTransactionGetResponse

// NewWalletTransactionGetResponse instantiates a new WalletTransactionGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionGetResponse(transactionId string, walletId string, reference string, type_ string, amount WalletTransactionAmount, counterparty WalletTransactionCounterparty, status WalletTransactionStatus, createdAt time.Time, lastStatusUpdate time.Time, requestId string) *WalletTransactionGetResponse {
	this := WalletTransactionGetResponse{}
	this.TransactionId = transactionId
	this.WalletId = walletId
	this.Reference = reference
	this.Type = type_
	this.Amount = amount
	this.Counterparty = counterparty
	this.Status = status
	this.CreatedAt = createdAt
	this.LastStatusUpdate = lastStatusUpdate
	this.RequestId = requestId
	return &this
}

// NewWalletTransactionGetResponseWithDefaults instantiates a new WalletTransactionGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionGetResponseWithDefaults() *WalletTransactionGetResponse {
	this := WalletTransactionGetResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *WalletTransactionGetResponse) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *WalletTransactionGetResponse) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetWalletId returns the WalletId field value
func (o *WalletTransactionGetResponse) GetWalletId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetWalletIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WalletId, true
}

// SetWalletId sets field value
func (o *WalletTransactionGetResponse) SetWalletId(v string) {
	o.WalletId = v
}

// GetReference returns the Reference field value
func (o *WalletTransactionGetResponse) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *WalletTransactionGetResponse) SetReference(v string) {
	o.Reference = v
}

// GetType returns the Type field value
func (o *WalletTransactionGetResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WalletTransactionGetResponse) SetType(v string) {
	o.Type = v
}

// GetAmount returns the Amount field value
func (o *WalletTransactionGetResponse) GetAmount() WalletTransactionAmount {
	if o == nil {
		var ret WalletTransactionAmount
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetAmountOk() (*WalletTransactionAmount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WalletTransactionGetResponse) SetAmount(v WalletTransactionAmount) {
	o.Amount = v
}

// GetCounterparty returns the Counterparty field value
func (o *WalletTransactionGetResponse) GetCounterparty() WalletTransactionCounterparty {
	if o == nil {
		var ret WalletTransactionCounterparty
		return ret
	}

	return o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetCounterpartyOk() (*WalletTransactionCounterparty, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Counterparty, true
}

// SetCounterparty sets field value
func (o *WalletTransactionGetResponse) SetCounterparty(v WalletTransactionCounterparty) {
	o.Counterparty = v
}

// GetStatus returns the Status field value
func (o *WalletTransactionGetResponse) GetStatus() WalletTransactionStatus {
	if o == nil {
		var ret WalletTransactionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetStatusOk() (*WalletTransactionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WalletTransactionGetResponse) SetStatus(v WalletTransactionStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WalletTransactionGetResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WalletTransactionGetResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetLastStatusUpdate returns the LastStatusUpdate field value
func (o *WalletTransactionGetResponse) GetLastStatusUpdate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastStatusUpdate
}

// GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetLastStatusUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastStatusUpdate, true
}

// SetLastStatusUpdate sets field value
func (o *WalletTransactionGetResponse) SetLastStatusUpdate(v time.Time) {
	o.LastStatusUpdate = v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletTransactionGetResponse) GetPaymentId() string {
	if o == nil || o.PaymentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletTransactionGetResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *WalletTransactionGetResponse) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *WalletTransactionGetResponse) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}
// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *WalletTransactionGetResponse) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *WalletTransactionGetResponse) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetRequestId returns the RequestId field value
func (o *WalletTransactionGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *WalletTransactionGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o WalletTransactionGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["wallet_id"] = o.WalletId
	}
	if true {
		toSerialize["reference"] = o.Reference
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["counterparty"] = o.Counterparty
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["last_status_update"] = o.LastStatusUpdate
	}
	if o.PaymentId.IsSet() {
		toSerialize["payment_id"] = o.PaymentId.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WalletTransactionGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWalletTransactionGetResponse := _WalletTransactionGetResponse{}

	if err = json.Unmarshal(bytes, &varWalletTransactionGetResponse); err == nil {
		*o = WalletTransactionGetResponse(varWalletTransactionGetResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "wallet_id")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "type")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "counterparty")
		delete(additionalProperties, "status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "last_status_update")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletTransactionGetResponse struct {
	value *WalletTransactionGetResponse
	isSet bool
}

func (v NullableWalletTransactionGetResponse) Get() *WalletTransactionGetResponse {
	return v.value
}

func (v *NullableWalletTransactionGetResponse) Set(val *WalletTransactionGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionGetResponse(val *WalletTransactionGetResponse) *NullableWalletTransactionGetResponse {
	return &NullableWalletTransactionGetResponse{value: val, isSet: true}
}

func (v NullableWalletTransactionGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


