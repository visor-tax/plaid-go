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

// IdentityVerificationCreateResponse A identity verification attempt represents a customer's attempt to verify their identity, reflecting the required steps for completing the session, the results for each step, and information collected in the process.
type IdentityVerificationCreateResponse struct {
	// ID of the associated Identity Verification attempt.
	Id string `json:"id"`
	// An identifier to help you connect this object to your internal systems. For example, your database ID corresponding to this object.
	ClientUserId string `json:"client_user_id"`
	// An ISO8601 formatted timestamp.
	CreatedAt time.Time `json:"created_at"`
	// An ISO8601 formatted timestamp.
	CompletedAt NullableTime `json:"completed_at"`
	// The ID for the Identity Verification preceding this session. This field will only be filled if the current Identity Verification is a retry of a previous attempt.
	PreviousAttemptId NullableString `json:"previous_attempt_id"`
	// A shareable URL that can be sent directly to the user to complete verification
	ShareableUrl NullableString `json:"shareable_url"`
	Template IdentityVerificationTemplateReference `json:"template"`
	User IdentityVerificationUserData `json:"user"`
	Status IdentityVerificationStatus `json:"status"`
	Steps IdentityVerificationStepSummary `json:"steps"`
	DocumentaryVerification NullableDocumentaryVerification `json:"documentary_verification"`
	KycCheck NullableKYCCheckDetails `json:"kyc_check"`
	RiskCheck NullableRiskCheckDetails `json:"risk_check"`
	// ID of the associated screening.
	WatchlistScreeningId NullableString `json:"watchlist_screening_id"`
	// An ISO8601 formatted timestamp.
	RedactedAt NullableTime `json:"redacted_at"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _IdentityVerificationCreateResponse IdentityVerificationCreateResponse

// NewIdentityVerificationCreateResponse instantiates a new IdentityVerificationCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityVerificationCreateResponse(id string, clientUserId string, createdAt time.Time, completedAt NullableTime, previousAttemptId NullableString, shareableUrl NullableString, template IdentityVerificationTemplateReference, user IdentityVerificationUserData, status IdentityVerificationStatus, steps IdentityVerificationStepSummary, documentaryVerification NullableDocumentaryVerification, kycCheck NullableKYCCheckDetails, riskCheck NullableRiskCheckDetails, watchlistScreeningId NullableString, redactedAt NullableTime, requestId string) *IdentityVerificationCreateResponse {
	this := IdentityVerificationCreateResponse{}
	this.Id = id
	this.ClientUserId = clientUserId
	this.CreatedAt = createdAt
	this.CompletedAt = completedAt
	this.PreviousAttemptId = previousAttemptId
	this.ShareableUrl = shareableUrl
	this.Template = template
	this.User = user
	this.Status = status
	this.Steps = steps
	this.DocumentaryVerification = documentaryVerification
	this.KycCheck = kycCheck
	this.RiskCheck = riskCheck
	this.WatchlistScreeningId = watchlistScreeningId
	this.RedactedAt = redactedAt
	this.RequestId = requestId
	return &this
}

// NewIdentityVerificationCreateResponseWithDefaults instantiates a new IdentityVerificationCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityVerificationCreateResponseWithDefaults() *IdentityVerificationCreateResponse {
	this := IdentityVerificationCreateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *IdentityVerificationCreateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IdentityVerificationCreateResponse) SetId(v string) {
	o.Id = v
}

// GetClientUserId returns the ClientUserId field value
func (o *IdentityVerificationCreateResponse) GetClientUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *IdentityVerificationCreateResponse) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IdentityVerificationCreateResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IdentityVerificationCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCompletedAt returns the CompletedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IdentityVerificationCreateResponse) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// SetCompletedAt sets field value
func (o *IdentityVerificationCreateResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}

// GetPreviousAttemptId returns the PreviousAttemptId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationCreateResponse) GetPreviousAttemptId() string {
	if o == nil || o.PreviousAttemptId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PreviousAttemptId.Get()
}

// GetPreviousAttemptIdOk returns a tuple with the PreviousAttemptId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetPreviousAttemptIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreviousAttemptId.Get(), o.PreviousAttemptId.IsSet()
}

// SetPreviousAttemptId sets field value
func (o *IdentityVerificationCreateResponse) SetPreviousAttemptId(v string) {
	o.PreviousAttemptId.Set(&v)
}

// GetShareableUrl returns the ShareableUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationCreateResponse) GetShareableUrl() string {
	if o == nil || o.ShareableUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareableUrl.Get()
}

// GetShareableUrlOk returns a tuple with the ShareableUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetShareableUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShareableUrl.Get(), o.ShareableUrl.IsSet()
}

// SetShareableUrl sets field value
func (o *IdentityVerificationCreateResponse) SetShareableUrl(v string) {
	o.ShareableUrl.Set(&v)
}

// GetTemplate returns the Template field value
func (o *IdentityVerificationCreateResponse) GetTemplate() IdentityVerificationTemplateReference {
	if o == nil {
		var ret IdentityVerificationTemplateReference
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetTemplateOk() (*IdentityVerificationTemplateReference, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *IdentityVerificationCreateResponse) SetTemplate(v IdentityVerificationTemplateReference) {
	o.Template = v
}

// GetUser returns the User field value
func (o *IdentityVerificationCreateResponse) GetUser() IdentityVerificationUserData {
	if o == nil {
		var ret IdentityVerificationUserData
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetUserOk() (*IdentityVerificationUserData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *IdentityVerificationCreateResponse) SetUser(v IdentityVerificationUserData) {
	o.User = v
}

// GetStatus returns the Status field value
func (o *IdentityVerificationCreateResponse) GetStatus() IdentityVerificationStatus {
	if o == nil {
		var ret IdentityVerificationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetStatusOk() (*IdentityVerificationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IdentityVerificationCreateResponse) SetStatus(v IdentityVerificationStatus) {
	o.Status = v
}

// GetSteps returns the Steps field value
func (o *IdentityVerificationCreateResponse) GetSteps() IdentityVerificationStepSummary {
	if o == nil {
		var ret IdentityVerificationStepSummary
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetStepsOk() (*IdentityVerificationStepSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value
func (o *IdentityVerificationCreateResponse) SetSteps(v IdentityVerificationStepSummary) {
	o.Steps = v
}

// GetDocumentaryVerification returns the DocumentaryVerification field value
// If the value is explicit nil, the zero value for DocumentaryVerification will be returned
func (o *IdentityVerificationCreateResponse) GetDocumentaryVerification() DocumentaryVerification {
	if o == nil || o.DocumentaryVerification.Get() == nil {
		var ret DocumentaryVerification
		return ret
	}

	return *o.DocumentaryVerification.Get()
}

// GetDocumentaryVerificationOk returns a tuple with the DocumentaryVerification field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetDocumentaryVerificationOk() (*DocumentaryVerification, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DocumentaryVerification.Get(), o.DocumentaryVerification.IsSet()
}

// SetDocumentaryVerification sets field value
func (o *IdentityVerificationCreateResponse) SetDocumentaryVerification(v DocumentaryVerification) {
	o.DocumentaryVerification.Set(&v)
}

// GetKycCheck returns the KycCheck field value
// If the value is explicit nil, the zero value for KYCCheckDetails will be returned
func (o *IdentityVerificationCreateResponse) GetKycCheck() KYCCheckDetails {
	if o == nil || o.KycCheck.Get() == nil {
		var ret KYCCheckDetails
		return ret
	}

	return *o.KycCheck.Get()
}

// GetKycCheckOk returns a tuple with the KycCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetKycCheckOk() (*KYCCheckDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.KycCheck.Get(), o.KycCheck.IsSet()
}

// SetKycCheck sets field value
func (o *IdentityVerificationCreateResponse) SetKycCheck(v KYCCheckDetails) {
	o.KycCheck.Set(&v)
}

// GetRiskCheck returns the RiskCheck field value
// If the value is explicit nil, the zero value for RiskCheckDetails will be returned
func (o *IdentityVerificationCreateResponse) GetRiskCheck() RiskCheckDetails {
	if o == nil || o.RiskCheck.Get() == nil {
		var ret RiskCheckDetails
		return ret
	}

	return *o.RiskCheck.Get()
}

// GetRiskCheckOk returns a tuple with the RiskCheck field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetRiskCheckOk() (*RiskCheckDetails, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RiskCheck.Get(), o.RiskCheck.IsSet()
}

// SetRiskCheck sets field value
func (o *IdentityVerificationCreateResponse) SetRiskCheck(v RiskCheckDetails) {
	o.RiskCheck.Set(&v)
}

// GetWatchlistScreeningId returns the WatchlistScreeningId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityVerificationCreateResponse) GetWatchlistScreeningId() string {
	if o == nil || o.WatchlistScreeningId.Get() == nil {
		var ret string
		return ret
	}

	return *o.WatchlistScreeningId.Get()
}

// GetWatchlistScreeningIdOk returns a tuple with the WatchlistScreeningId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetWatchlistScreeningIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WatchlistScreeningId.Get(), o.WatchlistScreeningId.IsSet()
}

// SetWatchlistScreeningId sets field value
func (o *IdentityVerificationCreateResponse) SetWatchlistScreeningId(v string) {
	o.WatchlistScreeningId.Set(&v)
}

// GetRedactedAt returns the RedactedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IdentityVerificationCreateResponse) GetRedactedAt() time.Time {
	if o == nil || o.RedactedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.RedactedAt.Get()
}

// GetRedactedAtOk returns a tuple with the RedactedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityVerificationCreateResponse) GetRedactedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RedactedAt.Get(), o.RedactedAt.IsSet()
}

// SetRedactedAt sets field value
func (o *IdentityVerificationCreateResponse) SetRedactedAt(v time.Time) {
	o.RedactedAt.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *IdentityVerificationCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *IdentityVerificationCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *IdentityVerificationCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o IdentityVerificationCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if true {
		toSerialize["previous_attempt_id"] = o.PreviousAttemptId.Get()
	}
	if true {
		toSerialize["shareable_url"] = o.ShareableUrl.Get()
	}
	if true {
		toSerialize["template"] = o.Template
	}
	if true {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["steps"] = o.Steps
	}
	if true {
		toSerialize["documentary_verification"] = o.DocumentaryVerification.Get()
	}
	if true {
		toSerialize["kyc_check"] = o.KycCheck.Get()
	}
	if true {
		toSerialize["risk_check"] = o.RiskCheck.Get()
	}
	if true {
		toSerialize["watchlist_screening_id"] = o.WatchlistScreeningId.Get()
	}
	if true {
		toSerialize["redacted_at"] = o.RedactedAt.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityVerificationCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityVerificationCreateResponse := _IdentityVerificationCreateResponse{}

	if err = json.Unmarshal(bytes, &varIdentityVerificationCreateResponse); err == nil {
		*o = IdentityVerificationCreateResponse(varIdentityVerificationCreateResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "previous_attempt_id")
		delete(additionalProperties, "shareable_url")
		delete(additionalProperties, "template")
		delete(additionalProperties, "user")
		delete(additionalProperties, "status")
		delete(additionalProperties, "steps")
		delete(additionalProperties, "documentary_verification")
		delete(additionalProperties, "kyc_check")
		delete(additionalProperties, "risk_check")
		delete(additionalProperties, "watchlist_screening_id")
		delete(additionalProperties, "redacted_at")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityVerificationCreateResponse struct {
	value *IdentityVerificationCreateResponse
	isSet bool
}

func (v NullableIdentityVerificationCreateResponse) Get() *IdentityVerificationCreateResponse {
	return v.value
}

func (v *NullableIdentityVerificationCreateResponse) Set(val *IdentityVerificationCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityVerificationCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityVerificationCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityVerificationCreateResponse(val *IdentityVerificationCreateResponse) *NullableIdentityVerificationCreateResponse {
	return &NullableIdentityVerificationCreateResponse{value: val, isSet: true}
}

func (v NullableIdentityVerificationCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityVerificationCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


