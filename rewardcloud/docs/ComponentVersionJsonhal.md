# ComponentVersionJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Component** | Pointer to **NullableString** |  | [optional] 
**ComponentVersionEnvVar** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to [**[]AbstractProjectJsonhal**](AbstractProjectJsonhal.md) |  | [optional] 
**EnvironmentComponent** | Pointer to **[]string** |  | [optional] 
**ComponentResourceLimit** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComponentVersionJsonhal

`func NewComponentVersionJsonhal() *ComponentVersionJsonhal`

NewComponentVersionJsonhal instantiates a new ComponentVersionJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionJsonhalWithDefaults

`func NewComponentVersionJsonhalWithDefaults() *ComponentVersionJsonhal`

NewComponentVersionJsonhalWithDefaults instantiates a new ComponentVersionJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ComponentVersionJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ComponentVersionJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ComponentVersionJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ComponentVersionJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ComponentVersionJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentVersionJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentVersionJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentVersionJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ComponentVersionJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComponentVersionJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComponentVersionJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComponentVersionJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ComponentVersionJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ComponentVersionJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetVersion

`func (o *ComponentVersionJsonhal) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ComponentVersionJsonhal) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ComponentVersionJsonhal) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ComponentVersionJsonhal) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ComponentVersionJsonhal) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ComponentVersionJsonhal) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetComponent

`func (o *ComponentVersionJsonhal) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ComponentVersionJsonhal) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ComponentVersionJsonhal) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ComponentVersionJsonhal) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### SetComponentNil

`func (o *ComponentVersionJsonhal) SetComponentNil(b bool)`

 SetComponentNil sets the value for Component to be an explicit nil

### UnsetComponent
`func (o *ComponentVersionJsonhal) UnsetComponent()`

UnsetComponent ensures that no value is present for Component, not even an explicit nil
### GetComponentVersionEnvVar

`func (o *ComponentVersionJsonhal) GetComponentVersionEnvVar() []string`

GetComponentVersionEnvVar returns the ComponentVersionEnvVar field if non-nil, zero value otherwise.

### GetComponentVersionEnvVarOk

`func (o *ComponentVersionJsonhal) GetComponentVersionEnvVarOk() (*[]string, bool)`

GetComponentVersionEnvVarOk returns a tuple with the ComponentVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersionEnvVar

`func (o *ComponentVersionJsonhal) SetComponentVersionEnvVar(v []string)`

SetComponentVersionEnvVar sets ComponentVersionEnvVar field to given value.

### HasComponentVersionEnvVar

`func (o *ComponentVersionJsonhal) HasComponentVersionEnvVar() bool`

HasComponentVersionEnvVar returns a boolean if a field has been set.

### GetProject

`func (o *ComponentVersionJsonhal) GetProject() []AbstractProjectJsonhal`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ComponentVersionJsonhal) GetProjectOk() (*[]AbstractProjectJsonhal, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ComponentVersionJsonhal) SetProject(v []AbstractProjectJsonhal)`

SetProject sets Project field to given value.

### HasProject

`func (o *ComponentVersionJsonhal) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *ComponentVersionJsonhal) GetEnvironmentComponent() []string`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *ComponentVersionJsonhal) GetEnvironmentComponentOk() (*[]string, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *ComponentVersionJsonhal) SetEnvironmentComponent(v []string)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *ComponentVersionJsonhal) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.

### GetComponentResourceLimit

`func (o *ComponentVersionJsonhal) GetComponentResourceLimit() []string`

GetComponentResourceLimit returns the ComponentResourceLimit field if non-nil, zero value otherwise.

### GetComponentResourceLimitOk

`func (o *ComponentVersionJsonhal) GetComponentResourceLimitOk() (*[]string, bool)`

GetComponentResourceLimitOk returns a tuple with the ComponentResourceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentResourceLimit

`func (o *ComponentVersionJsonhal) SetComponentResourceLimit(v []string)`

SetComponentResourceLimit sets ComponentResourceLimit field to given value.

### HasComponentResourceLimit

`func (o *ComponentVersionJsonhal) HasComponentResourceLimit() bool`

HasComponentResourceLimit returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ComponentVersionJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ComponentVersionJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ComponentVersionJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ComponentVersionJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ComponentVersionJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ComponentVersionJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ComponentVersionJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ComponentVersionJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ComponentVersionJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ComponentVersionJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ComponentVersionJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ComponentVersionJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ComponentVersionJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentVersionJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentVersionJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComponentVersionJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComponentVersionJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComponentVersionJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComponentVersionJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComponentVersionJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


