package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/segmentio/ksuid"
)

type Feed struct {
	ID        string
	Provider  string
	Category  string
	Url       string
	CreatedAt time.Time
}

func CreateFeedID() string {
	return "feed_" + ksuid.New().String()
}

func (feed *Feed) BeforeCreate(scope *gorm.Scope) error {
	if feed.ID == "" {
		feed.ID = CreateFeedID()
		scope.SetColumn("ID", feed.ID)
	}

	return nil
}
