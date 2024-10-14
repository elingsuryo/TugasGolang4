package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
    return middleware.JWTWithConfig(middleware.JWTConfig{
        SigningKey: []byte("secret"),
    })
}

func AdminRole(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        if claims["role"] != "Admin" {
            return c.JSON(http.StatusForbidden, "Access denied")
        }
        return next(c)
    }
}

func EditorRole(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        user := c.Get("user").(*jwt.Token)
        claims := user.Claims.(jwt.MapClaims)
        if claims["role"] != "Editor" {
            return c.JSON(http.StatusForbidden, "Access denied")
        }
        return next(c)
    }
}
