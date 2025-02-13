package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/errors"
	"github.com/1amDudman/faceit-wrapper/interfaces"
	"github.com/1amDudman/faceit-wrapper/models"
)

type MatchService struct {
	http interfaces.HTTPRequester
}

// NewMatchService is a MatchService constructor
func NewMatchService(http interfaces.HTTPRequester) *MatchService {
	return &MatchService{
		http: http,
	}
}

// Get data from a match by its ID
func (ms *MatchService) GetByMatchID(ctx context.Context, matchID string) (*models.Match, error) {
	if matchID == "" {
		return nil, errors.ErrMatchIDEmpty
	}

	endpoint := fmt.Sprintf("matches/%s", matchID)
	req, err := ms.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var match models.Match
	resp, err := ms.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}
