package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Route => handler
	e.GET("/photos", func(c echo.Context) error {
		pht := []byte(`{id: 1}`)
		return c.JSONBlob(http.StatusOK, pht)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
