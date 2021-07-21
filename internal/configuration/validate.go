package configuration

import "fmt"

func Validate(config *Configuration) []error {
	errors := make([]error, 0)

	if config.Name == "" {
		errors = append(errors, fmt.Errorf("name cannot be empty"))
	}

	return errors
}
