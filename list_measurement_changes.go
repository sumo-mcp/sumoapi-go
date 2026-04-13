package sumoapi

import "context"

// ListMeasurementChangesAPI defines the methods available for listing rikishi measurement changes across bashos.
type ListMeasurementChangesAPI interface {
	// ListMeasurementChanges calls the GET /api/measurements endpoint.
	//
	// Documented bugs:
	//   - When filtering by bashoId, the API rejects the sortOrder query
	//     parameter with 400 INVALID_QUERY_PARAMS, so callers cannot force a
	//     stable order on basho-scoped queries.
	ListMeasurementChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Measurement, error)
}

func (c *client) ListMeasurementChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Measurement, error) {
	return listRikishiChanges[Measurement](ctx, c, "/measurements", req)
}
