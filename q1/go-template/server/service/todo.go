package service

import (
	"simple-template/server/model"
	"simple-template/server/repo"
)

type TodoService struct {
	todoRepo repo.TodoRepo
}

func NewTodoService(todoRepo repo.TodoRepo) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
	}
}

func (service *TodoService) Create(input *model.TodoData) error {
	err := service.todoRepo.Create(input)
	if err != nil {
		return err
	}

	return nil
}

func (service *TodoService) Delete(unique model.TodoUnique) error {
	err := service.todoRepo.Delete(unique)

	if err != nil {
		return err
	}
	return nil
}

func (service *TodoService) FindAll() []model.TodoData {
	data, _ := service.todoRepo.FindMany(model.TodoFilter{})

	return data
}

func (service *TodoService) FindUnique(unique *model.TodoUnique) (*model.TodoData, error) {
	data, err := service.todoRepo.FindUnique(*unique)
	if err != nil {
		return nil, err
	}

	return data, nil
}
