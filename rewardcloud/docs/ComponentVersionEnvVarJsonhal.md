# ComponentVersionEnvVarJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**ComponentVersion** | Pointer to **NullableString** |  | [optional] 
**ComponentVersionEnvVarExample** | Pointer to **[]string** |  | [optional] 
**EnvVarType** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewComponentVersionEnvVarJsonhal

`func NewComponentVersionEnvVarJsonhal() *ComponentVersionEnvVarJsonhal`

NewComponentVersionEnvVarJsonhal instantiates a new ComponentVersionEnvVarJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentVersionEnvVarJsonhalWithDefaults

`func NewComponentVersionEnvVarJsonhalWithDefaults() *ComponentVersionEnvVarJsonhal`

NewComponentVersionEnvVarJsonhalWithDefaults instantiates a new ComponentVersionEnvVarJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ComponentVersionEnvVarJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ComponentVersionEnvVarJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ComponentVersionEnvVarJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ComponentVersionEnvVarJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ComponentVersionEnvVarJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentVersionEnvVarJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentVersionEnvVarJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentVersionEnvVarJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ComponentVersionEnvVarJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ComponentVersionEnvVarJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ComponentVersionEnvVarJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ComponentVersionEnvVarJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ComponentVersionEnvVarJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ComponentVersionEnvVarJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetKey

`func (o *ComponentVersionEnvVarJsonhal) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComponentVersionEnvVarJsonhal) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComponentVersionEnvVarJsonhal) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ComponentVersionEnvVarJsonhal) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ComponentVersionEnvVarJsonhal) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ComponentVersionEnvVarJsonhal) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ComponentVersionEnvVarJsonhal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComponentVersionEnvVarJsonhal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComponentVersionEnvVarJsonhal) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ComponentVersionEnvVarJsonhal) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ComponentVersionEnvVarJsonhal) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ComponentVersionEnvVarJsonhal) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEncrypted

`func (o *ComponentVersionEnvVarJsonhal) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *ComponentVersionEnvVarJsonhal) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *ComponentVersionEnvVarJsonhal) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *ComponentVersionEnvVarJsonhal) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *ComponentVersionEnvVarJsonhal) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *ComponentVersionEnvVarJsonhal) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetComponentVersion

`func (o *ComponentVersionEnvVarJsonhal) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *ComponentVersionEnvVarJsonhal) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *ComponentVersionEnvVarJsonhal) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.

### HasComponentVersion

`func (o *ComponentVersionEnvVarJsonhal) HasComponentVersion() bool`

HasComponentVersion returns a boolean if a field has been set.

### SetComponentVersionNil

`func (o *ComponentVersionEnvVarJsonhal) SetComponentVersionNil(b bool)`

 SetComponentVersionNil sets the value for ComponentVersion to be an explicit nil

### UnsetComponentVersion
`func (o *ComponentVersionEnvVarJsonhal) UnsetComponentVersion()`

UnsetComponentVersion ensures that no value is present for ComponentVersion, not even an explicit nil
### GetComponentVersionEnvVarExample

`func (o *ComponentVersionEnvVarJsonhal) GetComponentVersionEnvVarExample() []string`

GetComponentVersionEnvVarExample returns the ComponentVersionEnvVarExample field if non-nil, zero value otherwise.

### GetComponentVersionEnvVarExampleOk

`func (o *ComponentVersionEnvVarJsonhal) GetComponentVersionEnvVarExampleOk() (*[]string, bool)`

GetComponentVersionEnvVarExampleOk returns a tuple with the ComponentVersionEnvVarExample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersionEnvVarExample

`func (o *ComponentVersionEnvVarJsonhal) SetComponentVersionEnvVarExample(v []string)`

SetComponentVersionEnvVarExample sets ComponentVersionEnvVarExample field to given value.

### HasComponentVersionEnvVarExample

`func (o *ComponentVersionEnvVarJsonhal) HasComponentVersionEnvVarExample() bool`

HasComponentVersionEnvVarExample returns a boolean if a field has been set.

### GetEnvVarType

`func (o *ComponentVersionEnvVarJsonhal) GetEnvVarType() string`

GetEnvVarType returns the EnvVarType field if non-nil, zero value otherwise.

### GetEnvVarTypeOk

`func (o *ComponentVersionEnvVarJsonhal) GetEnvVarTypeOk() (*string, bool)`

GetEnvVarTypeOk returns a tuple with the EnvVarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarType

`func (o *ComponentVersionEnvVarJsonhal) SetEnvVarType(v string)`

SetEnvVarType sets EnvVarType field to given value.

### HasEnvVarType

`func (o *ComponentVersionEnvVarJsonhal) HasEnvVarType() bool`

HasEnvVarType returns a boolean if a field has been set.

### SetEnvVarTypeNil

`func (o *ComponentVersionEnvVarJsonhal) SetEnvVarTypeNil(b bool)`

 SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil

### UnsetEnvVarType
`func (o *ComponentVersionEnvVarJsonhal) UnsetEnvVarType()`

UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
### GetCreatedBy

`func (o *ComponentVersionEnvVarJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ComponentVersionEnvVarJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ComponentVersionEnvVarJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ComponentVersionEnvVarJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ComponentVersionEnvVarJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ComponentVersionEnvVarJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ComponentVersionEnvVarJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ComponentVersionEnvVarJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ComponentVersionEnvVarJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ComponentVersionEnvVarJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ComponentVersionEnvVarJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ComponentVersionEnvVarJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *ComponentVersionEnvVarJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComponentVersionEnvVarJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComponentVersionEnvVarJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComponentVersionEnvVarJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComponentVersionEnvVarJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComponentVersionEnvVarJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComponentVersionEnvVarJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComponentVersionEnvVarJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


