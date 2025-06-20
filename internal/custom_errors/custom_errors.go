package custom_errors

import "fmt"

type NotFoundError struct {
	Resource string
	ID       uint
	Message  string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s (ID: %d)", e.Resource, e.Message, e.ID)
}

func NewNotFoundError(resource string, id uint, message string) NotFoundError {
	return NotFoundError{
		Resource: resource,
		ID:       id,
		Message:  message,
	}
}
