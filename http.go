package ntnn

import (
	"context"
	"io"
	"net/http"
)

func do(ctx context.Context, method, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

// Get creates a context-aware request and submits it using the http.DefaultClient.
func Get(ctx context.Context, url string) (*http.Response, error) {
	return do(ctx, http.MethodGet, url)
}

// Head creates a context-aware request and submits it using the http.DefaultClient.
func Head(ctx context.Context, url string) (*http.Response, error) {
	return do(ctx, http.MethodHead, url)
}

// Post creates a context-aware request and submits it using the http.DefaultClient.
func Post(ctx context.Context, url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}
