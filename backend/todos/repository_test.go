package todos

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/rialfu/backendtry1/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&model.Todos{})
	assert.NoError(t, err)

	return db
}

func TestCreateTodo(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	task := "task 1"

	todo, err := repo.CreateTodos(task)
	assert.NoError(t, err)
	assert.NotNil(t, todo)
}

func TestGetTodo(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	repo.CreateTodos("task 1")

	todos, err := repo.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(todos))

	repo.CreateTodos("task 2")

	todos, err = repo.GetTodos()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(todos))
}
