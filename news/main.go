package main

import (
	_ "net/http/pprof"

	"github.com/bluelinecoding/news"
	"github.com/bluelinecoding/news/news/cmd"
	"github.com/bluelinecoding/news/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile/v2"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.NewsServer{}

	lile.Name("news")
	lile.Server(func(g *grpc.Server) {
		news.RegisterNewsServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
