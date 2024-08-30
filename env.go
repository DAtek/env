package env

import (
	"errors"
	"os"
	"reflect"

	"github.com/iancoleman/strcase"
)

/*
Returns a new loader which can be used to parse environmental variables into a struct.

Example:
*/
func NewLoader[T any](parsers ...ParserMap) func(default_ ...any) (T, error) {
	allParsers := []ParserMap{}

	switch len(parsers) {
	case 0:
		allParsers = append(allParsers, baseParsers)
	default:
		allParsers = append(allParsers, parsers...)
		allParsers = append(allParsers, baseParsers)
	}

	return func(default_ ...any) (T, error) {
		return load[T](allParsers, default_...)
	}
}

func load[T any](parserMaps []ParserMap, default_ ...any) (T, error) {
	var obj T

	errorCollection := ErrorCollection{
		Errors: []FieldError{},
	}
	reflection := reflect.ValueOf(&obj).Elem()

	for _, field := range reflect.VisibleFields(reflect.TypeOf(obj)) {
		upperSnakeCaseField := strcase.ToScreamingSnake(field.Name)
		value, envVarSet := os.LookupEnv(upperSnakeCaseField)
		targetField := reflection.FieldByName(field.Name)
		targetType := field.Type.Name()

		if envVarSet {
			parse, err := getParserForType(parserMaps, targetType)
			if err != nil {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:      upperSnakeCaseField,
					ErrorType:     ERROR_UNSUPPORTED_TYPE,
					VariableType:  targetType,
					OriginalError: err,
				})
				continue
			}

			convertedVal, err := parse(value)

			if err != nil {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:      upperSnakeCaseField,
					ErrorType:     ERROR_WORNG_TYPE,
					VariableType:  targetType,
					OriginalError: err,
				})
				continue
			}

			targetField.Set(reflect.ValueOf(convertedVal))
			continue
		}

		if len(default_) > 0 {
			defaultConfig := default_[0]
			defaultConfigReflection := reflect.ValueOf(defaultConfig)
			defaultField := defaultConfigReflection.FieldByName(field.Name)

			if !defaultField.IsValid() {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:     upperSnakeCaseField,
					ErrorType:    ERROR_REQUIRED,
					VariableType: targetType,
				})
				continue
			}

			targetField.Set(defaultField)
			continue
		}

		if targetField.Kind().String() == "ptr" {
			continue
		}

		errorCollection.Errors = append(errorCollection.Errors, FieldError{
			Location:     upperSnakeCaseField,
			ErrorType:    ERROR_REQUIRED,
			VariableType: targetType,
		})
	}

	if len(errorCollection.Errors) == 0 {
		return obj, nil
	}

	return obj, &errorCollection
}

func getParserForType(parserMaps []ParserMap, type_ string) (Parser, error) {
	for _, parsers := range parserMaps {
		parser, ok := parsers[type_]
		if ok {
			return parser, nil
		}
	}

	return nil, errors.New("PARSER_NOT_FOUND")
}
