package main

import (
	"go-dirty-workspace/typicode"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Route => handler
	e.GET("/photos", func(c echo.Context) error {
		p := []typicode.Photo{}
		t := typicode.New("/photos")
		err := t.Get(&p)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, p)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
