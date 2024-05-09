package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetErrorMessage(e validator.FieldError, obj any) (string, string) {
	errmsg := e.StructField() + ": " + e.Tag()
	fieldName := e.StructField()

	t := reflect.TypeOf(obj)
	field, _ := t.FieldByName(e.StructField())
	if msg, ok := field.Tag.Lookup(fmt.Sprintf("err_%s_msg", e.Tag())); ok {
		errmsg = msg
	}

	if msg, ok := field.Tag.Lookup("json"); ok {
		fieldName = msg
	}

	return fieldName, errmsg
}
