# ResourceTypeLimitJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
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

### NewResourceTypeLimitJsonhal

`func NewResourceTypeLimitJsonhal() *ResourceTypeLimitJsonhal`

NewResourceTypeLimitJsonhal instantiates a new ResourceTypeLimitJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceTypeLimitJsonhalWithDefaults

`func NewResourceTypeLimitJsonhalWithDefaults() *ResourceTypeLimitJsonhal`

NewResourceTypeLimitJsonhalWithDefaults instantiates a new ResourceTypeLimitJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ResourceTypeLimitJsonhal) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceTypeLimitJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceTypeLimitJsonhal) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceTypeLimitJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ResourceTypeLimitJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceTypeLimitJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceTypeLimitJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceTypeLimitJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ResourceTypeLimitJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ResourceTypeLimitJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ResourceTypeLimitJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ResourceTypeLimitJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ResourceTypeLimitJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ResourceTypeLimitJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetProjectMinValue

`func (o *ResourceTypeLimitJsonhal) GetProjectMinValue() int32`

GetProjectMinValue returns the ProjectMinValue field if non-nil, zero value otherwise.

### GetProjectMinValueOk

`func (o *ResourceTypeLimitJsonhal) GetProjectMinValueOk() (*int32, bool)`

GetProjectMinValueOk returns a tuple with the ProjectMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMinValue

`func (o *ResourceTypeLimitJsonhal) SetProjectMinValue(v int32)`

SetProjectMinValue sets ProjectMinValue field to given value.

### HasProjectMinValue

`func (o *ResourceTypeLimitJsonhal) HasProjectMinValue() bool`

HasProjectMinValue returns a boolean if a field has been set.

### SetProjectMinValueNil

`func (o *ResourceTypeLimitJsonhal) SetProjectMinValueNil(b bool)`

 SetProjectMinValueNil sets the value for ProjectMinValue to be an explicit nil

### UnsetProjectMinValue
`func (o *ResourceTypeLimitJsonhal) UnsetProjectMinValue()`

UnsetProjectMinValue ensures that no value is present for ProjectMinValue, not even an explicit nil
### GetProjectMaxValue

`func (o *ResourceTypeLimitJsonhal) GetProjectMaxValue() int32`

GetProjectMaxValue returns the ProjectMaxValue field if non-nil, zero value otherwise.

### GetProjectMaxValueOk

`func (o *ResourceTypeLimitJsonhal) GetProjectMaxValueOk() (*int32, bool)`

GetProjectMaxValueOk returns a tuple with the ProjectMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMaxValue

`func (o *ResourceTypeLimitJsonhal) SetProjectMaxValue(v int32)`

SetProjectMaxValue sets ProjectMaxValue field to given value.

### HasProjectMaxValue

`func (o *ResourceTypeLimitJsonhal) HasProjectMaxValue() bool`

HasProjectMaxValue returns a boolean if a field has been set.

### SetProjectMaxValueNil

`func (o *ResourceTypeLimitJsonhal) SetProjectMaxValueNil(b bool)`

 SetProjectMaxValueNil sets the value for ProjectMaxValue to be an explicit nil

### UnsetProjectMaxValue
`func (o *ResourceTypeLimitJsonhal) UnsetProjectMaxValue()`

UnsetProjectMaxValue ensures that no value is present for ProjectMaxValue, not even an explicit nil
### GetEnvironmentMinValue

`func (o *ResourceTypeLimitJsonhal) GetEnvironmentMinValue() int32`

GetEnvironmentMinValue returns the EnvironmentMinValue field if non-nil, zero value otherwise.

### GetEnvironmentMinValueOk

`func (o *ResourceTypeLimitJsonhal) GetEnvironmentMinValueOk() (*int32, bool)`

GetEnvironmentMinValueOk returns a tuple with the EnvironmentMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentMinValue

`func (o *ResourceTypeLimitJsonhal) SetEnvironmentMinValue(v int32)`

SetEnvironmentMinValue sets EnvironmentMinValue field to given value.

### HasEnvironmentMinValue

`func (o *ResourceTypeLimitJsonhal) HasEnvironmentMinValue() bool`

HasEnvironmentMinValue returns a boolean if a field has been set.

### SetEnvironmentMinValueNil

`func (o *ResourceTypeLimitJsonhal) SetEnvironmentMinValueNil(b bool)`

 SetEnvironmentMinValueNil sets the value for EnvironmentMinValue to be an explicit nil

### UnsetEnvironmentMinValue
`func (o *ResourceTypeLimitJsonhal) UnsetEnvironmentMinValue()`

UnsetEnvironmentMinValue ensures that no value is present for EnvironmentMinValue, not even an explicit nil
### GetEnvironmentMaxValue

`func (o *ResourceTypeLimitJsonhal) GetEnvironmentMaxValue() int32`

GetEnvironmentMaxValue returns the EnvironmentMaxValue field if non-nil, zero value otherwise.

### GetEnvironmentMaxValueOk

`func (o *ResourceTypeLimitJsonhal) GetEnvironmentMaxValueOk() (*int32, bool)`

GetEnvironmentMaxValueOk returns a tuple with the EnvironmentMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentMaxValue

`func (o *ResourceTypeLimitJsonhal) SetEnvironmentMaxValue(v int32)`

SetEnvironmentMaxValue sets EnvironmentMaxValue field to given value.

### HasEnvironmentMaxValue

`func (o *ResourceTypeLimitJsonhal) HasEnvironmentMaxValue() bool`

HasEnvironmentMaxValue returns a boolean if a field has been set.

### SetEnvironmentMaxValueNil

`func (o *ResourceTypeLimitJsonhal) SetEnvironmentMaxValueNil(b bool)`

 SetEnvironmentMaxValueNil sets the value for EnvironmentMaxValue to be an explicit nil

### UnsetEnvironmentMaxValue
`func (o *ResourceTypeLimitJsonhal) UnsetEnvironmentMaxValue()`

UnsetEnvironmentMaxValue ensures that no value is present for EnvironmentMaxValue, not even an explicit nil
### GetComponentMinValue

`func (o *ResourceTypeLimitJsonhal) GetComponentMinValue() int32`

GetComponentMinValue returns the ComponentMinValue field if non-nil, zero value otherwise.

### GetComponentMinValueOk

`func (o *ResourceTypeLimitJsonhal) GetComponentMinValueOk() (*int32, bool)`

GetComponentMinValueOk returns a tuple with the ComponentMinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMinValue

`func (o *ResourceTypeLimitJsonhal) SetComponentMinValue(v int32)`

SetComponentMinValue sets ComponentMinValue field to given value.

### HasComponentMinValue

`func (o *ResourceTypeLimitJsonhal) HasComponentMinValue() bool`

HasComponentMinValue returns a boolean if a field has been set.

### SetComponentMinValueNil

`func (o *ResourceTypeLimitJsonhal) SetComponentMinValueNil(b bool)`

 SetComponentMinValueNil sets the value for ComponentMinValue to be an explicit nil

### UnsetComponentMinValue
`func (o *ResourceTypeLimitJsonhal) UnsetComponentMinValue()`

UnsetComponentMinValue ensures that no value is present for ComponentMinValue, not even an explicit nil
### GetComponentMaxValue

`func (o *ResourceTypeLimitJsonhal) GetComponentMaxValue() int32`

GetComponentMaxValue returns the ComponentMaxValue field if non-nil, zero value otherwise.

### GetComponentMaxValueOk

`func (o *ResourceTypeLimitJsonhal) GetComponentMaxValueOk() (*int32, bool)`

GetComponentMaxValueOk returns a tuple with the ComponentMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMaxValue

`func (o *ResourceTypeLimitJsonhal) SetComponentMaxValue(v int32)`

SetComponentMaxValue sets ComponentMaxValue field to given value.

### HasComponentMaxValue

`func (o *ResourceTypeLimitJsonhal) HasComponentMaxValue() bool`

HasComponentMaxValue returns a boolean if a field has been set.

### SetComponentMaxValueNil

`func (o *ResourceTypeLimitJsonhal) SetComponentMaxValueNil(b bool)`

 SetComponentMaxValueNil sets the value for ComponentMaxValue to be an explicit nil

### UnsetComponentMaxValue
`func (o *ResourceTypeLimitJsonhal) UnsetComponentMaxValue()`

UnsetComponentMaxValue ensures that no value is present for ComponentMaxValue, not even an explicit nil
### GetCreatedBy

`func (o *ResourceTypeLimitJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ResourceTypeLimitJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ResourceTypeLimitJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ResourceTypeLimitJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ResourceTypeLimitJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ResourceTypeLimitJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ResourceTypeLimitJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ResourceTypeLimitJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ResourceTypeLimitJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ResourceTypeLimitJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ResourceTypeLimitJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ResourceTypeLimitJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ResourceTypeLimitJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceTypeLimitJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceTypeLimitJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceTypeLimitJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceTypeLimitJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceTypeLimitJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceTypeLimitJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceTypeLimitJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


