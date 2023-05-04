# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**AuthenticatorId** | Pointer to **NullableInt32** |  | [optional] 
**Fullname** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ChangePassword** | Pointer to **NullableString** |  | [optional] 
**CurrentPassword** | Pointer to **NullableString** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**RoleGroup** | Pointer to **NullableString** |  | [optional] 
**Team** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Salt** | Pointer to **NullableString** | not needed when using the \&quot;bcrypt\&quot; algorithm in security.yaml | [optional] [readonly] 
**UserIdentifier** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *User) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *User) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *User) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *User) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *User) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *User) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetPaymentId

`func (o *User) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *User) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *User) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *User) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *User) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *User) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetAuthenticatorId

`func (o *User) GetAuthenticatorId() int32`

GetAuthenticatorId returns the AuthenticatorId field if non-nil, zero value otherwise.

### GetAuthenticatorIdOk

`func (o *User) GetAuthenticatorIdOk() (*int32, bool)`

GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorId

`func (o *User) SetAuthenticatorId(v int32)`

SetAuthenticatorId sets AuthenticatorId field to given value.

### HasAuthenticatorId

`func (o *User) HasAuthenticatorId() bool`

HasAuthenticatorId returns a boolean if a field has been set.

### SetAuthenticatorIdNil

`func (o *User) SetAuthenticatorIdNil(b bool)`

 SetAuthenticatorIdNil sets the value for AuthenticatorId to be an explicit nil

### UnsetAuthenticatorId
`func (o *User) UnsetAuthenticatorId()`

UnsetAuthenticatorId ensures that no value is present for AuthenticatorId, not even an explicit nil
### GetFullname

`func (o *User) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *User) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *User) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *User) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### SetFullnameNil

`func (o *User) SetFullnameNil(b bool)`

 SetFullnameNil sets the value for Fullname to be an explicit nil

### UnsetFullname
`func (o *User) UnsetFullname()`

UnsetFullname ensures that no value is present for Fullname, not even an explicit nil
### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *User) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *User) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *User) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *User) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetCurrency

`func (o *User) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *User) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *User) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *User) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *User) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *User) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetChangePassword

`func (o *User) GetChangePassword() string`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *User) GetChangePasswordOk() (*string, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *User) SetChangePassword(v string)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *User) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### SetChangePasswordNil

`func (o *User) SetChangePasswordNil(b bool)`

 SetChangePasswordNil sets the value for ChangePassword to be an explicit nil

### UnsetChangePassword
`func (o *User) UnsetChangePassword()`

UnsetChangePassword ensures that no value is present for ChangePassword, not even an explicit nil
### GetCurrentPassword

`func (o *User) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *User) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *User) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *User) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### SetCurrentPasswordNil

`func (o *User) SetCurrentPasswordNil(b bool)`

 SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil

### UnsetCurrentPassword
`func (o *User) UnsetCurrentPassword()`

UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
### GetIsEnabled

`func (o *User) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *User) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *User) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *User) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetRoleGroup

`func (o *User) GetRoleGroup() string`

GetRoleGroup returns the RoleGroup field if non-nil, zero value otherwise.

### GetRoleGroupOk

`func (o *User) GetRoleGroupOk() (*string, bool)`

GetRoleGroupOk returns a tuple with the RoleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleGroup

`func (o *User) SetRoleGroup(v string)`

SetRoleGroup sets RoleGroup field to given value.

### HasRoleGroup

`func (o *User) HasRoleGroup() bool`

HasRoleGroup returns a boolean if a field has been set.

### SetRoleGroupNil

`func (o *User) SetRoleGroupNil(b bool)`

 SetRoleGroupNil sets the value for RoleGroup to be an explicit nil

### UnsetRoleGroup
`func (o *User) UnsetRoleGroup()`

UnsetRoleGroup ensures that no value is present for RoleGroup, not even an explicit nil
### GetTeam

`func (o *User) GetTeam() []string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *User) GetTeamOk() (*[]string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *User) SetTeam(v []string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *User) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetCreatedBy

`func (o *User) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *User) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *User) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *User) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *User) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *User) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *User) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *User) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *User) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *User) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *User) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *User) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSalt

`func (o *User) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *User) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *User) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *User) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### SetSaltNil

`func (o *User) SetSaltNil(b bool)`

 SetSaltNil sets the value for Salt to be an explicit nil

### UnsetSalt
`func (o *User) UnsetSalt()`

UnsetSalt ensures that no value is present for Salt, not even an explicit nil
### GetUserIdentifier

`func (o *User) GetUserIdentifier() string`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *User) GetUserIdentifierOk() (*string, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *User) SetUserIdentifier(v string)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *User) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


