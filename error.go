package env

import (
	"fmt"
	"strings"
)

type ErrorType string

func (e ErrorType) Error() string {
	return string(e)
}

const (
	ErrorRequired             ErrorType = "required"
	ErrorWrongType            ErrorType = "wrong_type"
	ErrorParserMissingType    ErrorType = "parser_missing"
	ErrorPointerSetterMissing ErrorType = "pointer_setter_missing"
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
	ErrorRequired: func(fe *FieldError) string { return fmt.Sprintf("Environmental variable '%s' is unset", fe.Location) },
	ErrorWrongType: func(fe *FieldError) string {
		return fmt.Sprintf("Environmental variable '%s' has wrong type. Required type: '%s'", fe.Location, fe.VariableType)
	},
	ErrorParserMissingType: func(fe *FieldError) string {
		return fmt.Sprintf("Parser missing for environmental variable '%s'. Required type: '%s'", fe.Location, fe.VariableType)
	},
	ErrorPointerSetterMissing: func(fe *FieldError) string {
		return fmt.Sprintf("Pointer setter missing for environmental variable '%s'. Required type: '%s'", fe.Location, fe.VariableType)
	},
}
