package http

import "gopkg.in/go-playground/validator.v9"

func validateRequest(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
