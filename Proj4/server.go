package main

import (
	"fmt"
	"net/http"
	"project/database"
	"project/handlers"

	"github.com/labstack/echo/v4"
)

func main() {

	// database.db_init()
	database.Init()
	fmt.Print(database.Database)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

	e.GET("/customer", func(c echo.Context) error {
		return handlers.GetCustomers(c)
	})
	e.POST("/customer", func(c echo.Context) error {
		return handlers.CreateCustomer(c)
	})

	e.GET("/worker", func(c echo.Context) error {
		return handlers.GetWorkers(c)
	})
	e.POST("/worker", func(c echo.Context) error {
		return handlers.CreateWorker(c)
	})

	e.GET("/shift", func(c echo.Context) error {
		return handlers.GetShifts(c)
	})
	e.POST("/shift", func(c echo.Context) error {
		return handlers.CreateShift(c)
	})

	e.GET("/product", func(c echo.Context) error {
		return handlers.GetProducts(c)
	})
	e.POST("/product", func(c echo.Context) error {
		return handlers.CreateProduct(c)
	})

	e.GET("/purchase", func(c echo.Context) error {
		return handlers.GetPurchases(c)
	})
	e.POST("/purchase", func(c echo.Context) error {
		return handlers.CreatePurchase(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
