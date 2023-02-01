# \DataTransferDataTypeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiDataTransferDataTypesGetCollection**](DataTransferDataTypeApi.md#ApiDataTransferDataTypesGetCollection) | **Get** /api/data_transfer_data_types | Retrieves the collection of DataTransferDataType resources.
[**ApiDataTransferDataTypesIdDelete**](DataTransferDataTypeApi.md#ApiDataTransferDataTypesIdDelete) | **Delete** /api/data_transfer_data_types/{id} | Removes the DataTransferDataType resource.
[**ApiDataTransferDataTypesIdGet**](DataTransferDataTypeApi.md#ApiDataTransferDataTypesIdGet) | **Get** /api/data_transfer_data_types/{id} | Retrieves a DataTransferDataType resource.
[**ApiDataTransferDataTypesIdPatch**](DataTransferDataTypeApi.md#ApiDataTransferDataTypesIdPatch) | **Patch** /api/data_transfer_data_types/{id} | Updates the DataTransferDataType resource.
[**ApiDataTransferDataTypesIdPut**](DataTransferDataTypeApi.md#ApiDataTransferDataTypesIdPut) | **Put** /api/data_transfer_data_types/{id} | Replaces the DataTransferDataType resource.
[**ApiDataTransferDataTypesPost**](DataTransferDataTypeApi.md#ApiDataTransferDataTypesPost) | **Post** /api/data_transfer_data_types | Creates a DataTransferDataType resource.



## ApiDataTransferDataTypesGetCollection

> []DataTransferDataType ApiDataTransferDataTypesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of DataTransferDataType resources.



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
    resp, r, err := apiClient.DataTransferDataTypeApi.ApiDataTransferDataTypesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferDataTypeApi.ApiDataTransferDataTypesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferDataTypesGetCollection`: []DataTransferDataType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferDataTypeApi.ApiDataTransferDataTypesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferDataTypesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]DataTransferDataType**](DataTransferDataType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferDataTypesIdDelete

> ApiDataTransferDataTypesIdDelete(ctx, id).Execute()

Removes the DataTransferDataType resource.



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
    id := "id_example" // string | DataTransferDataType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferDataTypeApi.ApiDataTransferDataTypesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferDataTypeApi.ApiDataTransferDataTypesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferDataType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferDataTypesIdDeleteRequest struct via the builder pattern


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


## ApiDataTransferDataTypesIdGet

> DataTransferDataType ApiDataTransferDataTypesIdGet(ctx, id).Execute()

Retrieves a DataTransferDataType resource.



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
    id := "id_example" // string | DataTransferDataType identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferDataTypeApi.ApiDataTransferDataTypesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferDataTypeApi.ApiDataTransferDataTypesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferDataTypesIdGet`: DataTransferDataType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferDataTypeApi.ApiDataTransferDataTypesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferDataType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferDataTypesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataTransferDataType**](DataTransferDataType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferDataTypesIdPatch

> DataTransferDataType ApiDataTransferDataTypesIdPatch(ctx, id).DataTransferDataType(dataTransferDataType).Execute()

Updates the DataTransferDataType resource.



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
    id := "id_example" // string | DataTransferDataType identifier
    dataTransferDataType := *openapiclient.NewDataTransferDataType() // DataTransferDataType | The updated DataTransferDataType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferDataTypeApi.ApiDataTransferDataTypesIdPatch(context.Background(), id).DataTransferDataType(dataTransferDataType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferDataTypeApi.ApiDataTransferDataTypesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferDataTypesIdPatch`: DataTransferDataType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferDataTypeApi.ApiDataTransferDataTypesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferDataType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferDataTypesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTransferDataType** | [**DataTransferDataType**](DataTransferDataType.md) | The updated DataTransferDataType resource | 

### Return type

[**DataTransferDataType**](DataTransferDataType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferDataTypesIdPut

> DataTransferDataType ApiDataTransferDataTypesIdPut(ctx, id).DataTransferDataType(dataTransferDataType).Execute()

Replaces the DataTransferDataType resource.



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
    id := "id_example" // string | DataTransferDataType identifier
    dataTransferDataType := *openapiclient.NewDataTransferDataType() // DataTransferDataType | The updated DataTransferDataType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferDataTypeApi.ApiDataTransferDataTypesIdPut(context.Background(), id).DataTransferDataType(dataTransferDataType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferDataTypeApi.ApiDataTransferDataTypesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferDataTypesIdPut`: DataTransferDataType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferDataTypeApi.ApiDataTransferDataTypesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | DataTransferDataType identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferDataTypesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataTransferDataType** | [**DataTransferDataType**](DataTransferDataType.md) | The updated DataTransferDataType resource | 

### Return type

[**DataTransferDataType**](DataTransferDataType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiDataTransferDataTypesPost

> DataTransferDataType ApiDataTransferDataTypesPost(ctx).DataTransferDataType(dataTransferDataType).Execute()

Creates a DataTransferDataType resource.



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
    dataTransferDataType := *openapiclient.NewDataTransferDataType() // DataTransferDataType | The new DataTransferDataType resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataTransferDataTypeApi.ApiDataTransferDataTypesPost(context.Background()).DataTransferDataType(dataTransferDataType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataTransferDataTypeApi.ApiDataTransferDataTypesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiDataTransferDataTypesPost`: DataTransferDataType
    fmt.Fprintf(os.Stdout, "Response from `DataTransferDataTypeApi.ApiDataTransferDataTypesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiDataTransferDataTypesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataTransferDataType** | [**DataTransferDataType**](DataTransferDataType.md) | The new DataTransferDataType resource | 

### Return type

[**DataTransferDataType**](DataTransferDataType.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

