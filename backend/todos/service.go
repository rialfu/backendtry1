package todos

import (
	"net/http"

	"github.com/rialfu/backendtry1/model"
)

type Service interface {
	GetTodos() ([]model.Todos, int, error)
	CreateTodos(req DataRequest) (model.Todos, int, error)
	UpdateChecked(req InputRequest) (int, error)
	DeleteTodo(req InputRequestOnlyId) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetTodos() ([]model.Todos, int, error) {

	todos, err := s.repo.GetTodos()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return todos, http.StatusOK, nil
}

func (s *service) CreateTodos(req DataRequest) (model.Todos, int, error) {

	todo, err := s.repo.CreateTodos(req.Task)
	if err != nil {
		return model.Todos{}, http.StatusInternalServerError, err
	}

	return todo, http.StatusOK, nil
}
func (s *service) UpdateChecked(req InputRequest) (int, error) {

	check := req.Check != "0"
	err := s.repo.UpdateChecked(check, req.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
func (s *service) DeleteTodo(req InputRequestOnlyId) (int, error) {
	err := s.repo.DeleteTodo(req.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
