package sumoapi

import "context"

// ListShikonaChangesAPI defines the methods available for listing Rikishi shikona (ring name) changes across bashos (sumo tournaments).
type ListShikonaChangesAPI interface {
	// ListShikonaChanges calls the GET /api/shikonas endpoint.
	ListShikonaChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Shikona, error)
}

func (c *client) ListShikonaChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Shikona, error) {
	return listRikishiChanges[Shikona](ctx, c, "/shikonas", req)
}
