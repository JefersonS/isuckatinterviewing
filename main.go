package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type Question struct {
	gorm.Model
	Id       int
	Question string
	Answer   string
}

func main() {
	dsn := "host=localhost user=postgres password=5234 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

	db.AutoMigrate(&Question{})

	router := gin.Default()
	router.Static("/templates", "./templates")
	router.LoadHTMLGlob("templates/*")
		// home page
	router.GET("/", func(c *gin.Context) {
		var questions []Question
		db.Find(&questions)
		c.HTML(http.StatusOK, "home.html", gin.H{
			"questions": questions,
		})
	})

	router.GET("/questions/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createQuestion.html", gin.H{})
	})

	router.GET("/questions/cancel", func(c *gin.Context) {
		c.HTML(http.StatusOK, "containerHeader.html", gin.H{})
	})

	router.POST("/questions/create", func(c *gin.Context) {
		fmt.Println("Create question")
		fmt.Println(c.PostForm("question"))
		question := c.PostForm("question")
		answer := c.PostForm("answer")
		result := db.Create(&Question{Question: question, Answer: answer})
		if result.Error != nil {
			fmt.Println(result.Error)
			c.HTML(http.StatusOK, "notCreated.html", gin.H{})
		} else {
			c.HTML(http.StatusOK, "successfullyCreated.html", gin.H{})
		}
	})

	router.GET("/questions/search", func(c *gin.Context) {
		searchQuestion := c.Query("question")
		fmt.Println(searchQuestion)
		var questions []Question
		result := db.Where("question LIKE ?", "%"+searchQuestion+"%").Find(&questions)
		if result.Error != nil {
			fmt.Println(result)
			c.HTML(http.StatusOK, "notCreated.html", gin.H{})
		} else {
			c.HTML(http.StatusOK, "questionsAndAnswers.html", gin.H{
				"questions": questions,
			})
		}
	})

	router.Run()
}
