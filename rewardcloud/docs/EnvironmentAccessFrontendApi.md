# \EnvironmentAccessFrontendApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessFrontendsGetCollection**](EnvironmentAccessFrontendApi.md#ApiEnvironmentAccessFrontendsGetCollection) | **Get** /api/environment_access_frontends | Retrieves the collection of EnvironmentAccessFrontend resources.
[**ApiEnvironmentAccessFrontendsIdDelete**](EnvironmentAccessFrontendApi.md#ApiEnvironmentAccessFrontendsIdDelete) | **Delete** /api/environment_access_frontends/{id} | Removes the EnvironmentAccessFrontend resource.
[**ApiEnvironmentAccessFrontendsIdGet**](EnvironmentAccessFrontendApi.md#ApiEnvironmentAccessFrontendsIdGet) | **Get** /api/environment_access_frontends/{id} | Retrieves a EnvironmentAccessFrontend resource.
[**ApiEnvironmentAccessFrontendsIdPatch**](EnvironmentAccessFrontendApi.md#ApiEnvironmentAccessFrontendsIdPatch) | **Patch** /api/environment_access_frontends/{id} | Updates the EnvironmentAccessFrontend resource.
[**ApiEnvironmentAccessFrontendsIdPut**](EnvironmentAccessFrontendApi.md#ApiEnvironmentAccessFrontendsIdPut) | **Put** /api/environment_access_frontends/{id} | Replaces the EnvironmentAccessFrontend resource.
[**ApiEnvironmentAccessFrontendsPost**](EnvironmentAccessFrontendApi.md#ApiEnvironmentAccessFrontendsPost) | **Post** /api/environment_access_frontends | Creates a EnvironmentAccessFrontend resource.



## ApiEnvironmentAccessFrontendsGetCollection

> []EnvironmentAccessFrontend ApiEnvironmentAccessFrontendsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvironmentAccessFrontend resources.



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
    resp, r, err := apiClient.EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessFrontendsGetCollection`: []EnvironmentAccessFrontend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessFrontendsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessFrontendsIdDelete

> ApiEnvironmentAccessFrontendsIdDelete(ctx, id).Execute()

Removes the EnvironmentAccessFrontend resource.



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
    id := "id_example" // string | EnvironmentAccessFrontend identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessFrontend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessFrontendsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessFrontendsIdGet

> EnvironmentAccessFrontend ApiEnvironmentAccessFrontendsIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccessFrontend resource.



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
    id := "id_example" // string | EnvironmentAccessFrontend identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessFrontendsIdGet`: EnvironmentAccessFrontend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessFrontend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessFrontendsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessFrontendsIdPatch

> EnvironmentAccessFrontend ApiEnvironmentAccessFrontendsIdPatch(ctx, id).EnvironmentAccessFrontend(environmentAccessFrontend).Execute()

Updates the EnvironmentAccessFrontend resource.



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
    id := "id_example" // string | EnvironmentAccessFrontend identifier
    environmentAccessFrontend := *openapiclient.NewEnvironmentAccessFrontend() // EnvironmentAccessFrontend | The updated EnvironmentAccessFrontend resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdPatch(context.Background(), id).EnvironmentAccessFrontend(environmentAccessFrontend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessFrontendsIdPatch`: EnvironmentAccessFrontend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessFrontend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessFrontendsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessFrontend** | [**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md) | The updated EnvironmentAccessFrontend resource | 

### Return type

[**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessFrontendsIdPut

> EnvironmentAccessFrontend ApiEnvironmentAccessFrontendsIdPut(ctx, id).EnvironmentAccessFrontend(environmentAccessFrontend).Execute()

Replaces the EnvironmentAccessFrontend resource.



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
    id := "id_example" // string | EnvironmentAccessFrontend identifier
    environmentAccessFrontend := *openapiclient.NewEnvironmentAccessFrontend() // EnvironmentAccessFrontend | The updated EnvironmentAccessFrontend resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdPut(context.Background(), id).EnvironmentAccessFrontend(environmentAccessFrontend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessFrontendsIdPut`: EnvironmentAccessFrontend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessFrontend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessFrontendsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessFrontend** | [**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md) | The updated EnvironmentAccessFrontend resource | 

### Return type

[**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessFrontendsPost

> EnvironmentAccessFrontend ApiEnvironmentAccessFrontendsPost(ctx).EnvironmentAccessFrontend(environmentAccessFrontend).Execute()

Creates a EnvironmentAccessFrontend resource.



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
    environmentAccessFrontend := *openapiclient.NewEnvironmentAccessFrontend() // EnvironmentAccessFrontend | The new EnvironmentAccessFrontend resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsPost(context.Background()).EnvironmentAccessFrontend(environmentAccessFrontend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessFrontendsPost`: EnvironmentAccessFrontend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessFrontendApi.ApiEnvironmentAccessFrontendsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessFrontendsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessFrontend** | [**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md) | The new EnvironmentAccessFrontend resource | 

### Return type

[**EnvironmentAccessFrontend**](EnvironmentAccessFrontend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

