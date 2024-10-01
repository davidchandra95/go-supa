package main

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(task NewTask) error {
	return s.repo.Insert(task)
}

func (s *Service) GetTasks() ([]Task, error) {
	return s.repo.Select()
}
