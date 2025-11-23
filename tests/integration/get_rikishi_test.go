package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_GetRikishi(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
		RikishiID:           45, // Terunofuji
		IncludeRanks:        true,
		IncludeShikonas:     true,
		IncludeMeasurements: true,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp.ID).To(Equal(45))
	g.Expect(resp.ShikonaEnglish).To(Equal("Terunofuji Haruo"))
	g.Expect(resp.ShikonaJapanese).To(Equal("照ノ富士　春雄"))
	g.Expect(resp.Heya).To(Equal("Isegahama"))

	g.Expect(resp.RankHistory).ToNot(BeEmpty())
	g.Expect(resp.ShikonaHistory).ToNot(BeEmpty())
	g.Expect(resp.MeasurementHistory).ToNot(BeEmpty())
}
