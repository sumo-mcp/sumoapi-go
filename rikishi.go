package sumoapi

import "time"

// Rikishi represents a sumo wrestler.
type Rikishi struct {
	ID                 int           `json:"id" jsonschema:"The unique identifier for the rikishi (sumo wrestler) in the API."`
	SumoDBID           int           `json:"sumodbId,omitempty" jsonschema:"The SumoDB ID of the rikishi (sumo wrestler)."`
	OfficialID         int           `json:"nskId,omitempty" jsonschema:"The official Nihon Sumo Kyokai (Japan Sumo Association) ID of the rikishi (sumo wrestler)."`
	ShikonaEnglish     string        `json:"shikonaEn,omitempty" jsonschema:"The shikona (ring name) of the rikishi (sumo wrestler) in English."`
	ShikonaJapanese    string        `json:"shikonaJp,omitempty" jsonschema:"The shikona (ring name) of the rikishi (sumo wrestler) in Japanese."`
	CurrentRank        string        `json:"currentRank,omitempty" jsonschema:"The current rank of the rikishi (sumo wrestler)."`
	Heya               string        `json:"heya,omitempty" jsonschema:"The heya (stable) of the rikishi (sumo wrestler)."`
	BirthDate          *time.Time    `json:"birthDate,omitempty" jsonschema:"The birth date of the rikishi (sumo wrestler)."`
	Shusshin           string        `json:"shusshin,omitempty" jsonschema:"The place of birth of the rikishi (sumo wrestler)."`
	Height             float64       `json:"height,omitempty" jsonschema:"The height of the rikishi (sumo wrestler) in centimeters."`
	Weight             float64       `json:"weight,omitempty" jsonschema:"The weight of the rikishi (sumo wrestler) in kilograms."`
	Debut              *BashoID      `json:"debut,omitempty" jsonschema:"The ID of the basho (sumo tournament) when the rikishi (sumo wrestler) made their debut in the format YYYYMM."`
	Intai              *time.Time    `json:"intai,omitempty" jsonschema:"The retirement date of the rikishi (sumo wrestler), if retired."`
	RankHistory        []Rank        `json:"rankHistory,omitempty" jsonschema:"The historical rank records of the rikishi (sumo wrestler) over time. Each record is the rank of the rikishi (sumo wrestler) in a specific basho (sumo tournament)."`
	ShikonaHistory     []Shikona     `json:"shikonaHistory,omitempty" jsonschema:"The historical shikona (ring name) records of the rikishi (sumo wrestler) over time. Each record is the shikona (ring name) of the rikishi (sumo wrestler) in a specific basho (sumo tournament)."`
	MeasurementHistory []Measurement `json:"measurementHistory,omitempty" jsonschema:"The historical measurement records of the rikishi (sumo wrestler) over time. Each record includes height and weight of the rikishi (sumo wrestler) in a specific basho (sumo tournament)."`
	CreatedAt          *time.Time    `json:"createdAt,omitempty" jsonschema:"The timestamp when the rikishi (sumo wrestler) record was created in the API."`
	UpdatedAt          *time.Time    `json:"updatedAt,omitempty" jsonschema:"The timestamp when the rikishi (sumo wrestler) record was last updated in the API."`
}
