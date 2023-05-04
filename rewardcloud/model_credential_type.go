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

// CredentialType Class CredentialType
type CredentialType struct {
	Id        *int32         `json:"id,omitempty"`
	Uuid      NullableString `json:"uuid,omitempty"`
	Name      NullableString `json:"name,omitempty"`
	CodeName  NullableString `json:"codeName,omitempty"`
	Git       []string       `json:"git,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	CreatedAt *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt *time.Time     `json:"updatedAt,omitempty"`
}

// NewCredentialType instantiates a new CredentialType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialType() *CredentialType {
	this := CredentialType{}
	return &this
}

// NewCredentialTypeWithDefaults instantiates a new CredentialType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialTypeWithDefaults() *CredentialType {
	this := CredentialType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialType) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialType) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialType) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CredentialType) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialType) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialType) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *CredentialType) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *CredentialType) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *CredentialType) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *CredentialType) UnsetUuid() {
	o.Uuid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialType) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CredentialType) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CredentialType) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CredentialType) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CredentialType) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialType) GetCodeName() string {
	if o == nil || isNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialType) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *CredentialType) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *CredentialType) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *CredentialType) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *CredentialType) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *CredentialType) GetGit() []string {
	if o == nil || isNil(o.Git) {
		var ret []string
		return ret
	}
	return o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialType) GetGitOk() ([]string, bool) {
	if o == nil || isNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *CredentialType) HasGit() bool {
	if o != nil && !isNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given []string and assigns it to the Git field.
func (o *CredentialType) SetGit(v []string) {
	o.Git = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialType) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialType) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CredentialType) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *CredentialType) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *CredentialType) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *CredentialType) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialType) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialType) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *CredentialType) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *CredentialType) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *CredentialType) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *CredentialType) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CredentialType) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CredentialType) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CredentialType) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CredentialType) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialType) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CredentialType) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CredentialType) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CredentialType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if !isNil(o.Git) {
		toSerialize["git"] = o.Git
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

type NullableCredentialType struct {
	value *CredentialType
	isSet bool
}

func (v NullableCredentialType) Get() *CredentialType {
	return v.value
}

func (v *NullableCredentialType) Set(val *CredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialType(val *CredentialType) *NullableCredentialType {
	return &NullableCredentialType{value: val, isSet: true}
}

func (v NullableCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
