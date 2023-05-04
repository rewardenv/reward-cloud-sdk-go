# AbstractEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**EnvVar** | Pointer to [**[]EnvironmentEnvVar**](EnvironmentEnvVar.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponent**](EnvironmentComponent.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAbstractEnvironment

`func NewAbstractEnvironment() *AbstractEnvironment`

NewAbstractEnvironment instantiates a new AbstractEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractEnvironmentWithDefaults

`func NewAbstractEnvironmentWithDefaults() *AbstractEnvironment`

NewAbstractEnvironmentWithDefaults instantiates a new AbstractEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AbstractEnvironment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractEnvironment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractEnvironment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AbstractEnvironment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AbstractEnvironment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AbstractEnvironment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AbstractEnvironment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *AbstractEnvironment) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *AbstractEnvironment) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *AbstractEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AbstractEnvironment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AbstractEnvironment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *AbstractEnvironment) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *AbstractEnvironment) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *AbstractEnvironment) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *AbstractEnvironment) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *AbstractEnvironment) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *AbstractEnvironment) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetCpu

`func (o *AbstractEnvironment) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *AbstractEnvironment) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *AbstractEnvironment) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *AbstractEnvironment) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *AbstractEnvironment) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *AbstractEnvironment) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *AbstractEnvironment) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *AbstractEnvironment) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *AbstractEnvironment) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *AbstractEnvironment) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *AbstractEnvironment) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *AbstractEnvironment) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *AbstractEnvironment) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AbstractEnvironment) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AbstractEnvironment) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *AbstractEnvironment) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *AbstractEnvironment) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *AbstractEnvironment) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *AbstractEnvironment) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *AbstractEnvironment) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *AbstractEnvironment) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *AbstractEnvironment) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *AbstractEnvironment) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *AbstractEnvironment) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *AbstractEnvironment) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *AbstractEnvironment) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *AbstractEnvironment) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *AbstractEnvironment) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *AbstractEnvironment) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *AbstractEnvironment) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *AbstractEnvironment) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *AbstractEnvironment) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *AbstractEnvironment) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *AbstractEnvironment) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *AbstractEnvironment) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *AbstractEnvironment) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *AbstractEnvironment) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *AbstractEnvironment) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *AbstractEnvironment) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *AbstractEnvironment) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *AbstractEnvironment) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *AbstractEnvironment) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *AbstractEnvironment) GetEnvVar() []EnvironmentEnvVar`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *AbstractEnvironment) GetEnvVarOk() (*[]EnvironmentEnvVar, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *AbstractEnvironment) SetEnvVar(v []EnvironmentEnvVar)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *AbstractEnvironment) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *AbstractEnvironment) GetEnvironmentComponent() []EnvironmentComponent`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *AbstractEnvironment) GetEnvironmentComponentOk() (*[]EnvironmentComponent, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *AbstractEnvironment) SetEnvironmentComponent(v []EnvironmentComponent)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *AbstractEnvironment) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AbstractEnvironment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AbstractEnvironment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AbstractEnvironment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AbstractEnvironment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AbstractEnvironment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AbstractEnvironment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AbstractEnvironment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AbstractEnvironment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AbstractEnvironment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AbstractEnvironment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AbstractEnvironment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AbstractEnvironment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *AbstractEnvironment) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *AbstractEnvironment) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *AbstractEnvironment) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AbstractEnvironment) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AbstractEnvironment) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AbstractEnvironment) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *AbstractEnvironment) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *AbstractEnvironment) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


