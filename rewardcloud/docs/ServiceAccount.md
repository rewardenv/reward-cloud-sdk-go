# ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**ServiceAccountGit** | Pointer to **NullableString** |  | [optional] 
**ServiceAccountRegistry** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceAccount

`func NewServiceAccount() *ServiceAccount`

NewServiceAccount instantiates a new ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountWithDefaults

`func NewServiceAccountWithDefaults() *ServiceAccount`

NewServiceAccountWithDefaults instantiates a new ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceAccount) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceAccount) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceAccount) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceAccount) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceAccount) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceAccount) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ServiceAccount) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ServiceAccount) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUsername

`func (o *ServiceAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ServiceAccount) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ServiceAccount) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ServiceAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServiceAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ServiceAccount) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ServiceAccount) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProject

`func (o *ServiceAccount) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceAccount) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceAccount) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceAccount) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *ServiceAccount) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *ServiceAccount) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetServiceAccountGit

`func (o *ServiceAccount) GetServiceAccountGit() string`

GetServiceAccountGit returns the ServiceAccountGit field if non-nil, zero value otherwise.

### GetServiceAccountGitOk

`func (o *ServiceAccount) GetServiceAccountGitOk() (*string, bool)`

GetServiceAccountGitOk returns a tuple with the ServiceAccountGit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountGit

`func (o *ServiceAccount) SetServiceAccountGit(v string)`

SetServiceAccountGit sets ServiceAccountGit field to given value.

### HasServiceAccountGit

`func (o *ServiceAccount) HasServiceAccountGit() bool`

HasServiceAccountGit returns a boolean if a field has been set.

### SetServiceAccountGitNil

`func (o *ServiceAccount) SetServiceAccountGitNil(b bool)`

 SetServiceAccountGitNil sets the value for ServiceAccountGit to be an explicit nil

### UnsetServiceAccountGit
`func (o *ServiceAccount) UnsetServiceAccountGit()`

UnsetServiceAccountGit ensures that no value is present for ServiceAccountGit, not even an explicit nil
### GetServiceAccountRegistry

`func (o *ServiceAccount) GetServiceAccountRegistry() string`

GetServiceAccountRegistry returns the ServiceAccountRegistry field if non-nil, zero value otherwise.

### GetServiceAccountRegistryOk

`func (o *ServiceAccount) GetServiceAccountRegistryOk() (*string, bool)`

GetServiceAccountRegistryOk returns a tuple with the ServiceAccountRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountRegistry

`func (o *ServiceAccount) SetServiceAccountRegistry(v string)`

SetServiceAccountRegistry sets ServiceAccountRegistry field to given value.

### HasServiceAccountRegistry

`func (o *ServiceAccount) HasServiceAccountRegistry() bool`

HasServiceAccountRegistry returns a boolean if a field has been set.

### SetServiceAccountRegistryNil

`func (o *ServiceAccount) SetServiceAccountRegistryNil(b bool)`

 SetServiceAccountRegistryNil sets the value for ServiceAccountRegistry to be an explicit nil

### UnsetServiceAccountRegistry
`func (o *ServiceAccount) UnsetServiceAccountRegistry()`

UnsetServiceAccountRegistry ensures that no value is present for ServiceAccountRegistry, not even an explicit nil
### GetCreatedBy

`func (o *ServiceAccount) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceAccount) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceAccount) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServiceAccount) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ServiceAccount) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ServiceAccount) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ServiceAccount) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServiceAccount) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServiceAccount) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServiceAccount) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ServiceAccount) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ServiceAccount) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


