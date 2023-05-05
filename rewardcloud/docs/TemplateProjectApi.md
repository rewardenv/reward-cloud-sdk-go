# \TemplateProjectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTemplateProjectsGetCollection**](TemplateProjectApi.md#ApiTemplateProjectsGetCollection) | **Get** /api/template_projects | Retrieves the collection of TemplateProject resources.
[**ApiTemplateProjectsIdDelete**](TemplateProjectApi.md#ApiTemplateProjectsIdDelete) | **Delete** /api/template_projects/{id} | Removes the TemplateProject resource.
[**ApiTemplateProjectsIdGet**](TemplateProjectApi.md#ApiTemplateProjectsIdGet) | **Get** /api/template_projects/{id} | Retrieves a TemplateProject resource.
[**ApiTemplateProjectsIdPatch**](TemplateProjectApi.md#ApiTemplateProjectsIdPatch) | **Patch** /api/template_projects/{id} | Updates the TemplateProject resource.
[**ApiTemplateProjectsIdPut**](TemplateProjectApi.md#ApiTemplateProjectsIdPut) | **Put** /api/template_projects/{id} | Replaces the TemplateProject resource.
[**ApiTemplateProjectsPost**](TemplateProjectApi.md#ApiTemplateProjectsPost) | **Post** /api/template_projects | Creates a TemplateProject resource.



## ApiTemplateProjectsGetCollection

> []TemplateProjectTemplateProjectOutput ApiTemplateProjectsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).OrderUpdatedAt(orderUpdatedAt).Execute()

Retrieves the collection of TemplateProject resources.



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
    projectTypeVersion := "projectTypeVersion_example" // string |  (optional)
    projectTypeVersion2 := []string{"Inner_example"} // []string |  (optional)
    orderUpdatedAt := "orderUpdatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateProjectApi.ApiTemplateProjectsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ProjectTypeVersion(projectTypeVersion).ProjectTypeVersion2(projectTypeVersion2).OrderUpdatedAt(orderUpdatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjectApi.ApiTemplateProjectsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateProjectsGetCollection`: []TemplateProjectTemplateProjectOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateProjectApi.ApiTemplateProjectsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateProjectsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **projectTypeVersion** | **string** |  | 
 **projectTypeVersion2** | **[]string** |  | 
 **orderUpdatedAt** | **string** |  | 

### Return type

[**[]TemplateProjectTemplateProjectOutput**](TemplateProjectTemplateProjectOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateProjectsIdDelete

> ApiTemplateProjectsIdDelete(ctx, id).Execute()

Removes the TemplateProject resource.



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
    id := "id_example" // string | TemplateProject identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplateProjectApi.ApiTemplateProjectsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjectApi.ApiTemplateProjectsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateProject identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateProjectsIdDeleteRequest struct via the builder pattern


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


## ApiTemplateProjectsIdGet

> TemplateProjectTemplateProjectOutput ApiTemplateProjectsIdGet(ctx, id).Execute()

Retrieves a TemplateProject resource.



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
    id := "id_example" // string | TemplateProject identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateProjectApi.ApiTemplateProjectsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjectApi.ApiTemplateProjectsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateProjectsIdGet`: TemplateProjectTemplateProjectOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateProjectApi.ApiTemplateProjectsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateProject identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateProjectsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateProjectTemplateProjectOutput**](TemplateProjectTemplateProjectOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateProjectsIdPatch

> TemplateProjectTemplateProjectOutput ApiTemplateProjectsIdPatch(ctx, id).TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput).Execute()

Updates the TemplateProject resource.



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
    templateProjectTemplateProjectInput := *openapiclient.NewTemplateProjectTemplateProjectInput() // TemplateProjectTemplateProjectInput | The updated TemplateProject resource
    id := "id_example" // string | TemplateProject identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateProjectApi.ApiTemplateProjectsIdPatch(context.Background(), id).TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjectApi.ApiTemplateProjectsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateProjectsIdPatch`: TemplateProjectTemplateProjectOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateProjectApi.ApiTemplateProjectsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateProject identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateProjectsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateProjectTemplateProjectInput** | [**TemplateProjectTemplateProjectInput**](TemplateProjectTemplateProjectInput.md) | The updated TemplateProject resource | 


### Return type

[**TemplateProjectTemplateProjectOutput**](TemplateProjectTemplateProjectOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateProjectsIdPut

> TemplateProjectTemplateProjectOutput ApiTemplateProjectsIdPut(ctx, id).TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput).Execute()

Replaces the TemplateProject resource.



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
    templateProjectTemplateProjectInput := *openapiclient.NewTemplateProjectTemplateProjectInput() // TemplateProjectTemplateProjectInput | The updated TemplateProject resource
    id := "id_example" // string | TemplateProject identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateProjectApi.ApiTemplateProjectsIdPut(context.Background(), id).TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjectApi.ApiTemplateProjectsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateProjectsIdPut`: TemplateProjectTemplateProjectOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateProjectApi.ApiTemplateProjectsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | TemplateProject identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateProjectsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateProjectTemplateProjectInput** | [**TemplateProjectTemplateProjectInput**](TemplateProjectTemplateProjectInput.md) | The updated TemplateProject resource | 


### Return type

[**TemplateProjectTemplateProjectOutput**](TemplateProjectTemplateProjectOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTemplateProjectsPost

> TemplateProjectTemplateProjectOutput ApiTemplateProjectsPost(ctx).TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput).Execute()

Creates a TemplateProject resource.



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
    templateProjectTemplateProjectInput := *openapiclient.NewTemplateProjectTemplateProjectInput() // TemplateProjectTemplateProjectInput | The new TemplateProject resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplateProjectApi.ApiTemplateProjectsPost(context.Background()).TemplateProjectTemplateProjectInput(templateProjectTemplateProjectInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplateProjectApi.ApiTemplateProjectsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTemplateProjectsPost`: TemplateProjectTemplateProjectOutput
    fmt.Fprintf(os.Stdout, "Response from `TemplateProjectApi.ApiTemplateProjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTemplateProjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateProjectTemplateProjectInput** | [**TemplateProjectTemplateProjectInput**](TemplateProjectTemplateProjectInput.md) | The new TemplateProject resource | 

### Return type

[**TemplateProjectTemplateProjectOutput**](TemplateProjectTemplateProjectOutput.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

