package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "admin:test1234@tcp(todo-rds.cngkce8yenk4.ap-northeast-2.rds.amazonaws.com:3306)/gotodo?charset=utf8mb4&parseTime=True&loc=Local"

type Todo struct {
    gorm.Model
	Task      string    `gorm:"column:task"`
	Done      bool      `gorm:"column:done"`
}

func getTodos(db *gorm.DB) []Todo {
	var todos []Todo
	db.Find(&todos)
	fmt.Println(todos)
	return todos
}

func main() {

	// initialize router
	r := gin.Default()
	// load HTML templates
	r.LoadHTMLGlob("templates/*")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// connect to db
	if err != nil {
		log.Fatal(err)
	}
	// 테이블 생성
	db.AutoMigrate(&Todo{})

	r.GET("/", func(c *gin.Context) {
		todos := getTodos(db)
		fmt.Println(todos)
		c.HTML(http.StatusOK, "index.html", todos)
	})

	r.POST("/add-task", func(c *gin.Context) {
		// 제출된 form에서 task란 이름의 값 찾음
		task := c.PostForm("task")
		newTodo := Todo{
			Task: task,
			Done: false,
		}
		res := db.Create(&newTodo)

		if res.Error != nil {
			fmt.Println(res.Error)
			log.Fatal(res.Error)
		}
		fmt.Println(res)
		// 루트로 리다이렉트
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/delete-task", func(c *gin.Context) {
		taskId := c.PostForm("taskId")
		// id가 taskId와 동일한 Todo 삭제
		db.Delete(&Todo{}, taskId)
		c.Redirect(http.StatusFound, "/")
	})

	r.Run("localhost:8080")

}
