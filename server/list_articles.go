package server

import (
	"errors"

	"context"
	"github.com/bluelinecoding/news"
)

func (s NewsServer) ListArticles(ctx context.Context, r *news.ListArticlesRequest) (*news.ListArticlesResponse, error) {
	return nil, errors.New("not yet implemented")
}
