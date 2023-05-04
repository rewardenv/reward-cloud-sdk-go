# \ProviderApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProvidersGetCollection**](ProviderApi.md#ApiProvidersGetCollection) | **Get** /api/providers | Retrieves the collection of Provider resources.
[**ApiProvidersIdDelete**](ProviderApi.md#ApiProvidersIdDelete) | **Delete** /api/providers/{id} | Removes the Provider resource.
[**ApiProvidersIdGet**](ProviderApi.md#ApiProvidersIdGet) | **Get** /api/providers/{id} | Retrieves a Provider resource.
[**ApiProvidersIdPatch**](ProviderApi.md#ApiProvidersIdPatch) | **Patch** /api/providers/{id} | Updates the Provider resource.
[**ApiProvidersIdPut**](ProviderApi.md#ApiProvidersIdPut) | **Put** /api/providers/{id} | Replaces the Provider resource.
[**ApiProvidersPost**](ProviderApi.md#ApiProvidersPost) | **Post** /api/providers | Creates a Provider resource.



## ApiProvidersGetCollection

> ApiProvidersGetCollection200Response ApiProvidersGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of Provider resources.



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
    resp, r, err := apiClient.ProviderApi.ApiProvidersGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ApiProvidersGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProvidersGetCollection`: ApiProvidersGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ApiProvidersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProvidersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**ApiProvidersGetCollection200Response**](ApiProvidersGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProvidersIdDelete

> ApiProvidersIdDelete(ctx, id).Execute()

Removes the Provider resource.



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
    id := "id_example" // string | Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProviderApi.ApiProvidersIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ApiProvidersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProvidersIdDeleteRequest struct via the builder pattern


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


## ApiProvidersIdGet

> ProviderJsonhal ApiProvidersIdGet(ctx, id).Execute()

Retrieves a Provider resource.



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
    id := "id_example" // string | Provider identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProviderApi.ApiProvidersIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ApiProvidersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProvidersIdGet`: ProviderJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ApiProvidersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProvidersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProviderJsonhal**](ProviderJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProvidersIdPatch

> ProviderJsonhal ApiProvidersIdPatch(ctx, id).Provider(provider).Execute()

Updates the Provider resource.



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
    id := "id_example" // string | Provider identifier
    provider := *openapiclient.NewProvider() // Provider | The updated Provider resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProviderApi.ApiProvidersIdPatch(context.Background(), id).Provider(provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ApiProvidersIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProvidersIdPatch`: ProviderJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ApiProvidersIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProvidersIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provider** | [**Provider**](Provider.md) | The updated Provider resource | 

### Return type

[**ProviderJsonhal**](ProviderJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProvidersIdPut

> ProviderJsonhal ApiProvidersIdPut(ctx, id).ProviderJsonhal(providerJsonhal).Execute()

Replaces the Provider resource.



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
    id := "id_example" // string | Provider identifier
    providerJsonhal := *openapiclient.NewProviderJsonhal() // ProviderJsonhal | The updated Provider resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProviderApi.ApiProvidersIdPut(context.Background(), id).ProviderJsonhal(providerJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ApiProvidersIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProvidersIdPut`: ProviderJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ApiProvidersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Provider identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProvidersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **providerJsonhal** | [**ProviderJsonhal**](ProviderJsonhal.md) | The updated Provider resource | 

### Return type

[**ProviderJsonhal**](ProviderJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProvidersPost

> ProviderJsonhal ApiProvidersPost(ctx).ProviderJsonhal(providerJsonhal).Execute()

Creates a Provider resource.



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
    providerJsonhal := *openapiclient.NewProviderJsonhal() // ProviderJsonhal | The new Provider resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProviderApi.ApiProvidersPost(context.Background()).ProviderJsonhal(providerJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProviderApi.ApiProvidersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProvidersPost`: ProviderJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ProviderApi.ApiProvidersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerJsonhal** | [**ProviderJsonhal**](ProviderJsonhal.md) | The new Provider resource | 

### Return type

[**ProviderJsonhal**](ProviderJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

