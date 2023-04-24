package validator

import (
	"regexp"
)

type Phone interface {
	IsDENumber(str string) bool
	IsUSNumber(str string) bool
	IsEuropeNumber(str string) bool
}

var (
	phoneNumberTR     = regexp.MustCompile(`^(\+?0?90\-?)?5\d{9}$`)
	phoneNumberDE     = regexp.MustCompile(`^(\+?0?49\-?)?1[345789]\d{9}$`)
	phoneNumberUS     = regexp.MustCompile(`^(\+?0?1\-?)?1[345789]\d{9}$`)
	phoneNumberEurope = regexp.MustCompile(`^(\+?0?44\-?)?1[345789]\d{9}$`)
	phoneNumberIR     = regexp.MustCompile(`^(\+?0?98\-?)?9\d{9}$`)
)

func IsDENumber(str string) bool {
	return phoneNumberDE.MatchString(str)
}

func IsUSNumber(str string) bool {
	return phoneNumberUS.MatchString(str)
}

func IsEuropeNumber(str string) bool {
	return phoneNumberEurope.MatchString(str)
}

func IsIRNumber(str string) bool {
	return phoneNumberIR.MatchString(str)
}

func IsTRNumber(str string) bool {
	return phoneNumberTR.MatchString(str)
}
