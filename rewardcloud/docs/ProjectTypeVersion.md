# ProjectTypeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **[]string** |  | [optional] 
**ProjectType** | Pointer to **NullableString** |  | [optional] 
**ProjectTypeVersionEnvVar** | Pointer to **[]string** |  | [optional] 
**ComponentResourceLimit** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectTypeVersion

`func NewProjectTypeVersion() *ProjectTypeVersion`

NewProjectTypeVersion instantiates a new ProjectTypeVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTypeVersionWithDefaults

`func NewProjectTypeVersionWithDefaults() *ProjectTypeVersion`

NewProjectTypeVersionWithDefaults instantiates a new ProjectTypeVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTypeVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTypeVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTypeVersion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTypeVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectTypeVersion) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectTypeVersion) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectTypeVersion) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectTypeVersion) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectTypeVersion) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectTypeVersion) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ProjectTypeVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectTypeVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectTypeVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProjectTypeVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProjectTypeVersion) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProjectTypeVersion) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetProject

`func (o *ProjectTypeVersion) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectTypeVersion) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectTypeVersion) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectTypeVersion) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectType

`func (o *ProjectTypeVersion) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *ProjectTypeVersion) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *ProjectTypeVersion) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *ProjectTypeVersion) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### SetProjectTypeNil

`func (o *ProjectTypeVersion) SetProjectTypeNil(b bool)`

 SetProjectTypeNil sets the value for ProjectType to be an explicit nil

### UnsetProjectType
`func (o *ProjectTypeVersion) UnsetProjectType()`

UnsetProjectType ensures that no value is present for ProjectType, not even an explicit nil
### GetProjectTypeVersionEnvVar

`func (o *ProjectTypeVersion) GetProjectTypeVersionEnvVar() []string`

GetProjectTypeVersionEnvVar returns the ProjectTypeVersionEnvVar field if non-nil, zero value otherwise.

### GetProjectTypeVersionEnvVarOk

`func (o *ProjectTypeVersion) GetProjectTypeVersionEnvVarOk() (*[]string, bool)`

GetProjectTypeVersionEnvVarOk returns a tuple with the ProjectTypeVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersionEnvVar

`func (o *ProjectTypeVersion) SetProjectTypeVersionEnvVar(v []string)`

SetProjectTypeVersionEnvVar sets ProjectTypeVersionEnvVar field to given value.

### HasProjectTypeVersionEnvVar

`func (o *ProjectTypeVersion) HasProjectTypeVersionEnvVar() bool`

HasProjectTypeVersionEnvVar returns a boolean if a field has been set.

### GetComponentResourceLimit

`func (o *ProjectTypeVersion) GetComponentResourceLimit() []string`

GetComponentResourceLimit returns the ComponentResourceLimit field if non-nil, zero value otherwise.

### GetComponentResourceLimitOk

`func (o *ProjectTypeVersion) GetComponentResourceLimitOk() (*[]string, bool)`

GetComponentResourceLimitOk returns a tuple with the ComponentResourceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentResourceLimit

`func (o *ProjectTypeVersion) SetComponentResourceLimit(v []string)`

SetComponentResourceLimit sets ComponentResourceLimit field to given value.

### HasComponentResourceLimit

`func (o *ProjectTypeVersion) HasComponentResourceLimit() bool`

HasComponentResourceLimit returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ProjectTypeVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectTypeVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectTypeVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectTypeVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProjectTypeVersion) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectTypeVersion) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ProjectTypeVersion) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ProjectTypeVersion) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ProjectTypeVersion) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ProjectTypeVersion) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ProjectTypeVersion) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ProjectTypeVersion) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ProjectTypeVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectTypeVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectTypeVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectTypeVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectTypeVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectTypeVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectTypeVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectTypeVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


