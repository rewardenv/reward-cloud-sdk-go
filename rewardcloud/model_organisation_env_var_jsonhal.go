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

// OrganisationEnvVarJsonhal Class OrganisationEnvVar
type OrganisationEnvVarJsonhal struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	Key NullableString `json:"key,omitempty"`
	Value NullableString `json:"value,omitempty"`
	IsEncrypted NullableBool `json:"isEncrypted,omitempty"`
	Organisation NullableString `json:"organisation,omitempty"`
	EnvVarType NullableString `json:"envVarType,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	RawValue NullableString `json:"rawValue,omitempty"`
}

// NewOrganisationEnvVarJsonhal instantiates a new OrganisationEnvVarJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationEnvVarJsonhal() *OrganisationEnvVarJsonhal {
	this := OrganisationEnvVarJsonhal{}
	return &this
}

// NewOrganisationEnvVarJsonhalWithDefaults instantiates a new OrganisationEnvVarJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationEnvVarJsonhalWithDefaults() *OrganisationEnvVarJsonhal {
	this := OrganisationEnvVarJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganisationEnvVarJsonhal) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationEnvVarJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *OrganisationEnvVarJsonhal) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganisationEnvVarJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationEnvVarJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganisationEnvVarJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *OrganisationEnvVarJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetKey() string {
	if o == nil || isNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *OrganisationEnvVarJsonhal) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetKey() {
	o.Key.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetValue() string {
	if o == nil || isNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *OrganisationEnvVarJsonhal) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetValue() {
	o.Value.Unset()
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetIsEncrypted() bool {
	if o == nil || isNil(o.IsEncrypted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted.Get()
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetIsEncryptedOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsEncrypted.Get(), o.IsEncrypted.IsSet()
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasIsEncrypted() bool {
	if o != nil && o.IsEncrypted.IsSet() {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given NullableBool and assigns it to the IsEncrypted field.
func (o *OrganisationEnvVarJsonhal) SetIsEncrypted(v bool) {
	o.IsEncrypted.Set(&v)
}
// SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetIsEncryptedNil() {
	o.IsEncrypted.Set(nil)
}

// UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetIsEncrypted() {
	o.IsEncrypted.Unset()
}

// GetOrganisation returns the Organisation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetOrganisation() string {
	if o == nil || isNil(o.Organisation.Get()) {
		var ret string
		return ret
	}
	return *o.Organisation.Get()
}

// GetOrganisationOk returns a tuple with the Organisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetOrganisationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Organisation.Get(), o.Organisation.IsSet()
}

// HasOrganisation returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasOrganisation() bool {
	if o != nil && o.Organisation.IsSet() {
		return true
	}

	return false
}

// SetOrganisation gets a reference to the given NullableString and assigns it to the Organisation field.
func (o *OrganisationEnvVarJsonhal) SetOrganisation(v string) {
	o.Organisation.Set(&v)
}
// SetOrganisationNil sets the value for Organisation to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetOrganisationNil() {
	o.Organisation.Set(nil)
}

// UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetOrganisation() {
	o.Organisation.Unset()
}

// GetEnvVarType returns the EnvVarType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetEnvVarType() string {
	if o == nil || isNil(o.EnvVarType.Get()) {
		var ret string
		return ret
	}
	return *o.EnvVarType.Get()
}

// GetEnvVarTypeOk returns a tuple with the EnvVarType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetEnvVarTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.EnvVarType.Get(), o.EnvVarType.IsSet()
}

// HasEnvVarType returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasEnvVarType() bool {
	if o != nil && o.EnvVarType.IsSet() {
		return true
	}

	return false
}

// SetEnvVarType gets a reference to the given NullableString and assigns it to the EnvVarType field.
func (o *OrganisationEnvVarJsonhal) SetEnvVarType(v string) {
	o.EnvVarType.Set(&v)
}
// SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetEnvVarTypeNil() {
	o.EnvVarType.Set(nil)
}

// UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetEnvVarType() {
	o.EnvVarType.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *OrganisationEnvVarJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *OrganisationEnvVarJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrganisationEnvVarJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationEnvVarJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrganisationEnvVarJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganisationEnvVarJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationEnvVarJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganisationEnvVarJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRawValue returns the RawValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationEnvVarJsonhal) GetRawValue() string {
	if o == nil || isNil(o.RawValue.Get()) {
		var ret string
		return ret
	}
	return *o.RawValue.Get()
}

// GetRawValueOk returns a tuple with the RawValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationEnvVarJsonhal) GetRawValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.RawValue.Get(), o.RawValue.IsSet()
}

// HasRawValue returns a boolean if a field has been set.
func (o *OrganisationEnvVarJsonhal) HasRawValue() bool {
	if o != nil && o.RawValue.IsSet() {
		return true
	}

	return false
}

// SetRawValue gets a reference to the given NullableString and assigns it to the RawValue field.
func (o *OrganisationEnvVarJsonhal) SetRawValue(v string) {
	o.RawValue.Set(&v)
}
// SetRawValueNil sets the value for RawValue to be an explicit nil
func (o *OrganisationEnvVarJsonhal) SetRawValueNil() {
	o.RawValue.Set(nil)
}

// UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil
func (o *OrganisationEnvVarJsonhal) UnsetRawValue() {
	o.RawValue.Unset()
}

func (o OrganisationEnvVarJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.IsEncrypted.IsSet() {
		toSerialize["isEncrypted"] = o.IsEncrypted.Get()
	}
	if o.Organisation.IsSet() {
		toSerialize["organisation"] = o.Organisation.Get()
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
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.RawValue.IsSet() {
		toSerialize["rawValue"] = o.RawValue.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganisationEnvVarJsonhal struct {
	value *OrganisationEnvVarJsonhal
	isSet bool
}

func (v NullableOrganisationEnvVarJsonhal) Get() *OrganisationEnvVarJsonhal {
	return v.value
}

func (v *NullableOrganisationEnvVarJsonhal) Set(val *OrganisationEnvVarJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationEnvVarJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationEnvVarJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationEnvVarJsonhal(val *OrganisationEnvVarJsonhal) *NullableOrganisationEnvVarJsonhal {
	return &NullableOrganisationEnvVarJsonhal{value: val, isSet: true}
}

func (v NullableOrganisationEnvVarJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationEnvVarJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


