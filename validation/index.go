package validation

import "github.com/go-playground/validator"

func GetError(err validator.ValidationErrors, message map[string]string) string {
	for _, item := range err {
		// struct column + form validation tag
		if val, ok := message[item.Field() + "." + item.Tag()]; ok {
			return val
		}
	}
	return "參數錯誤"
}