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

// Monitor struct for Monitor
type Monitor struct {
	// An identifier for the monitor, unique for the service
	Id string `json:"id"`
	// Name of the monitor
	Name string `json:"name"`
	// Description of the monitor
	Description string `json:"description"`
	// Type of the monitor
	Type string `json:"type"`
	// The id of the parent service
	Service string        `json:"service"`
	Status  MonitorStatus `json:"status"`
}

// NewMonitor instantiates a new Monitor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitor(id string, name string, description string, type_ string, service string, status MonitorStatus) *Monitor {
	this := Monitor{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Service = service
	this.Status = status
	return &this
}

// NewMonitorWithDefaults instantiates a new Monitor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorWithDefaults() *Monitor {
	this := Monitor{}
	return &this
}

// GetId returns the Id field value
func (o *Monitor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Monitor) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Monitor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Monitor) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Monitor) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Monitor) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *Monitor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Monitor) SetType(v string) {
	o.Type = v
}

// GetService returns the Service field value
func (o *Monitor) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *Monitor) SetService(v string) {
	o.Service = v
}

// GetStatus returns the Status field value
func (o *Monitor) GetStatus() MonitorStatus {
	if o == nil {
		var ret MonitorStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Monitor) GetStatusOk() (*MonitorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Monitor) SetStatus(v MonitorStatus) {
	o.Status = v
}

func (o Monitor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMonitor struct {
	value *Monitor
	isSet bool
}

func (v NullableMonitor) Get() *Monitor {
	return v.value
}

func (v *NullableMonitor) Set(val *Monitor) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitor) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitor(val *Monitor) *NullableMonitor {
	return &NullableMonitor{value: val, isSet: true}
}

func (v NullableMonitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}