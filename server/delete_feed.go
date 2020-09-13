package server

import (
	"context"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
)

func (s NewsServer) DeleteFeed(ctx context.Context, r *news.DeleteFeedRequest) (*news.DeleteFeedResponse, error) {
	err := r.Validate()
	if err != nil {
		validationErrors := news.ValidationErrToPBErrors(err)
		return nil, news.NewValidationError(validationErrors)
	}

	err = db.DeleteFeed(ctx, r.FeedId)
	if err != nil {
		return nil, err
	}

	return &news.DeleteFeedResponse{}, nil
}
