package err

import (
	"bytes"
	"fmt"

	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

const (
	ErrorInternalCode    = "internal"
	ErrorInternalMessage = "An internal error has occurred."
)

// Error defines a standard application error.
type Error struct {
	// Entity
	Entity string
	// Machine-readable error code.
	Code string
	// Human-readable message.
	Message string
	// Logical operation.
	Operation string
	// Nested error
	Err error
}

func NewErrorWithEntityCode(entity string, code string) *Error {
	return &Error{
		Entity: entity,
		Code:   code,
	}
}

func NewErrorWithEntityCodeMessage(entity string, code string, message string) *Error {
	return &Error{
		Entity:  entity,
		Code:    code,
		Message: message,
	}
}

func NewErrorWithOpAndError(op string, err error) *Error {
	return &Error{
		Operation: op,
		Err:       err,
	}
}

/* __________________________________________________ */

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Operation != std.Empty {
		_, _ = fmt.Fprintf(&buf, "%s: ", e.Operation)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise, print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != std.Empty {
			_, _ = fmt.Fprintf(&buf, "<%s> <%s> ", e.Entity, e.Code)
		}
		buf.WriteString(e.Message)
	}

	return buf.String()
}

/* __________________________________________________ */

func ErrorCode(err error) string {
	if err == nil {
		return std.Empty
	} else if e, ok := err.(*Error); ok && e.Code != std.Empty {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return ErrorInternalCode
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise, returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return std.Empty
	} else if e, ok := err.(*Error); ok && e.Message != std.Empty {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return ErrorInternalMessage
}

/* __________________________________________________ */
