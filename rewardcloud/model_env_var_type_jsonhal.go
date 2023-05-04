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

// EnvVarTypeJsonhal Class EnvVarType
type EnvVarTypeJsonhal struct {
	Links                  *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id                     *int32                           `json:"id,omitempty"`
	Uuid                   NullableString                   `json:"uuid,omitempty"`
	Name                   NullableString                   `json:"name,omitempty"`
	CodeName               NullableString                   `json:"codeName,omitempty"`
	IsDefault              NullableBool                     `json:"isDefault,omitempty"`
	ProjectEnvVar          []string                         `json:"projectEnvVar,omitempty"`
	EnvironmentEnvVar      []string                         `json:"environmentEnvVar,omitempty"`
	ComponentVersionEnvVar []string                         `json:"componentVersionEnvVar,omitempty"`
	OrganisationEnvVar     []string                         `json:"organisationEnvVar,omitempty"`
	TeamEnvVar             []string                         `json:"teamEnvVar,omitempty"`
	CreatedBy              NullableString                   `json:"createdBy,omitempty"`
	UpdatedBy              NullableString                   `json:"updatedBy,omitempty"`
	CreatedAt              *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt              *time.Time                       `json:"updatedAt,omitempty"`
}

// NewEnvVarTypeJsonhal instantiates a new EnvVarTypeJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvVarTypeJsonhal() *EnvVarTypeJsonhal {
	this := EnvVarTypeJsonhal{}
	return &this
}

// NewEnvVarTypeJsonhalWithDefaults instantiates a new EnvVarTypeJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvVarTypeJsonhalWithDefaults() *EnvVarTypeJsonhal {
	this := EnvVarTypeJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *EnvVarTypeJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EnvVarTypeJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVarTypeJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVarTypeJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *EnvVarTypeJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *EnvVarTypeJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *EnvVarTypeJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVarTypeJsonhal) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVarTypeJsonhal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EnvVarTypeJsonhal) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *EnvVarTypeJsonhal) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EnvVarTypeJsonhal) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVarTypeJsonhal) GetCodeName() string {
	if o == nil || isNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVarTypeJsonhal) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *EnvVarTypeJsonhal) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *EnvVarTypeJsonhal) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *EnvVarTypeJsonhal) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVarTypeJsonhal) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVarTypeJsonhal) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *EnvVarTypeJsonhal) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *EnvVarTypeJsonhal) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *EnvVarTypeJsonhal) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetProjectEnvVar returns the ProjectEnvVar field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetProjectEnvVar() []string {
	if o == nil || isNil(o.ProjectEnvVar) {
		var ret []string
		return ret
	}
	return o.ProjectEnvVar
}

// GetProjectEnvVarOk returns a tuple with the ProjectEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetProjectEnvVarOk() ([]string, bool) {
	if o == nil || isNil(o.ProjectEnvVar) {
		return nil, false
	}
	return o.ProjectEnvVar, true
}

// HasProjectEnvVar returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasProjectEnvVar() bool {
	if o != nil && !isNil(o.ProjectEnvVar) {
		return true
	}

	return false
}

// SetProjectEnvVar gets a reference to the given []string and assigns it to the ProjectEnvVar field.
func (o *EnvVarTypeJsonhal) SetProjectEnvVar(v []string) {
	o.ProjectEnvVar = v
}

// GetEnvironmentEnvVar returns the EnvironmentEnvVar field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetEnvironmentEnvVar() []string {
	if o == nil || isNil(o.EnvironmentEnvVar) {
		var ret []string
		return ret
	}
	return o.EnvironmentEnvVar
}

// GetEnvironmentEnvVarOk returns a tuple with the EnvironmentEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetEnvironmentEnvVarOk() ([]string, bool) {
	if o == nil || isNil(o.EnvironmentEnvVar) {
		return nil, false
	}
	return o.EnvironmentEnvVar, true
}

// HasEnvironmentEnvVar returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasEnvironmentEnvVar() bool {
	if o != nil && !isNil(o.EnvironmentEnvVar) {
		return true
	}

	return false
}

// SetEnvironmentEnvVar gets a reference to the given []string and assigns it to the EnvironmentEnvVar field.
func (o *EnvVarTypeJsonhal) SetEnvironmentEnvVar(v []string) {
	o.EnvironmentEnvVar = v
}

// GetComponentVersionEnvVar returns the ComponentVersionEnvVar field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetComponentVersionEnvVar() []string {
	if o == nil || isNil(o.ComponentVersionEnvVar) {
		var ret []string
		return ret
	}
	return o.ComponentVersionEnvVar
}

// GetComponentVersionEnvVarOk returns a tuple with the ComponentVersionEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetComponentVersionEnvVarOk() ([]string, bool) {
	if o == nil || isNil(o.ComponentVersionEnvVar) {
		return nil, false
	}
	return o.ComponentVersionEnvVar, true
}

// HasComponentVersionEnvVar returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasComponentVersionEnvVar() bool {
	if o != nil && !isNil(o.ComponentVersionEnvVar) {
		return true
	}

	return false
}

// SetComponentVersionEnvVar gets a reference to the given []string and assigns it to the ComponentVersionEnvVar field.
func (o *EnvVarTypeJsonhal) SetComponentVersionEnvVar(v []string) {
	o.ComponentVersionEnvVar = v
}

// GetOrganisationEnvVar returns the OrganisationEnvVar field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetOrganisationEnvVar() []string {
	if o == nil || isNil(o.OrganisationEnvVar) {
		var ret []string
		return ret
	}
	return o.OrganisationEnvVar
}

// GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetOrganisationEnvVarOk() ([]string, bool) {
	if o == nil || isNil(o.OrganisationEnvVar) {
		return nil, false
	}
	return o.OrganisationEnvVar, true
}

// HasOrganisationEnvVar returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasOrganisationEnvVar() bool {
	if o != nil && !isNil(o.OrganisationEnvVar) {
		return true
	}

	return false
}

// SetOrganisationEnvVar gets a reference to the given []string and assigns it to the OrganisationEnvVar field.
func (o *EnvVarTypeJsonhal) SetOrganisationEnvVar(v []string) {
	o.OrganisationEnvVar = v
}

// GetTeamEnvVar returns the TeamEnvVar field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetTeamEnvVar() []string {
	if o == nil || isNil(o.TeamEnvVar) {
		var ret []string
		return ret
	}
	return o.TeamEnvVar
}

// GetTeamEnvVarOk returns a tuple with the TeamEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetTeamEnvVarOk() ([]string, bool) {
	if o == nil || isNil(o.TeamEnvVar) {
		return nil, false
	}
	return o.TeamEnvVar, true
}

// HasTeamEnvVar returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasTeamEnvVar() bool {
	if o != nil && !isNil(o.TeamEnvVar) {
		return true
	}

	return false
}

// SetTeamEnvVar gets a reference to the given []string and assigns it to the TeamEnvVar field.
func (o *EnvVarTypeJsonhal) SetTeamEnvVar(v []string) {
	o.TeamEnvVar = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVarTypeJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVarTypeJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *EnvVarTypeJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *EnvVarTypeJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *EnvVarTypeJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvVarTypeJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvVarTypeJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *EnvVarTypeJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *EnvVarTypeJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *EnvVarTypeJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvVarTypeJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvVarTypeJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvVarTypeJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvVarTypeJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvVarTypeJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvVarTypeJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CodeName.IsSet() {
		toSerialize["codeName"] = o.CodeName.Get()
	}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if !isNil(o.ProjectEnvVar) {
		toSerialize["projectEnvVar"] = o.ProjectEnvVar
	}
	if !isNil(o.EnvironmentEnvVar) {
		toSerialize["environmentEnvVar"] = o.EnvironmentEnvVar
	}
	if !isNil(o.ComponentVersionEnvVar) {
		toSerialize["componentVersionEnvVar"] = o.ComponentVersionEnvVar
	}
	if !isNil(o.OrganisationEnvVar) {
		toSerialize["organisationEnvVar"] = o.OrganisationEnvVar
	}
	if !isNil(o.TeamEnvVar) {
		toSerialize["teamEnvVar"] = o.TeamEnvVar
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

type NullableEnvVarTypeJsonhal struct {
	value *EnvVarTypeJsonhal
	isSet bool
}

func (v NullableEnvVarTypeJsonhal) Get() *EnvVarTypeJsonhal {
	return v.value
}

func (v *NullableEnvVarTypeJsonhal) Set(val *EnvVarTypeJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvVarTypeJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvVarTypeJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvVarTypeJsonhal(val *EnvVarTypeJsonhal) *NullableEnvVarTypeJsonhal {
	return &NullableEnvVarTypeJsonhal{value: val, isSet: true}
}

func (v NullableEnvVarTypeJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvVarTypeJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
