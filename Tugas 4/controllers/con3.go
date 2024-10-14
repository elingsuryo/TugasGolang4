package controllers

import (
	"net/http"
	"tugas4/models"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")
    role := c.FormValue("role")

    user := models.User{Username: username, Password: password, Role: role}
    if err := models.DB.Create(&user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, user)
}
