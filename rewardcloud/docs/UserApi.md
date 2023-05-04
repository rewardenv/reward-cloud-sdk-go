# \UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUsersGetCollection**](UserApi.md#ApiUsersGetCollection) | **Get** /api/users | Retrieves the collection of User resources.
[**ApiUsersIdDelete**](UserApi.md#ApiUsersIdDelete) | **Delete** /api/users/{id} | Removes the User resource.
[**ApiUsersIdGet**](UserApi.md#ApiUsersIdGet) | **Get** /api/users/{id} | Retrieves a User resource.
[**ApiUsersIdPatch**](UserApi.md#ApiUsersIdPatch) | **Patch** /api/users/{id} | Updates the User resource.
[**ApiUsersIdPut**](UserApi.md#ApiUsersIdPut) | **Put** /api/users/{id} | Replaces the User resource.
[**ApiUsersIdadministrateGet**](UserApi.md#ApiUsersIdadministrateGet) | **Get** /api/users/{id}/administrate | Retrieves a User resource.
[**ApiUsersIdfinancecostsGet**](UserApi.md#ApiUsersIdfinancecostsGet) | **Get** /api/users/{id}/finance/costs | Retrieves a User resource.
[**ApiUsersIdfinanceinvoicesGet**](UserApi.md#ApiUsersIdfinanceinvoicesGet) | **Get** /api/users/{id}/finance/invoices | Retrieves a User resource.
[**ApiUsersIdfinancepayGet**](UserApi.md#ApiUsersIdfinancepayGet) | **Get** /api/users/{id}/finance/pay | Retrieves a User resource.
[**ApiUsersIdfinancesubscribeInfoGet**](UserApi.md#ApiUsersIdfinancesubscribeInfoGet) | **Get** /api/users/{id}/finance/subscribe-info | Retrieves a User resource.
[**ApiUsersIdremoveAccountDelete**](UserApi.md#ApiUsersIdremoveAccountDelete) | **Delete** /api/users/{id}/remove-account | Removes the User resource.
[**ApiUsersPost**](UserApi.md#ApiUsersPost) | **Post** /api/users | Creates a User resource.
[**ApiUserscheckEmailGet**](UserApi.md#ApiUserscheckEmailGet) | **Get** /api/users/check-email | It will check the email is exist in database or not
[**ApiUserscheckUsernameGet**](UserApi.md#ApiUserscheckUsernameGet) | **Get** /api/users/check-username | It will check the username is exist in database or not
[**ApiUsersdiscoverMercureGet**](UserApi.md#ApiUsersdiscoverMercureGet) | **Get** /api/users/discover_mercure | It will discovers for mercure JWT token for you - w/ take consideration your permissions
[**ApiUsersregisterPost**](UserApi.md#ApiUsersregisterPost) | **Post** /api/users/register | Creates a User resource.



## ApiUsersGetCollection

> ApiUsersGetCollection200Response ApiUsersGetCollection(ctx).Page(page).ItemsPerPage(itemsPerPage).RoleGroup(roleGroup).RoleGroup2(roleGroup2).Execute()

Retrieves the collection of User resources.



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
    roleGroup := "roleGroup_example" // string |  (optional)
    roleGroup2 := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersGetCollection(context.Background()).Page(page).ItemsPerPage(itemsPerPage).RoleGroup(roleGroup).RoleGroup2(roleGroup2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersGetCollection`: ApiUsersGetCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **itemsPerPage** | **int32** | The number of items per page | [default to 30]
 **roleGroup** | **string** |  | 
 **roleGroup2** | **[]string** |  | 

### Return type

[**ApiUsersGetCollection200Response**](ApiUsersGetCollection200Response.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdDelete

> ApiUsersIdDelete(ctx, id).Execute()

Removes the User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdDeleteRequest struct via the builder pattern


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


## ApiUsersIdGet

> UserJsonhal ApiUsersIdGet(ctx, id).Execute()

Retrieves a User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdPatch

> UserJsonhal ApiUsersIdPatch(ctx, id).User(user).Execute()

Updates the User resource.



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
    id := "id_example" // string | User identifier
    user := *openapiclient.NewUser() // User | The updated User resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdPatch(context.Background(), id).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdPatch`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | The updated User resource | 

### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdPut

> UserJsonhal ApiUsersIdPut(ctx, id).UserJsonhal(userJsonhal).Execute()

Replaces the User resource.



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
    id := "id_example" // string | User identifier
    userJsonhal := *openapiclient.NewUserJsonhal() // UserJsonhal | The updated User resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdPut(context.Background(), id).UserJsonhal(userJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdPut`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userJsonhal** | [**UserJsonhal**](UserJsonhal.md) | The updated User resource | 

### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdadministrateGet

> UserJsonhal ApiUsersIdadministrateGet(ctx, id).Execute()

Retrieves a User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdadministrateGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdadministrateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdadministrateGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdadministrateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdadministrateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdfinancecostsGet

> UserJsonhal ApiUsersIdfinancecostsGet(ctx, id).Execute()

Retrieves a User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdfinancecostsGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdfinancecostsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdfinancecostsGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdfinancecostsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdfinancecostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdfinanceinvoicesGet

> UserJsonhal ApiUsersIdfinanceinvoicesGet(ctx, id).Execute()

Retrieves a User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdfinanceinvoicesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdfinanceinvoicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdfinanceinvoicesGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdfinanceinvoicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdfinanceinvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdfinancepayGet

> UserJsonhal ApiUsersIdfinancepayGet(ctx, id).Execute()

Retrieves a User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdfinancepayGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdfinancepayGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdfinancepayGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdfinancepayGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdfinancepayGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdfinancesubscribeInfoGet

> UserJsonhal ApiUsersIdfinancesubscribeInfoGet(ctx, id).Execute()

Retrieves a User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdfinancesubscribeInfoGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdfinancesubscribeInfoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersIdfinancesubscribeInfoGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersIdfinancesubscribeInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdfinancesubscribeInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersIdremoveAccountDelete

> ApiUsersIdremoveAccountDelete(ctx, id).Execute()

Removes the User resource.



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
    id := "id_example" // string | User identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersIdremoveAccountDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersIdremoveAccountDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersIdremoveAccountDeleteRequest struct via the builder pattern


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


## ApiUsersPost

> UserJsonhal ApiUsersPost(ctx).UserJsonhal(userJsonhal).Execute()

Creates a User resource.



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
    userJsonhal := *openapiclient.NewUserJsonhal() // UserJsonhal | The new User resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersPost(context.Background()).UserJsonhal(userJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersPost`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userJsonhal** | [**UserJsonhal**](UserJsonhal.md) | The new User resource | 

### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserscheckEmailGet

> UserJsonhal ApiUserscheckEmailGet(ctx).Execute()

It will check the email is exist in database or not



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUserscheckEmailGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUserscheckEmailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserscheckEmailGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUserscheckEmailGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserscheckEmailGetRequest struct via the builder pattern


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUserscheckUsernameGet

> UserJsonhal ApiUserscheckUsernameGet(ctx).Execute()

It will check the username is exist in database or not



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUserscheckUsernameGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUserscheckUsernameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUserscheckUsernameGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUserscheckUsernameGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiUserscheckUsernameGetRequest struct via the builder pattern


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersdiscoverMercureGet

> UserJsonhal ApiUsersdiscoverMercureGet(ctx).Execute()

It will discovers for mercure JWT token for you - w/ take consideration your permissions



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersdiscoverMercureGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersdiscoverMercureGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersdiscoverMercureGet`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersdiscoverMercureGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersdiscoverMercureGetRequest struct via the builder pattern


### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiUsersregisterPost

> UserJsonhal ApiUsersregisterPost(ctx).UserJsonhal(userJsonhal).Execute()

Creates a User resource.



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
    userJsonhal := *openapiclient.NewUserJsonhal() // UserJsonhal | The new User resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ApiUsersregisterPost(context.Background()).UserJsonhal(userJsonhal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ApiUsersregisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiUsersregisterPost`: UserJsonhal
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ApiUsersregisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiUsersregisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userJsonhal** | [**UserJsonhal**](UserJsonhal.md) | The new User resource | 

### Return type

[**UserJsonhal**](UserJsonhal.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html
- **Accept**: application/hal+json, application/vnd.api+json, application/json, application/xml, text/xml, application/x-yaml, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

