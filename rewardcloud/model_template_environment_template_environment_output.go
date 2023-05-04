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

// TemplateEnvironmentTemplateEnvironmentOutput Class TemplateEnvironment
type TemplateEnvironmentTemplateEnvironmentOutput struct {
	TemplateProject       NullableString                                  `json:"templateProject,omitempty"`
	Id                    *int32                                          `json:"id,omitempty"`
	Uuid                  NullableString                                  `json:"uuid,omitempty"`
	Name                  NullableString                                  `json:"name,omitempty"`
	Cpu                   NullableInt32                                   `json:"cpu,omitempty"`
	Memory                NullableInt32                                   `json:"memory,omitempty"`
	Storage               NullableInt32                                   `json:"storage,omitempty"`
	DataTransferSettings  NullableString                                  `json:"dataTransferSettings,omitempty"`
	IsStripDatabase       NullableBool                                    `json:"isStripDatabase,omitempty"`
	IsAllowOutgoingEmails NullableBool                                    `json:"isAllowOutgoingEmails,omitempty"`
	IsInitSampleData      NullableBool                                    `json:"isInitSampleData,omitempty"`
	CreatedAt             *time.Time                                      `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time                                      `json:"updatedAt,omitempty"`
	EnvVar                []EnvironmentEnvVarTemplateEnvironmentOutput    `json:"envVar,omitempty"`
	EnvironmentComponent  []EnvironmentComponentTemplateEnvironmentOutput `json:"environmentComponent,omitempty"`
}

// NewTemplateEnvironmentTemplateEnvironmentOutput instantiates a new TemplateEnvironmentTemplateEnvironmentOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateEnvironmentTemplateEnvironmentOutput() *TemplateEnvironmentTemplateEnvironmentOutput {
	this := TemplateEnvironmentTemplateEnvironmentOutput{}
	return &this
}

// NewTemplateEnvironmentTemplateEnvironmentOutputWithDefaults instantiates a new TemplateEnvironmentTemplateEnvironmentOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateEnvironmentTemplateEnvironmentOutputWithDefaults() *TemplateEnvironmentTemplateEnvironmentOutput {
	this := TemplateEnvironmentTemplateEnvironmentOutput{}
	return &this
}

// GetTemplateProject returns the TemplateProject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetTemplateProject() string {
	if o == nil || isNil(o.TemplateProject.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateProject.Get()
}

// GetTemplateProjectOk returns a tuple with the TemplateProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetTemplateProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateProject.Get(), o.TemplateProject.IsSet()
}

// HasTemplateProject returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasTemplateProject() bool {
	if o != nil && o.TemplateProject.IsSet() {
		return true
	}

	return false
}

// SetTemplateProject gets a reference to the given NullableString and assigns it to the TemplateProject field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetTemplateProject(v string) {
	o.TemplateProject.Set(&v)
}

// SetTemplateProjectNil sets the value for TemplateProject to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetTemplateProjectNil() {
	o.TemplateProject.Set(nil)
}

// UnsetTemplateProject ensures that no value is present for TemplateProject, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetTemplateProject() {
	o.TemplateProject.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCpu() int32 {
	if o == nil || isNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetCpu(v int32) {
	o.Cpu.Set(&v)
}

// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetMemory() int32 {
	if o == nil || isNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetStorage() int32 {
	if o == nil || isNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetStorage(v int32) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetStorage() {
	o.Storage.Unset()
}

// GetDataTransferSettings returns the DataTransferSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetDataTransferSettings() string {
	if o == nil || isNil(o.DataTransferSettings.Get()) {
		var ret string
		return ret
	}
	return *o.DataTransferSettings.Get()
}

// GetDataTransferSettingsOk returns a tuple with the DataTransferSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetDataTransferSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataTransferSettings.Get(), o.DataTransferSettings.IsSet()
}

// HasDataTransferSettings returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasDataTransferSettings() bool {
	if o != nil && o.DataTransferSettings.IsSet() {
		return true
	}

	return false
}

// SetDataTransferSettings gets a reference to the given NullableString and assigns it to the DataTransferSettings field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetDataTransferSettings(v string) {
	o.DataTransferSettings.Set(&v)
}

// SetDataTransferSettingsNil sets the value for DataTransferSettings to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetDataTransferSettingsNil() {
	o.DataTransferSettings.Set(nil)
}

// UnsetDataTransferSettings ensures that no value is present for DataTransferSettings, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetDataTransferSettings() {
	o.DataTransferSettings.Unset()
}

// GetIsStripDatabase returns the IsStripDatabase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsStripDatabase() bool {
	if o == nil || isNil(o.IsStripDatabase.Get()) {
		var ret bool
		return ret
	}
	return *o.IsStripDatabase.Get()
}

// GetIsStripDatabaseOk returns a tuple with the IsStripDatabase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsStripDatabaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsStripDatabase.Get(), o.IsStripDatabase.IsSet()
}

// HasIsStripDatabase returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasIsStripDatabase() bool {
	if o != nil && o.IsStripDatabase.IsSet() {
		return true
	}

	return false
}

// SetIsStripDatabase gets a reference to the given NullableBool and assigns it to the IsStripDatabase field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsStripDatabase(v bool) {
	o.IsStripDatabase.Set(&v)
}

// SetIsStripDatabaseNil sets the value for IsStripDatabase to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsStripDatabaseNil() {
	o.IsStripDatabase.Set(nil)
}

// UnsetIsStripDatabase ensures that no value is present for IsStripDatabase, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetIsStripDatabase() {
	o.IsStripDatabase.Unset()
}

// GetIsAllowOutgoingEmails returns the IsAllowOutgoingEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsAllowOutgoingEmails() bool {
	if o == nil || isNil(o.IsAllowOutgoingEmails.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAllowOutgoingEmails.Get()
}

// GetIsAllowOutgoingEmailsOk returns a tuple with the IsAllowOutgoingEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsAllowOutgoingEmailsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAllowOutgoingEmails.Get(), o.IsAllowOutgoingEmails.IsSet()
}

// HasIsAllowOutgoingEmails returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasIsAllowOutgoingEmails() bool {
	if o != nil && o.IsAllowOutgoingEmails.IsSet() {
		return true
	}

	return false
}

// SetIsAllowOutgoingEmails gets a reference to the given NullableBool and assigns it to the IsAllowOutgoingEmails field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsAllowOutgoingEmails(v bool) {
	o.IsAllowOutgoingEmails.Set(&v)
}

// SetIsAllowOutgoingEmailsNil sets the value for IsAllowOutgoingEmails to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsAllowOutgoingEmailsNil() {
	o.IsAllowOutgoingEmails.Set(nil)
}

// UnsetIsAllowOutgoingEmails ensures that no value is present for IsAllowOutgoingEmails, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetIsAllowOutgoingEmails() {
	o.IsAllowOutgoingEmails.Unset()
}

// GetIsInitSampleData returns the IsInitSampleData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsInitSampleData() bool {
	if o == nil || isNil(o.IsInitSampleData.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitSampleData.Get()
}

// GetIsInitSampleDataOk returns a tuple with the IsInitSampleData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetIsInitSampleDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitSampleData.Get(), o.IsInitSampleData.IsSet()
}

// HasIsInitSampleData returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasIsInitSampleData() bool {
	if o != nil && o.IsInitSampleData.IsSet() {
		return true
	}

	return false
}

// SetIsInitSampleData gets a reference to the given NullableBool and assigns it to the IsInitSampleData field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsInitSampleData(v bool) {
	o.IsInitSampleData.Set(&v)
}

// SetIsInitSampleDataNil sets the value for IsInitSampleData to be an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetIsInitSampleDataNil() {
	o.IsInitSampleData.Set(nil)
}

// UnsetIsInitSampleData ensures that no value is present for IsInitSampleData, not even an explicit nil
func (o *TemplateEnvironmentTemplateEnvironmentOutput) UnsetIsInitSampleData() {
	o.IsInitSampleData.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvVar returns the EnvVar field value if set, zero value otherwise.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvVar() []EnvironmentEnvVarTemplateEnvironmentOutput {
	if o == nil || isNil(o.EnvVar) {
		var ret []EnvironmentEnvVarTemplateEnvironmentOutput
		return ret
	}
	return o.EnvVar
}

// GetEnvVarOk returns a tuple with the EnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvVarOk() ([]EnvironmentEnvVarTemplateEnvironmentOutput, bool) {
	if o == nil || isNil(o.EnvVar) {
		return nil, false
	}
	return o.EnvVar, true
}

// HasEnvVar returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasEnvVar() bool {
	if o != nil && !isNil(o.EnvVar) {
		return true
	}

	return false
}

// SetEnvVar gets a reference to the given []EnvironmentEnvVarTemplateEnvironmentOutput and assigns it to the EnvVar field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetEnvVar(v []EnvironmentEnvVarTemplateEnvironmentOutput) {
	o.EnvVar = v
}

// GetEnvironmentComponent returns the EnvironmentComponent field value if set, zero value otherwise.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvironmentComponent() []EnvironmentComponentTemplateEnvironmentOutput {
	if o == nil || isNil(o.EnvironmentComponent) {
		var ret []EnvironmentComponentTemplateEnvironmentOutput
		return ret
	}
	return o.EnvironmentComponent
}

// GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) GetEnvironmentComponentOk() ([]EnvironmentComponentTemplateEnvironmentOutput, bool) {
	if o == nil || isNil(o.EnvironmentComponent) {
		return nil, false
	}
	return o.EnvironmentComponent, true
}

// HasEnvironmentComponent returns a boolean if a field has been set.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) HasEnvironmentComponent() bool {
	if o != nil && !isNil(o.EnvironmentComponent) {
		return true
	}

	return false
}

// SetEnvironmentComponent gets a reference to the given []EnvironmentComponentTemplateEnvironmentOutput and assigns it to the EnvironmentComponent field.
func (o *TemplateEnvironmentTemplateEnvironmentOutput) SetEnvironmentComponent(v []EnvironmentComponentTemplateEnvironmentOutput) {
	o.EnvironmentComponent = v
}

func (o TemplateEnvironmentTemplateEnvironmentOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateProject.IsSet() {
		toSerialize["templateProject"] = o.TemplateProject.Get()
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
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.EnvVar) {
		toSerialize["envVar"] = o.EnvVar
	}
	if !isNil(o.EnvironmentComponent) {
		toSerialize["environmentComponent"] = o.EnvironmentComponent
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateEnvironmentTemplateEnvironmentOutput struct {
	value *TemplateEnvironmentTemplateEnvironmentOutput
	isSet bool
}

func (v NullableTemplateEnvironmentTemplateEnvironmentOutput) Get() *TemplateEnvironmentTemplateEnvironmentOutput {
	return v.value
}

func (v *NullableTemplateEnvironmentTemplateEnvironmentOutput) Set(val *TemplateEnvironmentTemplateEnvironmentOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateEnvironmentTemplateEnvironmentOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateEnvironmentTemplateEnvironmentOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateEnvironmentTemplateEnvironmentOutput(val *TemplateEnvironmentTemplateEnvironmentOutput) *NullableTemplateEnvironmentTemplateEnvironmentOutput {
	return &NullableTemplateEnvironmentTemplateEnvironmentOutput{value: val, isSet: true}
}

func (v NullableTemplateEnvironmentTemplateEnvironmentOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateEnvironmentTemplateEnvironmentOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
