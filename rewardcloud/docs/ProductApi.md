# \ProductApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiProductsGetCollection**](ProductApi.md#ApiProductsGetCollection) | **Get** /api/products | Retrieves the collection of Product resources.
[**ApiProductsIdDelete**](ProductApi.md#ApiProductsIdDelete) | **Delete** /api/products/{id} | Removes the Product resource.
[**ApiProductsIdGet**](ProductApi.md#ApiProductsIdGet) | **Get** /api/products/{id} | Retrieves a Product resource.
[**ApiProductsIdPatch**](ProductApi.md#ApiProductsIdPatch) | **Patch** /api/products/{id} | Updates the Product resource.
[**ApiProductsIdPut**](ProductApi.md#ApiProductsIdPut) | **Put** /api/products/{id} | Replaces the Product resource.
[**ApiProductsPost**](ProductApi.md#ApiProductsPost) | **Post** /api/products | Creates a Product resource.



## ApiProductsGetCollection

> []ProductProductOutput ApiProductsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).DefaultPrice(defaultPrice).DefaultPrice2(defaultPrice2).Region(region).Region2(region2).Execute()

Retrieves the collection of Product resources.



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
    defaultPrice := "defaultPrice_example" // string |  (optional)
    defaultPrice2 := []string{"Inner_example"} // []string |  (optional)
    region := "region_example" // string |  (optional)
    region2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ApiProductsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).DefaultPrice(defaultPrice).DefaultPrice2(defaultPrice2).Region(region).Region2(region2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ApiProductsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductsGetCollection`: []ProductProductOutput
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ApiProductsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **defaultPrice** | **string** |  | 
 **defaultPrice2** | **[]string** |  | 
 **region** | **string** |  | 
 **region2** | **[]string** |  | 

### Return type

[**[]ProductProductOutput**](ProductProductOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProductsIdDelete

> ApiProductsIdDelete(ctx, id).Execute()

Removes the Product resource.



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
    id := "id_example" // string | Product identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ApiProductsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ApiProductsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Product identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsIdDeleteRequest struct via the builder pattern


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


## ApiProductsIdGet

> ProductProductOutput ApiProductsIdGet(ctx, id).Execute()

Retrieves a Product resource.



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
    id := "id_example" // string | Product identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ApiProductsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ApiProductsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductsIdGet`: ProductProductOutput
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ApiProductsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Product identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductProductOutput**](ProductProductOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProductsIdPatch

> ProductProductOutput ApiProductsIdPatch(ctx, id).ProductProductInput(productProductInput).Execute()

Updates the Product resource.



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
    id := "id_example" // string | Product identifier
    productProductInput := *openapiclient.NewProductProductInput() // ProductProductInput | The updated Product resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ApiProductsIdPatch(context.Background(), id).ProductProductInput(productProductInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ApiProductsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductsIdPatch`: ProductProductOutput
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ApiProductsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Product identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productProductInput** | [**ProductProductInput**](ProductProductInput.md) | The updated Product resource | 

### Return type

[**ProductProductOutput**](ProductProductOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProductsIdPut

> ProductProductOutput ApiProductsIdPut(ctx, id).ProductProductInput(productProductInput).Execute()

Replaces the Product resource.



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
    id := "id_example" // string | Product identifier
    productProductInput := *openapiclient.NewProductProductInput() // ProductProductInput | The updated Product resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ApiProductsIdPut(context.Background(), id).ProductProductInput(productProductInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ApiProductsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductsIdPut`: ProductProductOutput
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ApiProductsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Product identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productProductInput** | [**ProductProductInput**](ProductProductInput.md) | The updated Product resource | 

### Return type

[**ProductProductOutput**](ProductProductOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiProductsPost

> ProductProductOutput ApiProductsPost(ctx).ProductProductInput(productProductInput).Execute()

Creates a Product resource.



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
    productProductInput := *openapiclient.NewProductProductInput() // ProductProductInput | The new Product resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductApi.ApiProductsPost(context.Background()).ProductProductInput(productProductInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.ApiProductsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiProductsPost`: ProductProductOutput
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.ApiProductsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiProductsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productProductInput** | [**ProductProductInput**](ProductProductInput.md) | The new Product resource | 

### Return type

[**ProductProductOutput**](ProductProductOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

