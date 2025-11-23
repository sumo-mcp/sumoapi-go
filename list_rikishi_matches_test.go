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

func TestClient_ListRikishiMatches(t *testing.T) {
	t.Run("list matches by rikishi ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 2,
			"records": [
				{
					"bashoId": "202501",
					"division": "Makuuchi",
					"day": 1,
					"matchNo": 5,
					"eastId": 45,
					"eastShikona": "Terunofuji",
					"eastRank": "Yokozuna 1 East",
					"westId": 123,
					"westShikona": "Takakeisho",
					"westRank": "Ozeki 1 West",
					"winnerId": 45,
					"winnerEn": "Terunofuji",
					"kimarite": "Yorikiri"
				},
				{
					"bashoId": "202501",
					"division": "Makuuchi",
					"day": 2,
					"matchNo": 6,
					"eastId": 45,
					"eastShikona": "Terunofuji",
					"eastRank": "Yokozuna 1 East",
					"westId": 456,
					"westShikona": "Hoshoryu",
					"westRank": "Sekiwake 1 West",
					"winnerId": 45,
					"winnerEn": "Terunofuji",
					"kimarite": "Oshidashi"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45/matches"))
				g.Expect(req.URL.Query().Has("bashoId")).To(BeFalse())
				g.Expect(req.URL.Query().Has("limit")).To(BeFalse())
				g.Expect(req.URL.Query().Has("skip")).To(BeFalse())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
			RikishiID: 45,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(2))
		g.Expect(resp.Matches).To(HaveLen(2))
		g.Expect(resp.Matches[0].EastID).To(Equal(45))
		g.Expect(resp.Matches[0].EastShikona).To(Equal("Terunofuji"))
		g.Expect(resp.Matches[1].WinnerID).To(Equal(45))
	})

	t.Run("list matches with basho ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": [
				{
					"bashoId": "202501",
					"division": "Makuuchi",
					"day": 1,
					"eastId": 45,
					"winnerId": 45
				}
			]
		}`

		bashoID := sumoapi.BashoID{Year: 2025, Month: 1}
		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45/matches"))
				g.Expect(req.URL.Query().Get("bashoId")).To(Equal("202501"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
			RikishiID: 45,
			BashoID:   &bashoID,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(1))
		g.Expect(resp.Matches).To(HaveLen(1))
	})

	t.Run("list matches with limit", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 100,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("limit")).To(Equal("10"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
			RikishiID: 45,
			Limit:     10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(100))
		g.Expect(resp.Matches).To(BeEmpty())
	})

	t.Run("list matches with skip", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 100,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("skip")).To(Equal("20"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
			RikishiID: 45,
			Skip:      20,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(100))
		g.Expect(resp.Matches).To(BeEmpty())
	})

	t.Run("list matches with all parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 5,
			"records": []
		}`

		bashoID := sumoapi.BashoID{Year: 2025, Month: 1}
		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45/matches"))
				g.Expect(query.Get("bashoId")).To(Equal("202501"))
				g.Expect(query.Get("limit")).To(Equal("5"))
				g.Expect(query.Get("skip")).To(Equal("10"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
			RikishiID: 45,
			BashoID:   &bashoID,
			Limit:     5,
			Skip:      10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(5))
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 0,
			"records": []
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
		_, err := client.ListRikishiMatches(ctx, sumoapi.ListRikishiMatchesRequest{
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
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
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
		resp, err := client.ListRikishiMatches(context.Background(), sumoapi.ListRikishiMatchesRequest{
			RikishiID: 45,
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
