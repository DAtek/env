package env

import (
	"errors"
	"os"
	"reflect"

	"github.com/iancoleman/strcase"
)

type Config struct {
	Parsers       ParserMap
	PointerSetter PointerSetter
}

type PointerSetter func(field *reflect.Value, value any) error

// Returns a new loader which can be used to parse environmental variables into a struct.
func NewLoader[T any](config ...Config) func(default_ ...any) (T, error) {
	allParsers := []ParserMap{}
	var pointerSetter PointerSetter = nil

	switch len(config) {
	case 0:
		allParsers = append(allParsers, baseParsers)
	default:
		allParsers = append(allParsers, config[0].Parsers)
		allParsers = append(allParsers, baseParsers)
		pointerSetter = config[0].PointerSetter
	}

	return func(default_ ...any) (T, error) {
		return load[T](allParsers, pointerSetter, default_...)
	}
}

func load[T any](parserMaps []ParserMap, extraPointerSetter PointerSetter, default_ ...any) (T, error) {
	var obj T

	pointerSetters := []PointerSetter{setPointerValue}
	if extraPointerSetter != nil {
		pointerSetters = append(pointerSetters, extraPointerSetter)
	}

	errorCollection := ErrorCollection{
		Errors: []FieldError{},
	}
	reflection := reflect.ValueOf(&obj).Elem()

	for _, field := range reflect.VisibleFields(reflect.TypeOf(obj)) {
		upperSnakeCaseField := strcase.ToScreamingSnake(field.Name)
		value, envVarSet := os.LookupEnv(upperSnakeCaseField)
		targetField := reflection.FieldByName(field.Name)
		targetType := getFieldTypeName(&field)

		if envVarSet {
			parse, err := getParserForType(parserMaps, targetType)
			if err != nil {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:      upperSnakeCaseField,
					ErrorType:     ErrorParserMissingType,
					VariableType:  targetType,
					OriginalError: err,
				})
				continue
			}

			convertedVal, err := parse(value)

			if err != nil {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:      upperSnakeCaseField,
					ErrorType:     ErrorWrongType,
					VariableType:  targetType,
					OriginalError: err,
				})
				continue
			}

			for _, setPointerValue := range pointerSetters {
				err = setPointerValue(&targetField, convertedVal)
				if err == nil {
					break
				}
			}

			if err != nil {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:      upperSnakeCaseField,
					ErrorType:     ErrorPointerSetterMissing,
					VariableType:  targetType,
					OriginalError: err,
				})
			}

			continue
		}

		if len(default_) > 0 {
			defaultConfig := default_[0]
			defaultConfigReflection := reflect.ValueOf(defaultConfig)
			defaultField := defaultConfigReflection.FieldByName(field.Name)

			if !defaultField.IsValid() {
				errorCollection.Errors = append(errorCollection.Errors, FieldError{
					Location:     upperSnakeCaseField,
					ErrorType:    ErrorRequired,
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
			ErrorType:    ErrorRequired,
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

func getFieldTypeName(field *reflect.StructField) string {
	if field.Type.Kind() == reflect.Pointer {
		return field.Type.Elem().Name()
	}

	return field.Type.Name()
}

func setPointerValue(field *reflect.Value, value any) error {
	if field.Kind() != reflect.Pointer {
		field.Set(reflect.ValueOf(value))
		return nil
	}

	switch value := value.(type) {
	case int:
		field.Set(reflect.ValueOf(&value))
	case int8:
		field.Set(reflect.ValueOf(&value))
	case int16:
		field.Set(reflect.ValueOf(&value))
	case int32:
		field.Set(reflect.ValueOf(&value))
	case int64:
		field.Set(reflect.ValueOf(&value))
	case uint:
		field.Set(reflect.ValueOf(&value))
	case uint8:
		field.Set(reflect.ValueOf(&value))
	case uint16:
		field.Set(reflect.ValueOf(&value))
	case uint32:
		field.Set(reflect.ValueOf(&value))
	case uint64:
		field.Set(reflect.ValueOf(&value))
	case float32:
		field.Set(reflect.ValueOf(&value))
	case float64:
		field.Set(reflect.ValueOf(&value))
	case string:
		field.Set(reflect.ValueOf(&value))
	case bool:
		field.Set(reflect.ValueOf(&value))
	default:
		return ErrorPointerSetterMissing
	}

	return nil
}

var _ PointerSetter = setPointerValue
