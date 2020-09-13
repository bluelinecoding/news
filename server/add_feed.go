package server

import (
	"errors"

	"context"
	"github.com/bluelinecoding/news"
)

func (s NewsServer) AddFeed(ctx context.Context, r *news.AddFeedRequest) (*news.AddFeedResponse, error) {
	return nil, errors.New("not yet implemented")
}
