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

// EnvironmentComponentEnvironmentPost 
type EnvironmentComponentEnvironmentPost struct {
	Cpu NullableInt32 `json:"cpu,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Storage NullableInt32 `json:"storage,omitempty"`
	Node NullableInt32 `json:"node,omitempty"`
	ComponentVersion NullableString `json:"componentVersion,omitempty"`
}

// NewEnvironmentComponentEnvironmentPost instantiates a new EnvironmentComponentEnvironmentPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentComponentEnvironmentPost() *EnvironmentComponentEnvironmentPost {
	this := EnvironmentComponentEnvironmentPost{}
	return &this
}

// NewEnvironmentComponentEnvironmentPostWithDefaults instantiates a new EnvironmentComponentEnvironmentPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentComponentEnvironmentPostWithDefaults() *EnvironmentComponentEnvironmentPost {
	this := EnvironmentComponentEnvironmentPost{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentPost) GetCpu() int32 {
	if o == nil || isNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentPost) GetCpuOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentPost) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *EnvironmentComponentEnvironmentPost) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *EnvironmentComponentEnvironmentPost) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *EnvironmentComponentEnvironmentPost) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentPost) GetMemory() int32 {
	if o == nil || isNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentPost) GetMemoryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentPost) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *EnvironmentComponentEnvironmentPost) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *EnvironmentComponentEnvironmentPost) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *EnvironmentComponentEnvironmentPost) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentPost) GetStorage() int32 {
	if o == nil || isNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentPost) GetStorageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentPost) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *EnvironmentComponentEnvironmentPost) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *EnvironmentComponentEnvironmentPost) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *EnvironmentComponentEnvironmentPost) UnsetStorage() {
	o.Storage.Unset()
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentPost) GetNode() int32 {
	if o == nil || isNil(o.Node.Get()) {
		var ret int32
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentPost) GetNodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentPost) HasNode() bool {
	if o != nil && o.Node.IsSet() {
		return true
	}

	return false
}

// SetNode gets a reference to the given NullableInt32 and assigns it to the Node field.
func (o *EnvironmentComponentEnvironmentPost) SetNode(v int32) {
	o.Node.Set(&v)
}
// SetNodeNil sets the value for Node to be an explicit nil
func (o *EnvironmentComponentEnvironmentPost) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil
func (o *EnvironmentComponentEnvironmentPost) UnsetNode() {
	o.Node.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentPost) GetComponentVersion() string {
	if o == nil || isNil(o.ComponentVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentVersion.Get()
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentPost) GetComponentVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentVersion.Get(), o.ComponentVersion.IsSet()
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentPost) HasComponentVersion() bool {
	if o != nil && o.ComponentVersion.IsSet() {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given NullableString and assigns it to the ComponentVersion field.
func (o *EnvironmentComponentEnvironmentPost) SetComponentVersion(v string) {
	o.ComponentVersion.Set(&v)
}
// SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil
func (o *EnvironmentComponentEnvironmentPost) SetComponentVersionNil() {
	o.ComponentVersion.Set(nil)
}

// UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
func (o *EnvironmentComponentEnvironmentPost) UnsetComponentVersion() {
	o.ComponentVersion.Unset()
}

func (o EnvironmentComponentEnvironmentPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu.IsSet() {
		toSerialize["cpu"] = o.Cpu.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Storage.IsSet() {
		toSerialize["storage"] = o.Storage.Get()
	}
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}
	if o.ComponentVersion.IsSet() {
		toSerialize["componentVersion"] = o.ComponentVersion.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentComponentEnvironmentPost struct {
	value *EnvironmentComponentEnvironmentPost
	isSet bool
}

func (v NullableEnvironmentComponentEnvironmentPost) Get() *EnvironmentComponentEnvironmentPost {
	return v.value
}

func (v *NullableEnvironmentComponentEnvironmentPost) Set(val *EnvironmentComponentEnvironmentPost) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentComponentEnvironmentPost) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentComponentEnvironmentPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentComponentEnvironmentPost(val *EnvironmentComponentEnvironmentPost) *NullableEnvironmentComponentEnvironmentPost {
	return &NullableEnvironmentComponentEnvironmentPost{value: val, isSet: true}
}

func (v NullableEnvironmentComponentEnvironmentPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentComponentEnvironmentPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


