package db

import (
	"context"

	"github.com/jinzhu/gorm"
)

func GetFeed(ctx context.Context, url string) (*Feed, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return nil, err
	}

	var result Feed

	err = db.Where("url = ?", url).First(&result).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func AddFeed(ctx context.Context, f *Feed) error {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	return db.Create(f).Error
}

func DeleteFeed(ctx context.Context, id string) error {
	db, err := GetDB(ctx)
	if err != nil {
		return err
	}

	return db.Where("id = ?", id).Delete(&Feed{}).Error
}

func ListFeeds(ctx context.Context) ([]*Feed, error) {
	db, err := GetDB(ctx)
	if err != nil {
		return nil, err
	}

	var feeds []*Feed

	err = db.Find(&feeds).Error
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
