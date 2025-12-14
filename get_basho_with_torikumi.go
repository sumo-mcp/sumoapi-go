package sumoapi

import (
	"context"
	"fmt"
)

// GetBashoWithTorikumiAPI defines the methods available for retrieving a basho.
type GetBashoWithTorikumiAPI interface {
	// GetBashoWithTorikumi calls the GET /api/basho/{bashoID}/torikumi/{division}/{day} endpoint.
	GetBashoWithTorikumi(ctx context.Context, req GetBashoWithTorikumiRequest) (*Basho, error)
}

// GetBashoWithTorikumiRequest represents the request parameters for the GetBashoWithTorikumi method.
type GetBashoWithTorikumiRequest struct {
	BashoID  BashoID `json:"bashoId" jsonschema:"The unique identifier of the basho (sumo tournament) to retrieve. Format: YYYYMM, e.g., 202401 for the January 2024 basho."`
	Division string  `json:"division" jsonschema:"The basho (sumo tournament) division to retrieve matches for. Valid values are Makuuchi, Juryo, Makushita, Sandanme, Jonidan, Jonokuchi."`
	Day      int     `json:"day" jsonschema:"The day of the basho (sumo tournament) to retrieve matches for. Values from 1 to 15 represent days, and 16 and above represent individual playoff matches."`
}

func (c *client) GetBashoWithTorikumi(ctx context.Context, req GetBashoWithTorikumiRequest) (*Basho, error) {
	path := fmt.Sprintf("/basho/%s/torikumi/%s/%d", req.BashoID.String(), req.Division, req.Day)
	return getObject[Basho](ctx, c, path, nil)
}
