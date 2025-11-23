package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListMeasurementChanges(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.ListMeasurementChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
		RikishiID: 3081, // Hakuho
		BashoID: &sumoapi.BashoID{
			Year:  2021,
			Month: 3,
		},
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(len(resp)).To(BeNumerically(">", 1)) // Bug: The bashoId filter is not working.

	expectedBashoID := sumoapi.BashoID{Year: 2021, Month: 3}

	hakuho := resp[0]
	g.Expect(hakuho.ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
	g.Expect(hakuho.BashoID).To(Equal(expectedBashoID))
	g.Expect(hakuho.RikishiID).To(Equal(3081))
	g.Expect(hakuho.Height).To(Equal(192.0))
	g.Expect(hakuho.Weight).To(Equal(151.0))
}
