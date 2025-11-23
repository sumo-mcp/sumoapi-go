package sumoapi_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestClient_GetRikishiStats(t *testing.T) {
	t.Run("get rikishi stats by ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"basho": 81,
			"yusho": 13,
			"totalMatches": 798,
			"totalWins": 523,
			"totalLosses": 275,
			"totalAbsences": 231,
			"sansho": {
				"Gino-sho": 3,
				"Kanto-sho": 3,
				"Shukun-sho": 3
			},
			"bashoByDivision": {
				"Makuuchi": 52,
				"Juryo": 7
			},
			"yushoByDivision": {
				"Makuuchi": 10,
				"Juryo": 2
			},
			"winsByDivision": {
				"Makuuchi": 366,
				"Juryo": 61
			},
			"lossByDivision": {
				"Makuuchi": 207,
				"Juryo": 38
			},
			"absenceByDivision": {
				"Makuuchi": 197,
				"Juryo": 6
			},
			"totalByDivision": {
				"Makuuchi": 573,
				"Juryo": 99
			}
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45/stats"))
				g.Expect(req.URL.RawQuery).To(BeEmpty())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishiStats(context.Background(), sumoapi.GetRikishiStatsRequest{
			RikishiID: 45,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Basho).To(Equal(81))
		g.Expect(resp.Yusho).To(Equal(13))
		g.Expect(resp.TotalMatches).To(Equal(798))
		g.Expect(resp.TotalWins).To(Equal(523))
		g.Expect(resp.TotalLosses).To(Equal(275))
		g.Expect(resp.TotalAbsences).To(Equal(231))

		// Check sansho
		g.Expect(resp.Sansho).To(HaveLen(3))
		g.Expect(resp.Sansho["Gino-sho"]).To(Equal(3))
		g.Expect(resp.Sansho["Kanto-sho"]).To(Equal(3))
		g.Expect(resp.Sansho["Shukun-sho"]).To(Equal(3))

		// Check division stats
		g.Expect(resp.BashoByDivision["Makuuchi"]).To(Equal(52))
		g.Expect(resp.YushoByDivision["Makuuchi"]).To(Equal(10))
		g.Expect(resp.WinsByDivision["Makuuchi"]).To(Equal(366))
		g.Expect(resp.LossByDivision["Makuuchi"]).To(Equal(207))
		g.Expect(resp.AbsenceByDivision["Makuuchi"]).To(Equal(197))
		g.Expect(resp.TotalMatchesByDivision["Makuuchi"]).To(Equal(573))
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"basho": 10,
			"yusho": 1
		}`

		type testKey struct{}
		ctx := context.WithValue(context.Background(), testKey{}, "test-value")

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Context().Value(testKey{})).To(Equal("test-value"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		_, err := client.GetRikishiStats(ctx, sumoapi.GetRikishiStatsRequest{
			RikishiID: 45,
		})

		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("http request error", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				return http.ErrAbortHandler
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishiStats(context.Background(), sumoapi.GetRikishiStatsRequest{
			RikishiID: 45,
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error making http request"))
		g.Expect(resp).To(BeNil())
	})

	t.Run("invalid JSON response", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("not valid json")),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishiStats(context.Background(), sumoapi.GetRikishiStatsRequest{
			RikishiID: 45,
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
