# EnvironmentAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewEnvironmentAccess

`func NewEnvironmentAccess() *EnvironmentAccess`

NewEnvironmentAccess instantiates a new EnvironmentAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAccessWithDefaults

`func NewEnvironmentAccessWithDefaults() *EnvironmentAccess`

NewEnvironmentAccessWithDefaults instantiates a new EnvironmentAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentAccess) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentAccess) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentAccess) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentAccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentAccess) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentAccess) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentAccess) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentAccess) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentAccess) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentAccess) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetEnvironment

`func (o *EnvironmentAccess) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentAccess) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentAccess) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentAccess) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvironmentAccess) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvironmentAccess) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetFrontend

`func (o *EnvironmentAccess) GetFrontend() string`

GetFrontend returns the Frontend field if non-nil, zero value otherwise.

### GetFrontendOk

`func (o *EnvironmentAccess) GetFrontendOk() (*string, bool)`

GetFrontendOk returns a tuple with the Frontend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontend

`func (o *EnvironmentAccess) SetFrontend(v string)`

SetFrontend sets Frontend field to given value.

### HasFrontend

`func (o *EnvironmentAccess) HasFrontend() bool`

HasFrontend returns a boolean if a field has been set.

### SetFrontendNil

`func (o *EnvironmentAccess) SetFrontendNil(b bool)`

 SetFrontendNil sets the value for Frontend to be an explicit nil

### UnsetFrontend
`func (o *EnvironmentAccess) UnsetFrontend()`

UnsetFrontend ensures that no value is present for Frontend, not even an explicit nil
### GetBackend

`func (o *EnvironmentAccess) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *EnvironmentAccess) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *EnvironmentAccess) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *EnvironmentAccess) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### SetBackendNil

`func (o *EnvironmentAccess) SetBackendNil(b bool)`

 SetBackendNil sets the value for Backend to be an explicit nil

### UnsetBackend
`func (o *EnvironmentAccess) UnsetBackend()`

UnsetBackend ensures that no value is present for Backend, not even an explicit nil
### GetDatabase

`func (o *EnvironmentAccess) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *EnvironmentAccess) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *EnvironmentAccess) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *EnvironmentAccess) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### SetDatabaseNil

`func (o *EnvironmentAccess) SetDatabaseNil(b bool)`

 SetDatabaseNil sets the value for Database to be an explicit nil

### UnsetDatabase
`func (o *EnvironmentAccess) UnsetDatabase()`

UnsetDatabase ensures that no value is present for Database, not even an explicit nil
### GetDevTools

`func (o *EnvironmentAccess) GetDevTools() string`

GetDevTools returns the DevTools field if non-nil, zero value otherwise.

### GetDevToolsOk

`func (o *EnvironmentAccess) GetDevToolsOk() (*string, bool)`

GetDevToolsOk returns a tuple with the DevTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevTools

`func (o *EnvironmentAccess) SetDevTools(v string)`

SetDevTools sets DevTools field to given value.

### HasDevTools

`func (o *EnvironmentAccess) HasDevTools() bool`

HasDevTools returns a boolean if a field has been set.

### SetDevToolsNil

`func (o *EnvironmentAccess) SetDevToolsNil(b bool)`

 SetDevToolsNil sets the value for DevTools to be an explicit nil

### UnsetDevTools
`func (o *EnvironmentAccess) UnsetDevTools()`

UnsetDevTools ensures that no value is present for DevTools, not even an explicit nil
### GetRedis

`func (o *EnvironmentAccess) GetRedis() string`

GetRedis returns the Redis field if non-nil, zero value otherwise.

### GetRedisOk

`func (o *EnvironmentAccess) GetRedisOk() (*string, bool)`

GetRedisOk returns a tuple with the Redis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedis

`func (o *EnvironmentAccess) SetRedis(v string)`

SetRedis sets Redis field to given value.

### HasRedis

`func (o *EnvironmentAccess) HasRedis() bool`

HasRedis returns a boolean if a field has been set.

### SetRedisNil

`func (o *EnvironmentAccess) SetRedisNil(b bool)`

 SetRedisNil sets the value for Redis to be an explicit nil

### UnsetRedis
`func (o *EnvironmentAccess) UnsetRedis()`

UnsetRedis ensures that no value is present for Redis, not even an explicit nil
### GetRabbit

`func (o *EnvironmentAccess) GetRabbit() string`

GetRabbit returns the Rabbit field if non-nil, zero value otherwise.

### GetRabbitOk

`func (o *EnvironmentAccess) GetRabbitOk() (*string, bool)`

GetRabbitOk returns a tuple with the Rabbit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbit

`func (o *EnvironmentAccess) SetRabbit(v string)`

SetRabbit sets Rabbit field to given value.

### HasRabbit

`func (o *EnvironmentAccess) HasRabbit() bool`

HasRabbit returns a boolean if a field has been set.

### SetRabbitNil

`func (o *EnvironmentAccess) SetRabbitNil(b bool)`

 SetRabbitNil sets the value for Rabbit to be an explicit nil

### UnsetRabbit
`func (o *EnvironmentAccess) UnsetRabbit()`

UnsetRabbit ensures that no value is present for Rabbit, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentAccess) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentAccess) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentAccess) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentAccess) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentAccess) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentAccess) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentAccess) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentAccess) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentAccess) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentAccess) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentAccess) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentAccess) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentAccess) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentAccess) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentAccess) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentAccess) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentAccess) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentAccess) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentAccess) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentAccess) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


