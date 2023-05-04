# \RoleGroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRoleGroupsGetCollection**](RoleGroupApi.md#ApiRoleGroupsGetCollection) | **Get** /api/role_groups | Retrieves the collection of RoleGroup resources.
[**ApiRoleGroupsIdDelete**](RoleGroupApi.md#ApiRoleGroupsIdDelete) | **Delete** /api/role_groups/{id} | Removes the RoleGroup resource.
[**ApiRoleGroupsIdGet**](RoleGroupApi.md#ApiRoleGroupsIdGet) | **Get** /api/role_groups/{id} | Retrieves a RoleGroup resource.
[**ApiRoleGroupsIdPatch**](RoleGroupApi.md#ApiRoleGroupsIdPatch) | **Patch** /api/role_groups/{id} | Updates the RoleGroup resource.
[**ApiRoleGroupsIdPut**](RoleGroupApi.md#ApiRoleGroupsIdPut) | **Put** /api/role_groups/{id} | Replaces the RoleGroup resource.
[**ApiRoleGroupsPost**](RoleGroupApi.md#ApiRoleGroupsPost) | **Post** /api/role_groups | Creates a RoleGroup resource.



## ApiRoleGroupsGetCollection

> ApiRoleGroupsGetCollection200Response ApiRoleGroupsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).Execute()

Retrieves the collection of RoleGroup resources.



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
    resp, r, err := apiClient.RoleGroupApi.ApiRoleGroupsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleGroupApi.ApiRoleGroupsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRoleGroupsGetCollection`: ApiRoleGroupsGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleGroupApi.ApiRoleGroupsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRoleGroupsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]

### Return type

[**ApiRoleGroupsGetCollection200Response**](ApiRoleGroupsGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRoleGroupsIdDelete

> ApiRoleGroupsIdDelete(ctx, id).Execute()

Removes the RoleGroup resource.



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
    id := "id_example" // string | RoleGroup identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleGroupApi.ApiRoleGroupsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleGroupApi.ApiRoleGroupsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RoleGroup identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRoleGroupsIdDeleteRequest struct via the builder pattern


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


## ApiRoleGroupsIdGet

> RoleGroupJsonhal ApiRoleGroupsIdGet(ctx, id).Execute()

Retrieves a RoleGroup resource.



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
    id := "id_example" // string | RoleGroup identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleGroupApi.ApiRoleGroupsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleGroupApi.ApiRoleGroupsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRoleGroupsIdGet`: RoleGroupJsonhal
    fmt.Fprintf(os.Stdout, "Response from `RoleGroupApi.ApiRoleGroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RoleGroup identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRoleGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleGroupJsonhal**](RoleGroupJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRoleGroupsIdPatch

> RoleGroupJsonhal ApiRoleGroupsIdPatch(ctx, id).RoleGroup(roleGroup).Execute()

Updates the RoleGroup resource.



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
    id := "id_example" // string | RoleGroup identifier
    roleGroup := *openapiclient.NewRoleGroup() // RoleGroup | The updated RoleGroup resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleGroupApi.ApiRoleGroupsIdPatch(context.Background(), id).RoleGroup(roleGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleGroupApi.ApiRoleGroupsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRoleGroupsIdPatch`: RoleGroupJsonhal
    fmt.Fprintf(os.Stdout, "Response from `RoleGroupApi.ApiRoleGroupsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RoleGroup identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRoleGroupsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleGroup** | [**RoleGroup**](RoleGroup.md) | The updated RoleGroup resource | 

### Return type

[**RoleGroupJsonhal**](RoleGroupJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRoleGroupsIdPut

> RoleGroupJsonhal ApiRoleGroupsIdPut(ctx, id).RoleGroupJsonhal(roleGroupJsonhal).Execute()

Replaces the RoleGroup resource.



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
    id := "id_example" // string | RoleGroup identifier
    roleGroupJsonhal := *openapiclient.NewRoleGroupJsonhal() // RoleGroupJsonhal | The updated RoleGroup resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleGroupApi.ApiRoleGroupsIdPut(context.Background(), id).RoleGroupJsonhal(roleGroupJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleGroupApi.ApiRoleGroupsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRoleGroupsIdPut`: RoleGroupJsonhal
    fmt.Fprintf(os.Stdout, "Response from `RoleGroupApi.ApiRoleGroupsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RoleGroup identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRoleGroupsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleGroupJsonhal** | [**RoleGroupJsonhal**](RoleGroupJsonhal.md) | The updated RoleGroup resource | 

### Return type

[**RoleGroupJsonhal**](RoleGroupJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRoleGroupsPost

> RoleGroupJsonhal ApiRoleGroupsPost(ctx).RoleGroupJsonhal(roleGroupJsonhal).Execute()

Creates a RoleGroup resource.



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
    roleGroupJsonhal := *openapiclient.NewRoleGroupJsonhal() // RoleGroupJsonhal | The new RoleGroup resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleGroupApi.ApiRoleGroupsPost(context.Background()).RoleGroupJsonhal(roleGroupJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleGroupApi.ApiRoleGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiRoleGroupsPost`: RoleGroupJsonhal
    fmt.Fprintf(os.Stdout, "Response from `RoleGroupApi.ApiRoleGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRoleGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleGroupJsonhal** | [**RoleGroupJsonhal**](RoleGroupJsonhal.md) | The new RoleGroup resource | 

### Return type

[**RoleGroupJsonhal**](RoleGroupJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

