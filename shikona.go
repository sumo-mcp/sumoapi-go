package sumoapi

// Shikona represents a sumo wrestler's shikona (ring name) in a specific Basho (sumo tournament).
type Shikona struct {
	ID              RikishiChangeID `json:"id" jsonschema:"The unique identifier for the Shikona in the format {bashoID}-{rikishiID} where {bashoID} is in the format YYYYMM and {rikishiID} is the unique identifier for the Rikishi in the API."`
	BashoID         BashoID         `json:"bashoId" jsonschema:"The ID of the Basho (sumo tournament) in the format YYYYMM."`
	RikishiID       int             `json:"rikishiId" jsonschema:"The unique identifier for the Rikishi (sumo wrestler) in the API."`
	ShikonaEnglish  string          `json:"shikonaEn,omitempty" jsonschema:"The shikona (ring name) in English of the Rikishi in the specific Basho (sumo tournament)."`
	ShikonaJapanese string          `json:"shikonaJp,omitempty" jsonschema:"The shikona (ring name) in Japanese of the Rikishi in the specific Basho (sumo tournament)."`
}
