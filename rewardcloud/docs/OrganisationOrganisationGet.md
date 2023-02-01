# OrganisationOrganisationGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**Team** | Pointer to **[]string** |  | [optional] 
**OrganisationEnvVar** | Pointer to [**[]OrganisationEnvVarOrganisationGet**](OrganisationEnvVarOrganisationGet.md) |  | [optional] 

## Methods

### NewOrganisationOrganisationGet

`func NewOrganisationOrganisationGet() *OrganisationOrganisationGet`

NewOrganisationOrganisationGet instantiates a new OrganisationOrganisationGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationOrganisationGetWithDefaults

`func NewOrganisationOrganisationGetWithDefaults() *OrganisationOrganisationGet`

NewOrganisationOrganisationGetWithDefaults instantiates a new OrganisationOrganisationGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganisationOrganisationGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganisationOrganisationGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganisationOrganisationGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganisationOrganisationGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *OrganisationOrganisationGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OrganisationOrganisationGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OrganisationOrganisationGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OrganisationOrganisationGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *OrganisationOrganisationGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganisationOrganisationGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganisationOrganisationGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganisationOrganisationGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganisationOrganisationGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganisationOrganisationGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *OrganisationOrganisationGet) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *OrganisationOrganisationGet) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *OrganisationOrganisationGet) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *OrganisationOrganisationGet) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *OrganisationOrganisationGet) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *OrganisationOrganisationGet) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetIsDefault

`func (o *OrganisationOrganisationGet) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *OrganisationOrganisationGet) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *OrganisationOrganisationGet) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *OrganisationOrganisationGet) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *OrganisationOrganisationGet) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *OrganisationOrganisationGet) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetTeam

`func (o *OrganisationOrganisationGet) GetTeam() []string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganisationOrganisationGet) GetTeamOk() (*[]string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganisationOrganisationGet) SetTeam(v []string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *OrganisationOrganisationGet) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetOrganisationEnvVar

`func (o *OrganisationOrganisationGet) GetOrganisationEnvVar() []OrganisationEnvVarOrganisationGet`

GetOrganisationEnvVar returns the OrganisationEnvVar field if non-nil, zero value otherwise.

### GetOrganisationEnvVarOk

`func (o *OrganisationOrganisationGet) GetOrganisationEnvVarOk() (*[]OrganisationEnvVarOrganisationGet, bool)`

GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationEnvVar

`func (o *OrganisationOrganisationGet) SetOrganisationEnvVar(v []OrganisationEnvVarOrganisationGet)`

SetOrganisationEnvVar sets OrganisationEnvVar field to given value.

### HasOrganisationEnvVar

`func (o *OrganisationOrganisationGet) HasOrganisationEnvVar() bool`

HasOrganisationEnvVar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


