# \ResourceTypeLimitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiResourceTypeLimitsGetCollection**](ResourceTypeLimitApi.md#ApiResourceTypeLimitsGetCollection) | **Get** /api/resource_type_limits | Retrieves the collection of ResourceTypeLimit resources.
[**ApiResourceTypeLimitsIdDelete**](ResourceTypeLimitApi.md#ApiResourceTypeLimitsIdDelete) | **Delete** /api/resource_type_limits/{id} | Removes the ResourceTypeLimit resource.
[**ApiResourceTypeLimitsIdGet**](ResourceTypeLimitApi.md#ApiResourceTypeLimitsIdGet) | **Get** /api/resource_type_limits/{id} | Retrieves a ResourceTypeLimit resource.
[**ApiResourceTypeLimitsIdPatch**](ResourceTypeLimitApi.md#ApiResourceTypeLimitsIdPatch) | **Patch** /api/resource_type_limits/{id} | Updates the ResourceTypeLimit resource.
[**ApiResourceTypeLimitsIdPut**](ResourceTypeLimitApi.md#ApiResourceTypeLimitsIdPut) | **Put** /api/resource_type_limits/{id} | Replaces the ResourceTypeLimit resource.
[**ApiResourceTypeLimitsPost**](ResourceTypeLimitApi.md#ApiResourceTypeLimitsPost) | **Post** /api/resource_type_limits | Creates a ResourceTypeLimit resource.



## ApiResourceTypeLimitsGetCollection

> []ResourceTypeLimit ApiResourceTypeLimitsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of ResourceTypeLimit resources.



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
    resp, r, err := apiClient.ResourceTypeLimitApi.ApiResourceTypeLimitsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeLimitApi.ApiResourceTypeLimitsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypeLimitsGetCollection`: []ResourceTypeLimit
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeLimitApi.ApiResourceTypeLimitsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypeLimitsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]ResourceTypeLimit**](ResourceTypeLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypeLimitsIdDelete

> ApiResourceTypeLimitsIdDelete(ctx, id).Execute()

Removes the ResourceTypeLimit resource.



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
    id := "id_example" // string | ResourceTypeLimit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceTypeLimitApi.ApiResourceTypeLimitsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeLimitApi.ApiResourceTypeLimitsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceTypeLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypeLimitsIdDeleteRequest struct via the builder pattern


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


## ApiResourceTypeLimitsIdGet

> ResourceTypeLimit ApiResourceTypeLimitsIdGet(ctx, id).Execute()

Retrieves a ResourceTypeLimit resource.



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
    id := "id_example" // string | ResourceTypeLimit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeLimitApi.ApiResourceTypeLimitsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeLimitApi.ApiResourceTypeLimitsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypeLimitsIdGet`: ResourceTypeLimit
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeLimitApi.ApiResourceTypeLimitsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceTypeLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypeLimitsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceTypeLimit**](ResourceTypeLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypeLimitsIdPatch

> ResourceTypeLimit ApiResourceTypeLimitsIdPatch(ctx, id).ResourceTypeLimit(resourceTypeLimit).Execute()

Updates the ResourceTypeLimit resource.



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
    resourceTypeLimit := *openapiclient.NewResourceTypeLimit() // ResourceTypeLimit | The updated ResourceTypeLimit resource
    id := "id_example" // string | ResourceTypeLimit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeLimitApi.ApiResourceTypeLimitsIdPatch(context.Background(), id).ResourceTypeLimit(resourceTypeLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeLimitApi.ApiResourceTypeLimitsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypeLimitsIdPatch`: ResourceTypeLimit
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeLimitApi.ApiResourceTypeLimitsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceTypeLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypeLimitsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceTypeLimit** | [**ResourceTypeLimit**](ResourceTypeLimit.md) | The updated ResourceTypeLimit resource | 


### Return type

[**ResourceTypeLimit**](ResourceTypeLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypeLimitsIdPut

> ResourceTypeLimit ApiResourceTypeLimitsIdPut(ctx, id).ResourceTypeLimit(resourceTypeLimit).Execute()

Replaces the ResourceTypeLimit resource.



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
    resourceTypeLimit := *openapiclient.NewResourceTypeLimit() // ResourceTypeLimit | The updated ResourceTypeLimit resource
    id := "id_example" // string | ResourceTypeLimit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeLimitApi.ApiResourceTypeLimitsIdPut(context.Background(), id).ResourceTypeLimit(resourceTypeLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeLimitApi.ApiResourceTypeLimitsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypeLimitsIdPut`: ResourceTypeLimit
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeLimitApi.ApiResourceTypeLimitsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ResourceTypeLimit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypeLimitsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceTypeLimit** | [**ResourceTypeLimit**](ResourceTypeLimit.md) | The updated ResourceTypeLimit resource | 


### Return type

[**ResourceTypeLimit**](ResourceTypeLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiResourceTypeLimitsPost

> ResourceTypeLimit ApiResourceTypeLimitsPost(ctx).ResourceTypeLimit(resourceTypeLimit).Execute()

Creates a ResourceTypeLimit resource.



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
    resourceTypeLimit := *openapiclient.NewResourceTypeLimit() // ResourceTypeLimit | The new ResourceTypeLimit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceTypeLimitApi.ApiResourceTypeLimitsPost(context.Background()).ResourceTypeLimit(resourceTypeLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceTypeLimitApi.ApiResourceTypeLimitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiResourceTypeLimitsPost`: ResourceTypeLimit
    fmt.Fprintf(os.Stdout, "Response from `ResourceTypeLimitApi.ApiResourceTypeLimitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiResourceTypeLimitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceTypeLimit** | [**ResourceTypeLimit**](ResourceTypeLimit.md) | The new ResourceTypeLimit resource | 

### Return type

[**ResourceTypeLimit**](ResourceTypeLimit.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

