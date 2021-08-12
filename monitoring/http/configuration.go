package http

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

type MonitorConfiguration struct {
	// URL of the service
	URL string `mapstructure:"url"`
	// Timeout is the duration to wait for a ping to finish, such as "1s"
	Timeout time.Duration `mapstructure:"timeout"`
	// Interval is the interval to use when watching the target, such as "1s"
	Interval time.Duration `mapstructure:"interval"`
	// Method is the HTTP method to use, such as "GET"
	Method string `mapstructure:"method"`
	// FollowRedirects specifies whether or not redirects should be followed
	FollowRedirects bool `mapstructure:"followRedirects"`
	// MaximumRedirects specifies the maximum number of redirects to follow before throwing an error
	MaximumRedirects int `mapstructure:"maximumRedirects"`
	// Expect is the matching clauses to determine an alive service
	Expect struct {
		// Status is the expected HTTP status
		Status int `mapstructure:"status"`
	} `mapstructure:"expect"`
}

func ParseConfiguration(options map[string]interface{}) (*MonitorConfiguration, error) {
	config := &MonitorConfiguration{}

	// Decode the arbitrary options with support for time.Duration parsing
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook:       mapstructure.ComposeDecodeHookFunc(mapstructure.StringToTimeDurationHookFunc()),
		WeaklyTypedInput: true,
		Metadata:         nil,
		Result:           &config,
	})
	if err != nil {
		return nil, err
	}
	err = decoder.Decode(options)
	if err != nil {
		return nil, err
	}

	// Set default timeout if non is specified
	if config.Timeout.Seconds() == 0 {
		config.Timeout = time.Second
	}

	// Set default interval if non is specified
	if config.Interval.Seconds() == 0 {
		config.Interval = time.Second
	}

	// Set the default HTTP method
	if config.Method == "" {
		config.Method = "GET"
	}

	// Set the default maximum redirect count
	if config.MaximumRedirects == 0 {
		config.MaximumRedirects = 10
	}

	return config, nil
}

func (config *MonitorConfiguration) Validate() []error {
	errors := make([]error, 0)

	if config.URL == "" {
		errors = append(errors, fmt.Errorf("url cannot be empty"))
	}

	if config.Method != "GET" {
		errors = append(errors, fmt.Errorf("unsupported HTTP method"))
	}

	return errors
}
