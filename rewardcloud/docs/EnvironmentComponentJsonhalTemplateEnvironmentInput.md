# EnvironmentComponentJsonhalTemplateEnvironmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**Node** | Pointer to **NullableInt32** |  | [optional] 
**ComponentVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEnvironmentComponentJsonhalTemplateEnvironmentInput

`func NewEnvironmentComponentJsonhalTemplateEnvironmentInput() *EnvironmentComponentJsonhalTemplateEnvironmentInput`

NewEnvironmentComponentJsonhalTemplateEnvironmentInput instantiates a new EnvironmentComponentJsonhalTemplateEnvironmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentComponentJsonhalTemplateEnvironmentInputWithDefaults

`func NewEnvironmentComponentJsonhalTemplateEnvironmentInputWithDefaults() *EnvironmentComponentJsonhalTemplateEnvironmentInput`

NewEnvironmentComponentJsonhalTemplateEnvironmentInputWithDefaults instantiates a new EnvironmentComponentJsonhalTemplateEnvironmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCpu

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetNode

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetNode() int32`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetNodeOk() (*int32, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetNode(v int32)`

SetNode sets Node field to given value.

### HasNode

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) HasNode() bool`

HasNode returns a boolean if a field has been set.

### SetNodeNil

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetComponentVersion

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *EnvironmentComponentJsonhalTemplateEnvironmentInput) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


