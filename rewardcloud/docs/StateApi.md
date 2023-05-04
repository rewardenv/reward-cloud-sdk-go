# \StateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiStatesGetCollection**](StateApi.md#ApiStatesGetCollection) | **Get** /api/states | Retrieves the collection of State resources.
[**ApiStatesIdDelete**](StateApi.md#ApiStatesIdDelete) | **Delete** /api/states/{id} | Removes the State resource.
[**ApiStatesIdGet**](StateApi.md#ApiStatesIdGet) | **Get** /api/states/{id} | Retrieves a State resource.
[**ApiStatesIdPatch**](StateApi.md#ApiStatesIdPatch) | **Patch** /api/states/{id} | Updates the State resource.
[**ApiStatesIdPut**](StateApi.md#ApiStatesIdPut) | **Put** /api/states/{id} | Replaces the State resource.
[**ApiStatesPost**](StateApi.md#ApiStatesPost) | **Post** /api/states | Creates a State resource.



## ApiStatesGetCollection

> ApiStatesGetCollection200Response ApiStatesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of State resources.



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
    resp, r, err := apiClient.StateApi.ApiStatesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.ApiStatesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStatesGetCollection`: ApiStatesGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `StateApi.ApiStatesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiStatesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**ApiStatesGetCollection200Response**](ApiStatesGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStatesIdDelete

> ApiStatesIdDelete(ctx, id).Execute()

Removes the State resource.



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
    id := "id_example" // string | State identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.ApiStatesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.ApiStatesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | State identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStatesIdDeleteRequest struct via the builder pattern


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


## ApiStatesIdGet

> StateJsonhal ApiStatesIdGet(ctx, id).Execute()

Retrieves a State resource.



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
    id := "id_example" // string | State identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.ApiStatesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.ApiStatesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStatesIdGet`: StateJsonhal
    fmt.Fprintf(os.Stdout, "Response from `StateApi.ApiStatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | State identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StateJsonhal**](StateJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStatesIdPatch

> StateJsonhal ApiStatesIdPatch(ctx, id).State(state).Execute()

Updates the State resource.



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
    id := "id_example" // string | State identifier
    state := *openapiclient.NewState() // State | The updated State resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.ApiStatesIdPatch(context.Background(), id).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.ApiStatesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStatesIdPatch`: StateJsonhal
    fmt.Fprintf(os.Stdout, "Response from `StateApi.ApiStatesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | State identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStatesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | [**State**](State.md) | The updated State resource | 

### Return type

[**StateJsonhal**](StateJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStatesIdPut

> StateJsonhal ApiStatesIdPut(ctx, id).StateJsonhal(stateJsonhal).Execute()

Replaces the State resource.



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
    id := "id_example" // string | State identifier
    stateJsonhal := *openapiclient.NewStateJsonhal() // StateJsonhal | The updated State resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.ApiStatesIdPut(context.Background(), id).StateJsonhal(stateJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.ApiStatesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStatesIdPut`: StateJsonhal
    fmt.Fprintf(os.Stdout, "Response from `StateApi.ApiStatesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | State identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiStatesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateJsonhal** | [**StateJsonhal**](StateJsonhal.md) | The updated State resource | 

### Return type

[**StateJsonhal**](StateJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiStatesPost

> StateJsonhal ApiStatesPost(ctx).StateJsonhal(stateJsonhal).Execute()

Creates a State resource.



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
    stateJsonhal := *openapiclient.NewStateJsonhal() // StateJsonhal | The new State resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StateApi.ApiStatesPost(context.Background()).StateJsonhal(stateJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StateApi.ApiStatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiStatesPost`: StateJsonhal
    fmt.Fprintf(os.Stdout, "Response from `StateApi.ApiStatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiStatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateJsonhal** | [**StateJsonhal**](StateJsonhal.md) | The new State resource | 

### Return type

[**StateJsonhal**](StateJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

