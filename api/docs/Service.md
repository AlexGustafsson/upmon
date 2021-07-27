# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the service | 
**Description** | Pointer to **string** | Description of the service | [optional] 
**Private** | **bool** | Whether or not the config is shared with the cluster | 
**Status** | **string** | The current status of the service | 
**LastSeen** | **string** | The timestamp at which the service was last seen responding | 

## Methods

### NewService

`func NewService(name string, private bool, status string, lastSeen string, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Service) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivate

`func (o *Service) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Service) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Service) SetPrivate(v bool)`

SetPrivate sets Private field to given value.


### GetStatus

`func (o *Service) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Service) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Service) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSeen

`func (o *Service) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Service) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Service) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


