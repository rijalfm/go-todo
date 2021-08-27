package service

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
	// guuid "github.com/google/uuid"
	"github.com/rijalfm/go-todo/internal/domain/entity"
)


// Create todo table if not already exists
func CreateTodoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&entity.Todo{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Todo table created")
	return nil
}

// Initialize DB connection
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}


// Get all todo
func GetAllTodo(c *gin.Context) {
	var todos []entity.Todo
	err := dbConnect.Model(&todos).Select()

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data": todos,
	})
}

// Create a new todo
func CreateTodo(c *gin.Context) {
	var todo entity.Todo
	c.BindJSON(&todo)
	// id := guuid.New().String()
	title := todo.Title
	description := todo.Description

	if err := todo.Validate(); err != nil {
		log.Printf("Error null value, Reason: %v\n", err["message"])
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"error": err["message"],
		})
		return
	}

	insertError := dbConnect.Insert(&entity.Todo{
		Title:       title,
		Description:      description,
		IsDone:     false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	if insertError != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"error": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created successfully",
	})
}
