# EnvVarType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ProjectEnvVar** | Pointer to **[]string** |  | [optional] 
**EnvironmentEnvVar** | Pointer to **[]string** |  | [optional] 
**ComponentVersionEnvVar** | Pointer to **[]string** |  | [optional] 
**OrganisationEnvVar** | Pointer to **[]string** |  | [optional] 
**TeamEnvVar** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvVarType

`func NewEnvVarType() *EnvVarType`

NewEnvVarType instantiates a new EnvVarType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvVarTypeWithDefaults

`func NewEnvVarTypeWithDefaults() *EnvVarType`

NewEnvVarTypeWithDefaults instantiates a new EnvVarType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvVarType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvVarType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvVarType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvVarType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvVarType) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvVarType) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvVarType) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvVarType) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvVarType) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvVarType) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *EnvVarType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvVarType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvVarType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvVarType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvVarType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvVarType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *EnvVarType) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *EnvVarType) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *EnvVarType) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *EnvVarType) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *EnvVarType) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *EnvVarType) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *EnvVarType) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *EnvVarType) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *EnvVarType) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *EnvVarType) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *EnvVarType) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *EnvVarType) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetProjectEnvVar

`func (o *EnvVarType) GetProjectEnvVar() []string`

GetProjectEnvVar returns the ProjectEnvVar field if non-nil, zero value otherwise.

### GetProjectEnvVarOk

`func (o *EnvVarType) GetProjectEnvVarOk() (*[]string, bool)`

GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectEnvVar

`func (o *EnvVarType) SetProjectEnvVar(v []string)`

SetProjectEnvVar sets ProjectEnvVar field to given value.

### HasProjectEnvVar

`func (o *EnvVarType) HasProjectEnvVar() bool`

HasProjectEnvVar returns a boolean if a field has been set.

### GetEnvironmentEnvVar

`func (o *EnvVarType) GetEnvironmentEnvVar() []string`

GetEnvironmentEnvVar returns the EnvironmentEnvVar field if non-nil, zero value otherwise.

### GetEnvironmentEnvVarOk

`func (o *EnvVarType) GetEnvironmentEnvVarOk() (*[]string, bool)`

GetEnvironmentEnvVarOk returns a tuple with the EnvironmentEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentEnvVar

`func (o *EnvVarType) SetEnvironmentEnvVar(v []string)`

SetEnvironmentEnvVar sets EnvironmentEnvVar field to given value.

### HasEnvironmentEnvVar

`func (o *EnvVarType) HasEnvironmentEnvVar() bool`

HasEnvironmentEnvVar returns a boolean if a field has been set.

### GetComponentVersionEnvVar

`func (o *EnvVarType) GetComponentVersionEnvVar() []string`

GetComponentVersionEnvVar returns the ComponentVersionEnvVar field if non-nil, zero value otherwise.

### GetComponentVersionEnvVarOk

`func (o *EnvVarType) GetComponentVersionEnvVarOk() (*[]string, bool)`

GetComponentVersionEnvVarOk returns a tuple with the ComponentVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersionEnvVar

`func (o *EnvVarType) SetComponentVersionEnvVar(v []string)`

SetComponentVersionEnvVar sets ComponentVersionEnvVar field to given value.

### HasComponentVersionEnvVar

`func (o *EnvVarType) HasComponentVersionEnvVar() bool`

HasComponentVersionEnvVar returns a boolean if a field has been set.

### GetOrganisationEnvVar

`func (o *EnvVarType) GetOrganisationEnvVar() []string`

GetOrganisationEnvVar returns the OrganisationEnvVar field if non-nil, zero value otherwise.

### GetOrganisationEnvVarOk

`func (o *EnvVarType) GetOrganisationEnvVarOk() (*[]string, bool)`

GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationEnvVar

`func (o *EnvVarType) SetOrganisationEnvVar(v []string)`

SetOrganisationEnvVar sets OrganisationEnvVar field to given value.

### HasOrganisationEnvVar

`func (o *EnvVarType) HasOrganisationEnvVar() bool`

HasOrganisationEnvVar returns a boolean if a field has been set.

### GetTeamEnvVar

`func (o *EnvVarType) GetTeamEnvVar() []string`

GetTeamEnvVar returns the TeamEnvVar field if non-nil, zero value otherwise.

### GetTeamEnvVarOk

`func (o *EnvVarType) GetTeamEnvVarOk() (*[]string, bool)`

GetTeamEnvVarOk returns a tuple with the TeamEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamEnvVar

`func (o *EnvVarType) SetTeamEnvVar(v []string)`

SetTeamEnvVar sets TeamEnvVar field to given value.

### HasTeamEnvVar

`func (o *EnvVarType) HasTeamEnvVar() bool`

HasTeamEnvVar returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EnvVarType) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvVarType) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvVarType) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvVarType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvVarType) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvVarType) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvVarType) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvVarType) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvVarType) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvVarType) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvVarType) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvVarType) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvVarType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvVarType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvVarType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvVarType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvVarType) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvVarType) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvVarType) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvVarType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


