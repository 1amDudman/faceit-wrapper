package interfaces

import (
	"context"
	"net/http"

	"github.com/1amDudman/faceit-wrapper/models"
)

type PlayerService interface {
	GetByPlayerID(ctx context.Context, playerID string) (*models.Player, error)
}

type MatchService interface {
	GetByMatchID(ctx context.Context, matchID string) (*models.Match, error)
}

type HTTPRequester interface {
	NewRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Request, error)
	Do(req *http.Request) (*http.Response, error)
}
