package integration_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestIntegration_ListKimarite(t *testing.T) {
	g := NewWithT(t)

	client := sumoapi.New()

	resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
		SortField: "count",
		SortOrder: "desc",
		Limit:     1,
		Skip:      1,
	})

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(resp).ToNot(BeNil())

	g.Expect(resp.SortField).To(Equal("count"))
	g.Expect(resp.SortOrder).To(Equal("desc"))
	g.Expect(resp.Limit).To(Equal(1))
	g.Expect(resp.Skip).To(Equal(1))
	g.Expect(resp.Kimarite).To(HaveLen(1))

	k := resp.Kimarite[0]
	g.Expect(k.Count).To(BeNumerically(">", 0))
	g.Expect(k.LastUsage.BashoID).ToNot(BeZero())
	g.Expect(k.LastUsage.Day).To(BeNumerically(">", 0))
}
