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

// HTTPClient is an Option type that allows you provide your own HTTP client
func WithHttpClient(client HttpClient) Option {
	return func(r *RpcClient) {
		r.httpClient = client
	}
}

// MODIFIERS
type (
	ModifierPayload      func(*JsonRpcRequest)
	ModifierHttpRequest  func(*http.Request)
	ModifierHttpResponse func(*http.Response, error)
)

func WithPayloadModifier(modifier ModifierPayload) Option {
	return func(r *RpcClient) {
		r.modifiers.payload = append(r.modifiers.payload, modifier)
	}
}
func WithHttpRequestModifier(modifier ModifierHttpRequest) Option {
	return func(r *RpcClient) {
		r.modifiers.httpRequest = append(r.modifiers.httpRequest, modifier)
	}
}
func WithHttpResponseModifier(modifier ModifierHttpResponse) Option {
	return func(r *RpcClient) {
		r.modifiers.httpResponse = append(r.modifiers.httpResponse, modifier)
	}
}

// MODIFIERS V2
func WithHeader(header http.Header) Option {
	return WithHttpRequestModifier(func(req *http.Request) {
		req.Header = header.Clone()
	})
}
func WithOnErrorOmitURL() Option {
	return WithHttpResponseModifier(func(_ *http.Response, err error) {
		redactErrorURL(err)
	})
}

func setDefaultOptions(r *RpcClient) {
	r.httpClient = &http.Client{}
	r.endpoint = MainnetRPCEndpoint
}
