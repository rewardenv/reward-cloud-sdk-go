# EnvironmentEnvironmentGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateIri** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**EnvironmentAccess** | Pointer to **NullableString** |  | [optional] 
**ExportedData** | Pointer to **[]string** |  | [optional] 
**ImportedData** | Pointer to **[]string** |  | [optional] 
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
**EnvVar** | Pointer to [**[]EnvironmentEnvVarEnvironmentGet**](EnvironmentEnvVarEnvironmentGet.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentEnvironmentGet**](EnvironmentComponentEnvironmentGet.md) |  | [optional] 

## Methods

### NewEnvironmentEnvironmentGet

`func NewEnvironmentEnvironmentGet() *EnvironmentEnvironmentGet`

NewEnvironmentEnvironmentGet instantiates a new EnvironmentEnvironmentGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvironmentGetWithDefaults

`func NewEnvironmentEnvironmentGetWithDefaults() *EnvironmentEnvironmentGet`

NewEnvironmentEnvironmentGetWithDefaults instantiates a new EnvironmentEnvironmentGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateIri

`func (o *EnvironmentEnvironmentGet) GetTemplateIri() string`

GetTemplateIri returns the TemplateIri field if non-nil, zero value otherwise.

### GetTemplateIriOk

`func (o *EnvironmentEnvironmentGet) GetTemplateIriOk() (*string, bool)`

GetTemplateIriOk returns a tuple with the TemplateIri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIri

`func (o *EnvironmentEnvironmentGet) SetTemplateIri(v string)`

SetTemplateIri sets TemplateIri field to given value.

### HasTemplateIri

`func (o *EnvironmentEnvironmentGet) HasTemplateIri() bool`

HasTemplateIri returns a boolean if a field has been set.

### SetTemplateIriNil

`func (o *EnvironmentEnvironmentGet) SetTemplateIriNil(b bool)`

 SetTemplateIriNil sets the value for TemplateIri to be an explicit nil

### UnsetTemplateIri
`func (o *EnvironmentEnvironmentGet) UnsetTemplateIri()`

UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
### GetProject

`func (o *EnvironmentEnvironmentGet) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentEnvironmentGet) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentEnvironmentGet) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentEnvironmentGet) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *EnvironmentEnvironmentGet) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *EnvironmentEnvironmentGet) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProvider

`func (o *EnvironmentEnvironmentGet) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnvironmentEnvironmentGet) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnvironmentEnvironmentGet) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EnvironmentEnvironmentGet) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *EnvironmentEnvironmentGet) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *EnvironmentEnvironmentGet) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetState

`func (o *EnvironmentEnvironmentGet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentEnvironmentGet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentEnvironmentGet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentEnvironmentGet) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *EnvironmentEnvironmentGet) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *EnvironmentEnvironmentGet) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetRegion

`func (o *EnvironmentEnvironmentGet) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentEnvironmentGet) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentEnvironmentGet) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnvironmentEnvironmentGet) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *EnvironmentEnvironmentGet) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *EnvironmentEnvironmentGet) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetEnvironmentAccess

`func (o *EnvironmentEnvironmentGet) GetEnvironmentAccess() string`

GetEnvironmentAccess returns the EnvironmentAccess field if non-nil, zero value otherwise.

### GetEnvironmentAccessOk

`func (o *EnvironmentEnvironmentGet) GetEnvironmentAccessOk() (*string, bool)`

GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccess

`func (o *EnvironmentEnvironmentGet) SetEnvironmentAccess(v string)`

SetEnvironmentAccess sets EnvironmentAccess field to given value.

### HasEnvironmentAccess

`func (o *EnvironmentEnvironmentGet) HasEnvironmentAccess() bool`

HasEnvironmentAccess returns a boolean if a field has been set.

### SetEnvironmentAccessNil

`func (o *EnvironmentEnvironmentGet) SetEnvironmentAccessNil(b bool)`

 SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil

### UnsetEnvironmentAccess
`func (o *EnvironmentEnvironmentGet) UnsetEnvironmentAccess()`

UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
### GetExportedData

`func (o *EnvironmentEnvironmentGet) GetExportedData() []string`

GetExportedData returns the ExportedData field if non-nil, zero value otherwise.

### GetExportedDataOk

`func (o *EnvironmentEnvironmentGet) GetExportedDataOk() (*[]string, bool)`

GetExportedDataOk returns a tuple with the ExportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedData

`func (o *EnvironmentEnvironmentGet) SetExportedData(v []string)`

SetExportedData sets ExportedData field to given value.

### HasExportedData

`func (o *EnvironmentEnvironmentGet) HasExportedData() bool`

HasExportedData returns a boolean if a field has been set.

### GetImportedData

`func (o *EnvironmentEnvironmentGet) GetImportedData() []string`

GetImportedData returns the ImportedData field if non-nil, zero value otherwise.

### GetImportedDataOk

`func (o *EnvironmentEnvironmentGet) GetImportedDataOk() (*[]string, bool)`

GetImportedDataOk returns a tuple with the ImportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedData

`func (o *EnvironmentEnvironmentGet) SetImportedData(v []string)`

SetImportedData sets ImportedData field to given value.

### HasImportedData

`func (o *EnvironmentEnvironmentGet) HasImportedData() bool`

HasImportedData returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentEnvironmentGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentEnvironmentGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentEnvironmentGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentEnvironmentGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentEnvironmentGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentEnvironmentGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentEnvironmentGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentEnvironmentGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentEnvironmentGet) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentEnvironmentGet) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *EnvironmentEnvironmentGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentEnvironmentGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentEnvironmentGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentEnvironmentGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentEnvironmentGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentEnvironmentGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *EnvironmentEnvironmentGet) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentEnvironmentGet) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentEnvironmentGet) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentEnvironmentGet) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentEnvironmentGet) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentEnvironmentGet) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentEnvironmentGet) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentEnvironmentGet) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentEnvironmentGet) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentEnvironmentGet) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentEnvironmentGet) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentEnvironmentGet) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentEnvironmentGet) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentEnvironmentGet) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentEnvironmentGet) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentEnvironmentGet) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentEnvironmentGet) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentEnvironmentGet) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *EnvironmentEnvironmentGet) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *EnvironmentEnvironmentGet) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *EnvironmentEnvironmentGet) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *EnvironmentEnvironmentGet) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *EnvironmentEnvironmentGet) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *EnvironmentEnvironmentGet) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *EnvironmentEnvironmentGet) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *EnvironmentEnvironmentGet) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *EnvironmentEnvironmentGet) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *EnvironmentEnvironmentGet) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *EnvironmentEnvironmentGet) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *EnvironmentEnvironmentGet) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *EnvironmentEnvironmentGet) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *EnvironmentEnvironmentGet) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *EnvironmentEnvironmentGet) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *EnvironmentEnvironmentGet) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *EnvironmentEnvironmentGet) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *EnvironmentEnvironmentGet) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *EnvironmentEnvironmentGet) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *EnvironmentEnvironmentGet) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *EnvironmentEnvironmentGet) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *EnvironmentEnvironmentGet) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *EnvironmentEnvironmentGet) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *EnvironmentEnvironmentGet) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentEnvironmentGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentEnvironmentGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentEnvironmentGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentEnvironmentGet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentEnvironmentGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentEnvironmentGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentEnvironmentGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentEnvironmentGet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvVar

`func (o *EnvironmentEnvironmentGet) GetEnvVar() []EnvironmentEnvVarEnvironmentGet`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *EnvironmentEnvironmentGet) GetEnvVarOk() (*[]EnvironmentEnvVarEnvironmentGet, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *EnvironmentEnvironmentGet) SetEnvVar(v []EnvironmentEnvVarEnvironmentGet)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *EnvironmentEnvironmentGet) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *EnvironmentEnvironmentGet) GetEnvironmentComponent() []EnvironmentComponentEnvironmentGet`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *EnvironmentEnvironmentGet) GetEnvironmentComponentOk() (*[]EnvironmentComponentEnvironmentGet, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *EnvironmentEnvironmentGet) SetEnvironmentComponent(v []EnvironmentComponentEnvironmentGet)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *EnvironmentEnvironmentGet) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


