package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/rialfu/backendtry1/api"
)

// var db []string

// type DataRequest struct {
// 	Text string `json:"text"`
// }

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

// type handleData struct {
// 	db []string
// }

// func (this *handleData) home(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": this.db,
// 	})
// }

//	func (this *handleData) postData(c *gin.Context) {
//		var data DataRequest
//		if err := c.ShouldBindJSON(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		this.db = append(this.db, data.Text)
//		c.JSON(http.StatusOK, gin.H{"message": "data berhasil terkirim", "data": data.Text})
//	}
//
// db := initial()
//
//	if err := db.AutoMigrate(&Todos{}); err != nil {
//		panic("failed to migrate database")
//	}
//
//	repo := repository{
//		db: db,
//	}
//
// // handle := new(handleData)
// // handle.db = make([]string, 0)
// r := gin.Default()
//
//	r.Use(cors.New(cors.Config{
//		AllowOrigins: []string{"*"},
//	}))
//
// r.GET("/", repo.handler)
// r.POST("/send", repo.postHandler)
// r.Run(":" + os.Getenv("PORT"))
func main() {

	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
}
