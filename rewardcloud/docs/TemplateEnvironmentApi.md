# \TemplateEnvironmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTemplateEnvironmentsGetCollection**](TemplateEnvironmentApi.md#ApiTemplateEnvironmentsGetCollection) | **Get** /api/template_environments | Retrieves the collection of TemplateEnvironment resources.
[**ApiTemplateEnvironmentsIdDelete**](TemplateEnvironmentApi.md#ApiTemplateEnvironmentsIdDelete) | **Delete** /api/template_environments/{id} | Removes the TemplateEnvironment resource.
[**ApiTemplateEnvironmentsIdGet**](TemplateEnvironmentApi.md#ApiTemplateEnvironmentsIdGet) | **Get** /api/template_environments/{id} | Retrieves a TemplateEnvironment resource.
[**ApiTemplateEnvironmentsIdPatch**](TemplateEnvironmentApi.md#ApiTemplateEnvironmentsIdPatch) | **Patch** /api/template_environments/{id} | Updates the TemplateEnvironment resource.
[**ApiTemplateEnvironmentsIdPut**](TemplateEnvironmentApi.md#ApiTemplateEnvironmentsIdPut) | **Put** /api/template_environments/{id} | Replaces the TemplateEnvironment resource.
[**ApiTemplateEnvironmentsPost**](TemplateEnvironmentApi.md#ApiTemplateEnvironmentsPost) | **Post** /api/template_environments | Creates a TemplateEnvironment resource.



## ApiTemplateEnvironmentsGetCollection

> []TemplateEnvironmentTemplateEnvironmentOutput ApiTemplateEnvironmentsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Id(id).Id2(id2).TemplateProject(templateProject).TemplateProject2(templateProject2).OrderUpdatedAt(orderUpdatedAt).Execute()

Retrieves the collection of TemplateEnvironment resources.



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
    id := int32(56) // int32 |  (optional)
    id2 := []int32{int32(123)} // []int32 |  (optional)
    templateProject := "templateProject_example" // string |  (optional)
    templateProject2 := []string{"Inner_example"} // []string |  (optional)
    orderUpdatedAt := "orderUpdatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateEnvironmentApi.ApiTemplateEnvironmentsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Id(id).Id2(id2).TemplateProject(templateProject).TemplateProject2(templateProject2).OrderUpdatedAt(orderUpdatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateEnvironmentApi.ApiTemplateEnvironmentsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateEnvironmentsGetCollection`: []TemplateEnvironmentTemplateEnvironmentOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateEnvironmentApi.ApiTemplateEnvironmentsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateEnvironmentsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **id** | **int32** |  | 
 **id2** | **[]int32** |  | 
 **templateProject** | **string** |  | 
 **templateProject2** | **[]string** |  | 
 **orderUpdatedAt** | **string** |  | 

### Return type

[**[]TemplateEnvironmentTemplateEnvironmentOutput**](TemplateEnvironmentTemplateEnvironmentOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateEnvironmentsIdDelete

> ApiTemplateEnvironmentsIdDelete(ctx, id).Execute()

Removes the TemplateEnvironment resource.



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
    id := "id_example" // string | TemplateEnvironment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplateEnvironmentApi.ApiTemplateEnvironmentsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateEnvironment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateEnvironmentsIdDeleteRequest struct via the builder pattern


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


## ApiTemplateEnvironmentsIdGet

> TemplateEnvironmentTemplateEnvironmentOutput ApiTemplateEnvironmentsIdGet(ctx, id).Execute()

Retrieves a TemplateEnvironment resource.



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
    id := "id_example" // string | TemplateEnvironment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateEnvironmentApi.ApiTemplateEnvironmentsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateEnvironmentsIdGet`: TemplateEnvironmentTemplateEnvironmentOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateEnvironment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateEnvironmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateEnvironmentTemplateEnvironmentOutput**](TemplateEnvironmentTemplateEnvironmentOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateEnvironmentsIdPatch

> TemplateEnvironmentTemplateEnvironmentOutput ApiTemplateEnvironmentsIdPatch(ctx, id).TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput).Execute()

Updates the TemplateEnvironment resource.



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
    templateEnvironmentTemplateEnvironmentInput := *openapiclient.NewTemplateEnvironmentTemplateEnvironmentInput() // TemplateEnvironmentTemplateEnvironmentInput | The updated TemplateEnvironment resource
    id := "id_example" // string | TemplateEnvironment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateEnvironmentApi.ApiTemplateEnvironmentsIdPatch(context.Background(), id).TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateEnvironmentsIdPatch`: TemplateEnvironmentTemplateEnvironmentOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateEnvironment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateEnvironmentsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateEnvironmentTemplateEnvironmentInput** | [**TemplateEnvironmentTemplateEnvironmentInput**](TemplateEnvironmentTemplateEnvironmentInput.md) | The updated TemplateEnvironment resource | 


### Return type

[**TemplateEnvironmentTemplateEnvironmentOutput**](TemplateEnvironmentTemplateEnvironmentOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateEnvironmentsIdPut

> TemplateEnvironmentTemplateEnvironmentOutput ApiTemplateEnvironmentsIdPut(ctx, id).TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput).Execute()

Replaces the TemplateEnvironment resource.



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
    templateEnvironmentTemplateEnvironmentInput := *openapiclient.NewTemplateEnvironmentTemplateEnvironmentInput() // TemplateEnvironmentTemplateEnvironmentInput | The updated TemplateEnvironment resource
    id := "id_example" // string | TemplateEnvironment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateEnvironmentApi.ApiTemplateEnvironmentsIdPut(context.Background(), id).TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateEnvironmentsIdPut`: TemplateEnvironmentTemplateEnvironmentOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateEnvironmentApi.ApiTemplateEnvironmentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateEnvironment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateEnvironmentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateEnvironmentTemplateEnvironmentInput** | [**TemplateEnvironmentTemplateEnvironmentInput**](TemplateEnvironmentTemplateEnvironmentInput.md) | The updated TemplateEnvironment resource | 


### Return type

[**TemplateEnvironmentTemplateEnvironmentOutput**](TemplateEnvironmentTemplateEnvironmentOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateEnvironmentsPost

> TemplateEnvironmentTemplateEnvironmentOutput ApiTemplateEnvironmentsPost(ctx).TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput).Execute()

Creates a TemplateEnvironment resource.



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
    templateEnvironmentTemplateEnvironmentInput := *openapiclient.NewTemplateEnvironmentTemplateEnvironmentInput() // TemplateEnvironmentTemplateEnvironmentInput | The new TemplateEnvironment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateEnvironmentApi.ApiTemplateEnvironmentsPost(context.Background()).TemplateEnvironmentTemplateEnvironmentInput(templateEnvironmentTemplateEnvironmentInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateEnvironmentApi.ApiTemplateEnvironmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateEnvironmentsPost`: TemplateEnvironmentTemplateEnvironmentOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateEnvironmentApi.ApiTemplateEnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateEnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateEnvironmentTemplateEnvironmentInput** | [**TemplateEnvironmentTemplateEnvironmentInput**](TemplateEnvironmentTemplateEnvironmentInput.md) | The new TemplateEnvironment resource | 

### Return type

[**TemplateEnvironmentTemplateEnvironmentOutput**](TemplateEnvironmentTemplateEnvironmentOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

