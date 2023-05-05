# \EnvironmentAccessBackendApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentAccessBackendsGetCollection**](EnvironmentAccessBackendApi.md#ApiEnvironmentAccessBackendsGetCollection) | **Get** /api/environment_access_backends | Retrieves the collection of EnvironmentAccessBackend resources.
[**ApiEnvironmentAccessBackendsIdDelete**](EnvironmentAccessBackendApi.md#ApiEnvironmentAccessBackendsIdDelete) | **Delete** /api/environment_access_backends/{id} | Removes the EnvironmentAccessBackend resource.
[**ApiEnvironmentAccessBackendsIdGet**](EnvironmentAccessBackendApi.md#ApiEnvironmentAccessBackendsIdGet) | **Get** /api/environment_access_backends/{id} | Retrieves a EnvironmentAccessBackend resource.
[**ApiEnvironmentAccessBackendsIdPatch**](EnvironmentAccessBackendApi.md#ApiEnvironmentAccessBackendsIdPatch) | **Patch** /api/environment_access_backends/{id} | Updates the EnvironmentAccessBackend resource.
[**ApiEnvironmentAccessBackendsIdPut**](EnvironmentAccessBackendApi.md#ApiEnvironmentAccessBackendsIdPut) | **Put** /api/environment_access_backends/{id} | Replaces the EnvironmentAccessBackend resource.
[**ApiEnvironmentAccessBackendsPost**](EnvironmentAccessBackendApi.md#ApiEnvironmentAccessBackendsPost) | **Post** /api/environment_access_backends | Creates a EnvironmentAccessBackend resource.



## ApiEnvironmentAccessBackendsGetCollection

> []EnvironmentAccessBackend ApiEnvironmentAccessBackendsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of EnvironmentAccessBackend resources.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    itemsPerPage := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessBackendsGetCollection`: []EnvironmentAccessBackend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessBackendsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]EnvironmentAccessBackend**](EnvironmentAccessBackend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessBackendsIdDelete

> ApiEnvironmentAccessBackendsIdDelete(ctx, id).Execute()

Removes the EnvironmentAccessBackend resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | EnvironmentAccessBackend identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessBackend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessBackendsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentAccessBackendsIdGet

> EnvironmentAccessBackend ApiEnvironmentAccessBackendsIdGet(ctx, id).Execute()

Retrieves a EnvironmentAccessBackend resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | EnvironmentAccessBackend identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessBackendsIdGet`: EnvironmentAccessBackend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessBackend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessBackendsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentAccessBackend**](EnvironmentAccessBackend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessBackendsIdPatch

> EnvironmentAccessBackend ApiEnvironmentAccessBackendsIdPatch(ctx, id).EnvironmentAccessBackend(environmentAccessBackend).Execute()

Updates the EnvironmentAccessBackend resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    environmentAccessBackend := *openapiclient.NewEnvironmentAccessBackend() // EnvironmentAccessBackend | The updated EnvironmentAccessBackend resource
    id := "id_example" // string | EnvironmentAccessBackend identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdPatch(context.Background(), id).EnvironmentAccessBackend(environmentAccessBackend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessBackendsIdPatch`: EnvironmentAccessBackend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessBackend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessBackendsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessBackend** | [**EnvironmentAccessBackend**](EnvironmentAccessBackend.md) | The updated EnvironmentAccessBackend resource | 


### Return type

[**EnvironmentAccessBackend**](EnvironmentAccessBackend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessBackendsIdPut

> EnvironmentAccessBackend ApiEnvironmentAccessBackendsIdPut(ctx, id).EnvironmentAccessBackend(environmentAccessBackend).Execute()

Replaces the EnvironmentAccessBackend resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    environmentAccessBackend := *openapiclient.NewEnvironmentAccessBackend() // EnvironmentAccessBackend | The updated EnvironmentAccessBackend resource
    id := "id_example" // string | EnvironmentAccessBackend identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdPut(context.Background(), id).EnvironmentAccessBackend(environmentAccessBackend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessBackendsIdPut`: EnvironmentAccessBackend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | EnvironmentAccessBackend identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessBackendsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessBackend** | [**EnvironmentAccessBackend**](EnvironmentAccessBackend.md) | The updated EnvironmentAccessBackend resource | 


### Return type

[**EnvironmentAccessBackend**](EnvironmentAccessBackend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentAccessBackendsPost

> EnvironmentAccessBackend ApiEnvironmentAccessBackendsPost(ctx).EnvironmentAccessBackend(environmentAccessBackend).Execute()

Creates a EnvironmentAccessBackend resource.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    environmentAccessBackend := *openapiclient.NewEnvironmentAccessBackend() // EnvironmentAccessBackend | The new EnvironmentAccessBackend resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsPost(context.Background()).EnvironmentAccessBackend(environmentAccessBackend).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentAccessBackendsPost`: EnvironmentAccessBackend
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentAccessBackendApi.ApiEnvironmentAccessBackendsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentAccessBackendsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentAccessBackend** | [**EnvironmentAccessBackend**](EnvironmentAccessBackend.md) | The new EnvironmentAccessBackend resource | 

### Return type

[**EnvironmentAccessBackend**](EnvironmentAccessBackend.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

