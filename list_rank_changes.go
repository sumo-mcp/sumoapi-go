package sumoapi

import "context"

// ListRankChangesAPI defines the methods available for listing rikishi rank changes across bashos.
type ListRankChangesAPI interface {
	// ListRankChanges calls the GET /api/ranks endpoint.
	//
	// Ordering: results are sorted by rank (always ascending) when filtering
	// by bashoId, and by bashoId when filtering by rikishiId.
	ListRankChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Rank, error)
}

func (c *client) ListRankChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Rank, error) {
	return listRikishiChanges[Rank](ctx, c, "/ranks", req)
}
