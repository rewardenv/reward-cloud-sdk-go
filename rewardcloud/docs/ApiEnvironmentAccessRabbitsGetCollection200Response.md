# ApiEnvironmentAccessRabbitsGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**[]EnvironmentAccessRabbitJsonhal**](EnvironmentAccessRabbitJsonhal.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Links** | [**ApiComponentResourceLimitsGetCollection200ResponseLinks**](ApiComponentResourceLimitsGetCollection200ResponseLinks.md) |  | 

## Methods

### NewApiEnvironmentAccessRabbitsGetCollection200Response

`func NewApiEnvironmentAccessRabbitsGetCollection200Response(embedded []EnvironmentAccessRabbitJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks, ) *ApiEnvironmentAccessRabbitsGetCollection200Response`

NewApiEnvironmentAccessRabbitsGetCollection200Response instantiates a new ApiEnvironmentAccessRabbitsGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEnvironmentAccessRabbitsGetCollection200ResponseWithDefaults

`func NewApiEnvironmentAccessRabbitsGetCollection200ResponseWithDefaults() *ApiEnvironmentAccessRabbitsGetCollection200Response`

NewApiEnvironmentAccessRabbitsGetCollection200ResponseWithDefaults instantiates a new ApiEnvironmentAccessRabbitsGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetEmbedded() []EnvironmentAccessRabbitJsonhal`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetEmbeddedOk() (*[]EnvironmentAccessRabbitJsonhal, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) SetEmbedded(v []EnvironmentAccessRabbitJsonhal)`

SetEmbedded sets Embedded field to given value.


### GetTotalItems

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetLinks

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiEnvironmentAccessRabbitsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


