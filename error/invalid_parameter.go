package error

import "fmt"

type InvalidParameterError struct {
	Parameter string
}

func (err InvalidParameterError) Error() string {
	return fmt.Sprintf("you used for the parameter '%s' a wrong syntax", err.Parameter)
}
