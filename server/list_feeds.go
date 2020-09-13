package server

import (
	"context"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/bluelinecoding/news/typeconv"
)

func (s NewsServer) ListFeeds(ctx context.Context, r *news.ListFeedsRequest) (*news.ListFeedsResponse, error) {
	feeds, err := db.ListFeeds(ctx)
	if err != nil {
		return nil, err
	}

	if len(feeds) == 0 {
		return &news.ListFeedsResponse{
			Feeds: []*news.Feed{},
		}, nil
	}

	pb, err := typeconv.DBFeedsToPBFeeds(feeds)
	if err != nil {
		return nil, err
	}

	return &news.ListFeedsResponse{
		Feeds: pb,
	}, nil
}
