package response

import "strings"

type Formatter interface {
	f(string) interface{}
}

type Err struct {
	msg string
}

func NewErr(msg string) *Err {
	return &Err{msg: msg}
}

// invalidMessageErrFormatter format invalid error message
// Sample err msg: `invalid: msg error`
// Formatted to: "msg error"
func (e *Err) invalidMessageErrFormatter() string {
	msg := strings.TrimPrefix(e.msg, "invalid:")
	msg = strings.TrimSpace(msg)
	return msg
}

// invalidFieldErrFormatter is formatter error string to fields errors
//
// Sample err msg: `invalidField: email:Email already used, password:Password is wrong`
// Formatted to:
// {
//   "errors": {
//     "email":
//       "Email already used"
//     },
//     "password":
//       "Password is wrong"
//    }
// }
func (e *Err) invalidFieldErrFormatter() map[string]string {
	trimmedErr := strings.TrimPrefix(e.msg, "invalidField:")
	errors := strings.Split(trimmedErr, ",")

	formattedErrors := make(map[string]string)
	for _, err := range errors {
		err = strings.TrimSpace(err)
		splitErr := strings.Split(err, ":")

		// Get key and error message
		key := strings.TrimSpace(splitErr[0])
		var errMessage string
		if len(splitErr) > 0 {
			errMessage = strings.TrimSpace(splitErr[1])
		}

		formattedErrors[key] = errMessage
	}
	return formattedErrors
}

// toMuchRetry is tooMuchRetry Error
func (e *Err) toMuchRetry() string {
	trimmedErr := strings.TrimPrefix(e.msg, "toMuchRetry:")
	msg := strings.TrimSpace(trimmedErr)
	return msg
}
