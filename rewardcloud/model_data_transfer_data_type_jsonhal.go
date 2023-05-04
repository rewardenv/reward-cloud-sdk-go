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

// DataTransferDataTypeJsonhal Class DataTransferDataType
type DataTransferDataTypeJsonhal struct {
	Links        *AbstractEnvironmentJsonhalLinks `json:"_links,omitempty"`
	Id           *int32                           `json:"id,omitempty"`
	Name         NullableString                   `json:"name,omitempty"`
	ExportedData []string                         `json:"exportedData,omitempty"`
	ImportedData []string                         `json:"importedData,omitempty"`
	CreatedBy    NullableString                   `json:"createdBy,omitempty"`
	UpdatedBy    NullableString                   `json:"updatedBy,omitempty"`
	CreatedAt    *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time                       `json:"updatedAt,omitempty"`
}

// NewDataTransferDataTypeJsonhal instantiates a new DataTransferDataTypeJsonhal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataTransferDataTypeJsonhal() *DataTransferDataTypeJsonhal {
	this := DataTransferDataTypeJsonhal{}
	return &this
}

// NewDataTransferDataTypeJsonhalWithDefaults instantiates a new DataTransferDataTypeJsonhal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataTransferDataTypeJsonhalWithDefaults() *DataTransferDataTypeJsonhal {
	this := DataTransferDataTypeJsonhal{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DataTransferDataTypeJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks {
	if o == nil || isNil(o.Links) {
		var ret AbstractEnvironmentJsonhalLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferDataTypeJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AbstractEnvironmentJsonhalLinks and assigns it to the Links field.
func (o *DataTransferDataTypeJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataTransferDataTypeJsonhal) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferDataTypeJsonhal) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DataTransferDataTypeJsonhal) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferDataTypeJsonhal) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferDataTypeJsonhal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DataTransferDataTypeJsonhal) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *DataTransferDataTypeJsonhal) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DataTransferDataTypeJsonhal) UnsetName() {
	o.Name.Unset()
}

// GetExportedData returns the ExportedData field value if set, zero value otherwise.
func (o *DataTransferDataTypeJsonhal) GetExportedData() []string {
	if o == nil || isNil(o.ExportedData) {
		var ret []string
		return ret
	}
	return o.ExportedData
}

// GetExportedDataOk returns a tuple with the ExportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferDataTypeJsonhal) GetExportedDataOk() ([]string, bool) {
	if o == nil || isNil(o.ExportedData) {
		return nil, false
	}
	return o.ExportedData, true
}

// HasExportedData returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasExportedData() bool {
	if o != nil && !isNil(o.ExportedData) {
		return true
	}

	return false
}

// SetExportedData gets a reference to the given []string and assigns it to the ExportedData field.
func (o *DataTransferDataTypeJsonhal) SetExportedData(v []string) {
	o.ExportedData = v
}

// GetImportedData returns the ImportedData field value if set, zero value otherwise.
func (o *DataTransferDataTypeJsonhal) GetImportedData() []string {
	if o == nil || isNil(o.ImportedData) {
		var ret []string
		return ret
	}
	return o.ImportedData
}

// GetImportedDataOk returns a tuple with the ImportedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferDataTypeJsonhal) GetImportedDataOk() ([]string, bool) {
	if o == nil || isNil(o.ImportedData) {
		return nil, false
	}
	return o.ImportedData, true
}

// HasImportedData returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasImportedData() bool {
	if o != nil && !isNil(o.ImportedData) {
		return true
	}

	return false
}

// SetImportedData gets a reference to the given []string and assigns it to the ImportedData field.
func (o *DataTransferDataTypeJsonhal) SetImportedData(v []string) {
	o.ImportedData = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferDataTypeJsonhal) GetCreatedBy() string {
	if o == nil || isNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferDataTypeJsonhal) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *DataTransferDataTypeJsonhal) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *DataTransferDataTypeJsonhal) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *DataTransferDataTypeJsonhal) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataTransferDataTypeJsonhal) GetUpdatedBy() string {
	if o == nil || isNil(o.UpdatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedBy.Get()
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataTransferDataTypeJsonhal) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedBy.Get(), o.UpdatedBy.IsSet()
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy.IsSet() {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given NullableString and assigns it to the UpdatedBy field.
func (o *DataTransferDataTypeJsonhal) SetUpdatedBy(v string) {
	o.UpdatedBy.Set(&v)
}

// SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil
func (o *DataTransferDataTypeJsonhal) SetUpdatedByNil() {
	o.UpdatedBy.Set(nil)
}

// UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
func (o *DataTransferDataTypeJsonhal) UnsetUpdatedBy() {
	o.UpdatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DataTransferDataTypeJsonhal) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferDataTypeJsonhal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DataTransferDataTypeJsonhal) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DataTransferDataTypeJsonhal) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataTransferDataTypeJsonhal) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DataTransferDataTypeJsonhal) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DataTransferDataTypeJsonhal) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DataTransferDataTypeJsonhal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.ExportedData) {
		toSerialize["exportedData"] = o.ExportedData
	}
	if !isNil(o.ImportedData) {
		toSerialize["importedData"] = o.ImportedData
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

type NullableDataTransferDataTypeJsonhal struct {
	value *DataTransferDataTypeJsonhal
	isSet bool
}

func (v NullableDataTransferDataTypeJsonhal) Get() *DataTransferDataTypeJsonhal {
	return v.value
}

func (v *NullableDataTransferDataTypeJsonhal) Set(val *DataTransferDataTypeJsonhal) {
	v.value = val
	v.isSet = true
}

func (v NullableDataTransferDataTypeJsonhal) IsSet() bool {
	return v.isSet
}

func (v *NullableDataTransferDataTypeJsonhal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataTransferDataTypeJsonhal(val *DataTransferDataTypeJsonhal) *NullableDataTransferDataTypeJsonhal {
	return &NullableDataTransferDataTypeJsonhal{value: val, isSet: true}
}

func (v NullableDataTransferDataTypeJsonhal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataTransferDataTypeJsonhal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
