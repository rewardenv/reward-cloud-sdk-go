/*
Reward Cloud

It is an API for Reward Cloud project in ITG Commerce

API version: v0.7.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rewardcloud

import (
	"encoding/json"
)

// checks if the OrganisationOrganisationOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganisationOrganisationOutput{}

// OrganisationOrganisationOutput Class Organisation
type OrganisationOrganisationOutput struct {
	Id       *int32         `json:"id,omitempty"`
	Uuid     *string        `json:"uuid,omitempty"`
	Name     NullableString `json:"name,omitempty"`
	CodeName NullableString `json:"codeName,omitempty"`
	// User related property - One user can only have 1 pcs default Organisation
	IsDefault          NullableBool                           `json:"isDefault,omitempty"`
	Team               []string                               `json:"team,omitempty"`
	OrganisationEnvVar []OrganisationEnvVarOrganisationOutput `json:"organisationEnvVar,omitempty"`
}

// NewOrganisationOrganisationOutput instantiates a new OrganisationOrganisationOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationOrganisationOutput() *OrganisationOrganisationOutput {
	this := OrganisationOrganisationOutput{}
	return &this
}

// NewOrganisationOrganisationOutputWithDefaults instantiates a new OrganisationOrganisationOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationOrganisationOutputWithDefaults() *OrganisationOrganisationOutput {
	this := OrganisationOrganisationOutput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganisationOrganisationOutput) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationOutput) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganisationOrganisationOutput) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *OrganisationOrganisationOutput) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationOutput) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *OrganisationOrganisationOutput) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationOrganisationOutput) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationOrganisationOutput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrganisationOrganisationOutput) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *OrganisationOrganisationOutput) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrganisationOrganisationOutput) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationOrganisationOutput) GetCodeName() string {
	if o == nil || IsNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationOrganisationOutput) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *OrganisationOrganisationOutput) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *OrganisationOrganisationOutput) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *OrganisationOrganisationOutput) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationOrganisationOutput) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationOrganisationOutput) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *OrganisationOrganisationOutput) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *OrganisationOrganisationOutput) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *OrganisationOrganisationOutput) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *OrganisationOrganisationOutput) GetTeam() []string {
	if o == nil || IsNil(o.Team) {
		var ret []string
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationOutput) GetTeamOk() ([]string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given []string and assigns it to the Team field.
func (o *OrganisationOrganisationOutput) SetTeam(v []string) {
	o.Team = v
}

// GetOrganisationEnvVar returns the OrganisationEnvVar field value if set, zero value otherwise.
func (o *OrganisationOrganisationOutput) GetOrganisationEnvVar() []OrganisationEnvVarOrganisationOutput {
	if o == nil || IsNil(o.OrganisationEnvVar) {
		var ret []OrganisationEnvVarOrganisationOutput
		return ret
	}
	return o.OrganisationEnvVar
}

// GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationOrganisationOutput) GetOrganisationEnvVarOk() ([]OrganisationEnvVarOrganisationOutput, bool) {
	if o == nil || IsNil(o.OrganisationEnvVar) {
		return nil, false
	}
	return o.OrganisationEnvVar, true
}

// HasOrganisationEnvVar returns a boolean if a field has been set.
func (o *OrganisationOrganisationOutput) HasOrganisationEnvVar() bool {
	if o != nil && !IsNil(o.OrganisationEnvVar) {
		return true
	}

	return false
}

// SetOrganisationEnvVar gets a reference to the given []OrganisationEnvVarOrganisationOutput and assigns it to the OrganisationEnvVar field.
func (o *OrganisationOrganisationOutput) SetOrganisationEnvVar(v []OrganisationEnvVarOrganisationOutput) {
	o.OrganisationEnvVar = v
}

func (o OrganisationOrganisationOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganisationOrganisationOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: uuid is readOnly
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CodeName.IsSet() {
		toSerialize["codeName"] = o.CodeName.Get()
	}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.OrganisationEnvVar) {
		toSerialize["organisationEnvVar"] = o.OrganisationEnvVar
	}
	return toSerialize, nil
}

type NullableOrganisationOrganisationOutput struct {
	value *OrganisationOrganisationOutput
	isSet bool
}

func (v NullableOrganisationOrganisationOutput) Get() *OrganisationOrganisationOutput {
	return v.value
}

func (v *NullableOrganisationOrganisationOutput) Set(val *OrganisationOrganisationOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationOrganisationOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationOrganisationOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationOrganisationOutput(val *OrganisationOrganisationOutput) *NullableOrganisationOrganisationOutput {
	return &NullableOrganisationOrganisationOutput{value: val, isSet: true}
}

func (v NullableOrganisationOrganisationOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationOrganisationOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
