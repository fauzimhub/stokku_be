package repository

import (
	"context"

	"database/sql"
	"github.com/fauzimhub/stokku/domain"
)

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}

func (cr *categoryRepository) Create(ctx context.Context, category *domain.Category) error {

	query := `
		INSERT INTO categories
		(name, description)
		VALUES ($1, $2)
	`
	_, err := cr.db.ExecContext(ctx, query,
		category.Name,
		category.Description,
	)

	return err
}

func (cr *categoryRepository) Fetch(ctx context.Context) ([]domain.Category, error) {

	query := "SELECT id, name, description FROM categories"

	rows, err := cr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		var category domain.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *categoryRepository) GetByName(ctx context.Context, name string) (domain.Category, error) {

	query := `
		SELECT id, name, description
		FROM categories
		WHERE name = $1
   `

	var category domain.Category
	err := cr.db.QueryRowContext(ctx, query, name).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Category{}, err
		}
		return domain.Category{}, err
	}

	return category, nil
}

func (cr *categoryRepository) GetByID(ctx context.Context, id int) (domain.Category, error) {

	query := `
		SELECT id, name, description
		FROM categories
		WHERE id = $1
   `

	var category domain.Category
	err := cr.db.QueryRowContext(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Category{}, err
		}
		return domain.Category{}, err
	}

	return category, nil
}
