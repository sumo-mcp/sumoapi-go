package sumoapi

// Shikona represents a rikishi's shikona in a specific basho.
type Shikona struct {
	ID              RikishiChangeID `json:"id" jsonschema:"The unique identifier for the shikona in the format {bashoID}-{rikishiID} where {bashoID} is in the format YYYYMM and {rikishiID} is the unique identifier for the rikishi in the API."`
	BashoID         BashoID         `json:"bashoId" jsonschema:"The ID of the basho (sumo tournament) in the format YYYYMM."`
	RikishiID       int             `json:"rikishiId" jsonschema:"The unique identifier for the rikishi (sumo wrestler) in the API."`
	ShikonaEnglish  string          `json:"shikonaEn,omitempty" jsonschema:"The shikona (ring name) in English of the rikishi in the specific basho (sumo tournament)."`
	ShikonaJapanese string          `json:"shikonaJp,omitempty" jsonschema:"The shikona (ring name) in Japanese of the rikishi in the specific basho (sumo tournament)."`
}
