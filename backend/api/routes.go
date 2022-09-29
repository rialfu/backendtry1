package api

import (
	"github.com/gin-contrib/cors"
	"github.com/rialfu/backendtry1/todos"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	todosRepo := todos.NewRepository(s.DB)
	todosService := todos.NewService(todosRepo)
	todosHandler := todos.NewHandler(todosService)

	s.Router.GET("/", todosHandler.GetTodos)
	s.Router.POST("/send", todosHandler.CreateTodo)
	s.Router.POST("/update-check", todosHandler.UpdateChecked)
	s.Router.POST("/delete-todo", todosHandler.DeleteTodo)
	s.Router.POST("/login", todosHandler.Login)
}
