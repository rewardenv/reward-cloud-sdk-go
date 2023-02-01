# ApiRoleGroupsGetCollection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | [**[]RoleGroupJsonhal**](RoleGroupJsonhal.md) |  | 
**TotalItems** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Links** | [**ApiComponentResourceLimitsGetCollection200ResponseLinks**](ApiComponentResourceLimitsGetCollection200ResponseLinks.md) |  | 

## Methods

### NewApiRoleGroupsGetCollection200Response

`func NewApiRoleGroupsGetCollection200Response(embedded []RoleGroupJsonhal, links ApiComponentResourceLimitsGetCollection200ResponseLinks, ) *ApiRoleGroupsGetCollection200Response`

NewApiRoleGroupsGetCollection200Response instantiates a new ApiRoleGroupsGetCollection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleGroupsGetCollection200ResponseWithDefaults

`func NewApiRoleGroupsGetCollection200ResponseWithDefaults() *ApiRoleGroupsGetCollection200Response`

NewApiRoleGroupsGetCollection200ResponseWithDefaults instantiates a new ApiRoleGroupsGetCollection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *ApiRoleGroupsGetCollection200Response) GetEmbedded() []RoleGroupJsonhal`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ApiRoleGroupsGetCollection200Response) GetEmbeddedOk() (*[]RoleGroupJsonhal, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ApiRoleGroupsGetCollection200Response) SetEmbedded(v []RoleGroupJsonhal)`

SetEmbedded sets Embedded field to given value.


### GetTotalItems

`func (o *ApiRoleGroupsGetCollection200Response) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ApiRoleGroupsGetCollection200Response) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ApiRoleGroupsGetCollection200Response) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ApiRoleGroupsGetCollection200Response) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ApiRoleGroupsGetCollection200Response) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ApiRoleGroupsGetCollection200Response) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ApiRoleGroupsGetCollection200Response) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ApiRoleGroupsGetCollection200Response) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetLinks

`func (o *ApiRoleGroupsGetCollection200Response) GetLinks() ApiComponentResourceLimitsGetCollection200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiRoleGroupsGetCollection200Response) GetLinksOk() (*ApiComponentResourceLimitsGetCollection200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiRoleGroupsGetCollection200Response) SetLinks(v ApiComponentResourceLimitsGetCollection200ResponseLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


