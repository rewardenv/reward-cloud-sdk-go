# ProjectProjectGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateIri** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Git** | Pointer to [**NullableProjectProjectGetGit**](ProjectProjectGetGit.md) |  | [optional] 
**ServiceAccount** | Pointer to [**NullableProjectProjectGetServiceAccount**](ProjectProjectGetServiceAccount.md) |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
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
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarProjectGet**](ProjectEnvVarProjectGet.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectProjectGet

`func NewProjectProjectGet() *ProjectProjectGet`

NewProjectProjectGet instantiates a new ProjectProjectGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectGetWithDefaults

`func NewProjectProjectGetWithDefaults() *ProjectProjectGet`

NewProjectProjectGetWithDefaults instantiates a new ProjectProjectGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateIri

`func (o *ProjectProjectGet) GetTemplateIri() string`

GetTemplateIri returns the TemplateIri field if non-nil, zero value otherwise.

### GetTemplateIriOk

`func (o *ProjectProjectGet) GetTemplateIriOk() (*string, bool)`

GetTemplateIriOk returns a tuple with the TemplateIri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIri

`func (o *ProjectProjectGet) SetTemplateIri(v string)`

SetTemplateIri sets TemplateIri field to given value.

### HasTemplateIri

`func (o *ProjectProjectGet) HasTemplateIri() bool`

HasTemplateIri returns a boolean if a field has been set.

### SetTemplateIriNil

`func (o *ProjectProjectGet) SetTemplateIriNil(b bool)`

 SetTemplateIriNil sets the value for TemplateIri to be an explicit nil

### UnsetTemplateIri
`func (o *ProjectProjectGet) UnsetTemplateIri()`

UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
### GetState

`func (o *ProjectProjectGet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectProjectGet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectProjectGet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProjectProjectGet) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ProjectProjectGet) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ProjectProjectGet) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetGit

`func (o *ProjectProjectGet) GetGit() ProjectProjectGetGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *ProjectProjectGet) GetGitOk() (*ProjectProjectGetGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *ProjectProjectGet) SetGit(v ProjectProjectGetGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *ProjectProjectGet) HasGit() bool`

HasGit returns a boolean if a field has been set.

### SetGitNil

`func (o *ProjectProjectGet) SetGitNil(b bool)`

 SetGitNil sets the value for Git to be an explicit nil

### UnsetGit
`func (o *ProjectProjectGet) UnsetGit()`

UnsetGit ensures that no value is present for Git, not even an explicit nil
### GetServiceAccount

`func (o *ProjectProjectGet) GetServiceAccount() ProjectProjectGetServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ProjectProjectGet) GetServiceAccountOk() (*ProjectProjectGetServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ProjectProjectGet) SetServiceAccount(v ProjectProjectGetServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *ProjectProjectGet) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### SetServiceAccountNil

`func (o *ProjectProjectGet) SetServiceAccountNil(b bool)`

 SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil

### UnsetServiceAccount
`func (o *ProjectProjectGet) UnsetServiceAccount()`

UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
### GetTeam

`func (o *ProjectProjectGet) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ProjectProjectGet) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ProjectProjectGet) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ProjectProjectGet) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ProjectProjectGet) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ProjectProjectGet) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetEnvironment

`func (o *ProjectProjectGet) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProjectProjectGet) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProjectProjectGet) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProjectProjectGet) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *ProjectProjectGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectProjectGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectProjectGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectProjectGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectProjectGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectProjectGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectProjectGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectProjectGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ProjectProjectGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectProjectGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectProjectGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectProjectGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectProjectGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectProjectGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *ProjectProjectGet) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectProjectGet) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectProjectGet) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProjectProjectGet) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProjectProjectGet) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProjectProjectGet) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *ProjectProjectGet) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProjectProjectGet) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProjectProjectGet) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProjectProjectGet) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ProjectProjectGet) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ProjectProjectGet) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ProjectProjectGet) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectProjectGet) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectProjectGet) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ProjectProjectGet) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ProjectProjectGet) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ProjectProjectGet) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *ProjectProjectGet) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectProjectGet) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectProjectGet) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProjectProjectGet) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ProjectProjectGet) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ProjectProjectGet) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *ProjectProjectGet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectProjectGet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectProjectGet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectProjectGet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ProjectProjectGet) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ProjectProjectGet) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *ProjectProjectGet) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ProjectProjectGet) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ProjectProjectGet) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ProjectProjectGet) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ProjectProjectGet) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ProjectProjectGet) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *ProjectProjectGet) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *ProjectProjectGet) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *ProjectProjectGet) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *ProjectProjectGet) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *ProjectProjectGet) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *ProjectProjectGet) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *ProjectProjectGet) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ProjectProjectGet) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ProjectProjectGet) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ProjectProjectGet) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *ProjectProjectGet) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ProjectProjectGet) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ProjectProjectGet) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ProjectProjectGet) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ProjectProjectGet) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ProjectProjectGet) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *ProjectProjectGet) GetProjectEnvVar() []ProjectEnvVarProjectGet`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *ProjectProjectGet) GetProjectEnvVarOk() (*[]ProjectEnvVarProjectGet, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *ProjectProjectGet) SetProjectEnvVar(v []ProjectEnvVarProjectGet)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *ProjectProjectGet) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectProjectGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectProjectGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectProjectGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectProjectGet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectProjectGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectProjectGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectProjectGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectProjectGet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


