# vi: ft=make
.PHONY: proto test

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . news.proto --lile-server_out=. --go_out=plugins=grpc,paths=source_relative:.

run:
	MIGRATIONS_DIR=migrations \
	go run news/main.go up

test: proto
	@go get github.com/rakyll/gotest
	MIGRATIONS_DIR=../migrations \
	gotest -v -p 1 ./... $${ARGS}