package response

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

// JSONResult return result in json format
// Accepted resultByte type data is slice byte
func JSONResult(w http.ResponseWriter, resultPtr interface{}) {
	resultBytePtr, code := parseToByte(resultPtr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(*resultBytePtr)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

// parseToByte parsing result data to slice byte json
func parseToByte(resultPtr interface{}) (*[]byte, int) {
	result, err := json.Marshal(resultPtr)
	if err != nil {
		response := ResultFormat{
			Status:  http.StatusBadRequest,
			Name:    http.StatusText(http.StatusBadRequest),
			Message: "Failed parse data to json",
		}
		errRes, _ := json.Marshal(response)
		return &errRes, http.StatusBadRequest
	}

	code := http.StatusOK
	if resultPtr != nil {
		code = getStatusCodeFromResult(resultPtr)
	}

	return &result, code
}

// getStatusCodeFromResult get Status code from result data
func getStatusCodeFromResult(resultPtr interface{}) int {
	resultReflect := reflect.ValueOf(resultPtr)
	// Get reflect data from pointer
	if resultReflect.Kind() == reflect.Ptr {
		resultReflect = resultReflect.Elem()
	}

	if resultReflect.Kind() == reflect.Struct {
		// Get info by field name
		if _, valid := resultReflect.Type().FieldByName("Status"); valid {
			Status := resultReflect.FieldByName("Status").Int()
			return int(Status)
		}
	}

	return http.StatusOK
}
