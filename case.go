package validator

import (
	"regexp"
)

type Case interface {
	IsLowercase(str string) bool
	IsUppercase(str string) bool
	IsNumber(str string) bool
	IsLetter(str string) bool
	IsLetterNumber(str string) bool
	IsLetterNumberLine(str string) bool
	HasInjectCharacters(str string) bool
	IsHTMLTag(str string) bool
}

var (
	lowercase        = regexp.MustCompile(`^[a-z]+$`)
	uppercase        = regexp.MustCompile(`^[A-Z]+$`)
	number           = regexp.MustCompile(`^[0-9]+$`)
	letter           = regexp.MustCompile(`^[a-zA-Z]+$`)
	letterNumber     = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	letterNumberLine = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	injectCharacters = regexp.MustCompile("^[^<>;`&]*$")
	htmlTag          = regexp.MustCompile(`^<.*>$`)
)

func IsLowercase(str string) bool {
	return lowercase.MatchString(str)
}

func IsUppercase(str string) bool {
	return uppercase.MatchString(str)
}

func IsNumber(str string) bool {
	return number.MatchString(str)
}

func IsLetter(str string) bool {
	return letter.MatchString(str)
}

func IsLetterNumber(str string) bool {
	return letterNumber.MatchString(str)
}

func IsLetterNumberLine(str string) bool {
	return letterNumberLine.MatchString(str)
}

func HasInjectCharacters(str string) bool {
	return injectCharacters.MatchString(str)
}

func IsHtmlTag(str string) bool {
	return htmlTag.MatchString(str)
}
