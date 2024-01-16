package repo

import (
	"app/crawler/model"
	"context"
)

type WebPage interface {
	// GetOneByID(ctx context.Context, id int) (model.Student, error)
	// GetAll(ctx context.Context) ([]model.Student, error)
	InsertOne(ctx context.Context, c *model.WebPage) error
	// UpdateOne(ctx context.Context, id int, student *model.Student) error
	// DeleteOne(ctx context.Context, id int) error
}
