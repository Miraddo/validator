package validator

type Empty interface {
	IsEmptyString(str string) bool
	IsEmptyArray(arr []string) bool
	IsEmptyMap(m map[string]string) bool
	IsEmptyInterface(i interface{}) bool
	IsEmptyNumber(n int) bool
	IsEmptyFloat(f float64) bool
}

func IsEmptyString(str string) bool {
	return str == ""
}

func IsEmptyArray(arr []string) bool {
	return len(arr) == 0
}

func IsEmptyMap(m map[string]string) bool {
	return len(m) == 0
}

func IsEmptyInterface(i interface{}) bool {
	return i == nil
}

func IsEmptyNumber(n int) bool {
	return n == 0
}

func IsEmptyFloat(f float64) bool {
	return f == 0
}
