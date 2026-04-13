package sumoapi

import "context"

// ListRankChangesAPI defines the methods available for listing rikishi rank changes across bashos.
type ListRankChangesAPI interface {
	// ListRankChanges calls the GET /api/ranks endpoint.
	//
	// Documented bugs:
	//   - When filtering by bashoId, the API does not return results in any
	//     deterministic order: the first record changes depending on the
	//     limit value (e.g. limit=3 returns rikishiId 8850 first, while
	//     limit>=50 returns rikishiId 19 first). Pagination via skip still
	//     yields the full set of unique records, but the order in which they
	//     arrive is not stable.
	//   - When filtering by bashoId, the API rejects the sortOrder query
	//     parameter with 400 INVALID_QUERY_PARAMS, so callers cannot force a
	//     stable order on basho-scoped queries.
	ListRankChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Rank, error)
}

func (c *client) ListRankChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Rank, error) {
	return listRikishiChanges[Rank](ctx, c, "/ranks", req)
}
