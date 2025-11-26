package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// ListKimariteAPI defines the methods available for listing kimarite.
type ListKimariteAPI interface {
	// ListKimarite calls the GET /api/kimarite endpoint.
	ListKimarite(ctx context.Context, req ListKimariteRequest) (*ListKimariteResponse, error)
}

// ListKimariteRequest represents the request parameters for the ListKimarite method.
type ListKimariteRequest struct {
	SortField string `json:"sortField" jsonschema:"The field by which to sort the results. Valid values are 'kimarite', 'count' and 'lastUsage'."`
	SortOrder string `json:"sortOrder,omitempty" jsonschema:"The order in which to sort the results by the sort field. Valid values are 'asc' for ascending and 'desc' for descending. Default is 'asc'."`
	Limit     int    `json:"limit,omitempty" jsonschema:"The maximum number of results to return."`
	Skip      int    `json:"skip,omitempty" jsonschema:"The number of results to skip over for pagination."`
}

// ListKimariteResponse represents the response from the ListKimarite method.
type ListKimariteResponse struct {
	Limit     int        `json:"limit" jsonschema:"The maximum number of results that were returned."`
	Skip      int        `json:"skip" jsonschema:"The number of results that were skipped over."`
	SortField string     `json:"sortField" jsonschema:"The field by which the results are sorted. Values are 'kimarite', 'count' and 'lastUsage'."`
	SortOrder string     `json:"sortOrder" jsonschema:"The order in which the results are sorted by the sort field. Values are 'asc' for ascending and 'desc' for descending."`
	Kimarite  []Kimarite `json:"records,omitempty" jsonschema:"The list of kimarite (winning technique) records matching the filters."`
}

func (c *client) ListKimarite(ctx context.Context, req ListKimariteRequest) (*ListKimariteResponse, error) {
	query := make(url.Values)
	query.Set("sortField", req.SortField)
	if order := getSortOrder(req.SortOrder); order != "" {
		query.Set("sortOrder", order)
	}
	if req.Limit > 0 {
		query.Set("limit", fmt.Sprint(req.Limit))
	}
	if req.Skip > 0 {
		query.Set("skip", fmt.Sprint(req.Skip))
	}
	return getObject[ListKimariteResponse](ctx, c, "/kimarite", query)
}
