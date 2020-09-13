package server

import (
	"testing"

	"context"
	"github.com/bluelinecoding/news"
	"github.com/stretchr/testify/assert"
)

func TestListArticles(t *testing.T) {
	ctx := context.Background()
	req := &news.ListArticlesRequest{}

	res, err := cli.ListArticles(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
