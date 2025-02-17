package interfaces

import (
	"context"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/models/player"
	"github.com/1amDudman/faceit-wrapper/models/match"
)

type PlayerService interface {
	GetPlayerDetailsByID(ctx context.Context, playerID string) (*player.PlayerDetails, error)
	GetPlayerDetailsByNickname(ctx context.Context, nickname string) (*player.PlayerDetails, error)
}

type MatchService interface {
	GetMatchDetailsByID(ctx context.Context, matchID string) (*match.MatchDetails, error)
}

type HTTPRequester interface {
	NewRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
}
