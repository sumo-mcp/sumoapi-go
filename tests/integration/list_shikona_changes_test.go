package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListShikonaChanges(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	// Here we test a specific Rikishi that had shikona changes to make sure the API
	// is returning exactly one change when filtering by RikishiID and BashoID.
	resp, err := client.ListShikonaChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
		RikishiID: 8857,
		BashoID: &sumoapi.BashoID{
			Year:  2025,
			Month: 11,
		},
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp).To(HaveLen(1))

	expectedBashoID := sumoapi.BashoID{Year: 2025, Month: 11}

	yoshinofuji := resp[0]
	g.Expect(yoshinofuji.ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 8857}))
	g.Expect(yoshinofuji.BashoID).To(Equal(expectedBashoID))
	g.Expect(yoshinofuji.RikishiID).To(Equal(8857))
	g.Expect(yoshinofuji.ShikonaEnglish).To(Equal("Yoshinofuji"))
	g.Expect(yoshinofuji.ShikonaJapanese).To(Equal("義ノ富士　直哉"))
}
