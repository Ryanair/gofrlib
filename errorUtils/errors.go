package errorUtils

import (
	"errors"
	"strings"
)

func MergeErrors(errs []error) error {
	var groupErrorMessage []string
	for _, err := range errs {
		if err != nil {
			groupErrorMessage = append(groupErrorMessage, err.Error())
		}
	}

	if len(groupErrorMessage) == 0 {
		return nil
	}
	return errors.New(strings.Join(groupErrorMessage, "\n"))
}
