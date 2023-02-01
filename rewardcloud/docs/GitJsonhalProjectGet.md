# GitJsonhalProjectGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
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

### NewGitJsonhalProjectGet

`func NewGitJsonhalProjectGet() *GitJsonhalProjectGet`

NewGitJsonhalProjectGet instantiates a new GitJsonhalProjectGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitJsonhalProjectGetWithDefaults

`func NewGitJsonhalProjectGetWithDefaults() *GitJsonhalProjectGet`

NewGitJsonhalProjectGetWithDefaults instantiates a new GitJsonhalProjectGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GitJsonhalProjectGet) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GitJsonhalProjectGet) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GitJsonhalProjectGet) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GitJsonhalProjectGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *GitJsonhalProjectGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitJsonhalProjectGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitJsonhalProjectGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GitJsonhalProjectGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GitJsonhalProjectGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GitJsonhalProjectGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GitJsonhalProjectGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GitJsonhalProjectGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *GitJsonhalProjectGet) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *GitJsonhalProjectGet) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetRepo

`func (o *GitJsonhalProjectGet) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *GitJsonhalProjectGet) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *GitJsonhalProjectGet) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *GitJsonhalProjectGet) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### SetRepoNil

`func (o *GitJsonhalProjectGet) SetRepoNil(b bool)`

 SetRepoNil sets the value for Repo to be an explicit nil

### UnsetRepo
`func (o *GitJsonhalProjectGet) UnsetRepo()`

UnsetRepo ensures that no value is present for Repo, not even an explicit nil
### GetUsername

`func (o *GitJsonhalProjectGet) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GitJsonhalProjectGet) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GitJsonhalProjectGet) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GitJsonhalProjectGet) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GitJsonhalProjectGet) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GitJsonhalProjectGet) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *GitJsonhalProjectGet) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GitJsonhalProjectGet) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GitJsonhalProjectGet) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GitJsonhalProjectGet) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *GitJsonhalProjectGet) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *GitJsonhalProjectGet) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetSshPrivateKey

`func (o *GitJsonhalProjectGet) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *GitJsonhalProjectGet) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *GitJsonhalProjectGet) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *GitJsonhalProjectGet) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *GitJsonhalProjectGet) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *GitJsonhalProjectGet) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
### GetSshPrivateKeyInput

`func (o *GitJsonhalProjectGet) GetSshPrivateKeyInput() string`

GetSshPrivateKeyInput returns the SshPrivateKeyInput field if non-nil, zero value otherwise.

### GetSshPrivateKeyInputOk

`func (o *GitJsonhalProjectGet) GetSshPrivateKeyInputOk() (*string, bool)`

GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKeyInput

`func (o *GitJsonhalProjectGet) SetSshPrivateKeyInput(v string)`

SetSshPrivateKeyInput sets SshPrivateKeyInput field to given value.

### HasSshPrivateKeyInput

`func (o *GitJsonhalProjectGet) HasSshPrivateKeyInput() bool`

HasSshPrivateKeyInput returns a boolean if a field has been set.

### SetSshPrivateKeyInputNil

`func (o *GitJsonhalProjectGet) SetSshPrivateKeyInputNil(b bool)`

 SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil

### UnsetSshPrivateKeyInput
`func (o *GitJsonhalProjectGet) UnsetSshPrivateKeyInput()`

UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
### GetGitType

`func (o *GitJsonhalProjectGet) GetGitType() GitJsonhalProjectGetGitType`

GetGitType returns the GitType field if non-nil, zero value otherwise.

### GetGitTypeOk

`func (o *GitJsonhalProjectGet) GetGitTypeOk() (*GitJsonhalProjectGetGitType, bool)`

GetGitTypeOk returns a tuple with the GitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitType

`func (o *GitJsonhalProjectGet) SetGitType(v GitJsonhalProjectGetGitType)`

SetGitType sets GitType field to given value.

### HasGitType

`func (o *GitJsonhalProjectGet) HasGitType() bool`

HasGitType returns a boolean if a field has been set.

### SetGitTypeNil

`func (o *GitJsonhalProjectGet) SetGitTypeNil(b bool)`

 SetGitTypeNil sets the value for GitType to be an explicit nil

### UnsetGitType
`func (o *GitJsonhalProjectGet) UnsetGitType()`

UnsetGitType ensures that no value is present for GitType, not even an explicit nil
### GetCredentialType

`func (o *GitJsonhalProjectGet) GetCredentialType() GitJsonhalProjectGetCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *GitJsonhalProjectGet) GetCredentialTypeOk() (*GitJsonhalProjectGetCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *GitJsonhalProjectGet) SetCredentialType(v GitJsonhalProjectGetCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *GitJsonhalProjectGet) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### SetCredentialTypeNil

`func (o *GitJsonhalProjectGet) SetCredentialTypeNil(b bool)`

 SetCredentialTypeNil sets the value for CredentialType to be an explicit nil

### UnsetCredentialType
`func (o *GitJsonhalProjectGet) UnsetCredentialType()`

UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


