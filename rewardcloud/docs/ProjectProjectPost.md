# ProjectProjectPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateIri** | Pointer to **NullableString** |  | [optional] 
**Git** | Pointer to [**NullableProjectProjectPostGit**](ProjectProjectPostGit.md) |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
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
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarProjectPost**](ProjectEnvVarProjectPost.md) |  | [optional] 

## Methods

### NewProjectProjectPost

`func NewProjectProjectPost() *ProjectProjectPost`

NewProjectProjectPost instantiates a new ProjectProjectPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectProjectPostWithDefaults

`func NewProjectProjectPostWithDefaults() *ProjectProjectPost`

NewProjectProjectPostWithDefaults instantiates a new ProjectProjectPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateIri

`func (o *ProjectProjectPost) GetTemplateIri() string`

GetTemplateIri returns the TemplateIri field if non-nil, zero value otherwise.

### GetTemplateIriOk

`func (o *ProjectProjectPost) GetTemplateIriOk() (*string, bool)`

GetTemplateIriOk returns a tuple with the TemplateIri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIri

`func (o *ProjectProjectPost) SetTemplateIri(v string)`

SetTemplateIri sets TemplateIri field to given value.

### HasTemplateIri

`func (o *ProjectProjectPost) HasTemplateIri() bool`

HasTemplateIri returns a boolean if a field has been set.

### SetTemplateIriNil

`func (o *ProjectProjectPost) SetTemplateIriNil(b bool)`

 SetTemplateIriNil sets the value for TemplateIri to be an explicit nil

### UnsetTemplateIri
`func (o *ProjectProjectPost) UnsetTemplateIri()`

UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
### GetGit

`func (o *ProjectProjectPost) GetGit() ProjectProjectPostGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *ProjectProjectPost) GetGitOk() (*ProjectProjectPostGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *ProjectProjectPost) SetGit(v ProjectProjectPostGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *ProjectProjectPost) HasGit() bool`

HasGit returns a boolean if a field has been set.

### SetGitNil

`func (o *ProjectProjectPost) SetGitNil(b bool)`

 SetGitNil sets the value for Git to be an explicit nil

### UnsetGit
`func (o *ProjectProjectPost) UnsetGit()`

UnsetGit ensures that no value is present for Git, not even an explicit nil
### GetTeam

`func (o *ProjectProjectPost) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ProjectProjectPost) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ProjectProjectPost) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ProjectProjectPost) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ProjectProjectPost) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ProjectProjectPost) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetEnvironment

`func (o *ProjectProjectPost) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProjectProjectPost) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProjectProjectPost) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProjectProjectPost) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *ProjectProjectPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectProjectPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectProjectPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectProjectPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectProjectPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectProjectPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *ProjectProjectPost) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectProjectPost) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectProjectPost) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProjectProjectPost) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProjectProjectPost) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProjectProjectPost) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *ProjectProjectPost) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProjectProjectPost) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProjectProjectPost) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProjectProjectPost) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ProjectProjectPost) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ProjectProjectPost) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ProjectProjectPost) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectProjectPost) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectProjectPost) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ProjectProjectPost) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ProjectProjectPost) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ProjectProjectPost) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *ProjectProjectPost) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectProjectPost) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectProjectPost) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProjectProjectPost) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ProjectProjectPost) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ProjectProjectPost) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *ProjectProjectPost) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectProjectPost) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectProjectPost) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectProjectPost) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ProjectProjectPost) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ProjectProjectPost) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *ProjectProjectPost) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ProjectProjectPost) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ProjectProjectPost) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ProjectProjectPost) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ProjectProjectPost) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ProjectProjectPost) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *ProjectProjectPost) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *ProjectProjectPost) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *ProjectProjectPost) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *ProjectProjectPost) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *ProjectProjectPost) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *ProjectProjectPost) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *ProjectProjectPost) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ProjectProjectPost) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ProjectProjectPost) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ProjectProjectPost) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *ProjectProjectPost) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ProjectProjectPost) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ProjectProjectPost) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ProjectProjectPost) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ProjectProjectPost) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ProjectProjectPost) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *ProjectProjectPost) GetProjectEnvVar() []ProjectEnvVarProjectPost`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *ProjectProjectPost) GetProjectEnvVarOk() (*[]ProjectEnvVarProjectPost, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *ProjectProjectPost) SetProjectEnvVar(v []ProjectEnvVarProjectPost)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *ProjectProjectPost) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


