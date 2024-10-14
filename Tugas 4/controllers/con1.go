package controllers

import (
	"net/http"
	"time"
	"tugas4/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    var user models.User
    models.DB.Where("username = ? AND password = ?", username, password).First(&user)

    if user.ID == 0 {
        return c.JSON(http.StatusUnauthorized, "Invalid credentials")
    }

    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = user.Username
    claims["role"] = user.Role
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Error generating token")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": tokenString,
    })
}
