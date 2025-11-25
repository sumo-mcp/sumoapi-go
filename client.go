package sumoapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client is a client for the Sumo API.
type Client interface {
	SearchRikishiAPI
	GetRikishiAPI
	GetRikishiStatsAPI
	ListRikishiMatchesAPI
	ListRikishiMatchesAgainstOpponentAPI
	GetBashoAPI
	GetBanzukeAPI
	GetBashoWithTorikumiAPI
	ListKimariteAPI
	ListKimariteMatchesAPI
	ListMeasurementChangesAPI
	ListRankChangesAPI
	ListShikonaChangesAPI
}

// Error represents an error returned by the Sumo API.
type Error struct {
	StatusCode  int
	Body        []byte
	ReadBodyErr error
}

func (e *Error) Error() string {
	switch {
	case e.ReadBodyErr != nil:
		return fmt.Sprintf("sumoapi: received HTTP %d response; additionally, error reading response body: %v", e.StatusCode, e.ReadBodyErr)
	case len(e.Body) == 0:
		return fmt.Sprintf("sumoapi: received HTTP %d response with empty body", e.StatusCode)
	default:
		return fmt.Sprintf("sumoapi: received HTTP %d response: %s", e.StatusCode, string(e.Body))
	}
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

type client struct {
	httpClient *http.Client
}

func (c *client) doRequest(ctx context.Context, method, path string, query url.Values, obj any) ([]byte, error) {
	u := fmt.Sprintf("https://sumo-api.com/api%s", path)
	if len(query) > 0 {
		u = fmt.Sprintf("%s?%s", u, query.Encode())
	}

	var body io.Reader
	if obj != nil {
		b, err := json.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		body = strings.NewReader(string(b))
	}

	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}
	if obj != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}
	defer resp.Body.Close()

	status := resp.StatusCode
	if status < 200 || status >= 300 {
		b, readErr := io.ReadAll(resp.Body)
		return nil, &Error{
			StatusCode:  status,
			Body:        b,
			ReadBodyErr: readErr,
		}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return b, nil
}

func getObject[obj any](ctx context.Context, c *client, path string, query url.Values) (*obj, error) {
	b, err := c.doRequest(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return nil, err
	}
	var o obj
	if err := json.Unmarshal(b, &o); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &o, nil
}

func listObjects[obj any](ctx context.Context, c *client, path string, query url.Values) ([]obj, error) {
	b, err := c.doRequest(ctx, http.MethodGet, path, query, nil)
	if err != nil {
		return nil, err
	}
	var l []obj
	if err := json.Unmarshal(b, &l); err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return l, nil
}

func getSortOrder(order string) string {
	if o := strings.ToLower(strings.TrimSpace(order)); o == "asc" || o == "desc" {
		return o
	}
	return ""
}
