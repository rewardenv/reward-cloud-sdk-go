# GitJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Repo** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKey** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKeyInput** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**GitType** | Pointer to **NullableString** |  | [optional] 
**CredentialType** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGitJsonhal

`func NewGitJsonhal() *GitJsonhal`

NewGitJsonhal instantiates a new GitJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitJsonhalWithDefaults

`func NewGitJsonhalWithDefaults() *GitJsonhal`

NewGitJsonhalWithDefaults instantiates a new GitJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GitJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GitJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GitJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GitJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *GitJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GitJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GitJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GitJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GitJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GitJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *GitJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *GitJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetRepo

`func (o *GitJsonhal) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitJsonhal) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitJsonhal) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitJsonhal) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *GitJsonhal) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *GitJsonhal) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetUsername

`func (o *GitJsonhal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GitJsonhal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GitJsonhal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GitJsonhal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GitJsonhal) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GitJsonhal) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *GitJsonhal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GitJsonhal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GitJsonhal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GitJsonhal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GitJsonhal) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GitJsonhal) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *GitJsonhal) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *GitJsonhal) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *GitJsonhal) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *GitJsonhal) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *GitJsonhal) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *GitJsonhal) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetSshPrivateKeyInput

`func (o *GitJsonhal) GetSshPrivateKeyInput() string`

GetSshPrivateKeyInput returns the SshPrivateKeyInput field if non-nil, zero value otherwise.

### GetSshPrivateKeyInputOk

`func (o *GitJsonhal) GetSshPrivateKeyInputOk() (*string, bool)`

GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyInput

`func (o *GitJsonhal) SetSshPrivateKeyInput(v string)`

SetSshPrivateKeyInput sets SshPrivateKeyInput field to given value.

### HasSshPrivateKeyInput

`func (o *GitJsonhal) HasSshPrivateKeyInput() bool`

HasSshPrivateKeyInput returns a boolean if a field has been set.

### SetSshPrivateKeyInputNil

`func (o *GitJsonhal) SetSshPrivateKeyInputNil(b bool)`

 SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil

### UnsetSshPrivateKeyInput
`func (o *GitJsonhal) UnsetSshPrivateKeyInput()`

UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
### GetProject

`func (o *GitJsonhal) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GitJsonhal) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GitJsonhal) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GitJsonhal) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *GitJsonhal) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *GitJsonhal) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetGitType

`func (o *GitJsonhal) GetGitType() string`

GetGitType returns the GitType field if non-nil, zero value otherwise.

### GetGitTypeOk

`func (o *GitJsonhal) GetGitTypeOk() (*string, bool)`

GetGitTypeOk returns a tuple with the GitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitType

`func (o *GitJsonhal) SetGitType(v string)`

SetGitType sets GitType field to given value.

### HasGitType

`func (o *GitJsonhal) HasGitType() bool`

HasGitType returns a boolean if a field has been set.

### SetGitTypeNil

`func (o *GitJsonhal) SetGitTypeNil(b bool)`

 SetGitTypeNil sets the value for GitType to be an explicit nil

### UnsetGitType
`func (o *GitJsonhal) UnsetGitType()`

UnsetGitType ensures that no value is present for GitType, not even an explicit nil
### GetCredentialType

`func (o *GitJsonhal) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *GitJsonhal) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *GitJsonhal) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *GitJsonhal) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### SetCredentialTypeNil

`func (o *GitJsonhal) SetCredentialTypeNil(b bool)`

 SetCredentialTypeNil sets the value for CredentialType to be an explicit nil

### UnsetCredentialType
`func (o *GitJsonhal) UnsetCredentialType()`

UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil
### GetCreatedBy

`func (o *GitJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GitJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GitJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GitJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GitJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GitJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *GitJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GitJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GitJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GitJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *GitJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *GitJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *GitJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GitJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GitJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GitJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GitJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


