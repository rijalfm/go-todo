package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rijalfm/go-todo/internal/domain/service"
)


// Setup Router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Todo API Endpoint
	todo := router.Group("api/v1/todo")
	{
		todo.GET("/", service.GetTodos)
		todo.GET("/:id", service.GetTodoById)
		todo.POST("/", service.SaveTodo)
		todo.DELETE("/:id", service.DeleteTodoById)
	}

	return router

}
