package response

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strings"
)

// NotFoundHandler is a middleware for handle 404 Error
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	response := ResultNotFound
	JSONResult(w, &response)
}

// PanicHandler is a middleware for catching panic and prevent app to crash
func PanicHandler() func(http.ResponseWriter, *http.Request, interface{}) {
	return func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		response := ResultInternalServerErr
		if env := viper.GetString(`env`); env == "development" {
			if rcv != nil {
				response.SetMessage(rcv)
			}
		}

		log.Printf("%s %s", r.Method, r.URL.Path)
		log.Printf("Panic Error: %+v", rcv)

		JSONResult(w, &response)
	}
}

// JSONError is a response handler for error result
func JSONError(w http.ResponseWriter, error string, code int) {
	switch e := NewErr(error); {
	case strings.HasPrefix(e.msg, "invalid:"):
		msg := e.invalidMessageErrFormatter()
		JSONErrorValidation(w, msg)
		return
	case strings.HasPrefix(error, "invalidField:"):
		formattedErrors := e.invalidFieldErrFormatter()
		JSONErrorValidationWithField(w, formattedErrors)
		return
	case strings.HasPrefix(error, "toMuchRetry:"):
		formattedErrors := e.toMuchRetry()
		JSONErrorValidationWithField(w, formattedErrors)
		return
	}

	response := ResultErr("")
	response.SetStatus(code)
	response.SetMessage(error)
	JSONResult(w, &response)
}

// JSONErrorValidation return result as an error validation.
// Error is in message field, only 1 error can contain in message.
func JSONErrorValidation(w http.ResponseWriter, msg string) {
	response := ResultValidationErr
	response.SetMessage(msg)
	JSONResult(w, &response)
}

// JSONErrorValidationWithField return result as an error validation with error per field.
// Error will be listed in the error field per input field.
func JSONErrorValidationWithField(w http.ResponseWriter, errors interface{}) {
	response := ResultValidationErr
	response.SetErrors(errors)
	JSONResult(w, &response)
}
