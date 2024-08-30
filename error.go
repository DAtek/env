package env

import (
	"fmt"
	"strings"
)

type ErrorType string

const (
	ERROR_REQUIRED         = ErrorType("required")
	ERROR_WORNG_TYPE       = ErrorType("wrong_type")
	ERROR_UNSUPPORTED_TYPE = ErrorType("unsupported_type")
)

type ErrorCollection struct {
	Errors []FieldError
}

func (collection *ErrorCollection) Error() string {
	msgParts := []string{}
	for _, err := range collection.Errors {
		msgParts = append(msgParts, err.Error())
	}
	return strings.Join(msgParts, "\n")
}

type FieldError struct {
	Location      string
	ErrorType     ErrorType
	VariableType  string
	OriginalError error
}

func (fieldError *FieldError) Error() string {
	createErrorMsg := errorMsgGenerators[fieldError.ErrorType]
	return createErrorMsg(fieldError)
}

var errorMsgGenerators = map[ErrorType]func(*FieldError) string{
	ERROR_REQUIRED: func(fe *FieldError) string { return fmt.Sprintf("Environmental variable '%s' is unset", fe.Location) },
	ERROR_WORNG_TYPE: func(fe *FieldError) string {
		return fmt.Sprintf("Environmental variable '%s' has wrong type. Required type: '%s'", fe.Location, fe.VariableType)
	},
	ERROR_UNSUPPORTED_TYPE: func(fe *FieldError) string {
		return fmt.Sprintf("Parser missing for environmental variable '%s'. Required type: '%s'", fe.Location, fe.VariableType)
	},
}
