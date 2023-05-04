# EnvironmentEnvVarJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**Environment** | Pointer to [**NullableEnvironmentComponentJsonhalEnvironment**](EnvironmentComponentJsonhalEnvironment.md) |  | [optional] 
**EnvVarType** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewEnvironmentEnvVarJsonhal

`func NewEnvironmentEnvVarJsonhal() *EnvironmentEnvVarJsonhal`

NewEnvironmentEnvVarJsonhal instantiates a new EnvironmentEnvVarJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvVarJsonhalWithDefaults

`func NewEnvironmentEnvVarJsonhalWithDefaults() *EnvironmentEnvVarJsonhal`

NewEnvironmentEnvVarJsonhalWithDefaults instantiates a new EnvironmentEnvVarJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EnvironmentEnvVarJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentEnvVarJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentEnvVarJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentEnvVarJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentEnvVarJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentEnvVarJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentEnvVarJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentEnvVarJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentEnvVarJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentEnvVarJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentEnvVarJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentEnvVarJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentEnvVarJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentEnvVarJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetKey

`func (o *EnvironmentEnvVarJsonhal) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentEnvVarJsonhal) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentEnvVarJsonhal) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EnvironmentEnvVarJsonhal) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EnvironmentEnvVarJsonhal) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EnvironmentEnvVarJsonhal) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *EnvironmentEnvVarJsonhal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentEnvVarJsonhal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentEnvVarJsonhal) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironmentEnvVarJsonhal) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EnvironmentEnvVarJsonhal) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EnvironmentEnvVarJsonhal) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEncrypted

`func (o *EnvironmentEnvVarJsonhal) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *EnvironmentEnvVarJsonhal) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *EnvironmentEnvVarJsonhal) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *EnvironmentEnvVarJsonhal) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *EnvironmentEnvVarJsonhal) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *EnvironmentEnvVarJsonhal) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetEnvironment

`func (o *EnvironmentEnvVarJsonhal) GetEnvironment() EnvironmentComponentJsonhalEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentEnvVarJsonhal) GetEnvironmentOk() (*EnvironmentComponentJsonhalEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentEnvVarJsonhal) SetEnvironment(v EnvironmentComponentJsonhalEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentEnvVarJsonhal) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvironmentEnvVarJsonhal) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvironmentEnvVarJsonhal) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEnvVarType

`func (o *EnvironmentEnvVarJsonhal) GetEnvVarType() string`

GetEnvVarType returns the EnvVarType field if non-nil, zero value otherwise.

### GetEnvVarTypeOk

`func (o *EnvironmentEnvVarJsonhal) GetEnvVarTypeOk() (*string, bool)`

GetEnvVarTypeOk returns a tuple with the EnvVarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarType

`func (o *EnvironmentEnvVarJsonhal) SetEnvVarType(v string)`

SetEnvVarType sets EnvVarType field to given value.

### HasEnvVarType

`func (o *EnvironmentEnvVarJsonhal) HasEnvVarType() bool`

HasEnvVarType returns a boolean if a field has been set.

### SetEnvVarTypeNil

`func (o *EnvironmentEnvVarJsonhal) SetEnvVarTypeNil(b bool)`

 SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil

### UnsetEnvVarType
`func (o *EnvironmentEnvVarJsonhal) UnsetEnvVarType()`

UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentEnvVarJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentEnvVarJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentEnvVarJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentEnvVarJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentEnvVarJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentEnvVarJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentEnvVarJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentEnvVarJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentEnvVarJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentEnvVarJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentEnvVarJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentEnvVarJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentEnvVarJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentEnvVarJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentEnvVarJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentEnvVarJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentEnvVarJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentEnvVarJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentEnvVarJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentEnvVarJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRawValue

`func (o *EnvironmentEnvVarJsonhal) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *EnvironmentEnvVarJsonhal) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *EnvironmentEnvVarJsonhal) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *EnvironmentEnvVarJsonhal) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *EnvironmentEnvVarJsonhal) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *EnvironmentEnvVarJsonhal) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


