package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_SearchRikishis(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
		Limit:               1,
		Shikona:             "Hakuho Sho",
		IncludeRetired:      true,
		IncludeRanks:        true,
		IncludeShikonas:     true,
		IncludeMeasurements: true,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.Rikishis).To(HaveLen(1))

	hakuho := resp.Rikishis[0]
	g.Expect(hakuho.ShikonaEnglish).To(Equal("Hakuho Sho"))
	g.Expect(hakuho.Heya).To(Equal("Miyagino"))
	g.Expect(hakuho.Debut).To(Equal(&sumoapi.BashoID{Year: 2001, Month: 3}))
	g.Expect(len(hakuho.RankHistory)).To(BeNumerically(">", 0))
	g.Expect(len(hakuho.ShikonaHistory)).To(BeNumerically(">", 0))
	g.Expect(len(hakuho.MeasurementHistory)).To(BeNumerically(">", 0))
}
