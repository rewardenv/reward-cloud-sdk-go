/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
	"time"
)

// EnvironmentJsonhalEnvironmentGet Class Environment
type EnvironmentJsonhalEnvironmentGet struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
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
	EnvVar []EnvironmentEnvVarJsonhalEnvironmentGet `json:"envVar,omitempty"`
	State NullableString `json:"state,omitempty"`
	EnvironmentComponent []EnvironmentComponentJsonhalEnvironmentGet `json:"environmentComponent,omitempty"`
	Region NullableString `json:"region,omitempty"`
	ExportedData []string `json:"exportedData,omitempty"`
	EnvironmentAccess NullableString `json:"environmentAccess,omitempty"`
	ImportedData []string `json:"importedData,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewEnvironmentJsonhalEnvironmentGet instantiates a new EnvironmentJsonhalEnvironmentGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentJsonhalEnvironmentGet() *EnvironmentJsonhalEnvironmentGet {
	this := EnvironmentJsonhalEnvironmentGet{}
	return &this
}

// NewEnvironmentJsonhalEnvironmentGetWithDefaults instantiates a new EnvironmentJsonhalEnvironmentGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentJsonhalEnvironmentGetWithDefaults() *EnvironmentJsonhalEnvironmentGet {
	this := EnvironmentJsonhalEnvironmentGet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *EnvironmentJsonhalEnvironmentGet) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EnvironmentJsonhalEnvironmentGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetUuidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *EnvironmentJsonhalEnvironmentGet) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EnvironmentJsonhalEnvironmentGet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetCpu() int32 {
	if o == nil || isNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetCpuOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *EnvironmentJsonhalEnvironmentGet) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetMemory() int32 {
	if o == nil || isNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetMemoryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *EnvironmentJsonhalEnvironmentGet) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetStorage() int32 {
	if o == nil || isNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetStorageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *EnvironmentJsonhalEnvironmentGet) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetStorage() {
	o.Storage.Unset()
}

// GetDataTransferSettings returns the DataTransferSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetDataTransferSettings() string {
	if o == nil || isNil(o.DataTransferSettings.Get()) {
		var ret string
		return ret
	}
	return *o.DataTransferSettings.Get()
}

// GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetDataTransferSettingsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.DataTransferSettings.Get(), o.DataTransferSettings.IsSet()
}

// HasDataTransferSettings returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasDataTransferSettings() bool {
	if o != nil && o.DataTransferSettings.IsSet() {
		return true
	}

	return false
}

// SetDataTransferSettings gets a reference to the given NullableString and assigns it to the DataTransferSettings field.
func (o *EnvironmentJsonhalEnvironmentGet) SetDataTransferSettings(v string) {
	o.DataTransferSettings.Set(&v)
}
// SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetDataTransferSettingsNil() {
	o.DataTransferSettings.Set(nil)
}

// UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetDataTransferSettings() {
	o.DataTransferSettings.Unset()
}

// GetIsStripDatabase returns the IsStripDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetIsStripDatabase() bool {
	if o == nil || isNil(o.IsStripDatabase.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStripDatabase.Get()
}

// GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetIsStripDatabaseOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsStripDatabase.Get(), o.IsStripDatabase.IsSet()
}

// HasIsStripDatabase returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasIsStripDatabase() bool {
	if o != nil && o.IsStripDatabase.IsSet() {
		return true
	}

	return false
}

// SetIsStripDatabase gets a reference to the given NullableBool and assigns it to the IsStripDatabase field.
func (o *EnvironmentJsonhalEnvironmentGet) SetIsStripDatabase(v bool) {
	o.IsStripDatabase.Set(&v)
}
// SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetIsStripDatabaseNil() {
	o.IsStripDatabase.Set(nil)
}

// UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetIsStripDatabase() {
	o.IsStripDatabase.Unset()
}

// GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetIsAllowOutgoingEmails() bool {
	if o == nil || isNil(o.IsAllowOutgoingEmails.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAllowOutgoingEmails.Get()
}

// GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetIsAllowOutgoingEmailsOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsAllowOutgoingEmails.Get(), o.IsAllowOutgoingEmails.IsSet()
}

// HasIsAllowOutgoingEmails returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasIsAllowOutgoingEmails() bool {
	if o != nil && o.IsAllowOutgoingEmails.IsSet() {
		return true
	}

	return false
}

// SetIsAllowOutgoingEmails gets a reference to the given NullableBool and assigns it to the IsAllowOutgoingEmails field.
func (o *EnvironmentJsonhalEnvironmentGet) SetIsAllowOutgoingEmails(v bool) {
	o.IsAllowOutgoingEmails.Set(&v)
}
// SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetIsAllowOutgoingEmailsNil() {
	o.IsAllowOutgoingEmails.Set(nil)
}

// UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetIsAllowOutgoingEmails() {
	o.IsAllowOutgoingEmails.Unset()
}

// GetIsInitSampleData returns the IsInitSampleData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetIsInitSampleData() bool {
	if o == nil || isNil(o.IsInitSampleData.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitSampleData.Get()
}

// GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetIsInitSampleDataOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsInitSampleData.Get(), o.IsInitSampleData.IsSet()
}

// HasIsInitSampleData returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasIsInitSampleData() bool {
	if o != nil && o.IsInitSampleData.IsSet() {
		return true
	}

	return false
}

// SetIsInitSampleData gets a reference to the given NullableBool and assigns it to the IsInitSampleData field.
func (o *EnvironmentJsonhalEnvironmentGet) SetIsInitSampleData(v bool) {
	o.IsInitSampleData.Set(&v)
}
// SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetIsInitSampleDataNil() {
	o.IsInitSampleData.Set(nil)
}

// UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetIsInitSampleData() {
	o.IsInitSampleData.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetProject() string {
	if o == nil || isNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetProjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *EnvironmentJsonhalEnvironmentGet) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetProject() {
	o.Project.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetProvider() string {
	if o == nil || isNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetProviderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *EnvironmentJsonhalEnvironmentGet) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetProvider() {
	o.Provider.Unset()
}

// GetEnvVar returns the EnvVar field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetEnvVar() []EnvironmentEnvVarJsonhalEnvironmentGet {
	if o == nil || isNil(o.EnvVar) {
		var ret []EnvironmentEnvVarJsonhalEnvironmentGet
		return ret
	}
	return o.EnvVar
}

// GetEnvVarOk returns a tuple with the EnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetEnvVarOk() ([]EnvironmentEnvVarJsonhalEnvironmentGet, bool) {
	if o == nil || isNil(o.EnvVar) {
    return nil, false
	}
	return o.EnvVar, true
}

// HasEnvVar returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasEnvVar() bool {
	if o != nil && !isNil(o.EnvVar) {
		return true
	}

	return false
}

// SetEnvVar gets a reference to the given []EnvironmentEnvVarJsonhalEnvironmentGet and assigns it to the EnvVar field.
func (o *EnvironmentJsonhalEnvironmentGet) SetEnvVar(v []EnvironmentEnvVarJsonhalEnvironmentGet) {
	o.EnvVar = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetState() string {
	if o == nil || isNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetStateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *EnvironmentJsonhalEnvironmentGet) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetState() {
	o.State.Unset()
}

// GetEnvironmentComponent returns the EnvironmentComponent field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentComponent() []EnvironmentComponentJsonhalEnvironmentGet {
	if o == nil || isNil(o.EnvironmentComponent) {
		var ret []EnvironmentComponentJsonhalEnvironmentGet
		return ret
	}
	return o.EnvironmentComponent
}

// GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentComponentOk() ([]EnvironmentComponentJsonhalEnvironmentGet, bool) {
	if o == nil || isNil(o.EnvironmentComponent) {
    return nil, false
	}
	return o.EnvironmentComponent, true
}

// HasEnvironmentComponent returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasEnvironmentComponent() bool {
	if o != nil && !isNil(o.EnvironmentComponent) {
		return true
	}

	return false
}

// SetEnvironmentComponent gets a reference to the given []EnvironmentComponentJsonhalEnvironmentGet and assigns it to the EnvironmentComponent field.
func (o *EnvironmentJsonhalEnvironmentGet) SetEnvironmentComponent(v []EnvironmentComponentJsonhalEnvironmentGet) {
	o.EnvironmentComponent = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetRegion() string {
	if o == nil || isNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetRegionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *EnvironmentJsonhalEnvironmentGet) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetRegion() {
	o.Region.Unset()
}

// GetExportedData returns the ExportedData field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetExportedData() []string {
	if o == nil || isNil(o.ExportedData) {
		var ret []string
		return ret
	}
	return o.ExportedData
}

// GetExportedDataOk returns a tuple with the ExportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetExportedDataOk() ([]string, bool) {
	if o == nil || isNil(o.ExportedData) {
    return nil, false
	}
	return o.ExportedData, true
}

// HasExportedData returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasExportedData() bool {
	if o != nil && !isNil(o.ExportedData) {
		return true
	}

	return false
}

// SetExportedData gets a reference to the given []string and assigns it to the ExportedData field.
func (o *EnvironmentJsonhalEnvironmentGet) SetExportedData(v []string) {
	o.ExportedData = v
}

// GetEnvironmentAccess returns the EnvironmentAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentAccess() string {
	if o == nil || isNil(o.EnvironmentAccess.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentAccess.Get()
}

// GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentGet) GetEnvironmentAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvironmentAccess.Get(), o.EnvironmentAccess.IsSet()
}

// HasEnvironmentAccess returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasEnvironmentAccess() bool {
	if o != nil && o.EnvironmentAccess.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentAccess gets a reference to the given NullableString and assigns it to the EnvironmentAccess field.
func (o *EnvironmentJsonhalEnvironmentGet) SetEnvironmentAccess(v string) {
	o.EnvironmentAccess.Set(&v)
}
// SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) SetEnvironmentAccessNil() {
	o.EnvironmentAccess.Set(nil)
}

// UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentGet) UnsetEnvironmentAccess() {
	o.EnvironmentAccess.Unset()
}

// GetImportedData returns the ImportedData field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetImportedData() []string {
	if o == nil || isNil(o.ImportedData) {
		var ret []string
		return ret
	}
	return o.ImportedData
}

// GetImportedDataOk returns a tuple with the ImportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetImportedDataOk() ([]string, bool) {
	if o == nil || isNil(o.ImportedData) {
    return nil, false
	}
	return o.ImportedData, true
}

// HasImportedData returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasImportedData() bool {
	if o != nil && !isNil(o.ImportedData) {
		return true
	}

	return false
}

// SetImportedData gets a reference to the given []string and assigns it to the ImportedData field.
func (o *EnvironmentJsonhalEnvironmentGet) SetImportedData(v []string) {
	o.ImportedData = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvironmentJsonhalEnvironmentGet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentGet) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentGet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentGet) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentJsonhalEnvironmentGet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentJsonhalEnvironmentGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
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
	if o.EnvironmentAccess.IsSet() {
		toSerialize["environmentAccess"] = o.EnvironmentAccess.Get()
	}
	if !isNil(o.ImportedData) {
		toSerialize["importedData"] = o.ImportedData
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentJsonhalEnvironmentGet struct {
	value *EnvironmentJsonhalEnvironmentGet
	isSet bool
}

func (v NullableEnvironmentJsonhalEnvironmentGet) Get() *EnvironmentJsonhalEnvironmentGet {
	return v.value
}

func (v *NullableEnvironmentJsonhalEnvironmentGet) Set(val *EnvironmentJsonhalEnvironmentGet) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentJsonhalEnvironmentGet) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentJsonhalEnvironmentGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentJsonhalEnvironmentGet(val *EnvironmentJsonhalEnvironmentGet) *NullableEnvironmentJsonhalEnvironmentGet {
	return &NullableEnvironmentJsonhalEnvironmentGet{value: val, isSet: true}
}

func (v NullableEnvironmentJsonhalEnvironmentGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentJsonhalEnvironmentGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


