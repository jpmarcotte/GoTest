package main

import (
	"github.com/labstack/echo/v4"
	"main/src/employee"
	"os"
)

func main() {
	e := echo.New()
	h, err := employee.NewHandler()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/employees", h.GetAllEmployees)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
