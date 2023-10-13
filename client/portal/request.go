package portal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type RequestFactory struct {
}

func NewRequestFactory() RequestFactory {
	return RequestFactory{}
}

func (f RequestFactory) NewRequest(ctx context.Context) (*http.Request, error) {
	rawURL := "https://portal.czarnetermosy.pl/menu"
	URL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	return req, nil
}
