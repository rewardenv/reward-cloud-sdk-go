# EnvironmentAccessFrontendJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Domain** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**EnvironmentAccess** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvironmentAccessFrontendJsonhal

`func NewEnvironmentAccessFrontendJsonhal() *EnvironmentAccessFrontendJsonhal`

NewEnvironmentAccessFrontendJsonhal instantiates a new EnvironmentAccessFrontendJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAccessFrontendJsonhalWithDefaults

`func NewEnvironmentAccessFrontendJsonhalWithDefaults() *EnvironmentAccessFrontendJsonhal`

NewEnvironmentAccessFrontendJsonhalWithDefaults instantiates a new EnvironmentAccessFrontendJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentAccessFrontendJsonhal) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentAccessFrontendJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentAccessFrontendJsonhal) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentAccessFrontendJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentAccessFrontendJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentAccessFrontendJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentAccessFrontendJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentAccessFrontendJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentAccessFrontendJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentAccessFrontendJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentAccessFrontendJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentAccessFrontendJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentAccessFrontendJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentAccessFrontendJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetDomain

`func (o *EnvironmentAccessFrontendJsonhal) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EnvironmentAccessFrontendJsonhal) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EnvironmentAccessFrontendJsonhal) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *EnvironmentAccessFrontendJsonhal) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *EnvironmentAccessFrontendJsonhal) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *EnvironmentAccessFrontendJsonhal) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetPath

`func (o *EnvironmentAccessFrontendJsonhal) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EnvironmentAccessFrontendJsonhal) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EnvironmentAccessFrontendJsonhal) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EnvironmentAccessFrontendJsonhal) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *EnvironmentAccessFrontendJsonhal) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *EnvironmentAccessFrontendJsonhal) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetUrl

`func (o *EnvironmentAccessFrontendJsonhal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EnvironmentAccessFrontendJsonhal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EnvironmentAccessFrontendJsonhal) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EnvironmentAccessFrontendJsonhal) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EnvironmentAccessFrontendJsonhal) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EnvironmentAccessFrontendJsonhal) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetEnvironmentAccess

`func (o *EnvironmentAccessFrontendJsonhal) GetEnvironmentAccess() string`

GetEnvironmentAccess returns the EnvironmentAccess field if non-nil, zero value otherwise.

### GetEnvironmentAccessOk

`func (o *EnvironmentAccessFrontendJsonhal) GetEnvironmentAccessOk() (*string, bool)`

GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccess

`func (o *EnvironmentAccessFrontendJsonhal) SetEnvironmentAccess(v string)`

SetEnvironmentAccess sets EnvironmentAccess field to given value.

### HasEnvironmentAccess

`func (o *EnvironmentAccessFrontendJsonhal) HasEnvironmentAccess() bool`

HasEnvironmentAccess returns a boolean if a field has been set.

### SetEnvironmentAccessNil

`func (o *EnvironmentAccessFrontendJsonhal) SetEnvironmentAccessNil(b bool)`

 SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil

### UnsetEnvironmentAccess
`func (o *EnvironmentAccessFrontendJsonhal) UnsetEnvironmentAccess()`

UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentAccessFrontendJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentAccessFrontendJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentAccessFrontendJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentAccessFrontendJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentAccessFrontendJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentAccessFrontendJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentAccessFrontendJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentAccessFrontendJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentAccessFrontendJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentAccessFrontendJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentAccessFrontendJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentAccessFrontendJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentAccessFrontendJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentAccessFrontendJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentAccessFrontendJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentAccessFrontendJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


