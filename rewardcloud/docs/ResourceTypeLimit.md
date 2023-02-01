# ResourceTypeLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**ProjectMinValue** | Pointer to **NullableInt32** |  | [optional] 
**ProjectMaxValue** | Pointer to **NullableInt32** |  | [optional] 
**EnvironmentMinValue** | Pointer to **NullableInt32** |  | [optional] 
**EnvironmentMaxValue** | Pointer to **NullableInt32** |  | [optional] 
**ComponentMinValue** | Pointer to **NullableInt32** |  | [optional] 
**ComponentMaxValue** | Pointer to **NullableInt32** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResourceTypeLimit

`func NewResourceTypeLimit() *ResourceTypeLimit`

NewResourceTypeLimit instantiates a new ResourceTypeLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeLimitWithDefaults

`func NewResourceTypeLimitWithDefaults() *ResourceTypeLimit`

NewResourceTypeLimitWithDefaults instantiates a new ResourceTypeLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceTypeLimit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceTypeLimit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceTypeLimit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceTypeLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ResourceTypeLimit) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ResourceTypeLimit) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ResourceTypeLimit) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ResourceTypeLimit) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ResourceTypeLimit) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ResourceTypeLimit) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetProjectMinValue

`func (o *ResourceTypeLimit) GetProjectMinValue() int32`

GetProjectMinValue returns the ProjectMinValue field if non-nil, zero value otherwise.

### GetProjectMinValueOk

`func (o *ResourceTypeLimit) GetProjectMinValueOk() (*int32, bool)`

GetProjectMinValueOk returns a tuple with the ProjectMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMinValue

`func (o *ResourceTypeLimit) SetProjectMinValue(v int32)`

SetProjectMinValue sets ProjectMinValue field to given value.

### HasProjectMinValue

`func (o *ResourceTypeLimit) HasProjectMinValue() bool`

HasProjectMinValue returns a boolean if a field has been set.

### SetProjectMinValueNil

`func (o *ResourceTypeLimit) SetProjectMinValueNil(b bool)`

 SetProjectMinValueNil sets the value for ProjectMinValue to be an explicit nil

### UnsetProjectMinValue
`func (o *ResourceTypeLimit) UnsetProjectMinValue()`

UnsetProjectMinValue ensures that no value is present for ProjectMinValue, not even an explicit nil
### GetProjectMaxValue

`func (o *ResourceTypeLimit) GetProjectMaxValue() int32`

GetProjectMaxValue returns the ProjectMaxValue field if non-nil, zero value otherwise.

### GetProjectMaxValueOk

`func (o *ResourceTypeLimit) GetProjectMaxValueOk() (*int32, bool)`

GetProjectMaxValueOk returns a tuple with the ProjectMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMaxValue

`func (o *ResourceTypeLimit) SetProjectMaxValue(v int32)`

SetProjectMaxValue sets ProjectMaxValue field to given value.

### HasProjectMaxValue

`func (o *ResourceTypeLimit) HasProjectMaxValue() bool`

HasProjectMaxValue returns a boolean if a field has been set.

### SetProjectMaxValueNil

`func (o *ResourceTypeLimit) SetProjectMaxValueNil(b bool)`

 SetProjectMaxValueNil sets the value for ProjectMaxValue to be an explicit nil

### UnsetProjectMaxValue
`func (o *ResourceTypeLimit) UnsetProjectMaxValue()`

UnsetProjectMaxValue ensures that no value is present for ProjectMaxValue, not even an explicit nil
### GetEnvironmentMinValue

`func (o *ResourceTypeLimit) GetEnvironmentMinValue() int32`

GetEnvironmentMinValue returns the EnvironmentMinValue field if non-nil, zero value otherwise.

### GetEnvironmentMinValueOk

`func (o *ResourceTypeLimit) GetEnvironmentMinValueOk() (*int32, bool)`

GetEnvironmentMinValueOk returns a tuple with the EnvironmentMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentMinValue

`func (o *ResourceTypeLimit) SetEnvironmentMinValue(v int32)`

SetEnvironmentMinValue sets EnvironmentMinValue field to given value.

### HasEnvironmentMinValue

`func (o *ResourceTypeLimit) HasEnvironmentMinValue() bool`

HasEnvironmentMinValue returns a boolean if a field has been set.

### SetEnvironmentMinValueNil

`func (o *ResourceTypeLimit) SetEnvironmentMinValueNil(b bool)`

 SetEnvironmentMinValueNil sets the value for EnvironmentMinValue to be an explicit nil

### UnsetEnvironmentMinValue
`func (o *ResourceTypeLimit) UnsetEnvironmentMinValue()`

UnsetEnvironmentMinValue ensures that no value is present for EnvironmentMinValue, not even an explicit nil
### GetEnvironmentMaxValue

`func (o *ResourceTypeLimit) GetEnvironmentMaxValue() int32`

GetEnvironmentMaxValue returns the EnvironmentMaxValue field if non-nil, zero value otherwise.

### GetEnvironmentMaxValueOk

`func (o *ResourceTypeLimit) GetEnvironmentMaxValueOk() (*int32, bool)`

GetEnvironmentMaxValueOk returns a tuple with the EnvironmentMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentMaxValue

`func (o *ResourceTypeLimit) SetEnvironmentMaxValue(v int32)`

SetEnvironmentMaxValue sets EnvironmentMaxValue field to given value.

### HasEnvironmentMaxValue

`func (o *ResourceTypeLimit) HasEnvironmentMaxValue() bool`

HasEnvironmentMaxValue returns a boolean if a field has been set.

### SetEnvironmentMaxValueNil

`func (o *ResourceTypeLimit) SetEnvironmentMaxValueNil(b bool)`

 SetEnvironmentMaxValueNil sets the value for EnvironmentMaxValue to be an explicit nil

### UnsetEnvironmentMaxValue
`func (o *ResourceTypeLimit) UnsetEnvironmentMaxValue()`

UnsetEnvironmentMaxValue ensures that no value is present for EnvironmentMaxValue, not even an explicit nil
### GetComponentMinValue

`func (o *ResourceTypeLimit) GetComponentMinValue() int32`

GetComponentMinValue returns the ComponentMinValue field if non-nil, zero value otherwise.

### GetComponentMinValueOk

`func (o *ResourceTypeLimit) GetComponentMinValueOk() (*int32, bool)`

GetComponentMinValueOk returns a tuple with the ComponentMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMinValue

`func (o *ResourceTypeLimit) SetComponentMinValue(v int32)`

SetComponentMinValue sets ComponentMinValue field to given value.

### HasComponentMinValue

`func (o *ResourceTypeLimit) HasComponentMinValue() bool`

HasComponentMinValue returns a boolean if a field has been set.

### SetComponentMinValueNil

`func (o *ResourceTypeLimit) SetComponentMinValueNil(b bool)`

 SetComponentMinValueNil sets the value for ComponentMinValue to be an explicit nil

### UnsetComponentMinValue
`func (o *ResourceTypeLimit) UnsetComponentMinValue()`

UnsetComponentMinValue ensures that no value is present for ComponentMinValue, not even an explicit nil
### GetComponentMaxValue

`func (o *ResourceTypeLimit) GetComponentMaxValue() int32`

GetComponentMaxValue returns the ComponentMaxValue field if non-nil, zero value otherwise.

### GetComponentMaxValueOk

`func (o *ResourceTypeLimit) GetComponentMaxValueOk() (*int32, bool)`

GetComponentMaxValueOk returns a tuple with the ComponentMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMaxValue

`func (o *ResourceTypeLimit) SetComponentMaxValue(v int32)`

SetComponentMaxValue sets ComponentMaxValue field to given value.

### HasComponentMaxValue

`func (o *ResourceTypeLimit) HasComponentMaxValue() bool`

HasComponentMaxValue returns a boolean if a field has been set.

### SetComponentMaxValueNil

`func (o *ResourceTypeLimit) SetComponentMaxValueNil(b bool)`

 SetComponentMaxValueNil sets the value for ComponentMaxValue to be an explicit nil

### UnsetComponentMaxValue
`func (o *ResourceTypeLimit) UnsetComponentMaxValue()`

UnsetComponentMaxValue ensures that no value is present for ComponentMaxValue, not even an explicit nil
### GetCreatedBy

`func (o *ResourceTypeLimit) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ResourceTypeLimit) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ResourceTypeLimit) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ResourceTypeLimit) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ResourceTypeLimit) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ResourceTypeLimit) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ResourceTypeLimit) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ResourceTypeLimit) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ResourceTypeLimit) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ResourceTypeLimit) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ResourceTypeLimit) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ResourceTypeLimit) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ResourceTypeLimit) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceTypeLimit) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceTypeLimit) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceTypeLimit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceTypeLimit) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceTypeLimit) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceTypeLimit) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceTypeLimit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


