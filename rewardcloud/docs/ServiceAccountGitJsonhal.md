# ServiceAccountGitJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**IsExternal** | Pointer to **NullableBool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKey** | Pointer to **NullableString** |  | [optional] 
**ServiceAccount** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceAccountGitJsonhal

`func NewServiceAccountGitJsonhal() *ServiceAccountGitJsonhal`

NewServiceAccountGitJsonhal instantiates a new ServiceAccountGitJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountGitJsonhalWithDefaults

`func NewServiceAccountGitJsonhalWithDefaults() *ServiceAccountGitJsonhal`

NewServiceAccountGitJsonhalWithDefaults instantiates a new ServiceAccountGitJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ServiceAccountGitJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccountGitJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccountGitJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceAccountGitJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ServiceAccountGitJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountGitJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountGitJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccountGitJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceAccountGitJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceAccountGitJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceAccountGitJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceAccountGitJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetIsExternal

`func (o *ServiceAccountGitJsonhal) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ServiceAccountGitJsonhal) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ServiceAccountGitJsonhal) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ServiceAccountGitJsonhal) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### SetIsExternalNil

`func (o *ServiceAccountGitJsonhal) SetIsExternalNil(b bool)`

 SetIsExternalNil sets the value for IsExternal to be an explicit nil

### UnsetIsExternal
`func (o *ServiceAccountGitJsonhal) UnsetIsExternal()`

UnsetIsExternal ensures that no value is present for IsExternal, not even an explicit nil
### GetUrl

`func (o *ServiceAccountGitJsonhal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceAccountGitJsonhal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceAccountGitJsonhal) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceAccountGitJsonhal) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServiceAccountGitJsonhal) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServiceAccountGitJsonhal) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUsername

`func (o *ServiceAccountGitJsonhal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceAccountGitJsonhal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceAccountGitJsonhal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceAccountGitJsonhal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ServiceAccountGitJsonhal) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ServiceAccountGitJsonhal) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *ServiceAccountGitJsonhal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceAccountGitJsonhal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceAccountGitJsonhal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceAccountGitJsonhal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ServiceAccountGitJsonhal) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ServiceAccountGitJsonhal) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *ServiceAccountGitJsonhal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceAccountGitJsonhal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceAccountGitJsonhal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServiceAccountGitJsonhal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ServiceAccountGitJsonhal) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ServiceAccountGitJsonhal) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *ServiceAccountGitJsonhal) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *ServiceAccountGitJsonhal) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *ServiceAccountGitJsonhal) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *ServiceAccountGitJsonhal) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *ServiceAccountGitJsonhal) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *ServiceAccountGitJsonhal) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetServiceAccount

`func (o *ServiceAccountGitJsonhal) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ServiceAccountGitJsonhal) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ServiceAccountGitJsonhal) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *ServiceAccountGitJsonhal) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### SetServiceAccountNil

`func (o *ServiceAccountGitJsonhal) SetServiceAccountNil(b bool)`

 SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil

### UnsetServiceAccount
`func (o *ServiceAccountGitJsonhal) UnsetServiceAccount()`

UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
### GetCreatedBy

`func (o *ServiceAccountGitJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceAccountGitJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceAccountGitJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServiceAccountGitJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ServiceAccountGitJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ServiceAccountGitJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ServiceAccountGitJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServiceAccountGitJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServiceAccountGitJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServiceAccountGitJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ServiceAccountGitJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ServiceAccountGitJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


