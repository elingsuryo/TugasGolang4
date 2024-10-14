package routes

import (
	controllers "tugas4/controllers"
	"tugas4/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    e.POST("/login", controllers.Login)

    todoGroup := e.Group("/todos")
    todoGroup.Use(middlewares.JWTMiddleware(), middlewares.EditorRole)
    todoGroup.POST("", controllers.CreateTodo)

    userGroup := e.Group("/users")
    userGroup.Use(middlewares.JWTMiddleware(), middlewares.AdminRole)
    userGroup.POST("", controllers.CreateUser)
}
