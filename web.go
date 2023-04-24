package validator

import (
	"regexp"

	pkgHttp "github.com/miraddo/validator/http"
)

type Web interface {
	IsEmail(str string) bool
	IsURL(str string) bool
	IsURLStatusOK(url string) (bool, error)
}

var (
	email = regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	url   = regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`)
)

func IsEmail(str string) bool {
	return email.MatchString(str)
}

func IsURL(str string) bool {
	return url.MatchString(str)
}

func IsURLStatusOK(url string) (bool, error) {
	return pkgHttp.CheckStatusOK(url)
}
