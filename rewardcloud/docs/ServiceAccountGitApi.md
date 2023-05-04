# \ServiceAccountGitApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServiceAccountGitsGetCollection**](ServiceAccountGitApi.md#ApiServiceAccountGitsGetCollection) | **Get** /api/service_account_gits | Retrieves the collection of ServiceAccountGit resources.
[**ApiServiceAccountGitsIdDelete**](ServiceAccountGitApi.md#ApiServiceAccountGitsIdDelete) | **Delete** /api/service_account_gits/{id} | Removes the ServiceAccountGit resource.
[**ApiServiceAccountGitsIdGet**](ServiceAccountGitApi.md#ApiServiceAccountGitsIdGet) | **Get** /api/service_account_gits/{id} | Retrieves a ServiceAccountGit resource.
[**ApiServiceAccountGitsIdPatch**](ServiceAccountGitApi.md#ApiServiceAccountGitsIdPatch) | **Patch** /api/service_account_gits/{id} | Updates the ServiceAccountGit resource.
[**ApiServiceAccountGitsIdPut**](ServiceAccountGitApi.md#ApiServiceAccountGitsIdPut) | **Put** /api/service_account_gits/{id} | Replaces the ServiceAccountGit resource.
[**ApiServiceAccountGitsPost**](ServiceAccountGitApi.md#ApiServiceAccountGitsPost) | **Post** /api/service_account_gits | Creates a ServiceAccountGit resource.



## ApiServiceAccountGitsGetCollection

> ApiServiceAccountGitsGetCollection200Response ApiServiceAccountGitsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ServiceAccount(serviceAccount).ServiceAccount2(serviceAccount2).Execute()

Retrieves the collection of ServiceAccountGit resources.



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
    serviceAccount := "serviceAccount_example" // string |  (optional)
    serviceAccount2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountGitApi.ApiServiceAccountGitsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ServiceAccount(serviceAccount).ServiceAccount2(serviceAccount2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountGitApi.ApiServiceAccountGitsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountGitsGetCollection`: ApiServiceAccountGitsGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountGitApi.ApiServiceAccountGitsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountGitsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **serviceAccount** | **string** |  | 
 **serviceAccount2** | **[]string** |  | 

### Return type

[**ApiServiceAccountGitsGetCollection200Response**](ApiServiceAccountGitsGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountGitsIdDelete

> ApiServiceAccountGitsIdDelete(ctx, id).Execute()

Removes the ServiceAccountGit resource.



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
    id := "id_example" // string | ServiceAccountGit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountGitApi.ApiServiceAccountGitsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountGitApi.ApiServiceAccountGitsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountGit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountGitsIdDeleteRequest struct via the builder pattern


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


## ApiServiceAccountGitsIdGet

> ServiceAccountGitJsonhal ApiServiceAccountGitsIdGet(ctx, id).Execute()

Retrieves a ServiceAccountGit resource.



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
    id := "id_example" // string | ServiceAccountGit identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountGitApi.ApiServiceAccountGitsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountGitApi.ApiServiceAccountGitsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountGitsIdGet`: ServiceAccountGitJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountGitApi.ApiServiceAccountGitsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountGit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountGitsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccountGitJsonhal**](ServiceAccountGitJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountGitsIdPatch

> ServiceAccountGitJsonhal ApiServiceAccountGitsIdPatch(ctx, id).ServiceAccountGit(serviceAccountGit).Execute()

Updates the ServiceAccountGit resource.



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
    id := "id_example" // string | ServiceAccountGit identifier
    serviceAccountGit := *openapiclient.NewServiceAccountGit() // ServiceAccountGit | The updated ServiceAccountGit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountGitApi.ApiServiceAccountGitsIdPatch(context.Background(), id).ServiceAccountGit(serviceAccountGit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountGitApi.ApiServiceAccountGitsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountGitsIdPatch`: ServiceAccountGitJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountGitApi.ApiServiceAccountGitsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountGit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountGitsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountGit** | [**ServiceAccountGit**](ServiceAccountGit.md) | The updated ServiceAccountGit resource | 

### Return type

[**ServiceAccountGitJsonhal**](ServiceAccountGitJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountGitsIdPut

> ServiceAccountGitJsonhal ApiServiceAccountGitsIdPut(ctx, id).ServiceAccountGitJsonhal(serviceAccountGitJsonhal).Execute()

Replaces the ServiceAccountGit resource.



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
    id := "id_example" // string | ServiceAccountGit identifier
    serviceAccountGitJsonhal := *openapiclient.NewServiceAccountGitJsonhal() // ServiceAccountGitJsonhal | The updated ServiceAccountGit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountGitApi.ApiServiceAccountGitsIdPut(context.Background(), id).ServiceAccountGitJsonhal(serviceAccountGitJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountGitApi.ApiServiceAccountGitsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountGitsIdPut`: ServiceAccountGitJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountGitApi.ApiServiceAccountGitsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccountGit identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountGitsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountGitJsonhal** | [**ServiceAccountGitJsonhal**](ServiceAccountGitJsonhal.md) | The updated ServiceAccountGit resource | 

### Return type

[**ServiceAccountGitJsonhal**](ServiceAccountGitJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountGitsPost

> ServiceAccountGitJsonhal ApiServiceAccountGitsPost(ctx).ServiceAccountGitJsonhal(serviceAccountGitJsonhal).Execute()

Creates a ServiceAccountGit resource.



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
    serviceAccountGitJsonhal := *openapiclient.NewServiceAccountGitJsonhal() // ServiceAccountGitJsonhal | The new ServiceAccountGit resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountGitApi.ApiServiceAccountGitsPost(context.Background()).ServiceAccountGitJsonhal(serviceAccountGitJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountGitApi.ApiServiceAccountGitsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountGitsPost`: ServiceAccountGitJsonhal
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountGitApi.ApiServiceAccountGitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountGitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountGitJsonhal** | [**ServiceAccountGitJsonhal**](ServiceAccountGitJsonhal.md) | The new ServiceAccountGit resource | 

### Return type

[**ServiceAccountGitJsonhal**](ServiceAccountGitJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

