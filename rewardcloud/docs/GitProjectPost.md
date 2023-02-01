# GitProjectPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repo** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKey** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKeyInput** | Pointer to **NullableString** |  | [optional] 
**GitType** | Pointer to [**NullableGitProjectPostGitType**](GitProjectPostGitType.md) |  | [optional] 
**CredentialType** | Pointer to [**NullableGitProjectPostCredentialType**](GitProjectPostCredentialType.md) |  | [optional] 

## Methods

### NewGitProjectPost

`func NewGitProjectPost() *GitProjectPost`

NewGitProjectPost instantiates a new GitProjectPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitProjectPostWithDefaults

`func NewGitProjectPostWithDefaults() *GitProjectPost`

NewGitProjectPostWithDefaults instantiates a new GitProjectPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepo

`func (o *GitProjectPost) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitProjectPost) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitProjectPost) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitProjectPost) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *GitProjectPost) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *GitProjectPost) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetUsername

`func (o *GitProjectPost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GitProjectPost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GitProjectPost) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GitProjectPost) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GitProjectPost) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GitProjectPost) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *GitProjectPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GitProjectPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GitProjectPost) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GitProjectPost) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GitProjectPost) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GitProjectPost) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *GitProjectPost) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *GitProjectPost) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *GitProjectPost) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *GitProjectPost) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *GitProjectPost) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *GitProjectPost) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetSshPrivateKeyInput

`func (o *GitProjectPost) GetSshPrivateKeyInput() string`

GetSshPrivateKeyInput returns the SshPrivateKeyInput field if non-nil, zero value otherwise.

### GetSshPrivateKeyInputOk

`func (o *GitProjectPost) GetSshPrivateKeyInputOk() (*string, bool)`

GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyInput

`func (o *GitProjectPost) SetSshPrivateKeyInput(v string)`

SetSshPrivateKeyInput sets SshPrivateKeyInput field to given value.

### HasSshPrivateKeyInput

`func (o *GitProjectPost) HasSshPrivateKeyInput() bool`

HasSshPrivateKeyInput returns a boolean if a field has been set.

### SetSshPrivateKeyInputNil

`func (o *GitProjectPost) SetSshPrivateKeyInputNil(b bool)`

 SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil

### UnsetSshPrivateKeyInput
`func (o *GitProjectPost) UnsetSshPrivateKeyInput()`

UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
### GetGitType

`func (o *GitProjectPost) GetGitType() GitProjectPostGitType`

GetGitType returns the GitType field if non-nil, zero value otherwise.

### GetGitTypeOk

`func (o *GitProjectPost) GetGitTypeOk() (*GitProjectPostGitType, bool)`

GetGitTypeOk returns a tuple with the GitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitType

`func (o *GitProjectPost) SetGitType(v GitProjectPostGitType)`

SetGitType sets GitType field to given value.

### HasGitType

`func (o *GitProjectPost) HasGitType() bool`

HasGitType returns a boolean if a field has been set.

### SetGitTypeNil

`func (o *GitProjectPost) SetGitTypeNil(b bool)`

 SetGitTypeNil sets the value for GitType to be an explicit nil

### UnsetGitType
`func (o *GitProjectPost) UnsetGitType()`

UnsetGitType ensures that no value is present for GitType, not even an explicit nil
### GetCredentialType

`func (o *GitProjectPost) GetCredentialType() GitProjectPostCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *GitProjectPost) GetCredentialTypeOk() (*GitProjectPostCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *GitProjectPost) SetCredentialType(v GitProjectPostCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *GitProjectPost) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### SetCredentialTypeNil

`func (o *GitProjectPost) SetCredentialTypeNil(b bool)`

 SetCredentialTypeNil sets the value for CredentialType to be an explicit nil

### UnsetCredentialType
`func (o *GitProjectPost) UnsetCredentialType()`

UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


