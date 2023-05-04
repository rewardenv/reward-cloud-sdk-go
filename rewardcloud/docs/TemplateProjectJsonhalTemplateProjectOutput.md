# TemplateProjectJsonhalTemplateProjectOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
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
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarJsonhalTemplateProjectOutput**](ProjectEnvVarJsonhalTemplateProjectOutput.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTemplateProjectJsonhalTemplateProjectOutput

`func NewTemplateProjectJsonhalTemplateProjectOutput() *TemplateProjectJsonhalTemplateProjectOutput`

NewTemplateProjectJsonhalTemplateProjectOutput instantiates a new TemplateProjectJsonhalTemplateProjectOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateProjectJsonhalTemplateProjectOutputWithDefaults

`func NewTemplateProjectJsonhalTemplateProjectOutputWithDefaults() *TemplateProjectJsonhalTemplateProjectOutput`

NewTemplateProjectJsonhalTemplateProjectOutputWithDefaults instantiates a new TemplateProjectJsonhalTemplateProjectOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateEnvironment

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetTemplateEnvironment() []string`

GetTemplateEnvironment returns the TemplateEnvironment field if non-nil, zero value otherwise.

### GetTemplateEnvironmentOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetTemplateEnvironmentOk() (*[]string, bool)`

GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironment

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetTemplateEnvironment(v []string)`

SetTemplateEnvironment sets TemplateEnvironment field to given value.

### HasTemplateEnvironment

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasTemplateEnvironment() bool`

HasTemplateEnvironment returns a boolean if a field has been set.

### GetId

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *TemplateProjectJsonhalTemplateProjectOutput) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetProjectEnvVar() []ProjectEnvVarJsonhalTemplateProjectOutput`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetProjectEnvVarOk() (*[]ProjectEnvVarJsonhalTemplateProjectOutput, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetProjectEnvVar(v []ProjectEnvVarJsonhalTemplateProjectOutput)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateProjectJsonhalTemplateProjectOutput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateProjectJsonhalTemplateProjectOutput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateProjectJsonhalTemplateProjectOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


