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

type mockTransport struct {
	validateRequest func(*http.Request) error
	response        *http.Response
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.validateRequest != nil {
		if err := m.validateRequest(req); err != nil {
			return nil, err
		}
	}
	return m.response, nil
}

func TestClient_SearchRikishis(t *testing.T) {
	t.Run("search by shikona", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 100,
			"records": [
				{
					"id": 1,
					"shikonaEn": "Terunofuji"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/rikishis"))
				g.Expect(req.URL.Query().Get("shikonaEn")).To(Equal("Terunofuji"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			Shikona: "Terunofuji",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(100))
		g.Expect(resp.Rikishis).To(HaveLen(1))
		g.Expect(resp.Rikishis[0].ID).To(Equal(1))
		g.Expect(resp.Rikishis[0].ShikonaEnglish).To(Equal("Terunofuji"))
	})

	t.Run("search by heya", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 50,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Path).To(Equal("/api/rikishis"))
				g.Expect(req.URL.Query().Get("heya")).To(Equal("Isegahama"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			Heya: "Isegahama",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Total).To(Equal(50))
		g.Expect(resp.Rikishis).To(BeEmpty())
	})

	t.Run("search by SumoDB ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": [
				{
					"id": 999,
					"sumodbId": 11927
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("sumodbId")).To(Equal("11927"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			SumoDBID: 11927,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Rikishis).To(HaveLen(1))
		g.Expect(resp.Rikishis[0].SumoDBID).To(Equal(11927))
	})

	t.Run("search by official ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": [
				{
					"id": 888,
					"nskId": 3321
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("nskId")).To(Equal("3321"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			OfficialID: 3321,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Rikishis).To(HaveLen(1))
		g.Expect(resp.Rikishis[0].OfficialID).To(Equal(3321))
	})

	t.Run("search with include retired flag", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 200,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("intai")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			IncludeRetired: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("search with include measurements flag", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("measurements")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			IncludeMeasurements: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("search with include ranks flag", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("ranks")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			IncludeRanks: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("search with include shikonas flag", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 1,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("shikonas")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			IncludeShikonas: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("search with limit", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 10,
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
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			Limit: 10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Limit).To(Equal(10))
	})

	t.Run("search with skip", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 20,
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
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			Skip: 20,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Skip).To(Equal(20))
	})

	t.Run("search with multiple parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 5,
			"skip": 10,
			"total": 30,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(query.Get("shikonaEn")).To(Equal("Hakuho"))
				g.Expect(query.Get("heya")).To(Equal("Miyagino"))
				g.Expect(query.Get("intai")).To(Equal("true"))
				g.Expect(query.Get("measurements")).To(Equal("true"))
				g.Expect(query.Get("ranks")).To(Equal("true"))
				g.Expect(query.Get("shikonas")).To(Equal("true"))
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
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{
			Shikona:             "Hakuho",
			Heya:                "Miyagino",
			IncludeRetired:      true,
			IncludeMeasurements: true,
			IncludeRanks:        true,
			IncludeShikonas:     true,
			Limit:               5,
			Skip:                10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Limit).To(Equal(5))
		g.Expect(resp.Skip).To(Equal(10))
		g.Expect(resp.Total).To(Equal(30))
	})

	t.Run("empty request excludes optional parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 100,
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(query.Has("shikonaEn")).To(BeFalse())
				g.Expect(query.Has("heya")).To(BeFalse())
				g.Expect(query.Has("sumodbId")).To(BeFalse())
				g.Expect(query.Has("nskId")).To(BeFalse())
				g.Expect(query.Has("intai")).To(BeFalse())
				g.Expect(query.Has("measurements")).To(BeFalse())
				g.Expect(query.Has("ranks")).To(BeFalse())
				g.Expect(query.Has("shikonas")).To(BeFalse())
				g.Expect(query.Has("limit")).To(BeFalse())
				g.Expect(query.Has("skip")).To(BeFalse())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"total": 0,
			"records": []
		}`

		ctx := context.WithValue(context.Background(), "test-key", "test-value")

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Context().Value("test-key")).To(Equal("test-value"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		_, err := client.SearchRikishis(ctx, sumoapi.SearchRikishisRequest{})

		g.Expect(err).ToNot(HaveOccurred())
	})

	// Error path tests
	t.Run("http request error", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				return http.ErrAbortHandler
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error making http request"))
		g.Expect(resp).To(BeNil())
	})

	t.Run("HTTP 404 error with body", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader(`{"error": "not found"}`)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("received HTTP 404 response"))
		g.Expect(err.Error()).To(ContainSubstring(`{"error": "not found"}`))
		g.Expect(resp).To(BeNil())
	})

	t.Run("HTTP 500 error with empty body", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("")),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("received HTTP 500 response with empty body"))
		g.Expect(resp).To(BeNil())
	})

	t.Run("HTTP 400 bad request", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader("invalid parameters")),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("received HTTP 400 response"))
		g.Expect(err.Error()).To(ContainSubstring("invalid parameters"))
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
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})

	t.Run("malformed JSON response", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"limit": "not a number"}`)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})

	t.Run("HTTP 403 forbidden", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       io.NopCloser(strings.NewReader("access denied")),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("received HTTP 403 response"))
		g.Expect(err.Error()).To(ContainSubstring("access denied"))
		g.Expect(resp).To(BeNil())
	})

	t.Run("HTTP 503 service unavailable", func(t *testing.T) {
		g := NewWithT(t)

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusServiceUnavailable,
				Body:       io.NopCloser(strings.NewReader("service temporarily unavailable")),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.SearchRikishis(context.Background(), sumoapi.SearchRikishisRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("received HTTP 503 response"))
		g.Expect(err.Error()).To(ContainSubstring("service temporarily unavailable"))
		g.Expect(resp).To(BeNil())
	})
}
