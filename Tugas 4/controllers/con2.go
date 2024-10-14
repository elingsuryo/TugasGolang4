package controllers

import (
	"net/http"
	"tugas4/models"

	"github.com/labstack/echo/v4"
)

func CreateTodo(c echo.Context) error {
    task := c.FormValue("task")
    todo := models.Todo{Task: task, Status: "Pending"}

    if err := models.DB.Create(&todo).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, todo)
}
