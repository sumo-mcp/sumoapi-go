package integration_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_GetBashoWithTorikumi(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	bashoID := sumoapi.BashoID{Year: 2025, Month: 11}
	resp, err := client.GetBashoWithTorikumi(context.Background(), sumoapi.GetBashoWithTorikumiRequest{
		BashoID:  bashoID,
		Division: "Makuuchi",
		Day:      1,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.ID).To(Equal(bashoID))

	startDate, _ := time.Parse(time.RFC3339, "2025-11-09T00:00:00Z")
	endDate, _ := time.Parse(time.RFC3339, "2025-11-23T00:00:00Z")
	g.Expect(resp.StartDate).ToNot(BeNil())
	g.Expect(*resp.StartDate).To(Equal(startDate))
	g.Expect(resp.EndDate).ToNot(BeNil())
	g.Expect(*resp.EndDate).To(Equal(endDate))

	g.Expect(resp.Yusho).To(HaveLen(6))
	g.Expect(resp.SpecialPrizes).To(HaveLen(5))

	g.Expect(resp.Torikumi).To(HaveLen(21))

	match := resp.Torikumi[0]
	g.Expect(match.ID).ToNot(BeNil())
	g.Expect(match.ID.BashoID).To(Equal(bashoID))
	g.Expect(match.ID.Day).To(Equal(1))
	g.Expect(match.ID.MatchNumber).To(Equal(0)) // Bug: The match number in the ID starts at 0.
	g.Expect(match.ID.EastID).To(Equal(111))
	g.Expect(match.ID.WestID).To(Equal(164))
	g.Expect(match.BashoID).To(Equal(bashoID))
	g.Expect(match.Division).To(Equal("Makuuchi"))
	g.Expect(match.Day).To(Equal(1))
	g.Expect(match.MatchNumber).To(Equal(1))
	g.Expect(match.EastID).To(Equal(111))
	g.Expect(match.EastShikona).To(Equal("Hitoshi"))
	g.Expect(match.EastRank).To(Equal("Juryo 1 East"))
	g.Expect(match.WestID).To(Equal(164))
	g.Expect(match.WestShikona).To(Equal("Asakoryu"))
	g.Expect(match.WestRank).To(Equal("Maegashira 17 West"))
	g.Expect(match.Kimarite).To(Equal("tsukidashi"))
	g.Expect(match.WinnerID).To(Equal(164))
	g.Expect(match.WinnerEnglish).To(Equal("Asakoryu"))
	g.Expect(match.WinnerJapanese).To(Equal("朝紅龍"))

	lastMatch := resp.Torikumi[20]
	g.Expect(lastMatch.ID).ToNot(BeNil())
	g.Expect(lastMatch.ID.MatchNumber).To(Equal(20)) // Bug: The match number in the ID starts at 0.
	g.Expect(lastMatch.MatchNumber).To(Equal(21))
	g.Expect(lastMatch.EastID).To(Equal(8850))
	g.Expect(lastMatch.EastShikona).To(Equal("Onosato"))
	g.Expect(lastMatch.EastRank).To(Equal("Yokozuna 1 East"))
	g.Expect(lastMatch.WestID).To(Equal(44))
	g.Expect(lastMatch.WestShikona).To(Equal("Takayasu"))
	g.Expect(lastMatch.WestRank).To(Equal("Komusubi 1 West"))
	g.Expect(lastMatch.Kimarite).To(Equal("yorikiri"))
	g.Expect(lastMatch.WinnerID).To(Equal(8850))
	g.Expect(lastMatch.WinnerEnglish).To(Equal("Onosato"))
	g.Expect(lastMatch.WinnerJapanese).To(Equal("大の里"))
}
