# CredentialTypeJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**Git** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCredentialTypeJsonhal

`func NewCredentialTypeJsonhal() *CredentialTypeJsonhal`

NewCredentialTypeJsonhal instantiates a new CredentialTypeJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeJsonhalWithDefaults

`func NewCredentialTypeJsonhalWithDefaults() *CredentialTypeJsonhal`

NewCredentialTypeJsonhalWithDefaults instantiates a new CredentialTypeJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CredentialTypeJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CredentialTypeJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CredentialTypeJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CredentialTypeJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *CredentialTypeJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialTypeJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialTypeJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialTypeJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *CredentialTypeJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CredentialTypeJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CredentialTypeJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CredentialTypeJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *CredentialTypeJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CredentialTypeJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *CredentialTypeJsonhal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialTypeJsonhal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialTypeJsonhal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CredentialTypeJsonhal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CredentialTypeJsonhal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CredentialTypeJsonhal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *CredentialTypeJsonhal) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *CredentialTypeJsonhal) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *CredentialTypeJsonhal) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *CredentialTypeJsonhal) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *CredentialTypeJsonhal) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *CredentialTypeJsonhal) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetGit

`func (o *CredentialTypeJsonhal) GetGit() []string`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *CredentialTypeJsonhal) GetGitOk() (*[]string, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *CredentialTypeJsonhal) SetGit(v []string)`

SetGit sets Git field to given value.

### HasGit

`func (o *CredentialTypeJsonhal) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CredentialTypeJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CredentialTypeJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CredentialTypeJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CredentialTypeJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *CredentialTypeJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *CredentialTypeJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *CredentialTypeJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CredentialTypeJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CredentialTypeJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CredentialTypeJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *CredentialTypeJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *CredentialTypeJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *CredentialTypeJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialTypeJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialTypeJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialTypeJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CredentialTypeJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialTypeJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialTypeJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialTypeJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


