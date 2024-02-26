package main

import (
	"fmt"
	"http-test/lib"
	"http-test/routes"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	lib.InitDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:id", routes.GetUser)
	e.POST("/user", routes.CreateUser)
	e.POST("/user/:id", routes.UpdateUser)
	e.DELETE("/user/:id", routes.DeleteUser)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
