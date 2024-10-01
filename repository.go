package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	postgrest_go "github.com/nedpals/postgrest-go/pkg"
)

type Repository struct {
	db *postgrest_go.Client
}

func NewRepository(db *postgrest_go.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(param NewTask) error {
	var result []Todo
	err := r.db.From("todos").Insert(param).Execute(&result)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Select(ctx context.Context) ([]Todo, error) {
	var results map[string]interface{}
	var todos Todo
	err := r.db.From("todos").Select("*").ExecuteWithContext(ctx, &results)
	if err != nil {
		return []Todo{}, err
	}

	mapstructure.Decode(results, &todos)

	return []Todo{todos}, nil
}
