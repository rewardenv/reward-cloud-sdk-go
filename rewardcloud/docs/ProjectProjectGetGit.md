# ProjectProjectGetGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Repo** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKey** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKeyInput** | Pointer to **NullableString** |  | [optional] 
**GitType** | Pointer to [**NullableGitProjectGetGitType**](GitProjectGetGitType.md) |  | [optional] 
**CredentialType** | Pointer to [**NullableGitProjectGetCredentialType**](GitProjectGetCredentialType.md) |  | [optional] 

## Methods

### NewProjectProjectGetGit

`func NewProjectProjectGetGit() *ProjectProjectGetGit`

NewProjectProjectGetGit instantiates a new ProjectProjectGetGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectGetGitWithDefaults

`func NewProjectProjectGetGitWithDefaults() *ProjectProjectGetGit`

NewProjectProjectGetGitWithDefaults instantiates a new ProjectProjectGetGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectProjectGetGit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectProjectGetGit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectProjectGetGit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectProjectGetGit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectProjectGetGit) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectProjectGetGit) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectProjectGetGit) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectProjectGetGit) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectProjectGetGit) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectProjectGetGit) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetRepo

`func (o *ProjectProjectGetGit) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectProjectGetGit) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectProjectGetGit) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectProjectGetGit) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *ProjectProjectGetGit) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *ProjectProjectGetGit) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetUsername

`func (o *ProjectProjectGetGit) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProjectProjectGetGit) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProjectProjectGetGit) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProjectProjectGetGit) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ProjectProjectGetGit) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ProjectProjectGetGit) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ProjectProjectGetGit) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProjectProjectGetGit) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProjectProjectGetGit) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProjectProjectGetGit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ProjectProjectGetGit) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ProjectProjectGetGit) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *ProjectProjectGetGit) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *ProjectProjectGetGit) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *ProjectProjectGetGit) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *ProjectProjectGetGit) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *ProjectProjectGetGit) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *ProjectProjectGetGit) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetSshPrivateKeyInput

`func (o *ProjectProjectGetGit) GetSshPrivateKeyInput() string`

GetSshPrivateKeyInput returns the SshPrivateKeyInput field if non-nil, zero value otherwise.

### GetSshPrivateKeyInputOk

`func (o *ProjectProjectGetGit) GetSshPrivateKeyInputOk() (*string, bool)`

GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyInput

`func (o *ProjectProjectGetGit) SetSshPrivateKeyInput(v string)`

SetSshPrivateKeyInput sets SshPrivateKeyInput field to given value.

### HasSshPrivateKeyInput

`func (o *ProjectProjectGetGit) HasSshPrivateKeyInput() bool`

HasSshPrivateKeyInput returns a boolean if a field has been set.

### SetSshPrivateKeyInputNil

`func (o *ProjectProjectGetGit) SetSshPrivateKeyInputNil(b bool)`

 SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil

### UnsetSshPrivateKeyInput
`func (o *ProjectProjectGetGit) UnsetSshPrivateKeyInput()`

UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
### GetGitType

`func (o *ProjectProjectGetGit) GetGitType() GitProjectGetGitType`

GetGitType returns the GitType field if non-nil, zero value otherwise.

### GetGitTypeOk

`func (o *ProjectProjectGetGit) GetGitTypeOk() (*GitProjectGetGitType, bool)`

GetGitTypeOk returns a tuple with the GitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitType

`func (o *ProjectProjectGetGit) SetGitType(v GitProjectGetGitType)`

SetGitType sets GitType field to given value.

### HasGitType

`func (o *ProjectProjectGetGit) HasGitType() bool`

HasGitType returns a boolean if a field has been set.

### SetGitTypeNil

`func (o *ProjectProjectGetGit) SetGitTypeNil(b bool)`

 SetGitTypeNil sets the value for GitType to be an explicit nil

### UnsetGitType
`func (o *ProjectProjectGetGit) UnsetGitType()`

UnsetGitType ensures that no value is present for GitType, not even an explicit nil
### GetCredentialType

`func (o *ProjectProjectGetGit) GetCredentialType() GitProjectGetCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *ProjectProjectGetGit) GetCredentialTypeOk() (*GitProjectGetCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *ProjectProjectGetGit) SetCredentialType(v GitProjectGetCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *ProjectProjectGetGit) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### SetCredentialTypeNil

`func (o *ProjectProjectGetGit) SetCredentialTypeNil(b bool)`

 SetCredentialTypeNil sets the value for CredentialType to be an explicit nil

### UnsetCredentialType
`func (o *ProjectProjectGetGit) UnsetCredentialType()`

UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


