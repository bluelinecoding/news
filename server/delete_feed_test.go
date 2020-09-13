package server

import (
	"testing"

	"context"
	"github.com/bluelinecoding/news"
	"github.com/stretchr/testify/assert"
)

func TestDeleteFeed(t *testing.T) {
	ctx := context.Background()
	req := &news.DeleteFeedRequest{}

	res, err := cli.DeleteFeed(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
