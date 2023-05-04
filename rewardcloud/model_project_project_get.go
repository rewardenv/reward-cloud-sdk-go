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

// checks if the ProjectProjectGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectProjectGet{}

// ProjectProjectGet Class Project
type ProjectProjectGet struct {
	TemplateIri NullableString `json:"templateIri,omitempty"`
	State NullableString `json:"state,omitempty"`
	Git NullableProjectProjectGetGit `json:"git,omitempty"`
	ServiceAccount NullableProjectProjectGetServiceAccount `json:"serviceAccount,omitempty"`
	Team NullableString `json:"team,omitempty"`
	Environment []string `json:"environment,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
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
	ProjectEnvVar []ProjectEnvVarProjectGet `json:"projectEnvVar,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewProjectProjectGet instantiates a new ProjectProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectProjectGet() *ProjectProjectGet {
	this := ProjectProjectGet{}
	return &this
}

// NewProjectProjectGetWithDefaults instantiates a new ProjectProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectProjectGetWithDefaults() *ProjectProjectGet {
	this := ProjectProjectGet{}
	return &this
}

// GetTemplateIri returns the TemplateIri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetTemplateIri() string {
	if o == nil || IsNil(o.TemplateIri.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateIri.Get()
}

// GetTemplateIriOk returns a tuple with the TemplateIri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetTemplateIriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateIri.Get(), o.TemplateIri.IsSet()
}

// HasTemplateIri returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasTemplateIri() bool {
	if o != nil && o.TemplateIri.IsSet() {
		return true
	}

	return false
}

// SetTemplateIri gets a reference to the given NullableString and assigns it to the TemplateIri field.
func (o *ProjectProjectGet) SetTemplateIri(v string) {
	o.TemplateIri.Set(&v)
}
// SetTemplateIriNil sets the value for TemplateIri to be an explicit nil
func (o *ProjectProjectGet) SetTemplateIriNil() {
	o.TemplateIri.Set(nil)
}

// UnsetTemplateIri ensures that no value is present for TemplateIri, not even an explicit nil
func (o *ProjectProjectGet) UnsetTemplateIri() {
	o.TemplateIri.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *ProjectProjectGet) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *ProjectProjectGet) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ProjectProjectGet) UnsetState() {
	o.State.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetGit() ProjectProjectGetGit {
	if o == nil || IsNil(o.Git.Get()) {
		var ret ProjectProjectGetGit
		return ret
	}
	return *o.Git.Get()
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetGitOk() (*ProjectProjectGetGit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Git.Get(), o.Git.IsSet()
}

// HasGit returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasGit() bool {
	if o != nil && o.Git.IsSet() {
		return true
	}

	return false
}

// SetGit gets a reference to the given NullableProjectProjectGetGit and assigns it to the Git field.
func (o *ProjectProjectGet) SetGit(v ProjectProjectGetGit) {
	o.Git.Set(&v)
}
// SetGitNil sets the value for Git to be an explicit nil
func (o *ProjectProjectGet) SetGitNil() {
	o.Git.Set(nil)
}

// UnsetGit ensures that no value is present for Git, not even an explicit nil
func (o *ProjectProjectGet) UnsetGit() {
	o.Git.Unset()
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetServiceAccount() ProjectProjectGetServiceAccount {
	if o == nil || IsNil(o.ServiceAccount.Get()) {
		var ret ProjectProjectGetServiceAccount
		return ret
	}
	return *o.ServiceAccount.Get()
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetServiceAccountOk() (*ProjectProjectGetServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceAccount.Get(), o.ServiceAccount.IsSet()
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasServiceAccount() bool {
	if o != nil && o.ServiceAccount.IsSet() {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given NullableProjectProjectGetServiceAccount and assigns it to the ServiceAccount field.
func (o *ProjectProjectGet) SetServiceAccount(v ProjectProjectGetServiceAccount) {
	o.ServiceAccount.Set(&v)
}
// SetServiceAccountNil sets the value for ServiceAccount to be an explicit nil
func (o *ProjectProjectGet) SetServiceAccountNil() {
	o.ServiceAccount.Set(nil)
}

// UnsetServiceAccount ensures that no value is present for ServiceAccount, not even an explicit nil
func (o *ProjectProjectGet) UnsetServiceAccount() {
	o.ServiceAccount.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetTeam() string {
	if o == nil || IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *ProjectProjectGet) SetTeam(v string) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *ProjectProjectGet) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *ProjectProjectGet) UnsetTeam() {
	o.Team.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetEnvironment() []string {
	if o == nil || IsNil(o.Environment) {
		var ret []string
		return ret
	}
	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetEnvironmentOk() ([]string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given []string and assigns it to the Environment field.
func (o *ProjectProjectGet) SetEnvironment(v []string) {
	o.Environment = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectProjectGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ProjectProjectGet) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProjectProjectGet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProjectProjectGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProjectProjectGet) UnsetName() {
	o.Name.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *ProjectProjectGet) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *ProjectProjectGet) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *ProjectProjectGet) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *ProjectProjectGet) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *ProjectProjectGet) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *ProjectProjectGet) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *ProjectProjectGet) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *ProjectProjectGet) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *ProjectProjectGet) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetStorage() int32 {
	if o == nil || IsNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetStorageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *ProjectProjectGet) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *ProjectProjectGet) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *ProjectProjectGet) UnsetStorage() {
	o.Storage.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *ProjectProjectGet) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ProjectProjectGet) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ProjectProjectGet) UnsetCode() {
	o.Code.Unset()
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *ProjectProjectGet) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *ProjectProjectGet) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *ProjectProjectGet) UnsetColor() {
	o.Color.Unset()
}

// GetIsInitProjectSkeleton returns the IsInitProjectSkeleton field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetIsInitProjectSkeleton() bool {
	if o == nil || IsNil(o.IsInitProjectSkeleton.Get()) {
		var ret bool
		return ret
	}
	return *o.IsInitProjectSkeleton.Get()
}

// GetIsInitProjectSkeletonOk returns a tuple with the IsInitProjectSkeleton field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetIsInitProjectSkeletonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsInitProjectSkeleton.Get(), o.IsInitProjectSkeleton.IsSet()
}

// HasIsInitProjectSkeleton returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasIsInitProjectSkeleton() bool {
	if o != nil && o.IsInitProjectSkeleton.IsSet() {
		return true
	}

	return false
}

// SetIsInitProjectSkeleton gets a reference to the given NullableBool and assigns it to the IsInitProjectSkeleton field.
func (o *ProjectProjectGet) SetIsInitProjectSkeleton(v bool) {
	o.IsInitProjectSkeleton.Set(&v)
}
// SetIsInitProjectSkeletonNil sets the value for IsInitProjectSkeleton to be an explicit nil
func (o *ProjectProjectGet) SetIsInitProjectSkeletonNil() {
	o.IsInitProjectSkeleton.Set(nil)
}

// UnsetIsInitProjectSkeleton ensures that no value is present for IsInitProjectSkeleton, not even an explicit nil
func (o *ProjectProjectGet) UnsetIsInitProjectSkeleton() {
	o.IsInitProjectSkeleton.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetComponentVersion() []string {
	if o == nil || IsNil(o.ComponentVersion) {
		var ret []string
		return ret
	}
	return o.ComponentVersion
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetComponentVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.ComponentVersion) {
		return nil, false
	}
	return o.ComponentVersion, true
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasComponentVersion() bool {
	if o != nil && !IsNil(o.ComponentVersion) {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given []string and assigns it to the ComponentVersion field.
func (o *ProjectProjectGet) SetComponentVersion(v []string) {
	o.ComponentVersion = v
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectProjectGet) GetProjectTypeVersion() string {
	if o == nil || IsNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectProjectGet) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *ProjectProjectGet) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}
// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *ProjectProjectGet) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *ProjectProjectGet) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetProjectEnvVar returns the ProjectEnvVar field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetProjectEnvVar() []ProjectEnvVarProjectGet {
	if o == nil || IsNil(o.ProjectEnvVar) {
		var ret []ProjectEnvVarProjectGet
		return ret
	}
	return o.ProjectEnvVar
}

// GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetProjectEnvVarOk() ([]ProjectEnvVarProjectGet, bool) {
	if o == nil || IsNil(o.ProjectEnvVar) {
		return nil, false
	}
	return o.ProjectEnvVar, true
}

// HasProjectEnvVar returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasProjectEnvVar() bool {
	if o != nil && !IsNil(o.ProjectEnvVar) {
		return true
	}

	return false
}

// SetProjectEnvVar gets a reference to the given []ProjectEnvVarProjectGet and assigns it to the ProjectEnvVar field.
func (o *ProjectProjectGet) SetProjectEnvVar(v []ProjectEnvVarProjectGet) {
	o.ProjectEnvVar = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectProjectGet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectProjectGet) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectProjectGet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectProjectGet) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectProjectGet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProjectProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectProjectGet) ToMap() (map[string]interface{}, error) {
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

type NullableProjectProjectGet struct {
	value *ProjectProjectGet
	isSet bool
}

func (v NullableProjectProjectGet) Get() *ProjectProjectGet {
	return v.value
}

func (v *NullableProjectProjectGet) Set(val *ProjectProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectProjectGet(val *ProjectProjectGet) *NullableProjectProjectGet {
	return &NullableProjectProjectGet{value: val, isSet: true}
}

func (v NullableProjectProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


