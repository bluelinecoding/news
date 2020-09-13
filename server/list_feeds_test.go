package server

import (
	"testing"

	"context"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/stretchr/testify/assert"
)

func TestListFeeds(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	err1 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_1",
		Provider: "provider_1",
		Category: "category_1",
		Url:      "url_1",
	})

	if err1 != nil {
		t.Fatal(err1)
	}

	err2 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_2",
		Provider: "provider_2",
		Category: "category_1",
		Url:      "url_2",
	})

	if err2 != nil {
		t.Fatal(err2)
	}

	req := &news.ListFeedsRequest{}

	res, err := cli.ListFeeds(ctx, req)

	assert.NotNil(t, res)
	assert.Nil(t, err)
	assert.Len(t, res.Feeds, 2)

	assert.Equal(t, "feed_1", res.Feeds[0].Id)
}

func TestListFeeds_NoRecords(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	req := &news.ListFeedsRequest{}

	res, err := cli.ListFeeds(ctx, req)
	assert.NotNil(t, res)
	assert.Len(t, res.Feeds, 0)
	assert.Nil(t, err)
}
