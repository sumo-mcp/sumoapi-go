package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_GetRikishiStats(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.GetRikishiStats(context.Background(), sumoapi.GetRikishiStatsRequest{
		RikishiID: 45, // Terunofuji
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())

	// Verify overall stats
	g.Expect(resp.Basho).To(Equal(81))
	g.Expect(resp.Yusho).To(Equal(13))
	g.Expect(resp.TotalMatches).To(Equal(798))
	g.Expect(resp.TotalWins).To(Equal(523))
	g.Expect(resp.TotalLosses).To(Equal(275))
	g.Expect(resp.TotalAbsences).To(Equal(231))

	// Verify sansho (special prizes)
	g.Expect(resp.Sansho).To(HaveLen(3))
	g.Expect(resp.Sansho["Gino-sho"]).To(Equal(3))
	g.Expect(resp.Sansho["Kanto-sho"]).To(Equal(3))
	g.Expect(resp.Sansho["Shukun-sho"]).To(Equal(3))

	// Verify basho by division
	g.Expect(resp.BashoByDivision["Jonokuchi"]).To(Equal(1))
	g.Expect(resp.BashoByDivision["Jonidan"]).To(Equal(2))
	g.Expect(resp.BashoByDivision["Sandanme"]).To(Equal(4))
	g.Expect(resp.BashoByDivision["Makushita"]).To(Equal(15))
	g.Expect(resp.BashoByDivision["Juryo"]).To(Equal(7))
	g.Expect(resp.BashoByDivision["Makuuchi"]).To(Equal(52))

	// Verify yusho by division
	g.Expect(resp.YushoByDivision["Makushita"]).To(Equal(1))
	g.Expect(resp.YushoByDivision["Juryo"]).To(Equal(2))
	g.Expect(resp.YushoByDivision["Makuuchi"]).To(Equal(10))

	// Verify wins by division
	g.Expect(resp.WinsByDivision["Jonokuchi"]).To(Equal(5))
	g.Expect(resp.WinsByDivision["Jonidan"]).To(Equal(13))
	g.Expect(resp.WinsByDivision["Sandanme"]).To(Equal(13))
	g.Expect(resp.WinsByDivision["Makushita"]).To(Equal(65))
	g.Expect(resp.WinsByDivision["Juryo"]).To(Equal(61))
	g.Expect(resp.WinsByDivision["Makuuchi"]).To(Equal(366))

	// Verify losses by division
	g.Expect(resp.LossByDivision["Jonokuchi"]).To(Equal(2))
	g.Expect(resp.LossByDivision["Jonidan"]).To(Equal(1))
	g.Expect(resp.LossByDivision["Sandanme"]).To(Equal(1))
	g.Expect(resp.LossByDivision["Makushita"]).To(Equal(26))
	g.Expect(resp.LossByDivision["Juryo"]).To(Equal(38))
	g.Expect(resp.LossByDivision["Makuuchi"]).To(Equal(207))

	// Verify absences by division
	g.Expect(resp.AbsenceByDivision["Jonokuchi"]).To(Equal(0))
	g.Expect(resp.AbsenceByDivision["Jonidan"]).To(Equal(0))
	g.Expect(resp.AbsenceByDivision["Sandanme"]).To(Equal(14))
	g.Expect(resp.AbsenceByDivision["Makushita"]).To(Equal(14))
	g.Expect(resp.AbsenceByDivision["Juryo"]).To(Equal(6))
	g.Expect(resp.AbsenceByDivision["Makuuchi"]).To(Equal(197))

	// Verify total matches by division
	g.Expect(resp.TotalMatchesByDivision["Jonokuchi"]).To(Equal(7))
	g.Expect(resp.TotalMatchesByDivision["Jonidan"]).To(Equal(14))
	g.Expect(resp.TotalMatchesByDivision["Sandanme"]).To(Equal(14))
	g.Expect(resp.TotalMatchesByDivision["Makushita"]).To(Equal(91))
	g.Expect(resp.TotalMatchesByDivision["Juryo"]).To(Equal(99))
	g.Expect(resp.TotalMatchesByDivision["Makuuchi"]).To(Equal(573))
}
