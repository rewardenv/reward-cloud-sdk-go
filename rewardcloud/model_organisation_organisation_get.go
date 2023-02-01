/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.6.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// OrganisationOrganisationGet Class Organisation
type OrganisationOrganisationGet struct {
	Id *int32 `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Name NullableString `json:"name,omitempty"`
	CodeName NullableString `json:"codeName,omitempty"`
	IsDefault NullableBool `json:"isDefault,omitempty"`
	Team []string `json:"team,omitempty"`
	OrganisationEnvVar []OrganisationEnvVarOrganisationGet `json:"organisationEnvVar,omitempty"`
}

// NewOrganisationOrganisationGet instantiates a new OrganisationOrganisationGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationOrganisationGet() *OrganisationOrganisationGet {
	this := OrganisationOrganisationGet{}
	return &this
}

// NewOrganisationOrganisationGetWithDefaults instantiates a new OrganisationOrganisationGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationOrganisationGetWithDefaults() *OrganisationOrganisationGet {
	this := OrganisationOrganisationGet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganisationOrganisationGet) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationGet) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganisationOrganisationGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *OrganisationOrganisationGet) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationGet) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *OrganisationOrganisationGet) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationOrganisationGet) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationOrganisationGet) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrganisationOrganisationGet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OrganisationOrganisationGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrganisationOrganisationGet) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationOrganisationGet) GetCodeName() string {
	if o == nil || isNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationOrganisationGet) GetCodeNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *OrganisationOrganisationGet) SetCodeName(v string) {
	o.CodeName.Set(&v)
}
// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *OrganisationOrganisationGet) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *OrganisationOrganisationGet) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationOrganisationGet) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationOrganisationGet) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *OrganisationOrganisationGet) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}
// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *OrganisationOrganisationGet) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *OrganisationOrganisationGet) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *OrganisationOrganisationGet) GetTeam() []string {
	if o == nil || isNil(o.Team) {
		var ret []string
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationGet) GetTeamOk() ([]string, bool) {
	if o == nil || isNil(o.Team) {
    return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasTeam() bool {
	if o != nil && !isNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given []string and assigns it to the Team field.
func (o *OrganisationOrganisationGet) SetTeam(v []string) {
	o.Team = v
}

// GetOrganisationEnvVar returns the OrganisationEnvVar field value if set, zero value otherwise.
func (o *OrganisationOrganisationGet) GetOrganisationEnvVar() []OrganisationEnvVarOrganisationGet {
	if o == nil || isNil(o.OrganisationEnvVar) {
		var ret []OrganisationEnvVarOrganisationGet
		return ret
	}
	return o.OrganisationEnvVar
}

// GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationGet) GetOrganisationEnvVarOk() ([]OrganisationEnvVarOrganisationGet, bool) {
	if o == nil || isNil(o.OrganisationEnvVar) {
    return nil, false
	}
	return o.OrganisationEnvVar, true
}

// HasOrganisationEnvVar returns a boolean if a field has been set.
func (o *OrganisationOrganisationGet) HasOrganisationEnvVar() bool {
	if o != nil && !isNil(o.OrganisationEnvVar) {
		return true
	}

	return false
}

// SetOrganisationEnvVar gets a reference to the given []OrganisationEnvVarOrganisationGet and assigns it to the OrganisationEnvVar field.
func (o *OrganisationOrganisationGet) SetOrganisationEnvVar(v []OrganisationEnvVarOrganisationGet) {
	o.OrganisationEnvVar = v
}

func (o OrganisationOrganisationGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CodeName.IsSet() {
		toSerialize["codeName"] = o.CodeName.Get()
	}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if !isNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !isNil(o.OrganisationEnvVar) {
		toSerialize["organisationEnvVar"] = o.OrganisationEnvVar
	}
	return json.Marshal(toSerialize)
}

type NullableOrganisationOrganisationGet struct {
	value *OrganisationOrganisationGet
	isSet bool
}

func (v NullableOrganisationOrganisationGet) Get() *OrganisationOrganisationGet {
	return v.value
}

func (v *NullableOrganisationOrganisationGet) Set(val *OrganisationOrganisationGet) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationOrganisationGet) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationOrganisationGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationOrganisationGet(val *OrganisationOrganisationGet) *NullableOrganisationOrganisationGet {
	return &NullableOrganisationOrganisationGet{value: val, isSet: true}
}

func (v NullableOrganisationOrganisationGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationOrganisationGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

