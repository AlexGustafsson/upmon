# Peer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the peer | 
**Bind** | **string** | The address and port of the peer | 
**Status** | **string** | The status of the peer | 

## Methods

### NewPeer

`func NewPeer(name string, bind string, status string, ) *Peer`

NewPeer instantiates a new Peer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeerWithDefaults

`func NewPeerWithDefaults() *Peer`

NewPeerWithDefaults instantiates a new Peer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Peer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Peer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Peer) SetName(v string)`

SetName sets Name field to given value.


### GetBind

`func (o *Peer) GetBind() string`

GetBind returns the Bind field if non-nil, zero value otherwise.

### GetBindOk

`func (o *Peer) GetBindOk() (*string, bool)`

GetBindOk returns a tuple with the Bind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBind

`func (o *Peer) SetBind(v string)`

SetBind sets Bind field to given value.


### GetStatus

`func (o *Peer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Peer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Peer) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


