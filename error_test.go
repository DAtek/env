package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorCollection(t *testing.T) {

	t.Run("Returns formatted error", func(t *testing.T) {
		// given
		errorCollection := ErrorCollection{
			Errors: []FieldError{
				{
					Location:     "VAR_A",
					ErrorType:    ERROR_WORNG_TYPE,
					VariableType: "int",
				},
				{
					Location:     "VAR_B",
					ErrorType:    ERROR_REQUIRED,
					VariableType: "float32",
				},
				{
					Location:     "VAR_C",
					ErrorType:    ERROR_UNSUPPORTED_TYPE,
					VariableType: "float32",
				},
			},
		}

		// when
		errorMsg := errorCollection.Error()

		// then
		expectedMsg := `Environmental variable 'VAR_A' has wrong type. Required type: 'int'
Environmental variable 'VAR_B' is unset
Parser missing for environmental variable 'VAR_C'. Required type: 'float32'`
		assert.Equal(t, expectedMsg, errorMsg)
	})
}
