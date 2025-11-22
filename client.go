package sumoapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client is a client for the Sumo API.
type Client interface {
}

type client struct {
	httpClient *http.Client
}

// Option is a function that configures a Client.
type Option func(*client)

// WithHTTPClient sets a custom HTTP client for the Sumo API client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.httpClient = httpClient
	}
}

// New creates a new Client with the given options.
func New(opts ...Option) Client {
	client := &client{
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func doRequest[obj any](ctx context.Context, c *client, method, path string, query url.Values) (*obj, error) {
	u := fmt.Sprintf("https://sumo-api.com/api%s", path)
	if len(query) > 0 {
		u = fmt.Sprintf("%s?%s", u, query.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, method, u, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	if status < 200 || status >= 300 {
		b, readErr := io.ReadAll(resp.Body)
		switch {
		case readErr != nil:
			return nil, fmt.Errorf("received HTTP %d response; additionally, error reading response body: %w", status, readErr)
		case len(b) == 0:
			return nil, fmt.Errorf("received HTTP %d response with empty body", status)
		default:
			return nil, fmt.Errorf("received HTTP %d response: %s", status, string(b))
		}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var o obj
	if err := json.Unmarshal(b, &o); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return &o, nil
}
