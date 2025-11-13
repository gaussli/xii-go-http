package http

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/url"

	"github.com/gaussli/xii-go-http/internal/utils"
	"github.com/goccy/go-yaml"
)

// XiiRequest represents an HTTP request with various configurable options.
type XiiRequest struct {
	method      string
	endpoint    string
	headers     map[string][]string
	body        io.Reader
	queryParams url.Values
	context     context.Context
}

// NewXiiRequest creates a new XiiRequest with default values.
func NewXiiRequest() *XiiRequest {
	return &XiiRequest{
		headers:     make(map[string][]string),
		queryParams: make(url.Values),
		context:     context.Background(),
	}
}

// Method sets the request method.
func (req *XiiRequest) Method(method string) *XiiRequest {
	req.method = method
	return req
}

// Endpoint sets the request endpoint.
func (req *XiiRequest) Endpoint(endpoint string) *XiiRequest {
	// Add / if not present
	if !utils.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}
	req.endpoint = endpoint
	return req
}

// Header adds a header to the request.
func (req *XiiRequest) Header(key, value string) *XiiRequest {
	if values, exists := req.headers[key]; exists {
		// Add value to existing header
		req.headers[key] = append(values, value)
	} else {
		// Create new header
		req.headers[key] = []string{value}
	}
	return req
}

// Body sets the request body to the provided io.Reader.
func (req *XiiRequest) Body(body io.Reader) *XiiRequest {
	req.body = body
	return req
}

// FormBody sets the request body to the form-urlencoded representation of formData.
func (req *XiiRequest) FormBody(formData url.Values) *XiiRequest {
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.body = bytes.NewReader([]byte(formData.Encode()))
	return req
}

// MultipartFormBody sets the request body to the multipart form representation of formData.
func (req *XiiRequest) MultipartFormBody(formData url.Values) *XiiRequest {
	req.Header("Content-Type", "multipart/form-data")
	req.body = bytes.NewReader([]byte(formData.Encode()))
	return req
}

// TextBody sets the request body to the plain text representation of text.
func (req *XiiRequest) TextBody(text string) *XiiRequest {
	req.Header("Content-Type", "text/plain")
	req.body = bytes.NewReader([]byte(text))
	return req
}

// JSONBody sets the request body to the JSON representation of jsonData.
func (req *XiiRequest) JSONBody(jsonData any) *XiiRequest {
	req.Header("Content-Type", "application/json")
	// Serialize jsonData to JSON
	jsonBytes, _ := json.Marshal(jsonData)
	req.body = bytes.NewReader(jsonBytes)
	return req
}

// XMLBody sets the request body to the XML representation of xmlData.
func (req *XiiRequest) XMLBody(xmlData any) *XiiRequest {
	req.Header("Content-Type", "application/xml")
	// Serialize xmlData to XML
	xmlBytes, _ := xml.Marshal(xmlData)
	req.body = bytes.NewReader(xmlBytes)
	return req
}

// YAMLBody sets the request body to the YAML representation of yamlData.
func (req *XiiRequest) YAMLBody(yamlData any) *XiiRequest {
	req.Header("Content-Type", "application/yaml")
	// Serialize yamlData to YAML
	yamlBytes, _ := yaml.Marshal(yamlData)
	req.body = bytes.NewReader(yamlBytes)
	return req
}

// QueryParam adds a query parameter to the request.
func (req *XiiRequest) QueryParam(key, value string) *XiiRequest {
	req.queryParams.Add(key, value)
	return req
}

// Context sets the request context.
func (req *XiiRequest) Context(ctx context.Context) *XiiRequest {
	req.context = ctx
	return req
}

// Shortcut methods for GET requests
func (req *XiiRequest) GET(endpoint string) *XiiRequest {
	return req.Method("GET").Endpoint(endpoint)
}

// Shortcut methods for POST requests
func (req *XiiRequest) POST(endpoint string) *XiiRequest {
	return req.Method("POST").Endpoint(endpoint)
}

// Shortcut methods for PUT requests
func (req *XiiRequest) PUT(endpoint string) *XiiRequest {
	return req.Method("PUT").Endpoint(endpoint)
}

// Shortcut methods for DELETE requests
func (req *XiiRequest) DELETE(endpoint string) *XiiRequest {
	return req.Method("DELETE").Endpoint(endpoint)
}

// Shortcut methods for PATCH requests
func (req *XiiRequest) PATCH(endpoint string) *XiiRequest {
	return req.Method("PATCH").Endpoint(endpoint)
}

// GetHeaders returns the headers of the request.
func (req *XiiRequest) GetHeaders() map[string][]string {
	return req.headers
}
