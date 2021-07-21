# \PeersApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeersGet**](PeersApi.md#PeersGet) | **Get** /peers | Retrieve all peers
[**PeersPeerIdGet**](PeersApi.md#PeersPeerIdGet) | **Get** /peers/{peerId} | Retrieve a peer



## PeersGet

> Peers PeersGet(ctx).Execute()

Retrieve all peers

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
    resp, r, err := api_client.PeersApi.PeersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeersApi.PeersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PeersGet`: Peers
    fmt.Fprintf(os.Stdout, "Response from `PeersApi.PeersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPeersGetRequest struct via the builder pattern


### Return type

[**Peers**](Peers.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeersPeerIdGet

> Peer PeersPeerIdGet(ctx, peerId).Execute()

Retrieve a peer

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
    peerId := "peerId_example" // string | The id of the target peer

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PeersApi.PeersPeerIdGet(context.Background(), peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeersApi.PeersPeerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PeersPeerIdGet`: Peer
    fmt.Fprintf(os.Stdout, "Response from `PeersApi.PeersPeerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerId** | **string** | The id of the target peer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPeersPeerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Peer**](Peer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

