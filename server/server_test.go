package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/bluelinecoding/news"
	"github.com/lileio/lile/v2"
)

var s = NewsServer{}
var cli news.NewsClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		news.RegisterNewsServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = news.NewNewsClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
