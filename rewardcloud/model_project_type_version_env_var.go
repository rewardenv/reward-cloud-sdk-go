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

// ProjectTypeVersionEnvVar Class ProjectTypeEnvVar
type ProjectTypeVersionEnvVar struct {
	Id                              *int32         `json:"id,omitempty"`
	Uuid                            NullableString `json:"uuid,omitempty"`
	Key                             NullableString `json:"key,omitempty"`
	Note                            NullableString `json:"note,omitempty"`
	IsRequired                      NullableBool   `json:"isRequired,omitempty"`
	IsEncrypted                     NullableBool   `json:"isEncrypted,omitempty"`
	ProjectTypeVersion              NullableString `json:"projectTypeVersion,omitempty"`
	ProjectTypeVersionEnvVarExample []string       `json:"projectTypeVersionEnvVarExample,omitempty"`
	CreatedBy                       NullableString `json:"createdBy,omitempty"`
	UpdatedBy                       NullableString `json:"updatedBy,omitempty"`
	CreatedAt                       *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt                       *time.Time     `json:"updatedAt,omitempty"`
}

// NewProjectTypeVersionEnvVar instantiates a new ProjectTypeVersionEnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTypeVersionEnvVar() *ProjectTypeVersionEnvVar {
	this := ProjectTypeVersionEnvVar{}
	return &this
}

// NewProjectTypeVersionEnvVarWithDefaults instantiates a new ProjectTypeVersionEnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTypeVersionEnvVarWithDefaults() *ProjectTypeVersionEnvVar {
	this := ProjectTypeVersionEnvVar{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVar) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVar) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectTypeVersionEnvVar) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ProjectTypeVersionEnvVar) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetUuid() {
	o.Uuid.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetKey() string {
	if o == nil || isNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *ProjectTypeVersionEnvVar) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetKey() {
	o.Key.Unset()
}

// GetNote returns the Note field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetNote() string {
	if o == nil || isNil(o.Note.Get()) {
		var ret string
		return ret
	}
	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// HasNote returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasNote() bool {
	if o != nil && o.Note.IsSet() {
		return true
	}

	return false
}

// SetNote gets a reference to the given NullableString and assigns it to the Note field.
func (o *ProjectTypeVersionEnvVar) SetNote(v string) {
	o.Note.Set(&v)
}

// SetNoteNil sets the value for Note to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetNoteNil() {
	o.Note.Set(nil)
}

// UnsetNote ensures that no value is present for Note, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetNote() {
	o.Note.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequired.Get()
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequired.Get(), o.IsRequired.IsSet()
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasIsRequired() bool {
	if o != nil && o.IsRequired.IsSet() {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given NullableBool and assigns it to the IsRequired field.
func (o *ProjectTypeVersionEnvVar) SetIsRequired(v bool) {
	o.IsRequired.Set(&v)
}

// SetIsRequiredNil sets the value for IsRequired to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetIsRequiredNil() {
	o.IsRequired.Set(nil)
}

// UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetIsRequired() {
	o.IsRequired.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetIsEncrypted() bool {
	if o == nil || isNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *ProjectTypeVersionEnvVar) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}

// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetProjectTypeVersion() string {
	if o == nil || isNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *ProjectTypeVersionEnvVar) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}

// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetProjectTypeVersionEnvVarExample returns the ProjectTypeVersionEnvVarExample field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVar) GetProjectTypeVersionEnvVarExample() []string {
	if o == nil || isNil(o.ProjectTypeVersionEnvVarExample) {
		var ret []string
		return ret
	}
	return o.ProjectTypeVersionEnvVarExample
}

// GetProjectTypeVersionEnvVarExampleOk returns a tuple with the ProjectTypeVersionEnvVarExample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVar) GetProjectTypeVersionEnvVarExampleOk() ([]string, bool) {
	if o == nil || isNil(o.ProjectTypeVersionEnvVarExample) {
		return nil, false
	}
	return o.ProjectTypeVersionEnvVarExample, true
}

// HasProjectTypeVersionEnvVarExample returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasProjectTypeVersionEnvVarExample() bool {
	if o != nil && !isNil(o.ProjectTypeVersionEnvVarExample) {
		return true
	}

	return false
}

// SetProjectTypeVersionEnvVarExample gets a reference to the given []string and assigns it to the ProjectTypeVersionEnvVarExample field.
func (o *ProjectTypeVersionEnvVar) SetProjectTypeVersionEnvVarExample(v []string) {
	o.ProjectTypeVersionEnvVarExample = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ProjectTypeVersionEnvVar) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVar) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVar) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *ProjectTypeVersionEnvVar) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *ProjectTypeVersionEnvVar) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *ProjectTypeVersionEnvVar) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVar) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVar) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectTypeVersionEnvVar) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVar) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVar) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVar) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectTypeVersionEnvVar) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProjectTypeVersionEnvVar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Note.IsSet() {
		toSerialize["note"] = o.Note.Get()
	}
	if o.IsRequired.IsSet() {
		toSerialize["isRequired"] = o.IsRequired.Get()
	}
	if o.IsEncrypted.IsSet() {
		toSerialize["isEncrypted"] = o.IsEncrypted.Get()
	}
	if o.ProjectTypeVersion.IsSet() {
		toSerialize["projectTypeVersion"] = o.ProjectTypeVersion.Get()
	}
	if !isNil(o.ProjectTypeVersionEnvVarExample) {
		toSerialize["projectTypeVersionEnvVarExample"] = o.ProjectTypeVersionEnvVarExample
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

type NullableProjectTypeVersionEnvVar struct {
	value *ProjectTypeVersionEnvVar
	isSet bool
}

func (v NullableProjectTypeVersionEnvVar) Get() *ProjectTypeVersionEnvVar {
	return v.value
}

func (v *NullableProjectTypeVersionEnvVar) Set(val *ProjectTypeVersionEnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTypeVersionEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTypeVersionEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTypeVersionEnvVar(val *ProjectTypeVersionEnvVar) *NullableProjectTypeVersionEnvVar {
	return &NullableProjectTypeVersionEnvVar{value: val, isSet: true}
}

func (v NullableProjectTypeVersionEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTypeVersionEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
