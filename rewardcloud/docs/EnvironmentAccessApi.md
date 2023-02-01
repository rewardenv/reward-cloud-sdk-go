# \EnvironmentAccessApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessesGetCollection**](EnvironmentAccessApi.md#ApiEnvironmentAccessesGetCollection) | **Get** /api/environment_accesses | Retrieves the collection of EnvironmentAccess resources.
[**ApiEnvironmentAccessesIdDelete**](EnvironmentAccessApi.md#ApiEnvironmentAccessesIdDelete) | **Delete** /api/environment_accesses/{id} | Removes the EnvironmentAccess resource.
[**ApiEnvironmentAccessesIdGet**](EnvironmentAccessApi.md#ApiEnvironmentAccessesIdGet) | **Get** /api/environment_accesses/{id} | Retrieves a EnvironmentAccess resource.
[**ApiEnvironmentAccessesIdPatch**](EnvironmentAccessApi.md#ApiEnvironmentAccessesIdPatch) | **Patch** /api/environment_accesses/{id} | Updates the EnvironmentAccess resource.
[**ApiEnvironmentAccessesIdPut**](EnvironmentAccessApi.md#ApiEnvironmentAccessesIdPut) | **Put** /api/environment_accesses/{id} | Replaces the EnvironmentAccess resource.
[**ApiEnvironmentAccessesIdallDataGet**](EnvironmentAccessApi.md#ApiEnvironmentAccessesIdallDataGet) | **Get** /api/environment_accesses/{id}/all-data | Retrieves a EnvironmentAccess resource.
[**ApiEnvironmentAccessesPost**](EnvironmentAccessApi.md#ApiEnvironmentAccessesPost) | **Post** /api/environment_accesses | Creates a EnvironmentAccess resource.



## ApiEnvironmentAccessesGetCollection

> []EnvironmentAccess ApiEnvironmentAccessesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Environment(environment).Environment2(environment2).Frontend(frontend).Frontend2(frontend2).Backend(backend).Backend2(backend2).Database(database).Database2(database2).DevTools(devTools).DevTools2(devTools2).Redis(redis).Redis2(redis2).Rabbit(rabbit).Rabbit2(rabbit2).Execute()

Retrieves the collection of EnvironmentAccess resources.



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
    environment := "environment_example" // string |  (optional)
    environment2 := []string{"Inner_example"} // []string |  (optional)
    frontend := "frontend_example" // string |  (optional)
    frontend2 := []string{"Inner_example"} // []string |  (optional)
    backend := "backend_example" // string |  (optional)
    backend2 := []string{"Inner_example"} // []string |  (optional)
    database := "database_example" // string |  (optional)
    database2 := []string{"Inner_example"} // []string |  (optional)
    devTools := "devTools_example" // string |  (optional)
    devTools2 := []string{"Inner_example"} // []string |  (optional)
    redis := "redis_example" // string |  (optional)
    redis2 := []string{"Inner_example"} // []string |  (optional)
    rabbit := "rabbit_example" // string |  (optional)
    rabbit2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Environment(environment).Environment2(environment2).Frontend(frontend).Frontend2(frontend2).Backend(backend).Backend2(backend2).Database(database).Database2(database2).DevTools(devTools).DevTools2(devTools2).Redis(redis).Redis2(redis2).Rabbit(rabbit).Rabbit2(rabbit2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessesGetCollection`: []EnvironmentAccess
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessApi.ApiEnvironmentAccessesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **environment** | **string** |  | 
 **environment2** | **[]string** |  | 
 **frontend** | **string** |  | 
 **frontend2** | **[]string** |  | 
 **backend** | **string** |  | 
 **backend2** | **[]string** |  | 
 **database** | **string** |  | 
 **database2** | **[]string** |  | 
 **devTools** | **string** |  | 
 **devTools2** | **[]string** |  | 
 **redis** | **string** |  | 
 **redis2** | **[]string** |  | 
 **rabbit** | **string** |  | 
 **rabbit2** | **[]string** |  | 

### Return type

[**[]EnvironmentAccess**](EnvironmentAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessesIdDelete

> ApiEnvironmentAccessesIdDelete(ctx, id).Execute()

Removes the EnvironmentAccess resource.



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
    id := "id_example" // string | EnvironmentAccess identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccess identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessesIdGet

> EnvironmentAccess ApiEnvironmentAccessesIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccess resource.



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
    id := "id_example" // string | EnvironmentAccess identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessesIdGet`: EnvironmentAccess
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessApi.ApiEnvironmentAccessesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccess identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccess**](EnvironmentAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessesIdPatch

> EnvironmentAccess ApiEnvironmentAccessesIdPatch(ctx, id).EnvironmentAccess(environmentAccess).Execute()

Updates the EnvironmentAccess resource.



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
    id := "id_example" // string | EnvironmentAccess identifier
    environmentAccess := *openapiclient.NewEnvironmentAccess() // EnvironmentAccess | The updated EnvironmentAccess resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesIdPatch(context.Background(), id).EnvironmentAccess(environmentAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessesIdPatch`: EnvironmentAccess
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessApi.ApiEnvironmentAccessesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccess identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccess** | [**EnvironmentAccess**](EnvironmentAccess.md) | The updated EnvironmentAccess resource | 

### Return type

[**EnvironmentAccess**](EnvironmentAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessesIdPut

> EnvironmentAccess ApiEnvironmentAccessesIdPut(ctx, id).EnvironmentAccess(environmentAccess).Execute()

Replaces the EnvironmentAccess resource.



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
    id := "id_example" // string | EnvironmentAccess identifier
    environmentAccess := *openapiclient.NewEnvironmentAccess() // EnvironmentAccess | The updated EnvironmentAccess resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesIdPut(context.Background(), id).EnvironmentAccess(environmentAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessesIdPut`: EnvironmentAccess
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessApi.ApiEnvironmentAccessesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccess identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccess** | [**EnvironmentAccess**](EnvironmentAccess.md) | The updated EnvironmentAccess resource | 

### Return type

[**EnvironmentAccess**](EnvironmentAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessesIdallDataGet

> EnvironmentAccess ApiEnvironmentAccessesIdallDataGet(ctx, id).Execute()

Retrieves a EnvironmentAccess resource.



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
    id := "id_example" // string | EnvironmentAccess identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesIdallDataGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesIdallDataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessesIdallDataGet`: EnvironmentAccess
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessApi.ApiEnvironmentAccessesIdallDataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccess identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesIdallDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccess**](EnvironmentAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessesPost

> EnvironmentAccess ApiEnvironmentAccessesPost(ctx).EnvironmentAccess(environmentAccess).Execute()

Creates a EnvironmentAccess resource.



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
    environmentAccess := *openapiclient.NewEnvironmentAccess() // EnvironmentAccess | The new EnvironmentAccess resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessApi.ApiEnvironmentAccessesPost(context.Background()).EnvironmentAccess(environmentAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessApi.ApiEnvironmentAccessesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessesPost`: EnvironmentAccess
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessApi.ApiEnvironmentAccessesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccess** | [**EnvironmentAccess**](EnvironmentAccess.md) | The new EnvironmentAccess resource | 

### Return type

[**EnvironmentAccess**](EnvironmentAccess.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

