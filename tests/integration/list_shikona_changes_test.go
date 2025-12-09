package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListShikonaChanges(t *testing.T) {
	client := sumoapi.New()

	t.Run("for rikishi", func(t *testing.T) {
		g := NewWithT(t)

		resp, err := client.ListShikonaChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			RikishiID: 3081, // Hakuho
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(1))

		expectedBashoID := sumoapi.BashoID{Year: 2001, Month: 3}

		g.Expect(resp[0].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
		g.Expect(resp[0].BashoID).To(Equal(expectedBashoID))
		g.Expect(resp[0].RikishiID).To(Equal(3081))
		g.Expect(resp[0].ShikonaEnglish).To(Equal("Hakuho Sho"))
		g.Expect(resp[0].ShikonaJapanese).To(Equal(""))
	})

	t.Run("for basho", func(t *testing.T) {
		g := NewWithT(t)

		bashoID := sumoapi.BashoID{
			Year:  2025,
			Month: 9,
		}

		resp, err := client.ListShikonaChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			BashoID: &bashoID,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(10))

		g.Expect(resp[0].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 8859}))
		g.Expect(resp[0].BashoID).To(Equal(bashoID))
		g.Expect(resp[0].RikishiID).To(Equal(8859))
		g.Expect(resp[0].ShikonaEnglish).To(Equal("Asasuiryu"))
		g.Expect(resp[0].ShikonaJapanese).To(Equal("朝翠龍　涼馬"))

		g.Expect(resp[9].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 594}))
		g.Expect(resp[9].BashoID).To(Equal(bashoID))
		g.Expect(resp[9].RikishiID).To(Equal(594))
		g.Expect(resp[9].ShikonaEnglish).To(Equal("Moriurara"))
		g.Expect(resp[9].ShikonaJapanese).To(Equal("森麗(もりうらら)"))
	})
}
