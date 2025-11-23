package sumoapi

import (
	"context"
	"fmt"
)

// GetRikishiStatsAPI defines the methods available for retrieving statistics for a single rikishi.
type GetRikishiStatsAPI interface {
	// GetRikishiStats calls the GET /api/rikishi/{rikishiID}/stats endpoint.
	GetRikishiStats(ctx context.Context, req GetRikishiStatsRequest) (*GetRikishiStatsResponse, error)
}

// GetRikishiStatsRequest represents the request parameters for the GetRikishiStats method.
type GetRikishiStatsRequest struct {
	RikishiID int `json:"rikishiId" jsonschema:"The unique identifier of the rikishi (sumo wrestler) to retrieve. Example: 45 = Terunofuji"`
}

// GetRikishiStatsResponse represents the response from the GetRikishiStats method.
type GetRikishiStatsResponse struct {
	Basho                  int            `json:"basho,omitempty" jsonschema:"The number of official tournaments (basho) the rikishi (sumo wrestler) has participated in."`
	Yusho                  int            `json:"yusho,omitempty" jsonschema:"The number of tournament championships (yusho) the rikishi (sumo wrestler) has won."`
	TotalMatches           int            `json:"totalMatches,omitempty" jsonschema:"The total number of matches the rikishi (sumo wrestler) has had."`
	TotalWins              int            `json:"totalWins,omitempty" jsonschema:"The total number of wins the rikishi (sumo wrestler) has achieved."`
	TotalLosses            int            `json:"totalLosses,omitempty" jsonschema:"The total number of losses the rikishi (sumo wrestler) has suffered."`
	TotalAbsences          int            `json:"totalAbsences,omitempty" jsonschema:"The total number of absences the rikishi (sumo wrestler) has had."`
	Sansho                 map[string]int `json:"sansho,omitempty" jsonschema:"A mapping of sansho (special prize) names to the number of times the rikishi (sumo wrestler) has won each prize."`
	BashoByDivision        map[string]int `json:"bashoByDivision,omitempty" jsonschema:"A mapping of division names to the number of basho (sumo tournaments) the rikishi (sumo wrestler) has participated in each division."`
	YushoByDivision        map[string]int `json:"yushoByDivision,omitempty" jsonschema:"A mapping of division names to the number of yusho (tournament championships) the rikishi (sumo wrestler) has won in each division."`
	WinsByDivision         map[string]int `json:"winsByDivision,omitempty" jsonschema:"A mapping of division names to the number of wins the rikishi (sumo wrestler) has had in each division."`
	LossByDivision         map[string]int `json:"lossByDivision,omitempty" jsonschema:"A mapping of division names to the number of losses the rikishi (sumo wrestler) has had in each division."`
	AbsenceByDivision      map[string]int `json:"absenceByDivision,omitempty" jsonschema:"A mapping of division names to the number of absences the rikishi (sumo wrestler) has had in each division."`
	TotalMatchesByDivision map[string]int `json:"totalByDivision,omitempty" jsonschema:"A mapping of division names to the total number of matches the rikishi (sumo wrestler) has had in each division."`
}

func (c *client) GetRikishiStats(ctx context.Context, req GetRikishiStatsRequest) (*GetRikishiStatsResponse, error) {
	path := fmt.Sprintf("/rikishi/%d/stats", req.RikishiID)
	return getObject[GetRikishiStatsResponse](ctx, c, path, nil)
}
