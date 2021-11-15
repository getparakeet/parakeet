package src

import (
	"net/http"
	"bytes"

	"github.com/getparakeet/parakeet/errors"
)

func GetHttp(url string) (*http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		errors.UnknownError(err)
	}
	return resp
}

func PostHttp(url string, body []byte) (*http.Response) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		errors.UnknownError(err)
	}
	return resp
}