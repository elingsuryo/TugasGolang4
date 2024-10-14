package main

import (
	"tugas4/models"
	"tugas4/routes"

	"github.com/labstack/echo/v4"
)




func main() {
    e := echo.New()

    // Inisialisasi Database
    models.InitDB()

    // Inisialisasi Routes
    routes.InitRoutes(e)

    // Jalankan server
    e.Logger.Fatal(e.Start(":8080"))
}
