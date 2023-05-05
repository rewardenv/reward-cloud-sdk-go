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

// checks if the ProjectTypeVersionJsonhal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectTypeVersionJsonhal{}

// ProjectTypeVersionJsonhal Class ProjectTypeVersion
type ProjectTypeVersionJsonhal struct {
	Links                    *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id                       *int32                           `json:"id,omitempty"`
	Uuid                     NullableString                   `json:"uuid,omitempty"`
	Version                  NullableString                   `json:"version,omitempty"`
	Project                  []AbstractProjectJsonhal         `json:"project,omitempty"`
	ProjectType              NullableString                   `json:"projectType,omitempty"`
	ProjectTypeVersionEnvVar []string                         `json:"projectTypeVersionEnvVar,omitempty"`
	ComponentResourceLimit   []string                         `json:"componentResourceLimit,omitempty"`
	CreatedBy                NullableString                   `json:"createdBy,omitempty"`
	UpdatedBy                NullableString                   `json:"updatedBy,omitempty"`
	CreatedAt                *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt                *time.Time                       `json:"updatedAt,omitempty"`
}

// NewProjectTypeVersionJsonhal instantiates a new ProjectTypeVersionJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTypeVersionJsonhal() *ProjectTypeVersionJsonhal {
	this := ProjectTypeVersionJsonhal{}
	return &this
}

// NewProjectTypeVersionJsonhalWithDefaults instantiates a new ProjectTypeVersionJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTypeVersionJsonhalWithDefaults() *ProjectTypeVersionJsonhal {
	this := ProjectTypeVersionJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *ProjectTypeVersionJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectTypeVersionJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionJsonhal) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ProjectTypeVersionJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ProjectTypeVersionJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ProjectTypeVersionJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionJsonhal) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionJsonhal) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ProjectTypeVersionJsonhal) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *ProjectTypeVersionJsonhal) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ProjectTypeVersionJsonhal) UnsetVersion() {
	o.Version.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetProject() []AbstractProjectJsonhal {
	if o == nil || IsNil(o.Project) {
		var ret []AbstractProjectJsonhal
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetProjectOk() ([]AbstractProjectJsonhal, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given []AbstractProjectJsonhal and assigns it to the Project field.
func (o *ProjectTypeVersionJsonhal) SetProject(v []AbstractProjectJsonhal) {
	o.Project = v
}

// GetProjectType returns the ProjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionJsonhal) GetProjectType() string {
	if o == nil || IsNil(o.ProjectType.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectType.Get()
}

// GetProjectTypeOk returns a tuple with the ProjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionJsonhal) GetProjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectType.Get(), o.ProjectType.IsSet()
}

// HasProjectType returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasProjectType() bool {
	if o != nil && o.ProjectType.IsSet() {
		return true
	}

	return false
}

// SetProjectType gets a reference to the given NullableString and assigns it to the ProjectType field.
func (o *ProjectTypeVersionJsonhal) SetProjectType(v string) {
	o.ProjectType.Set(&v)
}

// SetProjectTypeNil sets the value for ProjectType to be an explicit nil
func (o *ProjectTypeVersionJsonhal) SetProjectTypeNil() {
	o.ProjectType.Set(nil)
}

// UnsetProjectType ensures that no value is present for ProjectType, not even an explicit nil
func (o *ProjectTypeVersionJsonhal) UnsetProjectType() {
	o.ProjectType.Unset()
}

// GetProjectTypeVersionEnvVar returns the ProjectTypeVersionEnvVar field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetProjectTypeVersionEnvVar() []string {
	if o == nil || IsNil(o.ProjectTypeVersionEnvVar) {
		var ret []string
		return ret
	}
	return o.ProjectTypeVersionEnvVar
}

// GetProjectTypeVersionEnvVarOk returns a tuple with the ProjectTypeVersionEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetProjectTypeVersionEnvVarOk() ([]string, bool) {
	if o == nil || IsNil(o.ProjectTypeVersionEnvVar) {
		return nil, false
	}
	return o.ProjectTypeVersionEnvVar, true
}

// HasProjectTypeVersionEnvVar returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasProjectTypeVersionEnvVar() bool {
	if o != nil && !IsNil(o.ProjectTypeVersionEnvVar) {
		return true
	}

	return false
}

// SetProjectTypeVersionEnvVar gets a reference to the given []string and assigns it to the ProjectTypeVersionEnvVar field.
func (o *ProjectTypeVersionJsonhal) SetProjectTypeVersionEnvVar(v []string) {
	o.ProjectTypeVersionEnvVar = v
}

// GetComponentResourceLimit returns the ComponentResourceLimit field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetComponentResourceLimit() []string {
	if o == nil || IsNil(o.ComponentResourceLimit) {
		var ret []string
		return ret
	}
	return o.ComponentResourceLimit
}

// GetComponentResourceLimitOk returns a tuple with the ComponentResourceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetComponentResourceLimitOk() ([]string, bool) {
	if o == nil || IsNil(o.ComponentResourceLimit) {
		return nil, false
	}
	return o.ComponentResourceLimit, true
}

// HasComponentResourceLimit returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasComponentResourceLimit() bool {
	if o != nil && !IsNil(o.ComponentResourceLimit) {
		return true
	}

	return false
}

// SetComponentResourceLimit gets a reference to the given []string and assigns it to the ComponentResourceLimit field.
func (o *ProjectTypeVersionJsonhal) SetComponentResourceLimit(v []string) {
	o.ComponentResourceLimit = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionJsonhal) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ProjectTypeVersionJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ProjectTypeVersionJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ProjectTypeVersionJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionJsonhal) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *ProjectTypeVersionJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *ProjectTypeVersionJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *ProjectTypeVersionJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectTypeVersionJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectTypeVersionJsonhal) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectTypeVersionJsonhal) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectTypeVersionJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProjectTypeVersionJsonhal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectTypeVersionJsonhal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	// skip: id is readOnly
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if o.ProjectType.IsSet() {
		toSerialize["projectType"] = o.ProjectType.Get()
	}
	if !IsNil(o.ProjectTypeVersionEnvVar) {
		toSerialize["projectTypeVersionEnvVar"] = o.ProjectTypeVersionEnvVar
	}
	if !IsNil(o.ComponentResourceLimit) {
		toSerialize["componentResourceLimit"] = o.ComponentResourceLimit
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updatedBy"] = o.UpdatedBy.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableProjectTypeVersionJsonhal struct {
	value *ProjectTypeVersionJsonhal
	isSet bool
}

func (v NullableProjectTypeVersionJsonhal) Get() *ProjectTypeVersionJsonhal {
	return v.value
}

func (v *NullableProjectTypeVersionJsonhal) Set(val *ProjectTypeVersionJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTypeVersionJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTypeVersionJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTypeVersionJsonhal(val *ProjectTypeVersionJsonhal) *NullableProjectTypeVersionJsonhal {
	return &NullableProjectTypeVersionJsonhal{value: val, isSet: true}
}

func (v NullableProjectTypeVersionJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTypeVersionJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
