package usecase

import (
	"context"
	"time"

	"github.com/fauzimhub/stokku/domain"
)

type categoryUsecase struct {
	categoryRepository domain.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryUsecase(categoryRepository domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: categoryRepository,
		contextTimeout:     timeout,
	}
}

func (cu *categoryUsecase) Create(ctx context.Context, category *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Create(ctx, category)
}

func (cu *categoryUsecase) Fetch(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.Fetch(ctx)
}

func (cu *categoryUsecase) GetByName(ctx context.Context, name string) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.GetByName(ctx, name)
}

func (cu *categoryUsecase) GetByID(ctx context.Context, id int) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.categoryRepository.GetByID(ctx, id)
}
