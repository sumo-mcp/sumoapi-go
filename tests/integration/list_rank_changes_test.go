package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListRankChanges(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
		RikishiID: 3081,
		BashoID: &sumoapi.BashoID{
			Year:  2021,
			Month: 9,
		},
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())
	g.Expect(resp).To(HaveLen(1))

	expectedBashoID := sumoapi.BashoID{Year: 2021, Month: 9}

	hakuho := resp[0]
	g.Expect(hakuho.ID).To(Equal(sumoapi.RikishiChangeID{BashoID: expectedBashoID, RikishiID: 3081}))
	g.Expect(hakuho.BashoID).To(Equal(expectedBashoID))
	g.Expect(hakuho.RikishiID).To(Equal(3081))
	g.Expect(hakuho.HumanReadableName).To(Equal("Yokozuna 1 East"))
	g.Expect(hakuho.NumericName).To(Equal(101))
}
