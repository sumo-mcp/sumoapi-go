package sumoapi

import (
	"context"
	"fmt"
)

// GetBashoAPI defines the methods available for retrieving a basho.
type GetBashoAPI interface {
	// GetBasho calls the GET /api/basho/{bashoID} endpoint.
	GetBasho(ctx context.Context, req GetBashoRequest) (*Basho, error)
}

// GetBashoRequest represents the request parameters for the GetBasho method.
type GetBashoRequest struct {
	BashoID BashoID `json:"bashoId" jsonschema:"The unique identifier of the basho (sumo tournament) to retrieve. Format: YYYYMM, e.g., 202401 for the January 2024 basho."`
}

func (c *client) GetBasho(ctx context.Context, req GetBashoRequest) (*Basho, error) {
	path := fmt.Sprintf("/basho/%s", req.BashoID.String())
	return getObject[Basho](ctx, c, path, nil)
}
