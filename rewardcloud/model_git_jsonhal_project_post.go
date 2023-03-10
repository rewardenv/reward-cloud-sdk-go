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

// GitJsonhalProjectPost 
type GitJsonhalProjectPost struct {
	Links *ComponentJsonhalLinks `json:"_links,omitempty"`
	Repo NullableString `json:"repo,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	SshPrivateKey NullableString `json:"sshPrivateKey,omitempty"`
	SshPrivateKeyInput NullableString `json:"sshPrivateKeyInput,omitempty"`
	GitType NullableGitJsonhalProjectPostGitType `json:"gitType,omitempty"`
	CredentialType NullableGitJsonhalProjectPostCredentialType `json:"credentialType,omitempty"`
}

// NewGitJsonhalProjectPost instantiates a new GitJsonhalProjectPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitJsonhalProjectPost() *GitJsonhalProjectPost {
	this := GitJsonhalProjectPost{}
	return &this
}

// NewGitJsonhalProjectPostWithDefaults instantiates a new GitJsonhalProjectPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitJsonhalProjectPostWithDefaults() *GitJsonhalProjectPost {
	this := GitJsonhalProjectPost{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GitJsonhalProjectPost) GetLinks() ComponentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret ComponentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitJsonhalProjectPost) GetLinksOk() (*ComponentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ComponentJsonhalLinks and assigns it to the Links field.
func (o *GitJsonhalProjectPost) SetLinks(v ComponentJsonhalLinks) {
	o.Links = &v
}

// GetRepo returns the Repo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetRepo() string {
	if o == nil || isNil(o.Repo.Get()) {
		var ret string
		return ret
	}
	return *o.Repo.Get()
}

// GetRepoOk returns a tuple with the Repo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetRepoOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Repo.Get(), o.Repo.IsSet()
}

// HasRepo returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasRepo() bool {
	if o != nil && o.Repo.IsSet() {
		return true
	}

	return false
}

// SetRepo gets a reference to the given NullableString and assigns it to the Repo field.
func (o *GitJsonhalProjectPost) SetRepo(v string) {
	o.Repo.Set(&v)
}
// SetRepoNil sets the value for Repo to be an explicit nil
func (o *GitJsonhalProjectPost) SetRepoNil() {
	o.Repo.Set(nil)
}

// UnsetRepo ensures that no value is present for Repo, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetRepo() {
	o.Repo.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetUsername() string {
	if o == nil || isNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetUsernameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *GitJsonhalProjectPost) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *GitJsonhalProjectPost) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetPassword() string {
	if o == nil || isNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetPasswordOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *GitJsonhalProjectPost) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *GitJsonhalProjectPost) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetPassword() {
	o.Password.Unset()
}

// GetSshPrivateKey returns the SshPrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetSshPrivateKey() string {
	if o == nil || isNil(o.SshPrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKey.Get()
}

// GetSshPrivateKeyOk returns a tuple with the SshPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetSshPrivateKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SshPrivateKey.Get(), o.SshPrivateKey.IsSet()
}

// HasSshPrivateKey returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasSshPrivateKey() bool {
	if o != nil && o.SshPrivateKey.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKey gets a reference to the given NullableString and assigns it to the SshPrivateKey field.
func (o *GitJsonhalProjectPost) SetSshPrivateKey(v string) {
	o.SshPrivateKey.Set(&v)
}
// SetSshPrivateKeyNil sets the value for SshPrivateKey to be an explicit nil
func (o *GitJsonhalProjectPost) SetSshPrivateKeyNil() {
	o.SshPrivateKey.Set(nil)
}

// UnsetSshPrivateKey ensures that no value is present for SshPrivateKey, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetSshPrivateKey() {
	o.SshPrivateKey.Unset()
}

// GetSshPrivateKeyInput returns the SshPrivateKeyInput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetSshPrivateKeyInput() string {
	if o == nil || isNil(o.SshPrivateKeyInput.Get()) {
		var ret string
		return ret
	}
	return *o.SshPrivateKeyInput.Get()
}

// GetSshPrivateKeyInputOk returns a tuple with the SshPrivateKeyInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetSshPrivateKeyInputOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.SshPrivateKeyInput.Get(), o.SshPrivateKeyInput.IsSet()
}

// HasSshPrivateKeyInput returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasSshPrivateKeyInput() bool {
	if o != nil && o.SshPrivateKeyInput.IsSet() {
		return true
	}

	return false
}

// SetSshPrivateKeyInput gets a reference to the given NullableString and assigns it to the SshPrivateKeyInput field.
func (o *GitJsonhalProjectPost) SetSshPrivateKeyInput(v string) {
	o.SshPrivateKeyInput.Set(&v)
}
// SetSshPrivateKeyInputNil sets the value for SshPrivateKeyInput to be an explicit nil
func (o *GitJsonhalProjectPost) SetSshPrivateKeyInputNil() {
	o.SshPrivateKeyInput.Set(nil)
}

// UnsetSshPrivateKeyInput ensures that no value is present for SshPrivateKeyInput, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetSshPrivateKeyInput() {
	o.SshPrivateKeyInput.Unset()
}

// GetGitType returns the GitType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetGitType() GitJsonhalProjectPostGitType {
	if o == nil || isNil(o.GitType.Get()) {
		var ret GitJsonhalProjectPostGitType
		return ret
	}
	return *o.GitType.Get()
}

// GetGitTypeOk returns a tuple with the GitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetGitTypeOk() (*GitJsonhalProjectPostGitType, bool) {
	if o == nil {
    return nil, false
	}
	return o.GitType.Get(), o.GitType.IsSet()
}

// HasGitType returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasGitType() bool {
	if o != nil && o.GitType.IsSet() {
		return true
	}

	return false
}

// SetGitType gets a reference to the given NullableGitJsonhalProjectPostGitType and assigns it to the GitType field.
func (o *GitJsonhalProjectPost) SetGitType(v GitJsonhalProjectPostGitType) {
	o.GitType.Set(&v)
}
// SetGitTypeNil sets the value for GitType to be an explicit nil
func (o *GitJsonhalProjectPost) SetGitTypeNil() {
	o.GitType.Set(nil)
}

// UnsetGitType ensures that no value is present for GitType, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetGitType() {
	o.GitType.Unset()
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitJsonhalProjectPost) GetCredentialType() GitJsonhalProjectPostCredentialType {
	if o == nil || isNil(o.CredentialType.Get()) {
		var ret GitJsonhalProjectPostCredentialType
		return ret
	}
	return *o.CredentialType.Get()
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitJsonhalProjectPost) GetCredentialTypeOk() (*GitJsonhalProjectPostCredentialType, bool) {
	if o == nil {
    return nil, false
	}
	return o.CredentialType.Get(), o.CredentialType.IsSet()
}

// HasCredentialType returns a boolean if a field has been set.
func (o *GitJsonhalProjectPost) HasCredentialType() bool {
	if o != nil && o.CredentialType.IsSet() {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given NullableGitJsonhalProjectPostCredentialType and assigns it to the CredentialType field.
func (o *GitJsonhalProjectPost) SetCredentialType(v GitJsonhalProjectPostCredentialType) {
	o.CredentialType.Set(&v)
}
// SetCredentialTypeNil sets the value for CredentialType to be an explicit nil
func (o *GitJsonhalProjectPost) SetCredentialTypeNil() {
	o.CredentialType.Set(nil)
}

// UnsetCredentialType ensures that no value is present for CredentialType, not even an explicit nil
func (o *GitJsonhalProjectPost) UnsetCredentialType() {
	o.CredentialType.Unset()
}

func (o GitJsonhalProjectPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
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
	return json.Marshal(toSerialize)
}

type NullableGitJsonhalProjectPost struct {
	value *GitJsonhalProjectPost
	isSet bool
}

func (v NullableGitJsonhalProjectPost) Get() *GitJsonhalProjectPost {
	return v.value
}

func (v *NullableGitJsonhalProjectPost) Set(val *GitJsonhalProjectPost) {
	v.value = val
	v.isSet = true
}

func (v NullableGitJsonhalProjectPost) IsSet() bool {
	return v.isSet
}

func (v *NullableGitJsonhalProjectPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitJsonhalProjectPost(val *GitJsonhalProjectPost) *NullableGitJsonhalProjectPost {
	return &NullableGitJsonhalProjectPost{value: val, isSet: true}
}

func (v NullableGitJsonhalProjectPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitJsonhalProjectPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


