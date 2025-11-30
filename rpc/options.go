package rpc

import (
	"net/http"
)

// Option is a configuration type for the Client
type Option func(*RpcClient)

// WithEndpoint is an Option that allows you configure the rpc endpoint that our
// client will point to
func WithEndpoint(endpoint string) Option {
	return func(r *RpcClient) {
		r.endpoint = endpoint
	}
}

func WithHeader(header http.Header) Option {
	return func(r *RpcClient) {
		r.header = header
	}
}

// HTTPClient is an Option type that allows you provide your own HTTP client
func WithHttpClient(client HttpClient) Option {
	return func(r *RpcClient) {
		r.httpClient = client
	}
}

func WithOnErrorOmitURL(omit bool) Option {
	return func(r *RpcClient) {
		r.onErrorOmitURL = omit
	}
}

func setDefaultOptions(r *RpcClient) {
	r.httpClient = &http.Client{}
	r.endpoint = MainnetRPCEndpoint
}
