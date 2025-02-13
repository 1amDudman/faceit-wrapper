package client

import (
	"time"

	"github.com/1amDudman/faceit-wrapper/interfaces"
)

type Option func(*FaceitClient)

// In case we need mock or a "user" defined timeout
func WithTimeout(timeout time.Duration) Option {
	return func(fc *FaceitClient) {
		fc.httpClient.Timeout = timeout
	}
}

// In case we need mock or a "user" defined player service
func WithCustomPlayerService(playerService interfaces.PlayerService) Option {
	return func(fc *FaceitClient) {
		fc.PlayerService = playerService
	}
}

// In case we need mock or a "user" defined match service
func WithCustomMatchService(matchService interfaces.MatchService) Option {
	return func(fc *FaceitClient) {
		fc.MatchService = matchService
	}
}
