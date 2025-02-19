package services

import (
	"context"
	"encoding/json"
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
	req, err := ps.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var player models.PlayerDetails
	resp, err := ps.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

// Get player data by his nickname
func (ps *PlayerService) GetPlayerDetailsByNickname(ctx context.Context, nickname string) (*models.PlayerDetails, error) {
	if nickname == "" {
		return nil, errors.ErrPlayerNicknameEmpty
	}

	endpoint := fmt.Sprintf("players?nickname=%s", nickname)
	req, err := ps.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var playerDetails models.PlayerDetails
	resp, err := ps.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&playerDetails)
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
	req, err := ps.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var playerStats models.PlayerStats
	resp, err := ps.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&playerStats)
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

	req, err := ps.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var playerRangeStats models.PlayerRangeStats
	resp, err := ps.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&playerRangeStats)
	if err != nil {
		return nil, err
	}

	return &playerRangeStats, nil
}
