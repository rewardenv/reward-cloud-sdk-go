# EnvironmentJsonhalEnvironmentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
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
**EnvVar** | Pointer to [**[]EnvironmentEnvVarJsonhalEnvironmentPost**](EnvironmentEnvVarJsonhalEnvironmentPost.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentJsonhalEnvironmentPost**](EnvironmentComponentJsonhalEnvironmentPost.md) |  | [optional] 

## Methods

### NewEnvironmentJsonhalEnvironmentPost

`func NewEnvironmentJsonhalEnvironmentPost() *EnvironmentJsonhalEnvironmentPost`

NewEnvironmentJsonhalEnvironmentPost instantiates a new EnvironmentJsonhalEnvironmentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentJsonhalEnvironmentPostWithDefaults

`func NewEnvironmentJsonhalEnvironmentPostWithDefaults() *EnvironmentJsonhalEnvironmentPost`

NewEnvironmentJsonhalEnvironmentPostWithDefaults instantiates a new EnvironmentJsonhalEnvironmentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentJsonhalEnvironmentPost) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentJsonhalEnvironmentPost) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentJsonhalEnvironmentPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTemplateIri

`func (o *EnvironmentJsonhalEnvironmentPost) GetTemplateIri() string`

GetTemplateIri returns the TemplateIri field if non-nil, zero value otherwise.

### GetTemplateIriOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetTemplateIriOk() (*string, bool)`

GetTemplateIriOk returns a tuple with the TemplateIri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIri

`func (o *EnvironmentJsonhalEnvironmentPost) SetTemplateIri(v string)`

SetTemplateIri sets TemplateIri field to given value.

### HasTemplateIri

`func (o *EnvironmentJsonhalEnvironmentPost) HasTemplateIri() bool`

HasTemplateIri returns a boolean if a field has been set.

### SetTemplateIriNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetTemplateIriNil(b bool)`

 SetTemplateIriNil sets the value for TemplateIri to be an explicit nil

### UnsetTemplateIri
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetTemplateIri()`

UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
### GetProject

`func (o *EnvironmentJsonhalEnvironmentPost) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentJsonhalEnvironmentPost) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentJsonhalEnvironmentPost) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetProvider

`func (o *EnvironmentJsonhalEnvironmentPost) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnvironmentJsonhalEnvironmentPost) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EnvironmentJsonhalEnvironmentPost) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetState

`func (o *EnvironmentJsonhalEnvironmentPost) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentJsonhalEnvironmentPost) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentJsonhalEnvironmentPost) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetRegion

`func (o *EnvironmentJsonhalEnvironmentPost) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentJsonhalEnvironmentPost) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EnvironmentJsonhalEnvironmentPost) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetExportedData

`func (o *EnvironmentJsonhalEnvironmentPost) GetExportedData() []string`

GetExportedData returns the ExportedData field if non-nil, zero value otherwise.

### GetExportedDataOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetExportedDataOk() (*[]string, bool)`

GetExportedDataOk returns a tuple with the ExportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedData

`func (o *EnvironmentJsonhalEnvironmentPost) SetExportedData(v []string)`

SetExportedData sets ExportedData field to given value.

### HasExportedData

`func (o *EnvironmentJsonhalEnvironmentPost) HasExportedData() bool`

HasExportedData returns a boolean if a field has been set.

### GetImportedData

`func (o *EnvironmentJsonhalEnvironmentPost) GetImportedData() []string`

GetImportedData returns the ImportedData field if non-nil, zero value otherwise.

### GetImportedDataOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetImportedDataOk() (*[]string, bool)`

GetImportedDataOk returns a tuple with the ImportedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedData

`func (o *EnvironmentJsonhalEnvironmentPost) SetImportedData(v []string)`

SetImportedData sets ImportedData field to given value.

### HasImportedData

`func (o *EnvironmentJsonhalEnvironmentPost) HasImportedData() bool`

HasImportedData returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentJsonhalEnvironmentPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentJsonhalEnvironmentPost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentJsonhalEnvironmentPost) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *EnvironmentJsonhalEnvironmentPost) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentJsonhalEnvironmentPost) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentJsonhalEnvironmentPost) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentJsonhalEnvironmentPost) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentJsonhalEnvironmentPost) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentJsonhalEnvironmentPost) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentJsonhalEnvironmentPost) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentJsonhalEnvironmentPost) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentJsonhalEnvironmentPost) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *EnvironmentJsonhalEnvironmentPost) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *EnvironmentJsonhalEnvironmentPost) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *EnvironmentJsonhalEnvironmentPost) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *EnvironmentJsonhalEnvironmentPost) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *EnvironmentJsonhalEnvironmentPost) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *EnvironmentJsonhalEnvironmentPost) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *EnvironmentJsonhalEnvironmentPost) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *EnvironmentJsonhalEnvironmentPost) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *EnvironmentJsonhalEnvironmentPost) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *EnvironmentJsonhalEnvironmentPost) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *EnvironmentJsonhalEnvironmentPost) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *EnvironmentJsonhalEnvironmentPost) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *EnvironmentJsonhalEnvironmentPost) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *EnvironmentJsonhalEnvironmentPost) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *EnvironmentJsonhalEnvironmentPost) GetEnvVar() []EnvironmentEnvVarJsonhalEnvironmentPost`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetEnvVarOk() (*[]EnvironmentEnvVarJsonhalEnvironmentPost, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *EnvironmentJsonhalEnvironmentPost) SetEnvVar(v []EnvironmentEnvVarJsonhalEnvironmentPost)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *EnvironmentJsonhalEnvironmentPost) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *EnvironmentJsonhalEnvironmentPost) GetEnvironmentComponent() []EnvironmentComponentJsonhalEnvironmentPost`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *EnvironmentJsonhalEnvironmentPost) GetEnvironmentComponentOk() (*[]EnvironmentComponentJsonhalEnvironmentPost, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *EnvironmentJsonhalEnvironmentPost) SetEnvironmentComponent(v []EnvironmentComponentJsonhalEnvironmentPost)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *EnvironmentJsonhalEnvironmentPost) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


