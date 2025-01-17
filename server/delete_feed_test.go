package server

import (
	"testing"

	"context"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/stretchr/testify/assert"
)

func TestDeleteFeed(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	err1 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_1",
		Provider: "provider_1",
		Category: "category_1",
		Url:      "http://feeds.bbci.co.uk/news/uk/rss.xml",
	})

	if err1 != nil {
		t.Fatal(err1)
	}

	feedBefore, err := db.GetFeed(ctx, "http://feeds.bbci.co.uk/news/uk/rss.xml")

	assert.NotNil(t, feedBefore)

	req := &news.DeleteFeedRequest{
		FeedId: "feed_1",
	}

	deleteRes, err := cli.DeleteFeed(ctx, req)

	assert.NotNil(t, deleteRes)
	assert.Nil(t, err)

	feedAfter, err := db.GetFeed(ctx, "http://feeds.bbci.co.uk/news/uk/rss.xml")

	assert.Nil(t, feedAfter)
}
