package todos

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetTodos(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		messageUnauthorized(c)
		return
	}
	token := c.Request.Header["Authorization"][0]
	if token != "09598333" {
		messageUnauthorized(c)
		return
	}

	todos, status, err := h.Service.GetTodos()
	if err != nil {
		log.Println("getTodos", err.Error())
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func (h *Handler) CreateTodo(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		messageUnauthorized(c)
		return
	}
	token := c.Request.Header["Authorization"][0]
	if token != "09598333" {
		messageUnauthorized(c)
		return
	}
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("CreateTodo", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, status, err := h.Service.CreateTodos(req)
	if err != nil {
		log.Println("CreateTodo", err.Error())
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"data":    res,
	})
}
func (h *Handler) UpdateChecked(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		messageUnauthorized(c)
		return
	}
	token := c.Request.Header["Authorization"][0]
	if token != "09598333" {
		messageUnauthorized(c)
		return
	}
	var req InputRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("UpdateChecked", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := h.Service.UpdateChecked(req)
	if err != nil {
		log.Println("UpdateChecked", err.Error())
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success update",
	})

}
func (h *Handler) DeleteTodo(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		messageUnauthorized(c)
		return
	}
	token := c.Request.Header["Authorization"][0]
	if token != "09598333" {
		messageUnauthorized(c)
		return
	}
	var req InputRequestOnlyId
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("DeleteTodo", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := h.Service.DeleteTodo(req)
	if err != nil {
		log.Println("DeleteTodo", err.Error())
		c.JSON(status, gin.H{
			"message": err.Error(),
			"d":       "d",
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success delete",
	})
}
func (h *Handler) Login(c *gin.Context) {
	var req InputLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("login", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("lewat", req.Password, req.Username, "l")
	if req.Username != "rema" || req.Password != "12345" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "your account false"})
		return
	}
	token := "09598333"
	c.JSON(http.StatusOK, gin.H{
		"message": "success login",
		"token":   token,
	})
}
func messageUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "you not have access",
	})
}
