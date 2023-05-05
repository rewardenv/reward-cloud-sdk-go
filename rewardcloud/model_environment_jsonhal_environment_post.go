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

// checks if the EnvironmentJsonhalEnvironmentPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentJsonhalEnvironmentPost{}

// EnvironmentJsonhalEnvironmentPost Class Environment
type EnvironmentJsonhalEnvironmentPost struct {
	Links                 *AbstractEnvironmentJsonhalLinks             `json:"_links,omitempty"`
	TemplateIri           NullableString                               `json:"templateIri,omitempty"`
	Project               NullableString                               `json:"project,omitempty"`
	Provider              NullableString                               `json:"provider,omitempty"`
	State                 NullableString                               `json:"state,omitempty"`
	Region                NullableString                               `json:"region,omitempty"`
	ExportedData          []string                                     `json:"exportedData,omitempty"`
	ImportedData          []string                                     `json:"importedData,omitempty"`
	Name                  NullableString                               `json:"name,omitempty"`
	Cpu                   NullableInt32                                `json:"cpu,omitempty"`
	Memory                NullableInt32                                `json:"memory,omitempty"`
	Storage               NullableInt32                                `json:"storage,omitempty"`
	DataTransferSettings  NullableString                               `json:"dataTransferSettings,omitempty"`
	IsStripDatabase       NullableBool                                 `json:"isStripDatabase,omitempty"`
	IsAllowOutgoingEmails NullableBool                                 `json:"isAllowOutgoingEmails,omitempty"`
	IsInitSampleData      NullableBool                                 `json:"isInitSampleData,omitempty"`
	EnvVar                []EnvironmentEnvVarJsonhalEnvironmentPost    `json:"envVar,omitempty"`
	EnvironmentComponent  []EnvironmentComponentJsonhalEnvironmentPost `json:"environmentComponent,omitempty"`
}

// NewEnvironmentJsonhalEnvironmentPost instantiates a new EnvironmentJsonhalEnvironmentPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentJsonhalEnvironmentPost() *EnvironmentJsonhalEnvironmentPost {
	this := EnvironmentJsonhalEnvironmentPost{}
	return &this
}

// NewEnvironmentJsonhalEnvironmentPostWithDefaults instantiates a new EnvironmentJsonhalEnvironmentPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentJsonhalEnvironmentPostWithDefaults() *EnvironmentJsonhalEnvironmentPost {
	this := EnvironmentJsonhalEnvironmentPost{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentPost) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentPost) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *EnvironmentJsonhalEnvironmentPost) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetTemplateIri returns the TemplateIri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetTemplateIri() string {
	if o == nil || IsNil(o.TemplateIri.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateIri.Get()
}

// GetTemplateIriOk returns a tuple with the TemplateIri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetTemplateIriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateIri.Get(), o.TemplateIri.IsSet()
}

// HasTemplateIri returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasTemplateIri() bool {
	if o != nil && o.TemplateIri.IsSet() {
		return true
	}

	return false
}

// SetTemplateIri gets a reference to the given NullableString and assigns it to the TemplateIri field.
func (o *EnvironmentJsonhalEnvironmentPost) SetTemplateIri(v string) {
	o.TemplateIri.Set(&v)
}

// SetTemplateIriNil sets the value for TemplateIri to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetTemplateIriNil() {
	o.TemplateIri.Set(nil)
}

// UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetTemplateIri() {
	o.TemplateIri.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *EnvironmentJsonhalEnvironmentPost) SetProject(v string) {
	o.Project.Set(&v)
}

// SetProjectNil sets the value for Project to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetProject() {
	o.Project.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *EnvironmentJsonhalEnvironmentPost) SetProvider(v string) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetProvider() {
	o.Provider.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *EnvironmentJsonhalEnvironmentPost) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetState() {
	o.State.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *EnvironmentJsonhalEnvironmentPost) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetRegion() {
	o.Region.Unset()
}

// GetExportedData returns the ExportedData field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentPost) GetExportedData() []string {
	if o == nil || IsNil(o.ExportedData) {
		var ret []string
		return ret
	}
	return o.ExportedData
}

// GetExportedDataOk returns a tuple with the ExportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentPost) GetExportedDataOk() ([]string, bool) {
	if o == nil || IsNil(o.ExportedData) {
		return nil, false
	}
	return o.ExportedData, true
}

// HasExportedData returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasExportedData() bool {
	if o != nil && !IsNil(o.ExportedData) {
		return true
	}

	return false
}

// SetExportedData gets a reference to the given []string and assigns it to the ExportedData field.
func (o *EnvironmentJsonhalEnvironmentPost) SetExportedData(v []string) {
	o.ExportedData = v
}

// GetImportedData returns the ImportedData field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentPost) GetImportedData() []string {
	if o == nil || IsNil(o.ImportedData) {
		var ret []string
		return ret
	}
	return o.ImportedData
}

// GetImportedDataOk returns a tuple with the ImportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentPost) GetImportedDataOk() ([]string, bool) {
	if o == nil || IsNil(o.ImportedData) {
		return nil, false
	}
	return o.ImportedData, true
}

// HasImportedData returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasImportedData() bool {
	if o != nil && !IsNil(o.ImportedData) {
		return true
	}

	return false
}

// SetImportedData gets a reference to the given []string and assigns it to the ImportedData field.
func (o *EnvironmentJsonhalEnvironmentPost) SetImportedData(v []string) {
	o.ImportedData = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EnvironmentJsonhalEnvironmentPost) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *EnvironmentJsonhalEnvironmentPost) SetCpu(v int32) {
	o.Cpu.Set(&v)
}

// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *EnvironmentJsonhalEnvironmentPost) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetStorage() int32 {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *EnvironmentJsonhalEnvironmentPost) SetStorage(v int32) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetStorage() {
	o.Storage.Unset()
}

// GetDataTransferSettings returns the DataTransferSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetDataTransferSettings() string {
	if o == nil || IsNil(o.DataTransferSettings.Get()) {
		var ret string
		return ret
	}
	return *o.DataTransferSettings.Get()
}

// GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetDataTransferSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataTransferSettings.Get(), o.DataTransferSettings.IsSet()
}

// HasDataTransferSettings returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasDataTransferSettings() bool {
	if o != nil && o.DataTransferSettings.IsSet() {
		return true
	}

	return false
}

// SetDataTransferSettings gets a reference to the given NullableString and assigns it to the DataTransferSettings field.
func (o *EnvironmentJsonhalEnvironmentPost) SetDataTransferSettings(v string) {
	o.DataTransferSettings.Set(&v)
}

// SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetDataTransferSettingsNil() {
	o.DataTransferSettings.Set(nil)
}

// UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetDataTransferSettings() {
	o.DataTransferSettings.Unset()
}

// GetIsStripDatabase returns the IsStripDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetIsStripDatabase() bool {
	if o == nil || IsNil(o.IsStripDatabase.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStripDatabase.Get()
}

// GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetIsStripDatabaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsStripDatabase.Get(), o.IsStripDatabase.IsSet()
}

// HasIsStripDatabase returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasIsStripDatabase() bool {
	if o != nil && o.IsStripDatabase.IsSet() {
		return true
	}

	return false
}

// SetIsStripDatabase gets a reference to the given NullableBool and assigns it to the IsStripDatabase field.
func (o *EnvironmentJsonhalEnvironmentPost) SetIsStripDatabase(v bool) {
	o.IsStripDatabase.Set(&v)
}

// SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetIsStripDatabaseNil() {
	o.IsStripDatabase.Set(nil)
}

// UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetIsStripDatabase() {
	o.IsStripDatabase.Unset()
}

// GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetIsAllowOutgoingEmails() bool {
	if o == nil || IsNil(o.IsAllowOutgoingEmails.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAllowOutgoingEmails.Get()
}

// GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetIsAllowOutgoingEmailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAllowOutgoingEmails.Get(), o.IsAllowOutgoingEmails.IsSet()
}

// HasIsAllowOutgoingEmails returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasIsAllowOutgoingEmails() bool {
	if o != nil && o.IsAllowOutgoingEmails.IsSet() {
		return true
	}

	return false
}

// SetIsAllowOutgoingEmails gets a reference to the given NullableBool and assigns it to the IsAllowOutgoingEmails field.
func (o *EnvironmentJsonhalEnvironmentPost) SetIsAllowOutgoingEmails(v bool) {
	o.IsAllowOutgoingEmails.Set(&v)
}

// SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetIsAllowOutgoingEmailsNil() {
	o.IsAllowOutgoingEmails.Set(nil)
}

// UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetIsAllowOutgoingEmails() {
	o.IsAllowOutgoingEmails.Unset()
}

// GetIsInitSampleData returns the IsInitSampleData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentJsonhalEnvironmentPost) GetIsInitSampleData() bool {
	if o == nil || IsNil(o.IsInitSampleData.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitSampleData.Get()
}

// GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentJsonhalEnvironmentPost) GetIsInitSampleDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitSampleData.Get(), o.IsInitSampleData.IsSet()
}

// HasIsInitSampleData returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasIsInitSampleData() bool {
	if o != nil && o.IsInitSampleData.IsSet() {
		return true
	}

	return false
}

// SetIsInitSampleData gets a reference to the given NullableBool and assigns it to the IsInitSampleData field.
func (o *EnvironmentJsonhalEnvironmentPost) SetIsInitSampleData(v bool) {
	o.IsInitSampleData.Set(&v)
}

// SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) SetIsInitSampleDataNil() {
	o.IsInitSampleData.Set(nil)
}

// UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
func (o *EnvironmentJsonhalEnvironmentPost) UnsetIsInitSampleData() {
	o.IsInitSampleData.Unset()
}

// GetEnvVar returns the EnvVar field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentPost) GetEnvVar() []EnvironmentEnvVarJsonhalEnvironmentPost {
	if o == nil || IsNil(o.EnvVar) {
		var ret []EnvironmentEnvVarJsonhalEnvironmentPost
		return ret
	}
	return o.EnvVar
}

// GetEnvVarOk returns a tuple with the EnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentPost) GetEnvVarOk() ([]EnvironmentEnvVarJsonhalEnvironmentPost, bool) {
	if o == nil || IsNil(o.EnvVar) {
		return nil, false
	}
	return o.EnvVar, true
}

// HasEnvVar returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasEnvVar() bool {
	if o != nil && !IsNil(o.EnvVar) {
		return true
	}

	return false
}

// SetEnvVar gets a reference to the given []EnvironmentEnvVarJsonhalEnvironmentPost and assigns it to the EnvVar field.
func (o *EnvironmentJsonhalEnvironmentPost) SetEnvVar(v []EnvironmentEnvVarJsonhalEnvironmentPost) {
	o.EnvVar = v
}

// GetEnvironmentComponent returns the EnvironmentComponent field value if set, zero value otherwise.
func (o *EnvironmentJsonhalEnvironmentPost) GetEnvironmentComponent() []EnvironmentComponentJsonhalEnvironmentPost {
	if o == nil || IsNil(o.EnvironmentComponent) {
		var ret []EnvironmentComponentJsonhalEnvironmentPost
		return ret
	}
	return o.EnvironmentComponent
}

// GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentJsonhalEnvironmentPost) GetEnvironmentComponentOk() ([]EnvironmentComponentJsonhalEnvironmentPost, bool) {
	if o == nil || IsNil(o.EnvironmentComponent) {
		return nil, false
	}
	return o.EnvironmentComponent, true
}

// HasEnvironmentComponent returns a boolean if a field has been set.
func (o *EnvironmentJsonhalEnvironmentPost) HasEnvironmentComponent() bool {
	if o != nil && !IsNil(o.EnvironmentComponent) {
		return true
	}

	return false
}

// SetEnvironmentComponent gets a reference to the given []EnvironmentComponentJsonhalEnvironmentPost and assigns it to the EnvironmentComponent field.
func (o *EnvironmentJsonhalEnvironmentPost) SetEnvironmentComponent(v []EnvironmentComponentJsonhalEnvironmentPost) {
	o.EnvironmentComponent = v
}

func (o EnvironmentJsonhalEnvironmentPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentJsonhalEnvironmentPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if o.TemplateIri.IsSet() {
		toSerialize["templateIri"] = o.TemplateIri.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if !IsNil(o.ExportedData) {
		toSerialize["exportedData"] = o.ExportedData
	}
	if !IsNil(o.ImportedData) {
		toSerialize["importedData"] = o.ImportedData
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
	if !IsNil(o.EnvVar) {
		toSerialize["envVar"] = o.EnvVar
	}
	if !IsNil(o.EnvironmentComponent) {
		toSerialize["environmentComponent"] = o.EnvironmentComponent
	}
	return toSerialize, nil
}

type NullableEnvironmentJsonhalEnvironmentPost struct {
	value *EnvironmentJsonhalEnvironmentPost
	isSet bool
}

func (v NullableEnvironmentJsonhalEnvironmentPost) Get() *EnvironmentJsonhalEnvironmentPost {
	return v.value
}

func (v *NullableEnvironmentJsonhalEnvironmentPost) Set(val *EnvironmentJsonhalEnvironmentPost) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentJsonhalEnvironmentPost) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentJsonhalEnvironmentPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentJsonhalEnvironmentPost(val *EnvironmentJsonhalEnvironmentPost) *NullableEnvironmentJsonhalEnvironmentPost {
	return &NullableEnvironmentJsonhalEnvironmentPost{value: val, isSet: true}
}

func (v NullableEnvironmentJsonhalEnvironmentPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentJsonhalEnvironmentPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
