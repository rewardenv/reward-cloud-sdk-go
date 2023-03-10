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

// GitProjectPost 
type GitProjectPost struct {
	Repo NullableString `json:"repo,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	SshPrivateKey NullableString `json:"sshPrivateKey,omitempty"`
	SshPrivateKeyInput NullableString `json:"sshPrivateKeyInput,omitempty"`
	GitType NullableGitProjectPostGitType `json:"gitType,omitempty"`
	CredentialType NullableGitProjectPostCredentialType `json:"credentialType,omitempty"`
}

// NewGitProjectPost instantiates a new GitProjectPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitProjectPost() *GitProjectPost {
	this := GitProjectPost{}
	return &this
}

// NewGitProjectPostWithDefaults instantiates a new GitProjectPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitProjectPostWithDefaults() *GitProjectPost {
	this := GitProjectPost{}
	return &this
}

// GetRepo returns the Repo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetRepo() string {
	if o == nil || isNil(o.Repo.Get()) {
		var ret string
		return ret
	}
	return *o.Repo.Get()
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetRepoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Repo.Get(), o.Repo.IsSet()
}

// HasRepo returns a boolean if a field has been set.
func (o *GitProjectPost) HasRepo() bool {
	if o != nil && o.Repo.IsSet() {
		return true
	}

	return false
}

// SetRepo gets a reference to the given NullableString and assigns it to the Repo field.
func (o *GitProjectPost) SetRepo(v string) {
	o.Repo.Set(&v)
}
// SetRepoNil sets the value for Repo to be an explicit nil
func (o *GitProjectPost) SetRepoNil() {
	o.Repo.Set(nil)
}

// UnsetRepo ensures that no value is present for Repo, not even an explicit nil
func (o *GitProjectPost) UnsetRepo() {
	o.Repo.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetUsername() string {
	if o == nil || isNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *GitProjectPost) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *GitProjectPost) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *GitProjectPost) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *GitProjectPost) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetPassword() string {
	if o == nil || isNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *GitProjectPost) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *GitProjectPost) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *GitProjectPost) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *GitProjectPost) UnsetPassword() {
	o.Password.Unset()
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetSshPrivateKey() string {
	if o == nil || isNil(o.SshPrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey.Get()
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SshPrivateKey.Get(), o.SshPrivateKey.IsSet()
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *GitProjectPost) HasSshPrivateKey() bool {
	if o != nil && o.SshPrivateKey.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given NullableString and assigns it to the SshPrivateKey field.
func (o *GitProjectPost) SetSshPrivateKey(v string) {
	o.SshPrivateKey.Set(&v)
}
// SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil
func (o *GitProjectPost) SetSshPrivateKeyNil() {
	o.SshPrivateKey.Set(nil)
}

// UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
func (o *GitProjectPost) UnsetSshPrivateKey() {
	o.SshPrivateKey.Unset()
}

// GetSshPrivateKeyInput returns the SshPrivateKeyInput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetSshPrivateKeyInput() string {
	if o == nil || isNil(o.SshPrivateKeyInput.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKeyInput.Get()
}

// GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetSshPrivateKeyInputOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SshPrivateKeyInput.Get(), o.SshPrivateKeyInput.IsSet()
}

// HasSshPrivateKeyInput returns a boolean if a field has been set.
func (o *GitProjectPost) HasSshPrivateKeyInput() bool {
	if o != nil && o.SshPrivateKeyInput.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyInput gets a reference to the given NullableString and assigns it to the SshPrivateKeyInput field.
func (o *GitProjectPost) SetSshPrivateKeyInput(v string) {
	o.SshPrivateKeyInput.Set(&v)
}
// SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil
func (o *GitProjectPost) SetSshPrivateKeyInputNil() {
	o.SshPrivateKeyInput.Set(nil)
}

// UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
func (o *GitProjectPost) UnsetSshPrivateKeyInput() {
	o.SshPrivateKeyInput.Unset()
}

// GetGitType returns the GitType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetGitType() GitProjectPostGitType {
	if o == nil || isNil(o.GitType.Get()) {
		var ret GitProjectPostGitType
		return ret
	}
	return *o.GitType.Get()
}

// GetGitTypeOk returns a tuple with the GitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetGitTypeOk() (*GitProjectPostGitType, bool) {
	if o == nil {
    return nil, false
	}
	return o.GitType.Get(), o.GitType.IsSet()
}

// HasGitType returns a boolean if a field has been set.
func (o *GitProjectPost) HasGitType() bool {
	if o != nil && o.GitType.IsSet() {
		return true
	}

	return false
}

// SetGitType gets a reference to the given NullableGitProjectPostGitType and assigns it to the GitType field.
func (o *GitProjectPost) SetGitType(v GitProjectPostGitType) {
	o.GitType.Set(&v)
}
// SetGitTypeNil sets the value for GitType to be an explicit nil
func (o *GitProjectPost) SetGitTypeNil() {
	o.GitType.Set(nil)
}

// UnsetGitType ensures that no value is present for GitType, not even an explicit nil
func (o *GitProjectPost) UnsetGitType() {
	o.GitType.Unset()
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitProjectPost) GetCredentialType() GitProjectPostCredentialType {
	if o == nil || isNil(o.CredentialType.Get()) {
		var ret GitProjectPostCredentialType
		return ret
	}
	return *o.CredentialType.Get()
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitProjectPost) GetCredentialTypeOk() (*GitProjectPostCredentialType, bool) {
	if o == nil {
    return nil, false
	}
	return o.CredentialType.Get(), o.CredentialType.IsSet()
}

// HasCredentialType returns a boolean if a field has been set.
func (o *GitProjectPost) HasCredentialType() bool {
	if o != nil && o.CredentialType.IsSet() {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given NullableGitProjectPostCredentialType and assigns it to the CredentialType field.
func (o *GitProjectPost) SetCredentialType(v GitProjectPostCredentialType) {
	o.CredentialType.Set(&v)
}
// SetCredentialTypeNil sets the value for CredentialType to be an explicit nil
func (o *GitProjectPost) SetCredentialTypeNil() {
	o.CredentialType.Set(nil)
}

// UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil
func (o *GitProjectPost) UnsetCredentialType() {
	o.CredentialType.Unset()
}

func (o GitProjectPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableGitProjectPost struct {
	value *GitProjectPost
	isSet bool
}

func (v NullableGitProjectPost) Get() *GitProjectPost {
	return v.value
}

func (v *NullableGitProjectPost) Set(val *GitProjectPost) {
	v.value = val
	v.isSet = true
}

func (v NullableGitProjectPost) IsSet() bool {
	return v.isSet
}

func (v *NullableGitProjectPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitProjectPost(val *GitProjectPost) *NullableGitProjectPost {
	return &NullableGitProjectPost{value: val, isSet: true}
}

func (v NullableGitProjectPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitProjectPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


