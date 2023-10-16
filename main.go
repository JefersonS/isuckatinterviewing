package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	// "strconv"
	// "encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type Response struct {
	Message string `json:"message"`
}

func allQuestions() ([]Question, error) {
    var questions []Question

    rows, err := db.Query("SELECT * FROM questions")
    if err != nil {
        return nil, fmt.Errorf("allQuestions: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var question Question
        if err := rows.Scan(&question.ID, &question.Question, &question.Answer, &question.ItSucks, &question.YouSuck); err != nil {
            return nil, fmt.Errorf("allQuestionsScan: %v", err)
        }
        questions = append(questions, question)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("allQuestions: %v", err)
    }
    return questions, nil
}

func insertNewQuestion(question Question) (int64, error) {
	result, err := db.Exec("INSERT INTO questions (question, answer, itSucks, youSuck) VALUES (?, ?, ?, ?)", question.Question, question.Answer, question.ItSucks, question.YouSuck)
    if err != nil {
        return 0, fmt.Errorf("insertNewQuestion: %v", err)
    }
	id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("insertNewQuestion: %v", err)
    }
    return id, nil
}

var db *sql.DB

func sayHello(c *gin.Context) {
	var response = Response{
		Message: "I SUCK at interviewing!",
	}
	c.IndentedJSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.Static("/templates", "./templates")
	router.LoadHTMLGlob("templates/*")
	router.GET("/sayIt", sayHello)
	router.GET("/clicked", func(c *gin.Context) {
		var question = GetQuestion()
		c.HTML(http.StatusOK, "file.tmpl", gin.H{
			"question": question.Question,
			"answer":   question.Answer,
		})
	})

	router.GET("/index", func(c *gin.Context) {
		var questions, err = allQuestions()
		if err != nil {
			fmt.Println("errrrrrr: %v", err)
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"questions": questions,
		})
	})

	router.POST("/index/create-question", func(c *gin.Context) {
		newQuestion := Question {
			Question: c.PostForm("question"),
			Answer: c.PostForm("answer"),
		}

		// itSucks, _ := strconv.Atoi(params.Get("itSucks"))
		// youSuck, _ := strconv.Atoi(params.Get("youSuck"))

		// var newQuestion = Question {
		// 	Question: params.Get("question"),
		// 	Answer: params.Get("answer"),
		// 	ItSucks: int32(itSucks),
		// 	YouSuck: int32(youSuck),
		// }
		_, insertErr := insertNewQuestion(newQuestion)
		if insertErr != nil {
			fmt.Println("errrrrrr: %v", insertErr)
		}

		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	cfg := mysql.Config{
		User:   "root",
		Passwd: "Jebs#5234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "questions",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router.Run("localhost:8080")
}
