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

type PlayerService struct {
	http interfaces.HTTPRequester
}

// NewMatchService is a PlayerService constructor
func NewPlayerService(http interfaces.HTTPRequester) *PlayerService {
	return &PlayerService{
		http: http,
	}
}

// Get data from a match by its ID
func (ps *PlayerService) GetByPlayerID(ctx context.Context, playerID string) (*models.Player, error) {
	if playerID == "" {
		return nil, errors.ErrPlayerIDEmpty
	}

	endpoint := fmt.Sprintf("players/%s", playerID)
	req, err := ps.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var player models.Player
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

func (ps *PlayerService) GetByPlayerNickname(ctx context.Context, nickname string) (*models.Player, error) {
	if nickname == "" {
		return nil, errors.ErrPlayerNicknameEmpty
	}

	endpoint := fmt.Sprintf("players?nickname=%s", nickname)
	req, err := ps.http.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var player models.Player
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
