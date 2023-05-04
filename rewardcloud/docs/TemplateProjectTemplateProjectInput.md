# TemplateProjectTemplateProjectInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**TemplateEnvironment** | Pointer to **[]string** |  | [optional] 
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
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarTemplateProjectInput**](ProjectEnvVarTemplateProjectInput.md) |  | [optional] 

## Methods

### NewTemplateProjectTemplateProjectInput

`func NewTemplateProjectTemplateProjectInput() *TemplateProjectTemplateProjectInput`

NewTemplateProjectTemplateProjectInput instantiates a new TemplateProjectTemplateProjectInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateProjectTemplateProjectInputWithDefaults

`func NewTemplateProjectTemplateProjectInputWithDefaults() *TemplateProjectTemplateProjectInput`

NewTemplateProjectTemplateProjectInputWithDefaults instantiates a new TemplateProjectTemplateProjectInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TemplateProjectTemplateProjectInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateProjectTemplateProjectInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateProjectTemplateProjectInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateProjectTemplateProjectInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateProjectTemplateProjectInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateProjectTemplateProjectInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateEnvironment

`func (o *TemplateProjectTemplateProjectInput) GetTemplateEnvironment() []string`

GetTemplateEnvironment returns the TemplateEnvironment field if non-nil, zero value otherwise.

### GetTemplateEnvironmentOk

`func (o *TemplateProjectTemplateProjectInput) GetTemplateEnvironmentOk() (*[]string, bool)`

GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironment

`func (o *TemplateProjectTemplateProjectInput) SetTemplateEnvironment(v []string)`

SetTemplateEnvironment sets TemplateEnvironment field to given value.

### HasTemplateEnvironment

`func (o *TemplateProjectTemplateProjectInput) HasTemplateEnvironment() bool`

HasTemplateEnvironment returns a boolean if a field has been set.

### GetName

`func (o *TemplateProjectTemplateProjectInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateProjectTemplateProjectInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateProjectTemplateProjectInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateProjectTemplateProjectInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateProjectTemplateProjectInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateProjectTemplateProjectInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *TemplateProjectTemplateProjectInput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TemplateProjectTemplateProjectInput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TemplateProjectTemplateProjectInput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TemplateProjectTemplateProjectInput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *TemplateProjectTemplateProjectInput) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *TemplateProjectTemplateProjectInput) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *TemplateProjectTemplateProjectInput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateProjectTemplateProjectInput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateProjectTemplateProjectInput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateProjectTemplateProjectInput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateProjectTemplateProjectInput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateProjectTemplateProjectInput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateProjectTemplateProjectInput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateProjectTemplateProjectInput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateProjectTemplateProjectInput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateProjectTemplateProjectInput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateProjectTemplateProjectInput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateProjectTemplateProjectInput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateProjectTemplateProjectInput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateProjectTemplateProjectInput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateProjectTemplateProjectInput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateProjectTemplateProjectInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateProjectTemplateProjectInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateProjectTemplateProjectInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *TemplateProjectTemplateProjectInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TemplateProjectTemplateProjectInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TemplateProjectTemplateProjectInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TemplateProjectTemplateProjectInput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TemplateProjectTemplateProjectInput) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TemplateProjectTemplateProjectInput) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *TemplateProjectTemplateProjectInput) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TemplateProjectTemplateProjectInput) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TemplateProjectTemplateProjectInput) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TemplateProjectTemplateProjectInput) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TemplateProjectTemplateProjectInput) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TemplateProjectTemplateProjectInput) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *TemplateProjectTemplateProjectInput) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *TemplateProjectTemplateProjectInput) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *TemplateProjectTemplateProjectInput) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *TemplateProjectTemplateProjectInput) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *TemplateProjectTemplateProjectInput) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *TemplateProjectTemplateProjectInput) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *TemplateProjectTemplateProjectInput) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *TemplateProjectTemplateProjectInput) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *TemplateProjectTemplateProjectInput) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *TemplateProjectTemplateProjectInput) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *TemplateProjectTemplateProjectInput) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *TemplateProjectTemplateProjectInput) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *TemplateProjectTemplateProjectInput) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *TemplateProjectTemplateProjectInput) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *TemplateProjectTemplateProjectInput) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *TemplateProjectTemplateProjectInput) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *TemplateProjectTemplateProjectInput) GetProjectEnvVar() []ProjectEnvVarTemplateProjectInput`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *TemplateProjectTemplateProjectInput) GetProjectEnvVarOk() (*[]ProjectEnvVarTemplateProjectInput, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *TemplateProjectTemplateProjectInput) SetProjectEnvVar(v []ProjectEnvVarTemplateProjectInput)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *TemplateProjectTemplateProjectInput) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


