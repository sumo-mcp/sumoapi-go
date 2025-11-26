package sumoapi

// Banzuke represents the ranking list of rikishi in a basho division.
type Banzuke struct {
	BashoID  BashoID          `json:"bashoId" jsonschema:"The unique identifier for the basho (sumo tournament)."`
	Division string           `json:"division" jsonschema:"The division of the basho (sumo tournament). One of Makuuchi, Juryo, Makushita, Sandanme, Jonidan, Jonokuchi."`
	East     []RikishiBanzuke `json:"east,omitempty" jsonschema:"The banzuke (ranking list) for the east side of the division."`
	West     []RikishiBanzuke `json:"west,omitempty" jsonschema:"The banzuke (ranking list) for the west side of the division."`
}

// RikishiBanzuke represents a rikishi's ranking information in a basho division.
type RikishiBanzuke struct {
	Side                  string                `json:"side" jsonschema:"The side of the rikishi in the banzuke (ranking list). Either East or West."`
	RikishiID             int                   `json:"rikishiID" jsonschema:"The unique identifier for the rikishi (sumo wrestler)."`
	ShikonaEnglish        string                `json:"shikonaEn" jsonschema:"The shikona (ring name) in English of the rikishi."`
	ShikonaJapanese       string                `json:"shikonaJp" jsonschema:"The shikona (ring name) in Japanese of the rikishi."`
	HumanReadableRankName string                `json:"rank" jsonschema:"The human-readable name of the rank (e.g., Maegashira 1 East, Ozeki 2 West) of the rikishi (sumo wrestler) in the specific basho (sumo tournament)."`
	NumericRankName       int                   `json:"rankValue" jsonschema:"The numeric name of the rank of the rikishi (sumo wrestler) in the specific basho (sumo tournament)."`
	Wins                  int                   `json:"wins" jsonschema:"The number of wins the rikishi (sumo wrestler) achieved in the specific basho (sumo tournament)."`
	Losses                int                   `json:"losses" jsonschema:"The number of losses the rikishi (sumo wrestler) had in the specific basho (sumo tournament)."`
	Absences              int                   `json:"absences" jsonschema:"The number of absences the rikishi (sumo wrestler) had in the specific basho (sumo tournament)."`
	Matches               []RikishiBanzukeMatch `json:"record,omitempty" jsonschema:"The list of matches the rikishi (sumo wrestler) had or will have in the specific basho (sumo tournament)."`
}

// RikishiBanzukeMatch represents a match against an opponent in the banzuke.
type RikishiBanzukeMatch struct {
	OpponentShikonaEnglish  string `json:"opponentShikonaEn" jsonschema:"The shikona (ring name) in English of the opponent rikishi (sumo wrestler)."`
	OpponentShikonaJapanese string `json:"opponentShikonaJp" jsonschema:"The shikona (ring name) in Japanese of the opponent rikishi (sumo wrestler)."`
	OpponentID              int    `json:"opponentID" jsonschema:"The unique identifier for the opponent rikishi (sumo wrestler)."`
	Result                  string `json:"result,omitempty" jsonschema:"The result of the match for the rikishi (sumo wrestler). One of win, loss, absent, fusen win (forfeit win), fusen loss (forfeit loss). This field may be omitted if the match has not yet occurred."`
	Kimarite                string `json:"kimarite,omitempty" jsonschema:"The kimarite (technique) used in the match, if the match has already occurred."`
}
