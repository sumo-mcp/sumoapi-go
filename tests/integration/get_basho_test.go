package integration_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_GetBasho(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	bashoID := sumoapi.BashoID{Year: 2025, Month: 11}
	resp, err := client.GetBasho(context.Background(), sumoapi.GetBashoRequest{
		BashoID: bashoID,
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
	g.Expect(resp.Yusho[0].Type).To(Equal("Makuuchi"))
	g.Expect(resp.Yusho[0].RikishiID).To(Equal(8854))
	g.Expect(resp.Yusho[0].ShikonaEnglish).To(Equal("Aonishiki"))
	g.Expect(resp.Yusho[0].ShikonaJapanese).To(Equal("安青錦　新大"))

	g.Expect(resp.Yusho[1].Type).To(Equal("Juryo"))
	g.Expect(resp.Yusho[1].RikishiID).To(Equal(9051))
	g.Expect(resp.Yusho[1].ShikonaEnglish).To(Equal("Fujiryoga"))

	g.Expect(resp.Yusho[2].Type).To(Equal("Makushita"))
	g.Expect(resp.Yusho[2].RikishiID).To(Equal(8865))

	g.Expect(resp.Yusho[3].Type).To(Equal("Sandanme"))
	g.Expect(resp.Yusho[3].RikishiID).To(Equal(9088))

	g.Expect(resp.Yusho[4].Type).To(Equal("Jonidan"))
	g.Expect(resp.Yusho[4].RikishiID).To(Equal(9099))

	g.Expect(resp.Yusho[5].Type).To(Equal("Jonokuchi"))
	g.Expect(resp.Yusho[5].RikishiID).To(Equal(228))

	g.Expect(resp.SpecialPrizes).To(HaveLen(5))
	g.Expect(resp.SpecialPrizes[0].Type).To(Equal("Shukun-sho"))
	g.Expect(resp.SpecialPrizes[0].RikishiID).To(Equal(8854))
	g.Expect(resp.SpecialPrizes[0].ShikonaEnglish).To(Equal("Aonishiki"))

	g.Expect(resp.SpecialPrizes[1].Type).To(Equal("Kanto-sho"))
	g.Expect(resp.SpecialPrizes[1].RikishiID).To(Equal(7))
	g.Expect(resp.SpecialPrizes[1].ShikonaEnglish).To(Equal("Kirishima"))

	g.Expect(resp.SpecialPrizes[2].Type).To(Equal("Kanto-sho"))
	g.Expect(resp.SpecialPrizes[2].RikishiID).To(Equal(11))
	g.Expect(resp.SpecialPrizes[2].ShikonaEnglish).To(Equal("Ichiyamamoto"))

	g.Expect(resp.SpecialPrizes[3].Type).To(Equal("Gino-sho"))
	g.Expect(resp.SpecialPrizes[3].RikishiID).To(Equal(8854))
	g.Expect(resp.SpecialPrizes[3].ShikonaEnglish).To(Equal("Aonishiki"))

	g.Expect(resp.SpecialPrizes[4].Type).To(Equal("Gino-sho"))
	g.Expect(resp.SpecialPrizes[4].RikishiID).To(Equal(8857))
	g.Expect(resp.SpecialPrizes[4].ShikonaEnglish).To(Equal("Yoshinofuji"))
}
