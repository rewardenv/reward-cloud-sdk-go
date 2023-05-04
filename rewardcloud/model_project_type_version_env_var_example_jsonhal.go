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

// ProjectTypeVersionEnvVarExampleJsonhal Class ProjectTypeEnvVarExample
type ProjectTypeVersionEnvVarExampleJsonhal struct {
	Links                    *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id                       *int32                           `json:"id,omitempty"`
	Uuid                     NullableString                   `json:"uuid,omitempty"`
	Value                    NullableString                   `json:"value,omitempty"`
	IsDefault                NullableBool                     `json:"isDefault,omitempty"`
	ProjectTypeVersionEnvVar NullableString                   `json:"projectTypeVersionEnvVar,omitempty"`
	CreatedBy                NullableString                   `json:"createdBy,omitempty"`
	UpdatedBy                NullableString                   `json:"updatedBy,omitempty"`
	CreatedAt                *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt                *time.Time                       `json:"updatedAt,omitempty"`
}

// NewProjectTypeVersionEnvVarExampleJsonhal instantiates a new ProjectTypeVersionEnvVarExampleJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectTypeVersionEnvVarExampleJsonhal() *ProjectTypeVersionEnvVarExampleJsonhal {
	this := ProjectTypeVersionEnvVarExampleJsonhal{}
	return &this
}

// NewProjectTypeVersionEnvVarExampleJsonhalWithDefaults instantiates a new ProjectTypeVersionEnvVarExampleJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectTypeVersionEnvVarExampleJsonhalWithDefaults() *ProjectTypeVersionEnvVarExampleJsonhal {
	this := ProjectTypeVersionEnvVarExampleJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetValue() {
	o.Value.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetProjectTypeVersionEnvVar returns the ProjectTypeVersionEnvVar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetProjectTypeVersionEnvVar() string {
	if o == nil || isNil(o.ProjectTypeVersionEnvVar.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersionEnvVar.Get()
}

// GetProjectTypeVersionEnvVarOk returns a tuple with the ProjectTypeVersionEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetProjectTypeVersionEnvVarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectTypeVersionEnvVar.Get(), o.ProjectTypeVersionEnvVar.IsSet()
}

// HasProjectTypeVersionEnvVar returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasProjectTypeVersionEnvVar() bool {
	if o != nil && o.ProjectTypeVersionEnvVar.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersionEnvVar gets a reference to the given NullableString and assigns it to the ProjectTypeVersionEnvVar field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetProjectTypeVersionEnvVar(v string) {
	o.ProjectTypeVersionEnvVar.Set(&v)
}

// SetProjectTypeVersionEnvVarNil sets the value for ProjectTypeVersionEnvVar to be an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetProjectTypeVersionEnvVarNil() {
	o.ProjectTypeVersionEnvVar.Set(nil)
}

// UnsetProjectTypeVersionEnvVar ensures that no value is present for ProjectTypeVersionEnvVar, not even an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetProjectTypeVersionEnvVar() {
	o.ProjectTypeVersionEnvVar.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *ProjectTypeVersionEnvVarExampleJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProjectTypeVersionEnvVarExampleJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProjectTypeVersionEnvVarExampleJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if o.ProjectTypeVersionEnvVar.IsSet() {
		toSerialize["projectTypeVersionEnvVar"] = o.ProjectTypeVersionEnvVar.Get()
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

type NullableProjectTypeVersionEnvVarExampleJsonhal struct {
	value *ProjectTypeVersionEnvVarExampleJsonhal
	isSet bool
}

func (v NullableProjectTypeVersionEnvVarExampleJsonhal) Get() *ProjectTypeVersionEnvVarExampleJsonhal {
	return v.value
}

func (v *NullableProjectTypeVersionEnvVarExampleJsonhal) Set(val *ProjectTypeVersionEnvVarExampleJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectTypeVersionEnvVarExampleJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectTypeVersionEnvVarExampleJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectTypeVersionEnvVarExampleJsonhal(val *ProjectTypeVersionEnvVarExampleJsonhal) *NullableProjectTypeVersionEnvVarExampleJsonhal {
	return &NullableProjectTypeVersionEnvVarExampleJsonhal{value: val, isSet: true}
}

func (v NullableProjectTypeVersionEnvVarExampleJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectTypeVersionEnvVarExampleJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
