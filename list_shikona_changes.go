package sumoapi

import "context"

// ListShikonaChangesAPI defines the methods available for listing rikishi shikona changes across bashos.
type ListShikonaChangesAPI interface {
	// ListShikonaChanges calls the GET /api/shikonas endpoint.
	//
	// Documented bugs:
	//   - When filtering by bashoId, the API rejects the sortOrder query
	//     parameter with 400 INVALID_QUERY_PARAMS, so callers cannot force a
	//     stable order on basho-scoped queries.
	ListShikonaChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Shikona, error)
}

func (c *client) ListShikonaChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Shikona, error) {
	return listRikishiChanges[Shikona](ctx, c, "/shikonas", req)
}
