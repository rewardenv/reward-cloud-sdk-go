# \DataTransferTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDataTransferTypesGetCollection**](DataTransferTypeApi.md#ApiDataTransferTypesGetCollection) | **Get** /api/data_transfer_types | Retrieves the collection of DataTransferType resources.
[**ApiDataTransferTypesIdDelete**](DataTransferTypeApi.md#ApiDataTransferTypesIdDelete) | **Delete** /api/data_transfer_types/{id} | Removes the DataTransferType resource.
[**ApiDataTransferTypesIdGet**](DataTransferTypeApi.md#ApiDataTransferTypesIdGet) | **Get** /api/data_transfer_types/{id} | Retrieves a DataTransferType resource.
[**ApiDataTransferTypesIdPatch**](DataTransferTypeApi.md#ApiDataTransferTypesIdPatch) | **Patch** /api/data_transfer_types/{id} | Updates the DataTransferType resource.
[**ApiDataTransferTypesIdPut**](DataTransferTypeApi.md#ApiDataTransferTypesIdPut) | **Put** /api/data_transfer_types/{id} | Replaces the DataTransferType resource.
[**ApiDataTransferTypesPost**](DataTransferTypeApi.md#ApiDataTransferTypesPost) | **Post** /api/data_transfer_types | Creates a DataTransferType resource.



## ApiDataTransferTypesGetCollection

> []DataTransferType ApiDataTransferTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of DataTransferType resources.



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
    resp, r, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferTypeApi.ApiDataTransferTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferTypesGetCollection`: []DataTransferType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferTypeApi.ApiDataTransferTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]DataTransferType**](DataTransferType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferTypesIdDelete

> ApiDataTransferTypesIdDelete(ctx, id).Execute()

Removes the DataTransferType resource.



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
    id := "id_example" // string | DataTransferType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferTypeApi.ApiDataTransferTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferTypesIdDeleteRequest struct via the builder pattern


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


## ApiDataTransferTypesIdGet

> DataTransferType ApiDataTransferTypesIdGet(ctx, id).Execute()

Retrieves a DataTransferType resource.



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
    id := "id_example" // string | DataTransferType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferTypeApi.ApiDataTransferTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferTypesIdGet`: DataTransferType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferTypeApi.ApiDataTransferTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataTransferType**](DataTransferType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferTypesIdPatch

> DataTransferType ApiDataTransferTypesIdPatch(ctx, id).DataTransferType(dataTransferType).Execute()

Updates the DataTransferType resource.



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
    id := "id_example" // string | DataTransferType identifier
    dataTransferType := *openapiclient.NewDataTransferType() // DataTransferType | The updated DataTransferType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdPatch(context.Background(), id).DataTransferType(dataTransferType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferTypeApi.ApiDataTransferTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferTypesIdPatch`: DataTransferType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferTypeApi.ApiDataTransferTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTransferType** | [**DataTransferType**](DataTransferType.md) | The updated DataTransferType resource | 

### Return type

[**DataTransferType**](DataTransferType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferTypesIdPut

> DataTransferType ApiDataTransferTypesIdPut(ctx, id).DataTransferType(dataTransferType).Execute()

Replaces the DataTransferType resource.



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
    id := "id_example" // string | DataTransferType identifier
    dataTransferType := *openapiclient.NewDataTransferType() // DataTransferType | The updated DataTransferType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesIdPut(context.Background(), id).DataTransferType(dataTransferType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferTypeApi.ApiDataTransferTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferTypesIdPut`: DataTransferType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferTypeApi.ApiDataTransferTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTransferType** | [**DataTransferType**](DataTransferType.md) | The updated DataTransferType resource | 

### Return type

[**DataTransferType**](DataTransferType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferTypesPost

> DataTransferType ApiDataTransferTypesPost(ctx).DataTransferType(dataTransferType).Execute()

Creates a DataTransferType resource.



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
    dataTransferType := *openapiclient.NewDataTransferType() // DataTransferType | The new DataTransferType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferTypeApi.ApiDataTransferTypesPost(context.Background()).DataTransferType(dataTransferType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferTypeApi.ApiDataTransferTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferTypesPost`: DataTransferType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferTypeApi.ApiDataTransferTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataTransferType** | [**DataTransferType**](DataTransferType.md) | The new DataTransferType resource | 

### Return type

[**DataTransferType**](DataTransferType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

