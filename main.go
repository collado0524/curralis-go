packege main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net-http"
	"curralis-go/models"	
)


var DB *gorm.DB

func initDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("curralis.db"),&gorm.Config{})
	if err !=nil {
		panic("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Todo{})
}

func main() {
	r:= gin.Default()

	initDatabase()

	r.GET("/todos", GetTodos)
	r.POST("/todos", CreateTodo)
	r.GET("/todos/:id", GetTodo)
	r.PUT("/todos/:id",UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)

	r.Run(":8080")
}

func GetTodos(c *gin.Context){
	var todos []models.Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK. todos)
}