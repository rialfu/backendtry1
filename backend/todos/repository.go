package todos

import (
	"github.com/rialfu/backendtry1/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetTodos() ([]model.Todos, error)
	CreateTodos(data string) (model.Todos, error)
	UpdateChecked(checked bool, id string) error
	DeleteTodo(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTodos() ([]model.Todos, error) {
	var todos []model.Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		return nil, res.Error
	}

	return todos, nil
}

func (r *repository) CreateTodos(task string) (model.Todos, error) {
	todo := model.Todos{
		Task: task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		return model.Todos{}, res.Error
	}

	return todo, nil
}
func (r *repository) UpdateChecked(checked bool, id string) error {
	todo := model.Todos{}
	res := r.db.Model(todo).Where("id = ?", id).Update("done", checked)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (r *repository) DeleteTodo(id string) error {
	todo := model.Todos{}
	res := r.db.Where("id = ?", id).Delete(&todo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
