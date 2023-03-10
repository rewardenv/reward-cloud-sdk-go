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

// ProjectJsonhalProjectGet Class Project
type ProjectJsonhalProjectGet struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	Cpu NullableInt32 `json:"cpu,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Storage NullableInt32 `json:"storage,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Color NullableString `json:"color,omitempty"`
	IsInitProjectSkeleton NullableBool `json:"isInitProjectSkeleton,omitempty"`
	Environment []string `json:"environment,omitempty"`
	Team NullableString `json:"team,omitempty"`
	Git NullableProjectJsonhalProjectGetGit `json:"git,omitempty"`
	ServiceAccount NullableProjectJsonhalProjectGetServiceAccount `json:"serviceAccount,omitempty"`
	State NullableString `json:"state,omitempty"`
	ProjectEnvVar []ProjectEnvVarJsonhalProjectGet `json:"projectEnvVar,omitempty"`
	ProjectTypeVersion NullableString `json:"projectTypeVersion,omitempty"`
	ComponentVersion []string `json:"componentVersion,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewProjectJsonhalProjectGet instantiates a new ProjectJsonhalProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectJsonhalProjectGet() *ProjectJsonhalProjectGet {
	this := ProjectJsonhalProjectGet{}
	return &this
}

// NewProjectJsonhalProjectGetWithDefaults instantiates a new ProjectJsonhalProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectJsonhalProjectGetWithDefaults() *ProjectJsonhalProjectGet {
	this := ProjectJsonhalProjectGet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *ProjectJsonhalProjectGet) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectJsonhalProjectGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ProjectJsonhalProjectGet) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectJsonhalProjectGet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetName() {
	o.Name.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ProjectJsonhalProjectGet) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetCpu() int32 {
	if o == nil || isNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetCpuOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *ProjectJsonhalProjectGet) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetMemory() int32 {
	if o == nil || isNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetMemoryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *ProjectJsonhalProjectGet) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetStorage() int32 {
	if o == nil || isNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetStorageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *ProjectJsonhalProjectGet) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetStorage() {
	o.Storage.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetCode() string {
	if o == nil || isNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *ProjectJsonhalProjectGet) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetCode() {
	o.Code.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetColor() string {
	if o == nil || isNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetColorOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *ProjectJsonhalProjectGet) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetColor() {
	o.Color.Unset()
}

// GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetIsInitProjectSkeleton() bool {
	if o == nil || isNil(o.IsInitProjectSkeleton.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitProjectSkeleton.Get()
}

// GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetIsInitProjectSkeletonOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsInitProjectSkeleton.Get(), o.IsInitProjectSkeleton.IsSet()
}

// HasIsInitProjectSkeleton returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasIsInitProjectSkeleton() bool {
	if o != nil && o.IsInitProjectSkeleton.IsSet() {
		return true
	}

	return false
}

// SetIsInitProjectSkeleton gets a reference to the given NullableBool and assigns it to the IsInitProjectSkeleton field.
func (o *ProjectJsonhalProjectGet) SetIsInitProjectSkeleton(v bool) {
	o.IsInitProjectSkeleton.Set(&v)
}
// SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetIsInitProjectSkeletonNil() {
	o.IsInitProjectSkeleton.Set(nil)
}

// UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetIsInitProjectSkeleton() {
	o.IsInitProjectSkeleton.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetEnvironment() []string {
	if o == nil || isNil(o.Environment) {
		var ret []string
		return ret
	}
	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetEnvironmentOk() ([]string, bool) {
	if o == nil || isNil(o.Environment) {
    return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given []string and assigns it to the Environment field.
func (o *ProjectJsonhalProjectGet) SetEnvironment(v []string) {
	o.Environment = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetTeam() string {
	if o == nil || isNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetTeamOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *ProjectJsonhalProjectGet) SetTeam(v string) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetTeam() {
	o.Team.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetGit() ProjectJsonhalProjectGetGit {
	if o == nil || isNil(o.Git.Get()) {
		var ret ProjectJsonhalProjectGetGit
		return ret
	}
	return *o.Git.Get()
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetGitOk() (*ProjectJsonhalProjectGetGit, bool) {
	if o == nil {
    return nil, false
	}
	return o.Git.Get(), o.Git.IsSet()
}

// HasGit returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasGit() bool {
	if o != nil && o.Git.IsSet() {
		return true
	}

	return false
}

// SetGit gets a reference to the given NullableProjectJsonhalProjectGetGit and assigns it to the Git field.
func (o *ProjectJsonhalProjectGet) SetGit(v ProjectJsonhalProjectGetGit) {
	o.Git.Set(&v)
}
// SetGitNil sets the value for Git to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetGitNil() {
	o.Git.Set(nil)
}

// UnsetGit ensures that no value is present for Git, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetGit() {
	o.Git.Unset()
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetServiceAccount() ProjectJsonhalProjectGetServiceAccount {
	if o == nil || isNil(o.ServiceAccount.Get()) {
		var ret ProjectJsonhalProjectGetServiceAccount
		return ret
	}
	return *o.ServiceAccount.Get()
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetServiceAccountOk() (*ProjectJsonhalProjectGetServiceAccount, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServiceAccount.Get(), o.ServiceAccount.IsSet()
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount.IsSet() {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given NullableProjectJsonhalProjectGetServiceAccount and assigns it to the ServiceAccount field.
func (o *ProjectJsonhalProjectGet) SetServiceAccount(v ProjectJsonhalProjectGetServiceAccount) {
	o.ServiceAccount.Set(&v)
}
// SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetServiceAccountNil() {
	o.ServiceAccount.Set(nil)
}

// UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetServiceAccount() {
	o.ServiceAccount.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetState() string {
	if o == nil || isNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetStateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *ProjectJsonhalProjectGet) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetState() {
	o.State.Unset()
}

// GetProjectEnvVar returns the ProjectEnvVar field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetProjectEnvVar() []ProjectEnvVarJsonhalProjectGet {
	if o == nil || isNil(o.ProjectEnvVar) {
		var ret []ProjectEnvVarJsonhalProjectGet
		return ret
	}
	return o.ProjectEnvVar
}

// GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetProjectEnvVarOk() ([]ProjectEnvVarJsonhalProjectGet, bool) {
	if o == nil || isNil(o.ProjectEnvVar) {
    return nil, false
	}
	return o.ProjectEnvVar, true
}

// HasProjectEnvVar returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasProjectEnvVar() bool {
	if o != nil && !isNil(o.ProjectEnvVar) {
		return true
	}

	return false
}

// SetProjectEnvVar gets a reference to the given []ProjectEnvVarJsonhalProjectGet and assigns it to the ProjectEnvVar field.
func (o *ProjectJsonhalProjectGet) SetProjectEnvVar(v []ProjectEnvVarJsonhalProjectGet) {
	o.ProjectEnvVar = v
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectJsonhalProjectGet) GetProjectTypeVersion() string {
	if o == nil || isNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectJsonhalProjectGet) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *ProjectJsonhalProjectGet) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}
// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *ProjectJsonhalProjectGet) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *ProjectJsonhalProjectGet) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetComponentVersion() []string {
	if o == nil || isNil(o.ComponentVersion) {
		var ret []string
		return ret
	}
	return o.ComponentVersion
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetComponentVersionOk() ([]string, bool) {
	if o == nil || isNil(o.ComponentVersion) {
    return nil, false
	}
	return o.ComponentVersion, true
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasComponentVersion() bool {
	if o != nil && !isNil(o.ComponentVersion) {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given []string and assigns it to the ComponentVersion field.
func (o *ProjectJsonhalProjectGet) SetComponentVersion(v []string) {
	o.ComponentVersion = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectJsonhalProjectGet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectJsonhalProjectGet) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectJsonhalProjectGet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectJsonhalProjectGet) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectJsonhalProjectGet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProjectJsonhalProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
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
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.Git.IsSet() {
		toSerialize["git"] = o.Git.Get()
	}
	if o.ServiceAccount.IsSet() {
		toSerialize["serviceAccount"] = o.ServiceAccount.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !isNil(o.ProjectEnvVar) {
		toSerialize["projectEnvVar"] = o.ProjectEnvVar
	}
	if o.ProjectTypeVersion.IsSet() {
		toSerialize["projectTypeVersion"] = o.ProjectTypeVersion.Get()
	}
	if !isNil(o.ComponentVersion) {
		toSerialize["componentVersion"] = o.ComponentVersion
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableProjectJsonhalProjectGet struct {
	value *ProjectJsonhalProjectGet
	isSet bool
}

func (v NullableProjectJsonhalProjectGet) Get() *ProjectJsonhalProjectGet {
	return v.value
}

func (v *NullableProjectJsonhalProjectGet) Set(val *ProjectJsonhalProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectJsonhalProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectJsonhalProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectJsonhalProjectGet(val *ProjectJsonhalProjectGet) *NullableProjectJsonhalProjectGet {
	return &NullableProjectJsonhalProjectGet{value: val, isSet: true}
}

func (v NullableProjectJsonhalProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectJsonhalProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


