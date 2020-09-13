package feeds

import (
	"github.com/mmcdole/gofeed"
)

func NewParser() *gofeed.Parser {
	fp := gofeed.NewParser()
	fp.RSSTranslator = NewCustomRSSTranslator()
	return fp
}

func GetFeed(url string, parser *gofeed.Parser) ([]*gofeed.Item, error) {
	feed, err := parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}
