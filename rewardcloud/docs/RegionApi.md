# \RegionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRegionsGetCollection**](RegionApi.md#ApiRegionsGetCollection) | **Get** /api/regions | Retrieves the collection of Region resources.
[**ApiRegionsIdDelete**](RegionApi.md#ApiRegionsIdDelete) | **Delete** /api/regions/{id} | Removes the Region resource.
[**ApiRegionsIdGet**](RegionApi.md#ApiRegionsIdGet) | **Get** /api/regions/{id} | Retrieves a Region resource.
[**ApiRegionsIdPatch**](RegionApi.md#ApiRegionsIdPatch) | **Patch** /api/regions/{id} | Updates the Region resource.
[**ApiRegionsIdPut**](RegionApi.md#ApiRegionsIdPut) | **Put** /api/regions/{id} | Replaces the Region resource.
[**ApiRegionsIdgetWithProviderGet**](RegionApi.md#ApiRegionsIdgetWithProviderGet) | **Get** /api/regions/{id}/get-with-provider | Retrieves a Region resource.
[**ApiRegionsPost**](RegionApi.md#ApiRegionsPost) | **Post** /api/regions | Creates a Region resource.
[**ApiRegionsgetWithProviderGet**](RegionApi.md#ApiRegionsgetWithProviderGet) | **Get** /api/regions/get-with-provider | Retrieves a Region resource.



## ApiRegionsGetCollection

> []Region ApiRegionsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Provider(provider).Provider2(provider2).Execute()

Retrieves the collection of Region resources.



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
    provider := "provider_example" // string |  (optional)
    provider2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Provider(provider).Provider2(provider2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsGetCollection`: []Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **provider** | **string** |  | 
 **provider2** | **[]string** |  | 

### Return type

[**[]Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegionsIdDelete

> ApiRegionsIdDelete(ctx, id).Execute()

Removes the Region resource.



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
    id := "id_example" // string | Region identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Region identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsIdDeleteRequest struct via the builder pattern


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


## ApiRegionsIdGet

> Region ApiRegionsIdGet(ctx, id).Execute()

Retrieves a Region resource.



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
    id := "id_example" // string | Region identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsIdGet`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Region identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegionsIdPatch

> Region ApiRegionsIdPatch(ctx, id).Region(region).Execute()

Updates the Region resource.



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
    id := "id_example" // string | Region identifier
    region := *openapiclient.NewRegion() // Region | The updated Region resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsIdPatch(context.Background(), id).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsIdPatch`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Region identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | [**Region**](Region.md) | The updated Region resource | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegionsIdPut

> Region ApiRegionsIdPut(ctx, id).Region(region).Execute()

Replaces the Region resource.



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
    id := "id_example" // string | Region identifier
    region := *openapiclient.NewRegion() // Region | The updated Region resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsIdPut(context.Background(), id).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsIdPut`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Region identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | [**Region**](Region.md) | The updated Region resource | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegionsIdgetWithProviderGet

> Region ApiRegionsIdgetWithProviderGet(ctx, id).Execute()

Retrieves a Region resource.



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
    id := "id_example" // string | Region identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsIdgetWithProviderGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsIdgetWithProviderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsIdgetWithProviderGet`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsIdgetWithProviderGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Region identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsIdgetWithProviderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegionsPost

> Region ApiRegionsPost(ctx).Region(region).Execute()

Creates a Region resource.



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
    region := *openapiclient.NewRegion() // Region | The new Region resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsPost(context.Background()).Region(region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsPost`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**Region**](Region.md) | The new Region resource | 

### Return type

[**Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRegionsgetWithProviderGet

> Region ApiRegionsgetWithProviderGet(ctx).Execute()

Retrieves a Region resource.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegionApi.ApiRegionsgetWithProviderGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegionApi.ApiRegionsgetWithProviderGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRegionsgetWithProviderGet`: Region
    fmt.Fprintf(os.Stdout, "Response from `RegionApi.ApiRegionsgetWithProviderGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRegionsgetWithProviderGetRequest struct via the builder pattern


### Return type

[**Region**](Region.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

