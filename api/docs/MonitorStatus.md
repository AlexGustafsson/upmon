# MonitorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Up** | **int32** | The number of cluster members voting for up | 
**Down** | **int32** | The number of cluster members voting for down | 
**TransitioningUp** | **int32** | The number of cluster members voting for transitioning up | 
**TransitioningDown** | **int32** | The number of cluster members voting for transitioning down | 
**Unknown** | **int32** | The number of cluster members voting for unknown | 

## Methods

### NewMonitorStatus

`func NewMonitorStatus(up int32, down int32, transitioningUp int32, transitioningDown int32, unknown int32, ) *MonitorStatus`

NewMonitorStatus instantiates a new MonitorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorStatusWithDefaults

`func NewMonitorStatusWithDefaults() *MonitorStatus`

NewMonitorStatusWithDefaults instantiates a new MonitorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUp

`func (o *MonitorStatus) GetUp() int32`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MonitorStatus) GetUpOk() (*int32, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MonitorStatus) SetUp(v int32)`

SetUp sets Up field to given value.


### GetDown

`func (o *MonitorStatus) GetDown() int32`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *MonitorStatus) GetDownOk() (*int32, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *MonitorStatus) SetDown(v int32)`

SetDown sets Down field to given value.


### GetTransitioningUp

`func (o *MonitorStatus) GetTransitioningUp() int32`

GetTransitioningUp returns the TransitioningUp field if non-nil, zero value otherwise.

### GetTransitioningUpOk

`func (o *MonitorStatus) GetTransitioningUpOk() (*int32, bool)`

GetTransitioningUpOk returns a tuple with the TransitioningUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitioningUp

`func (o *MonitorStatus) SetTransitioningUp(v int32)`

SetTransitioningUp sets TransitioningUp field to given value.


### GetTransitioningDown

`func (o *MonitorStatus) GetTransitioningDown() int32`

GetTransitioningDown returns the TransitioningDown field if non-nil, zero value otherwise.

### GetTransitioningDownOk

`func (o *MonitorStatus) GetTransitioningDownOk() (*int32, bool)`

GetTransitioningDownOk returns a tuple with the TransitioningDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitioningDown

`func (o *MonitorStatus) SetTransitioningDown(v int32)`

SetTransitioningDown sets TransitioningDown field to given value.


### GetUnknown

`func (o *MonitorStatus) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *MonitorStatus) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *MonitorStatus) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


