package server

import (
	"testing"

	"context"
	"github.com/bluelinecoding/news"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	req := &news.GetRequest{}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
