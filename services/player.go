package services

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/1amDudman/faceit-wrapper/errors"
	"github.com/1amDudman/faceit-wrapper/interfaces"
	models "github.com/1amDudman/faceit-wrapper/models/player"
)

type PlayerService struct {
	http interfaces.HTTPRequester
}

// PlayerService constructor
func NewPlayerService(http interfaces.HTTPRequester) *PlayerService {
	return &PlayerService{
		http: http,
	}
}

// Get player data by his ID
func (ps *PlayerService) GetPlayerDetailsByID(ctx context.Context, playerID string) (*models.PlayerDetails, error) {
	if playerID == "" {
		return nil, errors.ErrPlayerIDEmpty
	}

	endpoint := fmt.Sprintf("players/%s", playerID)
	var playerDetails models.PlayerDetails
	err := ps.http.MakeRequest(ctx, http.MethodGet, endpoint, &playerDetails)
	if err != nil {
		return nil, err
	}

	return &playerDetails, nil
}

// Get player data by his nickname
func (ps *PlayerService) GetPlayerDetailsByNickname(ctx context.Context, nickname string) (*models.PlayerDetails, error) {
	if nickname == "" {
		return nil, errors.ErrPlayerNicknameEmpty
	}

	endpoint := fmt.Sprintf("players?nickname=%s", nickname)
	var playerDetails models.PlayerDetails
	err := ps.http.MakeRequest(ctx, http.MethodGet, endpoint, &playerDetails)
	if err != nil {
		return nil, err
	}

	return &playerDetails, nil
}

// Get total player stats(all time)
func (ps *PlayerService) GetPlayerStats(ctx context.Context, playerID, gameID string) (*models.PlayerStats, error) {
	if playerID == "" {
		return nil, errors.ErrPlayerIDEmpty
	}
	if gameID == "" {
		return nil, errors.ErrGameIDEmpty
	}

	endpoint := fmt.Sprintf("players/%s/stats/%s", playerID, gameID)
	var playerStats models.PlayerStats
	err := ps.http.MakeRequest(ctx, http.MethodGet, endpoint, &playerStats)
	if err != nil {
		return nil, err
	}

	return &playerStats, nil
}

// Get player stats by range of matches
func (ps *PlayerService) GetPlayerStatsByRange(ctx context.Context, playerID, gameID string, params map[string]string) (*models.PlayerRangeStats, error) {
	if playerID == "" {
		return nil, errors.ErrPlayerIDEmpty
	}
	if gameID == "" {
		return nil, errors.ErrGameIDEmpty
	}

	endpoint := fmt.Sprintf("players/%s/games/%s/stats", playerID, gameID)
	for k, v := range params {
		endpoint += fmt.Sprintf("?%s=%s&", k, v)
	}
	endpoint = strings.TrimSuffix(endpoint, "&")
	var playerRangeStats models.PlayerRangeStats
	err := ps.http.MakeRequest(ctx, http.MethodGet, endpoint, &playerRangeStats)
	if err != nil {
		return nil, err
	}

	return &playerRangeStats, nil
}
