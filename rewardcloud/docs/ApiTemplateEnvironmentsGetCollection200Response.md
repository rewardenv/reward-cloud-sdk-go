# ApiTemplateEnvironmentsGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**[]TemplateEnvironmentJsonhalTemplateEnvironmentOutput**](TemplateEnvironmentJsonhalTemplateEnvironmentOutput.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Links** | [**ApiComponentResourceLimitsGetCollection200ResponseLinks**](ApiComponentResourceLimitsGetCollection200ResponseLinks.md) |  | 

## Methods

### NewApiTemplateEnvironmentsGetCollection200Response

`func NewApiTemplateEnvironmentsGetCollection200Response(embedded []TemplateEnvironmentJsonhalTemplateEnvironmentOutput, links ApiComponentResourceLimitsGetCollection200ResponseLinks, ) *ApiTemplateEnvironmentsGetCollection200Response`

NewApiTemplateEnvironmentsGetCollection200Response instantiates a new ApiTemplateEnvironmentsGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTemplateEnvironmentsGetCollection200ResponseWithDefaults

`func NewApiTemplateEnvironmentsGetCollection200ResponseWithDefaults() *ApiTemplateEnvironmentsGetCollection200Response`

NewApiTemplateEnvironmentsGetCollection200ResponseWithDefaults instantiates a new ApiTemplateEnvironmentsGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetEmbedded() []TemplateEnvironmentJsonhalTemplateEnvironmentOutput`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetEmbeddedOk() (*[]TemplateEnvironmentJsonhalTemplateEnvironmentOutput, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiTemplateEnvironmentsGetCollection200Response) SetEmbedded(v []TemplateEnvironmentJsonhalTemplateEnvironmentOutput)`

SetEmbedded sets Embedded field to given value.


### GetTotalItems

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiTemplateEnvironmentsGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiTemplateEnvironmentsGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiTemplateEnvironmentsGetCollection200Response) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiTemplateEnvironmentsGetCollection200Response) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetLinks

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiTemplateEnvironmentsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiTemplateEnvironmentsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


