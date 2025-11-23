package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// GetRikishiAPI defines the methods available for retrieving a single Rikishi.
type GetRikishiAPI interface {
	// GetRikishi calls the GET /api/rikishi/{rikishiID} endpoint.
	GetRikishi(ctx context.Context, req GetRikishiRequest) (*Rikishi, error)
}

// GetRikishiRequest represents the request parameters for the GetRikishi method.
type GetRikishiRequest struct {
	RikishiID           int  `json:"rikishiId" jsonschema:"The unique identifier of the Rikishi to retrieve. Example: 45 = Terunofuji"`
	IncludeRanks        bool `json:"includeRanks,omitempty" jsonschema:"Whether to include rank records over time in the Rikishi data."`
	IncludeShikonas     bool `json:"includeShikonas,omitempty" jsonschema:"Whether to include shikona (ring name) records over time in the Rikishi data."`
	IncludeMeasurements bool `json:"includeMeasurements,omitempty" jsonschema:"Whether to include measurement records over time in the Rikishi data."`
}

func (c *client) GetRikishi(ctx context.Context, req GetRikishiRequest) (*Rikishi, error) {
	query := make(url.Values)
	if req.IncludeMeasurements {
		query.Set("measurements", "true")
	}
	if req.IncludeRanks {
		query.Set("ranks", "true")
	}
	if req.IncludeShikonas {
		query.Set("shikonas", "true")
	}
	path := fmt.Sprintf("/rikishi/%d", req.RikishiID)
	return getObject[Rikishi](ctx, c, path, query)
}
