package helpers

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
)

type Errors struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func ErrorValidation(s interface{}, err error) *Errors {
	typ := reflect.TypeOf(s).Elem()

	errs := map[string]string{}
	errsByField := govalidator.ErrorsByField(err)

	for i := 0; i < typ.NumField(); i++ {
		var message string
		key := typ.Field(i).Name
		tag := typ.Field(i).Tag

		if tag.Get("json") == "" {
			message = errsByField[key]
		} else {
			message = errsByField[tag.Get("json")]
		}

		if message != "" {
			name := strings.Split(tag.Get("json"), " ")[0]

			if name == "-" {
				continue
			}

			if name != "" {
				errs[name] = message
				continue
			}

			sKey := ToSnakeCase(key)
			errs[sKey] = message
		}
	}

	mapper := &Errors{Message: "Validation Error", Errors: errs}
	return mapper
}

func ToSnakeCase(in string) string {
	runes := []rune(in)

	var out []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && (unicode.IsUpper(runes[i]) || unicode.IsNumber(runes[i])) && ((i+1 < len(runes) && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
