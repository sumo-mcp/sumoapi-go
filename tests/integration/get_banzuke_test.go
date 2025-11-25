package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_GetBanzuke(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	bashoID := sumoapi.BashoID{Year: 2025, Month: 11}
	resp, err := client.GetBanzuke(context.Background(), sumoapi.GetBanzukeRequest{
		BashoID:  bashoID,
		Division: "Makuuchi",
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.BashoID).To(Equal(bashoID))
	g.Expect(resp.Division).To(Equal("Makuuchi"))
	g.Expect(resp.East).To(HaveLen(22))
	g.Expect(resp.West).To(HaveLen(20))

	g.Expect(resp.East[0].Side).To(Equal("East"))
	g.Expect(resp.East[0].RikishiID).To(Equal(8850))
	g.Expect(resp.East[0].ShikonaEnglish).To(Equal("Onosato"))
	g.Expect(resp.East[0].ShikonaJapanese).To(Equal("大の里　泰輝"))
	g.Expect(resp.East[0].HumanReadableRankName).To(Equal("Yokozuna 1 East"))
	g.Expect(resp.East[0].NumericRankName).To(Equal(101))
	g.Expect(resp.East[0].Wins).To(Equal(11))
	g.Expect(resp.East[0].Losses).To(Equal(4))
	g.Expect(resp.East[0].Absences).To(Equal(0))
	g.Expect(resp.East[0].Matches).To(HaveLen(15))

	g.Expect(resp.East[0].Matches[0].Result).To(Equal("win"))
	g.Expect(resp.East[0].Matches[0].OpponentShikonaEnglish).To(Equal("Takayasu"))
	g.Expect(resp.East[0].Matches[0].OpponentShikonaJapanese).To(Equal("高安"))
	g.Expect(resp.East[0].Matches[0].OpponentID).To(Equal(44))
	g.Expect(resp.East[0].Matches[0].Kimarite).To(Equal("yorikiri"))

	g.Expect(resp.East[0].Matches[14].Result).To(Equal("fusen loss"))
	g.Expect(resp.East[0].Matches[14].OpponentShikonaEnglish).To(Equal("Hoshoryu"))
	g.Expect(resp.East[0].Matches[14].Kimarite).To(Equal("fusen"))

	g.Expect(resp.West[0].Side).To(Equal("West"))
	g.Expect(resp.West[0].RikishiID).To(Equal(19))
	g.Expect(resp.West[0].ShikonaEnglish).To(Equal("Hoshoryu"))
	g.Expect(resp.West[0].ShikonaJapanese).To(Equal("豊昇龍　智勝"))
	g.Expect(resp.West[0].HumanReadableRankName).To(Equal("Yokozuna 1 West"))
	g.Expect(resp.West[0].NumericRankName).To(Equal(101))
	g.Expect(resp.West[0].Wins).To(Equal(12))
	g.Expect(resp.West[0].Losses).To(Equal(3))
	g.Expect(resp.West[0].Absences).To(Equal(0))

	g.Expect(resp.West[0].Matches[14].Result).To(Equal("fusen win"))
	g.Expect(resp.West[0].Matches[14].OpponentShikonaEnglish).To(Equal("Onosato"))
	g.Expect(resp.West[0].Matches[14].Kimarite).To(Equal("fusen"))

	var meisei *sumoapi.RikishiBanzuke
	for i := range resp.East {
		if resp.East[i].RikishiID == 38 {
			meisei = &resp.East[i]
			break
		}
	}
	g.Expect(meisei).ToNot(BeNil())
	g.Expect(meisei.ShikonaEnglish).To(Equal("Meisei"))
	g.Expect(meisei.Absences).To(Equal(9))
	g.Expect(meisei.Matches[0].Result).To(Equal("absent"))

	resultTypes := make(map[string]bool)
	for _, wrestler := range resp.East {
		for _, match := range wrestler.Matches {
			resultTypes[match.Result] = true
		}
	}
	for _, wrestler := range resp.West {
		for _, match := range wrestler.Matches {
			resultTypes[match.Result] = true
		}
	}
	g.Expect(resultTypes).To(HaveKey("win"))
	g.Expect(resultTypes).To(HaveKey("loss"))
	g.Expect(resultTypes).To(HaveKey("absent"))
	g.Expect(resultTypes).To(HaveKey("fusen win"))
	g.Expect(resultTypes).To(HaveKey("fusen loss"))
}
