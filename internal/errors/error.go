package errors

import "strings"

// Error represents a custom error type that encapsulates an error message.
type Error struct {
	Message string `json:"message,omitempty"`
}

func New(messages ...string) *Error {
	message := strings.Join(messages, ": ")
	return &Error{
		Message: message,
	}
}

// Error returns the error message as a string representation of the Error struct.
func (e *Error) Error() string {
	return e.Message
}

// NotSaved creates a new Error instance with a default message indicating that the entity failed to save.
func NotSaved(messages ...string) *Error {
	message := "failed to save the entity"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}

// NotDeleted creates a new Error instance with a default message indicating that the entity failed to delete.
func NotDeleted(messages ...string) *Error {
	message := "failed to delete the entity"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}

// NotFound creates a new Error instance with a default message indicating that the entity was not found.
func NotFound(messages ...string) *Error {
	message := "entity not found"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}

// Invalid creates a new Error instance with a default message indicating that the entity is invalid.
func Invalid(messages ...string) *Error {
	message := "invalid entity"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}
