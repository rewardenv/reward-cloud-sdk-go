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

// checks if the TeamJsonhalTeamGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamJsonhalTeamGet{}

// TeamJsonhalTeamGet Class Team
type TeamJsonhalTeamGet struct {
	Links        *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id           *int32                           `json:"id,omitempty"`
	Uuid         *string                          `json:"uuid,omitempty"`
	Name         NullableString                   `json:"name,omitempty"`
	CodeName     NullableString                   `json:"codeName,omitempty"`
	IsDefault    NullableBool                     `json:"isDefault,omitempty"`
	User         []string                         `json:"user,omitempty"`
	Project      []string                         `json:"project,omitempty"`
	Organisation NullableString                   `json:"organisation,omitempty"`
	TeamEnvVar   []TeamEnvVarJsonhalTeamGet       `json:"teamEnvVar,omitempty"`
}

// NewTeamJsonhalTeamGet instantiates a new TeamJsonhalTeamGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamJsonhalTeamGet() *TeamJsonhalTeamGet {
	this := TeamJsonhalTeamGet{}
	return &this
}

// NewTeamJsonhalTeamGetWithDefaults instantiates a new TeamJsonhalTeamGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamJsonhalTeamGetWithDefaults() *TeamJsonhalTeamGet {
	this := TeamJsonhalTeamGet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TeamJsonhalTeamGet) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamJsonhalTeamGet) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *TeamJsonhalTeamGet) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamJsonhalTeamGet) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamJsonhalTeamGet) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *TeamJsonhalTeamGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *TeamJsonhalTeamGet) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamJsonhalTeamGet) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *TeamJsonhalTeamGet) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamJsonhalTeamGet) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamJsonhalTeamGet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TeamJsonhalTeamGet) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TeamJsonhalTeamGet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TeamJsonhalTeamGet) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamJsonhalTeamGet) GetCodeName() string {
	if o == nil || IsNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamJsonhalTeamGet) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *TeamJsonhalTeamGet) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *TeamJsonhalTeamGet) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *TeamJsonhalTeamGet) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamJsonhalTeamGet) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamJsonhalTeamGet) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *TeamJsonhalTeamGet) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *TeamJsonhalTeamGet) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *TeamJsonhalTeamGet) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TeamJsonhalTeamGet) GetUser() []string {
	if o == nil || IsNil(o.User) {
		var ret []string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamJsonhalTeamGet) GetUserOk() ([]string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []string and assigns it to the User field.
func (o *TeamJsonhalTeamGet) SetUser(v []string) {
	o.User = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *TeamJsonhalTeamGet) GetProject() []string {
	if o == nil || IsNil(o.Project) {
		var ret []string
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamJsonhalTeamGet) GetProjectOk() ([]string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given []string and assigns it to the Project field.
func (o *TeamJsonhalTeamGet) SetProject(v []string) {
	o.Project = v
}

// GetOrganisation returns the Organisation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamJsonhalTeamGet) GetOrganisation() string {
	if o == nil || IsNil(o.Organisation.Get()) {
		var ret string
		return ret
	}
	return *o.Organisation.Get()
}

// GetOrganisationOk returns a tuple with the Organisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamJsonhalTeamGet) GetOrganisationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organisation.Get(), o.Organisation.IsSet()
}

// HasOrganisation returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasOrganisation() bool {
	if o != nil && o.Organisation.IsSet() {
		return true
	}

	return false
}

// SetOrganisation gets a reference to the given NullableString and assigns it to the Organisation field.
func (o *TeamJsonhalTeamGet) SetOrganisation(v string) {
	o.Organisation.Set(&v)
}

// SetOrganisationNil sets the value for Organisation to be an explicit nil
func (o *TeamJsonhalTeamGet) SetOrganisationNil() {
	o.Organisation.Set(nil)
}

// UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
func (o *TeamJsonhalTeamGet) UnsetOrganisation() {
	o.Organisation.Unset()
}

// GetTeamEnvVar returns the TeamEnvVar field value if set, zero value otherwise.
func (o *TeamJsonhalTeamGet) GetTeamEnvVar() []TeamEnvVarJsonhalTeamGet {
	if o == nil || IsNil(o.TeamEnvVar) {
		var ret []TeamEnvVarJsonhalTeamGet
		return ret
	}
	return o.TeamEnvVar
}

// GetTeamEnvVarOk returns a tuple with the TeamEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamJsonhalTeamGet) GetTeamEnvVarOk() ([]TeamEnvVarJsonhalTeamGet, bool) {
	if o == nil || IsNil(o.TeamEnvVar) {
		return nil, false
	}
	return o.TeamEnvVar, true
}

// HasTeamEnvVar returns a boolean if a field has been set.
func (o *TeamJsonhalTeamGet) HasTeamEnvVar() bool {
	if o != nil && !IsNil(o.TeamEnvVar) {
		return true
	}

	return false
}

// SetTeamEnvVar gets a reference to the given []TeamEnvVarJsonhalTeamGet and assigns it to the TeamEnvVar field.
func (o *TeamJsonhalTeamGet) SetTeamEnvVar(v []TeamEnvVarJsonhalTeamGet) {
	o.TeamEnvVar = v
}

func (o TeamJsonhalTeamGet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamJsonhalTeamGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
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

type NullableTeamJsonhalTeamGet struct {
	value *TeamJsonhalTeamGet
	isSet bool
}

func (v NullableTeamJsonhalTeamGet) Get() *TeamJsonhalTeamGet {
	return v.value
}

func (v *NullableTeamJsonhalTeamGet) Set(val *TeamJsonhalTeamGet) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamJsonhalTeamGet) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamJsonhalTeamGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamJsonhalTeamGet(val *TeamJsonhalTeamGet) *NullableTeamJsonhalTeamGet {
	return &NullableTeamJsonhalTeamGet{value: val, isSet: true}
}

func (v NullableTeamJsonhalTeamGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamJsonhalTeamGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
