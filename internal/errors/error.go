package errors

import "strings"

type Error struct {
	Message string `json:"message,omitempty"`
}

func New(messages ...string) *Error {
	message := strings.Join(messages, ": ")
	return &Error{
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func NotSaved(messages ...string) *Error {
	message := "failed to save the entity"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}

func NotDeleted(messages ...string) *Error {
	message := "failed to delete the entity"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}

func NotFound(messages ...string) *Error {
	message := "entity not found"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}

func Invalid(messages ...string) *Error {
	message := "invalid entity"
	if len(messages) > 0 {
		messages = append([]string{message}, messages...)
	}
	return New(messages...)
}
