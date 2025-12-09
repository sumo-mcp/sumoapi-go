package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListRankChanges(t *testing.T) {
	client := sumoapi.New()

	t.Run("for rikishi", func(t *testing.T) {
		g := NewWithT(t)

		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			RikishiID: 3081, // Hakuho
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(122))

		expectedBashoID := sumoapi.BashoID{Year: 2021, Month: 9}

		g.Expect(resp[0].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
		g.Expect(resp[0].BashoID).To(Equal(expectedBashoID))
		g.Expect(resp[0].RikishiID).To(Equal(3081))
		g.Expect(resp[0].HumanReadableName).To(Equal("Yokozuna 1 East"))
		g.Expect(resp[0].NumericName).To(Equal(101))

		expectedBashoID = sumoapi.BashoID{Year: 2001, Month: 3}

		g.Expect(resp[121].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
		g.Expect(resp[121].BashoID).To(Equal(expectedBashoID))
		g.Expect(resp[121].RikishiID).To(Equal(3081))
		g.Expect(resp[121].HumanReadableName).To(Equal("Mae-zumo"))
		g.Expect(resp[121].NumericName).To(Equal(2000))
	})

	t.Run("for basho", func(t *testing.T) {
		g := NewWithT(t)

		bashoID := sumoapi.BashoID{
			Year:  2025,
			Month: 9,
		}

		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			BashoID: &bashoID,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(611))

		g.Expect(resp[0].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 8850}))
		g.Expect(resp[0].BashoID).To(Equal(bashoID))
		g.Expect(resp[0].RikishiID).To(Equal(8850))
		g.Expect(resp[0].HumanReadableName).To(Equal("Yokozuna 1 East"))
		g.Expect(resp[0].NumericName).To(Equal(101))

		g.Expect(resp[610].ID).To(Equal(sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 9101}))
		g.Expect(resp[610].BashoID).To(Equal(bashoID))
		g.Expect(resp[610].RikishiID).To(Equal(9101))
		g.Expect(resp[610].HumanReadableName).To(Equal("Jonokuchi 26 East"))
		g.Expect(resp[610].NumericName).To(Equal(1026))
	})
}
