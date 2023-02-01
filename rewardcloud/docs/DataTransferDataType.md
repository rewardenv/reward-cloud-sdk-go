# DataTransferDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ExportedData** | Pointer to **[]string** |  | [optional] 
**ImportedData** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDataTransferDataType

`func NewDataTransferDataType() *DataTransferDataType`

NewDataTransferDataType instantiates a new DataTransferDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataTransferDataTypeWithDefaults

`func NewDataTransferDataTypeWithDefaults() *DataTransferDataType`

NewDataTransferDataTypeWithDefaults instantiates a new DataTransferDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataTransferDataType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataTransferDataType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataTransferDataType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DataTransferDataType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataTransferDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataTransferDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataTransferDataType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataTransferDataType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataTransferDataType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataTransferDataType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExportedData

`func (o *DataTransferDataType) GetExportedData() []string`

GetExportedData returns the ExportedData field if non-nil, zero value otherwise.

### GetExportedDataOk

`func (o *DataTransferDataType) GetExportedDataOk() (*[]string, bool)`

GetExportedDataOk returns a tuple with the ExportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedData

`func (o *DataTransferDataType) SetExportedData(v []string)`

SetExportedData sets ExportedData field to given value.

### HasExportedData

`func (o *DataTransferDataType) HasExportedData() bool`

HasExportedData returns a boolean if a field has been set.

### GetImportedData

`func (o *DataTransferDataType) GetImportedData() []string`

GetImportedData returns the ImportedData field if non-nil, zero value otherwise.

### GetImportedDataOk

`func (o *DataTransferDataType) GetImportedDataOk() (*[]string, bool)`

GetImportedDataOk returns a tuple with the ImportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedData

`func (o *DataTransferDataType) SetImportedData(v []string)`

SetImportedData sets ImportedData field to given value.

### HasImportedData

`func (o *DataTransferDataType) HasImportedData() bool`

HasImportedData returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DataTransferDataType) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DataTransferDataType) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DataTransferDataType) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DataTransferDataType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *DataTransferDataType) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *DataTransferDataType) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *DataTransferDataType) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DataTransferDataType) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DataTransferDataType) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DataTransferDataType) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *DataTransferDataType) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *DataTransferDataType) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *DataTransferDataType) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataTransferDataType) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataTransferDataType) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataTransferDataType) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataTransferDataType) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataTransferDataType) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataTransferDataType) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataTransferDataType) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


