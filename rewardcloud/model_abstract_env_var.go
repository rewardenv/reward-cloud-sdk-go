/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
	"time"
)

// checks if the AbstractEnvVar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbstractEnvVar{}

// AbstractEnvVar
type AbstractEnvVar struct {
	Id          *int32         `json:"id,omitempty"`
	Uuid        NullableString `json:"uuid,omitempty"`
	Key         NullableString `json:"key,omitempty"`
	Value       NullableString `json:"value,omitempty"`
	IsEncrypted NullableBool   `json:"isEncrypted,omitempty"`
	EnvVarType  NullableString `json:"envVarType,omitempty"`
	CreatedBy   NullableString `json:"createdBy,omitempty"`
	UpdatedBy   NullableString `json:"updatedBy,omitempty"`
	CreatedAt   *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time     `json:"updatedAt,omitempty"`
	RawValue    NullableString `json:"rawValue,omitempty"`
}

// NewAbstractEnvVar instantiates a new AbstractEnvVar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbstractEnvVar() *AbstractEnvVar {
	this := AbstractEnvVar{}
	return &this
}

// NewAbstractEnvVarWithDefaults instantiates a new AbstractEnvVar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbstractEnvVarWithDefaults() *AbstractEnvVar {
	this := AbstractEnvVar{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AbstractEnvVar) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractEnvVar) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AbstractEnvVar) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *AbstractEnvVar) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *AbstractEnvVar) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *AbstractEnvVar) UnsetUuid() {
	o.Uuid.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *AbstractEnvVar) SetKey(v string) {
	o.Key.Set(&v)
}

// SetKeyNil sets the value for Key to be an explicit nil
func (o *AbstractEnvVar) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *AbstractEnvVar) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *AbstractEnvVar) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *AbstractEnvVar) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *AbstractEnvVar) UnsetValue() {
	o.Value.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetIsEncrypted() bool {
	if o == nil || IsNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *AbstractEnvVar) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}

// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *AbstractEnvVar) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *AbstractEnvVar) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetEnvVarType() string {
	if o == nil || IsNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *AbstractEnvVar) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}

// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *AbstractEnvVar) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *AbstractEnvVar) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *AbstractEnvVar) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *AbstractEnvVar) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *AbstractEnvVar) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *AbstractEnvVar) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *AbstractEnvVar) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *AbstractEnvVar) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AbstractEnvVar) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractEnvVar) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AbstractEnvVar) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AbstractEnvVar) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbstractEnvVar) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AbstractEnvVar) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRawValue returns the RawValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AbstractEnvVar) GetRawValue() string {
	if o == nil || IsNil(o.RawValue.Get()) {
		var ret string
		return ret
	}
	return *o.RawValue.Get()
}

// GetRawValueOk returns a tuple with the RawValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AbstractEnvVar) GetRawValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawValue.Get(), o.RawValue.IsSet()
}

// HasRawValue returns a boolean if a field has been set.
func (o *AbstractEnvVar) HasRawValue() bool {
	if o != nil && o.RawValue.IsSet() {
		return true
	}

	return false
}

// SetRawValue gets a reference to the given NullableString and assigns it to the RawValue field.
func (o *AbstractEnvVar) SetRawValue(v string) {
	o.RawValue.Set(&v)
}

// SetRawValueNil sets the value for RawValue to be an explicit nil
func (o *AbstractEnvVar) SetRawValueNil() {
	o.RawValue.Set(nil)
}

// UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
func (o *AbstractEnvVar) UnsetRawValue() {
	o.RawValue.Unset()
}

func (o AbstractEnvVar) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbstractEnvVar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsEncrypted.IsSet() {
		toSerialize["isEncrypted"] = o.IsEncrypted.Get()
	}
	if o.EnvVarType.IsSet() {
		toSerialize["envVarType"] = o.EnvVarType.Get()
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
	if o.RawValue.IsSet() {
		toSerialize["rawValue"] = o.RawValue.Get()
	}
	return toSerialize, nil
}

type NullableAbstractEnvVar struct {
	value *AbstractEnvVar
	isSet bool
}

func (v NullableAbstractEnvVar) Get() *AbstractEnvVar {
	return v.value
}

func (v *NullableAbstractEnvVar) Set(val *AbstractEnvVar) {
	v.value = val
	v.isSet = true
}

func (v NullableAbstractEnvVar) IsSet() bool {
	return v.isSet
}

func (v *NullableAbstractEnvVar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbstractEnvVar(val *AbstractEnvVar) *NullableAbstractEnvVar {
	return &NullableAbstractEnvVar{value: val, isSet: true}
}

func (v NullableAbstractEnvVar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbstractEnvVar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
