# EnvironmentComponentJsonhalEnvironment

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

### NewEnvironmentComponentJsonhalEnvironment

`func NewEnvironmentComponentJsonhalEnvironment() *EnvironmentComponentJsonhalEnvironment`

NewEnvironmentComponentJsonhalEnvironment instantiates a new EnvironmentComponentJsonhalEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentComponentJsonhalEnvironmentWithDefaults

`func NewEnvironmentComponentJsonhalEnvironmentWithDefaults() *EnvironmentComponentJsonhalEnvironment`

NewEnvironmentComponentJsonhalEnvironmentWithDefaults instantiates a new EnvironmentComponentJsonhalEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentComponentJsonhalEnvironment) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentComponentJsonhalEnvironment) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentComponentJsonhalEnvironment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentComponentJsonhalEnvironment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentComponentJsonhalEnvironment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentComponentJsonhalEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentComponentJsonhalEnvironment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentComponentJsonhalEnvironment) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentComponentJsonhalEnvironment) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *EnvironmentComponentJsonhalEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentComponentJsonhalEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentComponentJsonhalEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCodeName

`func (o *EnvironmentComponentJsonhalEnvironment) GetCodeName() string`

GetCodeName returns the CodeName field if non-nil, zero value otherwise.

### GetCodeNameOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetCodeNameOk() (*string, bool)`

GetCodeNameOk returns a tuple with the CodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeName

`func (o *EnvironmentComponentJsonhalEnvironment) SetCodeName(v string)`

SetCodeName sets CodeName field to given value.

### HasCodeName

`func (o *EnvironmentComponentJsonhalEnvironment) HasCodeName() bool`

HasCodeName returns a boolean if a field has been set.

### SetCodeNameNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetCodeNameNil(b bool)`

 SetCodeNameNil sets the value for CodeName to be an explicit nil

### UnsetCodeName
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetCodeName()`

UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
### GetCpu

`func (o *EnvironmentComponentJsonhalEnvironment) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *EnvironmentComponentJsonhalEnvironment) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *EnvironmentComponentJsonhalEnvironment) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### SetCpuNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetCpuNil(b bool)`

 SetCpuNil sets the value for Cpu to be an explicit nil

### UnsetCpu
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetCpu()`

UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
### GetMemory

`func (o *EnvironmentComponentJsonhalEnvironment) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *EnvironmentComponentJsonhalEnvironment) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *EnvironmentComponentJsonhalEnvironment) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetStorage

`func (o *EnvironmentComponentJsonhalEnvironment) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *EnvironmentComponentJsonhalEnvironment) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *EnvironmentComponentJsonhalEnvironment) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### SetStorageNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetStorageNil(b bool)`

 SetStorageNil sets the value for Storage to be an explicit nil

### UnsetStorage
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetStorage()`

UnsetStorage ensures that no value is present for Storage, not even an explicit nil
### GetDataTransferSettings

`func (o *EnvironmentComponentJsonhalEnvironment) GetDataTransferSettings() string`

GetDataTransferSettings returns the DataTransferSettings field if non-nil, zero value otherwise.

### GetDataTransferSettingsOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetDataTransferSettingsOk() (*string, bool)`

GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTransferSettings

`func (o *EnvironmentComponentJsonhalEnvironment) SetDataTransferSettings(v string)`

SetDataTransferSettings sets DataTransferSettings field to given value.

### HasDataTransferSettings

`func (o *EnvironmentComponentJsonhalEnvironment) HasDataTransferSettings() bool`

HasDataTransferSettings returns a boolean if a field has been set.

### SetDataTransferSettingsNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetDataTransferSettingsNil(b bool)`

 SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil

### UnsetDataTransferSettings
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetDataTransferSettings()`

UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
### GetIsStripDatabase

`func (o *EnvironmentComponentJsonhalEnvironment) GetIsStripDatabase() bool`

GetIsStripDatabase returns the IsStripDatabase field if non-nil, zero value otherwise.

### GetIsStripDatabaseOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetIsStripDatabaseOk() (*bool, bool)`

GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStripDatabase

`func (o *EnvironmentComponentJsonhalEnvironment) SetIsStripDatabase(v bool)`

SetIsStripDatabase sets IsStripDatabase field to given value.

### HasIsStripDatabase

`func (o *EnvironmentComponentJsonhalEnvironment) HasIsStripDatabase() bool`

HasIsStripDatabase returns a boolean if a field has been set.

### SetIsStripDatabaseNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetIsStripDatabaseNil(b bool)`

 SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil

### UnsetIsStripDatabase
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetIsStripDatabase()`

UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
### GetIsAllowOutgoingEmails

`func (o *EnvironmentComponentJsonhalEnvironment) GetIsAllowOutgoingEmails() bool`

GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field if non-nil, zero value otherwise.

### GetIsAllowOutgoingEmailsOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetIsAllowOutgoingEmailsOk() (*bool, bool)`

GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowOutgoingEmails

`func (o *EnvironmentComponentJsonhalEnvironment) SetIsAllowOutgoingEmails(v bool)`

SetIsAllowOutgoingEmails sets IsAllowOutgoingEmails field to given value.

### HasIsAllowOutgoingEmails

`func (o *EnvironmentComponentJsonhalEnvironment) HasIsAllowOutgoingEmails() bool`

HasIsAllowOutgoingEmails returns a boolean if a field has been set.

### SetIsAllowOutgoingEmailsNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetIsAllowOutgoingEmailsNil(b bool)`

 SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil

### UnsetIsAllowOutgoingEmails
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetIsAllowOutgoingEmails()`

UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
### GetIsInitSampleData

`func (o *EnvironmentComponentJsonhalEnvironment) GetIsInitSampleData() bool`

GetIsInitSampleData returns the IsInitSampleData field if non-nil, zero value otherwise.

### GetIsInitSampleDataOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetIsInitSampleDataOk() (*bool, bool)`

GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitSampleData

`func (o *EnvironmentComponentJsonhalEnvironment) SetIsInitSampleData(v bool)`

SetIsInitSampleData sets IsInitSampleData field to given value.

### HasIsInitSampleData

`func (o *EnvironmentComponentJsonhalEnvironment) HasIsInitSampleData() bool`

HasIsInitSampleData returns a boolean if a field has been set.

### SetIsInitSampleDataNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetIsInitSampleDataNil(b bool)`

 SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil

### UnsetIsInitSampleData
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetIsInitSampleData()`

UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
### GetEnvVar

`func (o *EnvironmentComponentJsonhalEnvironment) GetEnvVar() []EnvironmentEnvVarJsonhal`

GetEnvVar returns the EnvVar field if non-nil, zero value otherwise.

### GetEnvVarOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetEnvVarOk() (*[]EnvironmentEnvVarJsonhal, bool)`

GetEnvVarOk returns a tuple with the EnvVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVar

`func (o *EnvironmentComponentJsonhalEnvironment) SetEnvVar(v []EnvironmentEnvVarJsonhal)`

SetEnvVar sets EnvVar field to given value.

### HasEnvVar

`func (o *EnvironmentComponentJsonhalEnvironment) HasEnvVar() bool`

HasEnvVar returns a boolean if a field has been set.

### GetEnvironmentComponent

`func (o *EnvironmentComponentJsonhalEnvironment) GetEnvironmentComponent() []EnvironmentComponentJsonhal`

GetEnvironmentComponent returns the EnvironmentComponent field if non-nil, zero value otherwise.

### GetEnvironmentComponentOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetEnvironmentComponentOk() (*[]EnvironmentComponentJsonhal, bool)`

GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentComponent

`func (o *EnvironmentComponentJsonhalEnvironment) SetEnvironmentComponent(v []EnvironmentComponentJsonhal)`

SetEnvironmentComponent sets EnvironmentComponent field to given value.

### HasEnvironmentComponent

`func (o *EnvironmentComponentJsonhalEnvironment) HasEnvironmentComponent() bool`

HasEnvironmentComponent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnvironmentComponentJsonhalEnvironment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentComponentJsonhalEnvironment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentComponentJsonhalEnvironment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentComponentJsonhalEnvironment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentComponentJsonhalEnvironment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentComponentJsonhalEnvironment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EnvironmentComponentJsonhalEnvironment) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentComponentJsonhalEnvironment) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentComponentJsonhalEnvironment) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentComponentJsonhalEnvironment) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentComponentJsonhalEnvironment) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentComponentJsonhalEnvironment) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentComponentJsonhalEnvironment) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentComponentJsonhalEnvironment) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentComponentJsonhalEnvironment) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


