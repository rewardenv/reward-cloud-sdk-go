# EnvironmentAccessJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Environment** | Pointer to **NullableString** |  | [optional] 
**Frontend** | Pointer to **NullableString** |  | [optional] 
**Backend** | Pointer to **NullableString** |  | [optional] 
**Database** | Pointer to **NullableString** |  | [optional] 
**DevTools** | Pointer to **NullableString** |  | [optional] 
**Redis** | Pointer to **NullableString** |  | [optional] 
**Rabbit** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvironmentAccessJsonhal

`func NewEnvironmentAccessJsonhal() *EnvironmentAccessJsonhal`

NewEnvironmentAccessJsonhal instantiates a new EnvironmentAccessJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAccessJsonhalWithDefaults

`func NewEnvironmentAccessJsonhalWithDefaults() *EnvironmentAccessJsonhal`

NewEnvironmentAccessJsonhalWithDefaults instantiates a new EnvironmentAccessJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentAccessJsonhal) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentAccessJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentAccessJsonhal) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentAccessJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentAccessJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentAccessJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentAccessJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentAccessJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentAccessJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentAccessJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentAccessJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentAccessJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentAccessJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentAccessJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetEnvironment

`func (o *EnvironmentAccessJsonhal) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentAccessJsonhal) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentAccessJsonhal) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentAccessJsonhal) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvironmentAccessJsonhal) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvironmentAccessJsonhal) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetFrontend

`func (o *EnvironmentAccessJsonhal) GetFrontend() string`

GetFrontend returns the Frontend field if non-nil, zero value otherwise.

### GetFrontendOk

`func (o *EnvironmentAccessJsonhal) GetFrontendOk() (*string, bool)`

GetFrontendOk returns a tuple with the Frontend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontend

`func (o *EnvironmentAccessJsonhal) SetFrontend(v string)`

SetFrontend sets Frontend field to given value.

### HasFrontend

`func (o *EnvironmentAccessJsonhal) HasFrontend() bool`

HasFrontend returns a boolean if a field has been set.

### SetFrontendNil

`func (o *EnvironmentAccessJsonhal) SetFrontendNil(b bool)`

 SetFrontendNil sets the value for Frontend to be an explicit nil

### UnsetFrontend
`func (o *EnvironmentAccessJsonhal) UnsetFrontend()`

UnsetFrontend ensures that no value is present for Frontend, not even an explicit nil
### GetBackend

`func (o *EnvironmentAccessJsonhal) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *EnvironmentAccessJsonhal) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *EnvironmentAccessJsonhal) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *EnvironmentAccessJsonhal) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### SetBackendNil

`func (o *EnvironmentAccessJsonhal) SetBackendNil(b bool)`

 SetBackendNil sets the value for Backend to be an explicit nil

### UnsetBackend
`func (o *EnvironmentAccessJsonhal) UnsetBackend()`

UnsetBackend ensures that no value is present for Backend, not even an explicit nil
### GetDatabase

`func (o *EnvironmentAccessJsonhal) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *EnvironmentAccessJsonhal) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *EnvironmentAccessJsonhal) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *EnvironmentAccessJsonhal) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *EnvironmentAccessJsonhal) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *EnvironmentAccessJsonhal) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetDevTools

`func (o *EnvironmentAccessJsonhal) GetDevTools() string`

GetDevTools returns the DevTools field if non-nil, zero value otherwise.

### GetDevToolsOk

`func (o *EnvironmentAccessJsonhal) GetDevToolsOk() (*string, bool)`

GetDevToolsOk returns a tuple with the DevTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTools

`func (o *EnvironmentAccessJsonhal) SetDevTools(v string)`

SetDevTools sets DevTools field to given value.

### HasDevTools

`func (o *EnvironmentAccessJsonhal) HasDevTools() bool`

HasDevTools returns a boolean if a field has been set.

### SetDevToolsNil

`func (o *EnvironmentAccessJsonhal) SetDevToolsNil(b bool)`

 SetDevToolsNil sets the value for DevTools to be an explicit nil

### UnsetDevTools
`func (o *EnvironmentAccessJsonhal) UnsetDevTools()`

UnsetDevTools ensures that no value is present for DevTools, not even an explicit nil
### GetRedis

`func (o *EnvironmentAccessJsonhal) GetRedis() string`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *EnvironmentAccessJsonhal) GetRedisOk() (*string, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *EnvironmentAccessJsonhal) SetRedis(v string)`

SetRedis sets Redis field to given value.

### HasRedis

`func (o *EnvironmentAccessJsonhal) HasRedis() bool`

HasRedis returns a boolean if a field has been set.

### SetRedisNil

`func (o *EnvironmentAccessJsonhal) SetRedisNil(b bool)`

 SetRedisNil sets the value for Redis to be an explicit nil

### UnsetRedis
`func (o *EnvironmentAccessJsonhal) UnsetRedis()`

UnsetRedis ensures that no value is present for Redis, not even an explicit nil
### GetRabbit

`func (o *EnvironmentAccessJsonhal) GetRabbit() string`

GetRabbit returns the Rabbit field if non-nil, zero value otherwise.

### GetRabbitOk

`func (o *EnvironmentAccessJsonhal) GetRabbitOk() (*string, bool)`

GetRabbitOk returns a tuple with the Rabbit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbit

`func (o *EnvironmentAccessJsonhal) SetRabbit(v string)`

SetRabbit sets Rabbit field to given value.

### HasRabbit

`func (o *EnvironmentAccessJsonhal) HasRabbit() bool`

HasRabbit returns a boolean if a field has been set.

### SetRabbitNil

`func (o *EnvironmentAccessJsonhal) SetRabbitNil(b bool)`

 SetRabbitNil sets the value for Rabbit to be an explicit nil

### UnsetRabbit
`func (o *EnvironmentAccessJsonhal) UnsetRabbit()`

UnsetRabbit ensures that no value is present for Rabbit, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentAccessJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentAccessJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentAccessJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentAccessJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentAccessJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentAccessJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentAccessJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentAccessJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentAccessJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentAccessJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentAccessJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentAccessJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentAccessJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentAccessJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentAccessJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentAccessJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentAccessJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentAccessJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentAccessJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentAccessJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


