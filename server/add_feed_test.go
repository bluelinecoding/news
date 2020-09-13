package server

import (
	"testing"

	"context"
	"github.com/bluelinecoding/news"
	"github.com/stretchr/testify/assert"
)

func TestAddFeed(t *testing.T) {
	ctx := context.Background()
	req := &news.AddFeedRequest{}

	res, err := cli.AddFeed(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
