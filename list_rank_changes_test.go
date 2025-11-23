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

func TestClient_ListRankChanges(t *testing.T) {
	t.Run("list by rikishi ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[
			{
				"id": "202501-123",
				"bashoId": "202501",
				"rikishiId": 123,
				"rank": "Maegashira 1 East",
				"rankValue": 1
			}
		]`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/ranks"))
				g.Expect(req.URL.Query().Get("rikishiId")).To(Equal("123"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			RikishiID: 123,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(1))
		g.Expect(resp[0].RikishiID).To(Equal(123))
		g.Expect(resp[0].HumanReadableName).To(Equal("Maegashira 1 East"))
	})

	t.Run("list by basho ID", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[
			{
				"id": "202501-456",
				"bashoId": "202501",
				"rikishiId": 456
			},
			{
				"id": "202501-789",
				"bashoId": "202501",
				"rikishiId": 789
			}
		]`

		bashoID := sumoapi.BashoID{Year: 2025, Month: 1}
		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Path).To(Equal("/api/ranks"))
				g.Expect(req.URL.Query().Get("bashoId")).To(Equal("202501"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			BashoID: &bashoID,
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(2))
		g.Expect(resp[0].RikishiID).To(Equal(456))
		g.Expect(resp[1].RikishiID).To(Equal(789))
	})

	t.Run("list with sort order asc", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[]`

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
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			SortOrder: "asc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(BeEmpty())
	})

	t.Run("list with sort order desc", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[]`

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
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			SortOrder: "desc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(BeEmpty())
	})

	t.Run("list with multiple parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[
			{
				"id": "202501-123",
				"bashoId": "202501",
				"rikishiId": 123,
				"rank": "Ozeki 2 West",
				"rankValue": 4
			}
		]`

		bashoID := sumoapi.BashoID{Year: 2025, Month: 1}
		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(query.Get("rikishiId")).To(Equal("123"))
				g.Expect(query.Get("bashoId")).To(Equal("202501"))
				g.Expect(query.Get("sortOrder")).To(Equal("asc"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{
			RikishiID: 123,
			BashoID:   &bashoID,
			SortOrder: "asc",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(HaveLen(1))
		g.Expect(resp[0].RikishiID).To(Equal(123))
		g.Expect(resp[0].NumericName).To(Equal(4))
	})

	t.Run("empty request excludes optional parameters", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[]`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				query := req.URL.Query()
				g.Expect(query.Has("rikishiId")).To(BeFalse())
				g.Expect(query.Has("bashoId")).To(BeFalse())
				g.Expect(query.Has("sortOrder")).To(BeFalse())
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp).To(BeEmpty())
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `[]`

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
		_, err := client.ListRankChanges(ctx, sumoapi.ListRikishiChangesRequest{})

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
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{})

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
		resp, err := client.ListRankChanges(context.Background(), sumoapi.ListRikishiChangesRequest{})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
