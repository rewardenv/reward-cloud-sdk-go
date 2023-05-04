# TemplateEnvironmentTemplateEnvironmentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateProject** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Cpu** | Pointer to **NullableInt32** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Storage** | Pointer to **NullableInt32** |  | [optional] 
**DataTransferSettings** | Pointer to **NullableString** |  | [optional] 
**IsStripDatabase** | Pointer to **NullableBool** |  | [optional] 
**IsAllowOutgoingEmails** | Pointer to **NullableBool** |  | [optional] 
**IsInitSampleData** | Pointer to **NullableBool** |  | [optional] 
**EnvVar** | Pointer to [**[]EnvironmentEnvVarTemplateEnvironmentInput**](EnvironmentEnvVarTemplateEnvironmentInput.md) |  | [optional] 
**EnvironmentComponent** | Pointer to [**[]EnvironmentComponentTemplateEnvironmentInput**](EnvironmentComponentTemplateEnvironmentInput.md) |  | [optional] 

## Methods

### NewTemplateEnvironmentTemplateEnvironmentInput

`func NewTemplateEnvironmentTemplateEnvironmentInput() *TemplateEnvironmentTemplateEnvironmentInput`

NewTemplateEnvironmentTemplateEnvironmentInput instantiates a new TemplateEnvironmentTemplateEnvironmentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEnvironmentTemplateEnvironmentInputWithDefaults

`func NewTemplateEnvironmentTemplateEnvironmentInputWithDefaults() *TemplateEnvironmentTemplateEnvironmentInput`

NewTemplateEnvironmentTemplateEnvironmentInputWithDefaults instantiates a new TemplateEnvironmentTemplateEnvironmentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateProject

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetTemplateProject() string`

GetTemplateProject returns the TemplateProject field if non-nil, zero value otherwise.

### GetTemplateProjectOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetTemplateProjectOk() (*string, bool)`

GetTemplateProjectOk returns a tuple with the TemplateProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateProject

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetTemplateProject(v string)`

SetTemplateProject sets TemplateProject field to given value.

### HasTemplateProject

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasTemplateProject() bool`

HasTemplateProject returns a boolean if a field has been set.

### SetTemplateProjectNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetTemplateProjectNil(b bool)`

 SetTemplateProjectNil sets the value for TemplateProject to be an explicit nil

### UnsetTemplateProject
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetTemplateProject()`

UnsetTemplateProject ensures that no value is present for TemplateProject, not even an explicit nil
### GetName

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCpu

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *TemplateEnvironmentTemplateEnvironmentInput) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetEnvVar() []EnvironmentEnvVarTemplateEnvironmentInput`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetEnvVarOk() (*[]EnvironmentEnvVarTemplateEnvironmentInput, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetEnvVar(v []EnvironmentEnvVarTemplateEnvironmentInput)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetEnvironmentComponent() []EnvironmentComponentTemplateEnvironmentInput`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *TemplateEnvironmentTemplateEnvironmentInput) GetEnvironmentComponentOk() (*[]EnvironmentComponentTemplateEnvironmentInput, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *TemplateEnvironmentTemplateEnvironmentInput) SetEnvironmentComponent(v []EnvironmentComponentTemplateEnvironmentInput)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *TemplateEnvironmentTemplateEnvironmentInput) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


