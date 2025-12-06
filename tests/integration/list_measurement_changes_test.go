package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListMeasurementChanges(t *testing.T) {
	client := sumoapi.New()

	t.Run("for rikishi", func(t *testing.T) {
		g := NewWithT(t)

		resp, err := client.ListMeasurementChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			RikishiID: 3081, // Hakuho
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(15))

		expectedBashoID := sumoapi.BashoID{Year: 2021, Month: 3}

		g.Expect(resp[0].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
		g.Expect(resp[0].BashoID).To(Equal(expectedBashoID))
		g.Expect(resp[0].RikishiID).To(Equal(3081))
		g.Expect(resp[0].Height).To(Equal(192.0))
		g.Expect(resp[0].Weight).To(Equal(151.0))

		expectedBashoID = sumoapi.BashoID{Year: 2001, Month: 3}

		g.Expect(resp[14].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
		g.Expect(resp[14].BashoID).To(Equal(expectedBashoID))
		g.Expect(resp[14].RikishiID).To(Equal(3081))
		g.Expect(resp[14].Height).To(Equal(180.0))
		g.Expect(resp[14].Weight).To(Equal(80.0))
	})

	t.Run("for basho", func(t *testing.T) {
		g := NewWithT(t)

		bashoID := sumoapi.BashoID{
			Year:  2025,
			Month: 9,
		}

		resp, err := client.ListMeasurementChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			BashoID: &bashoID,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(4))

		g.Expect(resp[0].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 9098}))
		g.Expect(resp[0].BashoID).To(Equal(bashoID))
		g.Expect(resp[0].RikishiID).To(Equal(9098))
		g.Expect(resp[0].Height).To(Equal(178.0))
		g.Expect(resp[0].Weight).To(Equal(119.0))

		g.Expect(resp[3].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 9101}))
		g.Expect(resp[3].BashoID).To(Equal(bashoID))
		g.Expect(resp[3].RikishiID).To(Equal(9101))
		g.Expect(resp[3].Height).To(Equal(175.0))
		g.Expect(resp[3].Weight).To(Equal(117.0))
	})
}
