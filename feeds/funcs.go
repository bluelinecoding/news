package feeds

import (
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/patrickmn/go-cache"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func setCache(key string, v interface{}) bool {
	Cache.Set(key, v, cache.DefaultExpiration)
	return true
}

func NewParser() *gofeed.Parser {
	fp := gofeed.NewParser()
	fp.RSSTranslator = NewCustomRSSTranslator()
	return fp
}

func GetFeed(url string, parser *gofeed.Parser) ([]*gofeed.Item, error) {
	cached, found := Cache.Get(url)
	if found {
		items := cached.([]*gofeed.Item)
		return items, nil
	}

	feed, err := parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	// TODO: find out the correct cache time, based on feed time to live details
	setCache(url, feed.Items)

	return feed.Items, nil
}
