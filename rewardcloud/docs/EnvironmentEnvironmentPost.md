# EnvironmentEnvironmentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateIri** | Pointer to **NullableString** |  | [optional] 
**Project** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**ExportedData** | Pointer to **[]string** |  | [optional] 
**ImportedData** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**DataTransferSettings** | Pointer to **NullableString** |  | [optional] 
**IsStripDatabase** | Pointer to **NullableBool** |  | [optional] 
**IsAllowOutgoingEmails** | Pointer to **NullableBool** |  | [optional] 
**IsInitSampleData** | Pointer to **NullableBool** |  | [optional] 
**EnvVar** | Pointer to [**[]EnvironmentEnvVarEnvironmentPost**](EnvironmentEnvVarEnvironmentPost.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentEnvironmentPost**](EnvironmentComponentEnvironmentPost.md) |  | [optional] 

## Methods

### NewEnvironmentEnvironmentPost

`func NewEnvironmentEnvironmentPost() *EnvironmentEnvironmentPost`

NewEnvironmentEnvironmentPost instantiates a new EnvironmentEnvironmentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvironmentPostWithDefaults

`func NewEnvironmentEnvironmentPostWithDefaults() *EnvironmentEnvironmentPost`

NewEnvironmentEnvironmentPostWithDefaults instantiates a new EnvironmentEnvironmentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateIri

`func (o *EnvironmentEnvironmentPost) GetTemplateIri() string`

GetTemplateIri returns the TemplateIri field if non-nil, zero value otherwise.

### GetTemplateIriOk

`func (o *EnvironmentEnvironmentPost) GetTemplateIriOk() (*string, bool)`

GetTemplateIriOk returns a tuple with the TemplateIri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIri

`func (o *EnvironmentEnvironmentPost) SetTemplateIri(v string)`

SetTemplateIri sets TemplateIri field to given value.

### HasTemplateIri

`func (o *EnvironmentEnvironmentPost) HasTemplateIri() bool`

HasTemplateIri returns a boolean if a field has been set.

### SetTemplateIriNil

`func (o *EnvironmentEnvironmentPost) SetTemplateIriNil(b bool)`

 SetTemplateIriNil sets the value for TemplateIri to be an explicit nil

### UnsetTemplateIri
`func (o *EnvironmentEnvironmentPost) UnsetTemplateIri()`

UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
### GetProject

`func (o *EnvironmentEnvironmentPost) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentEnvironmentPost) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentEnvironmentPost) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentEnvironmentPost) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *EnvironmentEnvironmentPost) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *EnvironmentEnvironmentPost) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProvider

`func (o *EnvironmentEnvironmentPost) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnvironmentEnvironmentPost) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnvironmentEnvironmentPost) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EnvironmentEnvironmentPost) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *EnvironmentEnvironmentPost) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *EnvironmentEnvironmentPost) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetState

`func (o *EnvironmentEnvironmentPost) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentEnvironmentPost) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentEnvironmentPost) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentEnvironmentPost) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *EnvironmentEnvironmentPost) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *EnvironmentEnvironmentPost) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetRegion

`func (o *EnvironmentEnvironmentPost) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentEnvironmentPost) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentEnvironmentPost) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnvironmentEnvironmentPost) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *EnvironmentEnvironmentPost) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *EnvironmentEnvironmentPost) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetExportedData

`func (o *EnvironmentEnvironmentPost) GetExportedData() []string`

GetExportedData returns the ExportedData field if non-nil, zero value otherwise.

### GetExportedDataOk

`func (o *EnvironmentEnvironmentPost) GetExportedDataOk() (*[]string, bool)`

GetExportedDataOk returns a tuple with the ExportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedData

`func (o *EnvironmentEnvironmentPost) SetExportedData(v []string)`

SetExportedData sets ExportedData field to given value.

### HasExportedData

`func (o *EnvironmentEnvironmentPost) HasExportedData() bool`

HasExportedData returns a boolean if a field has been set.

### GetImportedData

`func (o *EnvironmentEnvironmentPost) GetImportedData() []string`

GetImportedData returns the ImportedData field if non-nil, zero value otherwise.

### GetImportedDataOk

`func (o *EnvironmentEnvironmentPost) GetImportedDataOk() (*[]string, bool)`

GetImportedDataOk returns a tuple with the ImportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedData

`func (o *EnvironmentEnvironmentPost) SetImportedData(v []string)`

SetImportedData sets ImportedData field to given value.

### HasImportedData

`func (o *EnvironmentEnvironmentPost) HasImportedData() bool`

HasImportedData returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentEnvironmentPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentEnvironmentPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentEnvironmentPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentEnvironmentPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentEnvironmentPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentEnvironmentPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *EnvironmentEnvironmentPost) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentEnvironmentPost) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentEnvironmentPost) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentEnvironmentPost) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentEnvironmentPost) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentEnvironmentPost) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentEnvironmentPost) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentEnvironmentPost) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentEnvironmentPost) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentEnvironmentPost) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentEnvironmentPost) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentEnvironmentPost) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentEnvironmentPost) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentEnvironmentPost) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentEnvironmentPost) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentEnvironmentPost) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentEnvironmentPost) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentEnvironmentPost) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *EnvironmentEnvironmentPost) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *EnvironmentEnvironmentPost) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *EnvironmentEnvironmentPost) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *EnvironmentEnvironmentPost) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *EnvironmentEnvironmentPost) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *EnvironmentEnvironmentPost) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *EnvironmentEnvironmentPost) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *EnvironmentEnvironmentPost) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *EnvironmentEnvironmentPost) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *EnvironmentEnvironmentPost) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *EnvironmentEnvironmentPost) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *EnvironmentEnvironmentPost) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *EnvironmentEnvironmentPost) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *EnvironmentEnvironmentPost) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *EnvironmentEnvironmentPost) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *EnvironmentEnvironmentPost) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *EnvironmentEnvironmentPost) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *EnvironmentEnvironmentPost) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *EnvironmentEnvironmentPost) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *EnvironmentEnvironmentPost) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *EnvironmentEnvironmentPost) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *EnvironmentEnvironmentPost) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *EnvironmentEnvironmentPost) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *EnvironmentEnvironmentPost) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *EnvironmentEnvironmentPost) GetEnvVar() []EnvironmentEnvVarEnvironmentPost`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *EnvironmentEnvironmentPost) GetEnvVarOk() (*[]EnvironmentEnvVarEnvironmentPost, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *EnvironmentEnvironmentPost) SetEnvVar(v []EnvironmentEnvVarEnvironmentPost)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *EnvironmentEnvironmentPost) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *EnvironmentEnvironmentPost) GetEnvironmentComponent() []EnvironmentComponentEnvironmentPost`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *EnvironmentEnvironmentPost) GetEnvironmentComponentOk() (*[]EnvironmentComponentEnvironmentPost, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *EnvironmentEnvironmentPost) SetEnvironmentComponent(v []EnvironmentComponentEnvironmentPost)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *EnvironmentEnvironmentPost) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


