FROM golang:alpine AS build-env
WORKDIR /news
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /news/go.mod
COPY go.sum /news/go.sum
RUN go mod download
COPY . /news
RUN CGO_ENABLED=0 GOOS=linux go build -o build/news ./news


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /news/build/news /
ENTRYPOINT ["/news"]
CMD ["up", "--grpc-port=80"]
