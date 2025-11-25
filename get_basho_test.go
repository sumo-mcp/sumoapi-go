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

func TestClient_GetBasho(t *testing.T) {
	t.Run("get basho by ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501",
			"startDate": "2025-01-12T00:00:00Z",
			"endDate": "2025-01-26T00:00:00Z",
			"yusho": [
				{
					"type": "Makuuchi",
					"rikishiId": 45,
					"shikonaEn": "Terunofuji",
					"shikonaJp": "照ノ富士"
				}
			],
			"specialPrizes": [
				{
					"type": "Shukun-sho",
					"rikishiId": 123,
					"shikonaEn": "Takakeisho",
					"shikonaJp": "貴景勝"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/basho/202501"))
				g.Expect(req.URL.Query()).To(BeEmpty())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBasho(context.Background(), sumoapi.GetBashoRequest{
			BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.ID).To(Equal(sumoapi.BashoID{Year: 2025, Month: 1}))
		g.Expect(resp.StartDate).ToNot(BeNil())
		g.Expect(resp.EndDate).ToNot(BeNil())
		g.Expect(resp.Yusho).To(HaveLen(1))
		g.Expect(resp.Yusho[0].Type).To(Equal("Makuuchi"))
		g.Expect(resp.Yusho[0].RikishiID).To(Equal(45))
		g.Expect(resp.Yusho[0].ShikonaEnglish).To(Equal("Terunofuji"))
		g.Expect(resp.SpecialPrizes).To(HaveLen(1))
		g.Expect(resp.SpecialPrizes[0].Type).To(Equal("Shukun-sho"))
		g.Expect(resp.SpecialPrizes[0].RikishiID).To(Equal(123))
	})

	t.Run("get basho with minimal data", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "199903"
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/basho/199903"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBasho(context.Background(), sumoapi.GetBashoRequest{
			BashoID: sumoapi.BashoID{Year: 1999, Month: 3},
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.ID).To(Equal(sumoapi.BashoID{Year: 1999, Month: 3}))
		g.Expect(resp.StartDate).To(BeNil())
		g.Expect(resp.EndDate).To(BeNil())
		g.Expect(resp.Yusho).To(BeEmpty())
		g.Expect(resp.SpecialPrizes).To(BeEmpty())
		g.Expect(resp.Torikumi).To(BeEmpty())
	})

	t.Run("get basho with torikumi", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501",
			"torikumi": [
				{
					"bashoId": "202501",
					"division": "Makuuchi",
					"day": 1,
					"matchNo": 1,
					"eastId": 45,
					"eastShikona": "Terunofuji",
					"eastRank": "Yokozuna 1 East",
					"westId": 123,
					"westShikona": "Takakeisho",
					"westRank": "Ozeki 1 West",
					"kimarite": "Yorikiri",
					"winnerId": 45,
					"winnerEn": "Terunofuji"
				}
			]
		}`

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBasho(context.Background(), sumoapi.GetBashoRequest{
			BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Torikumi).To(HaveLen(1))
		g.Expect(resp.Torikumi[0].BashoID).To(Equal(sumoapi.BashoID{Year: 2025, Month: 1}))
		g.Expect(resp.Torikumi[0].Division).To(Equal("Makuuchi"))
		g.Expect(resp.Torikumi[0].Day).To(Equal(1))
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501"
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
		_, err := client.GetBasho(ctx, sumoapi.GetBashoRequest{
			BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
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
		resp, err := client.GetBasho(context.Background(), sumoapi.GetBashoRequest{
			BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
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
		resp, err := client.GetBasho(context.Background(), sumoapi.GetBashoRequest{
			BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
