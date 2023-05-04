# ProjectTypeVersionEnvVarExampleJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**ProjectTypeVersionEnvVar** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProjectTypeVersionEnvVarExampleJsonhal

`func NewProjectTypeVersionEnvVarExampleJsonhal() *ProjectTypeVersionEnvVarExampleJsonhal`

NewProjectTypeVersionEnvVarExampleJsonhal instantiates a new ProjectTypeVersionEnvVarExampleJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTypeVersionEnvVarExampleJsonhalWithDefaults

`func NewProjectTypeVersionEnvVarExampleJsonhalWithDefaults() *ProjectTypeVersionEnvVarExampleJsonhal`

NewProjectTypeVersionEnvVarExampleJsonhalWithDefaults instantiates a new ProjectTypeVersionEnvVarExampleJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetValue

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsDefault

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetProjectTypeVersionEnvVar

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetProjectTypeVersionEnvVar() string`

GetProjectTypeVersionEnvVar returns the ProjectTypeVersionEnvVar field if non-nil, zero value otherwise.

### GetProjectTypeVersionEnvVarOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetProjectTypeVersionEnvVarOk() (*string, bool)`

GetProjectTypeVersionEnvVarOk returns a tuple with the ProjectTypeVersionEnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeVersionEnvVar

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetProjectTypeVersionEnvVar(v string)`

SetProjectTypeVersionEnvVar sets ProjectTypeVersionEnvVar field to given value.

### HasProjectTypeVersionEnvVar

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasProjectTypeVersionEnvVar() bool`

HasProjectTypeVersionEnvVar returns a boolean if a field has been set.

### SetProjectTypeVersionEnvVarNil

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetProjectTypeVersionEnvVarNil(b bool)`

 SetProjectTypeVersionEnvVarNil sets the value for ProjectTypeVersionEnvVar to be an explicit nil

### UnsetProjectTypeVersionEnvVar
`func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetProjectTypeVersionEnvVar()`

UnsetProjectTypeVersionEnvVar ensures that no value is present for ProjectTypeVersionEnvVar, not even an explicit nil
### GetCreatedBy

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


