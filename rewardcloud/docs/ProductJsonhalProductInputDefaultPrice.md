# ProductJsonhalProductInputDefaultPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**AbstractEnvironmentJsonhalLinks**](AbstractEnvironmentJsonhalLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**Uuid** | Pointer to **NullableString** |  | [optional] [readonly] 
**PaymentId** | Pointer to **NullableString** |  | [optional] 
**Amount** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProductJsonhalProductInputDefaultPrice

`func NewProductJsonhalProductInputDefaultPrice() *ProductJsonhalProductInputDefaultPrice`

NewProductJsonhalProductInputDefaultPrice instantiates a new ProductJsonhalProductInputDefaultPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductJsonhalProductInputDefaultPriceWithDefaults

`func NewProductJsonhalProductInputDefaultPriceWithDefaults() *ProductJsonhalProductInputDefaultPrice`

NewProductJsonhalProductInputDefaultPriceWithDefaults instantiates a new ProductJsonhalProductInputDefaultPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProductJsonhalProductInputDefaultPrice) GetLinks() AbstractEnvironmentJsonhalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetLinksOk() (*AbstractEnvironmentJsonhalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProductJsonhalProductInputDefaultPrice) SetLinks(v AbstractEnvironmentJsonhalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProductJsonhalProductInputDefaultPrice) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ProductJsonhalProductInputDefaultPrice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductJsonhalProductInputDefaultPrice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductJsonhalProductInputDefaultPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ProductJsonhalProductInputDefaultPrice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ProductJsonhalProductInputDefaultPrice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ProductJsonhalProductInputDefaultPrice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *ProductJsonhalProductInputDefaultPrice) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *ProductJsonhalProductInputDefaultPrice) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetPaymentId

`func (o *ProductJsonhalProductInputDefaultPrice) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *ProductJsonhalProductInputDefaultPrice) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *ProductJsonhalProductInputDefaultPrice) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### SetPaymentIdNil

`func (o *ProductJsonhalProductInputDefaultPrice) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *ProductJsonhalProductInputDefaultPrice) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetAmount

`func (o *ProductJsonhalProductInputDefaultPrice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProductJsonhalProductInputDefaultPrice) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ProductJsonhalProductInputDefaultPrice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ProductJsonhalProductInputDefaultPrice) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ProductJsonhalProductInputDefaultPrice) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *ProductJsonhalProductInputDefaultPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductJsonhalProductInputDefaultPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductJsonhalProductInputDefaultPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ProductJsonhalProductInputDefaultPrice) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ProductJsonhalProductInputDefaultPrice) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetIsActive

`func (o *ProductJsonhalProductInputDefaultPrice) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductJsonhalProductInputDefaultPrice) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductJsonhalProductInputDefaultPrice) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProductJsonhalProductInputDefaultPrice) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *ProductJsonhalProductInputDefaultPrice) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProductJsonhalProductInputDefaultPrice) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


