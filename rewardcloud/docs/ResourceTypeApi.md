# \ResourceTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiResourceTypesGetCollection**](ResourceTypeApi.md#ApiResourceTypesGetCollection) | **Get** /api/resource_types | Retrieves the collection of ResourceType resources.
[**ApiResourceTypesIdDelete**](ResourceTypeApi.md#ApiResourceTypesIdDelete) | **Delete** /api/resource_types/{id} | Removes the ResourceType resource.
[**ApiResourceTypesIdGet**](ResourceTypeApi.md#ApiResourceTypesIdGet) | **Get** /api/resource_types/{id} | Retrieves a ResourceType resource.
[**ApiResourceTypesIdPatch**](ResourceTypeApi.md#ApiResourceTypesIdPatch) | **Patch** /api/resource_types/{id} | Updates the ResourceType resource.
[**ApiResourceTypesIdPut**](ResourceTypeApi.md#ApiResourceTypesIdPut) | **Put** /api/resource_types/{id} | Replaces the ResourceType resource.
[**ApiResourceTypesPost**](ResourceTypeApi.md#ApiResourceTypesPost) | **Post** /api/resource_types | Creates a ResourceType resource.



## ApiResourceTypesGetCollection

> []ResourceTypeResourceTypeOutput ApiResourceTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ResourceTypeLimit(resourceTypeLimit).ResourceTypeLimit2(resourceTypeLimit2).Execute()

Retrieves the collection of ResourceType resources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    itemsPerPage := int32(56) // int32 | The number of items per page (optional) (default to 30)
    resourceTypeLimit := "resourceTypeLimit_example" // string |  (optional)
    resourceTypeLimit2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeApi.ApiResourceTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ResourceTypeLimit(resourceTypeLimit).ResourceTypeLimit2(resourceTypeLimit2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeApi.ApiResourceTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypesGetCollection`: []ResourceTypeResourceTypeOutput
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeApi.ApiResourceTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **resourceTypeLimit** | **string** |  | 
 **resourceTypeLimit2** | **[]string** |  | 

### Return type

[**[]ResourceTypeResourceTypeOutput**](ResourceTypeResourceTypeOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypesIdDelete

> ApiResourceTypesIdDelete(ctx, id).Execute()

Removes the ResourceType resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ResourceType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeApi.ApiResourceTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeApi.ApiResourceTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypesIdGet

> ResourceTypeResourceTypeOutput ApiResourceTypesIdGet(ctx, id).Execute()

Retrieves a ResourceType resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ResourceType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeApi.ApiResourceTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeApi.ApiResourceTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypesIdGet`: ResourceTypeResourceTypeOutput
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeApi.ApiResourceTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceTypeResourceTypeOutput**](ResourceTypeResourceTypeOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypesIdPatch

> ResourceTypeResourceTypeOutput ApiResourceTypesIdPatch(ctx, id).ResourceTypeResourceTypeInput(resourceTypeResourceTypeInput).Execute()

Updates the ResourceType resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ResourceType identifier
    resourceTypeResourceTypeInput := *openapiclient.NewResourceTypeResourceTypeInput() // ResourceTypeResourceTypeInput | The updated ResourceType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeApi.ApiResourceTypesIdPatch(context.Background(), id).ResourceTypeResourceTypeInput(resourceTypeResourceTypeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeApi.ApiResourceTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypesIdPatch`: ResourceTypeResourceTypeOutput
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeApi.ApiResourceTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceTypeResourceTypeInput** | [**ResourceTypeResourceTypeInput**](ResourceTypeResourceTypeInput.md) | The updated ResourceType resource | 

### Return type

[**ResourceTypeResourceTypeOutput**](ResourceTypeResourceTypeOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypesIdPut

> ResourceTypeResourceTypeOutput ApiResourceTypesIdPut(ctx, id).ResourceTypeResourceTypeInput(resourceTypeResourceTypeInput).Execute()

Replaces the ResourceType resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ResourceType identifier
    resourceTypeResourceTypeInput := *openapiclient.NewResourceTypeResourceTypeInput() // ResourceTypeResourceTypeInput | The updated ResourceType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeApi.ApiResourceTypesIdPut(context.Background(), id).ResourceTypeResourceTypeInput(resourceTypeResourceTypeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeApi.ApiResourceTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypesIdPut`: ResourceTypeResourceTypeOutput
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeApi.ApiResourceTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceTypeResourceTypeInput** | [**ResourceTypeResourceTypeInput**](ResourceTypeResourceTypeInput.md) | The updated ResourceType resource | 

### Return type

[**ResourceTypeResourceTypeOutput**](ResourceTypeResourceTypeOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypesPost

> ResourceTypeResourceTypeOutput ApiResourceTypesPost(ctx).ResourceTypeResourceTypeInput(resourceTypeResourceTypeInput).Execute()

Creates a ResourceType resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    resourceTypeResourceTypeInput := *openapiclient.NewResourceTypeResourceTypeInput() // ResourceTypeResourceTypeInput | The new ResourceType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeApi.ApiResourceTypesPost(context.Background()).ResourceTypeResourceTypeInput(resourceTypeResourceTypeInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeApi.ApiResourceTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypesPost`: ResourceTypeResourceTypeOutput
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeApi.ApiResourceTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceTypeResourceTypeInput** | [**ResourceTypeResourceTypeInput**](ResourceTypeResourceTypeInput.md) | The new ResourceType resource | 

### Return type

[**ResourceTypeResourceTypeOutput**](ResourceTypeResourceTypeOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

