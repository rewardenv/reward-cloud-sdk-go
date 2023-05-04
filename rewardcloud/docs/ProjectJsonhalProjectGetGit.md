# ProjectJsonhalProjectGetGit

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
**GitType** | Pointer to [**NullableGitJsonhalProjectGetGitType**](GitJsonhalProjectGetGitType.md) |  | [optional] 
**CredentialType** | Pointer to [**NullableGitJsonhalProjectGetCredentialType**](GitJsonhalProjectGetCredentialType.md) |  | [optional] 

## Methods

### NewProjectJsonhalProjectGetGit

`func NewProjectJsonhalProjectGetGit() *ProjectJsonhalProjectGetGit`

NewProjectJsonhalProjectGetGit instantiates a new ProjectJsonhalProjectGetGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectJsonhalProjectGetGitWithDefaults

`func NewProjectJsonhalProjectGetGitWithDefaults() *ProjectJsonhalProjectGetGit`

NewProjectJsonhalProjectGetGitWithDefaults instantiates a new ProjectJsonhalProjectGetGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectJsonhalProjectGetGit) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectJsonhalProjectGetGit) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectJsonhalProjectGetGit) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectJsonhalProjectGetGit) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProjectJsonhalProjectGetGit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectJsonhalProjectGetGit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectJsonhalProjectGetGit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectJsonhalProjectGetGit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectJsonhalProjectGetGit) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectJsonhalProjectGetGit) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectJsonhalProjectGetGit) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectJsonhalProjectGetGit) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectJsonhalProjectGetGit) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectJsonhalProjectGetGit) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetRepo

`func (o *ProjectJsonhalProjectGetGit) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectJsonhalProjectGetGit) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectJsonhalProjectGetGit) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectJsonhalProjectGetGit) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *ProjectJsonhalProjectGetGit) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *ProjectJsonhalProjectGetGit) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetUsername

`func (o *ProjectJsonhalProjectGetGit) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProjectJsonhalProjectGetGit) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProjectJsonhalProjectGetGit) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProjectJsonhalProjectGetGit) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ProjectJsonhalProjectGetGit) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ProjectJsonhalProjectGetGit) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ProjectJsonhalProjectGetGit) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProjectJsonhalProjectGetGit) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProjectJsonhalProjectGetGit) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProjectJsonhalProjectGetGit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ProjectJsonhalProjectGetGit) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ProjectJsonhalProjectGetGit) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *ProjectJsonhalProjectGetGit) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *ProjectJsonhalProjectGetGit) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *ProjectJsonhalProjectGetGit) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *ProjectJsonhalProjectGetGit) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *ProjectJsonhalProjectGetGit) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *ProjectJsonhalProjectGetGit) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetSshPrivateKeyInput

`func (o *ProjectJsonhalProjectGetGit) GetSshPrivateKeyInput() string`

GetSshPrivateKeyInput returns the SshPrivateKeyInput field if non-nil, zero value otherwise.

### GetSshPrivateKeyInputOk

`func (o *ProjectJsonhalProjectGetGit) GetSshPrivateKeyInputOk() (*string, bool)`

GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyInput

`func (o *ProjectJsonhalProjectGetGit) SetSshPrivateKeyInput(v string)`

SetSshPrivateKeyInput sets SshPrivateKeyInput field to given value.

### HasSshPrivateKeyInput

`func (o *ProjectJsonhalProjectGetGit) HasSshPrivateKeyInput() bool`

HasSshPrivateKeyInput returns a boolean if a field has been set.

### SetSshPrivateKeyInputNil

`func (o *ProjectJsonhalProjectGetGit) SetSshPrivateKeyInputNil(b bool)`

 SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil

### UnsetSshPrivateKeyInput
`func (o *ProjectJsonhalProjectGetGit) UnsetSshPrivateKeyInput()`

UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
### GetGitType

`func (o *ProjectJsonhalProjectGetGit) GetGitType() GitJsonhalProjectGetGitType`

GetGitType returns the GitType field if non-nil, zero value otherwise.

### GetGitTypeOk

`func (o *ProjectJsonhalProjectGetGit) GetGitTypeOk() (*GitJsonhalProjectGetGitType, bool)`

GetGitTypeOk returns a tuple with the GitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitType

`func (o *ProjectJsonhalProjectGetGit) SetGitType(v GitJsonhalProjectGetGitType)`

SetGitType sets GitType field to given value.

### HasGitType

`func (o *ProjectJsonhalProjectGetGit) HasGitType() bool`

HasGitType returns a boolean if a field has been set.

### SetGitTypeNil

`func (o *ProjectJsonhalProjectGetGit) SetGitTypeNil(b bool)`

 SetGitTypeNil sets the value for GitType to be an explicit nil

### UnsetGitType
`func (o *ProjectJsonhalProjectGetGit) UnsetGitType()`

UnsetGitType ensures that no value is present for GitType, not even an explicit nil
### GetCredentialType

`func (o *ProjectJsonhalProjectGetGit) GetCredentialType() GitJsonhalProjectGetCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *ProjectJsonhalProjectGetGit) GetCredentialTypeOk() (*GitJsonhalProjectGetCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *ProjectJsonhalProjectGetGit) SetCredentialType(v GitJsonhalProjectGetCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *ProjectJsonhalProjectGetGit) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### SetCredentialTypeNil

`func (o *ProjectJsonhalProjectGetGit) SetCredentialTypeNil(b bool)`

 SetCredentialTypeNil sets the value for CredentialType to be an explicit nil

### UnsetCredentialType
`func (o *ProjectJsonhalProjectGetGit) UnsetCredentialType()`

UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


