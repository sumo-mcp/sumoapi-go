package sumoapi_test

import "net/http"

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
