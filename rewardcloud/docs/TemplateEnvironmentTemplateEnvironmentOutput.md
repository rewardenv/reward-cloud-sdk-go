# TemplateEnvironmentTemplateEnvironmentOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateProject** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**DataTransferSettings** | Pointer to **NullableString** |  | [optional] 
**IsStripDatabase** | Pointer to **NullableBool** |  | [optional] 
**IsAllowOutgoingEmails** | Pointer to **NullableBool** |  | [optional] 
**IsInitSampleData** | Pointer to **NullableBool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**EnvVar** | Pointer to [**[]EnvironmentEnvVarTemplateEnvironmentOutput**](EnvironmentEnvVarTemplateEnvironmentOutput.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentTemplateEnvironmentOutput**](EnvironmentComponentTemplateEnvironmentOutput.md) |  | [optional] 

## Methods

### NewTemplateEnvironmentTemplateEnvironmentOutput

`func NewTemplateEnvironmentTemplateEnvironmentOutput() *TemplateEnvironmentTemplateEnvironmentOutput`

NewTemplateEnvironmentTemplateEnvironmentOutput instantiates a new TemplateEnvironmentTemplateEnvironmentOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEnvironmentTemplateEnvironmentOutputWithDefaults

`func NewTemplateEnvironmentTemplateEnvironmentOutputWithDefaults() *TemplateEnvironmentTemplateEnvironmentOutput`

NewTemplateEnvironmentTemplateEnvironmentOutputWithDefaults instantiates a new TemplateEnvironmentTemplateEnvironmentOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateProject

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetTemplateProject() string`

GetTemplateProject returns the TemplateProject field if non-nil, zero value otherwise.

### GetTemplateProjectOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetTemplateProjectOk() (*string, bool)`

GetTemplateProjectOk returns a tuple with the TemplateProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProject

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetTemplateProject(v string)`

SetTemplateProject sets TemplateProject field to given value.

### HasTemplateProject

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasTemplateProject() bool`

HasTemplateProject returns a boolean if a field has been set.

### SetTemplateProjectNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetTemplateProjectNil(b bool)`

 SetTemplateProjectNil sets the value for TemplateProject to be an explicit nil

### UnsetTemplateProject
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetTemplateProject()`

UnsetTemplateProject ensures that no value is present for TemplateProject, not even an explicit nil
### GetId

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetCreatedAt

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvVar

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvVar() []EnvironmentEnvVarTemplateEnvironmentOutput`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvVarOk() (*[]EnvironmentEnvVarTemplateEnvironmentOutput, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetEnvVar(v []EnvironmentEnvVarTemplateEnvironmentOutput)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvironmentComponent() []EnvironmentComponentTemplateEnvironmentOutput`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvironmentComponentOk() (*[]EnvironmentComponentTemplateEnvironmentOutput, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetEnvironmentComponent(v []EnvironmentComponentTemplateEnvironmentOutput)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


