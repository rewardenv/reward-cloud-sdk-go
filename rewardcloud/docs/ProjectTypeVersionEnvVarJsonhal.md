# ProjectTypeVersionEnvVarJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Note** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **NullableBool** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**ProjectTypeVersion** | Pointer to **NullableString** |  | [optional] 
**ProjectTypeVersionEnvVarExample** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectTypeVersionEnvVarJsonhal

`func NewProjectTypeVersionEnvVarJsonhal() *ProjectTypeVersionEnvVarJsonhal`

NewProjectTypeVersionEnvVarJsonhal instantiates a new ProjectTypeVersionEnvVarJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTypeVersionEnvVarJsonhalWithDefaults

`func NewProjectTypeVersionEnvVarJsonhalWithDefaults() *ProjectTypeVersionEnvVarJsonhal`

NewProjectTypeVersionEnvVarJsonhalWithDefaults instantiates a new ProjectTypeVersionEnvVarJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectTypeVersionEnvVarJsonhal) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectTypeVersionEnvVarJsonhal) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectTypeVersionEnvVarJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProjectTypeVersionEnvVarJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTypeVersionEnvVarJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTypeVersionEnvVarJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectTypeVersionEnvVarJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectTypeVersionEnvVarJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectTypeVersionEnvVarJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetKey

`func (o *ProjectTypeVersionEnvVarJsonhal) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectTypeVersionEnvVarJsonhal) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectTypeVersionEnvVarJsonhal) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetNote

`func (o *ProjectTypeVersionEnvVarJsonhal) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ProjectTypeVersionEnvVarJsonhal) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ProjectTypeVersionEnvVarJsonhal) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetIsRequired

`func (o *ProjectTypeVersionEnvVarJsonhal) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ProjectTypeVersionEnvVarJsonhal) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ProjectTypeVersionEnvVarJsonhal) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetIsEncrypted

`func (o *ProjectTypeVersionEnvVarJsonhal) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *ProjectTypeVersionEnvVarJsonhal) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *ProjectTypeVersionEnvVarJsonhal) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetProjectTypeVersion

`func (o *ProjectTypeVersionEnvVarJsonhal) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ProjectTypeVersionEnvVarJsonhal) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ProjectTypeVersionEnvVarJsonhal) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectTypeVersionEnvVarExample

`func (o *ProjectTypeVersionEnvVarJsonhal) GetProjectTypeVersionEnvVarExample() []string`

GetProjectTypeVersionEnvVarExample returns the ProjectTypeVersionEnvVarExample field if non-nil, zero value otherwise.

### GetProjectTypeVersionEnvVarExampleOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetProjectTypeVersionEnvVarExampleOk() (*[]string, bool)`

GetProjectTypeVersionEnvVarExampleOk returns a tuple with the ProjectTypeVersionEnvVarExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersionEnvVarExample

`func (o *ProjectTypeVersionEnvVarJsonhal) SetProjectTypeVersionEnvVarExample(v []string)`

SetProjectTypeVersionEnvVarExample sets ProjectTypeVersionEnvVarExample field to given value.

### HasProjectTypeVersionEnvVarExample

`func (o *ProjectTypeVersionEnvVarJsonhal) HasProjectTypeVersionEnvVarExample() bool`

HasProjectTypeVersionEnvVarExample returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ProjectTypeVersionEnvVarJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectTypeVersionEnvVarJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectTypeVersionEnvVarJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ProjectTypeVersionEnvVarJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ProjectTypeVersionEnvVarJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ProjectTypeVersionEnvVarJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ProjectTypeVersionEnvVarJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ProjectTypeVersionEnvVarJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ProjectTypeVersionEnvVarJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectTypeVersionEnvVarJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectTypeVersionEnvVarJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectTypeVersionEnvVarJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectTypeVersionEnvVarJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectTypeVersionEnvVarJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectTypeVersionEnvVarJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


