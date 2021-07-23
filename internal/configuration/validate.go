package configuration

import "fmt"

// Validate a configuration
func Validate(config *Configuration) []error {
	errors := make([]error, 0)

	if config.Name == "" {
		errors = append(errors, fmt.Errorf("name cannot be empty"))
	}

	return errors
}