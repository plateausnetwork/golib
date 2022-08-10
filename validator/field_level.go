package validator

import reflect "reflect"

type FieldLevel interface {
	// Field returns current field for validation
	Field() reflect.Value
}
