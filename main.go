package main

import (
    "go-demo-server/products"
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, User API!")
	})

	repo := products.NewProductRepository()
	us := products.NewProductService(&repo)
	e.GET("/products", products.GetProductHandler(us))

	e.Logger.Fatal(e.Start("127.0.0.1:" + port))
}