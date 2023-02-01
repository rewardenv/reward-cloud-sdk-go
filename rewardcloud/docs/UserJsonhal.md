# UserJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
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

### NewUserJsonhal

`func NewUserJsonhal() *UserJsonhal`

NewUserJsonhal instantiates a new UserJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserJsonhalWithDefaults

`func NewUserJsonhalWithDefaults() *UserJsonhal`

NewUserJsonhalWithDefaults instantiates a new UserJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UserJsonhal) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserJsonhal) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *UserJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UserJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *UserJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *UserJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetPaymentId

`func (o *UserJsonhal) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UserJsonhal) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UserJsonhal) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UserJsonhal) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *UserJsonhal) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *UserJsonhal) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetAuthenticatorId

`func (o *UserJsonhal) GetAuthenticatorId() int32`

GetAuthenticatorId returns the AuthenticatorId field if non-nil, zero value otherwise.

### GetAuthenticatorIdOk

`func (o *UserJsonhal) GetAuthenticatorIdOk() (*int32, bool)`

GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorId

`func (o *UserJsonhal) SetAuthenticatorId(v int32)`

SetAuthenticatorId sets AuthenticatorId field to given value.

### HasAuthenticatorId

`func (o *UserJsonhal) HasAuthenticatorId() bool`

HasAuthenticatorId returns a boolean if a field has been set.

### SetAuthenticatorIdNil

`func (o *UserJsonhal) SetAuthenticatorIdNil(b bool)`

 SetAuthenticatorIdNil sets the value for AuthenticatorId to be an explicit nil

### UnsetAuthenticatorId
`func (o *UserJsonhal) UnsetAuthenticatorId()`

UnsetAuthenticatorId ensures that no value is present for AuthenticatorId, not even an explicit nil
### GetFullname

`func (o *UserJsonhal) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *UserJsonhal) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *UserJsonhal) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *UserJsonhal) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### SetFullnameNil

`func (o *UserJsonhal) SetFullnameNil(b bool)`

 SetFullnameNil sets the value for Fullname to be an explicit nil

### UnsetFullname
`func (o *UserJsonhal) UnsetFullname()`

UnsetFullname ensures that no value is present for Fullname, not even an explicit nil
### GetEmail

`func (o *UserJsonhal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserJsonhal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserJsonhal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserJsonhal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserJsonhal) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserJsonhal) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUsername

`func (o *UserJsonhal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserJsonhal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserJsonhal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserJsonhal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserJsonhal) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserJsonhal) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetRoles

`func (o *UserJsonhal) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserJsonhal) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserJsonhal) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserJsonhal) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UserJsonhal) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UserJsonhal) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetCurrency

`func (o *UserJsonhal) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UserJsonhal) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UserJsonhal) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UserJsonhal) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *UserJsonhal) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *UserJsonhal) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPassword

`func (o *UserJsonhal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserJsonhal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserJsonhal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserJsonhal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetChangePassword

`func (o *UserJsonhal) GetChangePassword() string`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *UserJsonhal) GetChangePasswordOk() (*string, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *UserJsonhal) SetChangePassword(v string)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *UserJsonhal) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### SetChangePasswordNil

`func (o *UserJsonhal) SetChangePasswordNil(b bool)`

 SetChangePasswordNil sets the value for ChangePassword to be an explicit nil

### UnsetChangePassword
`func (o *UserJsonhal) UnsetChangePassword()`

UnsetChangePassword ensures that no value is present for ChangePassword, not even an explicit nil
### GetIsEnabled

`func (o *UserJsonhal) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UserJsonhal) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UserJsonhal) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UserJsonhal) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetRoleGroup

`func (o *UserJsonhal) GetRoleGroup() string`

GetRoleGroup returns the RoleGroup field if non-nil, zero value otherwise.

### GetRoleGroupOk

`func (o *UserJsonhal) GetRoleGroupOk() (*string, bool)`

GetRoleGroupOk returns a tuple with the RoleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleGroup

`func (o *UserJsonhal) SetRoleGroup(v string)`

SetRoleGroup sets RoleGroup field to given value.

### HasRoleGroup

`func (o *UserJsonhal) HasRoleGroup() bool`

HasRoleGroup returns a boolean if a field has been set.

### SetRoleGroupNil

`func (o *UserJsonhal) SetRoleGroupNil(b bool)`

 SetRoleGroupNil sets the value for RoleGroup to be an explicit nil

### UnsetRoleGroup
`func (o *UserJsonhal) UnsetRoleGroup()`

UnsetRoleGroup ensures that no value is present for RoleGroup, not even an explicit nil
### GetTeam

`func (o *UserJsonhal) GetTeam() []string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *UserJsonhal) GetTeamOk() (*[]string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *UserJsonhal) SetTeam(v []string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *UserJsonhal) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UserJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UserJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *UserJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *UserJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *UserJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *UserJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *UserJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *UserJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *UserJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *UserJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *UserJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSalt

`func (o *UserJsonhal) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *UserJsonhal) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *UserJsonhal) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *UserJsonhal) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### SetSaltNil

`func (o *UserJsonhal) SetSaltNil(b bool)`

 SetSaltNil sets the value for Salt to be an explicit nil

### UnsetSalt
`func (o *UserJsonhal) UnsetSalt()`

UnsetSalt ensures that no value is present for Salt, not even an explicit nil
### GetUserIdentifier

`func (o *UserJsonhal) GetUserIdentifier() string`

GetUserIdentifier returns the UserIdentifier field if non-nil, zero value otherwise.

### GetUserIdentifierOk

`func (o *UserJsonhal) GetUserIdentifierOk() (*string, bool)`

GetUserIdentifierOk returns a tuple with the UserIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifier

`func (o *UserJsonhal) SetUserIdentifier(v string)`

SetUserIdentifier sets UserIdentifier field to given value.

### HasUserIdentifier

`func (o *UserJsonhal) HasUserIdentifier() bool`

HasUserIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


