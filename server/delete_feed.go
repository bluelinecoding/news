package server

import (
	"errors"

	"context"
	"github.com/bluelinecoding/news"
)

func (s NewsServer) DeleteFeed(ctx context.Context, r *news.DeleteFeedRequest) (*news.DeleteFeedResponse, error) {
	return nil, errors.New("not yet implemented")
}
