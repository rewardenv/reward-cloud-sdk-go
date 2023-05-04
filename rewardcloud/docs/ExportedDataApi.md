# \ExportedDataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiExportedDatasGetCollection**](ExportedDataApi.md#ApiExportedDatasGetCollection) | **Get** /api/exported_datas | Retrieves the collection of ExportedData resources.
[**ApiExportedDatasIdDelete**](ExportedDataApi.md#ApiExportedDatasIdDelete) | **Delete** /api/exported_datas/{id} | Removes the ExportedData resource.
[**ApiExportedDatasIdGet**](ExportedDataApi.md#ApiExportedDatasIdGet) | **Get** /api/exported_datas/{id} | Retrieves a ExportedData resource.
[**ApiExportedDatasIdPatch**](ExportedDataApi.md#ApiExportedDatasIdPatch) | **Patch** /api/exported_datas/{id} | Updates the ExportedData resource.
[**ApiExportedDatasIdPut**](ExportedDataApi.md#ApiExportedDatasIdPut) | **Put** /api/exported_datas/{id} | Replaces the ExportedData resource.
[**ApiExportedDatasPost**](ExportedDataApi.md#ApiExportedDatasPost) | **Post** /api/exported_datas | Creates a ExportedData resource.



## ApiExportedDatasGetCollection

> ApiExportedDatasGetCollection200Response ApiExportedDatasGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).DataTransferDataType(dataTransferDataType).DataTransferDataType2(dataTransferDataType2).Environment(environment).Environment2(environment2).OrderCreatedAt(orderCreatedAt).OrderUpdatedAt(orderUpdatedAt).Execute()

Retrieves the collection of ExportedData resources.



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
    dataTransferDataType := "dataTransferDataType_example" // string |  (optional)
    dataTransferDataType2 := []string{"Inner_example"} // []string |  (optional)
    environment := "environment_example" // string |  (optional)
    environment2 := []string{"Inner_example"} // []string |  (optional)
    orderCreatedAt := "orderCreatedAt_example" // string |  (optional)
    orderUpdatedAt := "orderUpdatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportedDataApi.ApiExportedDatasGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).DataTransferDataType(dataTransferDataType).DataTransferDataType2(dataTransferDataType2).Environment(environment).Environment2(environment2).OrderCreatedAt(orderCreatedAt).OrderUpdatedAt(orderUpdatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportedDataApi.ApiExportedDatasGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExportedDatasGetCollection`: ApiExportedDatasGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ExportedDataApi.ApiExportedDatasGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiExportedDatasGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **dataTransferDataType** | **string** |  | 
 **dataTransferDataType2** | **[]string** |  | 
 **environment** | **string** |  | 
 **environment2** | **[]string** |  | 
 **orderCreatedAt** | **string** |  | 
 **orderUpdatedAt** | **string** |  | 

### Return type

[**ApiExportedDatasGetCollection200Response**](ApiExportedDatasGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExportedDatasIdDelete

> ApiExportedDatasIdDelete(ctx, id).Execute()

Removes the ExportedData resource.



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
    id := "id_example" // string | ExportedData identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportedDataApi.ApiExportedDatasIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportedDataApi.ApiExportedDatasIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ExportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExportedDatasIdDeleteRequest struct via the builder pattern


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


## ApiExportedDatasIdGet

> ExportedDataJsonhal ApiExportedDatasIdGet(ctx, id).Execute()

Retrieves a ExportedData resource.



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
    id := "id_example" // string | ExportedData identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportedDataApi.ApiExportedDatasIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportedDataApi.ApiExportedDatasIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExportedDatasIdGet`: ExportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ExportedDataApi.ApiExportedDatasIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ExportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExportedDatasIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportedDataJsonhal**](ExportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExportedDatasIdPatch

> ExportedDataJsonhal ApiExportedDatasIdPatch(ctx, id).ExportedData(exportedData).Execute()

Updates the ExportedData resource.



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
    id := "id_example" // string | ExportedData identifier
    exportedData := *openapiclient.NewExportedData() // ExportedData | The updated ExportedData resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportedDataApi.ApiExportedDatasIdPatch(context.Background(), id).ExportedData(exportedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportedDataApi.ApiExportedDatasIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExportedDatasIdPatch`: ExportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ExportedDataApi.ApiExportedDatasIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ExportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExportedDatasIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportedData** | [**ExportedData**](ExportedData.md) | The updated ExportedData resource | 

### Return type

[**ExportedDataJsonhal**](ExportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExportedDatasIdPut

> ExportedDataJsonhal ApiExportedDatasIdPut(ctx, id).ExportedDataJsonhal(exportedDataJsonhal).Execute()

Replaces the ExportedData resource.



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
    id := "id_example" // string | ExportedData identifier
    exportedDataJsonhal := *openapiclient.NewExportedDataJsonhal() // ExportedDataJsonhal | The updated ExportedData resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportedDataApi.ApiExportedDatasIdPut(context.Background(), id).ExportedDataJsonhal(exportedDataJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportedDataApi.ApiExportedDatasIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExportedDatasIdPut`: ExportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ExportedDataApi.ApiExportedDatasIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ExportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiExportedDatasIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportedDataJsonhal** | [**ExportedDataJsonhal**](ExportedDataJsonhal.md) | The updated ExportedData resource | 

### Return type

[**ExportedDataJsonhal**](ExportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiExportedDatasPost

> ExportedDataJsonhal ApiExportedDatasPost(ctx).ExportedDataJsonhal(exportedDataJsonhal).Execute()

Creates a ExportedData resource.



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
    exportedDataJsonhal := *openapiclient.NewExportedDataJsonhal() // ExportedDataJsonhal | The new ExportedData resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportedDataApi.ApiExportedDatasPost(context.Background()).ExportedDataJsonhal(exportedDataJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportedDataApi.ApiExportedDatasPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiExportedDatasPost`: ExportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ExportedDataApi.ApiExportedDatasPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiExportedDatasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportedDataJsonhal** | [**ExportedDataJsonhal**](ExportedDataJsonhal.md) | The new ExportedData resource | 

### Return type

[**ExportedDataJsonhal**](ExportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

