package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// ListRikishiMatchesAPI defines the methods available for listing matches for a single rikishi.
type ListRikishiMatchesAPI interface {
	// ListRikishiMatches calls the GET /api/rikishi/{rikishiID}/matches endpoint.
	//
	// Documented bugs:
	//   - The API accepts and takes into account the limit and skip inputs, but they are not documented in the API guide.
	//   - The API response always returns 0 for the limit and skip outputs.
	//   - The API response does not return the match ID like in the GET /api/kimarite/{kimariteID} endpoint.
	ListRikishiMatches(ctx context.Context, req ListRikishiMatchesRequest) (*ListRikishiMatchesResponse, error)
}

// ListRikishiMatchesRequest represents the request parameters for the ListRikishiMatches method.
type ListRikishiMatchesRequest struct {
	RikishiID int      `json:"rikishiId" jsonschema:"The unique identifier for the rikishi (sumo wrestler)."`
	BashoID   *BashoID `json:"bashoId,omitempty" jsonschema:"The ID of the basho (sumo tournament) to filter matches by, in the format YYYYMM."`
	Limit     int      `json:"limit,omitempty" jsonschema:"The maximum number of results to return."`
	Skip      int      `json:"skip,omitempty" jsonschema:"The number of results to skip over for pagination."`
}

// ListRikishiMatchesResponse represents the response from the ListRikishiMatches method.
type ListRikishiMatchesResponse struct {
	Limit   int     `json:"limit" jsonschema:"The maximum number of results that were returned."`
	Skip    int     `json:"skip" jsonschema:"The number of results that were skipped over."`
	Total   int     `json:"total" jsonschema:"The total number of matching results."`
	Matches []Match `json:"records,omitempty" jsonschema:"The list of matches matching the filters."`
}

func (c *client) ListRikishiMatches(ctx context.Context, req ListRikishiMatchesRequest) (*ListRikishiMatchesResponse, error) {
	query := make(url.Values)
	if req.BashoID != nil {
		query.Set("bashoId", req.BashoID.String())
	}
	if req.Limit > 0 {
		query.Set("limit", fmt.Sprint(req.Limit))
	}
	if req.Skip > 0 {
		query.Set("skip", fmt.Sprint(req.Skip))
	}
	path := fmt.Sprintf("/rikishi/%d/matches", req.RikishiID)
	return getObject[ListRikishiMatchesResponse](ctx, c, path, query)
}
