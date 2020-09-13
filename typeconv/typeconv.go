package typeconv

import (
	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/golang/protobuf/ptypes"
	"github.com/mmcdole/gofeed"
)

func DBFeedsToPBFeeds(dbFeeds []*db.Feed) ([]*news.Feed, error) {
	feeds := []*news.Feed{}

	for _, f := range dbFeeds {
		pbFeed, err := DBFeedToPBFeed(f)
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, pbFeed)
	}

	return feeds, nil
}

func DBFeedToPBFeed(dbFeed *db.Feed) (*news.Feed, error) {
	createTime, err := ptypes.TimestampProto(dbFeed.CreatedAt)
	if err != nil {
		return nil, err
	}

	feed := &news.Feed{
		Id:         dbFeed.ID,
		Category:   dbFeed.Category,
		Provider:   dbFeed.Provider,
		Url:        dbFeed.Url,
		CreateTime: createTime,
	}

	return feed, nil
}

func GoFeedItemToPBArticle(item *gofeed.Item) *news.Article {
	article := &news.Article{
		Id:          item.GUID,
		Title:       item.Title,
		Description: item.Description,
	}

	if item.PublishedParsed != nil {
		publishedTime, err := ptypes.TimestampProto(*item.PublishedParsed)
		if err != nil {
			// Log error that we couldn't parse?
		}
		article.PublishedTime = publishedTime
	}

	if item.Image != nil {
		image := &news.Image{
			Url: item.Image.URL,
		}

		article.Image = image
	}

	return article
}
