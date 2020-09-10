package server

import (
	"errors"

	"context"
	"github.com/bluelinecoding/news"
)

func (s NewsServer) Get(ctx context.Context, r *news.GetRequest) (*news.GetResponse, error) {
	return nil, errors.New("not yet implemented")
}
