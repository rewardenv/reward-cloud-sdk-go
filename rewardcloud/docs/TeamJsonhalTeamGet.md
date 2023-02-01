# TeamJsonhalTeamGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**User** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to **[]string** |  | [optional] 
**Organisation** | Pointer to **NullableString** |  | [optional] 
**TeamEnvVar** | Pointer to [**[]TeamEnvVarJsonhalTeamGet**](TeamEnvVarJsonhalTeamGet.md) |  | [optional] 

## Methods

### NewTeamJsonhalTeamGet

`func NewTeamJsonhalTeamGet() *TeamJsonhalTeamGet`

NewTeamJsonhalTeamGet instantiates a new TeamJsonhalTeamGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamJsonhalTeamGetWithDefaults

`func NewTeamJsonhalTeamGetWithDefaults() *TeamJsonhalTeamGet`

NewTeamJsonhalTeamGetWithDefaults instantiates a new TeamJsonhalTeamGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TeamJsonhalTeamGet) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamJsonhalTeamGet) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamJsonhalTeamGet) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamJsonhalTeamGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *TeamJsonhalTeamGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamJsonhalTeamGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamJsonhalTeamGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamJsonhalTeamGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TeamJsonhalTeamGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TeamJsonhalTeamGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TeamJsonhalTeamGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TeamJsonhalTeamGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *TeamJsonhalTeamGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamJsonhalTeamGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamJsonhalTeamGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamJsonhalTeamGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TeamJsonhalTeamGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TeamJsonhalTeamGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *TeamJsonhalTeamGet) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *TeamJsonhalTeamGet) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *TeamJsonhalTeamGet) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *TeamJsonhalTeamGet) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *TeamJsonhalTeamGet) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *TeamJsonhalTeamGet) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *TeamJsonhalTeamGet) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *TeamJsonhalTeamGet) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *TeamJsonhalTeamGet) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *TeamJsonhalTeamGet) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *TeamJsonhalTeamGet) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *TeamJsonhalTeamGet) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetUser

`func (o *TeamJsonhalTeamGet) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TeamJsonhalTeamGet) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TeamJsonhalTeamGet) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *TeamJsonhalTeamGet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetProject

`func (o *TeamJsonhalTeamGet) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TeamJsonhalTeamGet) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TeamJsonhalTeamGet) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *TeamJsonhalTeamGet) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOrganisation

`func (o *TeamJsonhalTeamGet) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *TeamJsonhalTeamGet) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *TeamJsonhalTeamGet) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *TeamJsonhalTeamGet) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### SetOrganisationNil

`func (o *TeamJsonhalTeamGet) SetOrganisationNil(b bool)`

 SetOrganisationNil sets the value for Organisation to be an explicit nil

### UnsetOrganisation
`func (o *TeamJsonhalTeamGet) UnsetOrganisation()`

UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
### GetTeamEnvVar

`func (o *TeamJsonhalTeamGet) GetTeamEnvVar() []TeamEnvVarJsonhalTeamGet`

GetTeamEnvVar returns the TeamEnvVar field if non-nil, zero value otherwise.

### GetTeamEnvVarOk

`func (o *TeamJsonhalTeamGet) GetTeamEnvVarOk() (*[]TeamEnvVarJsonhalTeamGet, bool)`

GetTeamEnvVarOk returns a tuple with the TeamEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamEnvVar

`func (o *TeamJsonhalTeamGet) SetTeamEnvVar(v []TeamEnvVarJsonhalTeamGet)`

SetTeamEnvVar sets TeamEnvVar field to given value.

### HasTeamEnvVar

`func (o *TeamJsonhalTeamGet) HasTeamEnvVar() bool`

HasTeamEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


