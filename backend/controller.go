package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type DataRequest struct {
	Task string `json:"task" binding:"required"`
}

func (r *repository) handler(c *gin.Context) {
	var todos []Todos
	res := r.db.Find(&todos)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}
func (r *repository) postHandler(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := Todos{
		Task: data.Task,
		Done: false,
	}

	res := r.db.Create(&todo)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": todo})
}
