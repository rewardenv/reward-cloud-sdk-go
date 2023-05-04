# OrganisationEnvVarJsonhal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsEncrypted** | Pointer to **NullableBool** |  | [optional] 
**Organisation** | Pointer to **NullableString** |  | [optional] 
**EnvVarType** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RawValue** | Pointer to **NullableString** |  | [optional] [readonly] 

## Methods

### NewOrganisationEnvVarJsonhal

`func NewOrganisationEnvVarJsonhal() *OrganisationEnvVarJsonhal`

NewOrganisationEnvVarJsonhal instantiates a new OrganisationEnvVarJsonhal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationEnvVarJsonhalWithDefaults

`func NewOrganisationEnvVarJsonhalWithDefaults() *OrganisationEnvVarJsonhal`

NewOrganisationEnvVarJsonhalWithDefaults instantiates a new OrganisationEnvVarJsonhal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *OrganisationEnvVarJsonhal) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrganisationEnvVarJsonhal) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrganisationEnvVarJsonhal) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrganisationEnvVarJsonhal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *OrganisationEnvVarJsonhal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganisationEnvVarJsonhal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganisationEnvVarJsonhal) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrganisationEnvVarJsonhal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *OrganisationEnvVarJsonhal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OrganisationEnvVarJsonhal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OrganisationEnvVarJsonhal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *OrganisationEnvVarJsonhal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *OrganisationEnvVarJsonhal) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *OrganisationEnvVarJsonhal) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetKey

`func (o *OrganisationEnvVarJsonhal) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OrganisationEnvVarJsonhal) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OrganisationEnvVarJsonhal) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OrganisationEnvVarJsonhal) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *OrganisationEnvVarJsonhal) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *OrganisationEnvVarJsonhal) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *OrganisationEnvVarJsonhal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OrganisationEnvVarJsonhal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OrganisationEnvVarJsonhal) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OrganisationEnvVarJsonhal) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *OrganisationEnvVarJsonhal) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *OrganisationEnvVarJsonhal) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEncrypted

`func (o *OrganisationEnvVarJsonhal) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *OrganisationEnvVarJsonhal) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *OrganisationEnvVarJsonhal) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *OrganisationEnvVarJsonhal) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### SetIsEncryptedNil

`func (o *OrganisationEnvVarJsonhal) SetIsEncryptedNil(b bool)`

 SetIsEncryptedNil sets the value for IsEncrypted to be an explicit nil

### UnsetIsEncrypted
`func (o *OrganisationEnvVarJsonhal) UnsetIsEncrypted()`

UnsetIsEncrypted ensures that no value is present for IsEncrypted, not even an explicit nil
### GetOrganisation

`func (o *OrganisationEnvVarJsonhal) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *OrganisationEnvVarJsonhal) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *OrganisationEnvVarJsonhal) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.

### HasOrganisation

`func (o *OrganisationEnvVarJsonhal) HasOrganisation() bool`

HasOrganisation returns a boolean if a field has been set.

### SetOrganisationNil

`func (o *OrganisationEnvVarJsonhal) SetOrganisationNil(b bool)`

 SetOrganisationNil sets the value for Organisation to be an explicit nil

### UnsetOrganisation
`func (o *OrganisationEnvVarJsonhal) UnsetOrganisation()`

UnsetOrganisation ensures that no value is present for Organisation, not even an explicit nil
### GetEnvVarType

`func (o *OrganisationEnvVarJsonhal) GetEnvVarType() string`

GetEnvVarType returns the EnvVarType field if non-nil, zero value otherwise.

### GetEnvVarTypeOk

`func (o *OrganisationEnvVarJsonhal) GetEnvVarTypeOk() (*string, bool)`

GetEnvVarTypeOk returns a tuple with the EnvVarType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarType

`func (o *OrganisationEnvVarJsonhal) SetEnvVarType(v string)`

SetEnvVarType sets EnvVarType field to given value.

### HasEnvVarType

`func (o *OrganisationEnvVarJsonhal) HasEnvVarType() bool`

HasEnvVarType returns a boolean if a field has been set.

### SetEnvVarTypeNil

`func (o *OrganisationEnvVarJsonhal) SetEnvVarTypeNil(b bool)`

 SetEnvVarTypeNil sets the value for EnvVarType to be an explicit nil

### UnsetEnvVarType
`func (o *OrganisationEnvVarJsonhal) UnsetEnvVarType()`

UnsetEnvVarType ensures that no value is present for EnvVarType, not even an explicit nil
### GetCreatedBy

`func (o *OrganisationEnvVarJsonhal) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrganisationEnvVarJsonhal) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrganisationEnvVarJsonhal) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrganisationEnvVarJsonhal) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *OrganisationEnvVarJsonhal) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *OrganisationEnvVarJsonhal) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *OrganisationEnvVarJsonhal) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OrganisationEnvVarJsonhal) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OrganisationEnvVarJsonhal) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *OrganisationEnvVarJsonhal) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *OrganisationEnvVarJsonhal) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *OrganisationEnvVarJsonhal) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *OrganisationEnvVarJsonhal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganisationEnvVarJsonhal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganisationEnvVarJsonhal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganisationEnvVarJsonhal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrganisationEnvVarJsonhal) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganisationEnvVarJsonhal) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganisationEnvVarJsonhal) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrganisationEnvVarJsonhal) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRawValue

`func (o *OrganisationEnvVarJsonhal) GetRawValue() string`

GetRawValue returns the RawValue field if non-nil, zero value otherwise.

### GetRawValueOk

`func (o *OrganisationEnvVarJsonhal) GetRawValueOk() (*string, bool)`

GetRawValueOk returns a tuple with the RawValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawValue

`func (o *OrganisationEnvVarJsonhal) SetRawValue(v string)`

SetRawValue sets RawValue field to given value.

### HasRawValue

`func (o *OrganisationEnvVarJsonhal) HasRawValue() bool`

HasRawValue returns a boolean if a field has been set.

### SetRawValueNil

`func (o *OrganisationEnvVarJsonhal) SetRawValueNil(b bool)`

 SetRawValueNil sets the value for RawValue to be an explicit nil

### UnsetRawValue
`func (o *OrganisationEnvVarJsonhal) UnsetRawValue()`

UnsetRawValue ensures that no value is present for RawValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


