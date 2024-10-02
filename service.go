package main

type TaskRepository interface {
	Insert(task NewTask) error
	Select() ([]Task, error)
}

type Service struct {
	taskRepo TaskRepository
}

func NewService(repo TaskRepository) *Service {
	return &Service{taskRepo: repo}
}

func (s *Service) CreateTask(task NewTask) error {
	return s.taskRepo.Insert(task)
}

func (s *Service) GetTasks() ([]Task, error) {
	return s.taskRepo.Select()
}
