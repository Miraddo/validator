package validator

import (
	"path/filepath"
	"regexp"
	"strings"
	"time"

	pkgHttp "github.com/miraddo/validator/http"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Image interface {
	IsImage(url string) (bool, error)
	IsImageUrlLocaly(url string) (string, bool, error)
	ISImageUrlStatus(url string) (bool, error)
	IsImageUrl(url string) bool
}

var (
	regImg = regexp.MustCompile(`^https?:\/\/.*\.(?:png|jpg|jpeg|gif|bmp|webp)$`)
)

func IsImageURLStatus(url string) (bool, error) {
	return pkgHttp.CheckStatusOK(url)
}

func IsImageUrl(url string) bool {
	return regImg.MatchString(url)
}

func IsImage(url string) (bool, error) {
	if !IsImageUrl(url) {
		return false, nil
	}

	return IsImageURLStatus(url)
}

func IsImageLocaly(url string) (bool, error) {
	name := "temp" + time.Now().String() + ".jpg"

	err := pkgHttp.DownloadFile(name, url)
	if err != nil {
		return false, err
	}

	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true, nil
	default:
		return false, nil
	}
}
