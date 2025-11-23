package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListRikishiMatches(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
		RikishiID: 45, // Terunofuji
		Limit:     1,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.Total).To(Equal(1))
	g.Expect(resp.Matches).To(HaveLen(1))

	// Bug: The limit and skip outputs are always 0.
	g.Expect(resp.Limit).To(Equal(0))
	g.Expect(resp.Skip).To(Equal(0))

	match := resp.Matches[0]
	g.Expect(match.BashoID).To(Equal(sumoapi.BashoID{Year: 2025, Month: 1}))
	g.Expect(match.Division).To(Equal("Makuuchi"))
	g.Expect(match.Day).To(Equal(5))
	g.Expect(match.MatchNumber).To(Equal(21))
	g.Expect(match.EastID).To(Equal(45))
	g.Expect(match.EastShikona).To(Equal("Terunofuji"))
	g.Expect(match.EastRank).To(Equal("Yokozuna 1 East"))
	g.Expect(match.WestID).To(Equal(56))
	g.Expect(match.WestShikona).To(Equal("Gonoyama"))
	g.Expect(match.WestRank).To(Equal("Maegashira 3 East"))
	g.Expect(match.Kimarite).To(Equal("fusen"))
	g.Expect(match.WinnerID).To(Equal(56))
	g.Expect(match.WinnerEnglish).To(Equal("Gonoyama"))
	g.Expect(match.WinnerJapanese).To(Equal(""))

	// Bug: The match ID is not returned.
	g.Expect(match.ID).To(BeNil())
}
