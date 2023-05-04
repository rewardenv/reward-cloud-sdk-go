# EnvironmentEnvVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**Environment** | Pointer to [**NullableEnvironmentComponentEnvironment**](EnvironmentComponentEnvironment.md) |  | [optional] 
**EnvVarType** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewEnvironmentEnvVar

`func NewEnvironmentEnvVar() *EnvironmentEnvVar`

NewEnvironmentEnvVar instantiates a new EnvironmentEnvVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentEnvVarWithDefaults

`func NewEnvironmentEnvVarWithDefaults() *EnvironmentEnvVar`

NewEnvironmentEnvVarWithDefaults instantiates a new EnvironmentEnvVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentEnvVar) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentEnvVar) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentEnvVar) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentEnvVar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentEnvVar) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentEnvVar) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentEnvVar) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentEnvVar) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EnvironmentEnvVar) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EnvironmentEnvVar) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetKey

`func (o *EnvironmentEnvVar) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EnvironmentEnvVar) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EnvironmentEnvVar) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EnvironmentEnvVar) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EnvironmentEnvVar) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EnvironmentEnvVar) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *EnvironmentEnvVar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EnvironmentEnvVar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EnvironmentEnvVar) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EnvironmentEnvVar) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EnvironmentEnvVar) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EnvironmentEnvVar) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEncrypted

`func (o *EnvironmentEnvVar) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *EnvironmentEnvVar) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *EnvironmentEnvVar) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *EnvironmentEnvVar) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *EnvironmentEnvVar) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *EnvironmentEnvVar) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetEnvironment

`func (o *EnvironmentEnvVar) GetEnvironment() EnvironmentComponentEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentEnvVar) GetEnvironmentOk() (*EnvironmentComponentEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentEnvVar) SetEnvironment(v EnvironmentComponentEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentEnvVar) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *EnvironmentEnvVar) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *EnvironmentEnvVar) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetEnvVarType

`func (o *EnvironmentEnvVar) GetEnvVarType() string`

GetEnvVarType returns the EnvVarType field if non-nil, zero value otherwise.

### GetEnvVarTypeOk

`func (o *EnvironmentEnvVar) GetEnvVarTypeOk() (*string, bool)`

GetEnvVarTypeOk returns a tuple with the EnvVarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarType

`func (o *EnvironmentEnvVar) SetEnvVarType(v string)`

SetEnvVarType sets EnvVarType field to given value.

### HasEnvVarType

`func (o *EnvironmentEnvVar) HasEnvVarType() bool`

HasEnvVarType returns a boolean if a field has been set.

### SetEnvVarTypeNil

`func (o *EnvironmentEnvVar) SetEnvVarTypeNil(b bool)`

 SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil

### UnsetEnvVarType
`func (o *EnvironmentEnvVar) UnsetEnvVarType()`

UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
### GetCreatedBy

`func (o *EnvironmentEnvVar) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EnvironmentEnvVar) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EnvironmentEnvVar) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EnvironmentEnvVar) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *EnvironmentEnvVar) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *EnvironmentEnvVar) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *EnvironmentEnvVar) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *EnvironmentEnvVar) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *EnvironmentEnvVar) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *EnvironmentEnvVar) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *EnvironmentEnvVar) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *EnvironmentEnvVar) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *EnvironmentEnvVar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentEnvVar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentEnvVar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentEnvVar) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentEnvVar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentEnvVar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentEnvVar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentEnvVar) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRawValue

`func (o *EnvironmentEnvVar) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *EnvironmentEnvVar) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *EnvironmentEnvVar) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *EnvironmentEnvVar) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *EnvironmentEnvVar) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *EnvironmentEnvVar) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


