# ApiEnvironmentAccessBackendsGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**[]EnvironmentAccessBackendJsonhal**](EnvironmentAccessBackendJsonhal.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Links** | [**ApiComponentResourceLimitsGetCollection200ResponseLinks**](ApiComponentResourceLimitsGetCollection200ResponseLinks.md) |  | 

## Methods

### NewApiEnvironmentAccessBackendsGetCollection200Response

`func NewApiEnvironmentAccessBackendsGetCollection200Response(embedded []EnvironmentAccessBackendJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks, ) *ApiEnvironmentAccessBackendsGetCollection200Response`

NewApiEnvironmentAccessBackendsGetCollection200Response instantiates a new ApiEnvironmentAccessBackendsGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEnvironmentAccessBackendsGetCollection200ResponseWithDefaults

`func NewApiEnvironmentAccessBackendsGetCollection200ResponseWithDefaults() *ApiEnvironmentAccessBackendsGetCollection200Response`

NewApiEnvironmentAccessBackendsGetCollection200ResponseWithDefaults instantiates a new ApiEnvironmentAccessBackendsGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetEmbedded() []EnvironmentAccessBackendJsonhal`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetEmbeddedOk() (*[]EnvironmentAccessBackendJsonhal, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) SetEmbedded(v []EnvironmentAccessBackendJsonhal)`

SetEmbedded sets Embedded field to given value.


### GetTotalItems

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetLinks

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiEnvironmentAccessBackendsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


