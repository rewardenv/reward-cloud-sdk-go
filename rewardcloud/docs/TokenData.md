# TokenData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Fullname** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Teams** | Pointer to [**[]TokenDataTeamsInner**](TokenDataTeamsInner.md) |  | [optional] 

## Methods

### NewTokenData

`func NewTokenData() *TokenData`

NewTokenData instantiates a new TokenData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDataWithDefaults

`func NewTokenDataWithDefaults() *TokenData`

NewTokenDataWithDefaults instantiates a new TokenData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TokenData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFullname

`func (o *TokenData) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *TokenData) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *TokenData) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *TokenData) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### GetRoles

`func (o *TokenData) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TokenData) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TokenData) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TokenData) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTeams

`func (o *TokenData) GetTeams() []TokenDataTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *TokenData) GetTeamsOk() (*[]TokenDataTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *TokenData) SetTeams(v []TokenDataTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *TokenData) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


