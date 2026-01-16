package router

import (
	"github.com/gin-gonic/gin"
	"github.com/todo-app/internal/handler"
	"github.com/todo-app/internal/middleware"
	"github.com/todo-app/internal/service"
	"time"
)

type RouterSetup struct {
	Engine *gin.Engine
}

func Setup(svcGroup *service.ServiceGroup) *RouterSetup {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, "pong")
	})

	r.GET("/echo", func(c *gin.Context) {
		healthData := map[string]interface{}{
			"path": "echo",
			"time": time.Now().Format(time.DateTime),
		}
		c.JSON(200, healthData)
	})

	userHandler := handler.NewUserHandler(svcGroup.UserService)
	// example of github signin integration
	r.GET("/login", userHandler.GithubLogin)
	r.GET("/auth/github/callback", userHandler.GithubCallback)

	todoHandler := handler.NewTodoHandler(svcGroup.TodoService)
	api := r.Group("/api")
	api.Use(middleware.JwtAuth())
	{
		api.POST("/todo", todoHandler.CreateTodo)
		api.GET("/todo", todoHandler.ListTodos)
		api.DELETE("/todo/:id", todoHandler.DeleteTodo)
		api.PUT("/todo/:id/complete", todoHandler.CompleteTodo)
	}

	return &RouterSetup{
		Engine: r,
	}
}
