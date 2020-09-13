package server

import (
	"testing"

	"context"
	"github.com/bluelinecoding/news"
	"github.com/stretchr/testify/assert"
)

func TestListFeeds(t *testing.T) {
	ctx := context.Background()
	req := &news.ListFeedsRequest{}

	res, err := cli.ListFeeds(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
