package sumoapi

import "context"

// ListShikonaChangesAPI defines the methods available for listing rikishi shikona changes across bashos.
type ListShikonaChangesAPI interface {
	// ListShikonaChanges calls the GET /api/shikonas endpoint.
	ListShikonaChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Shikona, error)
}

func (c *client) ListShikonaChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Shikona, error) {
	return listRikishiChanges[Shikona](ctx, c, "/shikonas", req)
}
