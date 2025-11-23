package sumoapi

import "context"

// ListMeasurementChangesAPI defines the methods available for listing rikishi measurement changes across bashos.
type ListMeasurementChangesAPI interface {
	// ListMeasurementChanges calls the GET /api/measurements endpoint.
	//
	// Documented bugs:
	//   - The API is ignoring the bashoId input. Measurement changes are returned for all bashos instead of the specified one.
	ListMeasurementChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Measurement, error)
}

func (c *client) ListMeasurementChanges(ctx context.Context, req ListRikishiChangesRequest) ([]Measurement, error) {
	return listRikishiChanges[Measurement](ctx, c, "/measurements", req)
}
