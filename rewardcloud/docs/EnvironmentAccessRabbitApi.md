# \EnvironmentAccessRabbitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessRabbitsGetCollection**](EnvironmentAccessRabbitApi.md#ApiEnvironmentAccessRabbitsGetCollection) | **Get** /api/environment_access_rabbits | Retrieves the collection of EnvironmentAccessRabbit resources.
[**ApiEnvironmentAccessRabbitsIdDelete**](EnvironmentAccessRabbitApi.md#ApiEnvironmentAccessRabbitsIdDelete) | **Delete** /api/environment_access_rabbits/{id} | Removes the EnvironmentAccessRabbit resource.
[**ApiEnvironmentAccessRabbitsIdGet**](EnvironmentAccessRabbitApi.md#ApiEnvironmentAccessRabbitsIdGet) | **Get** /api/environment_access_rabbits/{id} | Retrieves a EnvironmentAccessRabbit resource.
[**ApiEnvironmentAccessRabbitsIdPatch**](EnvironmentAccessRabbitApi.md#ApiEnvironmentAccessRabbitsIdPatch) | **Patch** /api/environment_access_rabbits/{id} | Updates the EnvironmentAccessRabbit resource.
[**ApiEnvironmentAccessRabbitsIdPut**](EnvironmentAccessRabbitApi.md#ApiEnvironmentAccessRabbitsIdPut) | **Put** /api/environment_access_rabbits/{id} | Replaces the EnvironmentAccessRabbit resource.
[**ApiEnvironmentAccessRabbitsPost**](EnvironmentAccessRabbitApi.md#ApiEnvironmentAccessRabbitsPost) | **Post** /api/environment_access_rabbits | Creates a EnvironmentAccessRabbit resource.



## ApiEnvironmentAccessRabbitsGetCollection

> []EnvironmentAccessRabbit ApiEnvironmentAccessRabbitsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvironmentAccessRabbit resources.



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
    resp, r, err := apiClient.EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRabbitsGetCollection`: []EnvironmentAccessRabbit
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRabbitsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRabbitsIdDelete

> ApiEnvironmentAccessRabbitsIdDelete(ctx, id).Execute()

Removes the EnvironmentAccessRabbit resource.



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
    id := "id_example" // string | EnvironmentAccessRabbit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRabbit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRabbitsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessRabbitsIdGet

> EnvironmentAccessRabbit ApiEnvironmentAccessRabbitsIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccessRabbit resource.



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
    id := "id_example" // string | EnvironmentAccessRabbit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRabbitsIdGet`: EnvironmentAccessRabbit
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRabbit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRabbitsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRabbitsIdPatch

> EnvironmentAccessRabbit ApiEnvironmentAccessRabbitsIdPatch(ctx, id).EnvironmentAccessRabbit(environmentAccessRabbit).Execute()

Updates the EnvironmentAccessRabbit resource.



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
    id := "id_example" // string | EnvironmentAccessRabbit identifier
    environmentAccessRabbit := *openapiclient.NewEnvironmentAccessRabbit() // EnvironmentAccessRabbit | The updated EnvironmentAccessRabbit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdPatch(context.Background(), id).EnvironmentAccessRabbit(environmentAccessRabbit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRabbitsIdPatch`: EnvironmentAccessRabbit
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRabbit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRabbitsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessRabbit** | [**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md) | The updated EnvironmentAccessRabbit resource | 

### Return type

[**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRabbitsIdPut

> EnvironmentAccessRabbit ApiEnvironmentAccessRabbitsIdPut(ctx, id).EnvironmentAccessRabbit(environmentAccessRabbit).Execute()

Replaces the EnvironmentAccessRabbit resource.



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
    id := "id_example" // string | EnvironmentAccessRabbit identifier
    environmentAccessRabbit := *openapiclient.NewEnvironmentAccessRabbit() // EnvironmentAccessRabbit | The updated EnvironmentAccessRabbit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdPut(context.Background(), id).EnvironmentAccessRabbit(environmentAccessRabbit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRabbitsIdPut`: EnvironmentAccessRabbit
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessRabbit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRabbitsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessRabbit** | [**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md) | The updated EnvironmentAccessRabbit resource | 

### Return type

[**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessRabbitsPost

> EnvironmentAccessRabbit ApiEnvironmentAccessRabbitsPost(ctx).EnvironmentAccessRabbit(environmentAccessRabbit).Execute()

Creates a EnvironmentAccessRabbit resource.



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
    environmentAccessRabbit := *openapiclient.NewEnvironmentAccessRabbit() // EnvironmentAccessRabbit | The new EnvironmentAccessRabbit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsPost(context.Background()).EnvironmentAccessRabbit(environmentAccessRabbit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessRabbitsPost`: EnvironmentAccessRabbit
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessRabbitApi.ApiEnvironmentAccessRabbitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessRabbitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessRabbit** | [**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md) | The new EnvironmentAccessRabbit resource | 

### Return type

[**EnvironmentAccessRabbit**](EnvironmentAccessRabbit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

