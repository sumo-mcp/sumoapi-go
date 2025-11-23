package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListRikishiMatchesAgainstOpponent(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.ListRikishiMatchesAgainstOpponent(context.Background(), sumoapi.ListRikishiMatchesAgainstOpponentRequest{
		RikishiID:  45,   // Terunofuji
		OpponentID: 3081, // Hakuho
		Limit:      1,
		Skip:       1,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.Total).To(Equal(1))
	g.Expect(resp.Matches).To(HaveLen(1))
	g.Expect(resp.RikishiWins).To(Equal(0))
	g.Expect(resp.OpponentWins).To(Equal(1))
	g.Expect(resp.KimariteWins).To(Equal(map[string]int{}))
	g.Expect(resp.KimariteLosses).To(Equal(map[string]int{"yorikiri": 1}))

	// Bug: The limit and skip outputs are always 0.
	g.Expect(resp.Limit).To(Equal(0))
	g.Expect(resp.Skip).To(Equal(0))

	match := resp.Matches[0]
	g.Expect(match.BashoID).To(Equal(sumoapi.BashoID{Year: 2017, Month: 5}))
	g.Expect(match.Division).To(Equal("Makuuchi"))
	g.Expect(match.Day).To(Equal(14))
	g.Expect(match.MatchNumber).To(Equal(19))
	g.Expect(match.EastID).To(Equal(45))
	g.Expect(match.EastShikona).To(Equal("Terunofuji"))
	g.Expect(match.EastRank).To(Equal("Ozeki 1 East"))
	g.Expect(match.WestID).To(Equal(3081))
	g.Expect(match.WestShikona).To(Equal("Hakuho"))
	g.Expect(match.WestRank).To(Equal("Yokozuna 2 West"))
	g.Expect(match.Kimarite).To(Equal("yorikiri"))
	g.Expect(match.WinnerID).To(Equal(3081))
	g.Expect(match.WinnerEnglish).To(Equal("Hakuho"))
	g.Expect(match.WinnerJapanese).To(Equal(""))

	// Bug: The match ID is not returned.
	g.Expect(match.ID).To(BeNil())
}
