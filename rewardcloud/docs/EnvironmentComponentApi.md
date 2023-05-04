# \EnvironmentComponentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentComponentsGetCollection**](EnvironmentComponentApi.md#ApiEnvironmentComponentsGetCollection) | **Get** /api/environment_components | Retrieves the collection of EnvironmentComponent resources.
[**ApiEnvironmentComponentsIdDelete**](EnvironmentComponentApi.md#ApiEnvironmentComponentsIdDelete) | **Delete** /api/environment_components/{id} | Removes the EnvironmentComponent resource.
[**ApiEnvironmentComponentsIdGet**](EnvironmentComponentApi.md#ApiEnvironmentComponentsIdGet) | **Get** /api/environment_components/{id} | Retrieves a EnvironmentComponent resource.
[**ApiEnvironmentComponentsIdPatch**](EnvironmentComponentApi.md#ApiEnvironmentComponentsIdPatch) | **Patch** /api/environment_components/{id} | Updates the EnvironmentComponent resource.
[**ApiEnvironmentComponentsIdPut**](EnvironmentComponentApi.md#ApiEnvironmentComponentsIdPut) | **Put** /api/environment_components/{id} | Replaces the EnvironmentComponent resource.
[**ApiEnvironmentComponentsPost**](EnvironmentComponentApi.md#ApiEnvironmentComponentsPost) | **Post** /api/environment_components | Creates a EnvironmentComponent resource.



## ApiEnvironmentComponentsGetCollection

> ApiEnvironmentComponentsGetCollection200Response ApiEnvironmentComponentsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ComponentVersion(componentVersion).ComponentVersion2(componentVersion2).Environment(environment).Environment2(environment2).Execute()

Retrieves the collection of EnvironmentComponent resources.



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
    componentVersion := "componentVersion_example" // string |  (optional)
    componentVersion2 := []string{"Inner_example"} // []string |  (optional)
    environment := "environment_example" // string |  (optional)
    environment2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentComponentApi.ApiEnvironmentComponentsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ComponentVersion(componentVersion).ComponentVersion2(componentVersion2).Environment(environment).Environment2(environment2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentComponentApi.ApiEnvironmentComponentsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentComponentsGetCollection`: ApiEnvironmentComponentsGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentComponentApi.ApiEnvironmentComponentsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentComponentsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **componentVersion** | **string** |  | 
 **componentVersion2** | **[]string** |  | 
 **environment** | **string** |  | 
 **environment2** | **[]string** |  | 

### Return type

[**ApiEnvironmentComponentsGetCollection200Response**](ApiEnvironmentComponentsGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentComponentsIdDelete

> ApiEnvironmentComponentsIdDelete(ctx, id).Execute()

Removes the EnvironmentComponent resource.



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
    id := "id_example" // string | EnvironmentComponent identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentComponentApi.ApiEnvironmentComponentsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentComponentApi.ApiEnvironmentComponentsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentComponent identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentComponentsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentComponentsIdGet

> EnvironmentComponentJsonhal ApiEnvironmentComponentsIdGet(ctx, id).Execute()

Retrieves a EnvironmentComponent resource.



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
    id := "id_example" // string | EnvironmentComponent identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentComponentApi.ApiEnvironmentComponentsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentComponentApi.ApiEnvironmentComponentsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentComponentsIdGet`: EnvironmentComponentJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentComponentApi.ApiEnvironmentComponentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentComponent identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentComponentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentComponentsIdPatch

> EnvironmentComponentJsonhal ApiEnvironmentComponentsIdPatch(ctx, id).EnvironmentComponent(environmentComponent).Execute()

Updates the EnvironmentComponent resource.



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
    id := "id_example" // string | EnvironmentComponent identifier
    environmentComponent := *openapiclient.NewEnvironmentComponent() // EnvironmentComponent | The updated EnvironmentComponent resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentComponentApi.ApiEnvironmentComponentsIdPatch(context.Background(), id).EnvironmentComponent(environmentComponent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentComponentApi.ApiEnvironmentComponentsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentComponentsIdPatch`: EnvironmentComponentJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentComponentApi.ApiEnvironmentComponentsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentComponent identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentComponentsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentComponent** | [**EnvironmentComponent**](EnvironmentComponent.md) | The updated EnvironmentComponent resource | 

### Return type

[**EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentComponentsIdPut

> EnvironmentComponentJsonhal ApiEnvironmentComponentsIdPut(ctx, id).EnvironmentComponentJsonhal(environmentComponentJsonhal).Execute()

Replaces the EnvironmentComponent resource.



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
    id := "id_example" // string | EnvironmentComponent identifier
    environmentComponentJsonhal := *openapiclient.NewEnvironmentComponentJsonhal() // EnvironmentComponentJsonhal | The updated EnvironmentComponent resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentComponentApi.ApiEnvironmentComponentsIdPut(context.Background(), id).EnvironmentComponentJsonhal(environmentComponentJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentComponentApi.ApiEnvironmentComponentsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentComponentsIdPut`: EnvironmentComponentJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentComponentApi.ApiEnvironmentComponentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentComponent identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentComponentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentComponentJsonhal** | [**EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md) | The updated EnvironmentComponent resource | 

### Return type

[**EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentComponentsPost

> EnvironmentComponentJsonhal ApiEnvironmentComponentsPost(ctx).EnvironmentComponentJsonhal(environmentComponentJsonhal).Execute()

Creates a EnvironmentComponent resource.



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
    environmentComponentJsonhal := *openapiclient.NewEnvironmentComponentJsonhal() // EnvironmentComponentJsonhal | The new EnvironmentComponent resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentComponentApi.ApiEnvironmentComponentsPost(context.Background()).EnvironmentComponentJsonhal(environmentComponentJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentComponentApi.ApiEnvironmentComponentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentComponentsPost`: EnvironmentComponentJsonhal
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentComponentApi.ApiEnvironmentComponentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentComponentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentComponentJsonhal** | [**EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md) | The new EnvironmentComponent resource | 

### Return type

[**EnvironmentComponentJsonhal**](EnvironmentComponentJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

