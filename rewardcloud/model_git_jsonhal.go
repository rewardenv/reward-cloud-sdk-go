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

// GitJsonhal Class Git
type GitJsonhal struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	Repo NullableString `json:"repo,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	SshPrivateKey NullableString `json:"sshPrivateKey,omitempty"`
	SshPrivateKeyInput NullableString `json:"sshPrivateKeyInput,omitempty"`
	Project NullableString `json:"project,omitempty"`
	GitType NullableString `json:"gitType,omitempty"`
	CredentialType NullableString `json:"credentialType,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	UpdatedBy NullableString `json:"updatedBy,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewGitJsonhal instantiates a new GitJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitJsonhal() *GitJsonhal {
	this := GitJsonhal{}
	return &this
}

// NewGitJsonhalWithDefaults instantiates a new GitJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitJsonhalWithDefaults() *GitJsonhal {
	this := GitJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GitJsonhal) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitJsonhal) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GitJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *GitJsonhal) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GitJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GitJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GitJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetUuid() string {
	if o == nil || isNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetUuidOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *GitJsonhal) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *GitJsonhal) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *GitJsonhal) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *GitJsonhal) UnsetUuid() {
	o.Uuid.Unset()
}

// GetRepo returns the Repo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetRepo() string {
	if o == nil || isNil(o.Repo.Get()) {
		var ret string
		return ret
	}
	return *o.Repo.Get()
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetRepoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Repo.Get(), o.Repo.IsSet()
}

// HasRepo returns a boolean if a field has been set.
func (o *GitJsonhal) HasRepo() bool {
	if o != nil && o.Repo.IsSet() {
		return true
	}

	return false
}

// SetRepo gets a reference to the given NullableString and assigns it to the Repo field.
func (o *GitJsonhal) SetRepo(v string) {
	o.Repo.Set(&v)
}
// SetRepoNil sets the value for Repo to be an explicit nil
func (o *GitJsonhal) SetRepoNil() {
	o.Repo.Set(nil)
}

// UnsetRepo ensures that no value is present for Repo, not even an explicit nil
func (o *GitJsonhal) UnsetRepo() {
	o.Repo.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetUsername() string {
	if o == nil || isNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *GitJsonhal) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *GitJsonhal) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *GitJsonhal) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *GitJsonhal) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetPassword() string {
	if o == nil || isNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *GitJsonhal) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *GitJsonhal) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *GitJsonhal) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *GitJsonhal) UnsetPassword() {
	o.Password.Unset()
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetSshPrivateKey() string {
	if o == nil || isNil(o.SshPrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey.Get()
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SshPrivateKey.Get(), o.SshPrivateKey.IsSet()
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *GitJsonhal) HasSshPrivateKey() bool {
	if o != nil && o.SshPrivateKey.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given NullableString and assigns it to the SshPrivateKey field.
func (o *GitJsonhal) SetSshPrivateKey(v string) {
	o.SshPrivateKey.Set(&v)
}
// SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil
func (o *GitJsonhal) SetSshPrivateKeyNil() {
	o.SshPrivateKey.Set(nil)
}

// UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
func (o *GitJsonhal) UnsetSshPrivateKey() {
	o.SshPrivateKey.Unset()
}

// GetSshPrivateKeyInput returns the SshPrivateKeyInput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetSshPrivateKeyInput() string {
	if o == nil || isNil(o.SshPrivateKeyInput.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKeyInput.Get()
}

// GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetSshPrivateKeyInputOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SshPrivateKeyInput.Get(), o.SshPrivateKeyInput.IsSet()
}

// HasSshPrivateKeyInput returns a boolean if a field has been set.
func (o *GitJsonhal) HasSshPrivateKeyInput() bool {
	if o != nil && o.SshPrivateKeyInput.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyInput gets a reference to the given NullableString and assigns it to the SshPrivateKeyInput field.
func (o *GitJsonhal) SetSshPrivateKeyInput(v string) {
	o.SshPrivateKeyInput.Set(&v)
}
// SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil
func (o *GitJsonhal) SetSshPrivateKeyInputNil() {
	o.SshPrivateKeyInput.Set(nil)
}

// UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
func (o *GitJsonhal) UnsetSshPrivateKeyInput() {
	o.SshPrivateKeyInput.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetProject() string {
	if o == nil || isNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetProjectOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *GitJsonhal) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *GitJsonhal) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *GitJsonhal) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *GitJsonhal) UnsetProject() {
	o.Project.Unset()
}

// GetGitType returns the GitType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetGitType() string {
	if o == nil || isNil(o.GitType.Get()) {
		var ret string
		return ret
	}
	return *o.GitType.Get()
}

// GetGitTypeOk returns a tuple with the GitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetGitTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.GitType.Get(), o.GitType.IsSet()
}

// HasGitType returns a boolean if a field has been set.
func (o *GitJsonhal) HasGitType() bool {
	if o != nil && o.GitType.IsSet() {
		return true
	}

	return false
}

// SetGitType gets a reference to the given NullableString and assigns it to the GitType field.
func (o *GitJsonhal) SetGitType(v string) {
	o.GitType.Set(&v)
}
// SetGitTypeNil sets the value for GitType to be an explicit nil
func (o *GitJsonhal) SetGitTypeNil() {
	o.GitType.Set(nil)
}

// UnsetGitType ensures that no value is present for GitType, not even an explicit nil
func (o *GitJsonhal) UnsetGitType() {
	o.GitType.Unset()
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetCredentialType() string {
	if o == nil || isNil(o.CredentialType.Get()) {
		var ret string
		return ret
	}
	return *o.CredentialType.Get()
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetCredentialTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CredentialType.Get(), o.CredentialType.IsSet()
}

// HasCredentialType returns a boolean if a field has been set.
func (o *GitJsonhal) HasCredentialType() bool {
	if o != nil && o.CredentialType.IsSet() {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given NullableString and assigns it to the CredentialType field.
func (o *GitJsonhal) SetCredentialType(v string) {
	o.CredentialType.Set(&v)
}
// SetCredentialTypeNil sets the value for CredentialType to be an explicit nil
func (o *GitJsonhal) SetCredentialTypeNil() {
	o.CredentialType.Set(nil)
}

// UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil
func (o *GitJsonhal) UnsetCredentialType() {
	o.CredentialType.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GitJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *GitJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *GitJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *GitJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *GitJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *GitJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}
// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *GitJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *GitJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GitJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GitJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GitJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GitJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GitJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GitJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GitJsonhal) MarshalJSON() ([]byte, error) {
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
	if o.Repo.IsSet() {
		toSerialize["repo"] = o.Repo.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.SshPrivateKey.IsSet() {
		toSerialize["sshPrivateKey"] = o.SshPrivateKey.Get()
	}
	if o.SshPrivateKeyInput.IsSet() {
		toSerialize["sshPrivateKeyInput"] = o.SshPrivateKeyInput.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.GitType.IsSet() {
		toSerialize["gitType"] = o.GitType.Get()
	}
	if o.CredentialType.IsSet() {
		toSerialize["credentialType"] = o.CredentialType.Get()
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
	return json.Marshal(toSerialize)
}

type NullableGitJsonhal struct {
	value *GitJsonhal
	isSet bool
}

func (v NullableGitJsonhal) Get() *GitJsonhal {
	return v.value
}

func (v *NullableGitJsonhal) Set(val *GitJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableGitJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableGitJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitJsonhal(val *GitJsonhal) *NullableGitJsonhal {
	return &NullableGitJsonhal{value: val, isSet: true}
}

func (v NullableGitJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


