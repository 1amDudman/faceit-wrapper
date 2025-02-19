package interfaces

import (
	"context"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/models/match"
	"github.com/1amDudman/faceit-wrapper/models/player"
)

type PlayerService interface {
	GetPlayerDetailsByID(ctx context.Context, playerID string) (*player.PlayerDetails, error)
	GetPlayerDetailsByNickname(ctx context.Context, nickname string) (*player.PlayerDetails, error)
	GetPlayerStats(ctx context.Context, playerID, gameID string) (*player.PlayerStats, error)
	GetPlayerStatsByRange(ctx context.Context, playerID, gameID string, params map[string]string) (*player.PlayerRangeStats, error)
}

type MatchService interface {
	GetMatchDetailsByID(ctx context.Context, matchID string) (*match.MatchDetails, error)
	GetMatchStatsByID(ctx context.Context, matchID string) (*match.MatchStats, error)
}

type HTTPRequester interface {
	NewRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
	// MakeRequest(ctx context.Context, method, endpoint string, result any) error
}
