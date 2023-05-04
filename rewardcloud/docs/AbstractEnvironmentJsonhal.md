# AbstractEnvironmentJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CodeName** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**DataTransferSettings** | Pointer to **NullableString** |  | [optional] 
**IsStripDatabase** | Pointer to **NullableBool** |  | [optional] 
**IsAllowOutgoingEmails** | Pointer to **NullableBool** |  | [optional] 
**IsInitSampleData** | Pointer to **NullableBool** |  | [optional] 
**EnvVar** | Pointer to [**[]EnvironmentEnvVarJsonhal**](EnvironmentEnvVarJsonhal.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAbstractEnvironmentJsonhal

`func NewAbstractEnvironmentJsonhal() *AbstractEnvironmentJsonhal`

NewAbstractEnvironmentJsonhal instantiates a new AbstractEnvironmentJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractEnvironmentJsonhalWithDefaults

`func NewAbstractEnvironmentJsonhalWithDefaults() *AbstractEnvironmentJsonhal`

NewAbstractEnvironmentJsonhalWithDefaults instantiates a new AbstractEnvironmentJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AbstractEnvironmentJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AbstractEnvironmentJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AbstractEnvironmentJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AbstractEnvironmentJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AbstractEnvironmentJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractEnvironmentJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractEnvironmentJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractEnvironmentJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AbstractEnvironmentJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AbstractEnvironmentJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AbstractEnvironmentJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AbstractEnvironmentJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *AbstractEnvironmentJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *AbstractEnvironmentJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *AbstractEnvironmentJsonhal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractEnvironmentJsonhal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractEnvironmentJsonhal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractEnvironmentJsonhal) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AbstractEnvironmentJsonhal) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AbstractEnvironmentJsonhal) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *AbstractEnvironmentJsonhal) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *AbstractEnvironmentJsonhal) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *AbstractEnvironmentJsonhal) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *AbstractEnvironmentJsonhal) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *AbstractEnvironmentJsonhal) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *AbstractEnvironmentJsonhal) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetCpu

`func (o *AbstractEnvironmentJsonhal) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AbstractEnvironmentJsonhal) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AbstractEnvironmentJsonhal) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AbstractEnvironmentJsonhal) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *AbstractEnvironmentJsonhal) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *AbstractEnvironmentJsonhal) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *AbstractEnvironmentJsonhal) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AbstractEnvironmentJsonhal) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AbstractEnvironmentJsonhal) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AbstractEnvironmentJsonhal) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *AbstractEnvironmentJsonhal) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *AbstractEnvironmentJsonhal) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *AbstractEnvironmentJsonhal) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AbstractEnvironmentJsonhal) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AbstractEnvironmentJsonhal) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *AbstractEnvironmentJsonhal) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *AbstractEnvironmentJsonhal) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *AbstractEnvironmentJsonhal) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *AbstractEnvironmentJsonhal) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *AbstractEnvironmentJsonhal) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *AbstractEnvironmentJsonhal) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *AbstractEnvironmentJsonhal) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *AbstractEnvironmentJsonhal) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *AbstractEnvironmentJsonhal) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *AbstractEnvironmentJsonhal) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *AbstractEnvironmentJsonhal) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *AbstractEnvironmentJsonhal) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *AbstractEnvironmentJsonhal) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *AbstractEnvironmentJsonhal) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *AbstractEnvironmentJsonhal) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *AbstractEnvironmentJsonhal) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *AbstractEnvironmentJsonhal) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *AbstractEnvironmentJsonhal) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *AbstractEnvironmentJsonhal) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *AbstractEnvironmentJsonhal) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *AbstractEnvironmentJsonhal) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *AbstractEnvironmentJsonhal) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *AbstractEnvironmentJsonhal) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *AbstractEnvironmentJsonhal) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *AbstractEnvironmentJsonhal) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *AbstractEnvironmentJsonhal) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *AbstractEnvironmentJsonhal) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *AbstractEnvironmentJsonhal) GetEnvVar() []EnvironmentEnvVarJsonhal`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *AbstractEnvironmentJsonhal) GetEnvVarOk() (*[]EnvironmentEnvVarJsonhal, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *AbstractEnvironmentJsonhal) SetEnvVar(v []EnvironmentEnvVarJsonhal)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *AbstractEnvironmentJsonhal) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *AbstractEnvironmentJsonhal) GetEnvironmentComponent() []EnvironmentComponentJsonhal`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *AbstractEnvironmentJsonhal) GetEnvironmentComponentOk() (*[]EnvironmentComponentJsonhal, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *AbstractEnvironmentJsonhal) SetEnvironmentComponent(v []EnvironmentComponentJsonhal)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *AbstractEnvironmentJsonhal) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AbstractEnvironmentJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AbstractEnvironmentJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AbstractEnvironmentJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AbstractEnvironmentJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AbstractEnvironmentJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AbstractEnvironmentJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AbstractEnvironmentJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AbstractEnvironmentJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AbstractEnvironmentJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AbstractEnvironmentJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AbstractEnvironmentJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AbstractEnvironmentJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AbstractEnvironmentJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AbstractEnvironmentJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *AbstractEnvironmentJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AbstractEnvironmentJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AbstractEnvironmentJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AbstractEnvironmentJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *AbstractEnvironmentJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *AbstractEnvironmentJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


