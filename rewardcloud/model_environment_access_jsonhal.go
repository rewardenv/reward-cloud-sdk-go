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

// checks if the EnvironmentAccessJsonhal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentAccessJsonhal{}

// EnvironmentAccessJsonhal Class EnvironmentAccess
type EnvironmentAccessJsonhal struct {
	Links *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	Environment NullableString `json:"environment,omitempty"`
	Frontend NullableString `json:"frontend,omitempty"`
	Backend NullableString `json:"backend,omitempty"`
	Database NullableString `json:"database,omitempty"`
	DevTools NullableString `json:"devTools,omitempty"`
	Redis NullableString `json:"redis,omitempty"`
	Rabbit NullableString `json:"rabbit,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewEnvironmentAccessJsonhal instantiates a new EnvironmentAccessJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentAccessJsonhal() *EnvironmentAccessJsonhal {
	this := EnvironmentAccessJsonhal{}
	return &this
}

// NewEnvironmentAccessJsonhalWithDefaults instantiates a new EnvironmentAccessJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentAccessJsonhalWithDefaults() *EnvironmentAccessJsonhal {
	this := EnvironmentAccessJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnvironmentAccessJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *EnvironmentAccessJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentAccessJsonhal) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EnvironmentAccessJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *EnvironmentAccessJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetEnvironment() string {
	if o == nil || IsNil(o.Environment.Get()) {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *EnvironmentAccessJsonhal) SetEnvironment(v string) {
	o.Environment.Set(&v)
}
// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetFrontend returns the Frontend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetFrontend() string {
	if o == nil || IsNil(o.Frontend.Get()) {
		var ret string
		return ret
	}
	return *o.Frontend.Get()
}

// GetFrontendOk returns a tuple with the Frontend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetFrontendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frontend.Get(), o.Frontend.IsSet()
}

// HasFrontend returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasFrontend() bool {
	if o != nil && o.Frontend.IsSet() {
		return true
	}

	return false
}

// SetFrontend gets a reference to the given NullableString and assigns it to the Frontend field.
func (o *EnvironmentAccessJsonhal) SetFrontend(v string) {
	o.Frontend.Set(&v)
}
// SetFrontendNil sets the value for Frontend to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetFrontendNil() {
	o.Frontend.Set(nil)
}

// UnsetFrontend ensures that no value is present for Frontend, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetFrontend() {
	o.Frontend.Unset()
}

// GetBackend returns the Backend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetBackend() string {
	if o == nil || IsNil(o.Backend.Get()) {
		var ret string
		return ret
	}
	return *o.Backend.Get()
}

// GetBackendOk returns a tuple with the Backend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetBackendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Backend.Get(), o.Backend.IsSet()
}

// HasBackend returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasBackend() bool {
	if o != nil && o.Backend.IsSet() {
		return true
	}

	return false
}

// SetBackend gets a reference to the given NullableString and assigns it to the Backend field.
func (o *EnvironmentAccessJsonhal) SetBackend(v string) {
	o.Backend.Set(&v)
}
// SetBackendNil sets the value for Backend to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetBackendNil() {
	o.Backend.Set(nil)
}

// UnsetBackend ensures that no value is present for Backend, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetBackend() {
	o.Backend.Unset()
}

// GetDatabase returns the Database field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetDatabase() string {
	if o == nil || IsNil(o.Database.Get()) {
		var ret string
		return ret
	}
	return *o.Database.Get()
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Database.Get(), o.Database.IsSet()
}

// HasDatabase returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasDatabase() bool {
	if o != nil && o.Database.IsSet() {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given NullableString and assigns it to the Database field.
func (o *EnvironmentAccessJsonhal) SetDatabase(v string) {
	o.Database.Set(&v)
}
// SetDatabaseNil sets the value for Database to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetDatabaseNil() {
	o.Database.Set(nil)
}

// UnsetDatabase ensures that no value is present for Database, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetDatabase() {
	o.Database.Unset()
}

// GetDevTools returns the DevTools field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetDevTools() string {
	if o == nil || IsNil(o.DevTools.Get()) {
		var ret string
		return ret
	}
	return *o.DevTools.Get()
}

// GetDevToolsOk returns a tuple with the DevTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetDevToolsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DevTools.Get(), o.DevTools.IsSet()
}

// HasDevTools returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasDevTools() bool {
	if o != nil && o.DevTools.IsSet() {
		return true
	}

	return false
}

// SetDevTools gets a reference to the given NullableString and assigns it to the DevTools field.
func (o *EnvironmentAccessJsonhal) SetDevTools(v string) {
	o.DevTools.Set(&v)
}
// SetDevToolsNil sets the value for DevTools to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetDevToolsNil() {
	o.DevTools.Set(nil)
}

// UnsetDevTools ensures that no value is present for DevTools, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetDevTools() {
	o.DevTools.Unset()
}

// GetRedis returns the Redis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetRedis() string {
	if o == nil || IsNil(o.Redis.Get()) {
		var ret string
		return ret
	}
	return *o.Redis.Get()
}

// GetRedisOk returns a tuple with the Redis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetRedisOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Redis.Get(), o.Redis.IsSet()
}

// HasRedis returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasRedis() bool {
	if o != nil && o.Redis.IsSet() {
		return true
	}

	return false
}

// SetRedis gets a reference to the given NullableString and assigns it to the Redis field.
func (o *EnvironmentAccessJsonhal) SetRedis(v string) {
	o.Redis.Set(&v)
}
// SetRedisNil sets the value for Redis to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetRedisNil() {
	o.Redis.Set(nil)
}

// UnsetRedis ensures that no value is present for Redis, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetRedis() {
	o.Redis.Unset()
}

// GetRabbit returns the Rabbit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetRabbit() string {
	if o == nil || IsNil(o.Rabbit.Get()) {
		var ret string
		return ret
	}
	return *o.Rabbit.Get()
}

// GetRabbitOk returns a tuple with the Rabbit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetRabbitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rabbit.Get(), o.Rabbit.IsSet()
}

// HasRabbit returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasRabbit() bool {
	if o != nil && o.Rabbit.IsSet() {
		return true
	}

	return false
}

// SetRabbit gets a reference to the given NullableString and assigns it to the Rabbit field.
func (o *EnvironmentAccessJsonhal) SetRabbit(v string) {
	o.Rabbit.Set(&v)
}
// SetRabbitNil sets the value for Rabbit to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetRabbitNil() {
	o.Rabbit.Set(nil)
}

// UnsetRabbit ensures that no value is present for Rabbit, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetRabbit() {
	o.Rabbit.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *EnvironmentAccessJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessJsonhal) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *EnvironmentAccessJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *EnvironmentAccessJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *EnvironmentAccessJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentAccessJsonhal) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvironmentAccessJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentAccessJsonhal) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentAccessJsonhal) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentAccessJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentAccessJsonhal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentAccessJsonhal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	// skip: id is readOnly
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if o.Frontend.IsSet() {
		toSerialize["frontend"] = o.Frontend.Get()
	}
	if o.Backend.IsSet() {
		toSerialize["backend"] = o.Backend.Get()
	}
	if o.Database.IsSet() {
		toSerialize["database"] = o.Database.Get()
	}
	if o.DevTools.IsSet() {
		toSerialize["devTools"] = o.DevTools.Get()
	}
	if o.Redis.IsSet() {
		toSerialize["redis"] = o.Redis.Get()
	}
	if o.Rabbit.IsSet() {
		toSerialize["rabbit"] = o.Rabbit.Get()
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

type NullableEnvironmentAccessJsonhal struct {
	value *EnvironmentAccessJsonhal
	isSet bool
}

func (v NullableEnvironmentAccessJsonhal) Get() *EnvironmentAccessJsonhal {
	return v.value
}

func (v *NullableEnvironmentAccessJsonhal) Set(val *EnvironmentAccessJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentAccessJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentAccessJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentAccessJsonhal(val *EnvironmentAccessJsonhal) *NullableEnvironmentAccessJsonhal {
	return &NullableEnvironmentAccessJsonhal{value: val, isSet: true}
}

func (v NullableEnvironmentAccessJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentAccessJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


