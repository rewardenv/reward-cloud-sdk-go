# TemplateProjectTemplateProjectOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**TemplateEnvironment** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**IsInitProjectSkeleton** | Pointer to **NullableBool** |  | [optional] 
**ComponentVersion** | Pointer to **[]string** |  | [optional] 
**ProjectTypeVersion** | Pointer to **NullableString** |  | [optional] 
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarTemplateProjectOutput**](ProjectEnvVarTemplateProjectOutput.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTemplateProjectTemplateProjectOutput

`func NewTemplateProjectTemplateProjectOutput() *TemplateProjectTemplateProjectOutput`

NewTemplateProjectTemplateProjectOutput instantiates a new TemplateProjectTemplateProjectOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateProjectTemplateProjectOutputWithDefaults

`func NewTemplateProjectTemplateProjectOutputWithDefaults() *TemplateProjectTemplateProjectOutput`

NewTemplateProjectTemplateProjectOutputWithDefaults instantiates a new TemplateProjectTemplateProjectOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TemplateProjectTemplateProjectOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateProjectTemplateProjectOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateProjectTemplateProjectOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateProjectTemplateProjectOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateProjectTemplateProjectOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateProjectTemplateProjectOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateEnvironment

`func (o *TemplateProjectTemplateProjectOutput) GetTemplateEnvironment() []string`

GetTemplateEnvironment returns the TemplateEnvironment field if non-nil, zero value otherwise.

### GetTemplateEnvironmentOk

`func (o *TemplateProjectTemplateProjectOutput) GetTemplateEnvironmentOk() (*[]string, bool)`

GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironment

`func (o *TemplateProjectTemplateProjectOutput) SetTemplateEnvironment(v []string)`

SetTemplateEnvironment sets TemplateEnvironment field to given value.

### HasTemplateEnvironment

`func (o *TemplateProjectTemplateProjectOutput) HasTemplateEnvironment() bool`

HasTemplateEnvironment returns a boolean if a field has been set.

### GetId

`func (o *TemplateProjectTemplateProjectOutput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateProjectTemplateProjectOutput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateProjectTemplateProjectOutput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateProjectTemplateProjectOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TemplateProjectTemplateProjectOutput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TemplateProjectTemplateProjectOutput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TemplateProjectTemplateProjectOutput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TemplateProjectTemplateProjectOutput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *TemplateProjectTemplateProjectOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateProjectTemplateProjectOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateProjectTemplateProjectOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateProjectTemplateProjectOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateProjectTemplateProjectOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateProjectTemplateProjectOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *TemplateProjectTemplateProjectOutput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TemplateProjectTemplateProjectOutput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TemplateProjectTemplateProjectOutput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TemplateProjectTemplateProjectOutput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *TemplateProjectTemplateProjectOutput) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *TemplateProjectTemplateProjectOutput) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *TemplateProjectTemplateProjectOutput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateProjectTemplateProjectOutput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateProjectTemplateProjectOutput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateProjectTemplateProjectOutput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateProjectTemplateProjectOutput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateProjectTemplateProjectOutput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateProjectTemplateProjectOutput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateProjectTemplateProjectOutput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateProjectTemplateProjectOutput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateProjectTemplateProjectOutput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateProjectTemplateProjectOutput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateProjectTemplateProjectOutput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateProjectTemplateProjectOutput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateProjectTemplateProjectOutput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateProjectTemplateProjectOutput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateProjectTemplateProjectOutput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateProjectTemplateProjectOutput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateProjectTemplateProjectOutput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *TemplateProjectTemplateProjectOutput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TemplateProjectTemplateProjectOutput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TemplateProjectTemplateProjectOutput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TemplateProjectTemplateProjectOutput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TemplateProjectTemplateProjectOutput) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TemplateProjectTemplateProjectOutput) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *TemplateProjectTemplateProjectOutput) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TemplateProjectTemplateProjectOutput) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TemplateProjectTemplateProjectOutput) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TemplateProjectTemplateProjectOutput) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TemplateProjectTemplateProjectOutput) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TemplateProjectTemplateProjectOutput) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *TemplateProjectTemplateProjectOutput) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *TemplateProjectTemplateProjectOutput) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *TemplateProjectTemplateProjectOutput) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *TemplateProjectTemplateProjectOutput) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *TemplateProjectTemplateProjectOutput) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *TemplateProjectTemplateProjectOutput) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *TemplateProjectTemplateProjectOutput) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *TemplateProjectTemplateProjectOutput) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *TemplateProjectTemplateProjectOutput) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *TemplateProjectTemplateProjectOutput) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *TemplateProjectTemplateProjectOutput) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *TemplateProjectTemplateProjectOutput) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *TemplateProjectTemplateProjectOutput) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *TemplateProjectTemplateProjectOutput) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *TemplateProjectTemplateProjectOutput) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *TemplateProjectTemplateProjectOutput) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *TemplateProjectTemplateProjectOutput) GetProjectEnvVar() []ProjectEnvVarTemplateProjectOutput`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *TemplateProjectTemplateProjectOutput) GetProjectEnvVarOk() (*[]ProjectEnvVarTemplateProjectOutput, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *TemplateProjectTemplateProjectOutput) SetProjectEnvVar(v []ProjectEnvVarTemplateProjectOutput)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *TemplateProjectTemplateProjectOutput) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateProjectTemplateProjectOutput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateProjectTemplateProjectOutput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateProjectTemplateProjectOutput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateProjectTemplateProjectOutput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateProjectTemplateProjectOutput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateProjectTemplateProjectOutput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateProjectTemplateProjectOutput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateProjectTemplateProjectOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


