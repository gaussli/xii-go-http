package http

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// XiiMiddleware is a function that modifies the HTTP request before it is sent.
type XiiMiddleware func(*http.Request) error

// XiiClient is a client for making HTTP requests.
type XiiClient struct {
	client      *http.Client
	baseURL     string
	baseHeaders map[string][]string
	middleware  []XiiMiddleware
}

// Option is a function that modifies the XiiClient.
type Option func(*XiiClient)

// NewClient creates a new XiiClient with the provided options.
func NewClient(opts ...Option) *XiiClient {
	client := &XiiClient{
		client:      &http.Client{Timeout: 30 * time.Second},
		baseHeaders: make(map[string][]string),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// WithBaseURL sets the base URL for the XiiClient.
func WithBaseURL(baseURL string) Option {
	return func(hc *XiiClient) {
		hc.baseURL = baseURL
	}
}

// WithTimeout sets the timeout for the XiiClient.
func WithTimeout(timeout time.Duration) Option {
	return func(hc *XiiClient) {
		hc.client.Timeout = timeout
	}
}

// WithProxy sets the proxy for the XiiClient.
func WithProxy(proxyStr string) Option {
	return func(hc *XiiClient) {
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			panic(err)
		}
		if transport := hc.client.Transport; transport == nil {
			transport = &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
		} else if tr, ok := transport.(*http.Transport); ok {
			tr.Proxy = http.ProxyURL(proxyURL)
		}
	}
}

// WithHeader sets a header for the XiiClient.
func WithHeader(key, value string) Option {
	return func(hc *XiiClient) {
		if values, exists := hc.baseHeaders[key]; exists {
			hc.baseHeaders[key] = append(values, value)
		} else {
			hc.baseHeaders[key] = []string{value}
		}
	}
}

// Use adds a middleware to the XiiClient.
func (c *XiiClient) Use(mw XiiMiddleware) {
	c.middleware = append(c.middleware, mw)
}

// Do sends an HTTP request and returns an HTTP response.
func (c *XiiClient) Do(req *XiiRequest) (*XiiResponse, error) {
	// Build full URL
	fullURL := c.baseURL + req.endpoint

	log.Printf("Request URL: %s", fullURL)
	log.Printf("Request Body: %s", req.body)

	// Create HTTP request
	httpReq, err := http.NewRequest(req.method, fullURL, req.body)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	if len(req.queryParams) > 0 {
		httpReq.URL.RawQuery = req.queryParams.Encode()
	}

	// Add base headers
	for key, values := range c.baseHeaders {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	// Add request headers
	for key, values := range req.headers {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	for key, value := range httpReq.Header {
		log.Printf("Request Header: %s: %s", key, value)
	}

	// Apply middleware
	for _, mw := range c.middleware {
		if err := mw(httpReq); err != nil {
			return nil, err
		}
	}

	// Execute the request
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	// Read response body
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Response Body: %s", string(body))

	return &XiiResponse{
		Request:    req,
		StatusCode: httpResp.StatusCode,
		Proto:      httpResp.Proto,
		Headers:    httpResp.Header,
		Body:       body,
	}, nil
}
