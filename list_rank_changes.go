package sumoapi

import "context"

// ListRankChangesAPI defines the methods available for listing rikishi rank changes across bashos.
type ListRankChangesAPI interface {
	// ListRankChanges calls the GET /api/ranks endpoint.
	ListRankChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Rank, error)
}

func (c *client) ListRankChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Rank, error) {
	return listRikishiChanges[Rank](ctx, c, "/ranks", req)
}
