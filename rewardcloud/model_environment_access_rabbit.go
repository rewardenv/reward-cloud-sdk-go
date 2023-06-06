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

// checks if the EnvironmentAccessRabbit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentAccessRabbit{}

// EnvironmentAccessRabbit Class EnvironmentAccessRabbit
type EnvironmentAccessRabbit struct {
	Id                *int32         `json:"id,omitempty"`
	Uuid              NullableString `json:"uuid,omitempty"`
	Url               NullableString `json:"url,omitempty"`
	Host              NullableString `json:"host,omitempty"`
	Port              NullableInt32  `json:"port,omitempty"`
	Username          NullableString `json:"username,omitempty"`
	Password          NullableString `json:"password,omitempty"`
	EnvironmentAccess NullableString `json:"environmentAccess,omitempty"`
	CreatedBy         NullableString `json:"createdBy,omitempty"`
	UpdatedBy         NullableString `json:"updatedBy,omitempty"`
	CreatedAt         *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt         *time.Time     `json:"updatedAt,omitempty"`
}

// NewEnvironmentAccessRabbit instantiates a new EnvironmentAccessRabbit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentAccessRabbit() *EnvironmentAccessRabbit {
	this := EnvironmentAccessRabbit{}
	return &this
}

// NewEnvironmentAccessRabbitWithDefaults instantiates a new EnvironmentAccessRabbit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentAccessRabbitWithDefaults() *EnvironmentAccessRabbit {
	this := EnvironmentAccessRabbit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentAccessRabbit) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessRabbit) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EnvironmentAccessRabbit) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *EnvironmentAccessRabbit) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *EnvironmentAccessRabbit) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetUuid() {
	o.Uuid.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *EnvironmentAccessRabbit) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *EnvironmentAccessRabbit) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetUrl() {
	o.Url.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *EnvironmentAccessRabbit) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *EnvironmentAccessRabbit) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *EnvironmentAccessRabbit) SetPort(v int32) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *EnvironmentAccessRabbit) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetPort() {
	o.Port.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *EnvironmentAccessRabbit) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *EnvironmentAccessRabbit) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *EnvironmentAccessRabbit) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *EnvironmentAccessRabbit) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetPassword() {
	o.Password.Unset()
}

// GetEnvironmentAccess returns the EnvironmentAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetEnvironmentAccess() string {
	if o == nil || IsNil(o.EnvironmentAccess.Get()) {
		var ret string
		return ret
	}
	return *o.EnvironmentAccess.Get()
}

// GetEnvironmentAccessOk returns a tuple with the EnvironmentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetEnvironmentAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentAccess.Get(), o.EnvironmentAccess.IsSet()
}

// HasEnvironmentAccess returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasEnvironmentAccess() bool {
	if o != nil && o.EnvironmentAccess.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentAccess gets a reference to the given NullableString and assigns it to the EnvironmentAccess field.
func (o *EnvironmentAccessRabbit) SetEnvironmentAccess(v string) {
	o.EnvironmentAccess.Set(&v)
}

// SetEnvironmentAccessNil sets the value for EnvironmentAccess to be an explicit nil
func (o *EnvironmentAccessRabbit) SetEnvironmentAccessNil() {
	o.EnvironmentAccess.Set(nil)
}

// UnsetEnvironmentAccess ensures that no value is present for EnvironmentAccess, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetEnvironmentAccess() {
	o.EnvironmentAccess.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *EnvironmentAccessRabbit) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *EnvironmentAccessRabbit) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentAccessRabbit) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentAccessRabbit) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *EnvironmentAccessRabbit) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *EnvironmentAccessRabbit) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *EnvironmentAccessRabbit) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EnvironmentAccessRabbit) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessRabbit) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EnvironmentAccessRabbit) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EnvironmentAccessRabbit) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentAccessRabbit) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EnvironmentAccessRabbit) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *EnvironmentAccessRabbit) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o EnvironmentAccessRabbit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentAccessRabbit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableEnvironmentAccessRabbit struct {
	value *EnvironmentAccessRabbit
	isSet bool
}

func (v NullableEnvironmentAccessRabbit) Get() *EnvironmentAccessRabbit {
	return v.value
}

func (v *NullableEnvironmentAccessRabbit) Set(val *EnvironmentAccessRabbit) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentAccessRabbit) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentAccessRabbit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentAccessRabbit(val *EnvironmentAccessRabbit) *NullableEnvironmentAccessRabbit {
	return &NullableEnvironmentAccessRabbit{value: val, isSet: true}
}

func (v NullableEnvironmentAccessRabbit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentAccessRabbit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
