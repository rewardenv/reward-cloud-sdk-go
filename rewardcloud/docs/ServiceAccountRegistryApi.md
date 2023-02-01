# \ServiceAccountRegistryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServiceAccountRegistriesGetCollection**](ServiceAccountRegistryApi.md#ApiServiceAccountRegistriesGetCollection) | **Get** /api/service_account_registries | Retrieves the collection of ServiceAccountRegistry resources.
[**ApiServiceAccountRegistriesIdDelete**](ServiceAccountRegistryApi.md#ApiServiceAccountRegistriesIdDelete) | **Delete** /api/service_account_registries/{id} | Removes the ServiceAccountRegistry resource.
[**ApiServiceAccountRegistriesIdGet**](ServiceAccountRegistryApi.md#ApiServiceAccountRegistriesIdGet) | **Get** /api/service_account_registries/{id} | Retrieves a ServiceAccountRegistry resource.
[**ApiServiceAccountRegistriesIdPatch**](ServiceAccountRegistryApi.md#ApiServiceAccountRegistriesIdPatch) | **Patch** /api/service_account_registries/{id} | Updates the ServiceAccountRegistry resource.
[**ApiServiceAccountRegistriesIdPut**](ServiceAccountRegistryApi.md#ApiServiceAccountRegistriesIdPut) | **Put** /api/service_account_registries/{id} | Replaces the ServiceAccountRegistry resource.
[**ApiServiceAccountRegistriesPost**](ServiceAccountRegistryApi.md#ApiServiceAccountRegistriesPost) | **Post** /api/service_account_registries | Creates a ServiceAccountRegistry resource.



## ApiServiceAccountRegistriesGetCollection

> []ServiceAccountRegistry ApiServiceAccountRegistriesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ServiceAccount(serviceAccount).ServiceAccount2(serviceAccount2).Execute()

Retrieves the collection of ServiceAccountRegistry resources.



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
    serviceAccount := "serviceAccount_example" // string |  (optional)
    serviceAccount2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ServiceAccount(serviceAccount).ServiceAccount2(serviceAccount2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountRegistryApi.ApiServiceAccountRegistriesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountRegistriesGetCollection`: []ServiceAccountRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountRegistryApi.ApiServiceAccountRegistriesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountRegistriesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **serviceAccount** | **string** |  | 
 **serviceAccount2** | **[]string** |  | 

### Return type

[**[]ServiceAccountRegistry**](ServiceAccountRegistry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountRegistriesIdDelete

> ApiServiceAccountRegistriesIdDelete(ctx, id).Execute()

Removes the ServiceAccountRegistry resource.



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
    id := "id_example" // string | ServiceAccountRegistry identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountRegistry identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountRegistriesIdDeleteRequest struct via the builder pattern


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


## ApiServiceAccountRegistriesIdGet

> ServiceAccountRegistry ApiServiceAccountRegistriesIdGet(ctx, id).Execute()

Retrieves a ServiceAccountRegistry resource.



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
    id := "id_example" // string | ServiceAccountRegistry identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountRegistriesIdGet`: ServiceAccountRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountRegistry identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountRegistriesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccountRegistry**](ServiceAccountRegistry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountRegistriesIdPatch

> ServiceAccountRegistry ApiServiceAccountRegistriesIdPatch(ctx, id).ServiceAccountRegistry(serviceAccountRegistry).Execute()

Updates the ServiceAccountRegistry resource.



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
    id := "id_example" // string | ServiceAccountRegistry identifier
    serviceAccountRegistry := *openapiclient.NewServiceAccountRegistry() // ServiceAccountRegistry | The updated ServiceAccountRegistry resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPatch(context.Background(), id).ServiceAccountRegistry(serviceAccountRegistry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountRegistriesIdPatch`: ServiceAccountRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountRegistry identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountRegistriesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountRegistry** | [**ServiceAccountRegistry**](ServiceAccountRegistry.md) | The updated ServiceAccountRegistry resource | 

### Return type

[**ServiceAccountRegistry**](ServiceAccountRegistry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountRegistriesIdPut

> ServiceAccountRegistry ApiServiceAccountRegistriesIdPut(ctx, id).ServiceAccountRegistry(serviceAccountRegistry).Execute()

Replaces the ServiceAccountRegistry resource.



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
    id := "id_example" // string | ServiceAccountRegistry identifier
    serviceAccountRegistry := *openapiclient.NewServiceAccountRegistry() // ServiceAccountRegistry | The updated ServiceAccountRegistry resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPut(context.Background(), id).ServiceAccountRegistry(serviceAccountRegistry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountRegistriesIdPut`: ServiceAccountRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountRegistryApi.ApiServiceAccountRegistriesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountRegistry identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountRegistriesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountRegistry** | [**ServiceAccountRegistry**](ServiceAccountRegistry.md) | The updated ServiceAccountRegistry resource | 

### Return type

[**ServiceAccountRegistry**](ServiceAccountRegistry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountRegistriesPost

> ServiceAccountRegistry ApiServiceAccountRegistriesPost(ctx).ServiceAccountRegistry(serviceAccountRegistry).Execute()

Creates a ServiceAccountRegistry resource.



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
    serviceAccountRegistry := *openapiclient.NewServiceAccountRegistry() // ServiceAccountRegistry | The new ServiceAccountRegistry resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountRegistryApi.ApiServiceAccountRegistriesPost(context.Background()).ServiceAccountRegistry(serviceAccountRegistry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountRegistryApi.ApiServiceAccountRegistriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountRegistriesPost`: ServiceAccountRegistry
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountRegistryApi.ApiServiceAccountRegistriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountRegistriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountRegistry** | [**ServiceAccountRegistry**](ServiceAccountRegistry.md) | The new ServiceAccountRegistry resource | 

### Return type

[**ServiceAccountRegistry**](ServiceAccountRegistry.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

