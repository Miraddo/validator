package http

import (
	"net/http"
)

type HttpStatus interface {
	CheckStatusOK(url string) (bool, error)
	CheckStatusNotFound(url string) (bool, error)
	CheckStatusUnauthorized(url string) (bool, error)
	CheckStatusForbidden(url string) (bool, error)
	CheckStatusBadRequest(url string) (bool, error)
	CheckStatusInternalServerError(url string) (bool, error)
	CheckStatusServiceUnavailable(url string) (bool, error)
	CheckStatusGatewayTimeout(url string) (bool, error)
	CheckStatusHTTPVersionNotSupported(url string) (bool, error)
}

func CheckStatusOK(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, err
	}

	return true, nil
}

func CheckStatusNotFound(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	}

	return false, nil
}

func CheckStatusUnauthorized(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return true, nil
	}

	return false, nil
}

func CheckStatusForbidden(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		return true, nil
	}

	return false, nil
}

func CheckStatusBadRequest(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return true, nil
	}

	return false, nil
}

func CheckStatusInternalServerError(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return true, nil
	}

	return false, nil
}

func CheckStatusServiceUnavailable(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusServiceUnavailable {
		return true, nil
	}

	return false, nil
}

func CheckStatusGatewayTimeout(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusGatewayTimeout {
		return true, nil
	}

	return false, nil
}

func CheckStatusBadGateway(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadGateway {
		return true, nil
	}

	return false, nil
}

func CheckStatusHTTPVersionNotSupported(url string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusHTTPVersionNotSupported {
		return true, nil
	}

	return false, nil
}
