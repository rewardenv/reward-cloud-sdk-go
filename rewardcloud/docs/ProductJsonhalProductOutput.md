# ProductJsonhalProductOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **[]string** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Prices** | Pointer to [**[]PriceJsonhalProductOutput**](PriceJsonhalProductOutput.md) |  | [optional] 
**DefaultPrice** | Pointer to [**NullableProductJsonhalProductOutputDefaultPrice**](ProductJsonhalProductOutputDefaultPrice.md) |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProductJsonhalProductOutput

`func NewProductJsonhalProductOutput() *ProductJsonhalProductOutput`

NewProductJsonhalProductOutput instantiates a new ProductJsonhalProductOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductJsonhalProductOutputWithDefaults

`func NewProductJsonhalProductOutputWithDefaults() *ProductJsonhalProductOutput`

NewProductJsonhalProductOutputWithDefaults instantiates a new ProductJsonhalProductOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProductJsonhalProductOutput) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProductJsonhalProductOutput) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProductJsonhalProductOutput) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProductJsonhalProductOutput) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPaymentId

`func (o *ProductJsonhalProductOutput) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ProductJsonhalProductOutput) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ProductJsonhalProductOutput) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ProductJsonhalProductOutput) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *ProductJsonhalProductOutput) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *ProductJsonhalProductOutput) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetName

`func (o *ProductJsonhalProductOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductJsonhalProductOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductJsonhalProductOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductJsonhalProductOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ProductJsonhalProductOutput) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProductJsonhalProductOutput) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ProductJsonhalProductOutput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductJsonhalProductOutput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductJsonhalProductOutput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductJsonhalProductOutput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductJsonhalProductOutput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductJsonhalProductOutput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *ProductJsonhalProductOutput) GetMetadata() []string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductJsonhalProductOutput) GetMetadataOk() (*[]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductJsonhalProductOutput) SetMetadata(v []string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductJsonhalProductOutput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ProductJsonhalProductOutput) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ProductJsonhalProductOutput) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsActive

`func (o *ProductJsonhalProductOutput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductJsonhalProductOutput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductJsonhalProductOutput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProductJsonhalProductOutput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProductJsonhalProductOutput) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProductJsonhalProductOutput) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetPrices

`func (o *ProductJsonhalProductOutput) GetPrices() []PriceJsonhalProductOutput`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ProductJsonhalProductOutput) GetPricesOk() (*[]PriceJsonhalProductOutput, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ProductJsonhalProductOutput) SetPrices(v []PriceJsonhalProductOutput)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ProductJsonhalProductOutput) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetDefaultPrice

`func (o *ProductJsonhalProductOutput) GetDefaultPrice() ProductJsonhalProductOutputDefaultPrice`

GetDefaultPrice returns the DefaultPrice field if non-nil, zero value otherwise.

### GetDefaultPriceOk

`func (o *ProductJsonhalProductOutput) GetDefaultPriceOk() (*ProductJsonhalProductOutputDefaultPrice, bool)`

GetDefaultPriceOk returns a tuple with the DefaultPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPrice

`func (o *ProductJsonhalProductOutput) SetDefaultPrice(v ProductJsonhalProductOutputDefaultPrice)`

SetDefaultPrice sets DefaultPrice field to given value.

### HasDefaultPrice

`func (o *ProductJsonhalProductOutput) HasDefaultPrice() bool`

HasDefaultPrice returns a boolean if a field has been set.

### SetDefaultPriceNil

`func (o *ProductJsonhalProductOutput) SetDefaultPriceNil(b bool)`

 SetDefaultPriceNil sets the value for DefaultPrice to be an explicit nil

### UnsetDefaultPrice
`func (o *ProductJsonhalProductOutput) UnsetDefaultPrice()`

UnsetDefaultPrice ensures that no value is present for DefaultPrice, not even an explicit nil
### GetRegion

`func (o *ProductJsonhalProductOutput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProductJsonhalProductOutput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProductJsonhalProductOutput) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProductJsonhalProductOutput) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ProductJsonhalProductOutput) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ProductJsonhalProductOutput) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


