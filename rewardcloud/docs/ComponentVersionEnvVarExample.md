# ComponentVersionEnvVarExample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ComponentVersionEnvVar** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComponentVersionEnvVarExample

`func NewComponentVersionEnvVarExample() *ComponentVersionEnvVarExample`

NewComponentVersionEnvVarExample instantiates a new ComponentVersionEnvVarExample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionEnvVarExampleWithDefaults

`func NewComponentVersionEnvVarExampleWithDefaults() *ComponentVersionEnvVarExample`

NewComponentVersionEnvVarExampleWithDefaults instantiates a new ComponentVersionEnvVarExample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComponentVersionEnvVarExample) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentVersionEnvVarExample) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentVersionEnvVarExample) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentVersionEnvVarExample) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ComponentVersionEnvVarExample) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComponentVersionEnvVarExample) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComponentVersionEnvVarExample) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComponentVersionEnvVarExample) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ComponentVersionEnvVarExample) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ComponentVersionEnvVarExample) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetValue

`func (o *ComponentVersionEnvVarExample) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComponentVersionEnvVarExample) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComponentVersionEnvVarExample) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ComponentVersionEnvVarExample) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ComponentVersionEnvVarExample) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ComponentVersionEnvVarExample) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *ComponentVersionEnvVarExample) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ComponentVersionEnvVarExample) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ComponentVersionEnvVarExample) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ComponentVersionEnvVarExample) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *ComponentVersionEnvVarExample) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *ComponentVersionEnvVarExample) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetComponentVersionEnvVar

`func (o *ComponentVersionEnvVarExample) GetComponentVersionEnvVar() string`

GetComponentVersionEnvVar returns the ComponentVersionEnvVar field if non-nil, zero value otherwise.

### GetComponentVersionEnvVarOk

`func (o *ComponentVersionEnvVarExample) GetComponentVersionEnvVarOk() (*string, bool)`

GetComponentVersionEnvVarOk returns a tuple with the ComponentVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersionEnvVar

`func (o *ComponentVersionEnvVarExample) SetComponentVersionEnvVar(v string)`

SetComponentVersionEnvVar sets ComponentVersionEnvVar field to given value.

### HasComponentVersionEnvVar

`func (o *ComponentVersionEnvVarExample) HasComponentVersionEnvVar() bool`

HasComponentVersionEnvVar returns a boolean if a field has been set.

### SetComponentVersionEnvVarNil

`func (o *ComponentVersionEnvVarExample) SetComponentVersionEnvVarNil(b bool)`

 SetComponentVersionEnvVarNil sets the value for ComponentVersionEnvVar to be an explicit nil

### UnsetComponentVersionEnvVar
`func (o *ComponentVersionEnvVarExample) UnsetComponentVersionEnvVar()`

UnsetComponentVersionEnvVar ensures that no value is present for ComponentVersionEnvVar, not even an explicit nil
### GetCreatedBy

`func (o *ComponentVersionEnvVarExample) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ComponentVersionEnvVarExample) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ComponentVersionEnvVarExample) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ComponentVersionEnvVarExample) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ComponentVersionEnvVarExample) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ComponentVersionEnvVarExample) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ComponentVersionEnvVarExample) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ComponentVersionEnvVarExample) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ComponentVersionEnvVarExample) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ComponentVersionEnvVarExample) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ComponentVersionEnvVarExample) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ComponentVersionEnvVarExample) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ComponentVersionEnvVarExample) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentVersionEnvVarExample) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentVersionEnvVarExample) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComponentVersionEnvVarExample) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComponentVersionEnvVarExample) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComponentVersionEnvVarExample) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComponentVersionEnvVarExample) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComponentVersionEnvVarExample) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


