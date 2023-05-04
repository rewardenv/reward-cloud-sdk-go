/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
	"time"
)

// User TODO User can only delete himself!!
type User struct {
	Id              *int32         `json:"id,omitempty"`
	Uuid            NullableString `json:"uuid,omitempty"`
	PaymentId       NullableString `json:"paymentId,omitempty"`
	AuthenticatorId NullableInt32  `json:"authenticatorId,omitempty"`
	Fullname        NullableString `json:"fullname,omitempty"`
	Email           NullableString `json:"email,omitempty"`
	Username        NullableString `json:"username,omitempty"`
	Roles           []string       `json:"roles,omitempty"`
	Currency        NullableString `json:"currency,omitempty"`
	Password        *string        `json:"password,omitempty"`
	ChangePassword  NullableString `json:"changePassword,omitempty"`
	CurrentPassword NullableString `json:"currentPassword,omitempty"`
	IsEnabled       *bool          `json:"isEnabled,omitempty"`
	RoleGroup       NullableString `json:"roleGroup,omitempty"`
	Team            []string       `json:"team,omitempty"`
	CreatedBy       NullableString `json:"createdBy,omitempty"`
	UpdatedBy       NullableString `json:"updatedBy,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	// not needed when using the \"bcrypt\" algorithm in security.yaml
	Salt           NullableString `json:"salt,omitempty"`
	UserIdentifier *string        `json:"userIdentifier,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *User) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *User) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *User) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *User) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *User) UnsetUuid() {
	o.Uuid.Unset()
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPaymentId() string {
	if o == nil || isNil(o.PaymentId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// HasPaymentId returns a boolean if a field has been set.
func (o *User) HasPaymentId() bool {
	if o != nil && o.PaymentId.IsSet() {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given NullableString and assigns it to the PaymentId field.
func (o *User) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}

// SetPaymentIdNil sets the value for PaymentId to be an explicit nil
func (o *User) SetPaymentIdNil() {
	o.PaymentId.Set(nil)
}

// UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
func (o *User) UnsetPaymentId() {
	o.PaymentId.Unset()
}

// GetAuthenticatorId returns the AuthenticatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetAuthenticatorId() int32 {
	if o == nil || isNil(o.AuthenticatorId.Get()) {
		var ret int32
		return ret
	}
	return *o.AuthenticatorId.Get()
}

// GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetAuthenticatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticatorId.Get(), o.AuthenticatorId.IsSet()
}

// HasAuthenticatorId returns a boolean if a field has been set.
func (o *User) HasAuthenticatorId() bool {
	if o != nil && o.AuthenticatorId.IsSet() {
		return true
	}

	return false
}

// SetAuthenticatorId gets a reference to the given NullableInt32 and assigns it to the AuthenticatorId field.
func (o *User) SetAuthenticatorId(v int32) {
	o.AuthenticatorId.Set(&v)
}

// SetAuthenticatorIdNil sets the value for AuthenticatorId to be an explicit nil
func (o *User) SetAuthenticatorIdNil() {
	o.AuthenticatorId.Set(nil)
}

// UnsetAuthenticatorId ensures that no value is present for AuthenticatorId, not even an explicit nil
func (o *User) UnsetAuthenticatorId() {
	o.AuthenticatorId.Unset()
}

// GetFullname returns the Fullname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetFullname() string {
	if o == nil || isNil(o.Fullname.Get()) {
		var ret string
		return ret
	}
	return *o.Fullname.Get()
}

// GetFullnameOk returns a tuple with the Fullname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetFullnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fullname.Get(), o.Fullname.IsSet()
}

// HasFullname returns a boolean if a field has been set.
func (o *User) HasFullname() bool {
	if o != nil && o.Fullname.IsSet() {
		return true
	}

	return false
}

// SetFullname gets a reference to the given NullableString and assigns it to the Fullname field.
func (o *User) SetFullname(v string) {
	o.Fullname.Set(&v)
}

// SetFullnameNil sets the value for Fullname to be an explicit nil
func (o *User) SetFullnameNil() {
	o.Fullname.Set(nil)
}

// UnsetFullname ensures that no value is present for Fullname, not even an explicit nil
func (o *User) UnsetFullname() {
	o.Fullname.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetEmail() string {
	if o == nil || isNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *User) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *User) UnsetEmail() {
	o.Email.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUsername() string {
	if o == nil || isNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *User) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *User) UnsetUsername() {
	o.Username.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetRolesOk() ([]string, bool) {
	if o == nil || isNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *User) HasRoles() bool {
	if o != nil && isNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *User) SetRoles(v []string) {
	o.Roles = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCurrency() string {
	if o == nil || isNil(o.Currency.Get()) {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *User) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *User) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *User) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *User) UnsetCurrency() {
	o.Currency.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *User) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *User) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *User) SetPassword(v string) {
	o.Password = &v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetChangePassword() string {
	if o == nil || isNil(o.ChangePassword.Get()) {
		var ret string
		return ret
	}
	return *o.ChangePassword.Get()
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetChangePasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangePassword.Get(), o.ChangePassword.IsSet()
}

// HasChangePassword returns a boolean if a field has been set.
func (o *User) HasChangePassword() bool {
	if o != nil && o.ChangePassword.IsSet() {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given NullableString and assigns it to the ChangePassword field.
func (o *User) SetChangePassword(v string) {
	o.ChangePassword.Set(&v)
}

// SetChangePasswordNil sets the value for ChangePassword to be an explicit nil
func (o *User) SetChangePasswordNil() {
	o.ChangePassword.Set(nil)
}

// UnsetChangePassword ensures that no value is present for ChangePassword, not even an explicit nil
func (o *User) UnsetChangePassword() {
	o.ChangePassword.Unset()
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCurrentPassword() string {
	if o == nil || isNil(o.CurrentPassword.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentPassword.Get()
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCurrentPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPassword.Get(), o.CurrentPassword.IsSet()
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *User) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword.IsSet() {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given NullableString and assigns it to the CurrentPassword field.
func (o *User) SetCurrentPassword(v string) {
	o.CurrentPassword.Set(&v)
}

// SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil
func (o *User) SetCurrentPasswordNil() {
	o.CurrentPassword.Set(nil)
}

// UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
func (o *User) UnsetCurrentPassword() {
	o.CurrentPassword.Unset()
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *User) GetIsEnabled() bool {
	if o == nil || isNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *User) HasIsEnabled() bool {
	if o != nil && !isNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *User) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetRoleGroup returns the RoleGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetRoleGroup() string {
	if o == nil || isNil(o.RoleGroup.Get()) {
		var ret string
		return ret
	}
	return *o.RoleGroup.Get()
}

// GetRoleGroupOk returns a tuple with the RoleGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetRoleGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleGroup.Get(), o.RoleGroup.IsSet()
}

// HasRoleGroup returns a boolean if a field has been set.
func (o *User) HasRoleGroup() bool {
	if o != nil && o.RoleGroup.IsSet() {
		return true
	}

	return false
}

// SetRoleGroup gets a reference to the given NullableString and assigns it to the RoleGroup field.
func (o *User) SetRoleGroup(v string) {
	o.RoleGroup.Set(&v)
}

// SetRoleGroupNil sets the value for RoleGroup to be an explicit nil
func (o *User) SetRoleGroupNil() {
	o.RoleGroup.Set(nil)
}

// UnsetRoleGroup ensures that no value is present for RoleGroup, not even an explicit nil
func (o *User) UnsetRoleGroup() {
	o.RoleGroup.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *User) GetTeam() []string {
	if o == nil || isNil(o.Team) {
		var ret []string
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTeamOk() ([]string, bool) {
	if o == nil || isNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *User) HasTeam() bool {
	if o != nil && !isNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given []string and assigns it to the Team field.
func (o *User) SetTeam(v []string) {
	o.Team = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *User) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *User) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *User) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *User) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *User) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *User) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *User) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *User) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *User) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *User) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetSalt() string {
	if o == nil || isNil(o.Salt.Get()) {
		var ret string
		return ret
	}
	return *o.Salt.Get()
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetSaltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Salt.Get(), o.Salt.IsSet()
}

// HasSalt returns a boolean if a field has been set.
func (o *User) HasSalt() bool {
	if o != nil && o.Salt.IsSet() {
		return true
	}

	return false
}

// SetSalt gets a reference to the given NullableString and assigns it to the Salt field.
func (o *User) SetSalt(v string) {
	o.Salt.Set(&v)
}

// SetSaltNil sets the value for Salt to be an explicit nil
func (o *User) SetSaltNil() {
	o.Salt.Set(nil)
}

// UnsetSalt ensures that no value is present for Salt, not even an explicit nil
func (o *User) UnsetSalt() {
	o.Salt.Unset()
}

// GetUserIdentifier returns the UserIdentifier field value if set, zero value otherwise.
func (o *User) GetUserIdentifier() string {
	if o == nil || isNil(o.UserIdentifier) {
		var ret string
		return ret
	}
	return *o.UserIdentifier
}

// GetUserIdentifierOk returns a tuple with the UserIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUserIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.UserIdentifier) {
		return nil, false
	}
	return o.UserIdentifier, true
}

// HasUserIdentifier returns a boolean if a field has been set.
func (o *User) HasUserIdentifier() bool {
	if o != nil && !isNil(o.UserIdentifier) {
		return true
	}

	return false
}

// SetUserIdentifier gets a reference to the given string and assigns it to the UserIdentifier field.
func (o *User) SetUserIdentifier(v string) {
	o.UserIdentifier = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.PaymentId.IsSet() {
		toSerialize["paymentId"] = o.PaymentId.Get()
	}
	if o.AuthenticatorId.IsSet() {
		toSerialize["authenticatorId"] = o.AuthenticatorId.Get()
	}
	if o.Fullname.IsSet() {
		toSerialize["fullname"] = o.Fullname.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if o.ChangePassword.IsSet() {
		toSerialize["changePassword"] = o.ChangePassword.Get()
	}
	if o.CurrentPassword.IsSet() {
		toSerialize["currentPassword"] = o.CurrentPassword.Get()
	}
	if !isNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.RoleGroup.IsSet() {
		toSerialize["roleGroup"] = o.RoleGroup.Get()
	}
	if !isNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updatedBy"] = o.UpdatedBy.Get()
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Salt.IsSet() {
		toSerialize["salt"] = o.Salt.Get()
	}
	if !isNil(o.UserIdentifier) {
		toSerialize["userIdentifier"] = o.UserIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
