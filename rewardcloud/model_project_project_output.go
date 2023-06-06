/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
	"time"
)

// checks if the ProjectProjectOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectProjectOutput{}

// ProjectProjectOutput Class Project
type ProjectProjectOutput struct {
	TemplateIri           NullableString                             `json:"templateIri,omitempty"`
	State                 NullableString                             `json:"state,omitempty"`
	Git                   NullableProjectProjectOutputGit            `json:"git,omitempty"`
	ServiceAccount        NullableProjectProjectOutputServiceAccount `json:"serviceAccount,omitempty"`
	Team                  NullableString                             `json:"team,omitempty"`
	Environment           []string                                   `json:"environment,omitempty"`
	Id                    *int32                                     `json:"id,omitempty"`
	Uuid                  *string                                    `json:"uuid,omitempty"`
	Name                  NullableString                             `json:"name,omitempty"`
	CodeName              NullableString                             `json:"codeName,omitempty"`
	IsActive              NullableBool                               `json:"isActive,omitempty"`
	Cpu                   NullableInt32                              `json:"cpu,omitempty"`
	Memory                NullableInt32                              `json:"memory,omitempty"`
	Storage               NullableInt32                              `json:"storage,omitempty"`
	Code                  NullableString                             `json:"code,omitempty"`
	Color                 NullableString                             `json:"color,omitempty"`
	IsInitProjectSkeleton NullableBool                               `json:"isInitProjectSkeleton,omitempty"`
	ComponentVersion      []string                                   `json:"componentVersion,omitempty"`
	ProjectTypeVersion    NullableString                             `json:"projectTypeVersion,omitempty"`
	ProjectEnvVar         []ProjectEnvVarProjectOutput               `json:"projectEnvVar,omitempty"`
	CreatedAt             *time.Time                                 `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time                                 `json:"updatedAt,omitempty"`
}

// NewProjectProjectOutput instantiates a new ProjectProjectOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectOutput() *ProjectProjectOutput {
	this := ProjectProjectOutput{}
	return &this
}

// NewProjectProjectOutputWithDefaults instantiates a new ProjectProjectOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectOutputWithDefaults() *ProjectProjectOutput {
	this := ProjectProjectOutput{}
	return &this
}

// GetTemplateIri returns the TemplateIri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetTemplateIri() string {
	if o == nil || IsNil(o.TemplateIri.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateIri.Get()
}

// GetTemplateIriOk returns a tuple with the TemplateIri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetTemplateIriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateIri.Get(), o.TemplateIri.IsSet()
}

// HasTemplateIri returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasTemplateIri() bool {
	if o != nil && o.TemplateIri.IsSet() {
		return true
	}

	return false
}

// SetTemplateIri gets a reference to the given NullableString and assigns it to the TemplateIri field.
func (o *ProjectProjectOutput) SetTemplateIri(v string) {
	o.TemplateIri.Set(&v)
}

// SetTemplateIriNil sets the value for TemplateIri to be an explicit nil
func (o *ProjectProjectOutput) SetTemplateIriNil() {
	o.TemplateIri.Set(nil)
}

// UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
func (o *ProjectProjectOutput) UnsetTemplateIri() {
	o.TemplateIri.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *ProjectProjectOutput) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil
func (o *ProjectProjectOutput) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ProjectProjectOutput) UnsetState() {
	o.State.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetGit() ProjectProjectOutputGit {
	if o == nil || IsNil(o.Git.Get()) {
		var ret ProjectProjectOutputGit
		return ret
	}
	return *o.Git.Get()
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetGitOk() (*ProjectProjectOutputGit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Git.Get(), o.Git.IsSet()
}

// HasGit returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasGit() bool {
	if o != nil && o.Git.IsSet() {
		return true
	}

	return false
}

// SetGit gets a reference to the given NullableProjectProjectOutputGit and assigns it to the Git field.
func (o *ProjectProjectOutput) SetGit(v ProjectProjectOutputGit) {
	o.Git.Set(&v)
}

// SetGitNil sets the value for Git to be an explicit nil
func (o *ProjectProjectOutput) SetGitNil() {
	o.Git.Set(nil)
}

// UnsetGit ensures that no value is present for Git, not even an explicit nil
func (o *ProjectProjectOutput) UnsetGit() {
	o.Git.Unset()
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetServiceAccount() ProjectProjectOutputServiceAccount {
	if o == nil || IsNil(o.ServiceAccount.Get()) {
		var ret ProjectProjectOutputServiceAccount
		return ret
	}
	return *o.ServiceAccount.Get()
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetServiceAccountOk() (*ProjectProjectOutputServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceAccount.Get(), o.ServiceAccount.IsSet()
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount.IsSet() {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given NullableProjectProjectOutputServiceAccount and assigns it to the ServiceAccount field.
func (o *ProjectProjectOutput) SetServiceAccount(v ProjectProjectOutputServiceAccount) {
	o.ServiceAccount.Set(&v)
}

// SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil
func (o *ProjectProjectOutput) SetServiceAccountNil() {
	o.ServiceAccount.Set(nil)
}

// UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
func (o *ProjectProjectOutput) UnsetServiceAccount() {
	o.ServiceAccount.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetTeam() string {
	if o == nil || IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *ProjectProjectOutput) SetTeam(v string) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *ProjectProjectOutput) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *ProjectProjectOutput) UnsetTeam() {
	o.Team.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetEnvironment() []string {
	if o == nil || IsNil(o.Environment) {
		var ret []string
		return ret
	}
	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetEnvironmentOk() ([]string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given []string and assigns it to the Environment field.
func (o *ProjectProjectOutput) SetEnvironment(v []string) {
	o.Environment = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectProjectOutput) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ProjectProjectOutput) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectProjectOutput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectProjectOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectProjectOutput) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetCodeName() string {
	if o == nil || IsNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *ProjectProjectOutput) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *ProjectProjectOutput) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *ProjectProjectOutput) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *ProjectProjectOutput) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}

// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *ProjectProjectOutput) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *ProjectProjectOutput) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *ProjectProjectOutput) SetCpu(v int32) {
	o.Cpu.Set(&v)
}

// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *ProjectProjectOutput) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *ProjectProjectOutput) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *ProjectProjectOutput) SetMemory(v int32) {
	o.Memory.Set(&v)
}

// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *ProjectProjectOutput) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *ProjectProjectOutput) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetStorage() int32 {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *ProjectProjectOutput) SetStorage(v int32) {
	o.Storage.Set(&v)
}

// SetStorageNil sets the value for Storage to be an explicit nil
func (o *ProjectProjectOutput) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *ProjectProjectOutput) UnsetStorage() {
	o.Storage.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *ProjectProjectOutput) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *ProjectProjectOutput) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ProjectProjectOutput) UnsetCode() {
	o.Code.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *ProjectProjectOutput) SetColor(v string) {
	o.Color.Set(&v)
}

// SetColorNil sets the value for Color to be an explicit nil
func (o *ProjectProjectOutput) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *ProjectProjectOutput) UnsetColor() {
	o.Color.Unset()
}

// GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetIsInitProjectSkeleton() bool {
	if o == nil || IsNil(o.IsInitProjectSkeleton.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitProjectSkeleton.Get()
}

// GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetIsInitProjectSkeletonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitProjectSkeleton.Get(), o.IsInitProjectSkeleton.IsSet()
}

// HasIsInitProjectSkeleton returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasIsInitProjectSkeleton() bool {
	if o != nil && o.IsInitProjectSkeleton.IsSet() {
		return true
	}

	return false
}

// SetIsInitProjectSkeleton gets a reference to the given NullableBool and assigns it to the IsInitProjectSkeleton field.
func (o *ProjectProjectOutput) SetIsInitProjectSkeleton(v bool) {
	o.IsInitProjectSkeleton.Set(&v)
}

// SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil
func (o *ProjectProjectOutput) SetIsInitProjectSkeletonNil() {
	o.IsInitProjectSkeleton.Set(nil)
}

// UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
func (o *ProjectProjectOutput) UnsetIsInitProjectSkeleton() {
	o.IsInitProjectSkeleton.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetComponentVersion() []string {
	if o == nil || IsNil(o.ComponentVersion) {
		var ret []string
		return ret
	}
	return o.ComponentVersion
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetComponentVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ComponentVersion) {
		return nil, false
	}
	return o.ComponentVersion, true
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasComponentVersion() bool {
	if o != nil && !IsNil(o.ComponentVersion) {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given []string and assigns it to the ComponentVersion field.
func (o *ProjectProjectOutput) SetComponentVersion(v []string) {
	o.ComponentVersion = v
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectOutput) GetProjectTypeVersion() string {
	if o == nil || IsNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectOutput) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *ProjectProjectOutput) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}

// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *ProjectProjectOutput) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *ProjectProjectOutput) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetProjectEnvVar returns the ProjectEnvVar field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetProjectEnvVar() []ProjectEnvVarProjectOutput {
	if o == nil || IsNil(o.ProjectEnvVar) {
		var ret []ProjectEnvVarProjectOutput
		return ret
	}
	return o.ProjectEnvVar
}

// GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetProjectEnvVarOk() ([]ProjectEnvVarProjectOutput, bool) {
	if o == nil || IsNil(o.ProjectEnvVar) {
		return nil, false
	}
	return o.ProjectEnvVar, true
}

// HasProjectEnvVar returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasProjectEnvVar() bool {
	if o != nil && !IsNil(o.ProjectEnvVar) {
		return true
	}

	return false
}

// SetProjectEnvVar gets a reference to the given []ProjectEnvVarProjectOutput and assigns it to the ProjectEnvVar field.
func (o *ProjectProjectOutput) SetProjectEnvVar(v []ProjectEnvVarProjectOutput) {
	o.ProjectEnvVar = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectProjectOutput) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectProjectOutput) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectOutput) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectProjectOutput) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectProjectOutput) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProjectProjectOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectProjectOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TemplateIri.IsSet() {
		toSerialize["templateIri"] = o.TemplateIri.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Git.IsSet() {
		toSerialize["git"] = o.Git.Get()
	}
	if o.ServiceAccount.IsSet() {
		toSerialize["serviceAccount"] = o.ServiceAccount.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	// skip: uuid is readOnly
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CodeName.IsSet() {
		toSerialize["codeName"] = o.CodeName.Get()
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableProjectProjectOutput struct {
	value *ProjectProjectOutput
	isSet bool
}

func (v NullableProjectProjectOutput) Get() *ProjectProjectOutput {
	return v.value
}

func (v *NullableProjectProjectOutput) Set(val *ProjectProjectOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectOutput(val *ProjectProjectOutput) *NullableProjectProjectOutput {
	return &NullableProjectProjectOutput{value: val, isSet: true}
}

func (v NullableProjectProjectOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}