package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// ListRikishiChangesRequest represents a request to list Rikishi changes with optional filters.
type ListRikishiChangesRequest struct {
	RikishiID int      `json:"rikishiId,omitempty" jsonschema:"The ID of the rikishi whose changes are to be listed."`
	BashoID   *BashoID `json:"bashoId,omitempty" jsonschema:"The ID of the basho for which rikishi changes are to be listed."`
	SortOrder string   `json:"sortOrder,omitempty" jsonschema:"The order in which to sort the results by Basho (sumo tournament). Valid values are 'asc' for ascending and 'desc' for descending. Default is 'desc'."`
}

func listRikishiChanges[obj any](ctx context.Context, c *client, path string, req ListRikishiChangesRequest) ([]obj, error) {
	query := make(url.Values)
	if req.RikishiID > 0 {
		query.Set("rikishiId", fmt.Sprint(req.RikishiID))
	}
	if req.BashoID != nil {
		query.Set("bashoId", req.BashoID.String())
	}
	if order := getSortOrder(req.SortOrder); order != "" {
		query.Set("sortOrder", order)
	}
	return listObjects[obj](ctx, c, path, query)
}
