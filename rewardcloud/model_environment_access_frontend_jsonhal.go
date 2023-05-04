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

// EnvironmentAccessFrontendJsonhal Class EnvironmentAccessFrontend
type EnvironmentAccessFrontendJsonhal struct {
	Links             *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id                *int32                           `json:"id,omitempty"`
	Uuid              NullableString                   `json:"uuid,omitempty"`
	Domain            NullableString                   `json:"domain,omitempty"`
	Path              NullableString                   `json:"path,omitempty"`
	Url               NullableString                   `json:"url,omitempty"`
	EnvironmentAccess NullableString                   `json:"environmentAccess,omitempty"`
	CreatedBy         NullableString                   `json:"createdBy,omitempty"`
	UpdatedBy         NullableString                   `json:"updatedBy,omitempty"`
	CreatedAt         *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt         *time.Time                       `json:"updatedAt,omitempty"`
}

// NewEnvironmentAccessFrontendJsonhal instantiates a new EnvironmentAccessFrontendJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentAccessFrontendJsonhal() *EnvironmentAccessFrontendJsonhal {
	this := EnvironmentAccessFrontendJsonhal{}
	return &this
}

// NewEnvironmentAccessFrontendJsonhalWithDefaults instantiates a new EnvironmentAccessFrontendJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentAccessFrontendJsonhalWithDefaults() *EnvironmentAccessFrontendJsonhal {
	this := EnvironmentAccessFrontendJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvironmentAccessFrontendJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessFrontendJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *EnvironmentAccessFrontendJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentAccessFrontendJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessFrontendJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EnvironmentAccessFrontendJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *EnvironmentAccessFrontendJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetDomain() string {
	if o == nil || isNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *EnvironmentAccessFrontendJsonhal) SetDomain(v string) {
	o.Domain.Set(&v)
}

// SetDomainNil sets the value for Domain to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetDomain() {
	o.Domain.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetPath() string {
	if o == nil || isNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *EnvironmentAccessFrontendJsonhal) SetPath(v string) {
	o.Path.Set(&v)
}

// SetPathNil sets the value for Path to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetPath() {
	o.Path.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetUrl() string {
	if o == nil || isNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *EnvironmentAccessFrontendJsonhal) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetUrl() {
	o.Url.Unset()
}

// GetEnvironmentAccess returns the EnvironmentAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetEnvironmentAccess() string {
	if o == nil || isNil(o.EnvironmentAccess.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentAccess.Get()
}

// GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetEnvironmentAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentAccess.Get(), o.EnvironmentAccess.IsSet()
}

// HasEnvironmentAccess returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasEnvironmentAccess() bool {
	if o != nil && o.EnvironmentAccess.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentAccess gets a reference to the given NullableString and assigns it to the EnvironmentAccess field.
func (o *EnvironmentAccessFrontendJsonhal) SetEnvironmentAccess(v string) {
	o.EnvironmentAccess.Set(&v)
}

// SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetEnvironmentAccessNil() {
	o.EnvironmentAccess.Set(nil)
}

// UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetEnvironmentAccess() {
	o.EnvironmentAccess.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *EnvironmentAccessFrontendJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *EnvironmentAccessFrontendJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *EnvironmentAccessFrontendJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentAccessFrontendJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessFrontendJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvironmentAccessFrontendJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessFrontendJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentAccessFrontendJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentAccessFrontendJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentAccessFrontendJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.EnvironmentAccess.IsSet() {
		toSerialize["environmentAccess"] = o.EnvironmentAccess.Get()
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

type NullableEnvironmentAccessFrontendJsonhal struct {
	value *EnvironmentAccessFrontendJsonhal
	isSet bool
}

func (v NullableEnvironmentAccessFrontendJsonhal) Get() *EnvironmentAccessFrontendJsonhal {
	return v.value
}

func (v *NullableEnvironmentAccessFrontendJsonhal) Set(val *EnvironmentAccessFrontendJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentAccessFrontendJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentAccessFrontendJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentAccessFrontendJsonhal(val *EnvironmentAccessFrontendJsonhal) *NullableEnvironmentAccessFrontendJsonhal {
	return &NullableEnvironmentAccessFrontendJsonhal{value: val, isSet: true}
}

func (v NullableEnvironmentAccessFrontendJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentAccessFrontendJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
