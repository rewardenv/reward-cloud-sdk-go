# ProjectEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**Project** | Pointer to [**NullableProjectEnvVarProject**](ProjectEnvVarProject.md) |  | [optional] 
**EnvVarType** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewProjectEnvVar

`func NewProjectEnvVar() *ProjectEnvVar`

NewProjectEnvVar instantiates a new ProjectEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEnvVarWithDefaults

`func NewProjectEnvVarWithDefaults() *ProjectEnvVar`

NewProjectEnvVarWithDefaults instantiates a new ProjectEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectEnvVar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectEnvVar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectEnvVar) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectEnvVar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectEnvVar) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectEnvVar) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectEnvVar) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectEnvVar) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectEnvVar) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectEnvVar) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetKey

`func (o *ProjectEnvVar) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectEnvVar) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectEnvVar) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectEnvVar) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ProjectEnvVar) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ProjectEnvVar) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ProjectEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectEnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProjectEnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProjectEnvVar) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProjectEnvVar) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEncrypted

`func (o *ProjectEnvVar) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *ProjectEnvVar) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *ProjectEnvVar) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *ProjectEnvVar) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *ProjectEnvVar) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *ProjectEnvVar) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetProject

`func (o *ProjectEnvVar) GetProject() ProjectEnvVarProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectEnvVar) GetProjectOk() (*ProjectEnvVarProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectEnvVar) SetProject(v ProjectEnvVarProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectEnvVar) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *ProjectEnvVar) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *ProjectEnvVar) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetEnvVarType

`func (o *ProjectEnvVar) GetEnvVarType() string`

GetEnvVarType returns the EnvVarType field if non-nil, zero value otherwise.

### GetEnvVarTypeOk

`func (o *ProjectEnvVar) GetEnvVarTypeOk() (*string, bool)`

GetEnvVarTypeOk returns a tuple with the EnvVarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarType

`func (o *ProjectEnvVar) SetEnvVarType(v string)`

SetEnvVarType sets EnvVarType field to given value.

### HasEnvVarType

`func (o *ProjectEnvVar) HasEnvVarType() bool`

HasEnvVarType returns a boolean if a field has been set.

### SetEnvVarTypeNil

`func (o *ProjectEnvVar) SetEnvVarTypeNil(b bool)`

 SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil

### UnsetEnvVarType
`func (o *ProjectEnvVar) UnsetEnvVarType()`

UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
### GetCreatedBy

`func (o *ProjectEnvVar) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectEnvVar) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectEnvVar) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectEnvVar) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProjectEnvVar) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectEnvVar) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ProjectEnvVar) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ProjectEnvVar) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ProjectEnvVar) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ProjectEnvVar) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ProjectEnvVar) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ProjectEnvVar) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ProjectEnvVar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectEnvVar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectEnvVar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectEnvVar) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectEnvVar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectEnvVar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectEnvVar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectEnvVar) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRawValue

`func (o *ProjectEnvVar) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *ProjectEnvVar) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *ProjectEnvVar) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *ProjectEnvVar) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *ProjectEnvVar) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *ProjectEnvVar) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


