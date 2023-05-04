# \EnvironmentAccessDatabaseApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessDatabasesGetCollection**](EnvironmentAccessDatabaseApi.md#ApiEnvironmentAccessDatabasesGetCollection) | **Get** /api/environment_access_databases | Retrieves the collection of EnvironmentAccessDatabase resources.
[**ApiEnvironmentAccessDatabasesIdDelete**](EnvironmentAccessDatabaseApi.md#ApiEnvironmentAccessDatabasesIdDelete) | **Delete** /api/environment_access_databases/{id} | Removes the EnvironmentAccessDatabase resource.
[**ApiEnvironmentAccessDatabasesIdGet**](EnvironmentAccessDatabaseApi.md#ApiEnvironmentAccessDatabasesIdGet) | **Get** /api/environment_access_databases/{id} | Retrieves a EnvironmentAccessDatabase resource.
[**ApiEnvironmentAccessDatabasesIdPatch**](EnvironmentAccessDatabaseApi.md#ApiEnvironmentAccessDatabasesIdPatch) | **Patch** /api/environment_access_databases/{id} | Updates the EnvironmentAccessDatabase resource.
[**ApiEnvironmentAccessDatabasesIdPut**](EnvironmentAccessDatabaseApi.md#ApiEnvironmentAccessDatabasesIdPut) | **Put** /api/environment_access_databases/{id} | Replaces the EnvironmentAccessDatabase resource.
[**ApiEnvironmentAccessDatabasesPost**](EnvironmentAccessDatabaseApi.md#ApiEnvironmentAccessDatabasesPost) | **Post** /api/environment_access_databases | Creates a EnvironmentAccessDatabase resource.



## ApiEnvironmentAccessDatabasesGetCollection

> ApiEnvironmentAccessDatabasesGetCollection200Response ApiEnvironmentAccessDatabasesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvironmentAccessDatabase resources.



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
    resp, r, err := apiClient.EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDatabasesGetCollection`: ApiEnvironmentAccessDatabasesGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDatabasesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**ApiEnvironmentAccessDatabasesGetCollection200Response**](ApiEnvironmentAccessDatabasesGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDatabasesIdDelete

> ApiEnvironmentAccessDatabasesIdDelete(ctx, id).Execute()

Removes the EnvironmentAccessDatabase resource.



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
    id := "id_example" // string | EnvironmentAccessDatabase identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDatabase identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDatabasesIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessDatabasesIdGet

> EnvironmentAccessDatabaseJsonhal ApiEnvironmentAccessDatabasesIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccessDatabase resource.



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
    id := "id_example" // string | EnvironmentAccessDatabase identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDatabasesIdGet`: EnvironmentAccessDatabaseJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDatabase identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDatabasesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccessDatabaseJsonhal**](EnvironmentAccessDatabaseJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDatabasesIdPatch

> EnvironmentAccessDatabaseJsonhal ApiEnvironmentAccessDatabasesIdPatch(ctx, id).EnvironmentAccessDatabase(environmentAccessDatabase).Execute()

Updates the EnvironmentAccessDatabase resource.



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
    id := "id_example" // string | EnvironmentAccessDatabase identifier
    environmentAccessDatabase := *openapiclient.NewEnvironmentAccessDatabase() // EnvironmentAccessDatabase | The updated EnvironmentAccessDatabase resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdPatch(context.Background(), id).EnvironmentAccessDatabase(environmentAccessDatabase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDatabasesIdPatch`: EnvironmentAccessDatabaseJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDatabase identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDatabasesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessDatabase** | [**EnvironmentAccessDatabase**](EnvironmentAccessDatabase.md) | The updated EnvironmentAccessDatabase resource | 

### Return type

[**EnvironmentAccessDatabaseJsonhal**](EnvironmentAccessDatabaseJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDatabasesIdPut

> EnvironmentAccessDatabaseJsonhal ApiEnvironmentAccessDatabasesIdPut(ctx, id).EnvironmentAccessDatabaseJsonhal(environmentAccessDatabaseJsonhal).Execute()

Replaces the EnvironmentAccessDatabase resource.



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
    id := "id_example" // string | EnvironmentAccessDatabase identifier
    environmentAccessDatabaseJsonhal := *openapiclient.NewEnvironmentAccessDatabaseJsonhal() // EnvironmentAccessDatabaseJsonhal | The updated EnvironmentAccessDatabase resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdPut(context.Background(), id).EnvironmentAccessDatabaseJsonhal(environmentAccessDatabaseJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDatabasesIdPut`: EnvironmentAccessDatabaseJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessDatabase identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDatabasesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentAccessDatabaseJsonhal** | [**EnvironmentAccessDatabaseJsonhal**](EnvironmentAccessDatabaseJsonhal.md) | The updated EnvironmentAccessDatabase resource | 

### Return type

[**EnvironmentAccessDatabaseJsonhal**](EnvironmentAccessDatabaseJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessDatabasesPost

> EnvironmentAccessDatabaseJsonhal ApiEnvironmentAccessDatabasesPost(ctx).EnvironmentAccessDatabaseJsonhal(environmentAccessDatabaseJsonhal).Execute()

Creates a EnvironmentAccessDatabase resource.



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
    environmentAccessDatabaseJsonhal := *openapiclient.NewEnvironmentAccessDatabaseJsonhal() // EnvironmentAccessDatabaseJsonhal | The new EnvironmentAccessDatabase resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesPost(context.Background()).EnvironmentAccessDatabaseJsonhal(environmentAccessDatabaseJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessDatabasesPost`: EnvironmentAccessDatabaseJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessDatabaseApi.ApiEnvironmentAccessDatabasesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessDatabasesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessDatabaseJsonhal** | [**EnvironmentAccessDatabaseJsonhal**](EnvironmentAccessDatabaseJsonhal.md) | The new EnvironmentAccessDatabase resource | 

### Return type

[**EnvironmentAccessDatabaseJsonhal**](EnvironmentAccessDatabaseJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

