package services

import (
	"context"
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
	var matchDetails models.MatchDetails
	err := ms.http.MakeRequest(ctx, http.MethodGet, endpoint, &matchDetails)
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
	var matchStats models.MatchStats
	err := ms.http.MakeRequest(ctx, http.MethodGet, endpoint, &matchStats)
	if err != nil {
		return nil, err
	}

	return &matchStats, nil
}
