package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type State string

const (
	Available State = "available"
	Basket    State = "basket"
	Bough     State = "bought"
)

type Product struct {
	Name  string
	Price uint
	State State
	Id    uint
}

func find_product_by_id(id uint, arr []Product) int {
	for i := range arr {
		if arr[i].Id == id {
			return i
		}
	}
	return -1
}

var available = []Product{
	{"banana", 10, Available, 0},
	{"apple", 5, Available, 1},
	{"orange", 2, Available, 2}}

func main() {
	// available = []Product{{"banana", 10, 0}, {"apple", 5, 1}}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/products", func(c echo.Context) error {
		find_product_by_id(1, available)
		return c.JSON(http.StatusOK, available)
	})

	e.POST("/basket", func(c echo.Context) error {
		products := new([]Product)
		if err := c.Bind(products); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, products)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
