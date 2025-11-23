package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// ListKimariteMatchesAPI defines the methods available for listing matches for a single kimarite.
type ListKimariteMatchesAPI interface {
	// ListKimariteMatches calls the GET /api/kimarite/{kimarite} endpoint.
	//
	// Potential improvements:
	//   - Support filtering by basho ID like the other match listing endpoints.
	ListKimariteMatches(ctx context.Context, req ListKimariteMatchesRequest) (*ListKimariteMatchesResponse, error)
}

// ListKimariteMatchesRequest represents the request parameters for the ListKimariteMatches method.
type ListKimariteMatchesRequest struct {
	Kimarite  string `json:"kimarite" jsonschema:"The kimarite (winning technique) to list matches for."`
	SortOrder string `json:"sortOrder,omitempty" jsonschema:"The order in which to sort the results by basho (sumo tournament), then day (and is not guaranteed to be the actual use order on that day). Valid values are 'asc' for ascending and 'desc' for descending. Default is 'asc'."`
	Limit     int    `json:"limit,omitempty" jsonschema:"The maximum number of results to return."`
	Skip      int    `json:"skip,omitempty" jsonschema:"The number of results to skip over for pagination."`
}

// ListKimariteMatchesResponse represents the response from the ListKimariteMatches method.
type ListKimariteMatchesResponse struct {
	Limit   int     `json:"limit" jsonschema:"The maximum number of results that were returned."`
	Skip    int     `json:"skip" jsonschema:"The number of results that were skipped over."`
	Total   int     `json:"total" jsonschema:"The total number of matching results."`
	Matches []Match `json:"records" jsonschema:"The list of matches matching the filters."`
}

func (c *client) ListKimariteMatches(ctx context.Context, req ListKimariteMatchesRequest) (*ListKimariteMatchesResponse, error) {
	query := make(url.Values)
	if order := getSortOrder(req.SortOrder); order != "" {
		query.Set("sortOrder", order)
	}
	if req.Limit > 0 {
		query.Set("limit", fmt.Sprint(req.Limit))
	}
	if req.Skip > 0 {
		query.Set("skip", fmt.Sprint(req.Skip))
	}
	path := fmt.Sprintf("/kimarite/%s", req.Kimarite)
	return getObject[ListKimariteMatchesResponse](ctx, c, path, query)
}
