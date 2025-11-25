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

func TestClient_GetBashoWithTorikumi(t *testing.T) {
	t.Run("get basho with torikumi", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501",
			"startDate": "2025-01-12T00:00:00Z",
			"endDate": "2025-01-26T00:00:00Z",
			"torikumi": [
				{
					"id": "202501-1-1-45-123",
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
					"kimarite": "yorikiri",
					"winnerId": 45,
					"winnerEn": "Terunofuji",
					"winnerJp": "照ノ富士"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/basho/202501/torikumi/Makuuchi/1"))
				g.Expect(req.URL.Query()).To(BeEmpty())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBashoWithTorikumi(context.Background(), sumoapi.GetBashoWithTorikumiRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 1},
			Division: "Makuuchi",
			Day:      1,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.ID).To(Equal(sumoapi.BashoID{Year: 2025, Month: 1}))
		g.Expect(resp.Torikumi).To(HaveLen(1))
		g.Expect(resp.Torikumi[0].ID).ToNot(BeNil())
		g.Expect(resp.Torikumi[0].ID.BashoID).To(Equal(sumoapi.BashoID{Year: 2025, Month: 1}))
		g.Expect(resp.Torikumi[0].ID.Day).To(Equal(1))
		g.Expect(resp.Torikumi[0].ID.MatchNumber).To(Equal(1))
		g.Expect(resp.Torikumi[0].BashoID).To(Equal(sumoapi.BashoID{Year: 2025, Month: 1}))
		g.Expect(resp.Torikumi[0].Division).To(Equal("Makuuchi"))
		g.Expect(resp.Torikumi[0].Day).To(Equal(1))
		g.Expect(resp.Torikumi[0].MatchNumber).To(Equal(1))
		g.Expect(resp.Torikumi[0].EastID).To(Equal(45))
		g.Expect(resp.Torikumi[0].WestID).To(Equal(123))
		g.Expect(resp.Torikumi[0].Kimarite).To(Equal("yorikiri"))
		g.Expect(resp.Torikumi[0].WinnerID).To(Equal(45))
	})

	t.Run("get basho with torikumi for Juryo division", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501",
			"torikumi": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/basho/202501/torikumi/Juryo/5"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBashoWithTorikumi(context.Background(), sumoapi.GetBashoWithTorikumiRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 1},
			Division: "Juryo",
			Day:      5,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Torikumi).To(BeEmpty())
	})

	t.Run("get basho with torikumi for playoff day", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501",
			"torikumi": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/basho/202501/torikumi/Makuuchi/16"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBashoWithTorikumi(context.Background(), sumoapi.GetBashoWithTorikumiRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 1},
			Division: "Makuuchi",
			Day:      16,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"date": "202501",
			"torikumi": []
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
		_, err := client.GetBashoWithTorikumi(ctx, sumoapi.GetBashoWithTorikumiRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 1},
			Division: "Makuuchi",
			Day:      1,
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
		resp, err := client.GetBashoWithTorikumi(context.Background(), sumoapi.GetBashoWithTorikumiRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 1},
			Division: "Makuuchi",
			Day:      1,
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
		resp, err := client.GetBashoWithTorikumi(context.Background(), sumoapi.GetBashoWithTorikumiRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 1},
			Division: "Makuuchi",
			Day:      1,
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
