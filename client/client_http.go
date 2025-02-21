package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/errors"
)

// Request preparation
func (fc *FaceitClient) newRequest(ctx context.Context, method, endpoint string) (*http.Request, error) {
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

// Send request handling
func (fc *FaceitClient) do(req *http.Request) (*http.Response, error) {
	resp, err := fc.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrRequestFailed, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	return resp, nil
}

// Request process wrapper
func (fc *FaceitClient) MakeRequest(ctx context.Context, method, endpoint string, result any) error {
	req, err := fc.newRequest(ctx, method, endpoint)
	if err != nil {
		return err
	}

	resp, err := fc.do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
