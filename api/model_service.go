/*
 * upmon
 *
 * A cloud-native, distributed uptime monitor written in Go
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Service struct for Service
type Service struct {
	// An identifier for the service, unique for the origin
	Id string `json:"id"`
	// Name of the service
	Name string `json:"name"`
	// Description of the service
	Description *string `json:"description,omitempty"`
	// Whether or not the config is shared with the cluster
	Private bool `json:"private"`
	// The current status of the service
	Status string `json:"status"`
	// The timestamp at which the service was last seen responding
	LastSeen string `json:"lastSeen"`
	// The origin node from which this service is configured
	Origin string `json:"origin"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(id string, name string, private bool, status string, lastSeen string, origin string) *Service {
	this := Service{}
	this.Id = id
	this.Name = name
	this.Private = private
	this.Status = status
	this.LastSeen = lastSeen
	this.Origin = origin
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetId returns the Id field value
func (o *Service) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Service) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Service) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Service) SetDescription(v string) {
	o.Description = &v
}

// GetPrivate returns the Private field value
func (o *Service) GetPrivate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Private
}

// GetPrivateOk returns a tuple with the Private field value
// and a boolean to check if the value has been set.
func (o *Service) GetPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Private, true
}

// SetPrivate sets field value
func (o *Service) SetPrivate(v bool) {
	o.Private = v
}

// GetStatus returns the Status field value
func (o *Service) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Service) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Service) SetStatus(v string) {
	o.Status = v
}

// GetLastSeen returns the LastSeen field value
func (o *Service) GetLastSeen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *Service) GetLastSeenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value
func (o *Service) SetLastSeen(v string) {
	o.LastSeen = v
}

// GetOrigin returns the Origin field value
func (o *Service) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *Service) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *Service) SetOrigin(v string) {
	o.Origin = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["private"] = o.Private
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if true {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
