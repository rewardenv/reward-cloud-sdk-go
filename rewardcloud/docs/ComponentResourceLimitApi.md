# \ComponentResourceLimitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiComponentResourceLimitsGetCollection**](ComponentResourceLimitApi.md#ApiComponentResourceLimitsGetCollection) | **Get** /api/component_resource_limits | Retrieves the collection of ComponentResourceLimit resources.
[**ApiComponentResourceLimitsIdDelete**](ComponentResourceLimitApi.md#ApiComponentResourceLimitsIdDelete) | **Delete** /api/component_resource_limits/{id} | Removes the ComponentResourceLimit resource.
[**ApiComponentResourceLimitsIdGet**](ComponentResourceLimitApi.md#ApiComponentResourceLimitsIdGet) | **Get** /api/component_resource_limits/{id} | Retrieves a ComponentResourceLimit resource.
[**ApiComponentResourceLimitsIdPatch**](ComponentResourceLimitApi.md#ApiComponentResourceLimitsIdPatch) | **Patch** /api/component_resource_limits/{id} | Updates the ComponentResourceLimit resource.
[**ApiComponentResourceLimitsIdPut**](ComponentResourceLimitApi.md#ApiComponentResourceLimitsIdPut) | **Put** /api/component_resource_limits/{id} | Replaces the ComponentResourceLimit resource.
[**ApiComponentResourceLimitsPost**](ComponentResourceLimitApi.md#ApiComponentResourceLimitsPost) | **Post** /api/component_resource_limits | Creates a ComponentResourceLimit resource.



## ApiComponentResourceLimitsGetCollection

> []ComponentResourceLimit ApiComponentResourceLimitsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).ResourceType(resourceType).ResourceType2(resourceType2).ComponentVersion(componentVersion).ComponentVersion2(componentVersion2).Execute()

Retrieves the collection of ComponentResourceLimit resources.



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
    projectTypeVersion := "projectTypeVersion_example" // string |  (optional)
    projectTypeVersion2 := []string{"Inner_example"} // []string |  (optional)
    resourceType := "resourceType_example" // string |  (optional)
    resourceType2 := []string{"Inner_example"} // []string |  (optional)
    componentVersion := "componentVersion_example" // string |  (optional)
    componentVersion2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentResourceLimitApi.ApiComponentResourceLimitsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).ResourceType(resourceType).ResourceType2(resourceType2).ComponentVersion(componentVersion).ComponentVersion2(componentVersion2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentResourceLimitApi.ApiComponentResourceLimitsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentResourceLimitsGetCollection`: []ComponentResourceLimit
    fmt.Fprintf(os.Stdout, "Response from `ComponentResourceLimitApi.ApiComponentResourceLimitsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentResourceLimitsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **projectTypeVersion** | **string** |  | 
 **projectTypeVersion2** | **[]string** |  | 
 **resourceType** | **string** |  | 
 **resourceType2** | **[]string** |  | 
 **componentVersion** | **string** |  | 
 **componentVersion2** | **[]string** |  | 

### Return type

[**[]ComponentResourceLimit**](ComponentResourceLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentResourceLimitsIdDelete

> ApiComponentResourceLimitsIdDelete(ctx, id).Execute()

Removes the ComponentResourceLimit resource.



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
    id := "id_example" // string | ComponentResourceLimit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentResourceLimitApi.ApiComponentResourceLimitsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentResourceLimitApi.ApiComponentResourceLimitsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentResourceLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentResourceLimitsIdDeleteRequest struct via the builder pattern


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


## ApiComponentResourceLimitsIdGet

> ComponentResourceLimit ApiComponentResourceLimitsIdGet(ctx, id).Execute()

Retrieves a ComponentResourceLimit resource.



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
    id := "id_example" // string | ComponentResourceLimit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentResourceLimitApi.ApiComponentResourceLimitsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentResourceLimitApi.ApiComponentResourceLimitsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentResourceLimitsIdGet`: ComponentResourceLimit
    fmt.Fprintf(os.Stdout, "Response from `ComponentResourceLimitApi.ApiComponentResourceLimitsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentResourceLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentResourceLimitsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentResourceLimit**](ComponentResourceLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentResourceLimitsIdPatch

> ComponentResourceLimit ApiComponentResourceLimitsIdPatch(ctx, id).ComponentResourceLimit(componentResourceLimit).Execute()

Updates the ComponentResourceLimit resource.



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
    id := "id_example" // string | ComponentResourceLimit identifier
    componentResourceLimit := *openapiclient.NewComponentResourceLimit() // ComponentResourceLimit | The updated ComponentResourceLimit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentResourceLimitApi.ApiComponentResourceLimitsIdPatch(context.Background(), id).ComponentResourceLimit(componentResourceLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentResourceLimitApi.ApiComponentResourceLimitsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentResourceLimitsIdPatch`: ComponentResourceLimit
    fmt.Fprintf(os.Stdout, "Response from `ComponentResourceLimitApi.ApiComponentResourceLimitsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentResourceLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentResourceLimitsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentResourceLimit** | [**ComponentResourceLimit**](ComponentResourceLimit.md) | The updated ComponentResourceLimit resource | 

### Return type

[**ComponentResourceLimit**](ComponentResourceLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentResourceLimitsIdPut

> ComponentResourceLimit ApiComponentResourceLimitsIdPut(ctx, id).ComponentResourceLimit(componentResourceLimit).Execute()

Replaces the ComponentResourceLimit resource.



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
    id := "id_example" // string | ComponentResourceLimit identifier
    componentResourceLimit := *openapiclient.NewComponentResourceLimit() // ComponentResourceLimit | The updated ComponentResourceLimit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentResourceLimitApi.ApiComponentResourceLimitsIdPut(context.Background(), id).ComponentResourceLimit(componentResourceLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentResourceLimitApi.ApiComponentResourceLimitsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentResourceLimitsIdPut`: ComponentResourceLimit
    fmt.Fprintf(os.Stdout, "Response from `ComponentResourceLimitApi.ApiComponentResourceLimitsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentResourceLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentResourceLimitsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentResourceLimit** | [**ComponentResourceLimit**](ComponentResourceLimit.md) | The updated ComponentResourceLimit resource | 

### Return type

[**ComponentResourceLimit**](ComponentResourceLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentResourceLimitsPost

> ComponentResourceLimit ApiComponentResourceLimitsPost(ctx).ComponentResourceLimit(componentResourceLimit).Execute()

Creates a ComponentResourceLimit resource.



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
    componentResourceLimit := *openapiclient.NewComponentResourceLimit() // ComponentResourceLimit | The new ComponentResourceLimit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentResourceLimitApi.ApiComponentResourceLimitsPost(context.Background()).ComponentResourceLimit(componentResourceLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentResourceLimitApi.ApiComponentResourceLimitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentResourceLimitsPost`: ComponentResourceLimit
    fmt.Fprintf(os.Stdout, "Response from `ComponentResourceLimitApi.ApiComponentResourceLimitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentResourceLimitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentResourceLimit** | [**ComponentResourceLimit**](ComponentResourceLimit.md) | The new ComponentResourceLimit resource | 

### Return type

[**ComponentResourceLimit**](ComponentResourceLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

