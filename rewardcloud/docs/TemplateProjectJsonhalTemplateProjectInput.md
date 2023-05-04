# TemplateProjectJsonhalTemplateProjectInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
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
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarJsonhalTemplateProjectInput**](ProjectEnvVarJsonhalTemplateProjectInput.md) |  | [optional] 

## Methods

### NewTemplateProjectJsonhalTemplateProjectInput

`func NewTemplateProjectJsonhalTemplateProjectInput() *TemplateProjectJsonhalTemplateProjectInput`

NewTemplateProjectJsonhalTemplateProjectInput instantiates a new TemplateProjectJsonhalTemplateProjectInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateProjectJsonhalTemplateProjectInputWithDefaults

`func NewTemplateProjectJsonhalTemplateProjectInputWithDefaults() *TemplateProjectJsonhalTemplateProjectInput`

NewTemplateProjectJsonhalTemplateProjectInputWithDefaults instantiates a new TemplateProjectJsonhalTemplateProjectInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTemplateEnvironment

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetTemplateEnvironment() []string`

GetTemplateEnvironment returns the TemplateEnvironment field if non-nil, zero value otherwise.

### GetTemplateEnvironmentOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetTemplateEnvironmentOk() (*[]string, bool)`

GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateEnvironment

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetTemplateEnvironment(v []string)`

SetTemplateEnvironment sets TemplateEnvironment field to given value.

### HasTemplateEnvironment

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasTemplateEnvironment() bool`

HasTemplateEnvironment returns a boolean if a field has been set.

### GetName

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectEnvVar() []ProjectEnvVarJsonhalTemplateProjectInput`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectEnvVarOk() (*[]ProjectEnvVarJsonhalTemplateProjectInput, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *TemplateProjectJsonhalTemplateProjectInput) SetProjectEnvVar(v []ProjectEnvVarJsonhalTemplateProjectInput)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *TemplateProjectJsonhalTemplateProjectInput) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


