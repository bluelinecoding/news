package server

import (
	"context"
	"sort"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/db"
	"github.com/bluelinecoding/news/feeds"
	"github.com/bluelinecoding/news/typeconv"
	"github.com/mmcdole/gofeed"
)

var PAGE_SIZE_DEFAULT = 20

func (s NewsServer) ListArticles(ctx context.Context, r *news.ListArticlesRequest) (*news.ListArticlesResponse, error) {
	err := r.Validate()
	if err != nil {
		validationErrors := news.ValidationErrToPBErrors(err)
		return nil, news.NewValidationError(validationErrors)
	}

	fetchingFeeds, err := getFeedUrls(ctx, r.FeedProviders, r.FeedCategories)
	if err != nil {
		return nil, err
	}

	if len(fetchingFeeds) == 0 {
		return &news.ListArticlesResponse{}, nil
	}

	var itemList []*gofeed.Item

	parser := feeds.NewParser()

	for _, u := range fetchingFeeds {
		items, err := feeds.GetFeed(u, parser)
		if err != nil {
			// Todo: Log that we couldn't get articles from the URL
			continue
		}
		itemList = append(itemList, items...)
	}

	sort.Slice(itemList, func(i, j int) bool {
		if itemList[i].PublishedParsed != nil && itemList[j].PublishedParsed != nil {
			return itemList[i].PublishedParsed.Unix() > itemList[j].PublishedParsed.Unix()
		}
		return true
	})

	itemLength := int32(len(itemList))
	startIndex, endIndex := getPagedIndexes(itemLength, r.PageSize, r.PageIndex)

	var result []*news.Article

	for _, i := range itemList[startIndex:endIndex] {
		pbItem := typeconv.GoFeedItemToPBArticle(i)
		// TODO: Add social media links
		result = append(result, pbItem)
	}

	return &news.ListArticlesResponse{
		Articles: result,
	}, nil
}

func getPagedIndexes(length int32, pageSize int32, pageIndex int32) (int32, int32) {
	if pageSize == 0 {
		pageSize = int32(PAGE_SIZE_DEFAULT)
	}

	startIndex := int32(0)
	endIndex := int32(0)

	if pageIndex > 0 {
		startIndex = pageSize * pageIndex
		endIndex = (pageSize * pageIndex) + pageSize
	}

	if pageIndex == 0 {
		endIndex = pageSize
	}

	if startIndex > length {
		startIndex = length
	}

	if endIndex > length {
		endIndex = length
	}

	return startIndex, endIndex
}

func getFeedUrls(ctx context.Context, providers []string, categories []string) ([]string, error) {
	var result []string

	dbFeeds, err := db.ListFeeds(ctx)
	if err != nil {
		return nil, err
	}

	hasProviders := len(providers) > 0
	hasCategories := len(categories) > 0

	if !hasProviders && !hasCategories {
		for _, f := range dbFeeds {
			result = append(result, f.Url)
		}
	} else {
		for _, p := range providers {
			for _, f := range dbFeeds {
				if f.Provider == p {
					if !hasCategories {
						result = append(result, f.Url)
						continue
					}

					for _, c := range categories {
						if f.Category == c {
							result = append(result, f.Url)
						}
					}
				}
			}
		}
	}

	sort.Strings(result)

	return result, nil
}
