package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// SearchRikishiAPI defines the methods available for searching rikishi.
type SearchRikishiAPI interface {
	// SearchRikishi calls the GET /api/rikishis endpoint.
	SearchRikishi(ctx context.Context, req SearchRikishiRequest) (*SearchRikishiResponse, error)
}

// SearchRikishiRequest represents the request parameters for the SearchRikishi method.
type SearchRikishiRequest struct {
	Shikona             string `json:"shikona,omitempty" jsonschema:"The shikona (ring name) in English to search for. Example: Terunofuji"`
	Heya                string `json:"heya,omitempty" jsonschema:"The full name in English of the heya (stable) to search for. Example: Isegahama"`
	SumoDBID            int    `json:"sumoDBID,omitempty" jsonschema:"The SumoDB ID to search for. Example: 11927 = Terunofuji"`
	OfficialID          int    `json:"officialID,omitempty" jsonschema:"The official Nihon Sumo Kyokai (Japan Sumo Association) ID to search for. Example: 3321 = Terunofuji"`
	IncludeRetired      bool   `json:"includeRetired,omitempty" jsonschema:"Whether to include retired rikishi (sumo wrestlers) in the results."`
	IncludeRanks        bool   `json:"includeRanks,omitempty" jsonschema:"Whether to include rank records over time in the rikishi (sumo wrestler) data."`
	IncludeShikonas     bool   `json:"includeShikonas,omitempty" jsonschema:"Whether to include shikona (ring name) records over time in the rikishi (sumo wrestler) data."`
	IncludeMeasurements bool   `json:"includeMeasurements,omitempty" jsonschema:"Whether to include measurement records over time in the rikishi (sumo wrestler) data."`
	Limit               int    `json:"limit,omitempty" jsonschema:"The maximum number of results to return."`
	Skip                int    `json:"skip,omitempty" jsonschema:"The number of results to skip over for pagination."`
}

// SearchRikishiResponse represents the response from the SearchRikishi method.
type SearchRikishiResponse struct {
	Limit   int       `json:"limit" jsonschema:"The maximum number of results that were returned."`
	Skip    int       `json:"skip" jsonschema:"The number of results that were skipped over."`
	Total   int       `json:"total" jsonschema:"The total number of matching results."`
	Rikishi []Rikishi `json:"records,omitempty" jsonschema:"The list of rikishi (sumo wrestlers) matching the search criteria."`
}

func (c *client) SearchRikishi(ctx context.Context, req SearchRikishiRequest) (*SearchRikishiResponse, error) {
	query := make(url.Values)
	if req.Shikona != "" {
		query.Set("shikonaEn", req.Shikona)
	}
	if req.Heya != "" {
		query.Set("heya", req.Heya)
	}
	if req.SumoDBID > 0 {
		query.Set("sumodbId", fmt.Sprint(req.SumoDBID))
	}
	if req.OfficialID > 0 {
		query.Set("nskId", fmt.Sprint(req.OfficialID))
	}
	if req.IncludeRetired {
		query.Set("intai", "true")
	}
	if req.IncludeMeasurements {
		query.Set("measurements", "true")
	}
	if req.IncludeRanks {
		query.Set("ranks", "true")
	}
	if req.IncludeShikonas {
		query.Set("shikonas", "true")
	}
	if req.Limit > 0 {
		query.Set("limit", fmt.Sprint(req.Limit))
	}
	if req.Skip > 0 {
		query.Set("skip", fmt.Sprint(req.Skip))
	}
	return getObject[SearchRikishiResponse](ctx, c, "/rikishis", query)
}
