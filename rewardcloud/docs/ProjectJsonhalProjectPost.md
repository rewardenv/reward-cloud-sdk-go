# ProjectJsonhalProjectPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**IsInitProjectSkeleton** | Pointer to **NullableBool** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 
**Git** | Pointer to [**NullableProjectJsonhalProjectPostGit**](ProjectJsonhalProjectPostGit.md) |  | [optional] 
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarJsonhalProjectPost**](ProjectEnvVarJsonhalProjectPost.md) |  | [optional] 
**ProjectTypeVersion** | Pointer to **NullableString** |  | [optional] 
**ComponentVersion** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectJsonhalProjectPost

`func NewProjectJsonhalProjectPost() *ProjectJsonhalProjectPost`

NewProjectJsonhalProjectPost instantiates a new ProjectJsonhalProjectPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectJsonhalProjectPostWithDefaults

`func NewProjectJsonhalProjectPostWithDefaults() *ProjectJsonhalProjectPost`

NewProjectJsonhalProjectPostWithDefaults instantiates a new ProjectJsonhalProjectPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectJsonhalProjectPost) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectJsonhalProjectPost) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectJsonhalProjectPost) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectJsonhalProjectPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *ProjectJsonhalProjectPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectJsonhalProjectPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectJsonhalProjectPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectJsonhalProjectPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectJsonhalProjectPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectJsonhalProjectPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *ProjectJsonhalProjectPost) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectJsonhalProjectPost) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectJsonhalProjectPost) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProjectJsonhalProjectPost) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCpu

`func (o *ProjectJsonhalProjectPost) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProjectJsonhalProjectPost) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProjectJsonhalProjectPost) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProjectJsonhalProjectPost) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ProjectJsonhalProjectPost) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ProjectJsonhalProjectPost) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ProjectJsonhalProjectPost) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectJsonhalProjectPost) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectJsonhalProjectPost) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ProjectJsonhalProjectPost) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ProjectJsonhalProjectPost) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ProjectJsonhalProjectPost) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *ProjectJsonhalProjectPost) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectJsonhalProjectPost) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectJsonhalProjectPost) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProjectJsonhalProjectPost) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ProjectJsonhalProjectPost) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ProjectJsonhalProjectPost) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *ProjectJsonhalProjectPost) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectJsonhalProjectPost) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectJsonhalProjectPost) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectJsonhalProjectPost) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ProjectJsonhalProjectPost) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ProjectJsonhalProjectPost) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *ProjectJsonhalProjectPost) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ProjectJsonhalProjectPost) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ProjectJsonhalProjectPost) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ProjectJsonhalProjectPost) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ProjectJsonhalProjectPost) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ProjectJsonhalProjectPost) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *ProjectJsonhalProjectPost) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *ProjectJsonhalProjectPost) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *ProjectJsonhalProjectPost) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *ProjectJsonhalProjectPost) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *ProjectJsonhalProjectPost) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *ProjectJsonhalProjectPost) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetEnvironment

`func (o *ProjectJsonhalProjectPost) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProjectJsonhalProjectPost) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProjectJsonhalProjectPost) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProjectJsonhalProjectPost) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTeam

`func (o *ProjectJsonhalProjectPost) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ProjectJsonhalProjectPost) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ProjectJsonhalProjectPost) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ProjectJsonhalProjectPost) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ProjectJsonhalProjectPost) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ProjectJsonhalProjectPost) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetGit

`func (o *ProjectJsonhalProjectPost) GetGit() ProjectJsonhalProjectPostGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *ProjectJsonhalProjectPost) GetGitOk() (*ProjectJsonhalProjectPostGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *ProjectJsonhalProjectPost) SetGit(v ProjectJsonhalProjectPostGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *ProjectJsonhalProjectPost) HasGit() bool`

HasGit returns a boolean if a field has been set.

### SetGitNil

`func (o *ProjectJsonhalProjectPost) SetGitNil(b bool)`

 SetGitNil sets the value for Git to be an explicit nil

### UnsetGit
`func (o *ProjectJsonhalProjectPost) UnsetGit()`

UnsetGit ensures that no value is present for Git, not even an explicit nil
### GetProjectEnvVar

`func (o *ProjectJsonhalProjectPost) GetProjectEnvVar() []ProjectEnvVarJsonhalProjectPost`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *ProjectJsonhalProjectPost) GetProjectEnvVarOk() (*[]ProjectEnvVarJsonhalProjectPost, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *ProjectJsonhalProjectPost) SetProjectEnvVar(v []ProjectEnvVarJsonhalProjectPost)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *ProjectJsonhalProjectPost) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *ProjectJsonhalProjectPost) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ProjectJsonhalProjectPost) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ProjectJsonhalProjectPost) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ProjectJsonhalProjectPost) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ProjectJsonhalProjectPost) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ProjectJsonhalProjectPost) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetComponentVersion

`func (o *ProjectJsonhalProjectPost) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ProjectJsonhalProjectPost) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ProjectJsonhalProjectPost) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ProjectJsonhalProjectPost) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


