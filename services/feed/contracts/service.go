package contracts

import (
	"context"

	"github.com/trillyai/backend-microservices/services/feed/shared"
)

type Service interface {
	GenerateFeed(ctx context.Context, offset, limit uint32, username string) (shared.Feed, error)
}
