# ProjectEnvVarProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**IsInitProjectSkeleton** | Pointer to **NullableBool** |  | [optional] 
**ComponentVersion** | Pointer to [**[]ComponentVersion**](ComponentVersion.md) |  | [optional] 
**ProjectTypeVersion** | Pointer to **NullableString** |  | [optional] 
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVar**](ProjectEnvVar.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectEnvVarProject

`func NewProjectEnvVarProject() *ProjectEnvVarProject`

NewProjectEnvVarProject instantiates a new ProjectEnvVarProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEnvVarProjectWithDefaults

`func NewProjectEnvVarProjectWithDefaults() *ProjectEnvVarProject`

NewProjectEnvVarProjectWithDefaults instantiates a new ProjectEnvVarProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectEnvVarProject) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectEnvVarProject) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectEnvVarProject) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectEnvVarProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectEnvVarProject) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectEnvVarProject) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectEnvVarProject) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectEnvVarProject) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ProjectEnvVarProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectEnvVarProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectEnvVarProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectEnvVarProject) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectEnvVarProject) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectEnvVarProject) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *ProjectEnvVarProject) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *ProjectEnvVarProject) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *ProjectEnvVarProject) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *ProjectEnvVarProject) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *ProjectEnvVarProject) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *ProjectEnvVarProject) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsActive

`func (o *ProjectEnvVarProject) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectEnvVarProject) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectEnvVarProject) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProjectEnvVarProject) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProjectEnvVarProject) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProjectEnvVarProject) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *ProjectEnvVarProject) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProjectEnvVarProject) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProjectEnvVarProject) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProjectEnvVarProject) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ProjectEnvVarProject) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ProjectEnvVarProject) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ProjectEnvVarProject) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectEnvVarProject) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectEnvVarProject) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ProjectEnvVarProject) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ProjectEnvVarProject) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ProjectEnvVarProject) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *ProjectEnvVarProject) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectEnvVarProject) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectEnvVarProject) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProjectEnvVarProject) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ProjectEnvVarProject) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ProjectEnvVarProject) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *ProjectEnvVarProject) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectEnvVarProject) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectEnvVarProject) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectEnvVarProject) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ProjectEnvVarProject) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ProjectEnvVarProject) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *ProjectEnvVarProject) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ProjectEnvVarProject) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ProjectEnvVarProject) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ProjectEnvVarProject) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ProjectEnvVarProject) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ProjectEnvVarProject) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *ProjectEnvVarProject) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *ProjectEnvVarProject) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *ProjectEnvVarProject) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *ProjectEnvVarProject) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *ProjectEnvVarProject) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *ProjectEnvVarProject) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *ProjectEnvVarProject) GetComponentVersion() []ComponentVersion`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ProjectEnvVarProject) GetComponentVersionOk() (*[]ComponentVersion, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ProjectEnvVarProject) SetComponentVersion(v []ComponentVersion)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ProjectEnvVarProject) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *ProjectEnvVarProject) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ProjectEnvVarProject) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ProjectEnvVarProject) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ProjectEnvVarProject) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ProjectEnvVarProject) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ProjectEnvVarProject) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *ProjectEnvVarProject) GetProjectEnvVar() []ProjectEnvVar`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *ProjectEnvVarProject) GetProjectEnvVarOk() (*[]ProjectEnvVar, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *ProjectEnvVarProject) SetProjectEnvVar(v []ProjectEnvVar)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *ProjectEnvVarProject) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectEnvVarProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectEnvVarProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectEnvVarProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectEnvVarProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectEnvVarProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectEnvVarProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectEnvVarProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectEnvVarProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


