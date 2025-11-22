package sumoapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// SearchRikishisAPI defines the methods available for searching Rikishis.
type SearchRikishisAPI interface {
	// SearchRikishis calls the /rikishis endpoint.
	SearchRikishis(ctx context.Context, req SearchRikishisRequest) (*SearchRikishisResponse, error)
}

// SearchRikishisRequest represents the request parameters for the SearchRikishis method.
type SearchRikishisRequest struct {
	Shikona             string `json:"shikona,omitempty" jsonschema:"The English shikona (ring name) to search for. Example: Terunofuji"`
	Heya                string `json:"heya,omitempty" jsonschema:"The full name in English of the heya (stable) to search for. Example: Isegahama"`
	SumoDBID            int    `json:"sumoDBID,omitempty" jsonschema:"The SumoDB ID to search for. Example: 11927 = Terunofuji"`
	OfficialID          int    `json:"officialID,omitempty" jsonschema:"The official Nihon Sumo Kyokai (Japan Sumo Association) ID to search for. Example: 3321 = Terunofuji"`
	IncludeRetired      bool   `json:"includeRetired,omitempty" jsonschema:"Whether to include retired rikishis in the results."`
	IncludeRanks        bool   `json:"includeRanks,omitempty" jsonschema:"Whether to include rank records over time in the Rikishi data."`
	IncludeShikonas     bool   `json:"includeShikonas,omitempty" jsonschema:"Whether to include shikona (ring name) records over time in the Rikishi data."`
	IncludeMeasurements bool   `json:"includeMeasurements,omitempty" jsonschema:"Whether to include measurement records over time in the Rikishi data."`
	Limit               int    `json:"limit,omitempty" jsonschema:"The maximum number of results to return."`
	Skip                int    `json:"skip,omitempty" jsonschema:"The number of results to skip over for pagination."`
}

// SearchRikishisResponse represents the response from the SearchRikishis method.
type SearchRikishisResponse struct {
	Limit    int       `json:"limit" jsonschema:"The maximum number of results as specified in the request."`
	Skip     int       `json:"skip" jsonschema:"The number of results to skip over as specified in the request."`
	Total    int       `json:"total" jsonschema:"The overall total number of Rikishis available in the API."`
	Rikishis []Rikishi `json:"records" jsonschema:"The list of Rikishis matching the search criteria."`
}

func (c *client) SearchRikishis(ctx context.Context, req SearchRikishisRequest) (*SearchRikishisResponse, error) {
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
		query.Set("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Skip > 0 {
		query.Set("skip", fmt.Sprintf("%d", req.Skip))
	}
	return doRequest[SearchRikishisResponse](ctx, c, http.MethodGet, "/rikishis", query)
}
