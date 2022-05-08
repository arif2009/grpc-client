package handlers

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

const grpcTimeout = time.Second * 15

type validationField struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

type validationFields []*validationField

func validate(obj interface{}) validationFields {
	errors := validationFields{}
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	if err := validate.Struct(obj); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			namespaceSplit := strings.Split(err.Namespace(), ".")
			errors = append(errors, &validationField{
				Field: strings.Join(namespaceSplit[1:], "."),
				Rule:  err.Tag(),
			})
		}
	}

	return errors
}

type Include []string

func (inc Include) Has(toCheck string) bool {
	for _, item := range inc {
		if item == toCheck {
			return true
		}
	}

	return false
}
