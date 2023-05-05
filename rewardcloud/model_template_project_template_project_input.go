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

// checks if the TemplateProjectTemplateProjectInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateProjectTemplateProjectInput{}

// TemplateProjectTemplateProjectInput Class ProjectTemplate
type TemplateProjectTemplateProjectInput struct {
	Description           NullableString                      `json:"description,omitempty"`
	TemplateEnvironment   []string                            `json:"templateEnvironment,omitempty"`
	Name                  NullableString                      `json:"name,omitempty"`
	IsActive              NullableBool                        `json:"isActive,omitempty"`
	Cpu                   NullableInt32                       `json:"cpu,omitempty"`
	Memory                NullableInt32                       `json:"memory,omitempty"`
	Storage               NullableInt32                       `json:"storage,omitempty"`
	Code                  NullableString                      `json:"code,omitempty"`
	Color                 NullableString                      `json:"color,omitempty"`
	IsInitProjectSkeleton NullableBool                        `json:"isInitProjectSkeleton,omitempty"`
	ComponentVersion      []string                            `json:"componentVersion,omitempty"`
	ProjectTypeVersion    NullableString                      `json:"projectTypeVersion,omitempty"`
	ProjectEnvVar         []ProjectEnvVarTemplateProjectInput `json:"projectEnvVar,omitempty"`
}

// NewTemplateProjectTemplateProjectInput instantiates a new TemplateProjectTemplateProjectInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateProjectTemplateProjectInput() *TemplateProjectTemplateProjectInput {
	this := TemplateProjectTemplateProjectInput{}
	return &this
}

// NewTemplateProjectTemplateProjectInputWithDefaults instantiates a new TemplateProjectTemplateProjectInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateProjectTemplateProjectInputWithDefaults() *TemplateProjectTemplateProjectInput {
	this := TemplateProjectTemplateProjectInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TemplateProjectTemplateProjectInput) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetDescription() {
	o.Description.Unset()
}

// GetTemplateEnvironment returns the TemplateEnvironment field value if set, zero value otherwise.
func (o *TemplateProjectTemplateProjectInput) GetTemplateEnvironment() []string {
	if o == nil || IsNil(o.TemplateEnvironment) {
		var ret []string
		return ret
	}
	return o.TemplateEnvironment
}

// GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectTemplateProjectInput) GetTemplateEnvironmentOk() ([]string, bool) {
	if o == nil || IsNil(o.TemplateEnvironment) {
		return nil, false
	}
	return o.TemplateEnvironment, true
}

// HasTemplateEnvironment returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasTemplateEnvironment() bool {
	if o != nil && !IsNil(o.TemplateEnvironment) {
		return true
	}

	return false
}

// SetTemplateEnvironment gets a reference to the given []string and assigns it to the TemplateEnvironment field.
func (o *TemplateProjectTemplateProjectInput) SetTemplateEnvironment(v []string) {
	o.TemplateEnvironment = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TemplateProjectTemplateProjectInput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetName() {
	o.Name.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *TemplateProjectTemplateProjectInput) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}

// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *TemplateProjectTemplateProjectInput) SetCpu(v int32) {
	o.Cpu.Set(&v)
}

// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *TemplateProjectTemplateProjectInput) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetStorage() int32 {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *TemplateProjectTemplateProjectInput) SetStorage(v int32) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetStorage() {
	o.Storage.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TemplateProjectTemplateProjectInput) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetCode() {
	o.Code.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *TemplateProjectTemplateProjectInput) SetColor(v string) {
	o.Color.Set(&v)
}

// SetColorNil sets the value for Color to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetColor() {
	o.Color.Unset()
}

// GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetIsInitProjectSkeleton() bool {
	if o == nil || IsNil(o.IsInitProjectSkeleton.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitProjectSkeleton.Get()
}

// GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetIsInitProjectSkeletonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitProjectSkeleton.Get(), o.IsInitProjectSkeleton.IsSet()
}

// HasIsInitProjectSkeleton returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasIsInitProjectSkeleton() bool {
	if o != nil && o.IsInitProjectSkeleton.IsSet() {
		return true
	}

	return false
}

// SetIsInitProjectSkeleton gets a reference to the given NullableBool and assigns it to the IsInitProjectSkeleton field.
func (o *TemplateProjectTemplateProjectInput) SetIsInitProjectSkeleton(v bool) {
	o.IsInitProjectSkeleton.Set(&v)
}

// SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetIsInitProjectSkeletonNil() {
	o.IsInitProjectSkeleton.Set(nil)
}

// UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetIsInitProjectSkeleton() {
	o.IsInitProjectSkeleton.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise.
func (o *TemplateProjectTemplateProjectInput) GetComponentVersion() []string {
	if o == nil || IsNil(o.ComponentVersion) {
		var ret []string
		return ret
	}
	return o.ComponentVersion
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectTemplateProjectInput) GetComponentVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ComponentVersion) {
		return nil, false
	}
	return o.ComponentVersion, true
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasComponentVersion() bool {
	if o != nil && !IsNil(o.ComponentVersion) {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given []string and assigns it to the ComponentVersion field.
func (o *TemplateProjectTemplateProjectInput) SetComponentVersion(v []string) {
	o.ComponentVersion = v
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectTemplateProjectInput) GetProjectTypeVersion() string {
	if o == nil || IsNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectTemplateProjectInput) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *TemplateProjectTemplateProjectInput) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}

// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *TemplateProjectTemplateProjectInput) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *TemplateProjectTemplateProjectInput) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetProjectEnvVar returns the ProjectEnvVar field value if set, zero value otherwise.
func (o *TemplateProjectTemplateProjectInput) GetProjectEnvVar() []ProjectEnvVarTemplateProjectInput {
	if o == nil || IsNil(o.ProjectEnvVar) {
		var ret []ProjectEnvVarTemplateProjectInput
		return ret
	}
	return o.ProjectEnvVar
}

// GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectTemplateProjectInput) GetProjectEnvVarOk() ([]ProjectEnvVarTemplateProjectInput, bool) {
	if o == nil || IsNil(o.ProjectEnvVar) {
		return nil, false
	}
	return o.ProjectEnvVar, true
}

// HasProjectEnvVar returns a boolean if a field has been set.
func (o *TemplateProjectTemplateProjectInput) HasProjectEnvVar() bool {
	if o != nil && !IsNil(o.ProjectEnvVar) {
		return true
	}

	return false
}

// SetProjectEnvVar gets a reference to the given []ProjectEnvVarTemplateProjectInput and assigns it to the ProjectEnvVar field.
func (o *TemplateProjectTemplateProjectInput) SetProjectEnvVar(v []ProjectEnvVarTemplateProjectInput) {
	o.ProjectEnvVar = v
}

func (o TemplateProjectTemplateProjectInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateProjectTemplateProjectInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.TemplateEnvironment) {
		toSerialize["templateEnvironment"] = o.TemplateEnvironment
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
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
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.IsInitProjectSkeleton.IsSet() {
		toSerialize["isInitProjectSkeleton"] = o.IsInitProjectSkeleton.Get()
	}
	if !IsNil(o.ComponentVersion) {
		toSerialize["componentVersion"] = o.ComponentVersion
	}
	if o.ProjectTypeVersion.IsSet() {
		toSerialize["projectTypeVersion"] = o.ProjectTypeVersion.Get()
	}
	if !IsNil(o.ProjectEnvVar) {
		toSerialize["projectEnvVar"] = o.ProjectEnvVar
	}
	return toSerialize, nil
}

type NullableTemplateProjectTemplateProjectInput struct {
	value *TemplateProjectTemplateProjectInput
	isSet bool
}

func (v NullableTemplateProjectTemplateProjectInput) Get() *TemplateProjectTemplateProjectInput {
	return v.value
}

func (v *NullableTemplateProjectTemplateProjectInput) Set(val *TemplateProjectTemplateProjectInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateProjectTemplateProjectInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateProjectTemplateProjectInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateProjectTemplateProjectInput(val *TemplateProjectTemplateProjectInput) *NullableTemplateProjectTemplateProjectInput {
	return &NullableTemplateProjectTemplateProjectInput{value: val, isSet: true}
}

func (v NullableTemplateProjectTemplateProjectInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateProjectTemplateProjectInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
