package env

import (
	"fmt"
	"net/url"
	"os"
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
		assert.Equal(t, ERROR_REQUIRED, errorCollection.Errors[0].ErrorType)
		assert.Equal(t, ERROR_REQUIRED, errorCollection.Errors[1].ErrorType)
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
		assert.Equal(t, ERROR_WORNG_TYPE, errorCollection.Errors[0].ErrorType)
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
		assert.Equal(t, ERROR_REQUIRED, errorCollection.Errors[0].ErrorType)
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

		load := NewLoader[config](customParsers)
		os.Setenv(envSomeVariable, "2")

		// when
		conf := gotils.ResultOrPanic(load())

		// then
		assert.Equal(t, expectedResult, conf.SomeVariable)
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
		assert.Equal(t, ERROR_UNSUPPORTED_TYPE, errorCollection.Errors[0].ErrorType)
		assert.Equal(t, "URL", errorCollection.Errors[0].VariableType)
	})

}

const envSomeVariable = "SOME_VARIABLE"
