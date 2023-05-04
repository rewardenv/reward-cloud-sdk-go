# \PriceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiPricesGetCollection**](PriceApi.md#ApiPricesGetCollection) | **Get** /api/prices | Retrieves the collection of Price resources.
[**ApiPricesIdDelete**](PriceApi.md#ApiPricesIdDelete) | **Delete** /api/prices/{id} | Removes the Price resource.
[**ApiPricesIdGet**](PriceApi.md#ApiPricesIdGet) | **Get** /api/prices/{id} | Retrieves a Price resource.
[**ApiPricesIdPatch**](PriceApi.md#ApiPricesIdPatch) | **Patch** /api/prices/{id} | Updates the Price resource.
[**ApiPricesIdPut**](PriceApi.md#ApiPricesIdPut) | **Put** /api/prices/{id} | Replaces the Price resource.
[**ApiPricesPost**](PriceApi.md#ApiPricesPost) | **Post** /api/prices | Creates a Price resource.



## ApiPricesGetCollection

> ApiPricesGetCollection200Response ApiPricesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Product(product).Product2(product2).Execute()

Retrieves the collection of Price resources.



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
    product := "product_example" // string |  (optional)
    product2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceApi.ApiPricesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Product(product).Product2(product2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceApi.ApiPricesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPricesGetCollection`: ApiPricesGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceApi.ApiPricesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPricesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **product** | **string** |  | 
 **product2** | **[]string** |  | 

### Return type

[**ApiPricesGetCollection200Response**](ApiPricesGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPricesIdDelete

> ApiPricesIdDelete(ctx, id).Execute()

Removes the Price resource.



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
    id := "id_example" // string | Price identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceApi.ApiPricesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceApi.ApiPricesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPricesIdDeleteRequest struct via the builder pattern


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


## ApiPricesIdGet

> PriceJsonhal ApiPricesIdGet(ctx, id).Execute()

Retrieves a Price resource.



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
    id := "id_example" // string | Price identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceApi.ApiPricesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceApi.ApiPricesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPricesIdGet`: PriceJsonhal
    fmt.Fprintf(os.Stdout, "Response from `PriceApi.ApiPricesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPricesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PriceJsonhal**](PriceJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPricesIdPatch

> PriceJsonhal ApiPricesIdPatch(ctx, id).Price(price).Execute()

Updates the Price resource.



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
    id := "id_example" // string | Price identifier
    price := *openapiclient.NewPrice() // Price | The updated Price resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceApi.ApiPricesIdPatch(context.Background(), id).Price(price).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceApi.ApiPricesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPricesIdPatch`: PriceJsonhal
    fmt.Fprintf(os.Stdout, "Response from `PriceApi.ApiPricesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPricesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **price** | [**Price**](Price.md) | The updated Price resource | 

### Return type

[**PriceJsonhal**](PriceJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPricesIdPut

> PriceJsonhal ApiPricesIdPut(ctx, id).PriceJsonhal(priceJsonhal).Execute()

Replaces the Price resource.



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
    id := "id_example" // string | Price identifier
    priceJsonhal := *openapiclient.NewPriceJsonhal() // PriceJsonhal | The updated Price resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceApi.ApiPricesIdPut(context.Background(), id).PriceJsonhal(priceJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceApi.ApiPricesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPricesIdPut`: PriceJsonhal
    fmt.Fprintf(os.Stdout, "Response from `PriceApi.ApiPricesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiPricesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **priceJsonhal** | [**PriceJsonhal**](PriceJsonhal.md) | The updated Price resource | 

### Return type

[**PriceJsonhal**](PriceJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiPricesPost

> PriceJsonhal ApiPricesPost(ctx).PriceJsonhal(priceJsonhal).Execute()

Creates a Price resource.



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
    priceJsonhal := *openapiclient.NewPriceJsonhal() // PriceJsonhal | The new Price resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceApi.ApiPricesPost(context.Background()).PriceJsonhal(priceJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceApi.ApiPricesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiPricesPost`: PriceJsonhal
    fmt.Fprintf(os.Stdout, "Response from `PriceApi.ApiPricesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiPricesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceJsonhal** | [**PriceJsonhal**](PriceJsonhal.md) | The new Price resource | 

### Return type

[**PriceJsonhal**](PriceJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

