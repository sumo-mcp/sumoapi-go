package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListKimariteMatches(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
		Kimarite: "tsumatori",
		Limit:    1,
		Skip:     1,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.Limit).To(Equal(1))
	g.Expect(resp.Skip).To(Equal(1))
	g.Expect(resp.Total).To(BeNumerically(">", 1))
	g.Expect(resp.Matches).To(HaveLen(1))

	bashoID := sumoapi.BashoID{Year: 1958, Month: 11}

	match := resp.Matches[0]
	g.Expect(match.ID).NotTo(BeNil())
	g.Expect(match.ID.BashoID).To(Equal(bashoID))
	g.Expect(match.ID.Day).To(Equal(7))
	g.Expect(match.ID.MatchNumber).To(Equal(15))
	g.Expect(match.ID.EastID).To(Equal(1339))
	g.Expect(match.ID.WestID).To(Equal(1351))
	g.Expect(match.BashoID).To(Equal(bashoID))
	g.Expect(match.Division).To(Equal("Juryo"))
	g.Expect(match.Day).To(Equal(7))
	g.Expect(match.MatchNumber).To(Equal(15))
	g.Expect(match.EastID).To(Equal(1339))
	g.Expect(match.EastShikona).To(Equal("Nanatsuumi"))
	g.Expect(match.EastRank).To(Equal("Juryo 9 East"))
	g.Expect(match.WestID).To(Equal(1351))
	g.Expect(match.WestShikona).To(Equal("Maegashio"))
	g.Expect(match.WestRank).To(Equal("Juryo 13 West"))
	g.Expect(match.Kimarite).To(Equal("tsumatori"))
	g.Expect(match.WinnerID).To(Equal(1339))
	g.Expect(match.WinnerEnglish).To(Equal("Nanatsuumi"))
	g.Expect(match.WinnerJapanese).To(Equal(""))
}
