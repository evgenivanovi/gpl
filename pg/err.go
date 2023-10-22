package pg

import (
	"bytes"
	"fmt"

	"github.com/evgenivanovi/gpl/std"
)

const (
	ErrorUnknownEntity = "unknown"

	// ErrorInternalCode
	// Exception thrown when we can't classify an SQL exception.
	ErrorInternalCode    = "internal"
	ErrorInternalMessage = "An internal error has occurred"

	ErrorIntegrityCode    = "integrity"
	ErrorIntegrityMessage = "Integrity constraint violation has occurred"

	// ErrorUniqueCode
	// Exception thrown when an attempt to insert or update data results in violation of a primary key or unique constraint.
	ErrorUniqueCode    = "unique"
	ErrorUniqueMessage = "Primary key or unique integrity violation has occurred"

	// ErrorEmptyCode
	// Exception thrown when a result was expected to have at least one row (or element) but zero rows (or elements) were actually returned.
	ErrorEmptyCode    = "empty"
	ErrorEmptyMessage = "Result set is empty"

	// ErrorLockCode
	// Exception thrown on a pessimistic or optimistic locking violation.
	// Pessimistic lock: Thrown if a corresponding database error is encountered.
	// Optimistic lock: Thrown either by O/R mapping tools or by custom DAO implementations.
	// Optimistic locking failure is typically not detected by the database itself.
	ErrorLockCode    = "lock"
	ErrorLockMessage = "Pessimistic or optimistic locking violation has occurred"
)

// PersistenceError defines a standard persistence error.
type PersistenceError struct {
	// Entity
	Entity string
	// Machine-readable error code.
	Code string
	// Human-readable message.
	Message string
}

func NewErrorWithCode(code string) *PersistenceError {
	return &PersistenceError{
		Code: code,
	}
}

func NewErrorWithEntityCode(entity string, code string) *PersistenceError {
	return &PersistenceError{
		Entity: entity,
		Code:   code,
	}
}

func NewErrorWithCodeMessage(code string, message string) *PersistenceError {
	return &PersistenceError{
		Code:    code,
		Message: message,
	}
}

func NewErrorWithEntityCodeMessage(entity string, code string, message string) *PersistenceError {
	return &PersistenceError{
		Entity:  entity,
		Code:    code,
		Message: message,
	}
}

func (e *PersistenceError) WithEntity(entity string) *PersistenceError {
	e.Entity = entity
	return e
}

func (e *PersistenceError) WithCode(code string) *PersistenceError {
	e.Code = code
	return e
}

func (e *PersistenceError) WithMessage(message string) *PersistenceError {
	e.Message = message
	return e
}

// Error returns the string representation of the error message.
func (e *PersistenceError) Error() string {
	var buf bytes.Buffer

	// Print the error code & message.
	if e.Code != std.Empty {
		_, _ = fmt.Fprintf(&buf, "<%s> <%s> ", e.Entity, e.Code)
	}
	buf.WriteString(e.Message)

	return buf.String()
}

func ErrorEntity(err error) string {
	if err == nil {
		return std.Empty
	} else if e, ok := err.(*PersistenceError); ok && e.Entity != std.Empty {
		return e.Entity
	}
	return ErrorUnknownEntity
}

func ErrorCode(err error) string {
	if err == nil {
		return std.Empty
	} else if e, ok := err.(*PersistenceError); ok && e.Code != std.Empty {
		return e.Code
	}
	return ErrorInternalCode
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise, returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return std.Empty
	} else if e, ok := err.(*PersistenceError); ok && e.Message != std.Empty {
		return e.Message
	}
	return ErrorInternalMessage
}

func WithEntity(err error, entity string) error {
	if err == nil {
		return nil
	} else if e, ok := err.(*PersistenceError); ok {
		return e.WithEntity(entity)
	} else {
		return err
	}
}

func WithCode(err error, code string) error {
	if err == nil {
		return nil
	} else if e, ok := err.(*PersistenceError); ok {
		return e.WithCode(code)
	} else {
		return err
	}
}

func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	} else if e, ok := err.(*PersistenceError); ok {
		return e.WithMessage(message)
	} else {
		return err
	}
}
