# TemplateEnvironmentJsonhalTemplateEnvironmentOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
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
**EnvVar** | Pointer to [**[]EnvironmentEnvVarJsonhalTemplateEnvironmentOutput**](EnvironmentEnvVarJsonhalTemplateEnvironmentOutput.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentJsonhalTemplateEnvironmentOutput**](EnvironmentComponentJsonhalTemplateEnvironmentOutput.md) |  | [optional] 

## Methods

### NewTemplateEnvironmentJsonhalTemplateEnvironmentOutput

`func NewTemplateEnvironmentJsonhalTemplateEnvironmentOutput() *TemplateEnvironmentJsonhalTemplateEnvironmentOutput`

NewTemplateEnvironmentJsonhalTemplateEnvironmentOutput instantiates a new TemplateEnvironmentJsonhalTemplateEnvironmentOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEnvironmentJsonhalTemplateEnvironmentOutputWithDefaults

`func NewTemplateEnvironmentJsonhalTemplateEnvironmentOutputWithDefaults() *TemplateEnvironmentJsonhalTemplateEnvironmentOutput`

NewTemplateEnvironmentJsonhalTemplateEnvironmentOutputWithDefaults instantiates a new TemplateEnvironmentJsonhalTemplateEnvironmentOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTemplateProject

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetTemplateProject() string`

GetTemplateProject returns the TemplateProject field if non-nil, zero value otherwise.

### GetTemplateProjectOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetTemplateProjectOk() (*string, bool)`

GetTemplateProjectOk returns a tuple with the TemplateProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProject

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetTemplateProject(v string)`

SetTemplateProject sets TemplateProject field to given value.

### HasTemplateProject

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasTemplateProject() bool`

HasTemplateProject returns a boolean if a field has been set.

### SetTemplateProjectNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetTemplateProjectNil(b bool)`

 SetTemplateProjectNil sets the value for TemplateProject to be an explicit nil

### UnsetTemplateProject
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetTemplateProject()`

UnsetTemplateProject ensures that no value is present for TemplateProject, not even an explicit nil
### GetId

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetCreatedAt

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvVar

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvVar() []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvVarOk() (*[]EnvironmentEnvVarJsonhalTemplateEnvironmentOutput, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetEnvVar(v []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvironmentComponent() []EnvironmentComponentJsonhalTemplateEnvironmentOutput`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvironmentComponentOk() (*[]EnvironmentComponentJsonhalTemplateEnvironmentOutput, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetEnvironmentComponent(v []EnvironmentComponentJsonhalTemplateEnvironmentOutput)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


