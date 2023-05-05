# \ComponentVersionEnvVarExampleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiComponentVersionEnvVarExamplesGetCollection**](ComponentVersionEnvVarExampleApi.md#ApiComponentVersionEnvVarExamplesGetCollection) | **Get** /api/component_version_env_var_examples | Retrieves the collection of ComponentVersionEnvVarExample resources.
[**ApiComponentVersionEnvVarExamplesIdDelete**](ComponentVersionEnvVarExampleApi.md#ApiComponentVersionEnvVarExamplesIdDelete) | **Delete** /api/component_version_env_var_examples/{id} | Removes the ComponentVersionEnvVarExample resource.
[**ApiComponentVersionEnvVarExamplesIdGet**](ComponentVersionEnvVarExampleApi.md#ApiComponentVersionEnvVarExamplesIdGet) | **Get** /api/component_version_env_var_examples/{id} | Retrieves a ComponentVersionEnvVarExample resource.
[**ApiComponentVersionEnvVarExamplesIdPatch**](ComponentVersionEnvVarExampleApi.md#ApiComponentVersionEnvVarExamplesIdPatch) | **Patch** /api/component_version_env_var_examples/{id} | Updates the ComponentVersionEnvVarExample resource.
[**ApiComponentVersionEnvVarExamplesIdPut**](ComponentVersionEnvVarExampleApi.md#ApiComponentVersionEnvVarExamplesIdPut) | **Put** /api/component_version_env_var_examples/{id} | Replaces the ComponentVersionEnvVarExample resource.
[**ApiComponentVersionEnvVarExamplesPost**](ComponentVersionEnvVarExampleApi.md#ApiComponentVersionEnvVarExamplesPost) | **Post** /api/component_version_env_var_examples | Creates a ComponentVersionEnvVarExample resource.



## ApiComponentVersionEnvVarExamplesGetCollection

> []ComponentVersionEnvVarExample ApiComponentVersionEnvVarExamplesGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of ComponentVersionEnvVarExample resources.



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
    resp, r, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarExamplesGetCollection`: []ComponentVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarExamplesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarExamplesIdDelete

> ApiComponentVersionEnvVarExamplesIdDelete(ctx, id).Execute()

Removes the ComponentVersionEnvVarExample resource.



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
    id := "id_example" // string | ComponentVersionEnvVarExample identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarExamplesIdDeleteRequest struct via the builder pattern


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


## ApiComponentVersionEnvVarExamplesIdGet

> ComponentVersionEnvVarExample ApiComponentVersionEnvVarExamplesIdGet(ctx, id).Execute()

Retrieves a ComponentVersionEnvVarExample resource.



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
    id := "id_example" // string | ComponentVersionEnvVarExample identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarExamplesIdGet`: ComponentVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarExamplesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarExamplesIdPatch

> ComponentVersionEnvVarExample ApiComponentVersionEnvVarExamplesIdPatch(ctx, id).ComponentVersionEnvVarExample(componentVersionEnvVarExample).Execute()

Updates the ComponentVersionEnvVarExample resource.



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
    componentVersionEnvVarExample := *openapiclient.NewComponentVersionEnvVarExample() // ComponentVersionEnvVarExample | The updated ComponentVersionEnvVarExample resource
    id := "id_example" // string | ComponentVersionEnvVarExample identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPatch(context.Background(), id).ComponentVersionEnvVarExample(componentVersionEnvVarExample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarExamplesIdPatch`: ComponentVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarExamplesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentVersionEnvVarExample** | [**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md) | The updated ComponentVersionEnvVarExample resource | 


### Return type

[**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarExamplesIdPut

> ComponentVersionEnvVarExample ApiComponentVersionEnvVarExamplesIdPut(ctx, id).ComponentVersionEnvVarExample(componentVersionEnvVarExample).Execute()

Replaces the ComponentVersionEnvVarExample resource.



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
    componentVersionEnvVarExample := *openapiclient.NewComponentVersionEnvVarExample() // ComponentVersionEnvVarExample | The updated ComponentVersionEnvVarExample resource
    id := "id_example" // string | ComponentVersionEnvVarExample identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPut(context.Background(), id).ComponentVersionEnvVarExample(componentVersionEnvVarExample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarExamplesIdPut`: ComponentVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVarExample identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarExamplesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentVersionEnvVarExample** | [**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md) | The updated ComponentVersionEnvVarExample resource | 


### Return type

[**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarExamplesPost

> ComponentVersionEnvVarExample ApiComponentVersionEnvVarExamplesPost(ctx).ComponentVersionEnvVarExample(componentVersionEnvVarExample).Execute()

Creates a ComponentVersionEnvVarExample resource.



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
    componentVersionEnvVarExample := *openapiclient.NewComponentVersionEnvVarExample() // ComponentVersionEnvVarExample | The new ComponentVersionEnvVarExample resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesPost(context.Background()).ComponentVersionEnvVarExample(componentVersionEnvVarExample).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarExamplesPost`: ComponentVersionEnvVarExample
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarExampleApi.ApiComponentVersionEnvVarExamplesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarExamplesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentVersionEnvVarExample** | [**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md) | The new ComponentVersionEnvVarExample resource | 

### Return type

[**ComponentVersionEnvVarExample**](ComponentVersionEnvVarExample.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

