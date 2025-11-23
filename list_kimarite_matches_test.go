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

func TestClient_ListKimariteMatches(t *testing.T) {
	t.Run("list matches by kimarite", func(t *testing.T) {
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
					"day": 3,
					"matchNo": 8,
					"eastId": 100,
					"eastShikona": "Wrestler A",
					"eastRank": "Maegashira 1 East",
					"westId": 200,
					"westShikona": "Wrestler B",
					"westRank": "Maegashira 1 West",
					"winnerId": 100,
					"winnerEn": "Wrestler A",
					"kimarite": "Yorikiri"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/kimarite/Yorikiri"))
				g.Expect(req.URL.Query().Has("sortOrder")).To(BeFalse())
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
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite: "Yorikiri",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(2))
		g.Expect(resp.Matches).To(HaveLen(2))
		g.Expect(resp.Matches[0].Kimarite).To(Equal("Yorikiri"))
		g.Expect(resp.Matches[1].Kimarite).To(Equal("Yorikiri"))
	})

	t.Run("list matches with sort order asc", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/kimarite/Oshidashi"))
				g.Expect(req.URL.Query().Get("sortOrder")).To(Equal("asc"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite:  "Oshidashi",
			SortOrder: "asc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(1))
	})

	t.Run("list matches with sort order desc", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/kimarite/Hatakikomi"))
				g.Expect(req.URL.Query().Get("sortOrder")).To(Equal("desc"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite:  "Hatakikomi",
			SortOrder: "desc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(1))
	})

	t.Run("list matches with limit", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 5,
			"skip": 0,
			"total": 100,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("limit")).To(Equal("5"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite: "Yorikiri",
			Limit:    5,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(100))
		g.Expect(resp.Limit).To(Equal(5))
	})

	t.Run("list matches with skip", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 10,
			"total": 100,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("skip")).To(Equal("10"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite: "Yorikiri",
			Skip:     10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(100))
		g.Expect(resp.Skip).To(Equal(10))
	})

	t.Run("list matches with all parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 20,
			"skip": 5,
			"total": 50,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(req.URL.Path).To(Equal("/api/kimarite/Tsukiotoshi"))
				g.Expect(query.Get("sortOrder")).To(Equal("desc"))
				g.Expect(query.Get("limit")).To(Equal("20"))
				g.Expect(query.Get("skip")).To(Equal("5"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite:  "Tsukiotoshi",
			SortOrder: "desc",
			Limit:     20,
			Skip:      5,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(50))
		g.Expect(resp.Limit).To(Equal(20))
		g.Expect(resp.Skip).To(Equal(5))
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
		_, err := client.ListKimariteMatches(ctx, sumoapi.ListKimariteMatchesRequest{
			Kimarite: "Yorikiri",
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
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite: "Yorikiri",
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
		resp, err := client.ListKimariteMatches(context.Background(), sumoapi.ListKimariteMatchesRequest{
			Kimarite: "Yorikiri",
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
