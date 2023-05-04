# AbstractProjectJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
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
**ComponentVersion** | Pointer to [**[]ComponentVersionJsonhal**](ComponentVersionJsonhal.md) |  | [optional] 
**ProjectTypeVersion** | Pointer to **NullableString** |  | [optional] 
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarJsonhal**](ProjectEnvVarJsonhal.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAbstractProjectJsonhal

`func NewAbstractProjectJsonhal() *AbstractProjectJsonhal`

NewAbstractProjectJsonhal instantiates a new AbstractProjectJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractProjectJsonhalWithDefaults

`func NewAbstractProjectJsonhalWithDefaults() *AbstractProjectJsonhal`

NewAbstractProjectJsonhalWithDefaults instantiates a new AbstractProjectJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AbstractProjectJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AbstractProjectJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AbstractProjectJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AbstractProjectJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AbstractProjectJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractProjectJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractProjectJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractProjectJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AbstractProjectJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AbstractProjectJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AbstractProjectJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AbstractProjectJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AbstractProjectJsonhal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractProjectJsonhal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractProjectJsonhal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractProjectJsonhal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AbstractProjectJsonhal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AbstractProjectJsonhal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *AbstractProjectJsonhal) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *AbstractProjectJsonhal) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *AbstractProjectJsonhal) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *AbstractProjectJsonhal) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *AbstractProjectJsonhal) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *AbstractProjectJsonhal) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsActive

`func (o *AbstractProjectJsonhal) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AbstractProjectJsonhal) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AbstractProjectJsonhal) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AbstractProjectJsonhal) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *AbstractProjectJsonhal) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *AbstractProjectJsonhal) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *AbstractProjectJsonhal) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AbstractProjectJsonhal) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AbstractProjectJsonhal) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AbstractProjectJsonhal) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *AbstractProjectJsonhal) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *AbstractProjectJsonhal) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *AbstractProjectJsonhal) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AbstractProjectJsonhal) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AbstractProjectJsonhal) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AbstractProjectJsonhal) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *AbstractProjectJsonhal) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *AbstractProjectJsonhal) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *AbstractProjectJsonhal) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AbstractProjectJsonhal) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AbstractProjectJsonhal) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *AbstractProjectJsonhal) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *AbstractProjectJsonhal) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *AbstractProjectJsonhal) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *AbstractProjectJsonhal) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AbstractProjectJsonhal) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AbstractProjectJsonhal) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AbstractProjectJsonhal) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AbstractProjectJsonhal) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AbstractProjectJsonhal) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *AbstractProjectJsonhal) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AbstractProjectJsonhal) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AbstractProjectJsonhal) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AbstractProjectJsonhal) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *AbstractProjectJsonhal) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *AbstractProjectJsonhal) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *AbstractProjectJsonhal) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *AbstractProjectJsonhal) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *AbstractProjectJsonhal) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *AbstractProjectJsonhal) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *AbstractProjectJsonhal) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *AbstractProjectJsonhal) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *AbstractProjectJsonhal) GetComponentVersion() []ComponentVersionJsonhal`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *AbstractProjectJsonhal) GetComponentVersionOk() (*[]ComponentVersionJsonhal, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *AbstractProjectJsonhal) SetComponentVersion(v []ComponentVersionJsonhal)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *AbstractProjectJsonhal) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *AbstractProjectJsonhal) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *AbstractProjectJsonhal) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *AbstractProjectJsonhal) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *AbstractProjectJsonhal) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *AbstractProjectJsonhal) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *AbstractProjectJsonhal) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *AbstractProjectJsonhal) GetProjectEnvVar() []ProjectEnvVarJsonhal`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *AbstractProjectJsonhal) GetProjectEnvVarOk() (*[]ProjectEnvVarJsonhal, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *AbstractProjectJsonhal) SetProjectEnvVar(v []ProjectEnvVarJsonhal)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *AbstractProjectJsonhal) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AbstractProjectJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AbstractProjectJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AbstractProjectJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AbstractProjectJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AbstractProjectJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AbstractProjectJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AbstractProjectJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AbstractProjectJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


