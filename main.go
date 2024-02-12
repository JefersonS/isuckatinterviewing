package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Question struct {
	Id       string
	Question string
	Answer   string
}

func main() {
	questions := []Question{
		{
			Id: "1", Question: "What is your name?", Answer: "My name is Gopher",
		},
		{
			Id: "2", Question: "What is your age?", Answer: "I am 10 years old",
		},
		{
			Id: "3", Question: "What is your hobby?", Answer: "My hobby is coding",
		},
		{
			Id: "4", Question: "What is your favorite programming language?", Answer: "My favorite programming language is Go",
		},
	}
	router := gin.Default()
	router.Static("/templates", "./templates")
	router.LoadHTMLGlob("templates/*")
	router.GET("/sayIt", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"questions": questions,
		})
	})

	router.GET("/questions/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createQuestion.html", gin.H{})
	})

	router.GET("/questions/cancel", func(c *gin.Context) {
		c.HTML(http.StatusOK, "container-header.html", gin.H{})
	})

	router.Run()
}
