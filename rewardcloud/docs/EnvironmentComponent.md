# EnvironmentComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**Node** | Pointer to **NullableInt32** |  | [optional] 
**ComponentVersion** | Pointer to **NullableString** |  | [optional] 
**Environment** | Pointer to [**NullableEnvironmentComponentEnvironment**](EnvironmentComponentEnvironment.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvironmentComponent

`func NewEnvironmentComponent() *EnvironmentComponent`

NewEnvironmentComponent instantiates a new EnvironmentComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentComponentWithDefaults

`func NewEnvironmentComponentWithDefaults() *EnvironmentComponent`

NewEnvironmentComponentWithDefaults instantiates a new EnvironmentComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentComponent) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentComponent) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentComponent) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentComponent) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentComponent) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentComponent) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetCpu

`func (o *EnvironmentComponent) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentComponent) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentComponent) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentComponent) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentComponent) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentComponent) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentComponent) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentComponent) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentComponent) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentComponent) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentComponent) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentComponent) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentComponent) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentComponent) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentComponent) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentComponent) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentComponent) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentComponent) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetNode

`func (o *EnvironmentComponent) GetNode() int32`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *EnvironmentComponent) GetNodeOk() (*int32, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *EnvironmentComponent) SetNode(v int32)`

SetNode sets Node field to given value.

### HasNode

`func (o *EnvironmentComponent) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *EnvironmentComponent) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *EnvironmentComponent) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetComponentVersion

`func (o *EnvironmentComponent) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *EnvironmentComponent) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *EnvironmentComponent) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *EnvironmentComponent) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *EnvironmentComponent) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *EnvironmentComponent) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetEnvironment

`func (o *EnvironmentComponent) GetEnvironment() EnvironmentComponentEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentComponent) GetEnvironmentOk() (*EnvironmentComponentEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentComponent) SetEnvironment(v EnvironmentComponentEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentComponent) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvironmentComponent) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvironmentComponent) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentComponent) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentComponent) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentComponent) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentComponent) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentComponent) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentComponent) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentComponent) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentComponent) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentComponent) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentComponent) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentComponent) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentComponent) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentComponent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentComponent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentComponent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentComponent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentComponent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentComponent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentComponent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentComponent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


