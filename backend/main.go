package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

// var db []string

type DataRequest struct {
	Text string `json:"text"`
}

// func postHandler(c *gin.Context) {
// 	var data DataRequest
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db = append(db, data.Text)
// 	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": data.Text})
// }
// func handler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": db,
// 	})
// }

type handleData struct {
	db []string
}

func (this *handleData) home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": this.db,
	})
}
func (this *handleData) postData(c *gin.Context) {
	var data DataRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	this.db = append(this.db, data.Text)
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": data.Text})
}

func main() {
	handle := new(handleData)
	handle.db = make([]string, 0)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))
	r.GET("/", handle.home)
	r.POST("/send", handle.postData)
	r.Run(":" + os.Getenv("PORT"))
}
