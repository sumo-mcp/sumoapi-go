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

func TestClient_GetBanzuke(t *testing.T) {
	t.Run("get banzuke for division", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"bashoId": "202511",
			"division": "Makuuchi",
			"east": [
				{
					"side": "East",
					"rikishiID": 8850,
					"shikonaEn": "Onosato",
					"shikonaJp": "大の里",
					"rankValue": 101,
					"rank": "Yokozuna 1 East",
					"record": [
						{
							"result": "win",
							"opponentShikonaEn": "Takayasu",
							"opponentShikonaJp": "高安",
							"opponentID": 44,
							"kimarite": "yorikiri"
						}
					],
					"wins": 11,
					"losses": 4,
					"absences": 0
				}
			],
			"west": [
				{
					"side": "West",
					"rikishiID": 19,
					"shikonaEn": "Hoshoryu",
					"shikonaJp": "豊昇龍",
					"rankValue": 201,
					"rank": "Ozeki 1 West",
					"record": [],
					"wins": 12,
					"losses": 3,
					"absences": 0
				}
			]
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.Method).To(Equal(http.MethodGet))
				g.Expect(req.URL.Scheme).To(Equal("https"))
				g.Expect(req.URL.Host).To(Equal("sumo-api.com"))
				g.Expect(req.URL.Path).To(Equal("/api/basho/202511/banzuke/Makuuchi"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		bashoID := sumoapi.BashoID{Year: 2025, Month: 11}
		resp, err := client.GetBanzuke(context.Background(), sumoapi.GetBanzukeRequest{
			BashoID:  bashoID,
			Division: "Makuuchi",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.BashoID).To(Equal(bashoID))
		g.Expect(resp.Division).To(Equal("Makuuchi"))

		// Check east side
		g.Expect(resp.East).To(HaveLen(1))
		g.Expect(resp.East[0].Side).To(Equal("East"))
		g.Expect(resp.East[0].RikishiID).To(Equal(8850))
		g.Expect(resp.East[0].ShikonaEnglish).To(Equal("Onosato"))
		g.Expect(resp.East[0].ShikonaJapanese).To(Equal("大の里"))
		g.Expect(resp.East[0].HumanReadableRankName).To(Equal("Yokozuna 1 East"))
		g.Expect(resp.East[0].NumericRankName).To(Equal(101))
		g.Expect(resp.East[0].Wins).To(Equal(11))
		g.Expect(resp.East[0].Losses).To(Equal(4))
		g.Expect(resp.East[0].Absences).To(Equal(0))

		// Check match record
		g.Expect(resp.East[0].Matches).To(HaveLen(1))
		g.Expect(resp.East[0].Matches[0].Result).To(Equal("win"))
		g.Expect(resp.East[0].Matches[0].OpponentShikonaEnglish).To(Equal("Takayasu"))
		g.Expect(resp.East[0].Matches[0].OpponentShikonaJapanese).To(Equal("高安"))
		g.Expect(resp.East[0].Matches[0].OpponentID).To(Equal(44))
		g.Expect(resp.East[0].Matches[0].Kimarite).To(Equal("yorikiri"))

		// Check west side
		g.Expect(resp.West).To(HaveLen(1))
		g.Expect(resp.West[0].Side).To(Equal("West"))
		g.Expect(resp.West[0].RikishiID).To(Equal(19))
		g.Expect(resp.West[0].ShikonaEnglish).To(Equal("Hoshoryu"))
	})

	t.Run("get banzuke for different division", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"bashoId": "202501",
			"division": "Juryo",
			"east": [],
			"west": []
		}`

		transport := &mockTransport{
			validateRequest: func(req *http.Request) error {
				g.Expect(req.URL.Path).To(Equal("/api/basho/202501/banzuke/Juryo"))
				return nil
			},
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		bashoID := sumoapi.BashoID{Year: 2025, Month: 1}
		resp, err := client.GetBanzuke(context.Background(), sumoapi.GetBanzukeRequest{
			BashoID:  bashoID,
			Division: "Juryo",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.BashoID).To(Equal(bashoID))
		g.Expect(resp.Division).To(Equal("Juryo"))
	})

	t.Run("banzuke with all result types", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"bashoId": "202511",
			"division": "Makuuchi",
			"east": [
				{
					"side": "East",
					"rikishiID": 1,
					"shikonaEn": "TestRikishi",
					"shikonaJp": "テスト力士",
					"rankValue": 101,
					"rank": "Yokozuna 1 East",
					"record": [
						{"result": "win", "opponentShikonaEn": "A", "opponentShikonaJp": "あ", "opponentID": 2, "kimarite": "yorikiri"},
						{"result": "loss", "opponentShikonaEn": "B", "opponentShikonaJp": "い", "opponentID": 3, "kimarite": "oshidashi"},
						{"result": "absent", "opponentShikonaEn": "C", "opponentShikonaJp": "う", "opponentID": 4},
						{"result": "fusen win", "opponentShikonaEn": "D", "opponentShikonaJp": "え", "opponentID": 5, "kimarite": "fusen"},
						{"result": "fusen loss", "opponentShikonaEn": "E", "opponentShikonaJp": "お", "opponentID": 6, "kimarite": "fusen"}
					],
					"wins": 2,
					"losses": 2,
					"absences": 1
				}
			],
			"west": []
		}`

		transport := &mockTransport{
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(mockResp)),
			},
		}

		client := sumoapi.New(sumoapi.WithHTTPClient(&http.Client{Transport: transport}))
		resp, err := client.GetBanzuke(context.Background(), sumoapi.GetBanzukeRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 11},
			Division: "Makuuchi",
		})

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(resp).ToNot(BeNil())
		g.Expect(resp.East[0].Matches).To(HaveLen(5))
		g.Expect(resp.East[0].Matches[0].Result).To(Equal("win"))
		g.Expect(resp.East[0].Matches[1].Result).To(Equal("loss"))
		g.Expect(resp.East[0].Matches[2].Result).To(Equal("absent"))
		g.Expect(resp.East[0].Matches[3].Result).To(Equal("fusen win"))
		g.Expect(resp.East[0].Matches[4].Result).To(Equal("fusen loss"))
	})

	t.Run("context is propagated", func(t *testing.T) {
		g := NewWithT(t)

		mockResp := `{
			"bashoId": "202511",
			"division": "Makuuchi",
			"east": [],
			"west": []
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
		_, err := client.GetBanzuke(ctx, sumoapi.GetBanzukeRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 11},
			Division: "Makuuchi",
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
		resp, err := client.GetBanzuke(context.Background(), sumoapi.GetBanzukeRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 11},
			Division: "Makuuchi",
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
		resp, err := client.GetBanzuke(context.Background(), sumoapi.GetBanzukeRequest{
			BashoID:  sumoapi.BashoID{Year: 2025, Month: 11},
			Division: "Makuuchi",
		})

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("error unmarshaling response body"))
		g.Expect(resp).To(BeNil())
	})
}
