package service

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rijalfm/go-todo/internal/domain/entity"
)

// Todos Variable
var Todos = []entity.Todo{}

// Get All Todo Handler
func GetTodos(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, Todos)

}

// Get Todo By Id Handler
func GetTodoById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for _, todo := range Todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})

}

// New Todo Handler
func SaveTodo(c *gin.Context) {

	newTodo := new(entity.Todo)

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	if len(Todos) > 0 {
		sort.Slice(Todos, func(i, j int) bool {
			return Todos[i].ID < Todos[j].ID
		})
		newTodo.ID = Todos[len(Todos)-1].ID + 1
	} else {
		newTodo.ID = 1
	}

	newTodo.CreatedAt = time.Now()

	if err := newTodo.Validate(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err["message"]})
		return
	}

	Todos = append(Todos, *newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)

}

// Delete Handler
func DeleteTodoById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	for idx, todo := range Todos {

		if todo.ID == id {
			Todos = removeTodo(Todos, idx)
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})

}

// Remove Slice Elmenet By Index Function
func removeTodo(t []entity.Todo, index int) []entity.Todo {

	newTodos := t
	return append(newTodos[:index], newTodos[index+1:]...)

}
