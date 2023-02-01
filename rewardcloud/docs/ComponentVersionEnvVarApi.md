# \ComponentVersionEnvVarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiComponentVersionEnvVarsGetCollection**](ComponentVersionEnvVarApi.md#ApiComponentVersionEnvVarsGetCollection) | **Get** /api/component_version_env_vars | Retrieves the collection of ComponentVersionEnvVar resources.
[**ApiComponentVersionEnvVarsIdDelete**](ComponentVersionEnvVarApi.md#ApiComponentVersionEnvVarsIdDelete) | **Delete** /api/component_version_env_vars/{id} | Removes the ComponentVersionEnvVar resource.
[**ApiComponentVersionEnvVarsIdGet**](ComponentVersionEnvVarApi.md#ApiComponentVersionEnvVarsIdGet) | **Get** /api/component_version_env_vars/{id} | Retrieves a ComponentVersionEnvVar resource.
[**ApiComponentVersionEnvVarsIdPatch**](ComponentVersionEnvVarApi.md#ApiComponentVersionEnvVarsIdPatch) | **Patch** /api/component_version_env_vars/{id} | Updates the ComponentVersionEnvVar resource.
[**ApiComponentVersionEnvVarsIdPut**](ComponentVersionEnvVarApi.md#ApiComponentVersionEnvVarsIdPut) | **Put** /api/component_version_env_vars/{id} | Replaces the ComponentVersionEnvVar resource.
[**ApiComponentVersionEnvVarsPost**](ComponentVersionEnvVarApi.md#ApiComponentVersionEnvVarsPost) | **Post** /api/component_version_env_vars | Creates a ComponentVersionEnvVar resource.



## ApiComponentVersionEnvVarsGetCollection

> []ComponentVersionEnvVar ApiComponentVersionEnvVarsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ComponentVersion(componentVersion).ComponentVersion2(componentVersion2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()

Retrieves the collection of ComponentVersionEnvVar resources.



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
    envVarType := "envVarType_example" // string |  (optional)
    envVarType2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ComponentVersion(componentVersion).ComponentVersion2(componentVersion2).EnvVarType(envVarType).EnvVarType2(envVarType2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarsGetCollection`: []ComponentVersionEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **componentVersion** | **string** |  | 
 **componentVersion2** | **[]string** |  | 
 **envVarType** | **string** |  | 
 **envVarType2** | **[]string** |  | 

### Return type

[**[]ComponentVersionEnvVar**](ComponentVersionEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarsIdDelete

> ApiComponentVersionEnvVarsIdDelete(ctx, id).Execute()

Removes the ComponentVersionEnvVar resource.



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
    id := "id_example" // string | ComponentVersionEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarsIdDeleteRequest struct via the builder pattern


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


## ApiComponentVersionEnvVarsIdGet

> ComponentVersionEnvVar ApiComponentVersionEnvVarsIdGet(ctx, id).Execute()

Retrieves a ComponentVersionEnvVar resource.



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
    id := "id_example" // string | ComponentVersionEnvVar identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarsIdGet`: ComponentVersionEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentVersionEnvVar**](ComponentVersionEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarsIdPatch

> ComponentVersionEnvVar ApiComponentVersionEnvVarsIdPatch(ctx, id).ComponentVersionEnvVar(componentVersionEnvVar).Execute()

Updates the ComponentVersionEnvVar resource.



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
    id := "id_example" // string | ComponentVersionEnvVar identifier
    componentVersionEnvVar := *openapiclient.NewComponentVersionEnvVar() // ComponentVersionEnvVar | The updated ComponentVersionEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdPatch(context.Background(), id).ComponentVersionEnvVar(componentVersionEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarsIdPatch`: ComponentVersionEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentVersionEnvVar** | [**ComponentVersionEnvVar**](ComponentVersionEnvVar.md) | The updated ComponentVersionEnvVar resource | 

### Return type

[**ComponentVersionEnvVar**](ComponentVersionEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarsIdPut

> ComponentVersionEnvVar ApiComponentVersionEnvVarsIdPut(ctx, id).ComponentVersionEnvVar(componentVersionEnvVar).Execute()

Replaces the ComponentVersionEnvVar resource.



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
    id := "id_example" // string | ComponentVersionEnvVar identifier
    componentVersionEnvVar := *openapiclient.NewComponentVersionEnvVar() // ComponentVersionEnvVar | The updated ComponentVersionEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdPut(context.Background(), id).ComponentVersionEnvVar(componentVersionEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarsIdPut`: ComponentVersionEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ComponentVersionEnvVar identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentVersionEnvVar** | [**ComponentVersionEnvVar**](ComponentVersionEnvVar.md) | The updated ComponentVersionEnvVar resource | 

### Return type

[**ComponentVersionEnvVar**](ComponentVersionEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiComponentVersionEnvVarsPost

> ComponentVersionEnvVar ApiComponentVersionEnvVarsPost(ctx).ComponentVersionEnvVar(componentVersionEnvVar).Execute()

Creates a ComponentVersionEnvVar resource.



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
    componentVersionEnvVar := *openapiclient.NewComponentVersionEnvVar() // ComponentVersionEnvVar | The new ComponentVersionEnvVar resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsPost(context.Background()).ComponentVersionEnvVar(componentVersionEnvVar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiComponentVersionEnvVarsPost`: ComponentVersionEnvVar
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionEnvVarApi.ApiComponentVersionEnvVarsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiComponentVersionEnvVarsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **componentVersionEnvVar** | [**ComponentVersionEnvVar**](ComponentVersionEnvVar.md) | The new ComponentVersionEnvVar resource | 

### Return type

[**ComponentVersionEnvVar**](ComponentVersionEnvVar.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

