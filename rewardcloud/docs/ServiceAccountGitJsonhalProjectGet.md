# ServiceAccountGitJsonhalProjectGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**IsExternal** | Pointer to **NullableBool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**SshPrivateKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServiceAccountGitJsonhalProjectGet

`func NewServiceAccountGitJsonhalProjectGet() *ServiceAccountGitJsonhalProjectGet`

NewServiceAccountGitJsonhalProjectGet instantiates a new ServiceAccountGitJsonhalProjectGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountGitJsonhalProjectGetWithDefaults

`func NewServiceAccountGitJsonhalProjectGetWithDefaults() *ServiceAccountGitJsonhalProjectGet`

NewServiceAccountGitJsonhalProjectGetWithDefaults instantiates a new ServiceAccountGitJsonhalProjectGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ServiceAccountGitJsonhalProjectGet) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceAccountGitJsonhalProjectGet) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceAccountGitJsonhalProjectGet) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceAccountGitJsonhalProjectGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetIsExternal

`func (o *ServiceAccountGitJsonhalProjectGet) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ServiceAccountGitJsonhalProjectGet) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ServiceAccountGitJsonhalProjectGet) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ServiceAccountGitJsonhalProjectGet) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### SetIsExternalNil

`func (o *ServiceAccountGitJsonhalProjectGet) SetIsExternalNil(b bool)`

 SetIsExternalNil sets the value for IsExternal to be an explicit nil

### UnsetIsExternal
`func (o *ServiceAccountGitJsonhalProjectGet) UnsetIsExternal()`

UnsetIsExternal ensures that no value is present for IsExternal, not even an explicit nil
### GetUrl

`func (o *ServiceAccountGitJsonhalProjectGet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceAccountGitJsonhalProjectGet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceAccountGitJsonhalProjectGet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceAccountGitJsonhalProjectGet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServiceAccountGitJsonhalProjectGet) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServiceAccountGitJsonhalProjectGet) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSshPrivateKey

`func (o *ServiceAccountGitJsonhalProjectGet) GetSshPrivateKey() string`

GetSshPrivateKey returns the SshPrivateKey field if non-nil, zero value otherwise.

### GetSshPrivateKeyOk

`func (o *ServiceAccountGitJsonhalProjectGet) GetSshPrivateKeyOk() (*string, bool)`

GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrivateKey

`func (o *ServiceAccountGitJsonhalProjectGet) SetSshPrivateKey(v string)`

SetSshPrivateKey sets SshPrivateKey field to given value.

### HasSshPrivateKey

`func (o *ServiceAccountGitJsonhalProjectGet) HasSshPrivateKey() bool`

HasSshPrivateKey returns a boolean if a field has been set.

### SetSshPrivateKeyNil

`func (o *ServiceAccountGitJsonhalProjectGet) SetSshPrivateKeyNil(b bool)`

 SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil

### UnsetSshPrivateKey
`func (o *ServiceAccountGitJsonhalProjectGet) UnsetSshPrivateKey()`

UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


