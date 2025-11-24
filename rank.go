package sumoapi

// Rank represents a rikishi's rank in a specific basho.
type Rank struct {
	ID                RikishiChangeID `json:"id" jsonschema:"The unique identifier for the rank in the format {bashoID}-{rikishiID} where {bashoID} is in the format YYYYMM and {rikishiID} is the unique identifier for the rikishi in the API."`
	BashoID           BashoID         `json:"bashoId" jsonschema:"The ID of the basho (sumo tournament) in the format YYYYMM."`
	RikishiID         int             `json:"rikishiId" jsonschema:"The unique identifier for the rikishi (sumo wrestler) in the API."`
	HumanReadableName string          `json:"rank,omitempty" jsonschema:"The human-readable name of the rank (e.g., Maegashira 1 East, Ozeki 2 West) of the rikishi in the specific basho (sumo tournament)."`
	NumericName       int             `json:"rankValue,omitempty" jsonschema:"The numeric name of the rank of the rikishi in the specific basho (sumo tournament)."`
}
