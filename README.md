# News API

## First use
- First run `docker-compose stop && docker-compose rm && docker-compose up -d`. This will ensure a clean setup of the database
- To run `docker-compose up --force-recreate` which will start the docker container for the DB

## Tests  
- To run unit tests do `make test`

## Start service 
- Do `make run` which will start the gRPC service
- There are no feeds initally, so requests to add a feed are needed
- Use `grpcurl` (https://github.com/fullstorydev/grpcurl) to run requests e.g.

    - `grpcurl -d '{"provider":"SKY", "category":"uk", "url":"http://feeds.skynews.com/feeds/rss/uk.xml"}' -proto ./news.proto -plaintext  0.0.0.0:8000 news.News.AddFeed`

    - `grpcurl -d '{"feed_providers":["SKY", "BBC"], "page_index": 1, "page_size": 5}' -proto ./news.proto -plaintext  0.0.0.0:8000 news.News.ListArticles`

    - `grpcurl -d '{}' -proto ./news.proto -plaintext  0.0.0.0:8000 news.News.ListFeeds`

    - `grpcurl -d '{"feed_id":"[FEED_ID]"}' -proto ./news.proto -plaintext  0.0.0.0:8000 news.News.DeleteFeed`

## TODO
    - Social media proper links
    - Logging
    - Spans for tracing
    - Additional unit testing
