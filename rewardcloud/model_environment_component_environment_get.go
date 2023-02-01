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

// EnvironmentComponentEnvironmentGet 
type EnvironmentComponentEnvironmentGet struct {
	Cpu NullableInt32 `json:"cpu,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Storage NullableInt32 `json:"storage,omitempty"`
	Node NullableInt32 `json:"node,omitempty"`
	ComponentVersion NullableString `json:"componentVersion,omitempty"`
}

// NewEnvironmentComponentEnvironmentGet instantiates a new EnvironmentComponentEnvironmentGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentComponentEnvironmentGet() *EnvironmentComponentEnvironmentGet {
	this := EnvironmentComponentEnvironmentGet{}
	return &this
}

// NewEnvironmentComponentEnvironmentGetWithDefaults instantiates a new EnvironmentComponentEnvironmentGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentComponentEnvironmentGetWithDefaults() *EnvironmentComponentEnvironmentGet {
	this := EnvironmentComponentEnvironmentGet{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentGet) GetCpu() int32 {
	if o == nil || isNil(o.Cpu.Get()) {
		var ret int32
		return ret
	}
	return *o.Cpu.Get()
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentGet) GetCpuOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Cpu.Get(), o.Cpu.IsSet()
}

// HasCpu returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentGet) HasCpu() bool {
	if o != nil && o.Cpu.IsSet() {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NullableInt32 and assigns it to the Cpu field.
func (o *EnvironmentComponentEnvironmentGet) SetCpu(v int32) {
	o.Cpu.Set(&v)
}
// SetCpuNil sets the value for Cpu to be an explicit nil
func (o *EnvironmentComponentEnvironmentGet) SetCpuNil() {
	o.Cpu.Set(nil)
}

// UnsetCpu ensures that no value is present for Cpu, not even an explicit nil
func (o *EnvironmentComponentEnvironmentGet) UnsetCpu() {
	o.Cpu.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentGet) GetMemory() int32 {
	if o == nil || isNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentGet) GetMemoryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentGet) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *EnvironmentComponentEnvironmentGet) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *EnvironmentComponentEnvironmentGet) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *EnvironmentComponentEnvironmentGet) UnsetMemory() {
	o.Memory.Unset()
}

// GetStorage returns the Storage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentGet) GetStorage() int32 {
	if o == nil || isNil(o.Storage.Get()) {
		var ret int32
		return ret
	}
	return *o.Storage.Get()
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentGet) GetStorageOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Storage.Get(), o.Storage.IsSet()
}

// HasStorage returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentGet) HasStorage() bool {
	if o != nil && o.Storage.IsSet() {
		return true
	}

	return false
}

// SetStorage gets a reference to the given NullableInt32 and assigns it to the Storage field.
func (o *EnvironmentComponentEnvironmentGet) SetStorage(v int32) {
	o.Storage.Set(&v)
}
// SetStorageNil sets the value for Storage to be an explicit nil
func (o *EnvironmentComponentEnvironmentGet) SetStorageNil() {
	o.Storage.Set(nil)
}

// UnsetStorage ensures that no value is present for Storage, not even an explicit nil
func (o *EnvironmentComponentEnvironmentGet) UnsetStorage() {
	o.Storage.Unset()
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentGet) GetNode() int32 {
	if o == nil || isNil(o.Node.Get()) {
		var ret int32
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentGet) GetNodeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentGet) HasNode() bool {
	if o != nil && o.Node.IsSet() {
		return true
	}

	return false
}

// SetNode gets a reference to the given NullableInt32 and assigns it to the Node field.
func (o *EnvironmentComponentEnvironmentGet) SetNode(v int32) {
	o.Node.Set(&v)
}
// SetNodeNil sets the value for Node to be an explicit nil
func (o *EnvironmentComponentEnvironmentGet) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil
func (o *EnvironmentComponentEnvironmentGet) UnsetNode() {
	o.Node.Unset()
}

// GetComponentVersion returns the ComponentVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentComponentEnvironmentGet) GetComponentVersion() string {
	if o == nil || isNil(o.ComponentVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ComponentVersion.Get()
}

// GetComponentVersionOk returns a tuple with the ComponentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentComponentEnvironmentGet) GetComponentVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ComponentVersion.Get(), o.ComponentVersion.IsSet()
}

// HasComponentVersion returns a boolean if a field has been set.
func (o *EnvironmentComponentEnvironmentGet) HasComponentVersion() bool {
	if o != nil && o.ComponentVersion.IsSet() {
		return true
	}

	return false
}

// SetComponentVersion gets a reference to the given NullableString and assigns it to the ComponentVersion field.
func (o *EnvironmentComponentEnvironmentGet) SetComponentVersion(v string) {
	o.ComponentVersion.Set(&v)
}
// SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil
func (o *EnvironmentComponentEnvironmentGet) SetComponentVersionNil() {
	o.ComponentVersion.Set(nil)
}

// UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
func (o *EnvironmentComponentEnvironmentGet) UnsetComponentVersion() {
	o.ComponentVersion.Unset()
}

func (o EnvironmentComponentEnvironmentGet) MarshalJSON() ([]byte, error) {
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

type NullableEnvironmentComponentEnvironmentGet struct {
	value *EnvironmentComponentEnvironmentGet
	isSet bool
}

func (v NullableEnvironmentComponentEnvironmentGet) Get() *EnvironmentComponentEnvironmentGet {
	return v.value
}

func (v *NullableEnvironmentComponentEnvironmentGet) Set(val *EnvironmentComponentEnvironmentGet) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentComponentEnvironmentGet) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentComponentEnvironmentGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentComponentEnvironmentGet(val *EnvironmentComponentEnvironmentGet) *NullableEnvironmentComponentEnvironmentGet {
	return &NullableEnvironmentComponentEnvironmentGet{value: val, isSet: true}
}

func (v NullableEnvironmentComponentEnvironmentGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentComponentEnvironmentGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

