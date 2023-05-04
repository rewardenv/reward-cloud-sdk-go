# ServiceAccountJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
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

### NewServiceAccountJsonhal

`func NewServiceAccountJsonhal() *ServiceAccountJsonhal`

NewServiceAccountJsonhal instantiates a new ServiceAccountJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountJsonhalWithDefaults

`func NewServiceAccountJsonhalWithDefaults() *ServiceAccountJsonhal`

NewServiceAccountJsonhalWithDefaults instantiates a new ServiceAccountJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ServiceAccountJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccountJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccountJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceAccountJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ServiceAccountJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccountJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccountJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccountJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceAccountJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceAccountJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceAccountJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceAccountJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetEmail

`func (o *ServiceAccountJsonhal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ServiceAccountJsonhal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ServiceAccountJsonhal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ServiceAccountJsonhal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ServiceAccountJsonhal) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ServiceAccountJsonhal) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetUsername

`func (o *ServiceAccountJsonhal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ServiceAccountJsonhal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ServiceAccountJsonhal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ServiceAccountJsonhal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ServiceAccountJsonhal) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ServiceAccountJsonhal) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ServiceAccountJsonhal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ServiceAccountJsonhal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ServiceAccountJsonhal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ServiceAccountJsonhal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ServiceAccountJsonhal) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ServiceAccountJsonhal) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProject

`func (o *ServiceAccountJsonhal) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceAccountJsonhal) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceAccountJsonhal) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ServiceAccountJsonhal) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *ServiceAccountJsonhal) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *ServiceAccountJsonhal) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetServiceAccountGit

`func (o *ServiceAccountJsonhal) GetServiceAccountGit() string`

GetServiceAccountGit returns the ServiceAccountGit field if non-nil, zero value otherwise.

### GetServiceAccountGitOk

`func (o *ServiceAccountJsonhal) GetServiceAccountGitOk() (*string, bool)`

GetServiceAccountGitOk returns a tuple with the ServiceAccountGit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountGit

`func (o *ServiceAccountJsonhal) SetServiceAccountGit(v string)`

SetServiceAccountGit sets ServiceAccountGit field to given value.

### HasServiceAccountGit

`func (o *ServiceAccountJsonhal) HasServiceAccountGit() bool`

HasServiceAccountGit returns a boolean if a field has been set.

### SetServiceAccountGitNil

`func (o *ServiceAccountJsonhal) SetServiceAccountGitNil(b bool)`

 SetServiceAccountGitNil sets the value for ServiceAccountGit to be an explicit nil

### UnsetServiceAccountGit
`func (o *ServiceAccountJsonhal) UnsetServiceAccountGit()`

UnsetServiceAccountGit ensures that no value is present for ServiceAccountGit, not even an explicit nil
### GetServiceAccountRegistry

`func (o *ServiceAccountJsonhal) GetServiceAccountRegistry() string`

GetServiceAccountRegistry returns the ServiceAccountRegistry field if non-nil, zero value otherwise.

### GetServiceAccountRegistryOk

`func (o *ServiceAccountJsonhal) GetServiceAccountRegistryOk() (*string, bool)`

GetServiceAccountRegistryOk returns a tuple with the ServiceAccountRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountRegistry

`func (o *ServiceAccountJsonhal) SetServiceAccountRegistry(v string)`

SetServiceAccountRegistry sets ServiceAccountRegistry field to given value.

### HasServiceAccountRegistry

`func (o *ServiceAccountJsonhal) HasServiceAccountRegistry() bool`

HasServiceAccountRegistry returns a boolean if a field has been set.

### SetServiceAccountRegistryNil

`func (o *ServiceAccountJsonhal) SetServiceAccountRegistryNil(b bool)`

 SetServiceAccountRegistryNil sets the value for ServiceAccountRegistry to be an explicit nil

### UnsetServiceAccountRegistry
`func (o *ServiceAccountJsonhal) UnsetServiceAccountRegistry()`

UnsetServiceAccountRegistry ensures that no value is present for ServiceAccountRegistry, not even an explicit nil
### GetCreatedBy

`func (o *ServiceAccountJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceAccountJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceAccountJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ServiceAccountJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ServiceAccountJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ServiceAccountJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ServiceAccountJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServiceAccountJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServiceAccountJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServiceAccountJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ServiceAccountJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ServiceAccountJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


