package repo

import (
	"fmt"
	"simple-template/server/model"
	"sync"
)

type TodoRepo interface {
	Create(*model.TodoData) error
	Update(model.TodoUnique, *model.TodoData) error
	Delete(model.TodoUnique) error
	FindMany(model.TodoFilter) ([]model.TodoData, error)
	FindUnique(model.TodoUnique) (*model.TodoData, error)
}

func NewTodoRepo() TodoRepo {
	return &TodoImpl{
		dataInterface: make([]model.TodoData, 0),
		dataCounter:   1,
	}
}

type TodoImpl struct {
	m             *sync.Mutex
	dataInterface []model.TodoData
	dataCounter   int
}

func (repo *TodoImpl) Create(toCreate *model.TodoData) error {
	repo.m.Lock()
	defer repo.m.Unlock()
	if toCreate.Id != 0 {
		for _, v := range repo.dataInterface {
			// in case of using input id for custom id value
			if v.Id == toCreate.Id {
				return fmt.Errorf("%v", ERROR_DUPLICATE_UNIQUE+" id")
			}
		}
	}

	for _, v := range repo.dataInterface {
		if repo.dataCounter == v.Id || repo.dataCounter == toCreate.Id {
			// update auto increment id in case of duplicate with counter
			repo.dataCounter++
		}
	}

	if toCreate.Id == 0 {
		toCreate.Id = repo.dataCounter
		repo.dataCounter++
	}

	repo.dataInterface = append(repo.dataInterface, *toCreate)

	return nil
}
func (repo TodoImpl) validateUnique(unique model.TodoUnique) bool {
	isIdEmpty := false
	if unique.Id == nil {
		isIdEmpty = true
	}
	if isIdEmpty /* && otherUniqueIsEmpty */ {
		return false
	}

	return true
}
func (repo *TodoImpl) Update(unique model.TodoUnique, toUpdate *model.TodoData) error {

	if repo.validateUnique(unique) {
		return fmt.Errorf("%v", ERROR_MISSING_UNIQUE)
	}

	repo.m.Lock()
	defer repo.m.Unlock()
	updateIndex := -1
	for i, v := range repo.dataInterface {
		if v.Id == *unique.Id {
			updateIndex = i
			break
		}
	}
	if updateIndex == -1 {
		return fmt.Errorf("%v", ERROR_NOT_FOUND)
	}

	repo.dataInterface[updateIndex].Name = toUpdate.Name

	return nil
}
func (repo *TodoImpl) Delete(unique model.TodoUnique) error {
	if !repo.validateUnique(unique) {
		return fmt.Errorf("%v", ERROR_MISSING_UNIQUE)
	}
	repo.m.Lock()
	defer repo.m.Unlock()
	deleteIndex := -1
	for i, v := range repo.dataInterface {
		if v.Id == *unique.Id {
			deleteIndex = i
			break
		}
	}
	if deleteIndex == -1 {
		return fmt.Errorf("%v", ERROR_NOT_FOUND)
	}

	return nil
}
func (repo TodoImpl) FindMany(filter model.TodoFilter) ([]model.TodoData, error) {
	data := []model.TodoData{}
	if filter.Id == nil && filter.Name == nil {
		return repo.dataInterface, nil
	}
	for _, v := range repo.dataInterface {
		if filter.Id != nil {
			data = append(data, v)
			break
		}
		if v.Name == *filter.Name {
			data = append(data, v)
		}
	}
	return data, nil
}
func (repo TodoImpl) FindUnique(unique model.TodoUnique) (*model.TodoData, error) {
	if repo.validateUnique(unique) {
		return nil, fmt.Errorf("%v", ERROR_MISSING_UNIQUE)
	}
	for _, v := range repo.dataInterface {
		if v.Id == *unique.Id {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("%v", ERROR_NOT_FOUND)
}
