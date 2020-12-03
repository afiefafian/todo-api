package response

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// NotFoundHandler set response header type data to json
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	response := ResultNotFound
	JSONResult(w, &response)
}

func PanicHandler() func(http.ResponseWriter, *http.Request, interface{}) {
	return func(w http.ResponseWriter, r *http.Request, rcv interface{}) {
		response := ResultInternalServerErr
		if env := viper.GetString(`env`); env == "development" {
			if rcv != nil {
				response.Message = rcv
			}
		}

		log.Printf("%s %s", r.Method, r.URL.Path)
		log.Printf("Panic Error : %s", rcv)

		JSONResult(w, &response)
	}
}

func JSONError(w http.ResponseWriter, error string, code int) {
	response := ResultErr
	response.Status = code
	response.Message = error
	JSONResult(w, &response)
}

// JSONErrorValidation return result as an error validation
// Error is in message field, only 1 error can contain in message
func JSONErrorValidation(w http.ResponseWriter, msg string) {
	response := ResultValidationErr
	response.Message = msg
	JSONResult(w, &response)
}

// JSONErrorValidationWithField return result as an error validation with error per field
// Error will be listed in the error field per input field
// Sample
//
//
func JSONErrorValidationWithField(w http.ResponseWriter, errors interface{}) {
	response := ResultValidationErr
	response.Error = errors
	JSONResult(w, &response)
}
