package domain

import (
	"context"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	Fetch(ctx context.Context) ([]Category, error)
	GetByName(ctx context.Context, name string) (Category, error)
	GetByID(ctx context.Context, id int) (Category, error)
}

type CategoryUsecase interface {
	Create(ctx context.Context, category *Category) error
	Fetch(ctx context.Context) ([]Category, error)
	GetByName(ctx context.Context, name string) (Category, error)
	GetByID(ctx context.Context, id int) (Category, error)
}
