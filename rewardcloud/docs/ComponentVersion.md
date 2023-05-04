# ComponentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Component** | Pointer to **NullableString** |  | [optional] 
**ComponentVersionEnvVar** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to [**[]AbstractProject**](AbstractProject.md) |  | [optional] 
**EnvironmentComponent** | Pointer to **[]string** |  | [optional] 
**ComponentResourceLimit** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComponentVersion

`func NewComponentVersion() *ComponentVersion`

NewComponentVersion instantiates a new ComponentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionWithDefaults

`func NewComponentVersionWithDefaults() *ComponentVersion`

NewComponentVersionWithDefaults instantiates a new ComponentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentVersion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ComponentVersion) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComponentVersion) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComponentVersion) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComponentVersion) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ComponentVersion) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ComponentVersion) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ComponentVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComponentVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ComponentVersion) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ComponentVersion) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetComponent

`func (o *ComponentVersion) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentVersion) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentVersion) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ComponentVersion) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *ComponentVersion) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *ComponentVersion) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetComponentVersionEnvVar

`func (o *ComponentVersion) GetComponentVersionEnvVar() []string`

GetComponentVersionEnvVar returns the ComponentVersionEnvVar field if non-nil, zero value otherwise.

### GetComponentVersionEnvVarOk

`func (o *ComponentVersion) GetComponentVersionEnvVarOk() (*[]string, bool)`

GetComponentVersionEnvVarOk returns a tuple with the ComponentVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersionEnvVar

`func (o *ComponentVersion) SetComponentVersionEnvVar(v []string)`

SetComponentVersionEnvVar sets ComponentVersionEnvVar field to given value.

### HasComponentVersionEnvVar

`func (o *ComponentVersion) HasComponentVersionEnvVar() bool`

HasComponentVersionEnvVar returns a boolean if a field has been set.

### GetProject

`func (o *ComponentVersion) GetProject() []AbstractProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ComponentVersion) GetProjectOk() (*[]AbstractProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ComponentVersion) SetProject(v []AbstractProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ComponentVersion) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *ComponentVersion) GetEnvironmentComponent() []string`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *ComponentVersion) GetEnvironmentComponentOk() (*[]string, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *ComponentVersion) SetEnvironmentComponent(v []string)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *ComponentVersion) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.

### GetComponentResourceLimit

`func (o *ComponentVersion) GetComponentResourceLimit() []string`

GetComponentResourceLimit returns the ComponentResourceLimit field if non-nil, zero value otherwise.

### GetComponentResourceLimitOk

`func (o *ComponentVersion) GetComponentResourceLimitOk() (*[]string, bool)`

GetComponentResourceLimitOk returns a tuple with the ComponentResourceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentResourceLimit

`func (o *ComponentVersion) SetComponentResourceLimit(v []string)`

SetComponentResourceLimit sets ComponentResourceLimit field to given value.

### HasComponentResourceLimit

`func (o *ComponentVersion) HasComponentResourceLimit() bool`

HasComponentResourceLimit returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ComponentVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ComponentVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ComponentVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ComponentVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ComponentVersion) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ComponentVersion) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ComponentVersion) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ComponentVersion) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ComponentVersion) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ComponentVersion) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ComponentVersion) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ComponentVersion) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ComponentVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComponentVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComponentVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComponentVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComponentVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComponentVersion) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


