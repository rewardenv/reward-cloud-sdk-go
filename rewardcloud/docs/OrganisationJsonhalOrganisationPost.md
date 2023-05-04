# OrganisationJsonhalOrganisationPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**Team** | Pointer to **[]string** |  | [optional] 
**OrganisationEnvVar** | Pointer to [**[]OrganisationEnvVarJsonhalOrganisationPost**](OrganisationEnvVarJsonhalOrganisationPost.md) |  | [optional] 

## Methods

### NewOrganisationJsonhalOrganisationPost

`func NewOrganisationJsonhalOrganisationPost() *OrganisationJsonhalOrganisationPost`

NewOrganisationJsonhalOrganisationPost instantiates a new OrganisationJsonhalOrganisationPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationJsonhalOrganisationPostWithDefaults

`func NewOrganisationJsonhalOrganisationPostWithDefaults() *OrganisationJsonhalOrganisationPost`

NewOrganisationJsonhalOrganisationPostWithDefaults instantiates a new OrganisationJsonhalOrganisationPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *OrganisationJsonhalOrganisationPost) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganisationJsonhalOrganisationPost) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganisationJsonhalOrganisationPost) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganisationJsonhalOrganisationPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *OrganisationJsonhalOrganisationPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganisationJsonhalOrganisationPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganisationJsonhalOrganisationPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganisationJsonhalOrganisationPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganisationJsonhalOrganisationPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganisationJsonhalOrganisationPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *OrganisationJsonhalOrganisationPost) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *OrganisationJsonhalOrganisationPost) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *OrganisationJsonhalOrganisationPost) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *OrganisationJsonhalOrganisationPost) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *OrganisationJsonhalOrganisationPost) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *OrganisationJsonhalOrganisationPost) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *OrganisationJsonhalOrganisationPost) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OrganisationJsonhalOrganisationPost) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OrganisationJsonhalOrganisationPost) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OrganisationJsonhalOrganisationPost) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *OrganisationJsonhalOrganisationPost) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *OrganisationJsonhalOrganisationPost) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetTeam

`func (o *OrganisationJsonhalOrganisationPost) GetTeam() []string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganisationJsonhalOrganisationPost) GetTeamOk() (*[]string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganisationJsonhalOrganisationPost) SetTeam(v []string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *OrganisationJsonhalOrganisationPost) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOrganisationEnvVar

`func (o *OrganisationJsonhalOrganisationPost) GetOrganisationEnvVar() []OrganisationEnvVarJsonhalOrganisationPost`

GetOrganisationEnvVar returns the OrganisationEnvVar field if non-nil, zero value otherwise.

### GetOrganisationEnvVarOk

`func (o *OrganisationJsonhalOrganisationPost) GetOrganisationEnvVarOk() (*[]OrganisationEnvVarJsonhalOrganisationPost, bool)`

GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationEnvVar

`func (o *OrganisationJsonhalOrganisationPost) SetOrganisationEnvVar(v []OrganisationEnvVarJsonhalOrganisationPost)`

SetOrganisationEnvVar sets OrganisationEnvVar field to given value.

### HasOrganisationEnvVar

`func (o *OrganisationJsonhalOrganisationPost) HasOrganisationEnvVar() bool`

HasOrganisationEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


