# ApiComponentVersionEnvVarExamplesGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**[]ComponentVersionEnvVarExampleJsonhal**](ComponentVersionEnvVarExampleJsonhal.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Links** | [**ApiComponentResourceLimitsGetCollection200ResponseLinks**](ApiComponentResourceLimitsGetCollection200ResponseLinks.md) |  | 

## Methods

### NewApiComponentVersionEnvVarExamplesGetCollection200Response

`func NewApiComponentVersionEnvVarExamplesGetCollection200Response(embedded []ComponentVersionEnvVarExampleJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks, ) *ApiComponentVersionEnvVarExamplesGetCollection200Response`

NewApiComponentVersionEnvVarExamplesGetCollection200Response instantiates a new ApiComponentVersionEnvVarExamplesGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentVersionEnvVarExamplesGetCollection200ResponseWithDefaults

`func NewApiComponentVersionEnvVarExamplesGetCollection200ResponseWithDefaults() *ApiComponentVersionEnvVarExamplesGetCollection200Response`

NewApiComponentVersionEnvVarExamplesGetCollection200ResponseWithDefaults instantiates a new ApiComponentVersionEnvVarExamplesGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetEmbedded() []ComponentVersionEnvVarExampleJsonhal`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetEmbeddedOk() (*[]ComponentVersionEnvVarExampleJsonhal, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) SetEmbedded(v []ComponentVersionEnvVarExampleJsonhal)`

SetEmbedded sets Embedded field to given value.


### GetTotalItems

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetLinks

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiComponentVersionEnvVarExamplesGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


