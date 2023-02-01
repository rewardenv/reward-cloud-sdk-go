# TeamTeamGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**User** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to **[]string** |  | [optional] 
**Organisation** | Pointer to **NullableString** |  | [optional] 
**TeamEnvVar** | Pointer to [**[]TeamEnvVarTeamGet**](TeamEnvVarTeamGet.md) |  | [optional] 

## Methods

### NewTeamTeamGet

`func NewTeamTeamGet() *TeamTeamGet`

NewTeamTeamGet instantiates a new TeamTeamGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamGetWithDefaults

`func NewTeamTeamGetWithDefaults() *TeamTeamGet`

NewTeamTeamGetWithDefaults instantiates a new TeamTeamGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamTeamGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TeamTeamGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TeamTeamGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TeamTeamGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TeamTeamGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *TeamTeamGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TeamTeamGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TeamTeamGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *TeamTeamGet) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *TeamTeamGet) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *TeamTeamGet) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *TeamTeamGet) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *TeamTeamGet) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *TeamTeamGet) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *TeamTeamGet) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TeamTeamGet) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TeamTeamGet) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TeamTeamGet) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *TeamTeamGet) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *TeamTeamGet) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetUser

`func (o *TeamTeamGet) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamTeamGet) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamTeamGet) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *TeamTeamGet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProject

`func (o *TeamTeamGet) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TeamTeamGet) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TeamTeamGet) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TeamTeamGet) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganisation

`func (o *TeamTeamGet) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *TeamTeamGet) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *TeamTeamGet) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *TeamTeamGet) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### SetOrganisationNil

`func (o *TeamTeamGet) SetOrganisationNil(b bool)`

 SetOrganisationNil sets the value for Organisation to be an explicit nil

### UnsetOrganisation
`func (o *TeamTeamGet) UnsetOrganisation()`

UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
### GetTeamEnvVar

`func (o *TeamTeamGet) GetTeamEnvVar() []TeamEnvVarTeamGet`

GetTeamEnvVar returns the TeamEnvVar field if non-nil, zero value otherwise.

### GetTeamEnvVarOk

`func (o *TeamTeamGet) GetTeamEnvVarOk() (*[]TeamEnvVarTeamGet, bool)`

GetTeamEnvVarOk returns a tuple with the TeamEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamEnvVar

`func (o *TeamTeamGet) SetTeamEnvVar(v []TeamEnvVarTeamGet)`

SetTeamEnvVar sets TeamEnvVar field to given value.

### HasTeamEnvVar

`func (o *TeamTeamGet) HasTeamEnvVar() bool`

HasTeamEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


