package response

import (
	"fmt"
	"github.com/afiefafian/todo-api/src/delivery/http/helper"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

// ValidateAndReturnErrResultHTTP validate data and return result to http
func ValidateAndReturnErrResultHTTP(w http.ResponseWriter, m interface{}) bool {
	if ok, err := helper.ValidateStruct(m); !ok {
		var msg string
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			if len(castedObject) > 0 {
				msg = fmt.Sprintf("%s", castedObject[0])
			}
		}

		JSONErrorValidation(w, msg)
		return false
	}
	return true
}

// ValidateAndReturnErrResultHTTPWithErrField validate data and return result via http in error field
func ValidateAndReturnErrResultHTTPWithErrField(w http.ResponseWriter, m interface{}) bool {
	if ok, err := helper.ValidateStruct(m); !ok {
		var errors = make(map[string]string)

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				key := strings.ToLower(err.Field())
				errors[key] = fmt.Sprintf("%s", err)
			}
		}

		JSONErrorValidationWithField(w, errors)
		return false
	}
	return true
}
