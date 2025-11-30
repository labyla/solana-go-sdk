package rpc

import (
	"errors"
	"net/http"
	"net/url"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func redactErrorURL(err error) {
	var ue *url.Error
	if errors.As(err, &ue) {
		ue.URL = "<omitted>"
	}
}
