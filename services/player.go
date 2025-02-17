package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/errors"
	"github.com/1amDudman/faceit-wrapper/interfaces"
	models "github.com/1amDudman/faceit-wrapper/models/player"
)

type PlayerService struct {
	http interfaces.HTTPRequester
}

// NewMatchService is a PlayerService constructor
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