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

func TestClient_GetRikishi(t *testing.T) {
	t.Run("get rikishi by ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"id": 45,
			"shikonaEn": "Terunofuji",
			"shikonaJp": "照ノ富士",
			"currentRank": "Yokozuna",
			"heya": "Isegahama"
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45"))
				g.Expect(req.URL.Query().Has("measurements")).To(BeFalse())
				g.Expect(req.URL.Query().Has("ranks")).To(BeFalse())
				g.Expect(req.URL.Query().Has("shikonas")).To(BeFalse())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
			RikishiID: 45,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.ID).To(Equal(45))
		g.Expect(resp.ShikonaEnglish).To(Equal("Terunofuji"))
		g.Expect(resp.ShikonaJapanese).To(Equal("照ノ富士"))
		g.Expect(resp.CurrentRank).To(Equal("Yokozuna"))
		g.Expect(resp.Heya).To(Equal("Isegahama"))
	})

	t.Run("get rikishi with include ranks", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"id": 45,
			"shikonaEn": "Terunofuji",
			"rankHistory": [
				{
					"id": "202501-45",
					"bashoId": "202501",
					"rikishiId": 45,
					"rank": "Yokozuna 1 East"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45"))
				g.Expect(req.URL.Query().Get("ranks")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
			RikishiID:    45,
			IncludeRanks: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.RankHistory).To(HaveLen(1))
		g.Expect(resp.RankHistory[0].RikishiID).To(Equal(45))
	})

	t.Run("get rikishi with include shikonas", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"id": 45,
			"shikonaEn": "Terunofuji",
			"shikonaHistory": [
				{
					"id": "202501-45",
					"bashoId": "202501",
					"rikishiId": 45,
					"shikonaEn": "Terunofuji",
					"shikonaJp": "照ノ富士"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45"))
				g.Expect(req.URL.Query().Get("shikonas")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
			RikishiID:       45,
			IncludeShikonas: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.ShikonaHistory).To(HaveLen(1))
		g.Expect(resp.ShikonaHistory[0].RikishiID).To(Equal(45))
	})

	t.Run("get rikishi with include measurements", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"id": 45,
			"shikonaEn": "Terunofuji",
			"measurementHistory": [
				{
					"id": "202501-45",
					"bashoId": "202501",
					"rikishiId": 45,
					"height": 193.0,
					"weight": 180.0
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/rikishi/45"))
				g.Expect(req.URL.Query().Get("measurements")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
			RikishiID:           45,
			IncludeMeasurements: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.MeasurementHistory).To(HaveLen(1))
		g.Expect(resp.MeasurementHistory[0].RikishiID).To(Equal(45))
	})

	t.Run("get rikishi with all includes", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"id": 45,
			"shikonaEn": "Terunofuji",
			"rankHistory": [],
			"shikonaHistory": [],
			"measurementHistory": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(query.Get("ranks")).To(Equal("true"))
				g.Expect(query.Get("shikonas")).To(Equal("true"))
				g.Expect(query.Get("measurements")).To(Equal("true"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
			RikishiID:           45,
			IncludeRanks:        true,
			IncludeShikonas:     true,
			IncludeMeasurements: true,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"id": 45,
			"shikonaEn": "Terunofuji"
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
		_, err := client.GetRikishi(ctx, sumoapi.GetRikishiRequest{
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
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
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
		resp, err := client.GetRikishi(context.Background(), sumoapi.GetRikishiRequest{
			RikishiID: 45,
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
