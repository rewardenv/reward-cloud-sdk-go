# ProjectTypeVersionJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
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

### NewProjectTypeVersionJsonhal

`func NewProjectTypeVersionJsonhal() *ProjectTypeVersionJsonhal`

NewProjectTypeVersionJsonhal instantiates a new ProjectTypeVersionJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTypeVersionJsonhalWithDefaults

`func NewProjectTypeVersionJsonhalWithDefaults() *ProjectTypeVersionJsonhal`

NewProjectTypeVersionJsonhalWithDefaults instantiates a new ProjectTypeVersionJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectTypeVersionJsonhal) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectTypeVersionJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectTypeVersionJsonhal) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectTypeVersionJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProjectTypeVersionJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTypeVersionJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTypeVersionJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTypeVersionJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectTypeVersionJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectTypeVersionJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectTypeVersionJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectTypeVersionJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectTypeVersionJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectTypeVersionJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ProjectTypeVersionJsonhal) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProjectTypeVersionJsonhal) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProjectTypeVersionJsonhal) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProjectTypeVersionJsonhal) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ProjectTypeVersionJsonhal) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ProjectTypeVersionJsonhal) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetProject

`func (o *ProjectTypeVersionJsonhal) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectTypeVersionJsonhal) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectTypeVersionJsonhal) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectTypeVersionJsonhal) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectType

`func (o *ProjectTypeVersionJsonhal) GetProjectType() string`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *ProjectTypeVersionJsonhal) GetProjectTypeOk() (*string, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *ProjectTypeVersionJsonhal) SetProjectType(v string)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *ProjectTypeVersionJsonhal) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### SetProjectTypeNil

`func (o *ProjectTypeVersionJsonhal) SetProjectTypeNil(b bool)`

 SetProjectTypeNil sets the value for ProjectType to be an explicit nil

### UnsetProjectType
`func (o *ProjectTypeVersionJsonhal) UnsetProjectType()`

UnsetProjectType ensures that no value is present for ProjectType, not even an explicit nil
### GetProjectTypeVersionEnvVar

`func (o *ProjectTypeVersionJsonhal) GetProjectTypeVersionEnvVar() []string`

GetProjectTypeVersionEnvVar returns the ProjectTypeVersionEnvVar field if non-nil, zero value otherwise.

### GetProjectTypeVersionEnvVarOk

`func (o *ProjectTypeVersionJsonhal) GetProjectTypeVersionEnvVarOk() (*[]string, bool)`

GetProjectTypeVersionEnvVarOk returns a tuple with the ProjectTypeVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersionEnvVar

`func (o *ProjectTypeVersionJsonhal) SetProjectTypeVersionEnvVar(v []string)`

SetProjectTypeVersionEnvVar sets ProjectTypeVersionEnvVar field to given value.

### HasProjectTypeVersionEnvVar

`func (o *ProjectTypeVersionJsonhal) HasProjectTypeVersionEnvVar() bool`

HasProjectTypeVersionEnvVar returns a boolean if a field has been set.

### GetComponentResourceLimit

`func (o *ProjectTypeVersionJsonhal) GetComponentResourceLimit() []string`

GetComponentResourceLimit returns the ComponentResourceLimit field if non-nil, zero value otherwise.

### GetComponentResourceLimitOk

`func (o *ProjectTypeVersionJsonhal) GetComponentResourceLimitOk() (*[]string, bool)`

GetComponentResourceLimitOk returns a tuple with the ComponentResourceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentResourceLimit

`func (o *ProjectTypeVersionJsonhal) SetComponentResourceLimit(v []string)`

SetComponentResourceLimit sets ComponentResourceLimit field to given value.

### HasComponentResourceLimit

`func (o *ProjectTypeVersionJsonhal) HasComponentResourceLimit() bool`

HasComponentResourceLimit returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ProjectTypeVersionJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectTypeVersionJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectTypeVersionJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectTypeVersionJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProjectTypeVersionJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectTypeVersionJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ProjectTypeVersionJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ProjectTypeVersionJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ProjectTypeVersionJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ProjectTypeVersionJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ProjectTypeVersionJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ProjectTypeVersionJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ProjectTypeVersionJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectTypeVersionJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectTypeVersionJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectTypeVersionJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectTypeVersionJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectTypeVersionJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectTypeVersionJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectTypeVersionJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


