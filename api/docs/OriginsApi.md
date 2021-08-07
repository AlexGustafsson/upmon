# \OriginsApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OriginsGet**](OriginsApi.md#OriginsGet) | **Get** /origins | Retrieve all origins
[**OriginsOriginIdGet**](OriginsApi.md#OriginsOriginIdGet) | **Get** /origins/{originId} | Retrieve an origin



## OriginsGet

> Origins OriginsGet(ctx).Execute()

Retrieve all origins

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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.OriginsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.OriginsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsGet`: Origins
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.OriginsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsGetRequest struct via the builder pattern


### Return type

[**Origins**](Origins.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OriginsOriginIdGet

> Origin OriginsOriginIdGet(ctx, originId).Execute()

Retrieve an origin

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
    originId := "originId_example" // string | The id of the target origin

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OriginsApi.OriginsOriginIdGet(context.Background(), originId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OriginsApi.OriginsOriginIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OriginsOriginIdGet`: Origin
    fmt.Fprintf(os.Stdout, "Response from `OriginsApi.OriginsOriginIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | The id of the target origin | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginsOriginIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Origin**](Origin.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

