package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListRankChanges(t *testing.T) {
	client := sumoapi.New()

	// Use a non-default limit to exercise the Limit input as well.
	const pageLimit = 99

	listAllRanks := func(g *WithT, req sumoapi.ListRikishiChangesRequest) []sumoapi.Rank {
		req.Limit = pageLimit
		var all []sumoapi.Rank
		for {
			req.Skip = len(all)
			page, err := client.ListRankChanges(context.Background(), req)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(page).ToNot(BeNil())
			all = append(all, page...)
			if len(page) < pageLimit {
				return all
			}
		}
	}

	t.Run("for rikishi", func(t *testing.T) {
		g := NewWithT(t)

		resp := listAllRanks(g, sumoapi.ListRikishiChangesRequest{
			RikishiID: 3081, // Hakuho
		})

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

		resp := listAllRanks(g, sumoapi.ListRikishiChangesRequest{
			BashoID: &bashoID,
		})

		g.Expect(resp).To(HaveLen(611))

		g.Expect(resp[0]).To(Equal(sumoapi.Rank{
			ID:                sumoapi.RikishiChangeID{BashoID: bashoID, RikishiID: 19},
			BashoID:           bashoID,
			RikishiID:         19,
			HumanReadableName: "Yokozuna 1 West",
			NumericName:       101,
		}))
	})
}
