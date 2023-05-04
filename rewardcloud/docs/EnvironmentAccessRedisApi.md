# \EnvironmentAccessRedisApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessRedisGetCollection**](EnvironmentAccessRedisApi.md#ApiEnvironmentAccessRedisGetCollection) | **Get** /api/environment_access_redis | Retrieves the collection of EnvironmentAccessRedis resources.
[**ApiEnvironmentAccessRedisIdDelete**](EnvironmentAccessRedisApi.md#ApiEnvironmentAccessRedisIdDelete) | **Delete** /api/environment_access_redis/{id} | Removes the EnvironmentAccessRedis resource.
[**ApiEnvironmentAccessRedisIdGet**](EnvironmentAccessRedisApi.md#ApiEnvironmentAccessRedisIdGet) | **Get** /api/environment_access_redis/{id} | Retrieves a EnvironmentAccessRedis resource.
[**ApiEnvironmentAccessRedisIdPatch**](EnvironmentAccessRedisApi.md#ApiEnvironmentAccessRedisIdPatch) | **Patch** /api/environment_access_redis/{id} | Updates the EnvironmentAccessRedis resource.
[**ApiEnvironmentAccessRedisIdPut**](EnvironmentAccessRedisApi.md#ApiEnvironmentAccessRedisIdPut) | **Put** /api/environment_access_redis/{id} | Replaces the EnvironmentAccessRedis resource.
[**ApiEnvironmentAccessRedisPost**](EnvironmentAccessRedisApi.md#ApiEnvironmentAccessRedisPost) | **Post** /api/environment_access_redis | Creates a EnvironmentAccessRedis resource.



## ApiEnvironmentAccessRedisGetCollection

> ApiEnvironmentAccessRedisGetCollection200Response ApiEnvironmentAccessRedisGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvironmentAccessRedis resources.



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
    resp, r, err := apiClient.EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRedisGetCollection`: ApiEnvironmentAccessRedisGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRedisGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**ApiEnvironmentAccessRedisGetCollection200Response**](ApiEnvironmentAccessRedisGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRedisIdDelete

> ApiEnvironmentAccessRedisIdDelete(ctx, id).Execute()

Removes the EnvironmentAccessRedis resource.



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
    id := "id_example" // string | EnvironmentAccessRedis identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRedis identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRedisIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessRedisIdGet

> EnvironmentAccessRedisJsonhal ApiEnvironmentAccessRedisIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccessRedis resource.



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
    id := "id_example" // string | EnvironmentAccessRedis identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRedisIdGet`: EnvironmentAccessRedisJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRedis identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRedisIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccessRedisJsonhal**](EnvironmentAccessRedisJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRedisIdPatch

> EnvironmentAccessRedisJsonhal ApiEnvironmentAccessRedisIdPatch(ctx, id).EnvironmentAccessRedis(environmentAccessRedis).Execute()

Updates the EnvironmentAccessRedis resource.



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
    id := "id_example" // string | EnvironmentAccessRedis identifier
    environmentAccessRedis := *openapiclient.NewEnvironmentAccessRedis() // EnvironmentAccessRedis | The updated EnvironmentAccessRedis resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdPatch(context.Background(), id).EnvironmentAccessRedis(environmentAccessRedis).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRedisIdPatch`: EnvironmentAccessRedisJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRedis identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRedisIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessRedis** | [**EnvironmentAccessRedis**](EnvironmentAccessRedis.md) | The updated EnvironmentAccessRedis resource | 

### Return type

[**EnvironmentAccessRedisJsonhal**](EnvironmentAccessRedisJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRedisIdPut

> EnvironmentAccessRedisJsonhal ApiEnvironmentAccessRedisIdPut(ctx, id).EnvironmentAccessRedisJsonhal(environmentAccessRedisJsonhal).Execute()

Replaces the EnvironmentAccessRedis resource.



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
    id := "id_example" // string | EnvironmentAccessRedis identifier
    environmentAccessRedisJsonhal := *openapiclient.NewEnvironmentAccessRedisJsonhal() // EnvironmentAccessRedisJsonhal | The updated EnvironmentAccessRedis resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdPut(context.Background(), id).EnvironmentAccessRedisJsonhal(environmentAccessRedisJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRedisIdPut`: EnvironmentAccessRedisJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRedis identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRedisIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessRedisJsonhal** | [**EnvironmentAccessRedisJsonhal**](EnvironmentAccessRedisJsonhal.md) | The updated EnvironmentAccessRedis resource | 

### Return type

[**EnvironmentAccessRedisJsonhal**](EnvironmentAccessRedisJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRedisPost

> EnvironmentAccessRedisJsonhal ApiEnvironmentAccessRedisPost(ctx).EnvironmentAccessRedisJsonhal(environmentAccessRedisJsonhal).Execute()

Creates a EnvironmentAccessRedis resource.



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
    environmentAccessRedisJsonhal := *openapiclient.NewEnvironmentAccessRedisJsonhal() // EnvironmentAccessRedisJsonhal | The new EnvironmentAccessRedis resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisPost(context.Background()).EnvironmentAccessRedisJsonhal(environmentAccessRedisJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRedisPost`: EnvironmentAccessRedisJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRedisApi.ApiEnvironmentAccessRedisPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRedisPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessRedisJsonhal** | [**EnvironmentAccessRedisJsonhal**](EnvironmentAccessRedisJsonhal.md) | The new EnvironmentAccessRedis resource | 

### Return type

[**EnvironmentAccessRedisJsonhal**](EnvironmentAccessRedisJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

