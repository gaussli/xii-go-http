package http

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/goccy/go-yaml"
)

// XiiResponse represents an HTTP response with various configurable options.
type XiiResponse struct {
	Request    *XiiRequest
	StatusCode int
	Proto      string
	Headers    http.Header
	Body       []byte
}

// IsSuccess returns true if the response status code is in the 2xx range.
func (resp *XiiResponse) IsSuccess() bool {
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// IsRedirect returns true if the response status code is in the 3xx range.
func (resp *XiiResponse) IsRedirect() bool {
	return resp.StatusCode >= 300 && resp.StatusCode < 400
}

// IsClientError returns true if the response status code is in the 4xx range.
func (resp *XiiResponse) IsClientError() bool {
	return resp.StatusCode >= 400 && resp.StatusCode < 500
}

// IsServerError returns true if the response status code is in the 5xx range.
func (resp *XiiResponse) IsServerError() bool {
	return resp.StatusCode >= 500 && resp.StatusCode < 600
}

// IsError returns true if the response status code is in the 4xx or 5xx range.
func (resp *XiiResponse) IsError() bool {
	return resp.IsClientError() || resp.IsServerError()
}

// TextBody returns the response body as a string.
func (resp *XiiResponse) TextBody() string {
	return string(resp.Body)
}

// JSONBody unmarshals the response body into the provided struct.
func (resp *XiiResponse) JSONBody(v any) error {
	return json.Unmarshal(resp.Body, v)
}

// XMLBody unmarshals the response body into the provided struct.
func (resp *XiiResponse) XMLBody(v any) error {
	return xml.Unmarshal(resp.Body, v)
}

// YAMLBody unmarshals the response body into the provided struct.
func (resp *XiiResponse) YAMLBody(v any) error {
	return yaml.Unmarshal(resp.Body, v)
}
