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

// TeamTeamPost Class Team
type TeamTeamPost struct {
	Name         NullableString       `json:"name,omitempty"`
	CodeName     NullableString       `json:"codeName,omitempty"`
	IsDefault    NullableBool         `json:"isDefault,omitempty"`
	User         []string             `json:"user,omitempty"`
	Project      []string             `json:"project,omitempty"`
	Organisation NullableString       `json:"organisation,omitempty"`
	TeamEnvVar   []TeamEnvVarTeamPost `json:"teamEnvVar,omitempty"`
}

// NewTeamTeamPost instantiates a new TeamTeamPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamTeamPost() *TeamTeamPost {
	this := TeamTeamPost{}
	return &this
}

// NewTeamTeamPostWithDefaults instantiates a new TeamTeamPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamTeamPostWithDefaults() *TeamTeamPost {
	this := TeamTeamPost{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamPost) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TeamTeamPost) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TeamTeamPost) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *TeamTeamPost) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TeamTeamPost) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamPost) GetCodeName() string {
	if o == nil || isNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamPost) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *TeamTeamPost) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *TeamTeamPost) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *TeamTeamPost) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *TeamTeamPost) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamPost) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamPost) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *TeamTeamPost) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *TeamTeamPost) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *TeamTeamPost) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *TeamTeamPost) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TeamTeamPost) GetUser() []string {
	if o == nil || isNil(o.User) {
		var ret []string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamPost) GetUserOk() ([]string, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TeamTeamPost) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []string and assigns it to the User field.
func (o *TeamTeamPost) SetUser(v []string) {
	o.User = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *TeamTeamPost) GetProject() []string {
	if o == nil || isNil(o.Project) {
		var ret []string
		return ret
	}
	return o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamPost) GetProjectOk() ([]string, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *TeamTeamPost) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given []string and assigns it to the Project field.
func (o *TeamTeamPost) SetProject(v []string) {
	o.Project = v
}

// GetOrganisation returns the Organisation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamTeamPost) GetOrganisation() string {
	if o == nil || isNil(o.Organisation.Get()) {
		var ret string
		return ret
	}
	return *o.Organisation.Get()
}

// GetOrganisationOk returns a tuple with the Organisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamTeamPost) GetOrganisationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organisation.Get(), o.Organisation.IsSet()
}

// HasOrganisation returns a boolean if a field has been set.
func (o *TeamTeamPost) HasOrganisation() bool {
	if o != nil && o.Organisation.IsSet() {
		return true
	}

	return false
}

// SetOrganisation gets a reference to the given NullableString and assigns it to the Organisation field.
func (o *TeamTeamPost) SetOrganisation(v string) {
	o.Organisation.Set(&v)
}

// SetOrganisationNil sets the value for Organisation to be an explicit nil
func (o *TeamTeamPost) SetOrganisationNil() {
	o.Organisation.Set(nil)
}

// UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
func (o *TeamTeamPost) UnsetOrganisation() {
	o.Organisation.Unset()
}

// GetTeamEnvVar returns the TeamEnvVar field value if set, zero value otherwise.
func (o *TeamTeamPost) GetTeamEnvVar() []TeamEnvVarTeamPost {
	if o == nil || isNil(o.TeamEnvVar) {
		var ret []TeamEnvVarTeamPost
		return ret
	}
	return o.TeamEnvVar
}

// GetTeamEnvVarOk returns a tuple with the TeamEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamTeamPost) GetTeamEnvVarOk() ([]TeamEnvVarTeamPost, bool) {
	if o == nil || isNil(o.TeamEnvVar) {
		return nil, false
	}
	return o.TeamEnvVar, true
}

// HasTeamEnvVar returns a boolean if a field has been set.
func (o *TeamTeamPost) HasTeamEnvVar() bool {
	if o != nil && !isNil(o.TeamEnvVar) {
		return true
	}

	return false
}

// SetTeamEnvVar gets a reference to the given []TeamEnvVarTeamPost and assigns it to the TeamEnvVar field.
func (o *TeamTeamPost) SetTeamEnvVar(v []TeamEnvVarTeamPost) {
	o.TeamEnvVar = v
}

func (o TeamTeamPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CodeName.IsSet() {
		toSerialize["codeName"] = o.CodeName.Get()
	}
	if o.IsDefault.IsSet() {
		toSerialize["isDefault"] = o.IsDefault.Get()
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if o.Organisation.IsSet() {
		toSerialize["organisation"] = o.Organisation.Get()
	}
	if !isNil(o.TeamEnvVar) {
		toSerialize["teamEnvVar"] = o.TeamEnvVar
	}
	return json.Marshal(toSerialize)
}

type NullableTeamTeamPost struct {
	value *TeamTeamPost
	isSet bool
}

func (v NullableTeamTeamPost) Get() *TeamTeamPost {
	return v.value
}

func (v *NullableTeamTeamPost) Set(val *TeamTeamPost) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamTeamPost) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamTeamPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamTeamPost(val *TeamTeamPost) *NullableTeamTeamPost {
	return &NullableTeamTeamPost{value: val, isSet: true}
}

func (v NullableTeamTeamPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamTeamPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
