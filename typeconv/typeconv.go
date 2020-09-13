package typeconv

import (
	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/golang/protobuf/ptypes"
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
