# ProjectJsonhalProjectGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**TemplateIri** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Git** | Pointer to [**NullableProjectJsonhalProjectGetGit**](ProjectJsonhalProjectGetGit.md) |  | [optional] 
**ServiceAccount** | Pointer to [**NullableProjectJsonhalProjectGetServiceAccount**](ProjectJsonhalProjectGetServiceAccount.md) |  | [optional] 
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
**ProjectEnvVar** | Pointer to [**[]ProjectEnvVarJsonhalProjectGet**](ProjectEnvVarJsonhalProjectGet.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectJsonhalProjectGet

`func NewProjectJsonhalProjectGet() *ProjectJsonhalProjectGet`

NewProjectJsonhalProjectGet instantiates a new ProjectJsonhalProjectGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectJsonhalProjectGetWithDefaults

`func NewProjectJsonhalProjectGetWithDefaults() *ProjectJsonhalProjectGet`

NewProjectJsonhalProjectGetWithDefaults instantiates a new ProjectJsonhalProjectGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectJsonhalProjectGet) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectJsonhalProjectGet) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectJsonhalProjectGet) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectJsonhalProjectGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTemplateIri

`func (o *ProjectJsonhalProjectGet) GetTemplateIri() string`

GetTemplateIri returns the TemplateIri field if non-nil, zero value otherwise.

### GetTemplateIriOk

`func (o *ProjectJsonhalProjectGet) GetTemplateIriOk() (*string, bool)`

GetTemplateIriOk returns a tuple with the TemplateIri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIri

`func (o *ProjectJsonhalProjectGet) SetTemplateIri(v string)`

SetTemplateIri sets TemplateIri field to given value.

### HasTemplateIri

`func (o *ProjectJsonhalProjectGet) HasTemplateIri() bool`

HasTemplateIri returns a boolean if a field has been set.

### SetTemplateIriNil

`func (o *ProjectJsonhalProjectGet) SetTemplateIriNil(b bool)`

 SetTemplateIriNil sets the value for TemplateIri to be an explicit nil

### UnsetTemplateIri
`func (o *ProjectJsonhalProjectGet) UnsetTemplateIri()`

UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
### GetState

`func (o *ProjectJsonhalProjectGet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectJsonhalProjectGet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectJsonhalProjectGet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProjectJsonhalProjectGet) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ProjectJsonhalProjectGet) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ProjectJsonhalProjectGet) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetGit

`func (o *ProjectJsonhalProjectGet) GetGit() ProjectJsonhalProjectGetGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *ProjectJsonhalProjectGet) GetGitOk() (*ProjectJsonhalProjectGetGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *ProjectJsonhalProjectGet) SetGit(v ProjectJsonhalProjectGetGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *ProjectJsonhalProjectGet) HasGit() bool`

HasGit returns a boolean if a field has been set.

### SetGitNil

`func (o *ProjectJsonhalProjectGet) SetGitNil(b bool)`

 SetGitNil sets the value for Git to be an explicit nil

### UnsetGit
`func (o *ProjectJsonhalProjectGet) UnsetGit()`

UnsetGit ensures that no value is present for Git, not even an explicit nil
### GetServiceAccount

`func (o *ProjectJsonhalProjectGet) GetServiceAccount() ProjectJsonhalProjectGetServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ProjectJsonhalProjectGet) GetServiceAccountOk() (*ProjectJsonhalProjectGetServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ProjectJsonhalProjectGet) SetServiceAccount(v ProjectJsonhalProjectGetServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *ProjectJsonhalProjectGet) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### SetServiceAccountNil

`func (o *ProjectJsonhalProjectGet) SetServiceAccountNil(b bool)`

 SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil

### UnsetServiceAccount
`func (o *ProjectJsonhalProjectGet) UnsetServiceAccount()`

UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
### GetTeam

`func (o *ProjectJsonhalProjectGet) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ProjectJsonhalProjectGet) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ProjectJsonhalProjectGet) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ProjectJsonhalProjectGet) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ProjectJsonhalProjectGet) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ProjectJsonhalProjectGet) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetEnvironment

`func (o *ProjectJsonhalProjectGet) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProjectJsonhalProjectGet) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProjectJsonhalProjectGet) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProjectJsonhalProjectGet) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *ProjectJsonhalProjectGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectJsonhalProjectGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectJsonhalProjectGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectJsonhalProjectGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectJsonhalProjectGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectJsonhalProjectGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectJsonhalProjectGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectJsonhalProjectGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ProjectJsonhalProjectGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectJsonhalProjectGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectJsonhalProjectGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectJsonhalProjectGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProjectJsonhalProjectGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectJsonhalProjectGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsActive

`func (o *ProjectJsonhalProjectGet) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProjectJsonhalProjectGet) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProjectJsonhalProjectGet) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProjectJsonhalProjectGet) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProjectJsonhalProjectGet) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProjectJsonhalProjectGet) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCpu

`func (o *ProjectJsonhalProjectGet) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ProjectJsonhalProjectGet) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ProjectJsonhalProjectGet) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ProjectJsonhalProjectGet) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *ProjectJsonhalProjectGet) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *ProjectJsonhalProjectGet) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *ProjectJsonhalProjectGet) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ProjectJsonhalProjectGet) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ProjectJsonhalProjectGet) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ProjectJsonhalProjectGet) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *ProjectJsonhalProjectGet) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *ProjectJsonhalProjectGet) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *ProjectJsonhalProjectGet) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ProjectJsonhalProjectGet) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ProjectJsonhalProjectGet) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ProjectJsonhalProjectGet) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *ProjectJsonhalProjectGet) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *ProjectJsonhalProjectGet) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetCode

`func (o *ProjectJsonhalProjectGet) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectJsonhalProjectGet) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectJsonhalProjectGet) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectJsonhalProjectGet) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ProjectJsonhalProjectGet) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ProjectJsonhalProjectGet) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetColor

`func (o *ProjectJsonhalProjectGet) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ProjectJsonhalProjectGet) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ProjectJsonhalProjectGet) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ProjectJsonhalProjectGet) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ProjectJsonhalProjectGet) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ProjectJsonhalProjectGet) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetIsInitProjectSkeleton

`func (o *ProjectJsonhalProjectGet) GetIsInitProjectSkeleton() bool`

GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field if non-nil, zero value otherwise.

### GetIsInitProjectSkeletonOk

`func (o *ProjectJsonhalProjectGet) GetIsInitProjectSkeletonOk() (*bool, bool)`

GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitProjectSkeleton

`func (o *ProjectJsonhalProjectGet) SetIsInitProjectSkeleton(v bool)`

SetIsInitProjectSkeleton sets IsInitProjectSkeleton field to given value.

### HasIsInitProjectSkeleton

`func (o *ProjectJsonhalProjectGet) HasIsInitProjectSkeleton() bool`

HasIsInitProjectSkeleton returns a boolean if a field has been set.

### SetIsInitProjectSkeletonNil

`func (o *ProjectJsonhalProjectGet) SetIsInitProjectSkeletonNil(b bool)`

 SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil

### UnsetIsInitProjectSkeleton
`func (o *ProjectJsonhalProjectGet) UnsetIsInitProjectSkeleton()`

UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
### GetComponentVersion

`func (o *ProjectJsonhalProjectGet) GetComponentVersion() []string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ProjectJsonhalProjectGet) GetComponentVersionOk() (*[]string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ProjectJsonhalProjectGet) SetComponentVersion(v []string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ProjectJsonhalProjectGet) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### GetProjectTypeVersion

`func (o *ProjectJsonhalProjectGet) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ProjectJsonhalProjectGet) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ProjectJsonhalProjectGet) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ProjectJsonhalProjectGet) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ProjectJsonhalProjectGet) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ProjectJsonhalProjectGet) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetProjectEnvVar

`func (o *ProjectJsonhalProjectGet) GetProjectEnvVar() []ProjectEnvVarJsonhalProjectGet`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *ProjectJsonhalProjectGet) GetProjectEnvVarOk() (*[]ProjectEnvVarJsonhalProjectGet, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *ProjectJsonhalProjectGet) SetProjectEnvVar(v []ProjectEnvVarJsonhalProjectGet)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *ProjectJsonhalProjectGet) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProjectJsonhalProjectGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectJsonhalProjectGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectJsonhalProjectGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectJsonhalProjectGet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectJsonhalProjectGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectJsonhalProjectGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectJsonhalProjectGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectJsonhalProjectGet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


