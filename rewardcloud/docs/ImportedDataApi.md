# \ImportedDataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiImportedDatasGetCollection**](ImportedDataApi.md#ApiImportedDatasGetCollection) | **Get** /api/imported_datas | Retrieves the collection of ImportedData resources.
[**ApiImportedDatasIdDelete**](ImportedDataApi.md#ApiImportedDatasIdDelete) | **Delete** /api/imported_datas/{id} | Removes the ImportedData resource.
[**ApiImportedDatasIdGet**](ImportedDataApi.md#ApiImportedDatasIdGet) | **Get** /api/imported_datas/{id} | Retrieves a ImportedData resource.
[**ApiImportedDatasIdPatch**](ImportedDataApi.md#ApiImportedDatasIdPatch) | **Patch** /api/imported_datas/{id} | Updates the ImportedData resource.
[**ApiImportedDatasIdPut**](ImportedDataApi.md#ApiImportedDatasIdPut) | **Put** /api/imported_datas/{id} | Replaces the ImportedData resource.
[**ApiImportedDatasPost**](ImportedDataApi.md#ApiImportedDatasPost) | **Post** /api/imported_datas | Creates a ImportedData resource.



## ApiImportedDatasGetCollection

> ApiImportedDatasGetCollection200Response ApiImportedDatasGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Environment(environment).Environment2(environment2).DataTransferDataType(dataTransferDataType).DataTransferDataType2(dataTransferDataType2).State(state).State2(state2).OrderCreatedAt(orderCreatedAt).OrderUpdatedAt(orderUpdatedAt).Execute()

Retrieves the collection of ImportedData resources.



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
    environment := "environment_example" // string |  (optional)
    environment2 := []string{"Inner_example"} // []string |  (optional)
    dataTransferDataType := "dataTransferDataType_example" // string |  (optional)
    dataTransferDataType2 := []string{"Inner_example"} // []string |  (optional)
    state := "state_example" // string |  (optional)
    state2 := []string{"Inner_example"} // []string |  (optional)
    orderCreatedAt := "orderCreatedAt_example" // string |  (optional)
    orderUpdatedAt := "orderUpdatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportedDataApi.ApiImportedDatasGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Environment(environment).Environment2(environment2).DataTransferDataType(dataTransferDataType).DataTransferDataType2(dataTransferDataType2).State(state).State2(state2).OrderCreatedAt(orderCreatedAt).OrderUpdatedAt(orderUpdatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportedDataApi.ApiImportedDatasGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImportedDatasGetCollection`: ApiImportedDatasGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ImportedDataApi.ApiImportedDatasGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedDatasGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **environment** | **string** |  | 
 **environment2** | **[]string** |  | 
 **dataTransferDataType** | **string** |  | 
 **dataTransferDataType2** | **[]string** |  | 
 **state** | **string** |  | 
 **state2** | **[]string** |  | 
 **orderCreatedAt** | **string** |  | 
 **orderUpdatedAt** | **string** |  | 

### Return type

[**ApiImportedDatasGetCollection200Response**](ApiImportedDatasGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiImportedDatasIdDelete

> ApiImportedDatasIdDelete(ctx, id).Execute()

Removes the ImportedData resource.



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
    id := "id_example" // string | ImportedData identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportedDataApi.ApiImportedDatasIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportedDataApi.ApiImportedDatasIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ImportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedDatasIdDeleteRequest struct via the builder pattern


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


## ApiImportedDatasIdGet

> ImportedDataJsonhal ApiImportedDatasIdGet(ctx, id).Execute()

Retrieves a ImportedData resource.



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
    id := "id_example" // string | ImportedData identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportedDataApi.ApiImportedDatasIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportedDataApi.ApiImportedDatasIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImportedDatasIdGet`: ImportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ImportedDataApi.ApiImportedDatasIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ImportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedDatasIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportedDataJsonhal**](ImportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiImportedDatasIdPatch

> ImportedDataJsonhal ApiImportedDatasIdPatch(ctx, id).ImportedData(importedData).Execute()

Updates the ImportedData resource.



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
    id := "id_example" // string | ImportedData identifier
    importedData := *openapiclient.NewImportedData() // ImportedData | The updated ImportedData resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportedDataApi.ApiImportedDatasIdPatch(context.Background(), id).ImportedData(importedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportedDataApi.ApiImportedDatasIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImportedDatasIdPatch`: ImportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ImportedDataApi.ApiImportedDatasIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ImportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedDatasIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importedData** | [**ImportedData**](ImportedData.md) | The updated ImportedData resource | 

### Return type

[**ImportedDataJsonhal**](ImportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiImportedDatasIdPut

> ImportedDataJsonhal ApiImportedDatasIdPut(ctx, id).ImportedDataJsonhal(importedDataJsonhal).Execute()

Replaces the ImportedData resource.



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
    id := "id_example" // string | ImportedData identifier
    importedDataJsonhal := *openapiclient.NewImportedDataJsonhal() // ImportedDataJsonhal | The updated ImportedData resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportedDataApi.ApiImportedDatasIdPut(context.Background(), id).ImportedDataJsonhal(importedDataJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportedDataApi.ApiImportedDatasIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImportedDatasIdPut`: ImportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ImportedDataApi.ApiImportedDatasIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ImportedData identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedDatasIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importedDataJsonhal** | [**ImportedDataJsonhal**](ImportedDataJsonhal.md) | The updated ImportedData resource | 

### Return type

[**ImportedDataJsonhal**](ImportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiImportedDatasPost

> ImportedDataJsonhal ApiImportedDatasPost(ctx).ImportedDataJsonhal(importedDataJsonhal).Execute()

Creates a ImportedData resource.



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
    importedDataJsonhal := *openapiclient.NewImportedDataJsonhal() // ImportedDataJsonhal | The new ImportedData resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImportedDataApi.ApiImportedDatasPost(context.Background()).ImportedDataJsonhal(importedDataJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportedDataApi.ApiImportedDatasPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiImportedDatasPost`: ImportedDataJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ImportedDataApi.ApiImportedDatasPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiImportedDatasPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importedDataJsonhal** | [**ImportedDataJsonhal**](ImportedDataJsonhal.md) | The new ImportedData resource | 

### Return type

[**ImportedDataJsonhal**](ImportedDataJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

