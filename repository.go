package main

import (
	"encoding/json"
	postgrest_go "github.com/nedpals/postgrest-go/pkg"
)

type Repository struct {
	db *postgrest_go.Client
}

func NewRepository(db *postgrest_go.Client) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(param NewTask) error {
	var result []Task
	err := r.db.From("todos").Insert(param).Execute(&result)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Select() ([]Task, error) {
	var rows []map[string]interface{}
	err := r.db.From("todos").Select("*").Execute(&rows)
	if err != nil {
		return []Task{}, err
	}

	var results []Task
	for _, row := range rows {
		var todo Task
		err = ConvertToStruct(row, &todo)
		if err != nil {
			return []Task{}, nil
		}

		results = append(results, todo)
	}

	return results, nil
}

func ConvertToStruct(source map[string]interface{}, target interface{}) error {
	// Convert map to json string
	jsonStr, err := json.Marshal(source)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonStr, &target); err != nil {
		return err
	}
	return nil
}
