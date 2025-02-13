package client

import (
	"net/http"
	"time"

	"github.com/1amDudman/faceit-wrapper/errors"
	"github.com/1amDudman/faceit-wrapper/interfaces"
	"github.com/1amDudman/faceit-wrapper/services"
)

type FaceitClient struct {
	apiURL     string
	apiKey     string
	httpClient *http.Client

	PlayerService interfaces.PlayerService
	MatchService  interfaces.MatchService
}

// NewFaceitClient is a FaceitClient constructor
func NewFaceitClient(apiURL, apiKey string, opts ...Option) (*FaceitClient, error) {
	if apiURL == "" {
		return nil, errors.ErrApiUrlEmpty
	} else if apiKey == "" {
		return nil, errors.ErrApiKeyEmpty
	}

	fc := &FaceitClient{
		apiURL:     apiURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}

	fc.PlayerService = services.NewPlayerService(fc)
	fc.MatchService = services.NewMatchService(fc)

	for _, opt := range opts {
		opt(fc)
	}

	return fc, nil
}
