package server

import (
	"context"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
)

func (s NewsServer) AddFeed(ctx context.Context, r *news.AddFeedRequest) (*news.AddFeedResponse, error) {
	err := r.Validate()
	if err != nil {
		validationErrors := news.ValidationErrToPBErrors(err)
		return nil, news.NewValidationError(validationErrors)
	}

	feed, err := db.GetFeed(ctx, r.Url)
	if err != nil {
		return nil, err
	}

	if feed != nil {
		// Already exists
		return &news.AddFeedResponse{}, nil
	}

	err = db.AddFeed(ctx, &db.Feed{
		Provider: r.Provider,
		Category: r.Category,
		Url:      r.Url,
	})
	if err != nil {
		return nil, err
	}

	return &news.AddFeedResponse{}, nil
}
