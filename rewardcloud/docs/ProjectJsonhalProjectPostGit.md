# ProjectJsonhalProjectPostGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Repo** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKey** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKeyInput** | Pointer to **NullableString** |  | [optional] 
**GitType** | Pointer to [**NullableGitJsonhalProjectPostGitType**](GitJsonhalProjectPostGitType.md) |  | [optional] 
**CredentialType** | Pointer to [**NullableGitJsonhalProjectPostCredentialType**](GitJsonhalProjectPostCredentialType.md) |  | [optional] 

## Methods

### NewProjectJsonhalProjectPostGit

`func NewProjectJsonhalProjectPostGit() *ProjectJsonhalProjectPostGit`

NewProjectJsonhalProjectPostGit instantiates a new ProjectJsonhalProjectPostGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectJsonhalProjectPostGitWithDefaults

`func NewProjectJsonhalProjectPostGitWithDefaults() *ProjectJsonhalProjectPostGit`

NewProjectJsonhalProjectPostGitWithDefaults instantiates a new ProjectJsonhalProjectPostGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectJsonhalProjectPostGit) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectJsonhalProjectPostGit) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectJsonhalProjectPostGit) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectJsonhalProjectPostGit) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRepo

`func (o *ProjectJsonhalProjectPostGit) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ProjectJsonhalProjectPostGit) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ProjectJsonhalProjectPostGit) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ProjectJsonhalProjectPostGit) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *ProjectJsonhalProjectPostGit) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *ProjectJsonhalProjectPostGit) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetUsername

`func (o *ProjectJsonhalProjectPostGit) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProjectJsonhalProjectPostGit) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProjectJsonhalProjectPostGit) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProjectJsonhalProjectPostGit) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ProjectJsonhalProjectPostGit) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ProjectJsonhalProjectPostGit) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ProjectJsonhalProjectPostGit) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProjectJsonhalProjectPostGit) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProjectJsonhalProjectPostGit) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProjectJsonhalProjectPostGit) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ProjectJsonhalProjectPostGit) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ProjectJsonhalProjectPostGit) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *ProjectJsonhalProjectPostGit) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *ProjectJsonhalProjectPostGit) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *ProjectJsonhalProjectPostGit) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *ProjectJsonhalProjectPostGit) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *ProjectJsonhalProjectPostGit) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *ProjectJsonhalProjectPostGit) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetSshPrivateKeyInput

`func (o *ProjectJsonhalProjectPostGit) GetSshPrivateKeyInput() string`

GetSshPrivateKeyInput returns the SshPrivateKeyInput field if non-nil, zero value otherwise.

### GetSshPrivateKeyInputOk

`func (o *ProjectJsonhalProjectPostGit) GetSshPrivateKeyInputOk() (*string, bool)`

GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyInput

`func (o *ProjectJsonhalProjectPostGit) SetSshPrivateKeyInput(v string)`

SetSshPrivateKeyInput sets SshPrivateKeyInput field to given value.

### HasSshPrivateKeyInput

`func (o *ProjectJsonhalProjectPostGit) HasSshPrivateKeyInput() bool`

HasSshPrivateKeyInput returns a boolean if a field has been set.

### SetSshPrivateKeyInputNil

`func (o *ProjectJsonhalProjectPostGit) SetSshPrivateKeyInputNil(b bool)`

 SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil

### UnsetSshPrivateKeyInput
`func (o *ProjectJsonhalProjectPostGit) UnsetSshPrivateKeyInput()`

UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
### GetGitType

`func (o *ProjectJsonhalProjectPostGit) GetGitType() GitJsonhalProjectPostGitType`

GetGitType returns the GitType field if non-nil, zero value otherwise.

### GetGitTypeOk

`func (o *ProjectJsonhalProjectPostGit) GetGitTypeOk() (*GitJsonhalProjectPostGitType, bool)`

GetGitTypeOk returns a tuple with the GitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitType

`func (o *ProjectJsonhalProjectPostGit) SetGitType(v GitJsonhalProjectPostGitType)`

SetGitType sets GitType field to given value.

### HasGitType

`func (o *ProjectJsonhalProjectPostGit) HasGitType() bool`

HasGitType returns a boolean if a field has been set.

### SetGitTypeNil

`func (o *ProjectJsonhalProjectPostGit) SetGitTypeNil(b bool)`

 SetGitTypeNil sets the value for GitType to be an explicit nil

### UnsetGitType
`func (o *ProjectJsonhalProjectPostGit) UnsetGitType()`

UnsetGitType ensures that no value is present for GitType, not even an explicit nil
### GetCredentialType

`func (o *ProjectJsonhalProjectPostGit) GetCredentialType() GitJsonhalProjectPostCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *ProjectJsonhalProjectPostGit) GetCredentialTypeOk() (*GitJsonhalProjectPostCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *ProjectJsonhalProjectPostGit) SetCredentialType(v GitJsonhalProjectPostCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *ProjectJsonhalProjectPostGit) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### SetCredentialTypeNil

`func (o *ProjectJsonhalProjectPostGit) SetCredentialTypeNil(b bool)`

 SetCredentialTypeNil sets the value for CredentialType to be an explicit nil

### UnsetCredentialType
`func (o *ProjectJsonhalProjectPostGit) UnsetCredentialType()`

UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


