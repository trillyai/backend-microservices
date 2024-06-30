package repository

import (
	"context"

	"github.com/trillyai/backend-microservices/core/database/postgres"
	"github.com/trillyai/backend-microservices/core/database/tables"
	"github.com/trillyai/backend-microservices/core/logger"
	"github.com/trillyai/backend-microservices/core/utils"
	"github.com/trillyai/backend-microservices/services/feed/contracts"
	"github.com/trillyai/backend-microservices/services/feed/shared"
)

type repository struct {
	logger logger.Logger
}

func NewRepository() contracts.Repository {
	return repository{
		logger: *logger.NewLogger("feed-repository"),
	}
}

// GenerateFeed implements contracts.Repository.
func (r repository) GenerateFeed(ctx context.Context, offset uint32, limit uint32, username string) (shared.Feed, error) {
	resp, err := getRandomFeed(ctx, offset, limit)
	if err != nil {
		r.logger.Error(err.Error())
		return shared.Feed{}, err
	}
	return resp, nil
}

func getRandomFeed(ctx context.Context, offset, limit uint32) (shared.Feed, error) {
	var resp shared.Feed
	posts, err := postgres.PaginatedRead[[]tables.Post, tables.Post](ctx, map[string]interface{}{}, offset, limit)
	if err != nil {
		return shared.Feed{}, err
	}

	for _, post := range posts {
		resp.Posts = append(resp.Posts, getPostDetails(ctx, post))
	}

	return resp, nil
}

func getPostDetails(ctx context.Context, post tables.Post) shared.Post {
	likeCount, _ := postgres.Read[[]utils.JustId, tables.Like](ctx, map[string]interface{}{"Username": post.Username})
	commentCount, _ := postgres.Read[[]utils.JustId, tables.Comment](ctx, map[string]interface{}{"PostId": post.Id})

	return shared.Post{
		LikeCount:    uint(len(likeCount)),
		CommentCount: uint(len(commentCount)),
		Username:     post.Username,
		Description:  post.Description,
		CreatedDate:  post.CreatedDate,
	}
}
