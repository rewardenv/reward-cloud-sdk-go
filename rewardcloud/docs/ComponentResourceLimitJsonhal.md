# ComponentResourceLimitJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**MinValue** | Pointer to **NullableInt32** |  | [optional] 
**MaxValue** | Pointer to **NullableInt32** |  | [optional] 
**ProjectTypeVersion** | Pointer to **NullableString** |  | [optional] 
**ResourceType** | Pointer to **NullableString** |  | [optional] 
**ComponentVersion** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComponentResourceLimitJsonhal

`func NewComponentResourceLimitJsonhal() *ComponentResourceLimitJsonhal`

NewComponentResourceLimitJsonhal instantiates a new ComponentResourceLimitJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentResourceLimitJsonhalWithDefaults

`func NewComponentResourceLimitJsonhalWithDefaults() *ComponentResourceLimitJsonhal`

NewComponentResourceLimitJsonhalWithDefaults instantiates a new ComponentResourceLimitJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ComponentResourceLimitJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ComponentResourceLimitJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ComponentResourceLimitJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ComponentResourceLimitJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ComponentResourceLimitJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentResourceLimitJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentResourceLimitJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentResourceLimitJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ComponentResourceLimitJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComponentResourceLimitJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComponentResourceLimitJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComponentResourceLimitJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ComponentResourceLimitJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ComponentResourceLimitJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetMinValue

`func (o *ComponentResourceLimitJsonhal) GetMinValue() int32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *ComponentResourceLimitJsonhal) GetMinValueOk() (*int32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *ComponentResourceLimitJsonhal) SetMinValue(v int32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *ComponentResourceLimitJsonhal) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *ComponentResourceLimitJsonhal) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *ComponentResourceLimitJsonhal) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetMaxValue

`func (o *ComponentResourceLimitJsonhal) GetMaxValue() int32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *ComponentResourceLimitJsonhal) GetMaxValueOk() (*int32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *ComponentResourceLimitJsonhal) SetMaxValue(v int32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *ComponentResourceLimitJsonhal) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *ComponentResourceLimitJsonhal) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *ComponentResourceLimitJsonhal) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetProjectTypeVersion

`func (o *ComponentResourceLimitJsonhal) GetProjectTypeVersion() string`

GetProjectTypeVersion returns the ProjectTypeVersion field if non-nil, zero value otherwise.

### GetProjectTypeVersionOk

`func (o *ComponentResourceLimitJsonhal) GetProjectTypeVersionOk() (*string, bool)`

GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersion

`func (o *ComponentResourceLimitJsonhal) SetProjectTypeVersion(v string)`

SetProjectTypeVersion sets ProjectTypeVersion field to given value.

### HasProjectTypeVersion

`func (o *ComponentResourceLimitJsonhal) HasProjectTypeVersion() bool`

HasProjectTypeVersion returns a boolean if a field has been set.

### SetProjectTypeVersionNil

`func (o *ComponentResourceLimitJsonhal) SetProjectTypeVersionNil(b bool)`

 SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil

### UnsetProjectTypeVersion
`func (o *ComponentResourceLimitJsonhal) UnsetProjectTypeVersion()`

UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
### GetResourceType

`func (o *ComponentResourceLimitJsonhal) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ComponentResourceLimitJsonhal) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ComponentResourceLimitJsonhal) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ComponentResourceLimitJsonhal) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *ComponentResourceLimitJsonhal) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *ComponentResourceLimitJsonhal) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetComponentVersion

`func (o *ComponentResourceLimitJsonhal) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ComponentResourceLimitJsonhal) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ComponentResourceLimitJsonhal) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ComponentResourceLimitJsonhal) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *ComponentResourceLimitJsonhal) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *ComponentResourceLimitJsonhal) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetCreatedBy

`func (o *ComponentResourceLimitJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ComponentResourceLimitJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ComponentResourceLimitJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ComponentResourceLimitJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ComponentResourceLimitJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ComponentResourceLimitJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ComponentResourceLimitJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ComponentResourceLimitJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ComponentResourceLimitJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ComponentResourceLimitJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ComponentResourceLimitJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ComponentResourceLimitJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ComponentResourceLimitJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentResourceLimitJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentResourceLimitJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComponentResourceLimitJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComponentResourceLimitJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComponentResourceLimitJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComponentResourceLimitJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComponentResourceLimitJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


