# OrganisationJsonhalOrganisationGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**Team** | Pointer to **[]string** |  | [optional] 
**OrganisationEnvVar** | Pointer to [**[]OrganisationEnvVarJsonhalOrganisationGet**](OrganisationEnvVarJsonhalOrganisationGet.md) |  | [optional] 

## Methods

### NewOrganisationJsonhalOrganisationGet

`func NewOrganisationJsonhalOrganisationGet() *OrganisationJsonhalOrganisationGet`

NewOrganisationJsonhalOrganisationGet instantiates a new OrganisationJsonhalOrganisationGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationJsonhalOrganisationGetWithDefaults

`func NewOrganisationJsonhalOrganisationGetWithDefaults() *OrganisationJsonhalOrganisationGet`

NewOrganisationJsonhalOrganisationGetWithDefaults instantiates a new OrganisationJsonhalOrganisationGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *OrganisationJsonhalOrganisationGet) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganisationJsonhalOrganisationGet) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganisationJsonhalOrganisationGet) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganisationJsonhalOrganisationGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *OrganisationJsonhalOrganisationGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganisationJsonhalOrganisationGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganisationJsonhalOrganisationGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganisationJsonhalOrganisationGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *OrganisationJsonhalOrganisationGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OrganisationJsonhalOrganisationGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OrganisationJsonhalOrganisationGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OrganisationJsonhalOrganisationGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *OrganisationJsonhalOrganisationGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganisationJsonhalOrganisationGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganisationJsonhalOrganisationGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganisationJsonhalOrganisationGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganisationJsonhalOrganisationGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganisationJsonhalOrganisationGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *OrganisationJsonhalOrganisationGet) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *OrganisationJsonhalOrganisationGet) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *OrganisationJsonhalOrganisationGet) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *OrganisationJsonhalOrganisationGet) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *OrganisationJsonhalOrganisationGet) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *OrganisationJsonhalOrganisationGet) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *OrganisationJsonhalOrganisationGet) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OrganisationJsonhalOrganisationGet) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OrganisationJsonhalOrganisationGet) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OrganisationJsonhalOrganisationGet) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *OrganisationJsonhalOrganisationGet) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *OrganisationJsonhalOrganisationGet) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetTeam

`func (o *OrganisationJsonhalOrganisationGet) GetTeam() []string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganisationJsonhalOrganisationGet) GetTeamOk() (*[]string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganisationJsonhalOrganisationGet) SetTeam(v []string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *OrganisationJsonhalOrganisationGet) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOrganisationEnvVar

`func (o *OrganisationJsonhalOrganisationGet) GetOrganisationEnvVar() []OrganisationEnvVarJsonhalOrganisationGet`

GetOrganisationEnvVar returns the OrganisationEnvVar field if non-nil, zero value otherwise.

### GetOrganisationEnvVarOk

`func (o *OrganisationJsonhalOrganisationGet) GetOrganisationEnvVarOk() (*[]OrganisationEnvVarJsonhalOrganisationGet, bool)`

GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationEnvVar

`func (o *OrganisationJsonhalOrganisationGet) SetOrganisationEnvVar(v []OrganisationEnvVarJsonhalOrganisationGet)`

SetOrganisationEnvVar sets OrganisationEnvVar field to given value.

### HasOrganisationEnvVar

`func (o *OrganisationJsonhalOrganisationGet) HasOrganisationEnvVar() bool`

HasOrganisationEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


