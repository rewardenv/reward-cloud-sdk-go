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

// ResourceTypeLimit Class ResourceTypeLimit
type ResourceTypeLimit struct {
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	ProjectMinValue NullableInt32 `json:"projectMinValue,omitempty"`
	ProjectMaxValue NullableInt32 `json:"projectMaxValue,omitempty"`
	EnvironmentMinValue NullableInt32 `json:"environmentMinValue,omitempty"`
	EnvironmentMaxValue NullableInt32 `json:"environmentMaxValue,omitempty"`
	ComponentMinValue NullableInt32 `json:"componentMinValue,omitempty"`
	ComponentMaxValue NullableInt32 `json:"componentMaxValue,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewResourceTypeLimit instantiates a new ResourceTypeLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceTypeLimit() *ResourceTypeLimit {
	this := ResourceTypeLimit{}
	return &this
}

// NewResourceTypeLimitWithDefaults instantiates a new ResourceTypeLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceTypeLimitWithDefaults() *ResourceTypeLimit {
	this := ResourceTypeLimit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceTypeLimit) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeLimit) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResourceTypeLimit) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetUuidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ResourceTypeLimit) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ResourceTypeLimit) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ResourceTypeLimit) UnsetUuid() {
	o.Uuid.Unset()
}

// GetProjectMinValue returns the ProjectMinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetProjectMinValue() int32 {
	if o == nil || isNil(o.ProjectMinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectMinValue.Get()
}

// GetProjectMinValueOk returns a tuple with the ProjectMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetProjectMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectMinValue.Get(), o.ProjectMinValue.IsSet()
}

// HasProjectMinValue returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasProjectMinValue() bool {
	if o != nil && o.ProjectMinValue.IsSet() {
		return true
	}

	return false
}

// SetProjectMinValue gets a reference to the given NullableInt32 and assigns it to the ProjectMinValue field.
func (o *ResourceTypeLimit) SetProjectMinValue(v int32) {
	o.ProjectMinValue.Set(&v)
}
// SetProjectMinValueNil sets the value for ProjectMinValue to be an explicit nil
func (o *ResourceTypeLimit) SetProjectMinValueNil() {
	o.ProjectMinValue.Set(nil)
}

// UnsetProjectMinValue ensures that no value is present for ProjectMinValue, not even an explicit nil
func (o *ResourceTypeLimit) UnsetProjectMinValue() {
	o.ProjectMinValue.Unset()
}

// GetProjectMaxValue returns the ProjectMaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetProjectMaxValue() int32 {
	if o == nil || isNil(o.ProjectMaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectMaxValue.Get()
}

// GetProjectMaxValueOk returns a tuple with the ProjectMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetProjectMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectMaxValue.Get(), o.ProjectMaxValue.IsSet()
}

// HasProjectMaxValue returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasProjectMaxValue() bool {
	if o != nil && o.ProjectMaxValue.IsSet() {
		return true
	}

	return false
}

// SetProjectMaxValue gets a reference to the given NullableInt32 and assigns it to the ProjectMaxValue field.
func (o *ResourceTypeLimit) SetProjectMaxValue(v int32) {
	o.ProjectMaxValue.Set(&v)
}
// SetProjectMaxValueNil sets the value for ProjectMaxValue to be an explicit nil
func (o *ResourceTypeLimit) SetProjectMaxValueNil() {
	o.ProjectMaxValue.Set(nil)
}

// UnsetProjectMaxValue ensures that no value is present for ProjectMaxValue, not even an explicit nil
func (o *ResourceTypeLimit) UnsetProjectMaxValue() {
	o.ProjectMaxValue.Unset()
}

// GetEnvironmentMinValue returns the EnvironmentMinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetEnvironmentMinValue() int32 {
	if o == nil || isNil(o.EnvironmentMinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.EnvironmentMinValue.Get()
}

// GetEnvironmentMinValueOk returns a tuple with the EnvironmentMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetEnvironmentMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvironmentMinValue.Get(), o.EnvironmentMinValue.IsSet()
}

// HasEnvironmentMinValue returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasEnvironmentMinValue() bool {
	if o != nil && o.EnvironmentMinValue.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentMinValue gets a reference to the given NullableInt32 and assigns it to the EnvironmentMinValue field.
func (o *ResourceTypeLimit) SetEnvironmentMinValue(v int32) {
	o.EnvironmentMinValue.Set(&v)
}
// SetEnvironmentMinValueNil sets the value for EnvironmentMinValue to be an explicit nil
func (o *ResourceTypeLimit) SetEnvironmentMinValueNil() {
	o.EnvironmentMinValue.Set(nil)
}

// UnsetEnvironmentMinValue ensures that no value is present for EnvironmentMinValue, not even an explicit nil
func (o *ResourceTypeLimit) UnsetEnvironmentMinValue() {
	o.EnvironmentMinValue.Unset()
}

// GetEnvironmentMaxValue returns the EnvironmentMaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetEnvironmentMaxValue() int32 {
	if o == nil || isNil(o.EnvironmentMaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.EnvironmentMaxValue.Get()
}

// GetEnvironmentMaxValueOk returns a tuple with the EnvironmentMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetEnvironmentMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvironmentMaxValue.Get(), o.EnvironmentMaxValue.IsSet()
}

// HasEnvironmentMaxValue returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasEnvironmentMaxValue() bool {
	if o != nil && o.EnvironmentMaxValue.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentMaxValue gets a reference to the given NullableInt32 and assigns it to the EnvironmentMaxValue field.
func (o *ResourceTypeLimit) SetEnvironmentMaxValue(v int32) {
	o.EnvironmentMaxValue.Set(&v)
}
// SetEnvironmentMaxValueNil sets the value for EnvironmentMaxValue to be an explicit nil
func (o *ResourceTypeLimit) SetEnvironmentMaxValueNil() {
	o.EnvironmentMaxValue.Set(nil)
}

// UnsetEnvironmentMaxValue ensures that no value is present for EnvironmentMaxValue, not even an explicit nil
func (o *ResourceTypeLimit) UnsetEnvironmentMaxValue() {
	o.EnvironmentMaxValue.Unset()
}

// GetComponentMinValue returns the ComponentMinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetComponentMinValue() int32 {
	if o == nil || isNil(o.ComponentMinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ComponentMinValue.Get()
}

// GetComponentMinValueOk returns a tuple with the ComponentMinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetComponentMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentMinValue.Get(), o.ComponentMinValue.IsSet()
}

// HasComponentMinValue returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasComponentMinValue() bool {
	if o != nil && o.ComponentMinValue.IsSet() {
		return true
	}

	return false
}

// SetComponentMinValue gets a reference to the given NullableInt32 and assigns it to the ComponentMinValue field.
func (o *ResourceTypeLimit) SetComponentMinValue(v int32) {
	o.ComponentMinValue.Set(&v)
}
// SetComponentMinValueNil sets the value for ComponentMinValue to be an explicit nil
func (o *ResourceTypeLimit) SetComponentMinValueNil() {
	o.ComponentMinValue.Set(nil)
}

// UnsetComponentMinValue ensures that no value is present for ComponentMinValue, not even an explicit nil
func (o *ResourceTypeLimit) UnsetComponentMinValue() {
	o.ComponentMinValue.Unset()
}

// GetComponentMaxValue returns the ComponentMaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetComponentMaxValue() int32 {
	if o == nil || isNil(o.ComponentMaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ComponentMaxValue.Get()
}

// GetComponentMaxValueOk returns a tuple with the ComponentMaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetComponentMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentMaxValue.Get(), o.ComponentMaxValue.IsSet()
}

// HasComponentMaxValue returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasComponentMaxValue() bool {
	if o != nil && o.ComponentMaxValue.IsSet() {
		return true
	}

	return false
}

// SetComponentMaxValue gets a reference to the given NullableInt32 and assigns it to the ComponentMaxValue field.
func (o *ResourceTypeLimit) SetComponentMaxValue(v int32) {
	o.ComponentMaxValue.Set(&v)
}
// SetComponentMaxValueNil sets the value for ComponentMaxValue to be an explicit nil
func (o *ResourceTypeLimit) SetComponentMaxValueNil() {
	o.ComponentMaxValue.Set(nil)
}

// UnsetComponentMaxValue ensures that no value is present for ComponentMaxValue, not even an explicit nil
func (o *ResourceTypeLimit) UnsetComponentMaxValue() {
	o.ComponentMaxValue.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetCreatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ResourceTypeLimit) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ResourceTypeLimit) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ResourceTypeLimit) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceTypeLimit) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceTypeLimit) GetUpdatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *ResourceTypeLimit) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *ResourceTypeLimit) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *ResourceTypeLimit) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ResourceTypeLimit) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeLimit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ResourceTypeLimit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ResourceTypeLimit) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTypeLimit) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ResourceTypeLimit) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ResourceTypeLimit) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ResourceTypeLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.ProjectMinValue.IsSet() {
		toSerialize["projectMinValue"] = o.ProjectMinValue.Get()
	}
	if o.ProjectMaxValue.IsSet() {
		toSerialize["projectMaxValue"] = o.ProjectMaxValue.Get()
	}
	if o.EnvironmentMinValue.IsSet() {
		toSerialize["environmentMinValue"] = o.EnvironmentMinValue.Get()
	}
	if o.EnvironmentMaxValue.IsSet() {
		toSerialize["environmentMaxValue"] = o.EnvironmentMaxValue.Get()
	}
	if o.ComponentMinValue.IsSet() {
		toSerialize["componentMinValue"] = o.ComponentMinValue.Get()
	}
	if o.ComponentMaxValue.IsSet() {
		toSerialize["componentMaxValue"] = o.ComponentMaxValue.Get()
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

type NullableResourceTypeLimit struct {
	value *ResourceTypeLimit
	isSet bool
}

func (v NullableResourceTypeLimit) Get() *ResourceTypeLimit {
	return v.value
}

func (v *NullableResourceTypeLimit) Set(val *ResourceTypeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceTypeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceTypeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceTypeLimit(val *ResourceTypeLimit) *NullableResourceTypeLimit {
	return &NullableResourceTypeLimit{value: val, isSet: true}
}

func (v NullableResourceTypeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceTypeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

