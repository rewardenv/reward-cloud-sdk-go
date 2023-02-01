# \ComponentVersionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiComponentVersionsGetCollection**](ComponentVersionApi.md#ApiComponentVersionsGetCollection) | **Get** /api/component_versions | Retrieves the collection of ComponentVersion resources.
[**ApiComponentVersionsIdDelete**](ComponentVersionApi.md#ApiComponentVersionsIdDelete) | **Delete** /api/component_versions/{id} | Removes the ComponentVersion resource.
[**ApiComponentVersionsIdGet**](ComponentVersionApi.md#ApiComponentVersionsIdGet) | **Get** /api/component_versions/{id} | Retrieves a ComponentVersion resource.
[**ApiComponentVersionsIdPatch**](ComponentVersionApi.md#ApiComponentVersionsIdPatch) | **Patch** /api/component_versions/{id} | Updates the ComponentVersion resource.
[**ApiComponentVersionsIdPut**](ComponentVersionApi.md#ApiComponentVersionsIdPut) | **Put** /api/component_versions/{id} | Replaces the ComponentVersion resource.
[**ApiComponentVersionsPost**](ComponentVersionApi.md#ApiComponentVersionsPost) | **Post** /api/component_versions | Creates a ComponentVersion resource.



## ApiComponentVersionsGetCollection

> []ComponentVersion ApiComponentVersionsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Component(component).Component2(component2).Execute()

Retrieves the collection of ComponentVersion resources.



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
    component := "component_example" // string |  (optional)
    component2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionApi.ApiComponentVersionsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Component(component).Component2(component2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionApi.ApiComponentVersionsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionsGetCollection`: []ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionApi.ApiComponentVersionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **component** | **string** |  | 
 **component2** | **[]string** |  | 

### Return type

[**[]ComponentVersion**](ComponentVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionsIdDelete

> ApiComponentVersionsIdDelete(ctx, id).Execute()

Removes the ComponentVersion resource.



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
    id := "id_example" // string | ComponentVersion identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionApi.ApiComponentVersionsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionApi.ApiComponentVersionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionsIdDeleteRequest struct via the builder pattern


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


## ApiComponentVersionsIdGet

> ComponentVersion ApiComponentVersionsIdGet(ctx, id).Execute()

Retrieves a ComponentVersion resource.



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
    id := "id_example" // string | ComponentVersion identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionApi.ApiComponentVersionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionApi.ApiComponentVersionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionsIdGet`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionApi.ApiComponentVersionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionsIdPatch

> ComponentVersion ApiComponentVersionsIdPatch(ctx, id).ComponentVersion(componentVersion).Execute()

Updates the ComponentVersion resource.



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
    id := "id_example" // string | ComponentVersion identifier
    componentVersion := *openapiclient.NewComponentVersion() // ComponentVersion | The updated ComponentVersion resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionApi.ApiComponentVersionsIdPatch(context.Background(), id).ComponentVersion(componentVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionApi.ApiComponentVersionsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionsIdPatch`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionApi.ApiComponentVersionsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentVersion** | [**ComponentVersion**](ComponentVersion.md) | The updated ComponentVersion resource | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionsIdPut

> ComponentVersion ApiComponentVersionsIdPut(ctx, id).ComponentVersion(componentVersion).Execute()

Replaces the ComponentVersion resource.



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
    id := "id_example" // string | ComponentVersion identifier
    componentVersion := *openapiclient.NewComponentVersion() // ComponentVersion | The updated ComponentVersion resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionApi.ApiComponentVersionsIdPut(context.Background(), id).ComponentVersion(componentVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionApi.ApiComponentVersionsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionsIdPut`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionApi.ApiComponentVersionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersion identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentVersion** | [**ComponentVersion**](ComponentVersion.md) | The updated ComponentVersion resource | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionsPost

> ComponentVersion ApiComponentVersionsPost(ctx).ComponentVersion(componentVersion).Execute()

Creates a ComponentVersion resource.



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
    componentVersion := *openapiclient.NewComponentVersion() // ComponentVersion | The new ComponentVersion resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionApi.ApiComponentVersionsPost(context.Background()).ComponentVersion(componentVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionApi.ApiComponentVersionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionsPost`: ComponentVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionApi.ApiComponentVersionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentVersion** | [**ComponentVersion**](ComponentVersion.md) | The new ComponentVersion resource | 

### Return type

[**ComponentVersion**](ComponentVersion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

