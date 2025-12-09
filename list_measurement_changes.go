package sumoapi

import "context"

// ListMeasurementChangesAPI defines the methods available for listing rikishi measurement changes across bashos.
type ListMeasurementChangesAPI interface {
	// ListMeasurementChanges calls the GET /api/measurements endpoint.
	ListMeasurementChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Measurement, error)
}

func (c *client) ListMeasurementChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Measurement, error) {
	return listRikishiChanges[Measurement](ctx, c, "/measurements", req)
}
