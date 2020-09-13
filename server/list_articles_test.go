package server

import (
	"testing"

	"context"

	"github.com/bluelinecoding/news/db"
	"github.com/stretchr/testify/assert"
)

// TODO: Write tests with mocked news feeds

func createFeedData(ctx context.Context, t *testing.T) {
	// Out of order so we can be sure sorting works in asserts
	err1 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_4",
		Provider: "provider_2",
		Category: "category_2",
		Url:      "url_4",
	})

	if err1 != nil {
		t.Fatal(err1)
	}

	err2 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_2",
		Provider: "provider_1",
		Category: "category_2",
		Url:      "url_2",
	})

	if err2 != nil {
		t.Fatal(err2)
	}

	err3 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_3",
		Provider: "provider_2",
		Category: "category_1",
		Url:      "url_3",
	})

	if err3 != nil {
		t.Fatal(err3)
	}

	err4 := db.AddFeed(ctx, &db.Feed{
		ID:       "feed_1",
		Provider: "provider_1",
		Category: "category_1",
		Url:      "url_1",
	})

	if err4 != nil {
		t.Fatal(err4)
	}
}

func TestGetFeedUrlsWithNoProviderAndNoCategory(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	createFeedData(ctx, t)

	feedUrls, err := getFeedUrls(ctx, []string{}, []string{})
	assert.Nil(t, err)
	assert.Len(t, feedUrls, 4)

	assert.Equal(t, "url_1", feedUrls[0])
	assert.Equal(t, "url_2", feedUrls[1])
	assert.Equal(t, "url_3", feedUrls[2])
	assert.Equal(t, "url_4", feedUrls[3])
}

func TestGetFeedUrlsWithProvider(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	createFeedData(ctx, t)

	feedUrls, err := getFeedUrls(ctx, []string{"provider_1"}, []string{})
	assert.Nil(t, err)
	assert.Len(t, feedUrls, 2)

	assert.Equal(t, "url_1", feedUrls[0])
	assert.Equal(t, "url_2", feedUrls[1])
}

func TestGetFeedUrlsWithMultipleProviders(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	createFeedData(ctx, t)

	feedUrls, err := getFeedUrls(ctx, []string{"provider_1", "provider_2"}, []string{})
	assert.Nil(t, err)
	assert.Len(t, feedUrls, 4)

	assert.Equal(t, "url_1", feedUrls[0])
	assert.Equal(t, "url_2", feedUrls[1])
	assert.Equal(t, "url_3", feedUrls[2])
	assert.Equal(t, "url_4", feedUrls[3])
}

func TestGetFeedUrlsWithProviderAndCategory(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	createFeedData(ctx, t)

	feedUrls, err := getFeedUrls(ctx, []string{"provider_1"}, []string{"category_1"})
	assert.Nil(t, err)
	assert.Len(t, feedUrls, 1)

	assert.Equal(t, "url_1", feedUrls[0])
}

func TestGetFeedUrlsWithMultipleProvidersAndCategory(t *testing.T) {
	db.ResetDB()
	ctx := context.Background()

	createFeedData(ctx, t)

	feedUrls, err := getFeedUrls(ctx, []string{"provider_1", "provider_2"}, []string{"category_1"})
	assert.Nil(t, err)
	assert.Len(t, feedUrls, 2)

	assert.Equal(t, "url_1", feedUrls[0])
	assert.Equal(t, "url_3", feedUrls[1])
}

func TestGetPagedIndexes(t *testing.T) {
	start, end := getPagedIndexes(0, 10, 1)

	assert.Equal(t, int32(0), start)
	assert.Equal(t, int32(0), end)

	start, end = getPagedIndexes(3, 10, 1)

	assert.Equal(t, int32(3), start)
	assert.Equal(t, int32(3), end)

	start, end = getPagedIndexes(21, 10, 1)

	assert.Equal(t, int32(10), start)
	assert.Equal(t, int32(20), end)

	start, end = getPagedIndexes(21, 10, 2)

	assert.Equal(t, int32(20), start)
	assert.Equal(t, int32(21), end)

	start, end = getPagedIndexes(21, 2, 0)

	assert.Equal(t, int32(0), start)
	assert.Equal(t, int32(2), end)
}
