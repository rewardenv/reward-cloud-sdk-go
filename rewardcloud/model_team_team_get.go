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

// checks if the TeamTeamGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamTeamGet{}

// TeamTeamGet Class Team
type TeamTeamGet struct {
	Id *int32 `json:"id,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Name NullableString `json:"name,omitempty"`
	CodeName NullableString `json:"codeName,omitempty"`
	IsDefault NullableBool `json:"isDefault,omitempty"`
	User []string `json:"user,omitempty"`
	Project []string `json:"project,omitempty"`
	Organisation NullableString `json:"organisation,omitempty"`
	TeamEnvVar []TeamEnvVarTeamGet `json:"teamEnvVar,omitempty"`
}

// NewTeamTeamGet instantiates a new TeamTeamGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamTeamGet() *TeamTeamGet {
	this := TeamTeamGet{}
	return &this
}

// NewTeamTeamGetWithDefaults instantiates a new TeamTeamGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamTeamGetWithDefaults() *TeamTeamGet {
	this := TeamTeamGet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamTeamGet) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamGet) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamTeamGet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TeamTeamGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TeamTeamGet) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamGet) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TeamTeamGet) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TeamTeamGet) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamGet) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TeamTeamGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TeamTeamGet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TeamTeamGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TeamTeamGet) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamGet) GetCodeName() string {
	if o == nil || IsNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamGet) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *TeamTeamGet) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *TeamTeamGet) SetCodeName(v string) {
	o.CodeName.Set(&v)
}
// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *TeamTeamGet) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *TeamTeamGet) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamGet) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamGet) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *TeamTeamGet) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *TeamTeamGet) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}
// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *TeamTeamGet) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *TeamTeamGet) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TeamTeamGet) GetUser() []string {
	if o == nil || IsNil(o.User) {
		var ret []string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamGet) GetUserOk() ([]string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TeamTeamGet) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []string and assigns it to the User field.
func (o *TeamTeamGet) SetUser(v []string) {
	o.User = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *TeamTeamGet) GetProject() []string {
	if o == nil || IsNil(o.Project) {
		var ret []string
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamGet) GetProjectOk() ([]string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *TeamTeamGet) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given []string and assigns it to the Project field.
func (o *TeamTeamGet) SetProject(v []string) {
	o.Project = v
}

// GetOrganisation returns the Organisation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamGet) GetOrganisation() string {
	if o == nil || IsNil(o.Organisation.Get()) {
		var ret string
		return ret
	}
	return *o.Organisation.Get()
}

// GetOrganisationOk returns a tuple with the Organisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamGet) GetOrganisationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organisation.Get(), o.Organisation.IsSet()
}

// HasOrganisation returns a boolean if a field has been set.
func (o *TeamTeamGet) HasOrganisation() bool {
	if o != nil && o.Organisation.IsSet() {
		return true
	}

	return false
}

// SetOrganisation gets a reference to the given NullableString and assigns it to the Organisation field.
func (o *TeamTeamGet) SetOrganisation(v string) {
	o.Organisation.Set(&v)
}
// SetOrganisationNil sets the value for Organisation to be an explicit nil
func (o *TeamTeamGet) SetOrganisationNil() {
	o.Organisation.Set(nil)
}

// UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
func (o *TeamTeamGet) UnsetOrganisation() {
	o.Organisation.Unset()
}

// GetTeamEnvVar returns the TeamEnvVar field value if set, zero value otherwise.
func (o *TeamTeamGet) GetTeamEnvVar() []TeamEnvVarTeamGet {
	if o == nil || IsNil(o.TeamEnvVar) {
		var ret []TeamEnvVarTeamGet
		return ret
	}
	return o.TeamEnvVar
}

// GetTeamEnvVarOk returns a tuple with the TeamEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamGet) GetTeamEnvVarOk() ([]TeamEnvVarTeamGet, bool) {
	if o == nil || IsNil(o.TeamEnvVar) {
		return nil, false
	}
	return o.TeamEnvVar, true
}

// HasTeamEnvVar returns a boolean if a field has been set.
func (o *TeamTeamGet) HasTeamEnvVar() bool {
	if o != nil && !IsNil(o.TeamEnvVar) {
		return true
	}

	return false
}

// SetTeamEnvVar gets a reference to the given []TeamEnvVarTeamGet and assigns it to the TeamEnvVar field.
func (o *TeamTeamGet) SetTeamEnvVar(v []TeamEnvVarTeamGet) {
	o.TeamEnvVar = v
}

func (o TeamTeamGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamTeamGet) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if o.Organisation.IsSet() {
		toSerialize["organisation"] = o.Organisation.Get()
	}
	if !IsNil(o.TeamEnvVar) {
		toSerialize["teamEnvVar"] = o.TeamEnvVar
	}
	return toSerialize, nil
}

type NullableTeamTeamGet struct {
	value *TeamTeamGet
	isSet bool
}

func (v NullableTeamTeamGet) Get() *TeamTeamGet {
	return v.value
}

func (v *NullableTeamTeamGet) Set(val *TeamTeamGet) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamTeamGet) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamTeamGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamTeamGet(val *TeamTeamGet) *NullableTeamTeamGet {
	return &NullableTeamTeamGet{value: val, isSet: true}
}

func (v NullableTeamTeamGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamTeamGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


