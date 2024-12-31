package env

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/DAtek/gotils"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	testWithClearEnv := func(name string, testFunc func(*testing.T)) {
		os.Clearenv()
		t.Run(name, testFunc)
	}

	testWithClearEnv("Loads string", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable string
		}

		load := NewLoader[config]()

		v := "red"
		os.Setenv(envSomeVariable, v)

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, v, conf.SomeVariable)
	})

	testWithClearEnv("Loads int", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, 23, conf.SomeVariable)
	})

	testWithClearEnv("Loads int8", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int8
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "-100")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, int8(-100), conf.SomeVariable)
	})

	testWithClearEnv("Loads int16", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int16
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "-100")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, int16(-100), conf.SomeVariable)
	})

	testWithClearEnv("Loads int32", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int32
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, int32(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads int64", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int64
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, int64(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads uint", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable uint
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, uint(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads uint8", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable uint8
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, uint8(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads uint16", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable uint16
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, uint16(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads uint32", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable uint32
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, uint32(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads uint64", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable uint64
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, uint64(23), conf.SomeVariable)
	})

	testWithClearEnv("Loads float32", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable float32
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23.36")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, float32(23.36), conf.SomeVariable)
	})

	testWithClearEnv("Loads float64", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable float64
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "23.36")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, float64(23.36), conf.SomeVariable)
	})

	boolConversionScenarios := []struct {
		input    string
		expected bool
	}{
		{"Y", true},
		{"Y", true},
		{"yes", true},
		{"TrUe", true},
		{"false", false},
		{"asdasd", false},
		{"", false},
	}

	for _, scenario := range boolConversionScenarios {
		testWithClearEnv(fmt.Sprintf("Loads %s", scenario.input), func(t *testing.T) {
			// given
			type config struct {
				SomeVariable bool
			}
			load := NewLoader[config]()

			os.Setenv(envSomeVariable, scenario.input)

			// when
			conf := gotils.ResultOrPanic(load())

			// then
			assert.Equal(t, scenario.expected, conf.SomeVariable)
		})
	}

	testWithClearEnv("Returns error if uint conversion fails", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable uint
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "-23")

		// when
		_, err := load()

		// then
		assert.Error(t, err)
	})

	testWithClearEnv("Returns error if int conversion fails", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "asd")

		// when
		_, err := load()

		// then
		assert.Error(t, err)
	})

	testWithClearEnv("Returns error if float conversion fails", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable float32
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "asd")

		// when
		_, err := load()

		// then
		assert.Error(t, err)
	})

	testWithClearEnv("Error if fields are required", func(t *testing.T) {
		// given
		type config struct {
			FieldA float32
			FieldB int
		}
		load := NewLoader[config]()

		// when
		_, err := load()
		errorCollection := err.(*ErrorCollection)

		// then
		assert.Equal(t, 2, len(errorCollection.Errors))
		assert.Equal(t, "FIELD_A", errorCollection.Errors[0].Location)
		assert.Equal(t, "FIELD_B", errorCollection.Errors[1].Location)
		assert.Equal(t, ErrorRequired, errorCollection.Errors[0].ErrorType)
		assert.Equal(t, ErrorRequired, errorCollection.Errors[1].ErrorType)
	})

	testWithClearEnv("Error if variable has wrong type", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int
		}
		load := NewLoader[config]()

		os.Setenv(envSomeVariable, "asd")

		// when
		_, err := load()
		errorCollection := err.(*ErrorCollection)

		// then
		assert.Equal(t, envSomeVariable, errorCollection.Errors[0].Location)
		assert.Equal(t, ErrorWrongType, errorCollection.Errors[0].ErrorType)
	})

	testWithClearEnv("Loads default", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int
		}
		load := NewLoader[config]()

		defaultConfig := config{5}

		// when
		loadedConfig := gotils.ResultOrPanic(load(defaultConfig))

		// then
		assert.Equal(t, defaultConfig.SomeVariable, loadedConfig.SomeVariable)
	})

	testWithClearEnv("Loads default from pointer", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable *int
		}
		load := NewLoader[config]()

		defaultConfig := config{gotils.Pointer(5)}

		// when
		loadedConfig := gotils.ResultOrPanic(load(defaultConfig))

		// then
		assert.Equal(t, defaultConfig.SomeVariable, loadedConfig.SomeVariable)
	})

	testWithClearEnv("No error if parameter is optional", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable *int
		}
		load := NewLoader[config]()

		// when
		loadedConfig := gotils.ResultOrPanic(load())

		// then
		assert.Nil(t, loadedConfig.SomeVariable)
	})

	testWithClearEnv("Returns error if default config not contains field", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int
		}
		load := NewLoader[config]()

		defaultConfig := struct {
			OtherVariable int
		}{5}

		// when
		_, err := load(defaultConfig)
		errorCollection := err.(*ErrorCollection)

		// then
		assert.Error(t, err)
		assert.Equal(t, envSomeVariable, errorCollection.Errors[0].Location)
		assert.Equal(t, ErrorRequired, errorCollection.Errors[0].ErrorType)
	})

	testWithClearEnv("Custom parsers have precedence over base parsers", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable int
		}

		expectedResult := 6
		customParsers := ParserMap{
			"int": func(src string) (any, error) {
				return expectedResult, nil
			},
		}

		load := NewLoader[config](Config{Parsers: customParsers})
		os.Setenv(envSomeVariable, "2")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, expectedResult, conf.SomeVariable)
	})

	type config[T any] struct {
		SomeVariable T
	}

	optionalNumberScenarios := []struct {
		LoadConfig any
		Expected   any
	}{
		{NewLoader[config[*int]](), 2},
		{NewLoader[config[*int8]](), int8(2)},
		{NewLoader[config[*int16]](), int16(2)},
		{NewLoader[config[*int32]](), int32(2)},
		{NewLoader[config[*int64]](), int64(2)},
		{NewLoader[config[*uint]](), uint(2)},
		{NewLoader[config[*uint8]](), uint8(2)},
		{NewLoader[config[*uint16]](), uint16(2)},
		{NewLoader[config[*uint32]](), uint32(2)},
		{NewLoader[config[*uint64]](), uint64(2)},
		{NewLoader[config[*float32]](), float32(2)},
		{NewLoader[config[*float64]](), float64(2)},
		{NewLoader[config[*string]](), "2"},
	}
	for _, scenario := range optionalNumberScenarios {
		t.Run("Parses optional values", func(t *testing.T) {
			os.Setenv(envSomeVariable, "2")
			load := reflect.ValueOf(scenario.LoadConfig)

			// when
			result := load.Call(nil)[0]

			// then
			fmt.Printf("result: %v\n", result)
			field := result.FieldByName("SomeVariable")
			assert.True(t, reflect.ValueOf(scenario.Expected).Equal(field.Elem()))
		})
	}

	t.Run("Parses optional bool", func(t *testing.T) {
		os.Setenv(envSomeVariable, "true")
		load := NewLoader[config[*bool]]()

		// when
		result := gotils.ResultOrPanic(load())

		// then
		fmt.Printf("result: %v\n", result)
		assert.Equal(t, true, *result.SomeVariable)
	})

	testWithClearEnv("Parses custom child pointer struct", func(t *testing.T) {
		// given
		type ChildStruct struct {
			Age int
		}

		type config struct {
			SomeVariable *ChildStruct
		}

		expectedResult := ChildStruct{Age: 2}

		customParsers := ParserMap{
			"ChildStruct": func(src string) (any, error) {
				return expectedResult, nil
			},
		}

		load := NewLoader[config](Config{
			Parsers: customParsers,
			PointerSetter: func(field *reflect.Value, value any) error {
				child, ok := value.(ChildStruct)
				if ok {
					field.Set(reflect.ValueOf(&child))
					return nil
				}

				return ErrorPointerSetterMissing
			},
		})
		os.Setenv(envSomeVariable, "2")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, expectedResult, *conf.SomeVariable)
	})

	testWithClearEnv("Returns error if parser not implemented for type", func(t *testing.T) {
		// given
		type config struct {
			SomeVariable url.URL
		}
		load := NewLoader[config]()
		os.Setenv(envSomeVariable, "2")

		// when
		_, err := load()
		errorCollection := err.(*ErrorCollection)

		// then
		assert.Error(t, err)
		assert.Equal(t, envSomeVariable, errorCollection.Errors[0].Location)
		assert.Equal(t, ErrorParserMissingType, errorCollection.Errors[0].ErrorType)
		assert.Equal(t, "URL", errorCollection.Errors[0].VariableType)
	})

	testWithClearEnv("Returns error if pointer setter not implemented for custom type", func(t *testing.T) {
		// given
		type customInt int

		type config struct {
			SomeVariable *customInt
		}

		load := NewLoader[config](Config{
			Parsers: ParserMap{
				"customInt": func(src string) (any, error) {
					return customInt(2), nil
				},
			},
		})

		os.Setenv(envSomeVariable, "2")

		// when
		_, err := load()
		errorCollection := err.(*ErrorCollection)

		// then
		assert.Error(t, err)
		assert.Equal(t, envSomeVariable, errorCollection.Errors[0].Location)
		assert.Equal(t, ErrorPointerSetterMissing, errorCollection.Errors[0].ErrorType)
	})
}

const envSomeVariable = "SOME_VARIABLE"
