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

func TestClient_ListKimarite(t *testing.T) {
	t.Run("list by kimarite sort field", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"sortField": "kimarite",
			"sortOrder": "asc",
			"records": [
				{
					"kimarite": "hatakikomi",
					"count": 1234,
					"lastUsage": "202501-5"
				},
				{
					"kimarite": "oshidashi",
					"count": 5678,
					"lastUsage": "202501-10"
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/kimarite"))
				g.Expect(req.URL.Query().Get("sortField")).To(Equal("kimarite"))
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
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "kimarite",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.SortField).To(Equal("kimarite"))
		g.Expect(resp.SortOrder).To(Equal("asc"))
		g.Expect(resp.Kimarite).To(HaveLen(2))
		g.Expect(resp.Kimarite[0].Name).To(Equal("hatakikomi"))
		g.Expect(resp.Kimarite[0].Count).To(Equal(1234))
		g.Expect(resp.Kimarite[1].Name).To(Equal("oshidashi"))
		g.Expect(resp.Kimarite[1].Count).To(Equal(5678))
	})

	t.Run("list by count sort field", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"sortField": "count",
			"sortOrder": "asc",
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("sortField")).To(Equal("count"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "count",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.SortField).To(Equal("count"))
	})

	t.Run("list by lastUsage sort field", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"sortField": "lastUsage",
			"sortOrder": "asc",
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("sortField")).To(Equal("lastUsage"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "lastUsage",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.SortField).To(Equal("lastUsage"))
	})

	t.Run("list with sort order asc", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"sortField": "count",
			"sortOrder": "asc",
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("sortOrder")).To(Equal("asc"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "count",
			SortOrder: "asc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.SortOrder).To(Equal("asc"))
	})

	t.Run("list with sort order desc", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"sortField": "count",
			"sortOrder": "desc",
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("sortOrder")).To(Equal("desc"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "count",
			SortOrder: "desc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.SortOrder).To(Equal("desc"))
	})

	t.Run("list with limit", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 10,
			"skip": 0,
			"sortField": "kimarite",
			"sortOrder": "asc",
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
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "kimarite",
			Limit:     10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Limit).To(Equal(10))
	})

	t.Run("list with skip", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 5,
			"sortField": "kimarite",
			"sortOrder": "asc",
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Query().Get("skip")).To(Equal("5"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "kimarite",
			Skip:      5,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.Skip).To(Equal(5))
	})

	t.Run("list with all parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 20,
			"skip": 10,
			"sortField": "count",
			"sortOrder": "desc",
			"records": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(query.Get("sortField")).To(Equal("count"))
				g.Expect(query.Get("sortOrder")).To(Equal("desc"))
				g.Expect(query.Get("limit")).To(Equal("20"))
				g.Expect(query.Get("skip")).To(Equal("10"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "count",
			SortOrder: "desc",
			Limit:     20,
			Skip:      10,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.SortField).To(Equal("count"))
		g.Expect(resp.SortOrder).To(Equal("desc"))
		g.Expect(resp.Limit).To(Equal(20))
		g.Expect(resp.Skip).To(Equal(10))
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"limit": 0,
			"skip": 0,
			"sortField": "kimarite",
			"sortOrder": "asc",
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
		_, err := client.ListKimarite(ctx, sumoapi.ListKimariteRequest{
			SortField: "kimarite",
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
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "kimarite",
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
		resp, err := client.ListKimarite(context.Background(), sumoapi.ListKimariteRequest{
			SortField: "kimarite",
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
