package server

import (
	"errors"

	"context"
	"github.com/bluelinecoding/news"
)

func (s NewsServer) ListFeeds(ctx context.Context, r *news.ListFeedsRequest) (*news.ListFeedsResponse, error) {
	return nil, errors.New("not yet implemented")
}
