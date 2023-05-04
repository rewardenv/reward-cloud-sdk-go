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

// ComponentVersionJsonhal Class ComponentVersion
type ComponentVersionJsonhal struct {
	Links                  *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id                     *int32                           `json:"id,omitempty"`
	Uuid                   NullableString                   `json:"uuid,omitempty"`
	Version                NullableString                   `json:"version,omitempty"`
	Component              NullableString                   `json:"component,omitempty"`
	ComponentVersionEnvVar []string                         `json:"componentVersionEnvVar,omitempty"`
	Project                []AbstractProjectJsonhal         `json:"project,omitempty"`
	EnvironmentComponent   []string                         `json:"environmentComponent,omitempty"`
	ComponentResourceLimit []string                         `json:"componentResourceLimit,omitempty"`
	CreatedBy              NullableString                   `json:"createdBy,omitempty"`
	UpdatedBy              NullableString                   `json:"updatedBy,omitempty"`
	CreatedAt              *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt              *time.Time                       `json:"updatedAt,omitempty"`
}

// NewComponentVersionJsonhal instantiates a new ComponentVersionJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentVersionJsonhal() *ComponentVersionJsonhal {
	this := ComponentVersionJsonhal{}
	return &this
}

// NewComponentVersionJsonhalWithDefaults instantiates a new ComponentVersionJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentVersionJsonhalWithDefaults() *ComponentVersionJsonhal {
	this := ComponentVersionJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *ComponentVersionJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ComponentVersionJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentVersionJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ComponentVersionJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ComponentVersionJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ComponentVersionJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionJsonhal) GetVersion() string {
	if o == nil || isNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentVersionJsonhal) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ComponentVersionJsonhal) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *ComponentVersionJsonhal) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ComponentVersionJsonhal) UnsetVersion() {
	o.Version.Unset()
}

// GetComponent returns the Component field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionJsonhal) GetComponent() string {
	if o == nil || isNil(o.Component.Get()) {
		var ret string
		return ret
	}
	return *o.Component.Get()
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentVersionJsonhal) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Component.Get(), o.Component.IsSet()
}

// HasComponent returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasComponent() bool {
	if o != nil && o.Component.IsSet() {
		return true
	}

	return false
}

// SetComponent gets a reference to the given NullableString and assigns it to the Component field.
func (o *ComponentVersionJsonhal) SetComponent(v string) {
	o.Component.Set(&v)
}

// SetComponentNil sets the value for Component to be an explicit nil
func (o *ComponentVersionJsonhal) SetComponentNil() {
	o.Component.Set(nil)
}

// UnsetComponent ensures that no value is present for Component, not even an explicit nil
func (o *ComponentVersionJsonhal) UnsetComponent() {
	o.Component.Unset()
}

// GetComponentVersionEnvVar returns the ComponentVersionEnvVar field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetComponentVersionEnvVar() []string {
	if o == nil || isNil(o.ComponentVersionEnvVar) {
		var ret []string
		return ret
	}
	return o.ComponentVersionEnvVar
}

// GetComponentVersionEnvVarOk returns a tuple with the ComponentVersionEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetComponentVersionEnvVarOk() ([]string, bool) {
	if o == nil || isNil(o.ComponentVersionEnvVar) {
		return nil, false
	}
	return o.ComponentVersionEnvVar, true
}

// HasComponentVersionEnvVar returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasComponentVersionEnvVar() bool {
	if o != nil && !isNil(o.ComponentVersionEnvVar) {
		return true
	}

	return false
}

// SetComponentVersionEnvVar gets a reference to the given []string and assigns it to the ComponentVersionEnvVar field.
func (o *ComponentVersionJsonhal) SetComponentVersionEnvVar(v []string) {
	o.ComponentVersionEnvVar = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetProject() []AbstractProjectJsonhal {
	if o == nil || isNil(o.Project) {
		var ret []AbstractProjectJsonhal
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetProjectOk() ([]AbstractProjectJsonhal, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given []AbstractProjectJsonhal and assigns it to the Project field.
func (o *ComponentVersionJsonhal) SetProject(v []AbstractProjectJsonhal) {
	o.Project = v
}

// GetEnvironmentComponent returns the EnvironmentComponent field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetEnvironmentComponent() []string {
	if o == nil || isNil(o.EnvironmentComponent) {
		var ret []string
		return ret
	}
	return o.EnvironmentComponent
}

// GetEnvironmentComponentOk returns a tuple with the EnvironmentComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetEnvironmentComponentOk() ([]string, bool) {
	if o == nil || isNil(o.EnvironmentComponent) {
		return nil, false
	}
	return o.EnvironmentComponent, true
}

// HasEnvironmentComponent returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasEnvironmentComponent() bool {
	if o != nil && !isNil(o.EnvironmentComponent) {
		return true
	}

	return false
}

// SetEnvironmentComponent gets a reference to the given []string and assigns it to the EnvironmentComponent field.
func (o *ComponentVersionJsonhal) SetEnvironmentComponent(v []string) {
	o.EnvironmentComponent = v
}

// GetComponentResourceLimit returns the ComponentResourceLimit field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetComponentResourceLimit() []string {
	if o == nil || isNil(o.ComponentResourceLimit) {
		var ret []string
		return ret
	}
	return o.ComponentResourceLimit
}

// GetComponentResourceLimitOk returns a tuple with the ComponentResourceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetComponentResourceLimitOk() ([]string, bool) {
	if o == nil || isNil(o.ComponentResourceLimit) {
		return nil, false
	}
	return o.ComponentResourceLimit, true
}

// HasComponentResourceLimit returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasComponentResourceLimit() bool {
	if o != nil && !isNil(o.ComponentResourceLimit) {
		return true
	}

	return false
}

// SetComponentResourceLimit gets a reference to the given []string and assigns it to the ComponentResourceLimit field.
func (o *ComponentVersionJsonhal) SetComponentResourceLimit(v []string) {
	o.ComponentResourceLimit = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentVersionJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ComponentVersionJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ComponentVersionJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ComponentVersionJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentVersionJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *ComponentVersionJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *ComponentVersionJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *ComponentVersionJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ComponentVersionJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ComponentVersionJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentVersionJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ComponentVersionJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ComponentVersionJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ComponentVersionJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Component.IsSet() {
		toSerialize["component"] = o.Component.Get()
	}
	if !isNil(o.ComponentVersionEnvVar) {
		toSerialize["componentVersionEnvVar"] = o.ComponentVersionEnvVar
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.EnvironmentComponent) {
		toSerialize["environmentComponent"] = o.EnvironmentComponent
	}
	if !isNil(o.ComponentResourceLimit) {
		toSerialize["componentResourceLimit"] = o.ComponentResourceLimit
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.UpdatedBy.IsSet() {
		toSerialize["updatedBy"] = o.UpdatedBy.Get()
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableComponentVersionJsonhal struct {
	value *ComponentVersionJsonhal
	isSet bool
}

func (v NullableComponentVersionJsonhal) Get() *ComponentVersionJsonhal {
	return v.value
}

func (v *NullableComponentVersionJsonhal) Set(val *ComponentVersionJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentVersionJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentVersionJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentVersionJsonhal(val *ComponentVersionJsonhal) *NullableComponentVersionJsonhal {
	return &NullableComponentVersionJsonhal{value: val, isSet: true}
}

func (v NullableComponentVersionJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentVersionJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
