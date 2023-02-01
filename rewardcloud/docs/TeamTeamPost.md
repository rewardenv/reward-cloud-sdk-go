# TeamTeamPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**User** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to **[]string** |  | [optional] 
**Organisation** | Pointer to **NullableString** |  | [optional] 
**TeamEnvVar** | Pointer to [**[]TeamEnvVarTeamPost**](TeamEnvVarTeamPost.md) |  | [optional] 

## Methods

### NewTeamTeamPost

`func NewTeamTeamPost() *TeamTeamPost`

NewTeamTeamPost instantiates a new TeamTeamPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamPostWithDefaults

`func NewTeamTeamPostWithDefaults() *TeamTeamPost`

NewTeamTeamPostWithDefaults instantiates a new TeamTeamPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamTeamPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamTeamPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamTeamPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamTeamPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TeamTeamPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TeamTeamPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *TeamTeamPost) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *TeamTeamPost) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *TeamTeamPost) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *TeamTeamPost) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *TeamTeamPost) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *TeamTeamPost) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *TeamTeamPost) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TeamTeamPost) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TeamTeamPost) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TeamTeamPost) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *TeamTeamPost) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *TeamTeamPost) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetUser

`func (o *TeamTeamPost) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamTeamPost) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamTeamPost) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *TeamTeamPost) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProject

`func (o *TeamTeamPost) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TeamTeamPost) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TeamTeamPost) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TeamTeamPost) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganisation

`func (o *TeamTeamPost) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *TeamTeamPost) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *TeamTeamPost) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *TeamTeamPost) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### SetOrganisationNil

`func (o *TeamTeamPost) SetOrganisationNil(b bool)`

 SetOrganisationNil sets the value for Organisation to be an explicit nil

### UnsetOrganisation
`func (o *TeamTeamPost) UnsetOrganisation()`

UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
### GetTeamEnvVar

`func (o *TeamTeamPost) GetTeamEnvVar() []TeamEnvVarTeamPost`

GetTeamEnvVar returns the TeamEnvVar field if non-nil, zero value otherwise.

### GetTeamEnvVarOk

`func (o *TeamTeamPost) GetTeamEnvVarOk() (*[]TeamEnvVarTeamPost, bool)`

GetTeamEnvVarOk returns a tuple with the TeamEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamEnvVar

`func (o *TeamTeamPost) SetTeamEnvVar(v []TeamEnvVarTeamPost)`

SetTeamEnvVar sets TeamEnvVar field to given value.

### HasTeamEnvVar

`func (o *TeamTeamPost) HasTeamEnvVar() bool`

HasTeamEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


