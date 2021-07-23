package ping

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

type MonitorConfiguration struct {
	// hostname of the service
	Hostname string `mapstructure:"hostname"`
	// count is the number of pings to send
	Count int `mapstructure:"count"`
	// timeout is the duration to wait for a ping to finish, such as "1s"
	Timeout time.Duration `mapstructure:"timeout"`
	// interval is the interval to use when watching the target, such as "1s"
	Interval time.Duration `mapstructure:"interval"`
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

	// Set default count if non is specified
	if config.Count == 0 {
		config.Count = 1
	}

	// Set default timeout if non is specified
	if config.Timeout.Seconds() == 0 {
		config.Timeout = time.Second
	}

	// Set default interval if non is specified
	if config.Interval.Seconds() == 0 {
		config.Interval = time.Second
	}

	return config, nil
}

func (config *MonitorConfiguration) Validate() []error {
	errors := make([]error, 0)

	if config.Hostname == "" {
		errors = append(errors, fmt.Errorf("hostname cannot be empty"))
	}

	return errors
}
