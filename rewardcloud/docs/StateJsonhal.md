# StateJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SeqNumber** | Pointer to **NullableInt32** |  | [optional] 
**Project** | Pointer to **[]string** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
**ImportedData** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStateJsonhal

`func NewStateJsonhal() *StateJsonhal`

NewStateJsonhal instantiates a new StateJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateJsonhalWithDefaults

`func NewStateJsonhalWithDefaults() *StateJsonhal`

NewStateJsonhalWithDefaults instantiates a new StateJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StateJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StateJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StateJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StateJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *StateJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StateJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *StateJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StateJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StateJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StateJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *StateJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *StateJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *StateJsonhal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StateJsonhal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StateJsonhal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StateJsonhal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StateJsonhal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StateJsonhal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSeqNumber

`func (o *StateJsonhal) GetSeqNumber() int32`

GetSeqNumber returns the SeqNumber field if non-nil, zero value otherwise.

### GetSeqNumberOk

`func (o *StateJsonhal) GetSeqNumberOk() (*int32, bool)`

GetSeqNumberOk returns a tuple with the SeqNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqNumber

`func (o *StateJsonhal) SetSeqNumber(v int32)`

SetSeqNumber sets SeqNumber field to given value.

### HasSeqNumber

`func (o *StateJsonhal) HasSeqNumber() bool`

HasSeqNumber returns a boolean if a field has been set.

### SetSeqNumberNil

`func (o *StateJsonhal) SetSeqNumberNil(b bool)`

 SetSeqNumberNil sets the value for SeqNumber to be an explicit nil

### UnsetSeqNumber
`func (o *StateJsonhal) UnsetSeqNumber()`

UnsetSeqNumber ensures that no value is present for SeqNumber, not even an explicit nil
### GetProject

`func (o *StateJsonhal) GetProject() []string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *StateJsonhal) GetProjectOk() (*[]string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *StateJsonhal) SetProject(v []string)`

SetProject sets Project field to given value.

### HasProject

`func (o *StateJsonhal) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetEnvironment

`func (o *StateJsonhal) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *StateJsonhal) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *StateJsonhal) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *StateJsonhal) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetImportedData

`func (o *StateJsonhal) GetImportedData() []string`

GetImportedData returns the ImportedData field if non-nil, zero value otherwise.

### GetImportedDataOk

`func (o *StateJsonhal) GetImportedDataOk() (*[]string, bool)`

GetImportedDataOk returns a tuple with the ImportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedData

`func (o *StateJsonhal) SetImportedData(v []string)`

SetImportedData sets ImportedData field to given value.

### HasImportedData

`func (o *StateJsonhal) HasImportedData() bool`

HasImportedData returns a boolean if a field has been set.

### GetCreatedBy

`func (o *StateJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *StateJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *StateJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *StateJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *StateJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *StateJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *StateJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *StateJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *StateJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *StateJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *StateJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *StateJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *StateJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StateJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StateJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StateJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StateJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StateJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StateJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StateJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


