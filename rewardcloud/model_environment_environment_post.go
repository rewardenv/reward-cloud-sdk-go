/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// EnvironmentEnvironmentPost Class Environment
type EnvironmentEnvironmentPost struct {
	Name NullableString `json:"name,omitempty"`
	Cpu NullableInt32 `json:"cpu,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Storage NullableInt32 `json:"storage,omitempty"`
	DataTransferSettings NullableString `json:"dataTransferSettings,omitempty"`
	IsStripDatabase NullableBool `json:"isStripDatabase,omitempty"`
	IsAllowOutgoingEmails NullableBool `json:"isAllowOutgoingEmails,omitempty"`
	IsInitSampleData NullableBool `json:"isInitSampleData,omitempty"`
	Project NullableString `json:"project,omitempty"`
	Provider NullableString `json:"provider,omitempty"`
	EnvVar []EnvironmentEnvVarEnvironmentPost `json:"envVar,omitempty"`
	State NullableString `json:"state,omitempty"`
	EnvironmentComponent []EnvironmentComponentEnvironmentPost `json:"environmentComponent,omitempty"`
	Region NullableString `json:"region,omitempty"`
	ExportedData []string `json:"exportedData,omitempty"`
	ImportedData []string `json:"importedData,omitempty"`
}

// NewEnvironmentEnvironmentPost instantiates a new EnvironmentEnvironmentPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentEnvironmentPost() *EnvironmentEnvironmentPost {
	this := EnvironmentEnvironmentPost{}
	return &this
}

// NewEnvironmentEnvironmentPostWithDefaults instantiates a new EnvironmentEnvironmentPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentEnvironmentPostWithDefaults() *EnvironmentEnvironmentPost {
	this := EnvironmentEnvironmentPost{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EnvironmentEnvironmentPost) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetCpu() int32 {
	if o == nil || isNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetCpuOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *EnvironmentEnvironmentPost) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetMemory() int32 {
	if o == nil || isNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetMemoryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *EnvironmentEnvironmentPost) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetStorage() int32 {
	if o == nil || isNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetStorageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *EnvironmentEnvironmentPost) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetStorage() {
	o.Storage.Unset()
}

// GetDataTransferSettings returns the DataTransferSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetDataTransferSettings() string {
	if o == nil || isNil(o.DataTransferSettings.Get()) {
		var ret string
		return ret
	}
	return *o.DataTransferSettings.Get()
}

// GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetDataTransferSettingsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DataTransferSettings.Get(), o.DataTransferSettings.IsSet()
}

// HasDataTransferSettings returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasDataTransferSettings() bool {
	if o != nil && o.DataTransferSettings.IsSet() {
		return true
	}

	return false
}

// SetDataTransferSettings gets a reference to the given NullableString and assigns it to the DataTransferSettings field.
func (o *EnvironmentEnvironmentPost) SetDataTransferSettings(v string) {
	o.DataTransferSettings.Set(&v)
}
// SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetDataTransferSettingsNil() {
	o.DataTransferSettings.Set(nil)
}

// UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetDataTransferSettings() {
	o.DataTransferSettings.Unset()
}

// GetIsStripDatabase returns the IsStripDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetIsStripDatabase() bool {
	if o == nil || isNil(o.IsStripDatabase.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStripDatabase.Get()
}

// GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetIsStripDatabaseOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsStripDatabase.Get(), o.IsStripDatabase.IsSet()
}

// HasIsStripDatabase returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasIsStripDatabase() bool {
	if o != nil && o.IsStripDatabase.IsSet() {
		return true
	}

	return false
}

// SetIsStripDatabase gets a reference to the given NullableBool and assigns it to the IsStripDatabase field.
func (o *EnvironmentEnvironmentPost) SetIsStripDatabase(v bool) {
	o.IsStripDatabase.Set(&v)
}
// SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetIsStripDatabaseNil() {
	o.IsStripDatabase.Set(nil)
}

// UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetIsStripDatabase() {
	o.IsStripDatabase.Unset()
}

// GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetIsAllowOutgoingEmails() bool {
	if o == nil || isNil(o.IsAllowOutgoingEmails.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAllowOutgoingEmails.Get()
}

// GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetIsAllowOutgoingEmailsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsAllowOutgoingEmails.Get(), o.IsAllowOutgoingEmails.IsSet()
}

// HasIsAllowOutgoingEmails returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasIsAllowOutgoingEmails() bool {
	if o != nil && o.IsAllowOutgoingEmails.IsSet() {
		return true
	}

	return false
}

// SetIsAllowOutgoingEmails gets a reference to the given NullableBool and assigns it to the IsAllowOutgoingEmails field.
func (o *EnvironmentEnvironmentPost) SetIsAllowOutgoingEmails(v bool) {
	o.IsAllowOutgoingEmails.Set(&v)
}
// SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetIsAllowOutgoingEmailsNil() {
	o.IsAllowOutgoingEmails.Set(nil)
}

// UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetIsAllowOutgoingEmails() {
	o.IsAllowOutgoingEmails.Unset()
}

// GetIsInitSampleData returns the IsInitSampleData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetIsInitSampleData() bool {
	if o == nil || isNil(o.IsInitSampleData.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitSampleData.Get()
}

// GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetIsInitSampleDataOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsInitSampleData.Get(), o.IsInitSampleData.IsSet()
}

// HasIsInitSampleData returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasIsInitSampleData() bool {
	if o != nil && o.IsInitSampleData.IsSet() {
		return true
	}

	return false
}

// SetIsInitSampleData gets a reference to the given NullableBool and assigns it to the IsInitSampleData field.
func (o *EnvironmentEnvironmentPost) SetIsInitSampleData(v bool) {
	o.IsInitSampleData.Set(&v)
}
// SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetIsInitSampleDataNil() {
	o.IsInitSampleData.Set(nil)
}

// UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetIsInitSampleData() {
	o.IsInitSampleData.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetProject() string {
	if o == nil || isNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetProjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *EnvironmentEnvironmentPost) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetProject() {
	o.Project.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetProvider() string {
	if o == nil || isNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetProviderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *EnvironmentEnvironmentPost) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetProvider() {
	o.Provider.Unset()
}

// GetEnvVar returns the EnvVar field value if set, zero value otherwise.
func (o *EnvironmentEnvironmentPost) GetEnvVar() []EnvironmentEnvVarEnvironmentPost {
	if o == nil || isNil(o.EnvVar) {
		var ret []EnvironmentEnvVarEnvironmentPost
		return ret
	}
	return o.EnvVar
}

// GetEnvVarOk returns a tuple with the EnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEnvironmentPost) GetEnvVarOk() ([]EnvironmentEnvVarEnvironmentPost, bool) {
	if o == nil || isNil(o.EnvVar) {
    return nil, false
	}
	return o.EnvVar, true
}

// HasEnvVar returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasEnvVar() bool {
	if o != nil && !isNil(o.EnvVar) {
		return true
	}

	return false
}

// SetEnvVar gets a reference to the given []EnvironmentEnvVarEnvironmentPost and assigns it to the EnvVar field.
func (o *EnvironmentEnvironmentPost) SetEnvVar(v []EnvironmentEnvVarEnvironmentPost) {
	o.EnvVar = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetState() string {
	if o == nil || isNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetStateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *EnvironmentEnvironmentPost) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetState() {
	o.State.Unset()
}

// GetEnvironmentComponent returns the EnvironmentComponent field value if set, zero value otherwise.
func (o *EnvironmentEnvironmentPost) GetEnvironmentComponent() []EnvironmentComponentEnvironmentPost {
	if o == nil || isNil(o.EnvironmentComponent) {
		var ret []EnvironmentComponentEnvironmentPost
		return ret
	}
	return o.EnvironmentComponent
}

// GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEnvironmentPost) GetEnvironmentComponentOk() ([]EnvironmentComponentEnvironmentPost, bool) {
	if o == nil || isNil(o.EnvironmentComponent) {
    return nil, false
	}
	return o.EnvironmentComponent, true
}

// HasEnvironmentComponent returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasEnvironmentComponent() bool {
	if o != nil && !isNil(o.EnvironmentComponent) {
		return true
	}

	return false
}

// SetEnvironmentComponent gets a reference to the given []EnvironmentComponentEnvironmentPost and assigns it to the EnvironmentComponent field.
func (o *EnvironmentEnvironmentPost) SetEnvironmentComponent(v []EnvironmentComponentEnvironmentPost) {
	o.EnvironmentComponent = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentEnvironmentPost) GetRegion() string {
	if o == nil || isNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentEnvironmentPost) GetRegionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *EnvironmentEnvironmentPost) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *EnvironmentEnvironmentPost) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *EnvironmentEnvironmentPost) UnsetRegion() {
	o.Region.Unset()
}

// GetExportedData returns the ExportedData field value if set, zero value otherwise.
func (o *EnvironmentEnvironmentPost) GetExportedData() []string {
	if o == nil || isNil(o.ExportedData) {
		var ret []string
		return ret
	}
	return o.ExportedData
}

// GetExportedDataOk returns a tuple with the ExportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEnvironmentPost) GetExportedDataOk() ([]string, bool) {
	if o == nil || isNil(o.ExportedData) {
    return nil, false
	}
	return o.ExportedData, true
}

// HasExportedData returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasExportedData() bool {
	if o != nil && !isNil(o.ExportedData) {
		return true
	}

	return false
}

// SetExportedData gets a reference to the given []string and assigns it to the ExportedData field.
func (o *EnvironmentEnvironmentPost) SetExportedData(v []string) {
	o.ExportedData = v
}

// GetImportedData returns the ImportedData field value if set, zero value otherwise.
func (o *EnvironmentEnvironmentPost) GetImportedData() []string {
	if o == nil || isNil(o.ImportedData) {
		var ret []string
		return ret
	}
	return o.ImportedData
}

// GetImportedDataOk returns a tuple with the ImportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentEnvironmentPost) GetImportedDataOk() ([]string, bool) {
	if o == nil || isNil(o.ImportedData) {
    return nil, false
	}
	return o.ImportedData, true
}

// HasImportedData returns a boolean if a field has been set.
func (o *EnvironmentEnvironmentPost) HasImportedData() bool {
	if o != nil && !isNil(o.ImportedData) {
		return true
	}

	return false
}

// SetImportedData gets a reference to the given []string and assigns it to the ImportedData field.
func (o *EnvironmentEnvironmentPost) SetImportedData(v []string) {
	o.ImportedData = v
}

func (o EnvironmentEnvironmentPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.DataTransferSettings.IsSet() {
		toSerialize["dataTransferSettings"] = o.DataTransferSettings.Get()
	}
	if o.IsStripDatabase.IsSet() {
		toSerialize["isStripDatabase"] = o.IsStripDatabase.Get()
	}
	if o.IsAllowOutgoingEmails.IsSet() {
		toSerialize["isAllowOutgoingEmails"] = o.IsAllowOutgoingEmails.Get()
	}
	if o.IsInitSampleData.IsSet() {
		toSerialize["isInitSampleData"] = o.IsInitSampleData.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if !isNil(o.EnvVar) {
		toSerialize["envVar"] = o.EnvVar
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !isNil(o.EnvironmentComponent) {
		toSerialize["environmentComponent"] = o.EnvironmentComponent
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if !isNil(o.ExportedData) {
		toSerialize["exportedData"] = o.ExportedData
	}
	if !isNil(o.ImportedData) {
		toSerialize["importedData"] = o.ImportedData
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentEnvironmentPost struct {
	value *EnvironmentEnvironmentPost
	isSet bool
}

func (v NullableEnvironmentEnvironmentPost) Get() *EnvironmentEnvironmentPost {
	return v.value
}

func (v *NullableEnvironmentEnvironmentPost) Set(val *EnvironmentEnvironmentPost) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentEnvironmentPost) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentEnvironmentPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentEnvironmentPost(val *EnvironmentEnvironmentPost) *NullableEnvironmentEnvironmentPost {
	return &NullableEnvironmentEnvironmentPost{value: val, isSet: true}
}

func (v NullableEnvironmentEnvironmentPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentEnvironmentPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

