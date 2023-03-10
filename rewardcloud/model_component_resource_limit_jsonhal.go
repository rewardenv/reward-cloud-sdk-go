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

// ComponentResourceLimitJsonhal Class ComponentResourceLimit
type ComponentResourceLimitJsonhal struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	MinValue NullableInt32 `json:"minValue,omitempty"`
	MaxValue NullableInt32 `json:"maxValue,omitempty"`
	ProjectTypeVersion NullableString `json:"projectTypeVersion,omitempty"`
	ResourceType NullableString `json:"resourceType,omitempty"`
	ComponentVersion NullableString `json:"componentVersion,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewComponentResourceLimitJsonhal instantiates a new ComponentResourceLimitJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentResourceLimitJsonhal() *ComponentResourceLimitJsonhal {
	this := ComponentResourceLimitJsonhal{}
	return &this
}

// NewComponentResourceLimitJsonhalWithDefaults instantiates a new ComponentResourceLimitJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentResourceLimitJsonhalWithDefaults() *ComponentResourceLimitJsonhal {
	this := ComponentResourceLimitJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ComponentResourceLimitJsonhal) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentResourceLimitJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *ComponentResourceLimitJsonhal) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentResourceLimitJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentResourceLimitJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ComponentResourceLimitJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *ComponentResourceLimitJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetMinValue returns the MinValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetMinValue() int32 {
	if o == nil || isNil(o.MinValue.Get()) {
		var ret int32
		return ret
	}
	return *o.MinValue.Get()
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetMinValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinValue.Get(), o.MinValue.IsSet()
}

// HasMinValue returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasMinValue() bool {
	if o != nil && o.MinValue.IsSet() {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given NullableInt32 and assigns it to the MinValue field.
func (o *ComponentResourceLimitJsonhal) SetMinValue(v int32) {
	o.MinValue.Set(&v)
}
// SetMinValueNil sets the value for MinValue to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetMinValueNil() {
	o.MinValue.Set(nil)
}

// UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetMinValue() {
	o.MinValue.Unset()
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetMaxValue() int32 {
	if o == nil || isNil(o.MaxValue.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxValue.Get()
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetMaxValueOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxValue.Get(), o.MaxValue.IsSet()
}

// HasMaxValue returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasMaxValue() bool {
	if o != nil && o.MaxValue.IsSet() {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given NullableInt32 and assigns it to the MaxValue field.
func (o *ComponentResourceLimitJsonhal) SetMaxValue(v int32) {
	o.MaxValue.Set(&v)
}
// SetMaxValueNil sets the value for MaxValue to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetMaxValueNil() {
	o.MaxValue.Set(nil)
}

// UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetMaxValue() {
	o.MaxValue.Unset()
}

// GetProjectTypeVersion returns the ProjectTypeVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetProjectTypeVersion() string {
	if o == nil || isNil(o.ProjectTypeVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectTypeVersion.Get()
}

// GetProjectTypeVersionOk returns a tuple with the ProjectTypeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetProjectTypeVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ProjectTypeVersion.Get(), o.ProjectTypeVersion.IsSet()
}

// HasProjectTypeVersion returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasProjectTypeVersion() bool {
	if o != nil && o.ProjectTypeVersion.IsSet() {
		return true
	}

	return false
}

// SetProjectTypeVersion gets a reference to the given NullableString and assigns it to the ProjectTypeVersion field.
func (o *ComponentResourceLimitJsonhal) SetProjectTypeVersion(v string) {
	o.ProjectTypeVersion.Set(&v)
}
// SetProjectTypeVersionNil sets the value for ProjectTypeVersion to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetProjectTypeVersionNil() {
	o.ProjectTypeVersion.Set(nil)
}

// UnsetProjectTypeVersion ensures that no value is present for ProjectTypeVersion, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetProjectTypeVersion() {
	o.ProjectTypeVersion.Unset()
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetResourceType() string {
	if o == nil || isNil(o.ResourceType.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetResourceTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// HasResourceType returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasResourceType() bool {
	if o != nil && o.ResourceType.IsSet() {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given NullableString and assigns it to the ResourceType field.
func (o *ComponentResourceLimitJsonhal) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}
// SetResourceTypeNil sets the value for ResourceType to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetResourceTypeNil() {
	o.ResourceType.Set(nil)
}

// UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetResourceType() {
	o.ResourceType.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetComponentVersion() string {
	if o == nil || isNil(o.ComponentVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentVersion.Get()
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetComponentVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentVersion.Get(), o.ComponentVersion.IsSet()
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasComponentVersion() bool {
	if o != nil && o.ComponentVersion.IsSet() {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given NullableString and assigns it to the ComponentVersion field.
func (o *ComponentResourceLimitJsonhal) SetComponentVersion(v string) {
	o.ComponentVersion.Set(&v)
}
// SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetComponentVersionNil() {
	o.ComponentVersion.Set(nil)
}

// UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetComponentVersion() {
	o.ComponentVersion.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *ComponentResourceLimitJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentResourceLimitJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComponentResourceLimitJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *ComponentResourceLimitJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *ComponentResourceLimitJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *ComponentResourceLimitJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ComponentResourceLimitJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentResourceLimitJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ComponentResourceLimitJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ComponentResourceLimitJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentResourceLimitJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ComponentResourceLimitJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ComponentResourceLimitJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ComponentResourceLimitJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.MinValue.IsSet() {
		toSerialize["minValue"] = o.MinValue.Get()
	}
	if o.MaxValue.IsSet() {
		toSerialize["maxValue"] = o.MaxValue.Get()
	}
	if o.ProjectTypeVersion.IsSet() {
		toSerialize["projectTypeVersion"] = o.ProjectTypeVersion.Get()
	}
	if o.ResourceType.IsSet() {
		toSerialize["resourceType"] = o.ResourceType.Get()
	}
	if o.ComponentVersion.IsSet() {
		toSerialize["componentVersion"] = o.ComponentVersion.Get()
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

type NullableComponentResourceLimitJsonhal struct {
	value *ComponentResourceLimitJsonhal
	isSet bool
}

func (v NullableComponentResourceLimitJsonhal) Get() *ComponentResourceLimitJsonhal {
	return v.value
}

func (v *NullableComponentResourceLimitJsonhal) Set(val *ComponentResourceLimitJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentResourceLimitJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentResourceLimitJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentResourceLimitJsonhal(val *ComponentResourceLimitJsonhal) *NullableComponentResourceLimitJsonhal {
	return &NullableComponentResourceLimitJsonhal{value: val, isSet: true}
}

func (v NullableComponentResourceLimitJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentResourceLimitJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


