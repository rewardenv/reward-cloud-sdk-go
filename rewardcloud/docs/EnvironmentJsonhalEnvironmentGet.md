# EnvironmentJsonhalEnvironmentGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**ComponentJsonhalLinks**](ComponentJsonhalLinks.md) |  | [optional] 
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
**Project** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**EnvVar** | Pointer to [**[]EnvironmentEnvVarJsonhalEnvironmentGet**](EnvironmentEnvVarJsonhalEnvironmentGet.md) |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentJsonhalEnvironmentGet**](EnvironmentComponentJsonhalEnvironmentGet.md) |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**ExportedData** | Pointer to **[]string** |  | [optional] 
**EnvironmentAccess** | Pointer to **NullableString** |  | [optional] 
**ImportedData** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnvironmentJsonhalEnvironmentGet

`func NewEnvironmentJsonhalEnvironmentGet() *EnvironmentJsonhalEnvironmentGet`

NewEnvironmentJsonhalEnvironmentGet instantiates a new EnvironmentJsonhalEnvironmentGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentJsonhalEnvironmentGetWithDefaults

`func NewEnvironmentJsonhalEnvironmentGetWithDefaults() *EnvironmentJsonhalEnvironmentGet`

NewEnvironmentJsonhalEnvironmentGetWithDefaults instantiates a new EnvironmentJsonhalEnvironmentGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentJsonhalEnvironmentGet) GetLinks() ComponentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetLinksOk() (*ComponentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentJsonhalEnvironmentGet) SetLinks(v ComponentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentJsonhalEnvironmentGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentJsonhalEnvironmentGet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentJsonhalEnvironmentGet) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentJsonhalEnvironmentGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentJsonhalEnvironmentGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentJsonhalEnvironmentGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentJsonhalEnvironmentGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *EnvironmentJsonhalEnvironmentGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentJsonhalEnvironmentGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentJsonhalEnvironmentGet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *EnvironmentJsonhalEnvironmentGet) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentJsonhalEnvironmentGet) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentJsonhalEnvironmentGet) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentJsonhalEnvironmentGet) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentJsonhalEnvironmentGet) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentJsonhalEnvironmentGet) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentJsonhalEnvironmentGet) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentJsonhalEnvironmentGet) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentJsonhalEnvironmentGet) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *EnvironmentJsonhalEnvironmentGet) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *EnvironmentJsonhalEnvironmentGet) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *EnvironmentJsonhalEnvironmentGet) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *EnvironmentJsonhalEnvironmentGet) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *EnvironmentJsonhalEnvironmentGet) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *EnvironmentJsonhalEnvironmentGet) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *EnvironmentJsonhalEnvironmentGet) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *EnvironmentJsonhalEnvironmentGet) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *EnvironmentJsonhalEnvironmentGet) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *EnvironmentJsonhalEnvironmentGet) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *EnvironmentJsonhalEnvironmentGet) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *EnvironmentJsonhalEnvironmentGet) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetProject

`func (o *EnvironmentJsonhalEnvironmentGet) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentJsonhalEnvironmentGet) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentJsonhalEnvironmentGet) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProvider

`func (o *EnvironmentJsonhalEnvironmentGet) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnvironmentJsonhalEnvironmentGet) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EnvironmentJsonhalEnvironmentGet) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetEnvVar

`func (o *EnvironmentJsonhalEnvironmentGet) GetEnvVar() []EnvironmentEnvVarJsonhalEnvironmentGet`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetEnvVarOk() (*[]EnvironmentEnvVarJsonhalEnvironmentGet, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *EnvironmentJsonhalEnvironmentGet) SetEnvVar(v []EnvironmentEnvVarJsonhalEnvironmentGet)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *EnvironmentJsonhalEnvironmentGet) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetState

`func (o *EnvironmentJsonhalEnvironmentGet) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentJsonhalEnvironmentGet) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentJsonhalEnvironmentGet) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetEnvironmentComponent

`func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentComponent() []EnvironmentComponentJsonhalEnvironmentGet`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentComponentOk() (*[]EnvironmentComponentJsonhalEnvironmentGet, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *EnvironmentJsonhalEnvironmentGet) SetEnvironmentComponent(v []EnvironmentComponentJsonhalEnvironmentGet)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *EnvironmentJsonhalEnvironmentGet) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.

### GetRegion

`func (o *EnvironmentJsonhalEnvironmentGet) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentJsonhalEnvironmentGet) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnvironmentJsonhalEnvironmentGet) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetExportedData

`func (o *EnvironmentJsonhalEnvironmentGet) GetExportedData() []string`

GetExportedData returns the ExportedData field if non-nil, zero value otherwise.

### GetExportedDataOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetExportedDataOk() (*[]string, bool)`

GetExportedDataOk returns a tuple with the ExportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedData

`func (o *EnvironmentJsonhalEnvironmentGet) SetExportedData(v []string)`

SetExportedData sets ExportedData field to given value.

### HasExportedData

`func (o *EnvironmentJsonhalEnvironmentGet) HasExportedData() bool`

HasExportedData returns a boolean if a field has been set.

### GetEnvironmentAccess

`func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentAccess() string`

GetEnvironmentAccess returns the EnvironmentAccess field if non-nil, zero value otherwise.

### GetEnvironmentAccessOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentAccessOk() (*string, bool)`

GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentAccess

`func (o *EnvironmentJsonhalEnvironmentGet) SetEnvironmentAccess(v string)`

SetEnvironmentAccess sets EnvironmentAccess field to given value.

### HasEnvironmentAccess

`func (o *EnvironmentJsonhalEnvironmentGet) HasEnvironmentAccess() bool`

HasEnvironmentAccess returns a boolean if a field has been set.

### SetEnvironmentAccessNil

`func (o *EnvironmentJsonhalEnvironmentGet) SetEnvironmentAccessNil(b bool)`

 SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil

### UnsetEnvironmentAccess
`func (o *EnvironmentJsonhalEnvironmentGet) UnsetEnvironmentAccess()`

UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
### GetImportedData

`func (o *EnvironmentJsonhalEnvironmentGet) GetImportedData() []string`

GetImportedData returns the ImportedData field if non-nil, zero value otherwise.

### GetImportedDataOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetImportedDataOk() (*[]string, bool)`

GetImportedDataOk returns a tuple with the ImportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedData

`func (o *EnvironmentJsonhalEnvironmentGet) SetImportedData(v []string)`

SetImportedData sets ImportedData field to given value.

### HasImportedData

`func (o *EnvironmentJsonhalEnvironmentGet) HasImportedData() bool`

HasImportedData returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnvironmentJsonhalEnvironmentGet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentJsonhalEnvironmentGet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentJsonhalEnvironmentGet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentJsonhalEnvironmentGet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentJsonhalEnvironmentGet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentJsonhalEnvironmentGet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentJsonhalEnvironmentGet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


