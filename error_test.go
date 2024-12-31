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
					ErrorType:    ErrorWrongType,
					VariableType: "int",
				},
				{
					Location:     "VAR_B",
					ErrorType:    ErrorRequired,
					VariableType: "float32",
				},
				{
					Location:     "VAR_C",
					ErrorType:    ErrorParserMissingType,
					VariableType: "float32",
				},
				{
					Location:     "VAR_D",
					ErrorType:    ErrorPointerSetterMissing,
					VariableType: "customBool",
				},
			},
		}

		// when
		errorMsg := errorCollection.Error()

		// then
		expectedMsg := `Environmental variable 'VAR_A' has wrong type. Required type: 'int'
Environmental variable 'VAR_B' is unset
Parser missing for environmental variable 'VAR_C'. Required type: 'float32'
Pointer setter missing for environmental variable 'VAR_D'. Required type: 'customBool'`
		assert.Equal(t, expectedMsg, errorMsg)
	})

	t.Run("ErrorType is error", func(t *testing.T) {
		assert.NotEqual(t, "", ErrorWrongType.Error())
	})
}
