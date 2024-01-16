package mysql

import (
	"app/crawler/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type WebPageRepository struct {
	db *gorm.DB
}

func (w WebPageRepository) InsertOne(ctx context.Context, webpage *model.WebPage) error {
	if err := w.db.Create(&webpage).Error; err != nil {
		return fmt.Errorf("insert infomation error: %w", err)
	}
	return nil
}

var instance WebPageRepository

func NewWebPageRepository(db *gorm.DB) WebPageRepository { // Corrected the type name
	if instance.db == nil {
		instance.db = db
	}
	return instance
}
