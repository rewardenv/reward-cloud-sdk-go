# EnvironmentAccessRabbitJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**EnvironmentAccess** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvironmentAccessRabbitJsonhal

`func NewEnvironmentAccessRabbitJsonhal() *EnvironmentAccessRabbitJsonhal`

NewEnvironmentAccessRabbitJsonhal instantiates a new EnvironmentAccessRabbitJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAccessRabbitJsonhalWithDefaults

`func NewEnvironmentAccessRabbitJsonhalWithDefaults() *EnvironmentAccessRabbitJsonhal`

NewEnvironmentAccessRabbitJsonhalWithDefaults instantiates a new EnvironmentAccessRabbitJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentAccessRabbitJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentAccessRabbitJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentAccessRabbitJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentAccessRabbitJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentAccessRabbitJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentAccessRabbitJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentAccessRabbitJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentAccessRabbitJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentAccessRabbitJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentAccessRabbitJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentAccessRabbitJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentAccessRabbitJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentAccessRabbitJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentAccessRabbitJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetUrl

`func (o *EnvironmentAccessRabbitJsonhal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EnvironmentAccessRabbitJsonhal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EnvironmentAccessRabbitJsonhal) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EnvironmentAccessRabbitJsonhal) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EnvironmentAccessRabbitJsonhal) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EnvironmentAccessRabbitJsonhal) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetHost

`func (o *EnvironmentAccessRabbitJsonhal) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EnvironmentAccessRabbitJsonhal) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EnvironmentAccessRabbitJsonhal) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EnvironmentAccessRabbitJsonhal) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *EnvironmentAccessRabbitJsonhal) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *EnvironmentAccessRabbitJsonhal) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *EnvironmentAccessRabbitJsonhal) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EnvironmentAccessRabbitJsonhal) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EnvironmentAccessRabbitJsonhal) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *EnvironmentAccessRabbitJsonhal) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *EnvironmentAccessRabbitJsonhal) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *EnvironmentAccessRabbitJsonhal) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *EnvironmentAccessRabbitJsonhal) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *EnvironmentAccessRabbitJsonhal) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *EnvironmentAccessRabbitJsonhal) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *EnvironmentAccessRabbitJsonhal) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *EnvironmentAccessRabbitJsonhal) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *EnvironmentAccessRabbitJsonhal) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *EnvironmentAccessRabbitJsonhal) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnvironmentAccessRabbitJsonhal) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnvironmentAccessRabbitJsonhal) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *EnvironmentAccessRabbitJsonhal) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *EnvironmentAccessRabbitJsonhal) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *EnvironmentAccessRabbitJsonhal) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetEnvironmentAccess

`func (o *EnvironmentAccessRabbitJsonhal) GetEnvironmentAccess() string`

GetEnvironmentAccess returns the EnvironmentAccess field if non-nil, zero value otherwise.

### GetEnvironmentAccessOk

`func (o *EnvironmentAccessRabbitJsonhal) GetEnvironmentAccessOk() (*string, bool)`

GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccess

`func (o *EnvironmentAccessRabbitJsonhal) SetEnvironmentAccess(v string)`

SetEnvironmentAccess sets EnvironmentAccess field to given value.

### HasEnvironmentAccess

`func (o *EnvironmentAccessRabbitJsonhal) HasEnvironmentAccess() bool`

HasEnvironmentAccess returns a boolean if a field has been set.

### SetEnvironmentAccessNil

`func (o *EnvironmentAccessRabbitJsonhal) SetEnvironmentAccessNil(b bool)`

 SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil

### UnsetEnvironmentAccess
`func (o *EnvironmentAccessRabbitJsonhal) UnsetEnvironmentAccess()`

UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentAccessRabbitJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentAccessRabbitJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentAccessRabbitJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentAccessRabbitJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentAccessRabbitJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentAccessRabbitJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentAccessRabbitJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentAccessRabbitJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentAccessRabbitJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentAccessRabbitJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentAccessRabbitJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentAccessRabbitJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentAccessRabbitJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentAccessRabbitJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentAccessRabbitJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentAccessRabbitJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentAccessRabbitJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentAccessRabbitJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentAccessRabbitJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentAccessRabbitJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


