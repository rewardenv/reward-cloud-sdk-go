# EnvironmentAccessDevToolsJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**EnvironmentAccess** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvironmentAccessDevToolsJsonhal

`func NewEnvironmentAccessDevToolsJsonhal() *EnvironmentAccessDevToolsJsonhal`

NewEnvironmentAccessDevToolsJsonhal instantiates a new EnvironmentAccessDevToolsJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAccessDevToolsJsonhalWithDefaults

`func NewEnvironmentAccessDevToolsJsonhalWithDefaults() *EnvironmentAccessDevToolsJsonhal`

NewEnvironmentAccessDevToolsJsonhalWithDefaults instantiates a new EnvironmentAccessDevToolsJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentAccessDevToolsJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentAccessDevToolsJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentAccessDevToolsJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentAccessDevToolsJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentAccessDevToolsJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentAccessDevToolsJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentAccessDevToolsJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentAccessDevToolsJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentAccessDevToolsJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetUrl

`func (o *EnvironmentAccessDevToolsJsonhal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EnvironmentAccessDevToolsJsonhal) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EnvironmentAccessDevToolsJsonhal) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUsername

`func (o *EnvironmentAccessDevToolsJsonhal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnvironmentAccessDevToolsJsonhal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnvironmentAccessDevToolsJsonhal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *EnvironmentAccessDevToolsJsonhal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnvironmentAccessDevToolsJsonhal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnvironmentAccessDevToolsJsonhal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetEnvironmentAccess

`func (o *EnvironmentAccessDevToolsJsonhal) GetEnvironmentAccess() string`

GetEnvironmentAccess returns the EnvironmentAccess field if non-nil, zero value otherwise.

### GetEnvironmentAccessOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetEnvironmentAccessOk() (*string, bool)`

GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccess

`func (o *EnvironmentAccessDevToolsJsonhal) SetEnvironmentAccess(v string)`

SetEnvironmentAccess sets EnvironmentAccess field to given value.

### HasEnvironmentAccess

`func (o *EnvironmentAccessDevToolsJsonhal) HasEnvironmentAccess() bool`

HasEnvironmentAccess returns a boolean if a field has been set.

### SetEnvironmentAccessNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetEnvironmentAccessNil(b bool)`

 SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil

### UnsetEnvironmentAccess
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetEnvironmentAccess()`

UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentAccessDevToolsJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentAccessDevToolsJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentAccessDevToolsJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentAccessDevToolsJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentAccessDevToolsJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentAccessDevToolsJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentAccessDevToolsJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentAccessDevToolsJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentAccessDevToolsJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentAccessDevToolsJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentAccessDevToolsJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentAccessDevToolsJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentAccessDevToolsJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentAccessDevToolsJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentAccessDevToolsJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


