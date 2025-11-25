package sumoapi

import (
	"context"
	"fmt"
)

// GetBanzukeAPI defines the methods available for retrieving a banzuke.
type GetBanzukeAPI interface {
	// GetBanzuke calls the GET /api/basho/{bashoID}/banzuke/{division} endpoint.
	GetBanzuke(ctx context.Context, req GetBanzukeRequest) (*Banzuke, error)
}

// GetBanzukeRequest represents the request parameters for the GetBanzuke method.
type GetBanzukeRequest struct {
	BashoID  BashoID `json:"bashoId" jsonschema:"The unique identifier of the basho (sumo tournament) to retrieve the banzuke (ranking list) for. Format: YYYYMM, e.g., 202401 for the January 2024 basho."`
	Division string  `json:"division" jsonschema:"The division of the basho (sumo tournament) to retrieve the banzuke (ranking list) for. One of Makuuchi, Juryo, Makushita, Sandanme, Jonidan, Jonokuchi."`
}

func (c *client) GetBanzuke(ctx context.Context, req GetBanzukeRequest) (*Banzuke, error) {
	path := fmt.Sprintf("/basho/%s/banzuke/%s", req.BashoID.String(), req.Division)
	return getObject[Banzuke](ctx, c, path, nil)
}
