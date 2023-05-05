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

// checks if the OrganisationJsonhalOrganisationPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganisationJsonhalOrganisationPost{}

// OrganisationJsonhalOrganisationPost Class Organisation
type OrganisationJsonhalOrganisationPost struct {
	Links              *AbstractEnvironmentJsonhalLinks            `json:"_links,omitempty"`
	Name               NullableString                              `json:"name,omitempty"`
	CodeName           NullableString                              `json:"codeName,omitempty"`
	IsDefault          NullableBool                                `json:"isDefault,omitempty"`
	Team               []string                                    `json:"team,omitempty"`
	OrganisationEnvVar []OrganisationEnvVarJsonhalOrganisationPost `json:"organisationEnvVar,omitempty"`
}

// NewOrganisationJsonhalOrganisationPost instantiates a new OrganisationJsonhalOrganisationPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganisationJsonhalOrganisationPost() *OrganisationJsonhalOrganisationPost {
	this := OrganisationJsonhalOrganisationPost{}
	return &this
}

// NewOrganisationJsonhalOrganisationPostWithDefaults instantiates a new OrganisationJsonhalOrganisationPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganisationJsonhalOrganisationPostWithDefaults() *OrganisationJsonhalOrganisationPost {
	this := OrganisationJsonhalOrganisationPost{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganisationJsonhalOrganisationPost) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || IsNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationJsonhalOrganisationPost) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganisationJsonhalOrganisationPost) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *OrganisationJsonhalOrganisationPost) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationJsonhalOrganisationPost) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationJsonhalOrganisationPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrganisationJsonhalOrganisationPost) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrganisationJsonhalOrganisationPost) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *OrganisationJsonhalOrganisationPost) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrganisationJsonhalOrganisationPost) UnsetName() {
	o.Name.Unset()
}

// GetCodeName returns the CodeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationJsonhalOrganisationPost) GetCodeName() string {
	if o == nil || IsNil(o.CodeName.Get()) {
		var ret string
		return ret
	}
	return *o.CodeName.Get()
}

// GetCodeNameOk returns a tuple with the CodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationJsonhalOrganisationPost) GetCodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CodeName.Get(), o.CodeName.IsSet()
}

// HasCodeName returns a boolean if a field has been set.
func (o *OrganisationJsonhalOrganisationPost) HasCodeName() bool {
	if o != nil && o.CodeName.IsSet() {
		return true
	}

	return false
}

// SetCodeName gets a reference to the given NullableString and assigns it to the CodeName field.
func (o *OrganisationJsonhalOrganisationPost) SetCodeName(v string) {
	o.CodeName.Set(&v)
}

// SetCodeNameNil sets the value for CodeName to be an explicit nil
func (o *OrganisationJsonhalOrganisationPost) SetCodeNameNil() {
	o.CodeName.Set(nil)
}

// UnsetCodeName ensures that no value is present for CodeName, not even an explicit nil
func (o *OrganisationJsonhalOrganisationPost) UnsetCodeName() {
	o.CodeName.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganisationJsonhalOrganisationPost) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDefault.Get()
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganisationJsonhalOrganisationPost) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDefault.Get(), o.IsDefault.IsSet()
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OrganisationJsonhalOrganisationPost) HasIsDefault() bool {
	if o != nil && o.IsDefault.IsSet() {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given NullableBool and assigns it to the IsDefault field.
func (o *OrganisationJsonhalOrganisationPost) SetIsDefault(v bool) {
	o.IsDefault.Set(&v)
}

// SetIsDefaultNil sets the value for IsDefault to be an explicit nil
func (o *OrganisationJsonhalOrganisationPost) SetIsDefaultNil() {
	o.IsDefault.Set(nil)
}

// UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
func (o *OrganisationJsonhalOrganisationPost) UnsetIsDefault() {
	o.IsDefault.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *OrganisationJsonhalOrganisationPost) GetTeam() []string {
	if o == nil || IsNil(o.Team) {
		var ret []string
		return ret
	}
	return o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationJsonhalOrganisationPost) GetTeamOk() ([]string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *OrganisationJsonhalOrganisationPost) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given []string and assigns it to the Team field.
func (o *OrganisationJsonhalOrganisationPost) SetTeam(v []string) {
	o.Team = v
}

// GetOrganisationEnvVar returns the OrganisationEnvVar field value if set, zero value otherwise.
func (o *OrganisationJsonhalOrganisationPost) GetOrganisationEnvVar() []OrganisationEnvVarJsonhalOrganisationPost {
	if o == nil || IsNil(o.OrganisationEnvVar) {
		var ret []OrganisationEnvVarJsonhalOrganisationPost
		return ret
	}
	return o.OrganisationEnvVar
}

// GetOrganisationEnvVarOk returns a tuple with the OrganisationEnvVar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganisationJsonhalOrganisationPost) GetOrganisationEnvVarOk() ([]OrganisationEnvVarJsonhalOrganisationPost, bool) {
	if o == nil || IsNil(o.OrganisationEnvVar) {
		return nil, false
	}
	return o.OrganisationEnvVar, true
}

// HasOrganisationEnvVar returns a boolean if a field has been set.
func (o *OrganisationJsonhalOrganisationPost) HasOrganisationEnvVar() bool {
	if o != nil && !IsNil(o.OrganisationEnvVar) {
		return true
	}

	return false
}

// SetOrganisationEnvVar gets a reference to the given []OrganisationEnvVarJsonhalOrganisationPost and assigns it to the OrganisationEnvVar field.
func (o *OrganisationJsonhalOrganisationPost) SetOrganisationEnvVar(v []OrganisationEnvVarJsonhalOrganisationPost) {
	o.OrganisationEnvVar = v
}

func (o OrganisationJsonhalOrganisationPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganisationJsonhalOrganisationPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
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
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.OrganisationEnvVar) {
		toSerialize["organisationEnvVar"] = o.OrganisationEnvVar
	}
	return toSerialize, nil
}

type NullableOrganisationJsonhalOrganisationPost struct {
	value *OrganisationJsonhalOrganisationPost
	isSet bool
}

func (v NullableOrganisationJsonhalOrganisationPost) Get() *OrganisationJsonhalOrganisationPost {
	return v.value
}

func (v *NullableOrganisationJsonhalOrganisationPost) Set(val *OrganisationJsonhalOrganisationPost) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganisationJsonhalOrganisationPost) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganisationJsonhalOrganisationPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganisationJsonhalOrganisationPost(val *OrganisationJsonhalOrganisationPost) *NullableOrganisationJsonhalOrganisationPost {
	return &NullableOrganisationJsonhalOrganisationPost{value: val, isSet: true}
}

func (v NullableOrganisationJsonhalOrganisationPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganisationJsonhalOrganisationPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
