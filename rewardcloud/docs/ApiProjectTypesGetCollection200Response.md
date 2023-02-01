# ApiProjectTypesGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**[]ProjectTypeJsonhal**](ProjectTypeJsonhal.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Links** | [**ApiComponentResourceLimitsGetCollection200ResponseLinks**](ApiComponentResourceLimitsGetCollection200ResponseLinks.md) |  | 

## Methods

### NewApiProjectTypesGetCollection200Response

`func NewApiProjectTypesGetCollection200Response(embedded []ProjectTypeJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks, ) *ApiProjectTypesGetCollection200Response`

NewApiProjectTypesGetCollection200Response instantiates a new ApiProjectTypesGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiProjectTypesGetCollection200ResponseWithDefaults

`func NewApiProjectTypesGetCollection200ResponseWithDefaults() *ApiProjectTypesGetCollection200Response`

NewApiProjectTypesGetCollection200ResponseWithDefaults instantiates a new ApiProjectTypesGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ApiProjectTypesGetCollection200Response) GetEmbedded() []ProjectTypeJsonhal`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiProjectTypesGetCollection200Response) GetEmbeddedOk() (*[]ProjectTypeJsonhal, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiProjectTypesGetCollection200Response) SetEmbedded(v []ProjectTypeJsonhal)`

SetEmbedded sets Embedded field to given value.


### GetTotalItems

`func (o *ApiProjectTypesGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiProjectTypesGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiProjectTypesGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiProjectTypesGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ApiProjectTypesGetCollection200Response) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiProjectTypesGetCollection200Response) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiProjectTypesGetCollection200Response) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiProjectTypesGetCollection200Response) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetLinks

`func (o *ApiProjectTypesGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiProjectTypesGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiProjectTypesGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


