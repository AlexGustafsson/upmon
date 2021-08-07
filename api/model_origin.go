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

// Origin struct for Origin
type Origin struct {
	// A globally unique identifier for the origin
	Id string `json:"id"`
}

// NewOrigin instantiates a new Origin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrigin(id string) *Origin {
	this := Origin{}
	this.Id = id
	return &this
}

// NewOriginWithDefaults instantiates a new Origin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginWithDefaults() *Origin {
	this := Origin{}
	return &this
}

// GetId returns the Id field value
func (o *Origin) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Origin) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Origin) SetId(v string) {
	o.Id = v
}

func (o Origin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrigin struct {
	value *Origin
	isSet bool
}

func (v NullableOrigin) Get() *Origin {
	return v.value
}

func (v *NullableOrigin) Set(val *Origin) {
	v.value = val
	v.isSet = true
}

func (v NullableOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrigin(val *Origin) *NullableOrigin {
	return &NullableOrigin{value: val, isSet: true}
}

func (v NullableOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
