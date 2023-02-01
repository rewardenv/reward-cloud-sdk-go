# \EnvironmentAccessDevToolsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessDevToolsGetCollection**](EnvironmentAccessDevToolsApi.md#ApiEnvironmentAccessDevToolsGetCollection) | **Get** /api/environment_access_dev_tools | Retrieves the collection of EnvironmentAccessDevTools resources.
[**ApiEnvironmentAccessDevToolsIdDelete**](EnvironmentAccessDevToolsApi.md#ApiEnvironmentAccessDevToolsIdDelete) | **Delete** /api/environment_access_dev_tools/{id} | Removes the EnvironmentAccessDevTools resource.
[**ApiEnvironmentAccessDevToolsIdGet**](EnvironmentAccessDevToolsApi.md#ApiEnvironmentAccessDevToolsIdGet) | **Get** /api/environment_access_dev_tools/{id} | Retrieves a EnvironmentAccessDevTools resource.
[**ApiEnvironmentAccessDevToolsIdPatch**](EnvironmentAccessDevToolsApi.md#ApiEnvironmentAccessDevToolsIdPatch) | **Patch** /api/environment_access_dev_tools/{id} | Updates the EnvironmentAccessDevTools resource.
[**ApiEnvironmentAccessDevToolsIdPut**](EnvironmentAccessDevToolsApi.md#ApiEnvironmentAccessDevToolsIdPut) | **Put** /api/environment_access_dev_tools/{id} | Replaces the EnvironmentAccessDevTools resource.
[**ApiEnvironmentAccessDevToolsPost**](EnvironmentAccessDevToolsApi.md#ApiEnvironmentAccessDevToolsPost) | **Post** /api/environment_access_dev_tools | Creates a EnvironmentAccessDevTools resource.



## ApiEnvironmentAccessDevToolsGetCollection

> []EnvironmentAccessDevTools ApiEnvironmentAccessDevToolsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvironmentAccessDevTools resources.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDevToolsGetCollection`: []EnvironmentAccessDevTools
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDevToolsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDevToolsIdDelete

> ApiEnvironmentAccessDevToolsIdDelete(ctx, id).Execute()

Removes the EnvironmentAccessDevTools resource.



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
    id := "id_example" // string | EnvironmentAccessDevTools identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDevTools identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDevToolsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessDevToolsIdGet

> EnvironmentAccessDevTools ApiEnvironmentAccessDevToolsIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccessDevTools resource.



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
    id := "id_example" // string | EnvironmentAccessDevTools identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDevToolsIdGet`: EnvironmentAccessDevTools
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDevTools identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDevToolsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDevToolsIdPatch

> EnvironmentAccessDevTools ApiEnvironmentAccessDevToolsIdPatch(ctx, id).EnvironmentAccessDevTools(environmentAccessDevTools).Execute()

Updates the EnvironmentAccessDevTools resource.



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
    id := "id_example" // string | EnvironmentAccessDevTools identifier
    environmentAccessDevTools := *openapiclient.NewEnvironmentAccessDevTools() // EnvironmentAccessDevTools | The updated EnvironmentAccessDevTools resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPatch(context.Background(), id).EnvironmentAccessDevTools(environmentAccessDevTools).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDevToolsIdPatch`: EnvironmentAccessDevTools
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDevTools identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDevToolsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessDevTools** | [**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md) | The updated EnvironmentAccessDevTools resource | 

### Return type

[**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDevToolsIdPut

> EnvironmentAccessDevTools ApiEnvironmentAccessDevToolsIdPut(ctx, id).EnvironmentAccessDevTools(environmentAccessDevTools).Execute()

Replaces the EnvironmentAccessDevTools resource.



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
    id := "id_example" // string | EnvironmentAccessDevTools identifier
    environmentAccessDevTools := *openapiclient.NewEnvironmentAccessDevTools() // EnvironmentAccessDevTools | The updated EnvironmentAccessDevTools resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPut(context.Background(), id).EnvironmentAccessDevTools(environmentAccessDevTools).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDevToolsIdPut`: EnvironmentAccessDevTools
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDevTools identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDevToolsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessDevTools** | [**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md) | The updated EnvironmentAccessDevTools resource | 

### Return type

[**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDevToolsPost

> EnvironmentAccessDevTools ApiEnvironmentAccessDevToolsPost(ctx).EnvironmentAccessDevTools(environmentAccessDevTools).Execute()

Creates a EnvironmentAccessDevTools resource.



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
    environmentAccessDevTools := *openapiclient.NewEnvironmentAccessDevTools() // EnvironmentAccessDevTools | The new EnvironmentAccessDevTools resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsPost(context.Background()).EnvironmentAccessDevTools(environmentAccessDevTools).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDevToolsPost`: EnvironmentAccessDevTools
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDevToolsApi.ApiEnvironmentAccessDevToolsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDevToolsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessDevTools** | [**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md) | The new EnvironmentAccessDevTools resource | 

### Return type

[**EnvironmentAccessDevTools**](EnvironmentAccessDevTools.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

