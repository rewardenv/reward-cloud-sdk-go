# \ServiceAccountApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiServiceAccountsGetCollection**](ServiceAccountApi.md#ApiServiceAccountsGetCollection) | **Get** /api/service_accounts | Retrieves the collection of ServiceAccount resources.
[**ApiServiceAccountsIdDelete**](ServiceAccountApi.md#ApiServiceAccountsIdDelete) | **Delete** /api/service_accounts/{id} | Removes the ServiceAccount resource.
[**ApiServiceAccountsIdGet**](ServiceAccountApi.md#ApiServiceAccountsIdGet) | **Get** /api/service_accounts/{id} | Retrieves a ServiceAccount resource.
[**ApiServiceAccountsIdPatch**](ServiceAccountApi.md#ApiServiceAccountsIdPatch) | **Patch** /api/service_accounts/{id} | Updates the ServiceAccount resource.
[**ApiServiceAccountsIdPut**](ServiceAccountApi.md#ApiServiceAccountsIdPut) | **Put** /api/service_accounts/{id} | Replaces the ServiceAccount resource.
[**ApiServiceAccountsPost**](ServiceAccountApi.md#ApiServiceAccountsPost) | **Post** /api/service_accounts | Creates a ServiceAccount resource.



## ApiServiceAccountsGetCollection

> []ServiceAccount ApiServiceAccountsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Project(project).Project2(project2).ServiceAccountGit(serviceAccountGit).ServiceAccountGit2(serviceAccountGit2).ServiceAccountRegistry(serviceAccountRegistry).ServiceAccountRegistry2(serviceAccountRegistry2).Execute()

Retrieves the collection of ServiceAccount resources.



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
    project := "project_example" // string |  (optional)
    project2 := []string{"Inner_example"} // []string |  (optional)
    serviceAccountGit := "serviceAccountGit_example" // string |  (optional)
    serviceAccountGit2 := []string{"Inner_example"} // []string |  (optional)
    serviceAccountRegistry := "serviceAccountRegistry_example" // string |  (optional)
    serviceAccountRegistry2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ApiServiceAccountsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Project(project).Project2(project2).ServiceAccountGit(serviceAccountGit).ServiceAccountGit2(serviceAccountGit2).ServiceAccountRegistry(serviceAccountRegistry).ServiceAccountRegistry2(serviceAccountRegistry2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ApiServiceAccountsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsGetCollection`: []ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ApiServiceAccountsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **project** | **string** |  | 
 **project2** | **[]string** |  | 
 **serviceAccountGit** | **string** |  | 
 **serviceAccountGit2** | **[]string** |  | 
 **serviceAccountRegistry** | **string** |  | 
 **serviceAccountRegistry2** | **[]string** |  | 

### Return type

[**[]ServiceAccount**](ServiceAccount.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountsIdDelete

> ApiServiceAccountsIdDelete(ctx, id).Execute()

Removes the ServiceAccount resource.



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
    id := "id_example" // string | ServiceAccount identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceAccountApi.ApiServiceAccountsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ApiServiceAccountsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdDeleteRequest struct via the builder pattern


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


## ApiServiceAccountsIdGet

> ServiceAccount ApiServiceAccountsIdGet(ctx, id).Execute()

Retrieves a ServiceAccount resource.



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
    id := "id_example" // string | ServiceAccount identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ApiServiceAccountsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ApiServiceAccountsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsIdGet`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ApiServiceAccountsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountsIdPatch

> ServiceAccount ApiServiceAccountsIdPatch(ctx, id).ServiceAccount(serviceAccount).Execute()

Updates the ServiceAccount resource.



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
    serviceAccount := *openapiclient.NewServiceAccount() // ServiceAccount | The updated ServiceAccount resource
    id := "id_example" // string | ServiceAccount identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ApiServiceAccountsIdPatch(context.Background(), id).ServiceAccount(serviceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ApiServiceAccountsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsIdPatch`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ApiServiceAccountsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccount** | [**ServiceAccount**](ServiceAccount.md) | The updated ServiceAccount resource | 


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/vnd.api+json
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountsIdPut

> ServiceAccount ApiServiceAccountsIdPut(ctx, id).ServiceAccount(serviceAccount).Execute()

Replaces the ServiceAccount resource.



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
    serviceAccount := *openapiclient.NewServiceAccount() // ServiceAccount | The updated ServiceAccount resource
    id := "id_example" // string | ServiceAccount identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ApiServiceAccountsIdPut(context.Background(), id).ServiceAccount(serviceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ApiServiceAccountsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsIdPut`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ApiServiceAccountsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ServiceAccount identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccount** | [**ServiceAccount**](ServiceAccount.md) | The updated ServiceAccount resource | 


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiServiceAccountsPost

> ServiceAccount ApiServiceAccountsPost(ctx).ServiceAccount(serviceAccount).Execute()

Creates a ServiceAccount resource.



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
    serviceAccount := *openapiclient.NewServiceAccount() // ServiceAccount | The new ServiceAccount resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceAccountApi.ApiServiceAccountsPost(context.Background()).ServiceAccount(serviceAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountApi.ApiServiceAccountsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiServiceAccountsPost`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountApi.ApiServiceAccountsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiServiceAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccount** | [**ServiceAccount**](ServiceAccount.md) | The new ServiceAccount resource | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

