# ProductJsonhalProductInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Prices** | Pointer to [**[]PriceJsonhalProductInput**](PriceJsonhalProductInput.md) |  | [optional] 
**DefaultPrice** | Pointer to [**NullableProductJsonhalProductInputDefaultPrice**](ProductJsonhalProductInputDefaultPrice.md) |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProductJsonhalProductInput

`func NewProductJsonhalProductInput() *ProductJsonhalProductInput`

NewProductJsonhalProductInput instantiates a new ProductJsonhalProductInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductJsonhalProductInputWithDefaults

`func NewProductJsonhalProductInputWithDefaults() *ProductJsonhalProductInput`

NewProductJsonhalProductInputWithDefaults instantiates a new ProductJsonhalProductInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProductJsonhalProductInput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProductJsonhalProductInput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProductJsonhalProductInput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProductJsonhalProductInput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProductJsonhalProductInput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductJsonhalProductInput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductJsonhalProductInput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductJsonhalProductInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProductJsonhalProductInput) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProductJsonhalProductInput) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProductJsonhalProductInput) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProductJsonhalProductInput) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProductJsonhalProductInput) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProductJsonhalProductInput) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetPaymentId

`func (o *ProductJsonhalProductInput) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ProductJsonhalProductInput) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ProductJsonhalProductInput) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ProductJsonhalProductInput) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *ProductJsonhalProductInput) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *ProductJsonhalProductInput) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetName

`func (o *ProductJsonhalProductInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductJsonhalProductInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductJsonhalProductInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductJsonhalProductInput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductJsonhalProductInput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductJsonhalProductInput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ProductJsonhalProductInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductJsonhalProductInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductJsonhalProductInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductJsonhalProductInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductJsonhalProductInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductJsonhalProductInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *ProductJsonhalProductInput) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductJsonhalProductInput) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductJsonhalProductInput) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductJsonhalProductInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ProductJsonhalProductInput) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ProductJsonhalProductInput) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsActive

`func (o *ProductJsonhalProductInput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductJsonhalProductInput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductJsonhalProductInput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProductJsonhalProductInput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProductJsonhalProductInput) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProductJsonhalProductInput) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetPrices

`func (o *ProductJsonhalProductInput) GetPrices() []PriceJsonhalProductInput`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductJsonhalProductInput) GetPricesOk() (*[]PriceJsonhalProductInput, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductJsonhalProductInput) SetPrices(v []PriceJsonhalProductInput)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ProductJsonhalProductInput) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetDefaultPrice

`func (o *ProductJsonhalProductInput) GetDefaultPrice() ProductJsonhalProductInputDefaultPrice`

GetDefaultPrice returns the DefaultPrice field if non-nil, zero value otherwise.

### GetDefaultPriceOk

`func (o *ProductJsonhalProductInput) GetDefaultPriceOk() (*ProductJsonhalProductInputDefaultPrice, bool)`

GetDefaultPriceOk returns a tuple with the DefaultPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrice

`func (o *ProductJsonhalProductInput) SetDefaultPrice(v ProductJsonhalProductInputDefaultPrice)`

SetDefaultPrice sets DefaultPrice field to given value.

### HasDefaultPrice

`func (o *ProductJsonhalProductInput) HasDefaultPrice() bool`

HasDefaultPrice returns a boolean if a field has been set.

### SetDefaultPriceNil

`func (o *ProductJsonhalProductInput) SetDefaultPriceNil(b bool)`

 SetDefaultPriceNil sets the value for DefaultPrice to be an explicit nil

### UnsetDefaultPrice
`func (o *ProductJsonhalProductInput) UnsetDefaultPrice()`

UnsetDefaultPrice ensures that no value is present for DefaultPrice, not even an explicit nil
### GetRegion

`func (o *ProductJsonhalProductInput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProductJsonhalProductInput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProductJsonhalProductInput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProductJsonhalProductInput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ProductJsonhalProductInput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ProductJsonhalProductInput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


