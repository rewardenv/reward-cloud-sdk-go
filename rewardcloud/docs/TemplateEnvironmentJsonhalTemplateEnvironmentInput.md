# TemplateEnvironmentJsonhalTemplateEnvironmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**TemplateProject** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**DataTransferSettings** | Pointer to **NullableString** |  | [optional] 
**IsStripDatabase** | Pointer to **NullableBool** |  | [optional] 
**IsAllowOutgoingEmails** | Pointer to **NullableBool** |  | [optional] 
**IsInitSampleData** | Pointer to **NullableBool** |  | [optional] 
**EnvVar** | Pointer to [**[]EnvironmentEnvVarJsonhalTemplateEnvironmentInput**](EnvironmentEnvVarJsonhalTemplateEnvironmentInput.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentJsonhalTemplateEnvironmentInput**](EnvironmentComponentJsonhalTemplateEnvironmentInput.md) |  | [optional] 

## Methods

### NewTemplateEnvironmentJsonhalTemplateEnvironmentInput

`func NewTemplateEnvironmentJsonhalTemplateEnvironmentInput() *TemplateEnvironmentJsonhalTemplateEnvironmentInput`

NewTemplateEnvironmentJsonhalTemplateEnvironmentInput instantiates a new TemplateEnvironmentJsonhalTemplateEnvironmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEnvironmentJsonhalTemplateEnvironmentInputWithDefaults

`func NewTemplateEnvironmentJsonhalTemplateEnvironmentInputWithDefaults() *TemplateEnvironmentJsonhalTemplateEnvironmentInput`

NewTemplateEnvironmentJsonhalTemplateEnvironmentInputWithDefaults instantiates a new TemplateEnvironmentJsonhalTemplateEnvironmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTemplateProject

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetTemplateProject() string`

GetTemplateProject returns the TemplateProject field if non-nil, zero value otherwise.

### GetTemplateProjectOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetTemplateProjectOk() (*string, bool)`

GetTemplateProjectOk returns a tuple with the TemplateProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProject

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetTemplateProject(v string)`

SetTemplateProject sets TemplateProject field to given value.

### HasTemplateProject

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasTemplateProject() bool`

HasTemplateProject returns a boolean if a field has been set.

### SetTemplateProjectNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetTemplateProjectNil(b bool)`

 SetTemplateProjectNil sets the value for TemplateProject to be an explicit nil

### UnsetTemplateProject
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetTemplateProject()`

UnsetTemplateProject ensures that no value is present for TemplateProject, not even an explicit nil
### GetName

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetEnvVar() []EnvironmentEnvVarJsonhalTemplateEnvironmentInput`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetEnvVarOk() (*[]EnvironmentEnvVarJsonhalTemplateEnvironmentInput, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetEnvVar(v []EnvironmentEnvVarJsonhalTemplateEnvironmentInput)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetEnvironmentComponent() []EnvironmentComponentJsonhalTemplateEnvironmentInput`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) GetEnvironmentComponentOk() (*[]EnvironmentComponentJsonhalTemplateEnvironmentInput, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) SetEnvironmentComponent(v []EnvironmentComponentJsonhalTemplateEnvironmentInput)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *TemplateEnvironmentJsonhalTemplateEnvironmentInput) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


