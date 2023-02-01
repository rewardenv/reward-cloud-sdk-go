# \EnvVarTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvVarTypesGetCollection**](EnvVarTypeApi.md#ApiEnvVarTypesGetCollection) | **Get** /api/env_var_types | Retrieves the collection of EnvVarType resources.
[**ApiEnvVarTypesIdDelete**](EnvVarTypeApi.md#ApiEnvVarTypesIdDelete) | **Delete** /api/env_var_types/{id} | Removes the EnvVarType resource.
[**ApiEnvVarTypesIdGet**](EnvVarTypeApi.md#ApiEnvVarTypesIdGet) | **Get** /api/env_var_types/{id} | Retrieves a EnvVarType resource.
[**ApiEnvVarTypesIdPatch**](EnvVarTypeApi.md#ApiEnvVarTypesIdPatch) | **Patch** /api/env_var_types/{id} | Updates the EnvVarType resource.
[**ApiEnvVarTypesIdPut**](EnvVarTypeApi.md#ApiEnvVarTypesIdPut) | **Put** /api/env_var_types/{id} | Replaces the EnvVarType resource.
[**ApiEnvVarTypesPost**](EnvVarTypeApi.md#ApiEnvVarTypesPost) | **Post** /api/env_var_types | Creates a EnvVarType resource.



## ApiEnvVarTypesGetCollection

> []EnvVarType ApiEnvVarTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvVarType resources.



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
    resp, r, err := apiClient.EnvVarTypeApi.ApiEnvVarTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvVarTypeApi.ApiEnvVarTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvVarTypesGetCollection`: []EnvVarType
    fmt.Fprintf(os.Stdout, "Response from `EnvVarTypeApi.ApiEnvVarTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvVarTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]EnvVarType**](EnvVarType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvVarTypesIdDelete

> ApiEnvVarTypesIdDelete(ctx, id).Execute()

Removes the EnvVarType resource.



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
    id := "id_example" // string | EnvVarType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvVarTypeApi.ApiEnvVarTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvVarTypeApi.ApiEnvVarTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvVarType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvVarTypesIdDeleteRequest struct via the builder pattern


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


## ApiEnvVarTypesIdGet

> EnvVarType ApiEnvVarTypesIdGet(ctx, id).Execute()

Retrieves a EnvVarType resource.



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
    id := "id_example" // string | EnvVarType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvVarTypeApi.ApiEnvVarTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvVarTypeApi.ApiEnvVarTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvVarTypesIdGet`: EnvVarType
    fmt.Fprintf(os.Stdout, "Response from `EnvVarTypeApi.ApiEnvVarTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvVarType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvVarTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvVarType**](EnvVarType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvVarTypesIdPatch

> EnvVarType ApiEnvVarTypesIdPatch(ctx, id).EnvVarType(envVarType).Execute()

Updates the EnvVarType resource.



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
    id := "id_example" // string | EnvVarType identifier
    envVarType := *openapiclient.NewEnvVarType() // EnvVarType | The updated EnvVarType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvVarTypeApi.ApiEnvVarTypesIdPatch(context.Background(), id).EnvVarType(envVarType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvVarTypeApi.ApiEnvVarTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvVarTypesIdPatch`: EnvVarType
    fmt.Fprintf(os.Stdout, "Response from `EnvVarTypeApi.ApiEnvVarTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvVarType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvVarTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envVarType** | [**EnvVarType**](EnvVarType.md) | The updated EnvVarType resource | 

### Return type

[**EnvVarType**](EnvVarType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvVarTypesIdPut

> EnvVarType ApiEnvVarTypesIdPut(ctx, id).EnvVarType(envVarType).Execute()

Replaces the EnvVarType resource.



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
    id := "id_example" // string | EnvVarType identifier
    envVarType := *openapiclient.NewEnvVarType() // EnvVarType | The updated EnvVarType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvVarTypeApi.ApiEnvVarTypesIdPut(context.Background(), id).EnvVarType(envVarType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvVarTypeApi.ApiEnvVarTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvVarTypesIdPut`: EnvVarType
    fmt.Fprintf(os.Stdout, "Response from `EnvVarTypeApi.ApiEnvVarTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvVarType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvVarTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **envVarType** | [**EnvVarType**](EnvVarType.md) | The updated EnvVarType resource | 

### Return type

[**EnvVarType**](EnvVarType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvVarTypesPost

> EnvVarType ApiEnvVarTypesPost(ctx).EnvVarType(envVarType).Execute()

Creates a EnvVarType resource.



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
    envVarType := *openapiclient.NewEnvVarType() // EnvVarType | The new EnvVarType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvVarTypeApi.ApiEnvVarTypesPost(context.Background()).EnvVarType(envVarType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvVarTypeApi.ApiEnvVarTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvVarTypesPost`: EnvVarType
    fmt.Fprintf(os.Stdout, "Response from `EnvVarTypeApi.ApiEnvVarTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvVarTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **envVarType** | [**EnvVarType**](EnvVarType.md) | The new EnvVarType resource | 

### Return type

[**EnvVarType**](EnvVarType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

