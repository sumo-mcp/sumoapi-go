package sumoapi

import (
	"context"
	"fmt"
	"net/url"
)

// ListRikishiMatchesAgainstOpponentAPI defines the methods available for listing matches for a single rikishi and opponent pair.
type ListRikishiMatchesAgainstOpponentAPI interface {
	// ListRikishiMatchesAgainstOpponent calls the GET /api/rikishi/{rikishiID}/matches/{opponentID} endpoint.
	//
	// Documented bugs:
	//   - The API accepts and takes into account the limit and skip inputs, but they are not documented in the API guide.
	//   - The API response does not return the limit and skip outputs like in other endpoints.
	//   - The API response does not return the match ID like in the GET /api/kimarite/{kimariteID} endpoint.
	ListRikishiMatchesAgainstOpponent(ctx context.Context, req ListRikishiMatchesAgainstOpponentRequest) (*ListRikishiMatchesAgainstOpponentResponse, error)
}

// ListRikishiMatchesAgainstOpponentRequest represents the request parameters for the ListRikishiMatchesAgainstOpponent method.
type ListRikishiMatchesAgainstOpponentRequest struct {
	RikishiID  int      `json:"rikishiId" jsonschema:"The unique identifier for the rikishi (sumo wrestler)."`
	OpponentID int      `json:"opponentId" jsonschema:"The unique identifier for the opponent rikishi (sumo wrestler)."`
	BashoID    *BashoID `json:"bashoId,omitempty" jsonschema:"The ID of the basho (sumo tournament) to filter matches by, in the format YYYYMM."`
	Limit      int      `json:"limit,omitempty" jsonschema:"The maximum number of results to return."`
	Skip       int      `json:"skip,omitempty" jsonschema:"The number of results to skip over for pagination."`
}

// ListRikishiMatchesAgainstOpponentResponse represents the response from the ListRikishiMatchesAgainstOpponent method.
type ListRikishiMatchesAgainstOpponentResponse struct {
	RikishiWins    int            `json:"rikishiWins" jsonschema:"The total number of wins for the rikishi against the opponent in the matching results."`
	OpponentWins   int            `json:"opponentWins" jsonschema:"The total number of wins for the opponent against the rikishi in the matching results."`
	KimariteWins   map[string]int `json:"kimariteWins" jsonschema:"A breakdown of wins by kimarite (winning technique) for the rikishi against the opponent in the matching results."`
	KimariteLosses map[string]int `json:"kimariteLosses" jsonschema:"A breakdown of losses by kimarite (winning technique) for the rikishi against the opponent in the matching results."`
	Limit          int            `json:"limit" jsonschema:"The maximum number of results that were returned."`
	Skip           int            `json:"skip" jsonschema:"The number of results that were skipped over."`
	Total          int            `json:"total" jsonschema:"The total number of matching results."`
	Matches        []Match        `json:"matches,omitempty" jsonschema:"The list of matches matching the filters."`
}

func (c *client) ListRikishiMatchesAgainstOpponent(ctx context.Context, req ListRikishiMatchesAgainstOpponentRequest) (*ListRikishiMatchesAgainstOpponentResponse, error) {
	query := make(url.Values)
	if req.BashoID != nil {
		query.Set("bashoId", req.BashoID.String())
	}
	if req.Limit > 0 {
		query.Set("limit", fmt.Sprint(req.Limit))
	}
	if req.Skip > 0 {
		query.Set("skip", fmt.Sprint(req.Skip))
	}
	path := fmt.Sprintf("/rikishi/%d/matches/%d", req.RikishiID, req.OpponentID)
	return getObject[ListRikishiMatchesAgainstOpponentResponse](ctx, c, path, query)
}
