package server

import (
	"testing"

	"context"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/stretchr/testify/assert"
)

func TestAddFeed(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	req := &news.AddFeedRequest{
		Provider: "provider_1",
		Category: "category_1",
		Url:      "url_1",
	}

	res, err := cli.AddFeed(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	feedBefore, err := db.GetFeed(ctx, "url_1")
	assert.NotNil(t, feedBefore)
}

func TestAddFeed_NoDuplicates(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	req := &news.AddFeedRequest{
		Provider: "provider_1",
		Category: "category_1",
		Url:      "url_1",
	}

	res1, err := cli.AddFeed(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res1)

	res2, err := cli.AddFeed(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res2)

	feeds, err := db.ListFeeds(ctx)
	assert.Len(t, feeds, 1)
}
