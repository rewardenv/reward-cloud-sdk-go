# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**KubecostUrl** | Pointer to **NullableString** |  | [optional] 
**KubecostUser** | Pointer to **NullableString** |  | [optional] 
**KubecostPass** | Pointer to **NullableString** |  | [optional] 
**ChangeKubecostPass** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | Pointer to **NullableBool** |  | [optional] 
**Environment** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Products** | Pointer to **[]string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Region) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Region) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Region) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Region) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Region) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Region) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Region) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Region) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Region) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Region) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Region) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Region) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Region) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKubecostUrl

`func (o *Region) GetKubecostUrl() string`

GetKubecostUrl returns the KubecostUrl field if non-nil, zero value otherwise.

### GetKubecostUrlOk

`func (o *Region) GetKubecostUrlOk() (*string, bool)`

GetKubecostUrlOk returns a tuple with the KubecostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubecostUrl

`func (o *Region) SetKubecostUrl(v string)`

SetKubecostUrl sets KubecostUrl field to given value.

### HasKubecostUrl

`func (o *Region) HasKubecostUrl() bool`

HasKubecostUrl returns a boolean if a field has been set.

### SetKubecostUrlNil

`func (o *Region) SetKubecostUrlNil(b bool)`

 SetKubecostUrlNil sets the value for KubecostUrl to be an explicit nil

### UnsetKubecostUrl
`func (o *Region) UnsetKubecostUrl()`

UnsetKubecostUrl ensures that no value is present for KubecostUrl, not even an explicit nil
### GetKubecostUser

`func (o *Region) GetKubecostUser() string`

GetKubecostUser returns the KubecostUser field if non-nil, zero value otherwise.

### GetKubecostUserOk

`func (o *Region) GetKubecostUserOk() (*string, bool)`

GetKubecostUserOk returns a tuple with the KubecostUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubecostUser

`func (o *Region) SetKubecostUser(v string)`

SetKubecostUser sets KubecostUser field to given value.

### HasKubecostUser

`func (o *Region) HasKubecostUser() bool`

HasKubecostUser returns a boolean if a field has been set.

### SetKubecostUserNil

`func (o *Region) SetKubecostUserNil(b bool)`

 SetKubecostUserNil sets the value for KubecostUser to be an explicit nil

### UnsetKubecostUser
`func (o *Region) UnsetKubecostUser()`

UnsetKubecostUser ensures that no value is present for KubecostUser, not even an explicit nil
### GetKubecostPass

`func (o *Region) GetKubecostPass() string`

GetKubecostPass returns the KubecostPass field if non-nil, zero value otherwise.

### GetKubecostPassOk

`func (o *Region) GetKubecostPassOk() (*string, bool)`

GetKubecostPassOk returns a tuple with the KubecostPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubecostPass

`func (o *Region) SetKubecostPass(v string)`

SetKubecostPass sets KubecostPass field to given value.

### HasKubecostPass

`func (o *Region) HasKubecostPass() bool`

HasKubecostPass returns a boolean if a field has been set.

### SetKubecostPassNil

`func (o *Region) SetKubecostPassNil(b bool)`

 SetKubecostPassNil sets the value for KubecostPass to be an explicit nil

### UnsetKubecostPass
`func (o *Region) UnsetKubecostPass()`

UnsetKubecostPass ensures that no value is present for KubecostPass, not even an explicit nil
### GetChangeKubecostPass

`func (o *Region) GetChangeKubecostPass() string`

GetChangeKubecostPass returns the ChangeKubecostPass field if non-nil, zero value otherwise.

### GetChangeKubecostPassOk

`func (o *Region) GetChangeKubecostPassOk() (*string, bool)`

GetChangeKubecostPassOk returns a tuple with the ChangeKubecostPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKubecostPass

`func (o *Region) SetChangeKubecostPass(v string)`

SetChangeKubecostPass sets ChangeKubecostPass field to given value.

### HasChangeKubecostPass

`func (o *Region) HasChangeKubecostPass() bool`

HasChangeKubecostPass returns a boolean if a field has been set.

### SetChangeKubecostPassNil

`func (o *Region) SetChangeKubecostPassNil(b bool)`

 SetChangeKubecostPassNil sets the value for ChangeKubecostPass to be an explicit nil

### UnsetChangeKubecostPass
`func (o *Region) UnsetChangeKubecostPass()`

UnsetChangeKubecostPass ensures that no value is present for ChangeKubecostPass, not even an explicit nil
### GetIsDefault

`func (o *Region) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Region) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Region) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Region) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *Region) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *Region) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetEnvironment

`func (o *Region) GetEnvironment() []string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Region) GetEnvironmentOk() (*[]string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Region) SetEnvironment(v []string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Region) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProvider

`func (o *Region) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Region) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Region) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Region) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *Region) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *Region) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetProducts

`func (o *Region) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Region) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Region) SetProducts(v []string)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *Region) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Region) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Region) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Region) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Region) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Region) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Region) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *Region) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Region) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Region) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Region) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *Region) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *Region) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetCreatedAt

`func (o *Region) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Region) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Region) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Region) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Region) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Region) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Region) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Region) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


