package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/errors"
)

// NewRequest ia a http Request wrapper
func (fc *FaceitClient) NewRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Request, error) {
	url := fc.apiURL + endpoint
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrBadRequest, err)
	}

	req.Header.Set("Authorization", "Bearer "+fc.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Do ia a http Do wrapper
func (fc *FaceitClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := fc.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrRequestFailed, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	return resp, nil
}
