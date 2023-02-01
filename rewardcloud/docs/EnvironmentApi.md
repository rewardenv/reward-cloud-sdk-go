# \EnvironmentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEnvironmentsGetCollection**](EnvironmentApi.md#ApiEnvironmentsGetCollection) | **Get** /api/environments | Retrieves the collection of Environment resources.
[**ApiEnvironmentsIdDelete**](EnvironmentApi.md#ApiEnvironmentsIdDelete) | **Delete** /api/environments/{id} | Removes the Environment resource.
[**ApiEnvironmentsIdGet**](EnvironmentApi.md#ApiEnvironmentsIdGet) | **Get** /api/environments/{id} | Retrieves a Environment resource.
[**ApiEnvironmentsIdPatch**](EnvironmentApi.md#ApiEnvironmentsIdPatch) | **Patch** /api/environments/{id} | Updates the Environment resource.
[**ApiEnvironmentsIdPut**](EnvironmentApi.md#ApiEnvironmentsIdPut) | **Put** /api/environments/{id} | Replaces the Environment resource.
[**ApiEnvironmentsIdcostsGet**](EnvironmentApi.md#ApiEnvironmentsIdcostsGet) | **Get** /api/environments/{id}/costs | Retrieves a Environment resource.
[**ApiEnvironmentsIdexportDatabasePut**](EnvironmentApi.md#ApiEnvironmentsIdexportDatabasePut) | **Put** /api/environments/{id}/export-database | Replaces the Environment resource.
[**ApiEnvironmentsIdexportMediaPut**](EnvironmentApi.md#ApiEnvironmentsIdexportMediaPut) | **Put** /api/environments/{id}/export-media | Replaces the Environment resource.
[**ApiEnvironmentsIdimportDatabasePost**](EnvironmentApi.md#ApiEnvironmentsIdimportDatabasePost) | **Post** /api/environments/{id}/import-database | Creates a Environment resource.
[**ApiEnvironmentsIdimportMediaPost**](EnvironmentApi.md#ApiEnvironmentsIdimportMediaPost) | **Post** /api/environments/{id}/import-media | Creates a Environment resource.
[**ApiEnvironmentsPost**](EnvironmentApi.md#ApiEnvironmentsPost) | **Post** /api/environments | Creates a Environment resource.



## ApiEnvironmentsGetCollection

> []EnvironmentEnvironmentGet ApiEnvironmentsGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).ProjectTeamId(projectTeamId).ProjectTeamId2(projectTeamId2).Id(id).Id2(id2).Project(project).Project2(project2).Provider(provider).Provider2(provider2).State(state).State2(state2).Region(region).Region2(region2).EnvironmentAccess(environmentAccess).EnvironmentAccess2(environmentAccess2).OrderUpdatedAt(orderUpdatedAt).Execute()

Retrieves the collection of Environment resources.



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
    projectTeamId := int32(56) // int32 |  (optional)
    projectTeamId2 := []int32{int32(123)} // []int32 |  (optional)
    id := int32(56) // int32 |  (optional)
    id2 := []int32{int32(123)} // []int32 |  (optional)
    project := "project_example" // string |  (optional)
    project2 := []string{"Inner_example"} // []string |  (optional)
    provider := "provider_example" // string |  (optional)
    provider2 := []string{"Inner_example"} // []string |  (optional)
    state := "state_example" // string |  (optional)
    state2 := []string{"Inner_example"} // []string |  (optional)
    region := "region_example" // string |  (optional)
    region2 := []string{"Inner_example"} // []string |  (optional)
    environmentAccess := "environmentAccess_example" // string |  (optional)
    environmentAccess2 := []string{"Inner_example"} // []string |  (optional)
    orderUpdatedAt := "orderUpdatedAt_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).ProjectTeamId(projectTeamId).ProjectTeamId2(projectTeamId2).Id(id).Id2(id2).Project(project).Project2(project2).Provider(provider).Provider2(provider2).State(state).State2(state2).Region(region).Region2(region2).EnvironmentAccess(environmentAccess).EnvironmentAccess2(environmentAccess2).OrderUpdatedAt(orderUpdatedAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsGetCollection`: []EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **projectTeamId** | **int32** |  | 
 **projectTeamId2** | **[]int32** |  | 
 **id** | **int32** |  | 
 **id2** | **[]int32** |  | 
 **project** | **string** |  | 
 **project2** | **[]string** |  | 
 **provider** | **string** |  | 
 **provider2** | **[]string** |  | 
 **state** | **string** |  | 
 **state2** | **[]string** |  | 
 **region** | **string** |  | 
 **region2** | **[]string** |  | 
 **environmentAccess** | **string** |  | 
 **environmentAccess2** | **[]string** |  | 
 **orderUpdatedAt** | **string** |  | 

### Return type

[**[]EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdDelete

> ApiEnvironmentsIdDelete(ctx, id).Execute()

Removes the Environment resource.



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
    id := "id_example" // string | Environment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdDeleteRequest struct via the builder pattern


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


## ApiEnvironmentsIdGet

> EnvironmentEnvironmentGet ApiEnvironmentsIdGet(ctx, id).Execute()

Retrieves a Environment resource.



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
    id := "id_example" // string | Environment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdGet`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdPatch

> EnvironmentEnvironmentGet ApiEnvironmentsIdPatch(ctx, id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()

Updates the Environment resource.



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
    id := "id_example" // string | Environment identifier
    environmentEnvironmentPost := *openapiclient.NewEnvironmentEnvironmentPost() // EnvironmentEnvironmentPost | The updated Environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdPatch(context.Background(), id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdPatch`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvironmentPost** | [**EnvironmentEnvironmentPost**](EnvironmentEnvironmentPost.md) | The updated Environment resource | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdPut

> EnvironmentEnvironmentGet ApiEnvironmentsIdPut(ctx, id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()

Replaces the Environment resource.



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
    id := "id_example" // string | Environment identifier
    environmentEnvironmentPost := *openapiclient.NewEnvironmentEnvironmentPost() // EnvironmentEnvironmentPost | The updated Environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdPut(context.Background(), id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdPut`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvironmentPost** | [**EnvironmentEnvironmentPost**](EnvironmentEnvironmentPost.md) | The updated Environment resource | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdcostsGet

> EnvironmentEnvironmentGet ApiEnvironmentsIdcostsGet(ctx, id).Execute()

Retrieves a Environment resource.



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
    id := "id_example" // string | Environment identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdcostsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdcostsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdcostsGet`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdcostsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdcostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdexportDatabasePut

> EnvironmentEnvironmentGet ApiEnvironmentsIdexportDatabasePut(ctx, id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()

Replaces the Environment resource.



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
    id := "id_example" // string | Environment identifier
    environmentEnvironmentPost := *openapiclient.NewEnvironmentEnvironmentPost() // EnvironmentEnvironmentPost | The updated Environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdexportDatabasePut(context.Background(), id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdexportDatabasePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdexportDatabasePut`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdexportDatabasePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdexportDatabasePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvironmentPost** | [**EnvironmentEnvironmentPost**](EnvironmentEnvironmentPost.md) | The updated Environment resource | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdexportMediaPut

> EnvironmentEnvironmentGet ApiEnvironmentsIdexportMediaPut(ctx, id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()

Replaces the Environment resource.



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
    id := "id_example" // string | Environment identifier
    environmentEnvironmentPost := *openapiclient.NewEnvironmentEnvironmentPost() // EnvironmentEnvironmentPost | The updated Environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdexportMediaPut(context.Background(), id).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdexportMediaPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdexportMediaPut`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdexportMediaPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdexportMediaPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentEnvironmentPost** | [**EnvironmentEnvironmentPost**](EnvironmentEnvironmentPost.md) | The updated Environment resource | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdimportDatabasePost

> EnvironmentEnvironmentGet ApiEnvironmentsIdimportDatabasePost(ctx, id).Name(name).Cpu(cpu).Memory(memory).Storage(storage).DataTransferSettings(dataTransferSettings).IsStripDatabase(isStripDatabase).IsAllowOutgoingEmails(isAllowOutgoingEmails).IsInitSampleData(isInitSampleData).Project(project).Provider(provider).EnvVar(envVar).State(state).EnvironmentComponent(environmentComponent).Region(region).ExportedData(exportedData).ImportedData(importedData).Execute()

Creates a Environment resource.



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
    id := "id_example" // string | Environment identifier
    name := "name_example" // string |  (optional)
    cpu := int32(56) // int32 |  (optional)
    memory := int32(56) // int32 |  (optional)
    storage := int32(56) // int32 |  (optional)
    dataTransferSettings := "dataTransferSettings_example" // string |  (optional)
    isStripDatabase := true // bool |  (optional)
    isAllowOutgoingEmails := true // bool |  (optional)
    isInitSampleData := true // bool |  (optional)
    project := "project_example" // string |  (optional)
    provider := "provider_example" // string |  (optional)
    envVar := []openapiclient.EnvironmentEnvVarEnvironmentPost{*openapiclient.NewEnvironmentEnvVarEnvironmentPost()} // []EnvironmentEnvVarEnvironmentPost |  (optional)
    state := "state_example" // string |  (optional)
    environmentComponent := []openapiclient.EnvironmentComponentEnvironmentPost{*openapiclient.NewEnvironmentComponentEnvironmentPost()} // []EnvironmentComponentEnvironmentPost |  (optional)
    region := "region_example" // string |  (optional)
    exportedData := []string{"Inner_example"} // []string |  (optional)
    importedData := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdimportDatabasePost(context.Background(), id).Name(name).Cpu(cpu).Memory(memory).Storage(storage).DataTransferSettings(dataTransferSettings).IsStripDatabase(isStripDatabase).IsAllowOutgoingEmails(isAllowOutgoingEmails).IsInitSampleData(isInitSampleData).Project(project).Provider(provider).EnvVar(envVar).State(state).EnvironmentComponent(environmentComponent).Region(region).ExportedData(exportedData).ImportedData(importedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdimportDatabasePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdimportDatabasePost`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdimportDatabasePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdimportDatabasePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **cpu** | **int32** |  | 
 **memory** | **int32** |  | 
 **storage** | **int32** |  | 
 **dataTransferSettings** | **string** |  | 
 **isStripDatabase** | **bool** |  | 
 **isAllowOutgoingEmails** | **bool** |  | 
 **isInitSampleData** | **bool** |  | 
 **project** | **string** |  | 
 **provider** | **string** |  | 
 **envVar** | [**[]EnvironmentEnvVarEnvironmentPost**](EnvironmentEnvVarEnvironmentPost.md) |  | 
 **state** | **string** |  | 
 **environmentComponent** | [**[]EnvironmentComponentEnvironmentPost**](EnvironmentComponentEnvironmentPost.md) |  | 
 **region** | **string** |  | 
 **exportedData** | **[]string** |  | 
 **importedData** | **[]string** |  | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsIdimportMediaPost

> EnvironmentEnvironmentGet ApiEnvironmentsIdimportMediaPost(ctx, id).Name(name).Cpu(cpu).Memory(memory).Storage(storage).DataTransferSettings(dataTransferSettings).IsStripDatabase(isStripDatabase).IsAllowOutgoingEmails(isAllowOutgoingEmails).IsInitSampleData(isInitSampleData).Project(project).Provider(provider).EnvVar(envVar).State(state).EnvironmentComponent(environmentComponent).Region(region).ExportedData(exportedData).ImportedData(importedData).Execute()

Creates a Environment resource.



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
    id := "id_example" // string | Environment identifier
    name := "name_example" // string |  (optional)
    cpu := int32(56) // int32 |  (optional)
    memory := int32(56) // int32 |  (optional)
    storage := int32(56) // int32 |  (optional)
    dataTransferSettings := "dataTransferSettings_example" // string |  (optional)
    isStripDatabase := true // bool |  (optional)
    isAllowOutgoingEmails := true // bool |  (optional)
    isInitSampleData := true // bool |  (optional)
    project := "project_example" // string |  (optional)
    provider := "provider_example" // string |  (optional)
    envVar := []openapiclient.EnvironmentEnvVarEnvironmentPost{*openapiclient.NewEnvironmentEnvVarEnvironmentPost()} // []EnvironmentEnvVarEnvironmentPost |  (optional)
    state := "state_example" // string |  (optional)
    environmentComponent := []openapiclient.EnvironmentComponentEnvironmentPost{*openapiclient.NewEnvironmentComponentEnvironmentPost()} // []EnvironmentComponentEnvironmentPost |  (optional)
    region := "region_example" // string |  (optional)
    exportedData := []string{"Inner_example"} // []string |  (optional)
    importedData := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsIdimportMediaPost(context.Background(), id).Name(name).Cpu(cpu).Memory(memory).Storage(storage).DataTransferSettings(dataTransferSettings).IsStripDatabase(isStripDatabase).IsAllowOutgoingEmails(isAllowOutgoingEmails).IsInitSampleData(isInitSampleData).Project(project).Provider(provider).EnvVar(envVar).State(state).EnvironmentComponent(environmentComponent).Region(region).ExportedData(exportedData).ImportedData(importedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsIdimportMediaPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsIdimportMediaPost`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsIdimportMediaPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsIdimportMediaPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **cpu** | **int32** |  | 
 **memory** | **int32** |  | 
 **storage** | **int32** |  | 
 **dataTransferSettings** | **string** |  | 
 **isStripDatabase** | **bool** |  | 
 **isAllowOutgoingEmails** | **bool** |  | 
 **isInitSampleData** | **bool** |  | 
 **project** | **string** |  | 
 **provider** | **string** |  | 
 **envVar** | [**[]EnvironmentEnvVarEnvironmentPost**](EnvironmentEnvVarEnvironmentPost.md) |  | 
 **state** | **string** |  | 
 **environmentComponent** | [**[]EnvironmentComponentEnvironmentPost**](EnvironmentComponentEnvironmentPost.md) |  | 
 **region** | **string** |  | 
 **exportedData** | **[]string** |  | 
 **importedData** | **[]string** |  | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEnvironmentsPost

> EnvironmentEnvironmentGet ApiEnvironmentsPost(ctx).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()

Creates a Environment resource.



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
    environmentEnvironmentPost := *openapiclient.NewEnvironmentEnvironmentPost() // EnvironmentEnvironmentPost | The new Environment resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentApi.ApiEnvironmentsPost(context.Background()).EnvironmentEnvironmentPost(environmentEnvironmentPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ApiEnvironmentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEnvironmentsPost`: EnvironmentEnvironmentGet
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ApiEnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentEnvironmentPost** | [**EnvironmentEnvironmentPost**](EnvironmentEnvironmentPost.md) | The new Environment resource | 

### Return type

[**EnvironmentEnvironmentGet**](EnvironmentEnvironmentGet.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json, application/hal+json, text/html
- **Accept**: application/json, application/hal+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

