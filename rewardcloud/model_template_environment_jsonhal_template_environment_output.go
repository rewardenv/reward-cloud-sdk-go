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

// checks if the TemplateEnvironmentJsonhalTemplateEnvironmentOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateEnvironmentJsonhalTemplateEnvironmentOutput{}

// TemplateEnvironmentJsonhalTemplateEnvironmentOutput Class TemplateEnvironment
type TemplateEnvironmentJsonhalTemplateEnvironmentOutput struct {
	Links                 *AbstractEnvironmentJsonhalLinks                       `json:"_links,omitempty"`
	TemplateProject       NullableString                                         `json:"templateProject,omitempty"`
	Id                    *int32                                                 `json:"id,omitempty"`
	Uuid                  NullableString                                         `json:"uuid,omitempty"`
	Name                  NullableString                                         `json:"name,omitempty"`
	Cpu                   NullableInt32                                          `json:"cpu,omitempty"`
	Memory                NullableInt32                                          `json:"memory,omitempty"`
	Storage               NullableInt32                                          `json:"storage,omitempty"`
	DataTransferSettings  NullableString                                         `json:"dataTransferSettings,omitempty"`
	IsStripDatabase       NullableBool                                           `json:"isStripDatabase,omitempty"`
	IsAllowOutgoingEmails NullableBool                                           `json:"isAllowOutgoingEmails,omitempty"`
	IsInitSampleData      NullableBool                                           `json:"isInitSampleData,omitempty"`
	CreatedAt             *time.Time                                             `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time                                             `json:"updatedAt,omitempty"`
	EnvVar                []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput    `json:"envVar,omitempty"`
	EnvironmentComponent  []EnvironmentComponentJsonhalTemplateEnvironmentOutput `json:"environmentComponent,omitempty"`
}

// NewTemplateEnvironmentJsonhalTemplateEnvironmentOutput instantiates a new TemplateEnvironmentJsonhalTemplateEnvironmentOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateEnvironmentJsonhalTemplateEnvironmentOutput() *TemplateEnvironmentJsonhalTemplateEnvironmentOutput {
	this := TemplateEnvironmentJsonhalTemplateEnvironmentOutput{}
	return &this
}

// NewTemplateEnvironmentJsonhalTemplateEnvironmentOutputWithDefaults instantiates a new TemplateEnvironmentJsonhalTemplateEnvironmentOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateEnvironmentJsonhalTemplateEnvironmentOutputWithDefaults() *TemplateEnvironmentJsonhalTemplateEnvironmentOutput {
	this := TemplateEnvironmentJsonhalTemplateEnvironmentOutput{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetTemplateProject returns the TemplateProject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetTemplateProject() string {
	if o == nil || IsNil(o.TemplateProject.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateProject.Get()
}

// GetTemplateProjectOk returns a tuple with the TemplateProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetTemplateProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateProject.Get(), o.TemplateProject.IsSet()
}

// HasTemplateProject returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasTemplateProject() bool {
	if o != nil && o.TemplateProject.IsSet() {
		return true
	}

	return false
}

// SetTemplateProject gets a reference to the given NullableString and assigns it to the TemplateProject field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetTemplateProject(v string) {
	o.TemplateProject.Set(&v)
}

// SetTemplateProjectNil sets the value for TemplateProject to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetTemplateProjectNil() {
	o.TemplateProject.Set(nil)
}

// UnsetTemplateProject ensures that no value is present for TemplateProject, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetTemplateProject() {
	o.TemplateProject.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetCpu(v int32) {
	o.Cpu.Set(&v)
}

// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetStorage() int32 {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetStorage(v int32) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetStorage() {
	o.Storage.Unset()
}

// GetDataTransferSettings returns the DataTransferSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetDataTransferSettings() string {
	if o == nil || IsNil(o.DataTransferSettings.Get()) {
		var ret string
		return ret
	}
	return *o.DataTransferSettings.Get()
}

// GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetDataTransferSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataTransferSettings.Get(), o.DataTransferSettings.IsSet()
}

// HasDataTransferSettings returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasDataTransferSettings() bool {
	if o != nil && o.DataTransferSettings.IsSet() {
		return true
	}

	return false
}

// SetDataTransferSettings gets a reference to the given NullableString and assigns it to the DataTransferSettings field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetDataTransferSettings(v string) {
	o.DataTransferSettings.Set(&v)
}

// SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetDataTransferSettingsNil() {
	o.DataTransferSettings.Set(nil)
}

// UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetDataTransferSettings() {
	o.DataTransferSettings.Unset()
}

// GetIsStripDatabase returns the IsStripDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsStripDatabase() bool {
	if o == nil || IsNil(o.IsStripDatabase.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStripDatabase.Get()
}

// GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsStripDatabaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsStripDatabase.Get(), o.IsStripDatabase.IsSet()
}

// HasIsStripDatabase returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasIsStripDatabase() bool {
	if o != nil && o.IsStripDatabase.IsSet() {
		return true
	}

	return false
}

// SetIsStripDatabase gets a reference to the given NullableBool and assigns it to the IsStripDatabase field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsStripDatabase(v bool) {
	o.IsStripDatabase.Set(&v)
}

// SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsStripDatabaseNil() {
	o.IsStripDatabase.Set(nil)
}

// UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetIsStripDatabase() {
	o.IsStripDatabase.Unset()
}

// GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsAllowOutgoingEmails() bool {
	if o == nil || IsNil(o.IsAllowOutgoingEmails.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAllowOutgoingEmails.Get()
}

// GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsAllowOutgoingEmailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAllowOutgoingEmails.Get(), o.IsAllowOutgoingEmails.IsSet()
}

// HasIsAllowOutgoingEmails returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasIsAllowOutgoingEmails() bool {
	if o != nil && o.IsAllowOutgoingEmails.IsSet() {
		return true
	}

	return false
}

// SetIsAllowOutgoingEmails gets a reference to the given NullableBool and assigns it to the IsAllowOutgoingEmails field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsAllowOutgoingEmails(v bool) {
	o.IsAllowOutgoingEmails.Set(&v)
}

// SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsAllowOutgoingEmailsNil() {
	o.IsAllowOutgoingEmails.Set(nil)
}

// UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetIsAllowOutgoingEmails() {
	o.IsAllowOutgoingEmails.Unset()
}

// GetIsInitSampleData returns the IsInitSampleData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsInitSampleData() bool {
	if o == nil || IsNil(o.IsInitSampleData.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitSampleData.Get()
}

// GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetIsInitSampleDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitSampleData.Get(), o.IsInitSampleData.IsSet()
}

// HasIsInitSampleData returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasIsInitSampleData() bool {
	if o != nil && o.IsInitSampleData.IsSet() {
		return true
	}

	return false
}

// SetIsInitSampleData gets a reference to the given NullableBool and assigns it to the IsInitSampleData field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsInitSampleData(v bool) {
	o.IsInitSampleData.Set(&v)
}

// SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetIsInitSampleDataNil() {
	o.IsInitSampleData.Set(nil)
}

// UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnsetIsInitSampleData() {
	o.IsInitSampleData.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvVar returns the EnvVar field value if set, zero value otherwise.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvVar() []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput {
	if o == nil || IsNil(o.EnvVar) {
		var ret []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput
		return ret
	}
	return o.EnvVar
}

// GetEnvVarOk returns a tuple with the EnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvVarOk() ([]EnvironmentEnvVarJsonhalTemplateEnvironmentOutput, bool) {
	if o == nil || IsNil(o.EnvVar) {
		return nil, false
	}
	return o.EnvVar, true
}

// HasEnvVar returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasEnvVar() bool {
	if o != nil && !IsNil(o.EnvVar) {
		return true
	}

	return false
}

// SetEnvVar gets a reference to the given []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput and assigns it to the EnvVar field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetEnvVar(v []EnvironmentEnvVarJsonhalTemplateEnvironmentOutput) {
	o.EnvVar = v
}

// GetEnvironmentComponent returns the EnvironmentComponent field value if set, zero value otherwise.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvironmentComponent() []EnvironmentComponentJsonhalTemplateEnvironmentOutput {
	if o == nil || IsNil(o.EnvironmentComponent) {
		var ret []EnvironmentComponentJsonhalTemplateEnvironmentOutput
		return ret
	}
	return o.EnvironmentComponent
}

// GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) GetEnvironmentComponentOk() ([]EnvironmentComponentJsonhalTemplateEnvironmentOutput, bool) {
	if o == nil || IsNil(o.EnvironmentComponent) {
		return nil, false
	}
	return o.EnvironmentComponent, true
}

// HasEnvironmentComponent returns a boolean if a field has been set.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) HasEnvironmentComponent() bool {
	if o != nil && !IsNil(o.EnvironmentComponent) {
		return true
	}

	return false
}

// SetEnvironmentComponent gets a reference to the given []EnvironmentComponentJsonhalTemplateEnvironmentOutput and assigns it to the EnvironmentComponent field.
func (o *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) SetEnvironmentComponent(v []EnvironmentComponentJsonhalTemplateEnvironmentOutput) {
	o.EnvironmentComponent = v
}

func (o TemplateEnvironmentJsonhalTemplateEnvironmentOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateEnvironmentJsonhalTemplateEnvironmentOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if o.TemplateProject.IsSet() {
		toSerialize["templateProject"] = o.TemplateProject.Get()
	}
	// skip: id is readOnly
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.EnvVar) {
		toSerialize["envVar"] = o.EnvVar
	}
	if !IsNil(o.EnvironmentComponent) {
		toSerialize["environmentComponent"] = o.EnvironmentComponent
	}
	return toSerialize, nil
}

type NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput struct {
	value *TemplateEnvironmentJsonhalTemplateEnvironmentOutput
	isSet bool
}

func (v NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput) Get() *TemplateEnvironmentJsonhalTemplateEnvironmentOutput {
	return v.value
}

func (v *NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput) Set(val *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput(val *TemplateEnvironmentJsonhalTemplateEnvironmentOutput) *NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput {
	return &NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput{value: val, isSet: true}
}

func (v NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateEnvironmentJsonhalTemplateEnvironmentOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
