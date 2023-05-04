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

// checks if the TemplateProjectJsonhalTemplateProjectInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateProjectJsonhalTemplateProjectInput{}

// TemplateProjectJsonhalTemplateProjectInput Class ProjectTemplate
type TemplateProjectJsonhalTemplateProjectInput struct {
	Links *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Description NullableString `json:"description,omitempty"`
	TemplateEnvironment []string `json:"templateEnvironment,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IsActive NullableBool `json:"isActive,omitempty"`
	Cpu NullableInt32 `json:"cpu,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Storage NullableInt32 `json:"storage,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Color NullableString `json:"color,omitempty"`
	IsInitProjectSkeleton NullableBool `json:"isInitProjectSkeleton,omitempty"`
	ComponentVersion []string `json:"componentVersion,omitempty"`
	ProjectTypeVersion NullableString `json:"projectTypeVersion,omitempty"`
	ProjectEnvVar []ProjectEnvVarJsonhalTemplateProjectInput `json:"projectEnvVar,omitempty"`
}

// NewTemplateProjectJsonhalTemplateProjectInput instantiates a new TemplateProjectJsonhalTemplateProjectInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateProjectJsonhalTemplateProjectInput() *TemplateProjectJsonhalTemplateProjectInput {
	this := TemplateProjectJsonhalTemplateProjectInput{}
	return &this
}

// NewTemplateProjectJsonhalTemplateProjectInputWithDefaults instantiates a new TemplateProjectJsonhalTemplateProjectInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateProjectJsonhalTemplateProjectInputWithDefaults() *TemplateProjectJsonhalTemplateProjectInput {
	this := TemplateProjectJsonhalTemplateProjectInput{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetDescription() {
	o.Description.Unset()
}

// GetTemplateEnvironment returns the TemplateEnvironment field value if set, zero value otherwise.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetTemplateEnvironment() []string {
	if o == nil || IsNil(o.TemplateEnvironment) {
		var ret []string
		return ret
	}
	return o.TemplateEnvironment
}

// GetTemplateEnvironmentOk returns a tuple with the TemplateEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetTemplateEnvironmentOk() ([]string, bool) {
	if o == nil || IsNil(o.TemplateEnvironment) {
		return nil, false
	}
	return o.TemplateEnvironment, true
}

// HasTemplateEnvironment returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasTemplateEnvironment() bool {
	if o != nil && !IsNil(o.TemplateEnvironment) {
		return true
	}

	return false
}

// SetTemplateEnvironment gets a reference to the given []string and assigns it to the TemplateEnvironment field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetTemplateEnvironment(v []string) {
	o.TemplateEnvironment = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetName() {
	o.Name.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetStorage() int32 {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetStorage() {
	o.Storage.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetCode() {
	o.Code.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetColor() {
	o.Color.Unset()
}

// GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsInitProjectSkeleton() bool {
	if o == nil || IsNil(o.IsInitProjectSkeleton.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitProjectSkeleton.Get()
}

// GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetIsInitProjectSkeletonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitProjectSkeleton.Get(), o.IsInitProjectSkeleton.IsSet()
}

// HasIsInitProjectSkeleton returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasIsInitProjectSkeleton() bool {
	if o != nil && o.IsInitProjectSkeleton.IsSet() {
		return true
	}

	return false
}

// SetIsInitProjectSkeleton gets a reference to the given NullableBool and assigns it to the IsInitProjectSkeleton field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsInitProjectSkeleton(v bool) {
	o.IsInitProjectSkeleton.Set(&v)
}
// SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetIsInitProjectSkeletonNil() {
	o.IsInitProjectSkeleton.Set(nil)
}

// UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetIsInitProjectSkeleton() {
	o.IsInitProjectSkeleton.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetComponentVersion() []string {
	if o == nil || IsNil(o.ComponentVersion) {
		var ret []string
		return ret
	}
	return o.ComponentVersion
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetComponentVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ComponentVersion) {
		return nil, false
	}
	return o.ComponentVersion, true
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasComponentVersion() bool {
	if o != nil && !IsNil(o.ComponentVersion) {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given []string and assigns it to the ComponentVersion field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetComponentVersion(v []string) {
	o.ComponentVersion = v
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectTypeVersion() string {
	if o == nil || IsNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}
// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *TemplateProjectJsonhalTemplateProjectInput) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetProjectEnvVar returns the ProjectEnvVar field value if set, zero value otherwise.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectEnvVar() []ProjectEnvVarJsonhalTemplateProjectInput {
	if o == nil || IsNil(o.ProjectEnvVar) {
		var ret []ProjectEnvVarJsonhalTemplateProjectInput
		return ret
	}
	return o.ProjectEnvVar
}

// GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) GetProjectEnvVarOk() ([]ProjectEnvVarJsonhalTemplateProjectInput, bool) {
	if o == nil || IsNil(o.ProjectEnvVar) {
		return nil, false
	}
	return o.ProjectEnvVar, true
}

// HasProjectEnvVar returns a boolean if a field has been set.
func (o *TemplateProjectJsonhalTemplateProjectInput) HasProjectEnvVar() bool {
	if o != nil && !IsNil(o.ProjectEnvVar) {
		return true
	}

	return false
}

// SetProjectEnvVar gets a reference to the given []ProjectEnvVarJsonhalTemplateProjectInput and assigns it to the ProjectEnvVar field.
func (o *TemplateProjectJsonhalTemplateProjectInput) SetProjectEnvVar(v []ProjectEnvVarJsonhalTemplateProjectInput) {
	o.ProjectEnvVar = v
}

func (o TemplateProjectJsonhalTemplateProjectInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateProjectJsonhalTemplateProjectInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
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

type NullableTemplateProjectJsonhalTemplateProjectInput struct {
	value *TemplateProjectJsonhalTemplateProjectInput
	isSet bool
}

func (v NullableTemplateProjectJsonhalTemplateProjectInput) Get() *TemplateProjectJsonhalTemplateProjectInput {
	return v.value
}

func (v *NullableTemplateProjectJsonhalTemplateProjectInput) Set(val *TemplateProjectJsonhalTemplateProjectInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateProjectJsonhalTemplateProjectInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateProjectJsonhalTemplateProjectInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateProjectJsonhalTemplateProjectInput(val *TemplateProjectJsonhalTemplateProjectInput) *NullableTemplateProjectJsonhalTemplateProjectInput {
	return &NullableTemplateProjectJsonhalTemplateProjectInput{value: val, isSet: true}
}

func (v NullableTemplateProjectJsonhalTemplateProjectInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateProjectJsonhalTemplateProjectInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


