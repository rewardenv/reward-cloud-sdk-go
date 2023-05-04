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

// checks if the GitProjectGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitProjectGet{}

// GitProjectGet 
type GitProjectGet struct {
	Id *int32 `json:"id,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	Repo NullableString `json:"repo,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	SshPrivateKey NullableString `json:"sshPrivateKey,omitempty"`
	SshPrivateKeyInput NullableString `json:"sshPrivateKeyInput,omitempty"`
	GitType NullableGitProjectGetGitType `json:"gitType,omitempty"`
	CredentialType NullableGitProjectGetCredentialType `json:"credentialType,omitempty"`
}

// NewGitProjectGet instantiates a new GitProjectGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitProjectGet() *GitProjectGet {
	this := GitProjectGet{}
	return &this
}

// NewGitProjectGetWithDefaults instantiates a new GitProjectGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitProjectGetWithDefaults() *GitProjectGet {
	this := GitProjectGet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GitProjectGet) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitProjectGet) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GitProjectGet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GitProjectGet) SetId(v int32) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *GitProjectGet) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *GitProjectGet) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *GitProjectGet) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *GitProjectGet) UnsetUuid() {
	o.Uuid.Unset()
}

// GetRepo returns the Repo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetRepo() string {
	if o == nil || IsNil(o.Repo.Get()) {
		var ret string
		return ret
	}
	return *o.Repo.Get()
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetRepoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repo.Get(), o.Repo.IsSet()
}

// HasRepo returns a boolean if a field has been set.
func (o *GitProjectGet) HasRepo() bool {
	if o != nil && o.Repo.IsSet() {
		return true
	}

	return false
}

// SetRepo gets a reference to the given NullableString and assigns it to the Repo field.
func (o *GitProjectGet) SetRepo(v string) {
	o.Repo.Set(&v)
}
// SetRepoNil sets the value for Repo to be an explicit nil
func (o *GitProjectGet) SetRepoNil() {
	o.Repo.Set(nil)
}

// UnsetRepo ensures that no value is present for Repo, not even an explicit nil
func (o *GitProjectGet) UnsetRepo() {
	o.Repo.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *GitProjectGet) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *GitProjectGet) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *GitProjectGet) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *GitProjectGet) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *GitProjectGet) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *GitProjectGet) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *GitProjectGet) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *GitProjectGet) UnsetPassword() {
	o.Password.Unset()
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetSshPrivateKey() string {
	if o == nil || IsNil(o.SshPrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey.Get()
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPrivateKey.Get(), o.SshPrivateKey.IsSet()
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *GitProjectGet) HasSshPrivateKey() bool {
	if o != nil && o.SshPrivateKey.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given NullableString and assigns it to the SshPrivateKey field.
func (o *GitProjectGet) SetSshPrivateKey(v string) {
	o.SshPrivateKey.Set(&v)
}
// SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil
func (o *GitProjectGet) SetSshPrivateKeyNil() {
	o.SshPrivateKey.Set(nil)
}

// UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
func (o *GitProjectGet) UnsetSshPrivateKey() {
	o.SshPrivateKey.Unset()
}

// GetSshPrivateKeyInput returns the SshPrivateKeyInput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetSshPrivateKeyInput() string {
	if o == nil || IsNil(o.SshPrivateKeyInput.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKeyInput.Get()
}

// GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetSshPrivateKeyInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPrivateKeyInput.Get(), o.SshPrivateKeyInput.IsSet()
}

// HasSshPrivateKeyInput returns a boolean if a field has been set.
func (o *GitProjectGet) HasSshPrivateKeyInput() bool {
	if o != nil && o.SshPrivateKeyInput.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyInput gets a reference to the given NullableString and assigns it to the SshPrivateKeyInput field.
func (o *GitProjectGet) SetSshPrivateKeyInput(v string) {
	o.SshPrivateKeyInput.Set(&v)
}
// SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil
func (o *GitProjectGet) SetSshPrivateKeyInputNil() {
	o.SshPrivateKeyInput.Set(nil)
}

// UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
func (o *GitProjectGet) UnsetSshPrivateKeyInput() {
	o.SshPrivateKeyInput.Unset()
}

// GetGitType returns the GitType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetGitType() GitProjectGetGitType {
	if o == nil || IsNil(o.GitType.Get()) {
		var ret GitProjectGetGitType
		return ret
	}
	return *o.GitType.Get()
}

// GetGitTypeOk returns a tuple with the GitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetGitTypeOk() (*GitProjectGetGitType, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitType.Get(), o.GitType.IsSet()
}

// HasGitType returns a boolean if a field has been set.
func (o *GitProjectGet) HasGitType() bool {
	if o != nil && o.GitType.IsSet() {
		return true
	}

	return false
}

// SetGitType gets a reference to the given NullableGitProjectGetGitType and assigns it to the GitType field.
func (o *GitProjectGet) SetGitType(v GitProjectGetGitType) {
	o.GitType.Set(&v)
}
// SetGitTypeNil sets the value for GitType to be an explicit nil
func (o *GitProjectGet) SetGitTypeNil() {
	o.GitType.Set(nil)
}

// UnsetGitType ensures that no value is present for GitType, not even an explicit nil
func (o *GitProjectGet) UnsetGitType() {
	o.GitType.Unset()
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectGet) GetCredentialType() GitProjectGetCredentialType {
	if o == nil || IsNil(o.CredentialType.Get()) {
		var ret GitProjectGetCredentialType
		return ret
	}
	return *o.CredentialType.Get()
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectGet) GetCredentialTypeOk() (*GitProjectGetCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return o.CredentialType.Get(), o.CredentialType.IsSet()
}

// HasCredentialType returns a boolean if a field has been set.
func (o *GitProjectGet) HasCredentialType() bool {
	if o != nil && o.CredentialType.IsSet() {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given NullableGitProjectGetCredentialType and assigns it to the CredentialType field.
func (o *GitProjectGet) SetCredentialType(v GitProjectGetCredentialType) {
	o.CredentialType.Set(&v)
}
// SetCredentialTypeNil sets the value for CredentialType to be an explicit nil
func (o *GitProjectGet) SetCredentialTypeNil() {
	o.CredentialType.Set(nil)
}

// UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil
func (o *GitProjectGet) UnsetCredentialType() {
	o.CredentialType.Unset()
}

func (o GitProjectGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitProjectGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
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
	if o.GitType.IsSet() {
		toSerialize["gitType"] = o.GitType.Get()
	}
	if o.CredentialType.IsSet() {
		toSerialize["credentialType"] = o.CredentialType.Get()
	}
	return toSerialize, nil
}

type NullableGitProjectGet struct {
	value *GitProjectGet
	isSet bool
}

func (v NullableGitProjectGet) Get() *GitProjectGet {
	return v.value
}

func (v *NullableGitProjectGet) Set(val *GitProjectGet) {
	v.value = val
	v.isSet = true
}

func (v NullableGitProjectGet) IsSet() bool {
	return v.isSet
}

func (v *NullableGitProjectGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitProjectGet(val *GitProjectGet) *NullableGitProjectGet {
	return &NullableGitProjectGet{value: val, isSet: true}
}

func (v NullableGitProjectGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitProjectGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


