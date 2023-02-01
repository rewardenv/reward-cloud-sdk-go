# TeamJsonhalTeamPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**User** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to **[]string** |  | [optional] 
**Organisation** | Pointer to **NullableString** |  | [optional] 
**TeamEnvVar** | Pointer to [**[]TeamEnvVarJsonhalTeamPost**](TeamEnvVarJsonhalTeamPost.md) |  | [optional] 

## Methods

### NewTeamJsonhalTeamPost

`func NewTeamJsonhalTeamPost() *TeamJsonhalTeamPost`

NewTeamJsonhalTeamPost instantiates a new TeamJsonhalTeamPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamJsonhalTeamPostWithDefaults

`func NewTeamJsonhalTeamPostWithDefaults() *TeamJsonhalTeamPost`

NewTeamJsonhalTeamPostWithDefaults instantiates a new TeamJsonhalTeamPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TeamJsonhalTeamPost) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamJsonhalTeamPost) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamJsonhalTeamPost) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamJsonhalTeamPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *TeamJsonhalTeamPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamJsonhalTeamPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamJsonhalTeamPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamJsonhalTeamPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TeamJsonhalTeamPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TeamJsonhalTeamPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *TeamJsonhalTeamPost) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *TeamJsonhalTeamPost) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *TeamJsonhalTeamPost) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *TeamJsonhalTeamPost) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *TeamJsonhalTeamPost) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *TeamJsonhalTeamPost) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *TeamJsonhalTeamPost) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TeamJsonhalTeamPost) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TeamJsonhalTeamPost) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TeamJsonhalTeamPost) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *TeamJsonhalTeamPost) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *TeamJsonhalTeamPost) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetUser

`func (o *TeamJsonhalTeamPost) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamJsonhalTeamPost) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamJsonhalTeamPost) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *TeamJsonhalTeamPost) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProject

`func (o *TeamJsonhalTeamPost) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TeamJsonhalTeamPost) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TeamJsonhalTeamPost) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TeamJsonhalTeamPost) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganisation

`func (o *TeamJsonhalTeamPost) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *TeamJsonhalTeamPost) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *TeamJsonhalTeamPost) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *TeamJsonhalTeamPost) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### SetOrganisationNil

`func (o *TeamJsonhalTeamPost) SetOrganisationNil(b bool)`

 SetOrganisationNil sets the value for Organisation to be an explicit nil

### UnsetOrganisation
`func (o *TeamJsonhalTeamPost) UnsetOrganisation()`

UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
### GetTeamEnvVar

`func (o *TeamJsonhalTeamPost) GetTeamEnvVar() []TeamEnvVarJsonhalTeamPost`

GetTeamEnvVar returns the TeamEnvVar field if non-nil, zero value otherwise.

### GetTeamEnvVarOk

`func (o *TeamJsonhalTeamPost) GetTeamEnvVarOk() (*[]TeamEnvVarJsonhalTeamPost, bool)`

GetTeamEnvVarOk returns a tuple with the TeamEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamEnvVar

`func (o *TeamJsonhalTeamPost) SetTeamEnvVar(v []TeamEnvVarJsonhalTeamPost)`

SetTeamEnvVar sets TeamEnvVar field to given value.

### HasTeamEnvVar

`func (o *TeamJsonhalTeamPost) HasTeamEnvVar() bool`

HasTeamEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


