package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/errors"
	"github.com/1amDudman/faceit-wrapper/interfaces"
	models "github.com/1amDudman/faceit-wrapper/models/match"
)

type MatchService struct {
	http interfaces.HTTPRequester
}

// MatchService constructor
func NewMatchService(http interfaces.HTTPRequester) *MatchService {
	return &MatchService{
		http: http,
	}
}

// Get match details(without deep match stats) by ID
func (ms *MatchService) GetMatchDetailsByID(ctx context.Context, matchID string) (*models.MatchDetails, error) {
	if matchID == "" {
		return nil, errors.ErrMatchIDEmpty
	}

	endpoint := fmt.Sprintf("matches/%s", matchID)
	req, err := ms.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var matchDetails models.MatchDetails
	resp, err := ms.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&matchDetails)
	if err != nil {
		return nil, err
	}

	return &matchDetails, nil
}

// Get players' stats in a match by its ID
func (ms *MatchService) GetMatchStatsByID(ctx context.Context, matchID string) (*models.MatchStats, error) {
	if matchID == "" {
		return nil, errors.ErrMatchIDEmpty
	}

	endpoint := fmt.Sprintf("matches/%s/stats", matchID)
	req, err := ms.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var matchStats models.MatchStats
	resp, err := ms.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&matchStats)
	if err != nil {
		return nil, err
	}

	return &matchStats, nil
}
