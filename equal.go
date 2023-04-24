package validator

import "reflect"

type Equal interface {
	IsEqual(a, b interface{}) bool
	IsNotEqual(a, b interface{}) bool
	IsEqualType(a, b interface{}) bool
}

func IsEqual(a, b interface{}) bool {
	return a == b
}

func IsNotEqual(a, b interface{}) bool {
	return a != b
}

func IsEqualType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
