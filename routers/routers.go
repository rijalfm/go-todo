package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rijalfm/go-todo/internal/domain/service"
)


// Setup Router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Todo API Endpoint
	todo := router.Group("api/v1/todo")
	{
		todo.GET("/", service.GetAllTodo)
		// todo.GET("/:id", service.GetTodoById)
		todo.POST("/", service.CreateTodo)
		// todo.DELETE("/:id", service.DeleteTodoById)
	}

	router.NoRoute(notFound)

	return router

}

// Handle unknown route (Not Found Page)
func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"error": "Page Not Found",
	})
}